// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/prove-go/pkg/models/shared"
	"net/http"
)

type GetIdentityResponse struct {
	ContentType string
	// Success
	IdentityResponse *shared.IdentityResponse
	StatusCode       int
	RawResponse      *http.Response
}
