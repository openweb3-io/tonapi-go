# \JettonsAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetJettonHolders**](JettonsAPI.md#GetJettonHolders) | **Get** /v2/jettons/{account_id}/holders | 
[**GetJettonInfo**](JettonsAPI.md#GetJettonInfo) | **Get** /v2/jettons/{account_id} | 
[**GetJettons**](JettonsAPI.md#GetJettons) | **Get** /v2/jettons | 
[**GetJettonsEvents**](JettonsAPI.md#GetJettonsEvents) | **Get** /v2/events/{event_id}/jettons | 



## GetJettonHolders

> JettonHolders GetJettonHolders(ctx, accountId).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.JettonsAPI.GetJettonHolders(context.Background(), accountId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JettonsAPI.GetJettonHolders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJettonHolders`: JettonHolders
	fmt.Fprintf(os.Stdout, "Response from `JettonsAPI.GetJettonHolders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJettonHoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 1000]
 **offset** | **int32** |  | [default to 0]

### Return type

[**JettonHolders**](JettonHolders.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJettonInfo

> JettonInfo GetJettonInfo(ctx, accountId).Execute()





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
	resp, r, err := apiClient.JettonsAPI.GetJettonInfo(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JettonsAPI.GetJettonInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJettonInfo`: JettonInfo
	fmt.Fprintf(os.Stdout, "Response from `JettonsAPI.GetJettonInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJettonInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JettonInfo**](JettonInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJettons

> Jettons GetJettons(ctx).Limit(limit).Offset(offset).Execute()





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
	limit := int32(15) // int32 |  (optional) (default to 100)
	offset := int32(10) // int32 |  (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JettonsAPI.GetJettons(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JettonsAPI.GetJettons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJettons`: Jettons
	fmt.Fprintf(os.Stdout, "Response from `JettonsAPI.GetJettons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJettonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**Jettons**](Jettons.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJettonsEvents

> Event GetJettonsEvents(ctx, eventId).AcceptLanguage(acceptLanguage).Execute()





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
	eventId := "97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | event ID or transaction hash in hex (without 0x) or base64url format
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JettonsAPI.GetJettonsEvents(context.Background(), eventId).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JettonsAPI.GetJettonsEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJettonsEvents`: Event
	fmt.Fprintf(os.Stdout, "Response from `JettonsAPI.GetJettonsEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | event ID or transaction hash in hex (without 0x) or base64url format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJettonsEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

