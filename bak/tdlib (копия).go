package tdlib

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
	"strings"

	"sync"
	"time"
	"unsafe"
)

// UpdateData alias for use in UpdateMsg
type UpdateData map[string]interface{}

// UpdateMsg is used to unmarshal recieved json strings into
type UpdateMsg struct {
	Data UpdateData
	Raw  []byte
}

// EventHandler ....
//type EventHandler func(eventType string, update *UpdateMsg) *Error
type EventHandler func(eventType string, update interface{}) *Error

const (
	EventTypeRequest  = "request"
	EventTypeResponse = "response"
	EventTypeUpdate   = "update"
	EventTypeError    = "error"
)

const (
	//Error codes
	ErrorCodeNotFound        = 101
	ErrorCodeSystem          = 102
	ErrorCodeTimeout         = 103
	ErrorCodeClose           = 104
	ErrorCodeStopped         = 105
	ErrorCodeNotInit         = 106
	ErrorCodeLogout          = 107
	ErrorCodeAborted         = 108
	ErrorCodePhoneBanned     = 109
	ErrorCodePhoneInvalid    = 110
	ErrorCodePassInvalid     = 111
	ErrorCodeManyRequests    = 112
	ErrorCodeUsernameInvalid = 113
	ErrorCodeFloodLock       = 114
	ErrorCodeNoAccess        = 115
)

// Client is the Telegram TdLib client
type Client struct {
	Client          unsafe.Pointer
	Config          Config
	receiversUpdate []chan *UpdateMsg
	eventHandlers   []EventHandler
	waiters         map[string]chan UpdateMsg
	receiverLock    *sync.Mutex
	waitersLock     *sync.RWMutex
	IsStopped       bool
	StopWork        chan bool
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
	client.waiters = make(map[string]chan UpdateMsg)
	client.StopWork = make(chan bool)

	/*
		if err := client.run(); err != nil {
			return nil, err
		}
	*/
	return &client
}

// Start
func (client *Client) Start() error {
	//client.Destroy()
	client.IsStopped = false
	client.Client = C.td_json_client_create()
	go func() {
		for {

			select {
			case <-client.StopWork:
				goto Stop

			default:
				// get update
				updateBytes := client.Receive(1)
				if len(updateBytes) == 0 {
					break
				}

				//fmt.Printf("UP %#v\n", updateBytes)
				var updateData UpdateData
				json.Unmarshal(updateBytes, &updateData)

				// does new update has @extra field?
				if extra, hasExtra := updateData["@extra"].(string); hasExtra {
					client.waitersLock.RLock()
					waiter, found := client.waiters[extra]
					client.waitersLock.RUnlock()

					// trying to load update with this salt
					if found {
						// found? send it to waiter channel
						waiter <- UpdateMsg{Data: updateData, Raw: updateBytes}

						// trying to prevent memory leak
						close(waiter)
					}
				} else {
					//fmt.Printf("Publish Update : %#v\n", updateData)
					client.publishUpdate(&UpdateMsg{Data: updateData, Raw: updateBytes})
				}

				//fmt.Printf("UP %#v\n", updateData)
			}
		}

	Stop:
		client.IsStopped = true
		//fmt.Println("Client stoped")
	}()

	client.sendTdLibParams()
	_, err := client.Authorize()
	return err

}

func (client *Client) Stop() {

	if client.IsStopped || client.StopWork == nil {
		return
	}

	//TEST!!!!
	/*
		for client.WaitersLen() != 0 {
			fmt.Println("Waiters count : ", client.WaitersLen())
			time.Sleep(time.Second * 1)
		}
	*/
	client.StopWork <- true
	for !client.IsStopped {
		time.Sleep(1 * time.Second)
	}
	//time.Sleep(1 * time.Second)

	//client.Close()
	client.DestroyInstance()
	//fmt.Println("Client stoped")
	/*
		for {
			l := client.WaitersLen()
			fmt.Println("Waiters", l)
			if l != 0 {
				time.Sleep(time.Second * 1)
			} else {
				break
			}
		}
	*/
	//client.ResetEventHandlers()
	//time.Sleep(1 * time.Second)

	//client.DestroyInstance()
}

func (client *Client) WaitersLen() int {

	return len(client.waiters)

	/*
		length := 0
		client.waiters.Range(func(_, _ interface{}) bool {
			length++

			return true
		})

		return length
	*/
}

// AddEventHandler ....
//--------------------------------------------------------------------
func (client *Client) AddEventHandler(event EventHandler) {
	//блокируется при рестарте
	//client.receiverLock.Lock()
	//defer client.receiverLock.Unlock()
	client.eventHandlers = append(client.eventHandlers, event)

}

