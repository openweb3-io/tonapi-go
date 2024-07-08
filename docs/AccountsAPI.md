# \AccountsAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountDnsBackResolve**](AccountsAPI.md#AccountDnsBackResolve) | **Get** /v2/accounts/{account_id}/dns/backresolve | 
[**AddressParse**](AccountsAPI.md#AddressParse) | **Get** /v2/address/{account_id}/parse | 
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /v2/accounts/{account_id} | 
[**GetAccountDiff**](AccountsAPI.md#GetAccountDiff) | **Get** /v2/accounts/{account_id}/diff | 
[**GetAccountDnsExpiring**](AccountsAPI.md#GetAccountDnsExpiring) | **Get** /v2/accounts/{account_id}/dns/expiring | 
[**GetAccountEvent**](AccountsAPI.md#GetAccountEvent) | **Get** /v2/accounts/{account_id}/events/{event_id} | 
[**GetAccountEvents**](AccountsAPI.md#GetAccountEvents) | **Get** /v2/accounts/{account_id}/events | 
[**GetAccountJettonBalance**](AccountsAPI.md#GetAccountJettonBalance) | **Get** /v2/accounts/{account_id}/jettons/{jetton_id} | 
[**GetAccountJettonHistoryByID**](AccountsAPI.md#GetAccountJettonHistoryByID) | **Get** /v2/accounts/{account_id}/jettons/{jetton_id}/history | 
[**GetAccountJettonsBalances**](AccountsAPI.md#GetAccountJettonsBalances) | **Get** /v2/accounts/{account_id}/jettons | 
[**GetAccountJettonsHistory**](AccountsAPI.md#GetAccountJettonsHistory) | **Get** /v2/accounts/{account_id}/jettons/history | 
[**GetAccountMultisigs**](AccountsAPI.md#GetAccountMultisigs) | **Get** /v2/accounts/{account_id}/multisigs | 
[**GetAccountNftItems**](AccountsAPI.md#GetAccountNftItems) | **Get** /v2/accounts/{account_id}/nfts | 
[**GetAccountPublicKey**](AccountsAPI.md#GetAccountPublicKey) | **Get** /v2/accounts/{account_id}/publickey | 
[**GetAccountSubscriptions**](AccountsAPI.md#GetAccountSubscriptions) | **Get** /v2/accounts/{account_id}/subscriptions | 
[**GetAccountTraces**](AccountsAPI.md#GetAccountTraces) | **Get** /v2/accounts/{account_id}/traces | 
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Post** /v2/accounts/_bulk | 
[**ReindexAccount**](AccountsAPI.md#ReindexAccount) | **Post** /v2/accounts/{account_id}/reindex | 
[**SearchAccounts**](AccountsAPI.md#SearchAccounts) | **Get** /v2/accounts/search | 



## AccountDnsBackResolve

> DomainNames AccountDnsBackResolve(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AccountDnsBackResolve(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AccountDnsBackResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccountDnsBackResolve`: DomainNames
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AccountDnsBackResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountDnsBackResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainNames**](DomainNames.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressParse

> AddressParse200Response AddressParse(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.AddressParse(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.AddressParse``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressParse`: AddressParse200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.AddressParse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddressParseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddressParse200Response**](AddressParse200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDiff

> GetAccountDiff200Response GetAccountDiff(ctx, accountId).StartDate(startDate).EndDate(endDate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	startDate := int64(1668436763) // int64 | 
	endDate := int64(1668436763) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountDiff(context.Background(), accountId).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountDiff`: GetAccountDiff200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **int64** |  | 
 **endDate** | **int64** |  | 

### Return type

[**GetAccountDiff200Response**](GetAccountDiff200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountDnsExpiring

> DnsExpiring GetAccountDnsExpiring(ctx, accountId).Period(period).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	period := int32(56) // int32 | number of days before expiration (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountDnsExpiring(context.Background(), accountId).Period(period).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountDnsExpiring``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountDnsExpiring`: DnsExpiring
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountDnsExpiring`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDnsExpiringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **int32** | number of days before expiration | 

### Return type

[**DnsExpiring**](DnsExpiring.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountEvent

> AccountEvent GetAccountEvent(ctx, accountId, eventId).AcceptLanguage(acceptLanguage).SubjectOnly(subjectOnly).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	eventId := "97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | event ID or transaction hash in hex (without 0x) or base64url format
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	subjectOnly := true // bool | filter actions where requested account is not real subject (for example sender or receiver jettons) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountEvent(context.Background(), accountId, eventId).AcceptLanguage(acceptLanguage).SubjectOnly(subjectOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEvent`: AccountEvent
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 
**eventId** | **string** | event ID or transaction hash in hex (without 0x) or base64url format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **subjectOnly** | **bool** | filter actions where requested account is not real subject (for example sender or receiver jettons) | [default to false]

### Return type

[**AccountEvent**](AccountEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountEvents

> AccountEvents GetAccountEvents(ctx, accountId).Limit(limit).AcceptLanguage(acceptLanguage).Initiator(initiator).SubjectOnly(subjectOnly).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	limit := int32(20) // int32 | 
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	initiator := true // bool | Show only events that are initiated by this account (optional) (default to false)
	subjectOnly := true // bool | filter actions where requested account is not real subject (for example sender or receiver jettons) (optional) (default to false)
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	startDate := int64(1668436763) // int64 |  (optional)
	endDate := int64(1668436763) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountEvents(context.Background(), accountId).Limit(limit).AcceptLanguage(acceptLanguage).Initiator(initiator).SubjectOnly(subjectOnly).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountEvents`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **initiator** | **bool** | Show only events that are initiated by this account | [default to false]
 **subjectOnly** | **bool** | filter actions where requested account is not real subject (for example sender or receiver jettons) | [default to false]
 **beforeLt** | **int64** | omit this parameter to get last events | 
 **startDate** | **int64** |  | 
 **endDate** | **int64** |  | 

### Return type

[**AccountEvents**](AccountEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountJettonBalance

> JettonBalance GetAccountJettonBalance(ctx, accountId, jettonId).Currencies(currencies).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	jettonId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | jetton ID
	currencies := []string{"Inner_example"} // []string | accept ton and all possible fiat currencies, separated by commas (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountJettonBalance(context.Background(), accountId, jettonId).Currencies(currencies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountJettonBalance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountJettonBalance`: JettonBalance
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountJettonBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 
**jettonId** | **string** | jetton ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountJettonBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **currencies** | **[]string** | accept ton and all possible fiat currencies, separated by commas | 

### Return type

[**JettonBalance**](JettonBalance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountJettonHistoryByID

> AccountEvents GetAccountJettonHistoryByID(ctx, accountId, jettonId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	jettonId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | jetton ID
	limit := int32(100) // int32 | 
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	startDate := int64(1668436763) // int64 |  (optional)
	endDate := int64(1668436763) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountJettonHistoryByID(context.Background(), accountId, jettonId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountJettonHistoryByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountJettonHistoryByID`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountJettonHistoryByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 
**jettonId** | **string** | jetton ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountJettonHistoryByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** |  | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **beforeLt** | **int64** | omit this parameter to get last events | 
 **startDate** | **int64** |  | 
 **endDate** | **int64** |  | 

### Return type

[**AccountEvents**](AccountEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountJettonsBalances

> JettonsBalances GetAccountJettonsBalances(ctx, accountId).Currencies(currencies).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	currencies := []string{"Inner_example"} // []string | accept ton and all possible fiat currencies, separated by commas (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountJettonsBalances(context.Background(), accountId).Currencies(currencies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountJettonsBalances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountJettonsBalances`: JettonsBalances
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountJettonsBalances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountJettonsBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currencies** | **[]string** | accept ton and all possible fiat currencies, separated by commas | 

### Return type

[**JettonsBalances**](JettonsBalances.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountJettonsHistory

> AccountEvents GetAccountJettonsHistory(ctx, accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	limit := int32(100) // int32 | 
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	startDate := int64(1668436763) // int64 |  (optional)
	endDate := int64(1668436763) // int64 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountJettonsHistory(context.Background(), accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountJettonsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountJettonsHistory`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountJettonsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountJettonsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **beforeLt** | **int64** | omit this parameter to get last events | 
 **startDate** | **int64** |  | 
 **endDate** | **int64** |  | 

### Return type

[**AccountEvents**](AccountEvents.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountMultisigs

> Multisigs GetAccountMultisigs(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountMultisigs(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountMultisigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountMultisigs`: Multisigs
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountMultisigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountMultisigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Multisigs**](Multisigs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNftItems

> NftItems GetAccountNftItems(ctx, accountId).Collection(collection).Limit(limit).Offset(offset).IndirectOwnership(indirectOwnership).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	collection := "0:06d811f426598591b32b2c49f29f66c821368e4acb1de16762b04e0174532465" // string | nft collection (optional)
	limit := int32(56) // int32 |  (optional) (default to 1000)
	offset := int32(56) // int32 |  (optional) (default to 0)
	indirectOwnership := true // bool | Selling nft items in ton implemented usually via transfer items to special selling account. This option enables including items which owned not directly. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountNftItems(context.Background(), accountId).Collection(collection).Limit(limit).Offset(offset).IndirectOwnership(indirectOwnership).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountNftItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNftItems`: NftItems
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountNftItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNftItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collection** | **string** | nft collection | 
 **limit** | **int32** |  | [default to 1000]
 **offset** | **int32** |  | [default to 0]
 **indirectOwnership** | **bool** | Selling nft items in ton implemented usually via transfer items to special selling account. This option enables including items which owned not directly. | [default to false]

### Return type

[**NftItems**](NftItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountPublicKey

> GetAccountPublicKey200Response GetAccountPublicKey(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountPublicKey(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountPublicKey`: GetAccountPublicKey200Response
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAccountPublicKey200Response**](GetAccountPublicKey200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountSubscriptions

> Subscriptions GetAccountSubscriptions(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountSubscriptions(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSubscriptions`: Subscriptions
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriptions**](Subscriptions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountTraces

> TraceIDs GetAccountTraces(ctx, accountId).BeforeLt(beforeLt).Limit(limit).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountTraces(context.Background(), accountId).BeforeLt(beforeLt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountTraces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountTraces`: TraceIDs
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountTraces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountTracesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beforeLt** | **int64** | omit this parameter to get last events | 
 **limit** | **int32** |  | [default to 100]

### Return type

[**TraceIDs**](TraceIDs.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> Accounts GetAccounts(ctx).Currency(currency).GetAccountsRequest(getAccountsRequest).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	currency := "usd" // string |  (optional)
	getAccountsRequest := *openapiclient.NewGetAccountsRequest([]string{"0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621"}) // GetAccountsRequest | a list of account ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).Currency(currency).GetAccountsRequest(getAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **currency** | **string** |  | 
 **getAccountsRequest** | [**GetAccountsRequest**](GetAccountsRequest.md) | a list of account ids | 

### Return type

[**Accounts**](Accounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReindexAccount

> ReindexAccount(ctx, accountId).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccountsAPI.ReindexAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.ReindexAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReindexAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAccounts

> FoundAccounts SearchAccounts(ctx).Name(name).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/tonapi"
)

func main() {
	name := "name_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.SearchAccounts(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.SearchAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAccounts`: FoundAccounts
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.SearchAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** |  | 

### Return type

[**FoundAccounts**](FoundAccounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

