# \TracesAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrace**](TracesAPI.md#GetTrace) | **Get** /v2/traces/{trace_id} | 



## GetTrace

> Trace GetTrace(ctx, traceId).Execute()





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
	traceId := "97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | trace ID or transaction hash in hex (without 0x) or base64url format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TracesAPI.GetTrace(context.Background(), traceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TracesAPI.GetTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrace`: Trace
	fmt.Fprintf(os.Stdout, "Response from `TracesAPI.GetTrace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**traceId** | **string** | trace ID or transaction hash in hex (without 0x) or base64url format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trace**](Trace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

