# InstantLink

## Overview

Operations or actions related to retrieving an Instant Link.

### Available Operations

* [GetAuthURL](#getauthurl) - Authorization Url
* [GetResult](#getresult) - Instant Link Result

## GetAuthURL

As the starting point for Instant Link, this endpoint provides an authentication URL (appended with the first, unique verification fingerprint) that can be sent via SMS to initiate the middle authentication step.


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
    res, err := s.InstantLink.GetAuthURL(ctx, shared.AuthURLRequest{
        APIClientID: "C6f1j294x70dY3l76xU6",
        FinalTargetURL: "http://www.google.com",
        MobileNumber: "8167434789",
        RequestID: "7f83-b0c4-90e0-90b3-11e1-0800200c9a66",
        SessionID: "0109083438",
        SourceIP: prove.String("2607:fb90:be01:4122:e118:813f:736a:b7b9"),
        SubClientID: prove.String("D6hy5294x70dY3l76xU6"),
        SubscriptionCustomerID: prove.String("ThisIsMyCustomerId222"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AuthURLResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [shared.AuthURLRequest](../../models/shared/authurlrequest.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |


### Response

**[*operations.GetAuthURLResponse](../../models/operations/getauthurlresponse.md), error**


## GetResult

To complete the Instant Link flow, this endpoint passes the second, unique verification fingerprint returned by the mobile device, and identifies whether Instant Link was completed with or without carrier authentication, whether the link was clicked within the expiration time, and whether the input device was where the link was clicked.


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
    res, err := s.InstantLink.GetResult(ctx, shared.InstantLinkRequest{
        RequestID: prove.String("7f83-b0c4-90e0-90b3-11e10800200c9a66"),
        SessionID: prove.String("0109083438"),
        VerificationFingerprint: prove.String("F22440010AC782406249CFE0560F68EF"),
        APIClientID: "C6f1j294x70dY3l76xU6",
        SubClientID: prove.String("G6x1j294x70dY3l76xU6"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.InstantLinkResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [shared.InstantLinkRequest](../../models/shared/instantlinkrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[*operations.GetResultResponse](../../models/operations/getresultresponse.md), error**

