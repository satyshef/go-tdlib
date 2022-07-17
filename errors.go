package tdlib

import (
	"encoding/json"
	"strings"
)

const (
	//Error codes
	// clients errors
	ErrorCodeSystem                = 400
	ErrorCodeNoAccess              = 403
	ErrorCodeNotFound              = 404
	ErrorCodeUsernameNotOccupied   = 405
	ErrorCodeUsernameInvalid       = 406
	ErrorCodeUserPrivacyRestricted = 407
	ErrorCodeUserNotFound          = 408
	ErrorCodeMemberNotFound        = 409
	ErrorCodeUserNotMutual         = 410
	ErrorCodeUserTooMuch           = 411
	ErrorCodeChatNotFound          = 412

	// server errors
	ErrorCodeManyRequests        = 501
	ErrorCodeUserBannedInChannel = 502
	ErrorCodePhoneInvalid        = 503
	ErrorCodePassInvalid         = 504
	ErrorCodeStopped             = 505
	ErrorCodeFloodLock           = 506
	ErrorCodeTimeout             = 507
	ErrorCodeClose               = 508
	ErrorCodeNotInit             = 509
	ErrorCodeLogout              = 510
	ErrorCodeAborted             = 511
	ErrorCodePhoneBanned         = 512
	ErrorCodeUserKickedFromChat  = 513
	ErrorCodeAuthKeyDublicated   = 514
)

// Error An object of this type can be returned on every function call, in case of an error
type Error struct {
	tdCommon
	Code    int    `json:"code"`    // Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
	Message string `json:"message"` // Error message; subject to future changes
}

// MessageType return the string telegram-type of Error
func (error *Error) MessageType() string {
	return "error"
}

func (error *Error) Error() string {
	result, _ := json.Marshal(error)
	return string(result)
}

// NewError creates a new Error
// @param code Error code; subject to future changes. If the error code is 406, the error message must not be processed in any way and must not be displayed to the user
// @param message Error message; subject to future changes
func NewError(code int, message, extra string) *Error {
	errorTemp := Error{
		tdCommon: tdCommon{Type: "error", Extra: extra},
		Code:     code,
		Message:  message,
	}

	return &errorTemp
}

//Конвертируем ответ в ошибку
func responseToError(response UpdateMsg, update UpdateData) *Error {

	var e *Error
	if err := json.Unmarshal(response.Raw, &e); err != nil {
		return NewError(ErrorCodeSystem, "SYSTEM_JSON", err.Error())
	}
	//В качестве типа ошибки устанавливаем название запроса
	e.Type = update["@type"].(string)

	if strings.Contains(e.Message, "Too Many Requests") {
		e.Code = ErrorCodeManyRequests
		return e
	}

	//Переназначаем коды ошибок телеграм на свои
	switch e.Message {
	case "USERNAME_NOT_OCCUPIED":
		e.Code = ErrorCodeUsernameNotOccupied
	case "USER_PRIVACY_RESTRICTED":
		e.Code = ErrorCodeUserPrivacyRestricted
	case "USER_NOT_MUTUAL_CONTACT":
		e.Code = ErrorCodeUserNotMutual
	case "USER_CHANNELS_TOO_MUCH":
		e.Code = ErrorCodeUserTooMuch
	case "AUTH_KEY_DUPLICATED":
		e.Code = ErrorCodeAuthKeyDublicated
	case "PHONE_NUMBER_BANNED":
		e.Code = ErrorCodePhoneBanned
	case "PASSWORD_HASH_INVALID":
		e.Code = ErrorCodePassInvalid
	case "PHONE_NUMBER_INVALID":
		e.Code = ErrorCodePhoneInvalid
	case "USER_DEACTIVATED_BAN",
		"Unauthorized":
		e.Code = ErrorCodeLogout
	case "USERNAME_INVALID":
		e.Code = ErrorCodeUsernameInvalid
	case "USER_BANNED_IN_CHANNEL":
		e.Code = ErrorCodeUserBannedInChannel
	case "PEER_FLOOD":
		e.Code = ErrorCodeFloodLock
	case "Have no write access to the chat":
		e.Code = ErrorCodeNoAccess
	case "Have no rights to send a message":
		e.Code = ErrorCodeNoAccess
	case "USER_KICKED":
		e.Code = ErrorCodeUserKickedFromChat
	case "Can't return to kicked from chat":
		e.Code = ErrorCodeUserKickedFromChat
	case "User not found":
		e.Code = ErrorCodeUserNotFound
	case "Member not found":
		e.Code = ErrorCodeMemberNotFound
	case "Chat not found":
		e.Code = ErrorCodeChatNotFound
	case "Request aborted":
		e.Code = ErrorCodeAborted
	}

	return e
}
