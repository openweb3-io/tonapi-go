# \InscriptionsAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInscriptions**](InscriptionsAPI.md#GetAccountInscriptions) | **Get** /v2/experimental/accounts/{account_id}/inscriptions | 
[**GetAccountInscriptionsHistory**](InscriptionsAPI.md#GetAccountInscriptionsHistory) | **Get** /v2/experimental/accounts/{account_id}/inscriptions/history | 
[**GetAccountInscriptionsHistoryByTicker**](InscriptionsAPI.md#GetAccountInscriptionsHistoryByTicker) | **Get** /v2/experimental/accounts/{account_id}/inscriptions/{ticker}/history | 
[**GetInscriptionOpTemplate**](InscriptionsAPI.md#GetInscriptionOpTemplate) | **Get** /v2/experimental/inscriptions/op-template | 



## GetAccountInscriptions

> InscriptionBalances GetAccountInscriptions(ctx, accountId).Limit(limit).Offset(offset).Execute()





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
	limit := int32(56) // int32 |  (optional) (default to 1000)
	offset := int32(56) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InscriptionsAPI.GetAccountInscriptions(context.Background(), accountId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InscriptionsAPI.GetAccountInscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInscriptions`: InscriptionBalances
	fmt.Fprintf(os.Stdout, "Response from `InscriptionsAPI.GetAccountInscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 1000]
 **offset** | **int32** |  | [default to 0]

### Return type

[**InscriptionBalances**](InscriptionBalances.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountInscriptionsHistory

> AccountEvents GetAccountInscriptionsHistory(ctx, accountId).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).Limit(limit).Execute()





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
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InscriptionsAPI.GetAccountInscriptionsHistory(context.Background(), accountId).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InscriptionsAPI.GetAccountInscriptionsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInscriptionsHistory`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `InscriptionsAPI.GetAccountInscriptionsHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInscriptionsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **beforeLt** | **int64** | omit this parameter to get last events | 
 **limit** | **int32** |  | [default to 100]

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


## GetAccountInscriptionsHistoryByTicker

> AccountEvents GetAccountInscriptionsHistoryByTicker(ctx, accountId, ticker).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).Limit(limit).Execute()





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
	ticker := "nano" // string | 
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	beforeLt := int64(25758317000002) // int64 | omit this parameter to get last events (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InscriptionsAPI.GetAccountInscriptionsHistoryByTicker(context.Background(), accountId, ticker).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InscriptionsAPI.GetAccountInscriptionsHistoryByTicker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInscriptionsHistoryByTicker`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `InscriptionsAPI.GetAccountInscriptionsHistoryByTicker`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 
**ticker** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInscriptionsHistoryByTickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **beforeLt** | **int64** | omit this parameter to get last events | 
 **limit** | **int32** |  | [default to 100]

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


## GetInscriptionOpTemplate

> GetInscriptionOpTemplate200Response GetInscriptionOpTemplate(ctx).Type_(type_).Operation(operation).Amount(amount).Ticker(ticker).Who(who).Destination(destination).Comment(comment).Execute()





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
	type_ := "ton20" // string | 
	operation := "transfer" // string | 
	amount := "1000000000" // string | 
	ticker := "nano" // string | 
	who := "UQAs87W4yJHlF8mt29ocA4agnMrLsOP69jC1HPyBUjJay7Mg" // string | 
	destination := "destination_example" // string |  (optional)
	comment := "comment_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InscriptionsAPI.GetInscriptionOpTemplate(context.Background()).Type_(type_).Operation(operation).Amount(amount).Ticker(ticker).Who(who).Destination(destination).Comment(comment).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InscriptionsAPI.GetInscriptionOpTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInscriptionOpTemplate`: GetInscriptionOpTemplate200Response
	fmt.Fprintf(os.Stdout, "Response from `InscriptionsAPI.GetInscriptionOpTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInscriptionOpTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** |  | 
 **operation** | **string** |  | 
 **amount** | **string** |  | 
 **ticker** | **string** |  | 
 **who** | **string** |  | 
 **destination** | **string** |  | 
 **comment** | **string** |  | 

### Return type

[**GetInscriptionOpTemplate200Response**](GetInscriptionOpTemplate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

