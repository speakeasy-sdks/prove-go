# ConsentStatus

Denotes whether consent has been provided by the mobile subscriber for this transaction. Acceptable values are:
- `optedIn` - The end user has provided consent for the collection of their personal data. - `optedOut` - The end user has refused to allow collection of their personal data. - `notCollected` - No attempt has been made to obtain consent from the end user. - `unknown` - The status of consent collection is unknown.
Note: This value must be optedIn in order to access MNO data. Legacy customers' configurations are not affected by this requirement at this time.


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ConsentStatusOptedIn`      | optedIn                     |
| `ConsentStatusOptedOut`     | optedOut                    |
| `ConsentStatusNotCollected` | notCollected                |
| `ConsentStatusUnknown`      | unknown                     |