func (client *Client) ResetEventHandlers() {
	client.eventHandlers = nil
}

func (client *Client) PublishEvent(eventType string, update interface{}) error {

	if client == nil {
		return NewError(ErrorCodeSystem, "CLIENT_DESTROED", "Client is destroed")
	}

	if update == nil {
		return NewError(ErrorCodeSystem, "CLIENT_WRONG_DATA", "No required parameters. Event is empty")
	}

	if client.IsStopped {
		return NewError(ErrorCodeStopped, "CLIENT_STOPPED", "Publish Event Error : Client stopped")
	}

	client.receiverLock.Lock()
	defer client.receiverLock.Unlock()
	// Отправляем событие подписавшимся обработчикам
	for _, u := range client.eventHandlers {
		err := u(eventType, update)
		if err != nil {
			return err
		}
	}

	return nil

}

// GetRawUpdatesChannel creates a general channel that fetches every update comming from tdlib
func (client *Client) GetRawUpdatesChannel(capacity int) chan *UpdateMsg {
	c := make(chan *UpdateMsg, capacity)
	client.receiversUpdate = append(client.receiversUpdate, c)
	return c
}

// публикуем обновления
func (client *Client) publishUpdate(update *UpdateMsg) {
	if update.Data == nil || update.Raw == nil {
		return
	}
	for _, c := range client.receiversUpdate {
		c <- update
	}
}

// DestroyInstance Destroys the TDLib client instance.
// After this is called the client instance shouldn't be used anymore.
func (client *Client) DestroyInstance() {
	C.td_json_client_destroy(client.Client)
}

// Send Sends request to the TDLib client.
// You can provide string or UpdateData.
func (client *Client) Send(jsonQuery interface{}, force bool) {

	if client.IsStopped && !force {
		fmt.Printf("Query not send, client stopped : %s\n", jsonQuery)
		return
	}

	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	C.td_json_client_send(client.Client, query)

}

// Receive Receives incoming updates and request responses from the TDLib client.
// You can provide string or UpdateData.
func (client *Client) Receive(timeout float64) []byte {
	/*
		if IsStopped {
			fmt.Printf("Incoming updates not receives, client stopped\n")
			return nil
		}
	*/
	result := C.td_json_client_receive(client.Client, C.double(timeout))

	return []byte(C.GoString(result))
}

// Execute Synchronously executes TDLib request.
// Only a few requests can be executed synchronously.
func (client *Client) Execute(jsonQuery interface{}) UpdateMsg {
	/*
		if IsStopped {
			fmt.Printf("Request not execute, client stopped\n")
			return UpdateMsg{}
		}
	*/
	var query *C.char

	switch jsonQuery.(type) {
	case string:
		query = C.CString(jsonQuery.(string))
	case UpdateData:
		jsonBytes, _ := json.Marshal(jsonQuery.(UpdateData))
		query = C.CString(string(jsonBytes))
	}

	defer C.free(unsafe.Pointer(query))
	result := C.td_json_client_execute(client.Client, query)

	var update UpdateData
	json.Unmarshal([]byte(C.GoString(result)), &update)
	return UpdateMsg{Data: update, Raw: []byte(C.GoString(result))}
}

