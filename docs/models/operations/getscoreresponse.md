# GetScoreResponse


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           |
| `ScoreResponse`                                               | [*shared.ScoreResponse](../../models/shared/scoreresponse.md) | :heavy_minus_sign:                                            | Success - consentStatus=optedIn, details=true                 |
| `StatusCode`                                                  | *int*                                                         | :heavy_check_mark:                                            | N/A                                                           |
| `RawResponse`                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)        | :heavy_minus_sign:                                            | N/A                                                           |