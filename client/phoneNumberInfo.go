// AUTOGENERATED - DO NOT EDIT

package client

import (
	"encoding/json"

	"github.com/satyshef/go-tdlib/tdlib"
)

// GetPhoneNumberInfo Returns information about a phone number by its prefix. Can be called before authorization
// @param phoneNumberPrefix The phone number prefix
func (client *Client) GetPhoneNumberInfo(phoneNumberPrefix string) (*tdlib.PhoneNumberInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":               "getPhoneNumberInfo",
		"phone_number_prefix": phoneNumberPrefix,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var phoneNumberInfo tdlib.PhoneNumberInfo
	err = json.Unmarshal(result.Raw, &phoneNumberInfo)
	return &phoneNumberInfo, err

}

// GetPhoneNumberInfoSync Returns information about a phone number by its prefix synchronously. getCountries must be called at least once after changing localization to the specified language if properly localized country information is expected. Can be called synchronously
// @param languageCode A two-letter ISO 639-1 country code for country information localization
// @param phoneNumberPrefix The phone number prefix
func (client *Client) GetPhoneNumberInfoSync(languageCode string, phoneNumberPrefix string) (*tdlib.PhoneNumberInfo, error) {
	result, err := client.SendAndCatch(tdlib.UpdateData{
		"@type":               "getPhoneNumberInfoSync",
		"language_code":       languageCode,
		"phone_number_prefix": phoneNumberPrefix,
	})

	if err != nil {
		return nil, err
	}

	if result.Data["@type"].(string) == "error" {
		return nil, tdlib.RequestError{Code: int(result.Data["code"].(float64)), Message: result.Data["message"].(string)}
	}

	var phoneNumberInfo tdlib.PhoneNumberInfo
	err = json.Unmarshal(result.Raw, &phoneNumberInfo)
	return &phoneNumberInfo, err

}
