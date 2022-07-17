// AUTOGENERATED - DO NOT EDIT

package tdlib

import (
	"encoding/json"
)

// ChatEvent Represents a chat event
type ChatEvent struct {
	tdCommon
	ID     JSONInt64       `json:"id"`      // Chat event identifier
	Date   int32           `json:"date"`    // Point in time (Unix timestamp) when the event happened
	UserID int64           `json:"user_id"` // Identifier of the user who performed the action that triggered the event
	Action ChatEventAction `json:"action"`  // Action performed by the user
}

// MessageType return the string telegram-type of ChatEvent
func (chatEvent *ChatEvent) MessageType() string {
	return "chatEvent"
}

// NewChatEvent creates a new ChatEvent
//
// @param iD Chat event identifier
// @param date Point in time (Unix timestamp) when the event happened
// @param userID Identifier of the user who performed the action that triggered the event
// @param action Action performed by the user
func NewChatEvent(iD JSONInt64, date int32, userID int64, action ChatEventAction) *ChatEvent {
	chatEventTemp := ChatEvent{
		tdCommon: tdCommon{Type: "chatEvent"},
		ID:       iD,
		Date:     date,
		UserID:   userID,
		Action:   action,
	}

	return &chatEventTemp
}

// UnmarshalJSON unmarshal to json
func (chatEvent *ChatEvent) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	tempObj := struct {
		tdCommon
		ID     JSONInt64 `json:"id"`      // Chat event identifier
		Date   int32     `json:"date"`    // Point in time (Unix timestamp) when the event happened
		UserID int64     `json:"user_id"` // Identifier of the user who performed the action that triggered the event

	}{}
	err = json.Unmarshal(b, &tempObj)
	if err != nil {
		return err
	}

	chatEvent.tdCommon = tempObj.tdCommon
	chatEvent.ID = tempObj.ID
	chatEvent.Date = tempObj.Date
	chatEvent.UserID = tempObj.UserID

	fieldAction, _ := unmarshalChatEventAction(objMap["action"])
	chatEvent.Action = fieldAction

	return nil
}
