// AUTOGENERATED - DO NOT EDIT

package tdlib

// BankCardActionOpenURL Describes an action associated with a bank card number
type BankCardActionOpenURL struct {
	tdCommon
	Text string `json:"text"` // Action text
	URL  string `json:"url"`  // The URL to be opened
}

// MessageType return the string telegram-type of BankCardActionOpenURL
func (bankCardActionOpenURL *BankCardActionOpenURL) MessageType() string {
	return "bankCardActionOpenUrl"
}

// NewBankCardActionOpenURL creates a new BankCardActionOpenURL
//
// @param text Action text
// @param uRL The URL to be opened
func NewBankCardActionOpenURL(text string, uRL string) *BankCardActionOpenURL {
	bankCardActionOpenURLTemp := BankCardActionOpenURL{
		tdCommon: tdCommon{Type: "bankCardActionOpenUrl"},
		Text:     text,
		URL:      uRL,
	}

	return &bankCardActionOpenURLTemp
}
