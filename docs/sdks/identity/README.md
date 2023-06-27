# Identity

## Overview

Operations or actions related to identity

### Available Operations

* [Verify](#verify) - Verify Consumer Information

## Verify

This endpoint submits the consumer's personal identifying information (PII) for verification; the response returns that verification decision and, if requested, detailed field-level information on each PII piece submitted.

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

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `request`                                                    | [shared.VerifyRequest](../../models/shared/verifyrequest.md) | :heavy_check_mark:                                           | The request object to use for the request.                   |


### Response

**[*operations.VerifyResponse](../../models/operations/verifyresponse.md), error**

