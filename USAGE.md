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