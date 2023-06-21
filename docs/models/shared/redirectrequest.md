# RedirectRequest


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `APIClientID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Prove-issued unique, private key that identifies the API Client.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | C6f1j294x70dY3l76xU6                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `ConsentCollectedTimestamp`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Date/Time field where the original consent was collected by the client. The string format is YYYY-MM-DD. **Required** if `ConsentStatus` = `optedIn`                                                                                                                                                                                                                                                                                                                                                                                                                                          | 2022-05-11                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| `ConsentDescription`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | The client describes the type of consent (electronic/paper), use case and reference to the version of Terms and Conditions (T&Cs), if applicable. **Required** if `ConsentStatus` = `optedIn`                                                                                                                                                                                                                                                                                                                                                                                                 | Test Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| `ConsentStatus`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | [ConsentStatus](../../models/shared/consentstatus.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Denotes whether consent has been provided by the mobile subscriber for this transaction. Acceptable values are:<br/>- `optedIn` - The end user has provided consent for the collection of their personal data. - `optedOut` - The end user has refused to allow collection of their personal data. - `notCollected` - No attempt has been made to obtain consent from the end user. - `unknown` - The status of consent collection is unknown.<br/>Note: This value must be optedIn in order to access MNO data. Legacy customers' configurations are not affected by this requirement at this time.<br/> | optedOut                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| `ConsentTransactionID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Uniquely identify the consent collected by the client. **Required** if `ConsentStatus` = `optedIn`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | EWSrelease-01092020-testTMO5                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |
| `DeviceIP`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | The IP address of the mobile device to be authenticated.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | 2607:fb90:be01:4122:e118:813f:736a:b7b9                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| `FinalTargetURL`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | The URL of the client server that will be called back by the phone, providing the VFP for the result call.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    | http://www.google.com                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| `MobileNetworkOperator`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | The mobile network operator that currently operates the phone number. Internal note: This is a legacy parameter and is no longer used.                                                                                                                                                                                                                                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `RequestID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Unique identifier associated with this request. This identifier must be unique for each transaction. The RequestId input field is rejected as invalid if it contains any special characters other than dashes, periods, underbars, plus signs, equal signs, or forward slashes. Max length is 128 bytes.                                                                                                                                                                                                                                                                                      | 7f83-b0c4-90e0-90b3-11e10800200c9a66                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |
| `SubClientID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Prove-issued unique, private key that identifies Sub Client.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| `SubscriptionCustomerID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | **string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | A unique identifier for our customer's subscribers for Identity Manager.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | ThisIsMyCustomerId222                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |