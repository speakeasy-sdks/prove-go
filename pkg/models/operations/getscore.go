// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/prove-go/pkg/models/shared"
	"net/http"
)

type GetScoreResponse struct {
	ContentType string
	// Success - consentStatus=optedIn, details=true
	ScoreResponse *shared.ScoreResponse
	StatusCode    int
	RawResponse   *http.Response
}
