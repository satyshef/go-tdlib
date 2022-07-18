// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetMessageLink Returns an HTTPS link to a message in a chat. Available only for already sent messages in supergroups and channels, or if message.can_get_media_timestamp_links and a media timestamp link is generated. This is an offline request
// @param chatID Identifier of the chat to which the message belongs
// @param messageID Identifier of the message
// @param mediaTimestamp If not 0, timestamp from which the video/audio/video note/voice note playing should start, in seconds. The media can be in the message content or in its web page preview
// @param forAlbum Pass true to create a link for the whole media album
// @param forComment Pass true to create a link to the message as a channel post comment, or from a message thread
func (client *Client) GetMessageLink(chatID int64, messageID int64, mediaTimestamp int32, forAlbum bool, forComment bool) (*tdlib.MessageLink, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":           "getMessageLink",
		"chat_id":         chatID,
		"message_id":      messageID,
		"media_timestamp": mediaTimestamp,
		"for_album":       forAlbum,
		"for_comment":     forComment,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var messageLink tdlib.MessageLink
	err = json.Unmarshal(result.Raw, &messageLink)
	return &messageLink, err

}
