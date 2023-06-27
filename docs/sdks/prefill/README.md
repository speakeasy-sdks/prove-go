# PreFill

## Overview

Operations or actions related to identity

### Available Operations

* [GetIdentity](#getidentity) - Get the Identity Information

## GetIdentity

This endpoint request sends the phone number previously checked in the Eligibility API, along with either the SSN (full or last four digits) or date of birth, to then return the correct identity information for the consumer to autofill the customer's form/application.


### Example Usage

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
    res, err := s.PreFill.GetIdentity(ctx, shared.IdentityRequest{
        ConsentStatus: shared.ConsentStatusOptedOut,
        Dob: prove.String("1982-07-19"),
        Last4: prove.String("6227"),
        NumberOfAddresses: prove.String("3"),
        NumberOfEmails: prove.String("3"),
        PhoneNumber: "6464778753",
        RequestID: "14f3-b0c4-90e0-90b3-11e1-0800200c9a66",
        Ssn: prove.String("470806227"),
        SubscriptionCustomerID: prove.String("ThisIsMyCustomerId222"),
        TrustScore: prove.String("true"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IdentityResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.IdentityRequest](../../models/shared/identityrequest.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.GetIdentityResponse](../../models/operations/getidentityresponse.md), error**

