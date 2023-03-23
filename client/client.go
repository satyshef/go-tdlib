package client

//#cgo linux CFLAGS: -I/usr/local/include
//#cgo darwin CFLAGS: -I/usr/local/include
//#cgo windows CFLAGS: -IC:/src/td -IC:/src/td/build
//#cgo linux LDFLAGS: -L/usr/local/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
//#cgo darwin LDFLAGS: -L/usr/local/lib -L/usr/local/opt/openssl/lib -ltdjson_static -ltdjson_private -ltdclient -ltdcore -ltdapi -ltdactor -ltddb -ltdsqlite -ltdnet -ltdutils -lstdc++ -lssl -lcrypto -ldl -lz -lm
//#cgo windows LDFLAGS: -LC:/src/td/build/Debug -ltdjson
//#include <stdlib.h>
//#include <td/telegram/td_json_client.h>
//#include <td/telegram/td_log.h>
import "C"

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
	"unsafe"

	"github.com/satyshef/go-tdlib/tdlib"
)

// EventHandler ....
type EventHandler func(event *SystemEvent) *tdlib.Error

//type ClientStatus string

/*
const (
	StatusInit     ClientStatus = "init"
	StatusRun      ClientStatus = "run"
	StatusStopping ClientStatus = "stopping"
	StatusStopped  ClientStatus = "stopped"
)
*/

// Client is the Telegram TdLib client
type Client struct {
	Client          unsafe.Pointer
	Config          Config
	receiversUpdate []chan *tdlib.UpdateMsg
	eventHandlers   []EventHandler
	waiters         map[string]chan tdlib.UpdateMsg
	receiverLock    *sync.Mutex
	waitersLock     *sync.RWMutex
	isRun           bool
	//Status          ClientStatus
	//IsStopped       bool
	//StopWork        chan bool
}

// Config holds tdlibParameters
type Config struct {
	APIID              string // Application identifier for Telegram API access, which can be obtained at https://my.telegram.org   --- must be non-empty..
	APIHash            string // Application identifier hash for Telegram API access, which can be obtained at https://my.telegram.org  --- must be non-empty..
	SystemLanguageCode string // IETF language tag of the user's operating system language; must be non-empty.
	DeviceModel        string // Model of the device the application is being run on; must be non-empty.
	SystemVersion      string // Version of the operating system the application is being run on; must be non-empty.
	ApplicationVersion string // Application version; must be non-empty.
	// Optional fields
	UseTestDataCenter      bool   // if set to true, the Telegram test environment will be used instead of the production environment.
	DatabaseDirectory      string // The path to the directory for the persistent database; if empty, the current working directory will be used.
	FileDirectory          string // The path to the directory for storing files; if empty, database_directory will be used.
	UseFileDatabase        bool   // If set to true, information about downloaded and uploaded files will be saved between application restarts.
	UseChatInfoDatabase    bool   // If set to true, the library will maintain a cache of users, basic groups, supergroups, channels and secret chats. Implies use_file_database.
	UseMessageDatabase     bool   // If set to true, the library will maintain a cache of chats and messages. Implies use_chat_info_database.
	UseSecretChats         bool   // If set to true, support for secret chats will be enabled.
	EnableStorageOptimizer bool   // If set to true, old files will automatically be deleted.
	IgnoreFileNames        bool   // If set to true, original file names will be ignored. Otherwise, downloaded files will be saved under names as close as possible to the original name.
}

// NewClient Creates a new instance of TDLib.
func NewClient(config Config) *Client {
	// Seed rand with time
	rand.Seed(time.Now().UnixNano())
	//client := Client{Client: C.td_json_client_create()}
	client := Client{}
	client.Config = config
	client.receiverLock = &sync.Mutex{}
	client.waitersLock = &sync.RWMutex{}
	client.waiters = make(map[string]chan tdlib.UpdateMsg)
	//client.Status = StatusInit
	//client.StopWork = make(chan bool)

	/*
		if err := client.run(); err != nil {
			return nil, err
		}
	*/
	return &client
}

// Start
func (client *Client) Run() error {
	//client.Destroy()
	//client.Status = StatusInit
	if client.IsRun() {
		return fmt.Errorf("%s", "Client is running")
	}
	client.isRun = true
	runGetUpdates(client)
	time.Sleep(time.Second * 1)
	client.sendTdLibParams()
	for state, err := client.Authorize(); state == nil; {
		fmt.Printf("ERROR %#v\n", err)
		if err != nil {
			client.Stop()
			return err
		}
		//time.Sleep(time.Second * 1)
	}
	//client.Status = StatusRun
	return nil
}

