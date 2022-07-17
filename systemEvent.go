package tdlib

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type SystemEvent struct {
	Type EventType   `json:"type"`
	Name string      `json:"name"`
	Time int64       `json:"time"`
	Data interface{} `json:"data"`
}

type EventType string

const (
	EventTypeRequest  EventType = "request"
	EventTypeResponse EventType = "response"
	EventTypeUpdate   EventType = "update"
	EventTypeError    EventType = "error"
)

// New SystemEvent
func NewEvent(eventType EventType, eventName string, eventTime int64, eventData string) *SystemEvent {
	if eventTime == 0 {
		eventTime = time.Now().UnixNano()
	}
	return &SystemEvent{
		Type: eventType,
		Name: eventName,
		Data: eventData,
		Time: eventTime,
	}
}

func (ev *SystemEvent) DataType() string {
	t := fmt.Sprintf("%T", ev.Data)
	t = strings.Trim(t, " {}")
	return t
}

func (ev *SystemEvent) DataJSON() string {
	result, _ := json.Marshal(ev.Data)
	return string(result)
}
