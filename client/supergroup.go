// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetSupergroup Returns information about a supergroup or a channel by its identifier. This is an offline request if the current user is not a bot
// @param supergroupID Supergroup or channel identifier
func (client *Client) GetSupergroup(supergroupID int64) (*tdlib.Supergroup, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "getSupergroup",
		"supergroup_id": supergroupID,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var supergroupDummy tdlib.Supergroup
	err = json.Unmarshal(result.Raw, &supergroupDummy)
	return &supergroupDummy, err

}