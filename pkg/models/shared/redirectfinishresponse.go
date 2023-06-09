// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type RedirectFinishResponseResponse struct {
	// A unique identifier generated by Prove that can be used to troubleshoot any problems with this transaction.
	AuthenticateFinishTransactionID *string `json:"AuthenticateFinishTransactionId,omitempty"`
	// A transaction ID denoting the authentication event.
	AuthenticationCode *string `json:"AuthenticationCode,omitempty"`
	// The expiration timestamp in UTC of the authentication code. Timestamp is in ISO-8601 format.
	AuthenticationExpiration *string `json:"AuthenticationExpiration,omitempty"`
	// The country code associated with this phone number.
	MobileCountryCode *string `json:"MobileCountryCode,omitempty"`
	// The phone number of the originating phone.
	MobileNumber *string `json:"MobileNumber,omitempty"`
	// The carrier related to the phone number.
	MobileOperatorName *string `json:"MobileOperatorName,omitempty"`
	// A persistent ID that uniquely identifies a telephone subscriber.
	PayfoneAlias *string `json:"PayfoneAlias,omitempty"`
}

// RedirectFinishResponse - Success
type RedirectFinishResponse struct {
	// A text string that defines the cause of the status code.
	Description *string `json:"Description,omitempty"`
	// The requestId from the request, reflected back for tracking purposes.
	RequestID *string                         `json:"RequestId,omitempty"`
	Response  *RedirectFinishResponseResponse `json:"Response,omitempty"`
	// The status of the request. A response of 0 indicates success. Any non-0 response is an error indication. For more information on status codes, see the Error and Status Codes section.
	Status *int64 `json:"Status,omitempty"`
}
