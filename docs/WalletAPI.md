# \WalletAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountSeqno**](WalletAPI.md#GetAccountSeqno) | **Get** /v2/wallet/{account_id}/seqno | 
[**GetWalletBackup**](WalletAPI.md#GetWalletBackup) | **Get** /v2/wallet/backup | 
[**GetWalletsByPublicKey**](WalletAPI.md#GetWalletsByPublicKey) | **Get** /v2/pubkeys/{public_key}/wallets | 
[**SetWalletBackup**](WalletAPI.md#SetWalletBackup) | **Put** /v2/wallet/backup | 
[**TonConnectProof**](WalletAPI.md#TonConnectProof) | **Post** /v2/wallet/auth/proof | 



## GetAccountSeqno

> Seqno GetAccountSeqno(ctx, accountId).Execute()





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
	resp, r, err := apiClient.WalletAPI.GetAccountSeqno(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetAccountSeqno``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountSeqno`: Seqno
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetAccountSeqno`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountSeqnoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Seqno**](Seqno.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletBackup

> GetWalletBackup200Response GetWalletBackup(ctx).XTonConnectAuth(xTonConnectAuth).Execute()





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
	xTonConnectAuth := "xTonConnectAuth_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetWalletBackup(context.Background()).XTonConnectAuth(xTonConnectAuth).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetWalletBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletBackup`: GetWalletBackup200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetWalletBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTonConnectAuth** | **string** |  | 

### Return type

[**GetWalletBackup200Response**](GetWalletBackup200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWalletsByPublicKey

> Accounts GetWalletsByPublicKey(ctx, publicKey).Execute()





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
	publicKey := "NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQ3..." // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.GetWalletsByPublicKey(context.Background(), publicKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.GetWalletsByPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWalletsByPublicKey`: Accounts
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.GetWalletsByPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletsByPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Accounts**](Accounts.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWalletBackup

> SetWalletBackup(ctx).XTonConnectAuth(xTonConnectAuth).Body(body).Execute()





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
	xTonConnectAuth := "xTonConnectAuth_example" // string | 
	body := os.NewFile(1234, "some_file") // *os.File | Information for saving backup

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WalletAPI.SetWalletBackup(context.Background()).XTonConnectAuth(xTonConnectAuth).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.SetWalletBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetWalletBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xTonConnectAuth** | **string** |  | 
 **body** | ***os.File** | Information for saving backup | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/octet-stream
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TonConnectProof

> TonConnectProof200Response TonConnectProof(ctx).TonConnectProofRequest(tonConnectProofRequest).Execute()





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
	tonConnectProofRequest := *openapiclient.NewTonConnectProofRequest("0:97146a46acc2654y27947f14c4a4b14273e954f78bc017790b41208b0043200b", *openapiclient.NewTonConnectProofRequestProof(int64(1678275313), *openapiclient.NewTonConnectProofRequestProofDomain("Value_example"), "Signature_example", "84jHVNLQmZsAAAAAZB0Zryi2wqVJI-KaKNXOvCijEi46YyYzkaSHyJrMPBMOkVZa")) // TonConnectProofRequest | Data that is expected from TON Connect

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WalletAPI.TonConnectProof(context.Background()).TonConnectProofRequest(tonConnectProofRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WalletAPI.TonConnectProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TonConnectProof`: TonConnectProof200Response
	fmt.Fprintf(os.Stdout, "Response from `WalletAPI.TonConnectProof`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTonConnectProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tonConnectProofRequest** | [**TonConnectProofRequest**](TonConnectProofRequest.md) | Data that is expected from TON Connect | 

### Return type

[**TonConnectProof200Response**](TonConnectProof200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