func (client *Client) Stop() {
	// WORK VERSION 1
	/*
		defer func() {
			client.Client = nil
		}()
	*/

	if !client.IsRun() {
		fmt.Println("DEBUG: Stop in Stop")
		return
	}

	// WORK VERSION 1
	/*
		client.isRun = false
		for client.WaitersLen() != 0 || client.Client != nil {
			//fmt.Println("Waiters count : ", client.WaitersLen())
			time.Sleep(time.Second * 1)
		}

	*/
	client.isRun = false
	timeout := 20
	for client.WaitersLen() != 0 && timeout > 0 {
		//fmt.Println("Waiters count : ", client.WaitersLen())
		time.Sleep(time.Second * 1)
		timeout -= 1
	}
}

func (client *Client) IsRun() bool {
	if client.isRun {
		return true
	}
	return false
}

func (client *Client) WaitersLen() int {
	return len(client.waiters)
}

/*
func (client *Client) IsStopped() bool {
	if client.Status != StatusStopped {
		return false
	}
	return true
}
*/

func runGetUpdates(client *Client) {
	go func() {
		client.Client = C.td_json_client_create()
		for client.IsRun() {
			updateBytes := client.Receive(10)
			if len(updateBytes) == 0 {
				continue
			}
			var updateData tdlib.UpdateData
			json.Unmarshal(updateBytes, &updateData)
			// does new update has @extra field?
			if extra, hasExtra := updateData["@extra"].(string); hasExtra {
				client.waitersLock.RLock()
				waiter, found := client.waiters[extra]
				client.waitersLock.RUnlock()
				// trying to load update with this salt
				if found {
					// found? send it to waiter channel
					waiter <- tdlib.UpdateMsg{Data: updateData, Raw: updateBytes}
					// trying to prevent memory leak
					close(waiter)
				}
			} else {
				client.publishUpdate(&tdlib.UpdateMsg{Data: updateData, Raw: updateBytes})
			}
		}
		C.td_json_client_destroy(client.Client)
		client.Client = nil
	}()
}

// AddEventHandler ....
func (client *Client) AddEventHandler(event EventHandler) {
	//блокируется при рестарте
	client.eventHandlers = append(client.eventHandlers, event)
}

func (client *Client) ResetEventHandlers() {
	client.eventHandlers = nil
}

func (client *Client) PublishEvent(event *SystemEvent) error {
	/*
		if client == nil {
			return NewError(ErrorCodeSystem, "CLIENT_DESTROED", "Client is destroed")
		}

		if event == (SystemEvent{}) {
			return NewError(ErrorCodeSystem, "CLIENT_WRONG_DATA", "No required parameters. Event is empty")
		}
	*/
	if !client.IsRun() {
		return tdlib.NewError(ErrorCodeStopped, "CLIENT_STOPPED", "Publish Event Error : Client stopped")
	}
	// TODO: test disable mutex
	/*
		client.receiverLock.Lock()
		defer client.receiverLock.Unlock()
	*/
	// Отправляем событие подписавшимся обработчикам
	for _, h := range client.eventHandlers {
		err := h(event)
		if err != nil {
			return err
		}
	}

	return nil

}

// GetRawUpdatesChannel creates a general channel that fetches every update comming from tdlib
func (client *Client) GetRawUpdatesChannel(capacity int) chan *tdlib.UpdateMsg {
	c := make(chan *tdlib.UpdateMsg, capacity)
	client.receiversUpdate = append(client.receiversUpdate, c)
	return c
}

// публикуем обновления
func (client *Client) publishUpdate(update *tdlib.UpdateMsg) {
	if update.Data == nil || update.Raw == nil {
		return
	}
	for _, c := range client.receiversUpdate {
		c <- update
	}
}

// DestroyInstance Destroys the TDLib client instance.
// After this is called the client instance shouldn't be used anymore.
func (client *Client) destroyInstance11() {
	if client.Client == nil {
		return
	}

	state, err := client.Authorize()
	if err != nil {
		e := err.(*tdlib.Error)
		if e.Code != ErrorCodeStopped {
			fmt.Printf("Destroy client error %#v\n", err)
			return
		}
	}
	fmt.Printf("STATE #%v\n\n", state)
	if state == nil || state.GetAuthorizationStateEnum() != tdlib.AuthorizationStateClosedType && state.GetAuthorizationStateEnum() != tdlib.AuthorizationStateClosingType {
		fmt.Println("Destroy !!!")
		C.td_json_client_destroy(client.Client)
		//time.Sleep(time.Second * 1)
		//client.Client = nil
	}
}

