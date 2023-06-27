# Trust

## Overview

Operations or actions related to Trust score

### Available Operations

* [GetScore](#getscore) - Trust Score

## GetScore

This endpoint queries the consumer's phone number (or PayfoneAlias) to return a Prove Trust Score and, potentially, other attributes to indicate the trustworthiness of a phone number and its identity.

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
    res, err := s.Trust.GetScore(ctx, shared.ScoreRequest{
        ConsentStatus: shared.ConsentStatusOptedOut,
        Details: prove.String("true"),
        PayfoneAlias: prove.String("FB9990AC4VKRKKX801F617431E9C9B5D10MEK1KQZ3T9PCC265D01B991F09559712BEF30B7115C258F6G38FB0EF56226588191F2011FF1246401DF6B6"),
        PhoneNumber: "8167434789",
        RequestID: "59c04ea5-5aae-4934-a96f-f32051f5078b",
        SubscriptionCustomerID: prove.String("ThisIsMyCustomerId222"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ScoreResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `request`                                                  | [shared.ScoreRequest](../../models/shared/scorerequest.md) | :heavy_check_mark:                                         | The request object to use for the request.                 |


### Response

**[*operations.GetScoreResponse](../../models/operations/getscoreresponse.md), error**

