// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type InstantLinkRequest struct {
	// The requestId from the request, reflected back for tracking purposes.
	RequestID *string `json:"RequestId,omitempty"`
	// Session identifier associated with this request flow. It should be the same session ID as the /getAuthUrl call.
	SessionID *string `json:"SessionId,omitempty"`
	// The second, unique verification fingerprint returned by the mobile device
	VerificationFingerprint *string `json:"VerificationFingerprint,omitempty"`
	// Prove-issued unique, private key that identifies the API Client.
	APIClientID string `json:"apiClientId"`
	// Prove-issued unique, private key that identifies the API Client.
	SubClientID *string `json:"subClientId,omitempty"`
}