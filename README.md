# AuthByRedirect

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/prove-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import(
	"context"
	"log"
	"AuthByRedirect"
	"AuthByRedirect/pkg/models/shared"
)

func main() {
    s := authbyredirect.New()

    ctx := context.Background()
    res, err := s.MobileAuth.AuthByRedirect(ctx, shared.RedirectRequest{
        APIClientID: "C6f1j294x70dY3l76xU6",
        ConsentCollectedTimestamp: authbyredirect.String("2022-05-11"),
        ConsentDescription: authbyredirect.String("Test Description"),
        ConsentStatus: shared.ConsentStatusOptedOut,
        ConsentTransactionID: authbyredirect.String("EWSrelease-01092020-testTMO5"),
        DeviceIP: "2607:fb90:be01:4122:e118:813f:736a:b7b9",
        FinalTargetURL: "http://www.google.com",
        MobileNetworkOperator: authbyredirect.String("corrupti"),
        RequestID: "7f83-b0c4-90e0-90b3-11e10800200c9a66",
        SubClientID: authbyredirect.String("provident"),
        SubscriptionCustomerID: authbyredirect.String("ThisIsMyCustomerId222"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RedirectResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [MobileAuth](docs/sdks/mobileauth/README.md)

* [AuthByRedirect](docs/sdks/mobileauth/README.md#authbyredirect) - Authenticate by Redirect
* [AuthByRedirectFinish](docs/sdks/mobileauth/README.md#authbyredirectfinish) - Authenticate by Redirect Finish
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
