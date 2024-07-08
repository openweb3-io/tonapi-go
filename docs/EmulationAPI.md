# \EmulationAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecodeMessage**](EmulationAPI.md#DecodeMessage) | **Post** /v2/message/decode | 
[**EmulateMessageToAccountEvent**](EmulationAPI.md#EmulateMessageToAccountEvent) | **Post** /v2/accounts/{account_id}/events/emulate | 
[**EmulateMessageToEvent**](EmulationAPI.md#EmulateMessageToEvent) | **Post** /v2/events/emulate | 
[**EmulateMessageToTrace**](EmulationAPI.md#EmulateMessageToTrace) | **Post** /v2/traces/emulate | 
[**EmulateMessageToWallet**](EmulationAPI.md#EmulateMessageToWallet) | **Post** /v2/wallet/emulate | 



## DecodeMessage

> DecodedMessage DecodeMessage(ctx).DecodeMessageRequest(decodeMessageRequest).Execute()





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
	decodeMessageRequest := *openapiclient.NewDecodeMessageRequest("te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // DecodeMessageRequest | bag-of-cells serialized to base64

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmulationAPI.DecodeMessage(context.Background()).DecodeMessageRequest(decodeMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.DecodeMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecodeMessage`: DecodedMessage
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.DecodeMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecodeMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decodeMessageRequest** | [**DecodeMessageRequest**](DecodeMessageRequest.md) | bag-of-cells serialized to base64 | 

### Return type

[**DecodedMessage**](DecodedMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmulateMessageToAccountEvent

> AccountEvent EmulateMessageToAccountEvent(ctx, accountId).DecodeMessageRequest(decodeMessageRequest).AcceptLanguage(acceptLanguage).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()





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
	decodeMessageRequest := *openapiclient.NewDecodeMessageRequest("te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // DecodeMessageRequest | bag-of-cells serialized to base64
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	ignoreSignatureCheck := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmulationAPI.EmulateMessageToAccountEvent(context.Background(), accountId).DecodeMessageRequest(decodeMessageRequest).AcceptLanguage(acceptLanguage).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.EmulateMessageToAccountEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmulateMessageToAccountEvent`: AccountEvent
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.EmulateMessageToAccountEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEmulateMessageToAccountEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decodeMessageRequest** | [**DecodeMessageRequest**](DecodeMessageRequest.md) | bag-of-cells serialized to base64 | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **ignoreSignatureCheck** | **bool** |  | 

### Return type

[**AccountEvent**](AccountEvent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmulateMessageToEvent

> Event EmulateMessageToEvent(ctx).DecodeMessageRequest(decodeMessageRequest).AcceptLanguage(acceptLanguage).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()





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
	decodeMessageRequest := *openapiclient.NewDecodeMessageRequest("te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // DecodeMessageRequest | bag-of-cells serialized to base64
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")
	ignoreSignatureCheck := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmulationAPI.EmulateMessageToEvent(context.Background()).DecodeMessageRequest(decodeMessageRequest).AcceptLanguage(acceptLanguage).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.EmulateMessageToEvent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmulateMessageToEvent`: Event
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.EmulateMessageToEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmulateMessageToEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decodeMessageRequest** | [**DecodeMessageRequest**](DecodeMessageRequest.md) | bag-of-cells serialized to base64 | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]
 **ignoreSignatureCheck** | **bool** |  | 

### Return type

[**Event**](Event.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmulateMessageToTrace

> Trace EmulateMessageToTrace(ctx).DecodeMessageRequest(decodeMessageRequest).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()





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
	decodeMessageRequest := *openapiclient.NewDecodeMessageRequest("te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // DecodeMessageRequest | bag-of-cells serialized to base64
	ignoreSignatureCheck := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmulationAPI.EmulateMessageToTrace(context.Background()).DecodeMessageRequest(decodeMessageRequest).IgnoreSignatureCheck(ignoreSignatureCheck).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.EmulateMessageToTrace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmulateMessageToTrace`: Trace
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.EmulateMessageToTrace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmulateMessageToTraceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decodeMessageRequest** | [**DecodeMessageRequest**](DecodeMessageRequest.md) | bag-of-cells serialized to base64 | 
 **ignoreSignatureCheck** | **bool** |  | 

### Return type

[**Trace**](Trace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmulateMessageToWallet

> MessageConsequences EmulateMessageToWallet(ctx).EmulateMessageToWalletRequest(emulateMessageToWalletRequest).AcceptLanguage(acceptLanguage).Execute()





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
	emulateMessageToWalletRequest := *openapiclient.NewEmulateMessageToWalletRequest("te6ccgECBQEAARUAAkWIAWTtae+KgtbrX26Bep8JSq8lFLfGOoyGR/xwdjfvpvEaHg") // EmulateMessageToWalletRequest | bag-of-cells serialized to base64 and additional parameters to configure emulation
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmulationAPI.EmulateMessageToWallet(context.Background()).EmulateMessageToWalletRequest(emulateMessageToWalletRequest).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmulationAPI.EmulateMessageToWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EmulateMessageToWallet`: MessageConsequences
	fmt.Fprintf(os.Stdout, "Response from `EmulationAPI.EmulateMessageToWallet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmulateMessageToWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emulateMessageToWalletRequest** | [**EmulateMessageToWalletRequest**](EmulateMessageToWalletRequest.md) | bag-of-cells serialized to base64 and additional parameters to configure emulation | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]

### Return type

[**MessageConsequences**](MessageConsequences.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

