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
	"github.com/speakeasy-sdks/prove-go"
	"github.com/speakeasy-sdks/prove-go/pkg/models/shared"
)

func main() {
    s := prove.New(
        prove.WithSecurity(shared.Security{
            BearerAuth: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Identity.Verify(ctx, shared.VerifyRequest{
        Address: prove.String("00 Mallard Park"),
        City: prove.String("New Orleans"),
        ConsentStatus: shared.ConsentStatusOptedOut.ToPointer(),
        Details: prove.String("true"),
        Dob: prove.String("1982-07-19"),
        DriversLicenseNumber: prove.String("123456789"),
        DriversLicenseState: prove.String("CO"),
        EmailAddress: prove.String("mlongok@amazonaws.com"),
        ExtendedAddress: prove.String("Apartment 3A"),
        FirstName: "Marcia",
        Last4: prove.String("6227"),
        LastName: "Longo-Jones",
        LastVerified: prove.String("2022-07-19"),
        PayfoneAlias: prove.String("8EDE1ACC4VKRKKX8B91612DE9DCFB77DF0MEK1KQZ3T9PA44306E401F13F8B12A283E6E941AC03B46F6G3FD4CED48D730FC618931737A5FEE6E31E447"),
        PhoneNumber: prove.String("6464778753"),
        PhoneUpdate: prove.String("true"),
        PostalCode: prove.String("70165"),
        Region: prove.String("LA"),
        RequestID: "14f3-b0c4-90e0-90b3-11e1-0800200c9a66",
        Ssn: prove.String("470806227"),
        SubscriptionCustomerID: prove.String("ThisIsMyCustomerId222"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VerifyResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Identity](docs/sdks/identity/README.md)

* [Verify](docs/sdks/identity/README.md#verify) - Verify Consumer Information

### [InstantLink](docs/sdks/instantlink/README.md)

* [GetAuthURL](docs/sdks/instantlink/README.md#getauthurl) - Authorization Url
* [GetResult](docs/sdks/instantlink/README.md#getresult) - Instant Link Result

### [MobileAuth](docs/sdks/mobileauth/README.md)

* [AuthByRedirect](docs/sdks/mobileauth/README.md#authbyredirect) - Authenticate by Redirect
* [AuthByRedirectFinish](docs/sdks/mobileauth/README.md#authbyredirectfinish) - Authenticate by Redirect Finish

### [PreFill](docs/sdks/prefill/README.md)

* [GetIdentity](docs/sdks/prefill/README.md#getidentity) - Get the Identity Information

### [Trust](docs/sdks/trust/README.md)

* [GetScore](docs/sdks/trust/README.md#getscore) - Trust Score
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
