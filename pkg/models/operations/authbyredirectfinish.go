// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"AuthByRedirect/pkg/models/shared"
	"net/http"
)

type AuthByRedirectFinishResponse struct {
	ContentType string
	// Success
	RedirectFinishResponse *shared.RedirectFinishResponse
	StatusCode             int
	RawResponse            *http.Response
}