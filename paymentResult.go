// AUTOGENERATED - DO NOT EDIT

package tdlib

// PaymentResult Contains the result of a payment request
type PaymentResult struct {
	tdCommon
	Success         bool   `json:"success"`          // True, if the payment request was successful; otherwise the verification_url will be non-empty
	VerificationURL string `json:"verification_url"` // URL for additional payment credentials verification
}

// MessageType return the string telegram-type of PaymentResult
func (paymentResult *PaymentResult) MessageType() string {
	return "paymentResult"
}

// NewPaymentResult creates a new PaymentResult
//
// @param success True, if the payment request was successful; otherwise the verification_url will be non-empty
// @param verificationURL URL for additional payment credentials verification
func NewPaymentResult(success bool, verificationURL string) *PaymentResult {
	paymentResultTemp := PaymentResult{
		tdCommon:        tdCommon{Type: "paymentResult"},
		Success:         success,
		VerificationURL: verificationURL,
	}

	return &paymentResultTemp
}