// Send Sends request to the TDLib client.
// You can provide string or tdlib.UpdateData.
func (client *Client) Send(jsonQuery interface{}, force bool) {

	if !client.IsRun() && !force {
		fmt.Printf("Query not send, client stopped : %s\n", jsonQuery)
		return
	}

	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case tdlib.UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(tdlib.UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(client.Client, query)

}

// Receive Receives incoming updates and request responses from the TDLib client.
// You can provide string or tdlib.UpdateData.
func (client *Client) Receive(timeout float64) []byte {

	if !client.IsRun() {
		fmt.Printf("Incoming updates not receives, client stopped\n")
		return nil
	}

	result := C.td_json_client_receive(client.Client, C.double(timeout))

	return []byte(C.GoString(result))
}

// Execute Synchronously executes TDLib request.
// Only a few requests can be executed synchronously.
func (client *Client) Execute(jsonQuery interface{}) tdlib.UpdateMsg {
	/*
		if IsStopped {
			fmt.Printf("Request not execute, client stopped\n")
			return tdlib.UpdateMsg{}
		}
	*/
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case tdlib.UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(tdlib.UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	result := C.td_json_client_execute(client.Client, query)

	var update tdlib.UpdateData
	json.Unmarshal([]byte(C.GoString(result)), &update)
	return tdlib.UpdateMsg{Data: update, Raw: []byte(C.GoString(result))}
}

// SetFilePath Sets the path to the file to where the internal TDLib log will be written.
// By default TDLib writes logs to stderr or an OS specific log.
// Use this method to write the log to a file instead.
func SetFilePath(path string) {
	bytes, _ := json.Marshal(tdlib.UpdateData{
		"@type": "setLogStream",
		"log_stream": tdlib.UpdateData{
			"@type":         "logStreamFile",
			"path":          path,
			"max_file_size": 10485760,
		},
	})

	query := C.CString(string(bytes))
	C.td_json_client_execute(nil, query)
	C.free(unsafe.Pointer(query))
}

// SetLogVerbosityLevel Sets the verbosity level of the internal logging of TDLib.
// By default the TDLib uses a verbosity level of 5 for logging.
func SetLogVerbosityLevel(level int) {
	bytes, _ := json.Marshal(tdlib.UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": level,
	})

	query := C.CString(string(bytes))
	C.td_json_client_execute(nil, query)
	C.free(unsafe.Pointer(query))
}

// SendAndCatch Sends request to the TDLib client and catches the result in updates channel.
// You can provide string or tdlib.UpdateData.
func (client *Client) SendAndCatch(jsonQuery interface{}) (tdlib.UpdateMsg, error) {

	var update tdlib.UpdateData

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case tdlib.UpdateData:
		update = jsonQuery.(tdlib.UpdateData)
	}
	//Публикуем запрос
	ev := UpdateToEvent(update)
	//fmt.Printf("UP %#v\n\n", ev)
	//var err error
	err := client.PublishEvent(ev)
	if err != nil {
		return tdlib.UpdateMsg{}, err
	}
	// letters for generating random string
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// generate random string for @extra field
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	randomString := string(b)
	// set @extra field
	update["@extra"] = randomString
	// create waiter chan and save it in Waiters
	waiter := make(chan tdlib.UpdateMsg, 1)

	client.waitersLock.Lock()
	client.waiters[randomString] = waiter
	client.waitersLock.Unlock()
	// send it through already implemented method
	if update["@type"] == "close" {
		client.Send(update, true)
	} else {
		client.Send(update, false)
	}

	select {
	// wait response from main loop in NewClient()
	case response := <-waiter:
		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()

		// if response is error
		if response.Data["@type"].(string) == "error" {
			respErr := responseToError(response, update)
			ev := ErrorToEvent(respErr)
			if err := client.PublishEvent(ev); err != nil {
				return tdlib.UpdateMsg{}, err
			}
			return tdlib.UpdateMsg{}, respErr
		}

		ev := ResponseToEvent(response, update)
		//ev := responseToEvent(response, update)
		//fmt.Printf("Response Type %s\n\n", ev2.DataType())

		err := client.PublishEvent(ev)
		return response, err

		// or timeout
	case <-time.After(15 * time.Second):
		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()
		if update["@type"].(string) == "close" {
			err := tdlib.NewError(ErrorCodeClose, "CLIENT_CLOSE", "Client closed")
			return tdlib.UpdateMsg{}, err
		}
		e := tdlib.NewError(ErrorCodeTimeout, "CLIENT_TIMEOUT", "Receive answer timeout")
		ev := ErrorToEvent(e)
		//e.Extra = randomString
		client.PublishEvent(ev)
		return tdlib.UpdateMsg{}, e
	}
}

