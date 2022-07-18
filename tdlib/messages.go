// AUTOGENERATED - DO NOT EDIT

package tdlib

// Messages Contains a list of messages
type Messages struct {
	tdCommon
	TotalCount int32     `json:"total_count"` // Approximate total count of messages found
	Messages   []Message `json:"messages"`    // List of messages; messages may be null
}

// MessageType return the string telegram-type of Messages
func (messages *Messages) MessageType() string {
	return "messages"
}

// NewMessages creates a new Messages
//
// @param totalCount Approximate total count of messages found
// @param messages List of messages; messages may be null
func NewMessages(totalCount int32, messages []Message) *Messages {
	messagesTemp := Messages{
		tdCommon:   tdCommon{Type: "messages"},
		TotalCount: totalCount,
		Messages:   messages,
	}

	return &messagesTemp
}