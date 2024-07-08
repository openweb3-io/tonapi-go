# \EventsAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /v2/events/{event_id} | 



## GetEvent

> Event GetEvent(ctx, eventId).AcceptLanguage(acceptLanguage).Execute()





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
	resp, r, err := apiClient.EventsAPI.GetEvent(context.Background(), eventId).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | event ID or transaction hash in hex (without 0x) or base64url format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


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

