# \GaslessAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GaslessConfig**](GaslessAPI.md#GaslessConfig) | **Get** /v2/gasless/config | 
[**GaslessEstimate**](GaslessAPI.md#GaslessEstimate) | **Post** /v2/gasless/estimate/{master_id} | 
[**GaslessSend**](GaslessAPI.md#GaslessSend) | **Post** /v2/gasless/send | 



## GaslessConfig

> GaslessConfig GaslessConfig(ctx).Execute()





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
	resp, r, err := apiClient.GaslessAPI.GaslessConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GaslessAPI.GaslessConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GaslessConfig`: GaslessConfig
	fmt.Fprintf(os.Stdout, "Response from `GaslessAPI.GaslessConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGaslessConfigRequest struct via the builder pattern


### Return type

[**GaslessConfig**](GaslessConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaslessEstimate

> SignRawParams GaslessEstimate(ctx, masterId).GaslessEstimateRequest(gaslessEstimateRequest).Execute()





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
	masterId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | jetton to pay commission
	gaslessEstimateRequest := *openapiclient.NewGaslessEstimateRequest("WalletAddress_example", "WalletPublicKey_example", []openapiclient.GaslessEstimateRequestMessagesInner{*openapiclient.NewGaslessEstimateRequestMessagesInner("B5EE9C7201010101001100001D00048656C6C6F2C20776F726C64218")}) // GaslessEstimateRequest | bag-of-cells serialized to base64

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GaslessAPI.GaslessEstimate(context.Background(), masterId).GaslessEstimateRequest(gaslessEstimateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GaslessAPI.GaslessEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GaslessEstimate`: SignRawParams
	fmt.Fprintf(os.Stdout, "Response from `GaslessAPI.GaslessEstimate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterId** | **string** | jetton to pay commission | 

### Other Parameters

Other parameters are passed through a pointer to a apiGaslessEstimateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **gaslessEstimateRequest** | [**GaslessEstimateRequest**](GaslessEstimateRequest.md) | bag-of-cells serialized to base64 | 

### Return type

[**SignRawParams**](SignRawParams.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GaslessSend

> GaslessSend(ctx).GaslessSendRequest(gaslessSendRequest).Execute()



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
	gaslessSendRequest := *openapiclient.NewGaslessSendRequest("WalletPublicKey_example", "te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // GaslessSendRequest | bag-of-cells serialized to base64

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GaslessAPI.GaslessSend(context.Background()).GaslessSendRequest(gaslessSendRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GaslessAPI.GaslessSend``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGaslessSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gaslessSendRequest** | [**GaslessSendRequest**](GaslessSendRequest.md) | bag-of-cells serialized to base64 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

