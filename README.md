<div align="center">
   <img src="https://user-images.githubusercontent.com/6267663/232879662-b4423dfa-4430-4dcd-86eb-641ba05ad440.png">
   <h1>Go SDK</h1>
   <p><strong>The modern way of proving identity</strong></p>
   <a href="https://developer.prove.com/public/reference/introduction"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/prove-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/prove-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

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