// Authorize is used to authorize the users
func (client *Client) Authorize() (tdlib.AuthorizationState, error) {
	state, err := client.GetAuthorizationState()
	if err != nil {
		return nil, err
	}
	if state.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitEncryptionKeyType {
		ok, err := client.CheckDatabaseEncryptionKey(nil)
		if ok == nil || err != nil {
			return nil, err
		}
	} else if state.GetAuthorizationStateEnum() == tdlib.AuthorizationStateWaitTdlibParametersType {
		client.sendTdLibParams()
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

func (client *Client) sendTdLibParams() {

	client.Send(tdlib.UpdateData{
		"@type":                    "setTdlibParameters",
		"use_test_dc":              client.Config.UseTestDataCenter,
		"database_directory":       client.Config.DatabaseDirectory,
		"files_directory":          client.Config.FileDirectory,
		"use_file_database":        client.Config.UseFileDatabase,
		"use_chat_info_database":   client.Config.UseChatInfoDatabase,
		"use_message_database":     client.Config.UseMessageDatabase,
		"use_secret_chats":         client.Config.UseSecretChats,
		"api_id":                   client.Config.APIID,
		"api_hash":                 client.Config.APIHash,
		"system_language_code":     client.Config.SystemLanguageCode,
		"device_model":             client.Config.DeviceModel,
		"system_version":           client.Config.SystemVersion,
		"application_version":      client.Config.ApplicationVersion,
		"enable_storage_optimizer": client.Config.EnableStorageOptimizer,
		"ignore_file_names":        client.Config.IgnoreFileNames,
	}, true)
}

// SendPhoneNumber sends phone number to tdlib
func (client *Client) SendPhoneNumber(phoneNumber string) (tdlib.AuthorizationState, error) {
	phoneNumberConfig := tdlib.PhoneNumberAuthenticationSettings{AllowFlashCall: false, AllowMissedCall: false, IsCurrentPhoneNumber: true, AllowSmsRetrieverAPI: true}
	_, err := client.SetAuthenticationPhoneNumber(phoneNumber, &phoneNumberConfig)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthCode sends auth code to tdlib
func (client *Client) SendAuthCode(code string) (tdlib.AuthorizationState, error) {
	_, err := client.CheckAuthenticationCode(code)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthPassword sends two-step verification password (user defined)to tdlib
func (client *Client) SendAuthPassword(password string) (tdlib.AuthorizationState, error) {
	_, err := client.CheckAuthenticationPassword(password)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

//convert tdlib.UpdateData to SystemEvent
func UpdateToEvent(update tdlib.UpdateData) *SystemEvent {
	name := update["@type"].(string)
	data := make(map[string]interface{})
	for key, val := range update {
		if key == "@extra" || key == "@type" {
			continue
		}
		data[key] = val
	}
	return &SystemEvent{
		Type: EventTypeRequest,
		Name: name,
		Data: data,
	}
}

func ResponseToEvent(response tdlib.UpdateMsg, update tdlib.UpdateData) *SystemEvent {
	/*
		response.Data["@extra"] = update["@type"]
		return response.Data
	*/
	data := make(map[string]interface{})
	for key, val := range response.Data {
		if key == "@extra" {
			continue
		}
		data[key] = val
	}
	//result
	r := &SystemEvent{
		Type: EventTypeResponse,
		Name: update["@type"].(string),
		Data: data,
	}
	/*
		switch data["@type"] {
		case "ok":
			r.Data = "OK"
		case "text":
			r.Data = data["text"]
		default:
			r.Data = data
		}
	*/
	return r
}

func ErrorToEvent(err *tdlib.Error) *SystemEvent {
	err.Extra = ""
	return &SystemEvent{
		Type: EventTypeError,
		Name: err.Type,
		//DataType: "tdlib_err",
		Data: *err,
	}

}
