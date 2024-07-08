# \ConnectAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountInfoByStateInit**](ConnectAPI.md#GetAccountInfoByStateInit) | **Post** /v2/tonconnect/stateinit | 
[**GetTonConnectPayload**](ConnectAPI.md#GetTonConnectPayload) | **Get** /v2/tonconnect/payload | 



## GetAccountInfoByStateInit

> AccountInfoByStateInit GetAccountInfoByStateInit(ctx).GetAccountInfoByStateInitRequest(getAccountInfoByStateInitRequest).Execute()





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
	getAccountInfoByStateInitRequest := *openapiclient.NewGetAccountInfoByStateInitRequest("StateInit_example") // GetAccountInfoByStateInitRequest | Data that is expected

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetAccountInfoByStateInit(context.Background()).GetAccountInfoByStateInitRequest(getAccountInfoByStateInitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetAccountInfoByStateInit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountInfoByStateInit`: AccountInfoByStateInit
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetAccountInfoByStateInit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountInfoByStateInitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAccountInfoByStateInitRequest** | [**GetAccountInfoByStateInitRequest**](GetAccountInfoByStateInitRequest.md) | Data that is expected | 

### Return type

[**AccountInfoByStateInit**](AccountInfoByStateInit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTonConnectPayload

> GetTonConnectPayload200Response GetTonConnectPayload(ctx).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConnectAPI.GetTonConnectPayload(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConnectAPI.GetTonConnectPayload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTonConnectPayload`: GetTonConnectPayload200Response
	fmt.Fprintf(os.Stdout, "Response from `ConnectAPI.GetTonConnectPayload`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTonConnectPayloadRequest struct via the builder pattern


### Return type

[**GetTonConnectPayload200Response**](GetTonConnectPayload200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