// SetFilePath Sets the path to the file to where the internal TDLib log will be written.
// By default TDLib writes logs to stderr or an OS specific log.
// Use this method to write the log to a file instead.
func SetFilePath(path string) {
	bytes, _ := json.Marshal(UpdateData{
		"@type": "setLogStream",
		"log_stream": UpdateData{
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
	bytes, _ := json.Marshal(UpdateData{
		"@type":               "setLogVerbosityLevel",
		"new_verbosity_level": level,
	})

	query := C.CString(string(bytes))
	C.td_json_client_execute(nil, query)
	C.free(unsafe.Pointer(query))
}

// SendAndCatch Sends request to the TDLib client and catches the result in updates channel.
// You can provide string or UpdateData.
func (client *Client) SendAndCatch(jsonQuery interface{}) (UpdateMsg, error) {

	var update UpdateData

	switch jsonQuery.(type) {
	case string:
		// unmarshal JSON into map, we don't have @extra field, if user don't set it
		json.Unmarshal([]byte(jsonQuery.(string)), &update)
	case UpdateData:
		update = jsonQuery.(UpdateData)
	}

	//Публикуем запрос
	raw, _ := json.Marshal(jsonQuery)
	err := client.PublishEvent(EventTypeRequest, UpdateMsg{Data: update, Raw: raw})
	if err != nil {
		return UpdateMsg{}, err
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
	waiter := make(chan UpdateMsg, 1)

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
		/*
			if IsStopped {
				return UpdateMsg{}, fmt.Errorf("Request not execute, client stopped\n")
			}
		*/

		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()

		if e := ResponceToError(response); e != nil {
			//В качестве типа ошибки устанавливаем название запроса
			e.Type = update["@type"].(string)

			//client.publishError(e)
			if err := client.PublishEvent(EventTypeError, *e); err != nil {
				return UpdateMsg{}, err
			}
			return UpdateMsg{}, e
		}
		/*
			fmt.Println("====================================")
			fmt.Printf("UPDATE : %#v\n", update)
			fmt.Println("------------------------------------")
			fmt.Printf("RESPONSE : %#v\n", response.Data)
			fmt.Println("====================================")
		*/
		//публикуем результат выполнения запроса
		//update["@extra"] = response.Data["@type"]
		response.Data["@extra"] = update["@type"]
		err := client.PublishEvent(EventTypeResponse, UpdateMsg{Data: response.Data, Raw: raw})
		//fmt.Printf("RESP %s\n\n", err)
		return response, err

		// or timeout
	case <-time.After(10 * time.Second):
		/*
			if IsStopped {
				return UpdateMsg{}, fmt.Errorf("Request not execute, client stopped\n")
			}
		*/
		client.waitersLock.Lock()
		delete(client.waiters, randomString)
		client.waitersLock.Unlock()

		if update["@type"].(string) == "close" {
			err := NewError(ErrorCodeClose, "CLIENT_CLOSE", "Client closed")
			return UpdateMsg{}, err
		}
		// 501 - таймаут
		e := NewError(ErrorCodeTimeout, "CLIENT_TIMEOUT", "Receive answer timeout")
		e.Extra = randomString
		client.PublishEvent(EventTypeError, *e)
		//fmt.Printf("EEEE %#v\n", e)
		return UpdateMsg{}, e
	}
}

// Authorize is used to authorize the users
func (client *Client) Authorize() (AuthorizationState, error) {

	state, err := client.GetAuthorizationState()
	if err != nil {
		return nil, err
	}

	if state.GetAuthorizationStateEnum() == AuthorizationStateWaitEncryptionKeyType {
		ok, err := client.CheckDatabaseEncryptionKey(nil)
		if ok == nil || err != nil {
			return nil, err
		}
	} else if state.GetAuthorizationStateEnum() == AuthorizationStateWaitTdlibParametersType {
		client.sendTdLibParams()
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

func (client *Client) sendTdLibParams() {

	client.Send(UpdateData{
		"@type": "setTdlibParameters",
		"parameters": UpdateData{
			"@type":                    "tdlibParameters",
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
		},
	}, true)
}

// SendPhoneNumber sends phone number to tdlib
func (client *Client) SendPhoneNumber(phoneNumber string) (AuthorizationState, error) {
	phoneNumberConfig := PhoneNumberAuthenticationSettings{AllowFlashCall: false, IsCurrentPhoneNumber: false, AllowSmsRetrieverAPI: false}
	_, err := client.SetAuthenticationPhoneNumber(phoneNumber, &phoneNumberConfig)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthCode sends auth code to tdlib
func (client *Client) SendAuthCode(code string) (AuthorizationState, error) {
	_, err := client.CheckAuthenticationCode(code)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

// SendAuthPassword sends two-step verification password (user defined)to tdlib
func (client *Client) SendAuthPassword(password string) (AuthorizationState, error) {
	_, err := client.CheckAuthenticationPassword(password)

	if err != nil {
		return nil, err
	}

	authState, err := client.GetAuthorizationState()
	return authState, err
}

//Конвертируем ответ в ошибку
func ResponceToError(response UpdateMsg) *Error {

	if response.Data["@type"].(string) != "error" {
		return nil
	}

	var e *Error
	json.Unmarshal(response.Raw, &e)

	if strings.Contains(e.Message, "Too Many Requests") {
		e.Code = ErrorCodeManyRequests
		return e
	}

	//Переназначаем коды ошибок телеграм на свои
	switch e.Message {
	case "PHONE_NUMBER_BANNED":
		e.Code = ErrorCodePhoneBanned
	case "PASSWORD_HASH_INVALID":
		e.Code = ErrorCodePassInvalid
	case "PHONE_NUMBER_INVALID":
		e.Code = ErrorCodePhoneInvalid
	case "Unauthorized":
		e.Code = ErrorCodeLogout
	case "USERNAME_INVALID":
		e.Code = ErrorCodeUsernameInvalid
	case "PEER_FLOOD":
		e.Code = ErrorCodeFloodLock
	case "Have no write access to the chat":
		e.Code = ErrorCodeNoAccess
	}

	return e
}
