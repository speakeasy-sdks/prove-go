// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// ScoreResponseResponseCarrierTenure - The amount of time Prove has been able to track the carrier associated with the subscriber. Defined with a minimum and maximum date.
type ScoreResponseResponseCarrierTenure struct {
	MaximumDate *time.Time `json:"maximumDate,omitempty"`
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

// ScoreResponseResponseDeviceTenure - The amount of time Prove has been able to track the device association with the subscriber. Defined with a minimum and maximum date.
type ScoreResponseResponseDeviceTenure struct {
	MaximumDate *time.Time `json:"maximumDate,omitempty"`
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

// ScoreResponseResponsePayfoneTenure - The amount of time Prove has been able to track the subscriber. Defined with a minimum and maximum date.
type ScoreResponseResponsePayfoneTenure struct {
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

// ScoreResponseResponsePhoneNumberTenure - The amount of time Prove has been able to track the phone number associated with the subscriber. Defined with a minimum and maximum date.
type ScoreResponseResponsePhoneNumberTenure struct {
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

// ScoreResponseResponsePortedDate - The date associated with a port of the phone number. Defined with a minimum and maximum date.
type ScoreResponseResponsePortedDate struct {
	MaximumDate *time.Time `json:"maximumDate,omitempty"`
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

// ScoreResponseResponseSimTenure - The amount of time Prove has been able to track the SIM association with the subscriber. Defined with a minimum and maximum date.
type ScoreResponseResponseSimTenure struct {
	MaximumDate *time.Time `json:"maximumDate,omitempty"`
	MinimumDate *time.Time `json:"minimumDate,omitempty"`
}

type ScoreResponseResponse struct {
	// The carrier related to the phone number.
	Carrier *string `json:"carrier,omitempty"`
	// The current phone number network status associated with the carriers. Values include: Active, Suspended, Disconnected, Unknown
	CarrierStatus *string `json:"carrierStatus,omitempty"`
	// The amount of time Prove has been able to track the carrier associated with the subscriber. Defined with a minimum and maximum date.
	CarrierTenure *ScoreResponseResponseCarrierTenure `json:"carrierTenure,omitempty"`
	// The country code associated with the phone number
	CountryCode *string `json:"countryCode,omitempty"`
	// The amount of time Prove has been able to track the device association with the subscriber. Defined with a minimum and maximum date.
	DeviceTenure *ScoreResponseResponseDeviceTenure `json:"deviceTenure,omitempty"`
	// The number of times the mobile subscriber has changed their device within the last 90 days.
	DeviceVelocity *int64 `json:"deviceVelocity,omitempty"`
	// An indicator to communicate whether Prove has ever transacted with the input phone number before the current transaction.
	IsBaselined *bool `json:"isBaselined,omitempty"`
	// Line type associated with the phone number. Possible values are: Mobile, Landline, FixedVoIP, NonFixedVoIP
	LineType *string `json:"lineType,omitempty"`
	// A persistent ID that uniquely identifies a telephone subscriber.
	PayfoneAlias *string `json:"payfoneAlias,omitempty"`
	// The amount of time Prove has been able to track the subscriber. Defined with a minimum and maximum date.
	PayfoneTenure *ScoreResponseResponsePayfoneTenure `json:"payfoneTenure,omitempty"`
	// The phone number associated with the subscriber.
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// The amount of time Prove has been able to track the phone number associated with the subscriber. Defined with a minimum and maximum date.
	PhoneNumberTenure *ScoreResponseResponsePhoneNumberTenure `json:"phoneNumberTenure,omitempty"`
	// The number of times the subscriber has changed their phone number.
	PhoneNumberVelocity *int64 `json:"phoneNumberVelocity,omitempty"`
	// The number of times the subscriber has changed their carrier within the last 90 days.
	PortVelocity *int64 `json:"portVelocity,omitempty"`
	// The date associated with a port of the phone number. Defined with a minimum and maximum date.
	PortedDate *ScoreResponseResponsePortedDate `json:"portedDate,omitempty"`
	// An array of indicators provide additional context about the transaction. See Reason Codes Reference Information for detailed reason codes.
	ReasonCodes []string `json:"reasonCodes,omitempty"`
	// The amount of time Prove has been able to track the SIM association with the subscriber. Defined with a minimum and maximum date.
	SimTenure *ScoreResponseResponseSimTenure `json:"simTenure,omitempty"`
	// The number of times the mobile subscriber has changed their device within the last 90 days.
	SimVelocity *int64 `json:"simVelocity,omitempty"`
	// A bitmapped value that represents networkStatus, accountType, accountRole, and customerType. For more information, see the Trust Score Integration Guide.
	StatusIndex *string `json:"statusIndex,omitempty"`
	// Unique transaction identifier used to identify the results of the request.
	TransactionID *string `json:"transactionId,omitempty"`
	// An integer value ranging from 0–1000 denotes the real-time trustworthiness of a phone number. 1000 indicates perfect trust, while 0 indicates a complete lack of trust. For more information on Trust Score, see the Trust Score Integration Guide.
	TrustScore *int64 `json:"trustScore,omitempty"`
}

// ScoreResponse - Success - consentStatus=optedIn, details=true
type ScoreResponse struct {
	// A text string that defines the cause of the status code.
	Description *string `json:"description,omitempty"`
	// The requestId from the request, reflected back for tracking purposes.
	RequestID *string                `json:"requestId,omitempty"`
	Response  *ScoreResponseResponse `json:"response,omitempty"`
	// The status of the request. A response of 0 indicates success. Any non-0 response is an error indication. For more information on status codes, see the Error and Status Codes section.
	Status *int64 `json:"status,omitempty"`
}