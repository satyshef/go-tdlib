// AUTOGENERATED - DO NOT EDIT

package tdlib

// PhoneNumberAuthenticationSettings Contains settings for the authentication of the user's phone number
type PhoneNumberAuthenticationSettings struct {
	tdCommon
	AllowFlashCall       bool `json:"allow_flash_call"`        // Pass true if the authentication code may be sent via flash call to the specified phone number
	AllowMissedCall 	 bool `json:"allow_missed_call"`	   // Pass true if the authentication code may be sent via a missed call to the specified phone number.
	IsCurrentPhoneNumber bool `json:"is_current_phone_number"` // Pass true if the authenticated phone number is used on the current device
	AllowSmsRetrieverAPI bool `json:"allow_sms_retriever_api"` // For official applications only. True, if the application can use Android SMS Retriever API (requires Google Play Services >= 10.2) to automatically receive the authentication code from the SMS. See https://developers.google.com/identity/sms-retriever/ for more details
}

// MessageType return the string telegram-type of PhoneNumberAuthenticationSettings
func (phoneNumberAuthenticationSettings *PhoneNumberAuthenticationSettings) MessageType() string {
	return "phoneNumberAuthenticationSettings"
}

// NewPhoneNumberAuthenticationSettings creates a new PhoneNumberAuthenticationSettings
//
// @param allowFlashCall Pass true if the authentication code may be sent via flash call to the specified phone number
// @param allowMissedCall Pass true if the authentication code may be sent via a missed call to the specified phone number.
// @param isCurrentPhoneNumber Pass true if the authenticated phone number is used on the current device
// @param allowSmsRetrieverAPI For official applications only. True, if the application can use Android SMS Retriever API (requires Google Play Services >= 10.2) to automatically receive the authentication code from the SMS. See https://developers.google.com/identity/sms-retriever/ for more details
func NewPhoneNumberAuthenticationSettings(allowFlashCall bool, alloMissedCall bool, isCurrentPhoneNumber bool, allowSmsRetrieverAPI bool) *PhoneNumberAuthenticationSettings {
	phoneNumberAuthenticationSettingsTemp := PhoneNumberAuthenticationSettings{
		tdCommon:             tdCommon{Type: "phoneNumberAuthenticationSettings"},
		AllowFlashCall:       allowFlashCall,
		AllowMissedCall: 	  alloMissedCall,
		IsCurrentPhoneNumber: isCurrentPhoneNumber,
		AllowSmsRetrieverAPI: allowSmsRetrieverAPI,
	}

	return &phoneNumberAuthenticationSettingsTemp
}
