// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// RequestPasswordRecovery Requests to send a 2-step verification password recovery code to an email address that was previously set up
func (client *Client) RequestPasswordRecovery() (*tdlib.EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "requestPasswordRecovery",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var emailAddressAuthenticationCodeInfo tdlib.EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// SendEmailAddressVerificationCode Sends a code to verify an email address to be added to a user's Telegram Passport
// @param emailAddress Email address
func (client *Client) SendEmailAddressVerificationCode(emailAddress string) (*tdlib.EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":         "sendEmailAddressVerificationCode",
		"email_address": emailAddress,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var emailAddressAuthenticationCodeInfo tdlib.EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}

// ResendEmailAddressVerificationCode Re-sends the code to verify an email address to be added to a user's Telegram Passport
func (client *Client) ResendEmailAddressVerificationCode() (*tdlib.EmailAddressAuthenticationCodeInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type": "resendEmailAddressVerificationCode",
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var emailAddressAuthenticationCodeInfo tdlib.EmailAddressAuthenticationCodeInfo
	err = json.Unmarshal(result.Raw, &emailAddressAuthenticationCodeInfo)
	return &emailAddressAuthenticationCodeInfo, err

}
