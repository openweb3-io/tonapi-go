# \NFTAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountNftHistory**](NFTAPI.md#GetAccountNftHistory) | **Get** /v2/accounts/{account_id}/nfts/history | 
[**GetItemsFromCollection**](NFTAPI.md#GetItemsFromCollection) | **Get** /v2/nfts/collections/{account_id}/items | 
[**GetNftCollection**](NFTAPI.md#GetNftCollection) | **Get** /v2/nfts/collections/{account_id} | 
[**GetNftCollections**](NFTAPI.md#GetNftCollections) | **Get** /v2/nfts/collections | 
[**GetNftHistoryByID**](NFTAPI.md#GetNftHistoryByID) | **Get** /v2/nfts/{account_id}/history | 
[**GetNftItemByAddress**](NFTAPI.md#GetNftItemByAddress) | **Get** /v2/nfts/{account_id} | 
[**GetNftItemsByAddresses**](NFTAPI.md#GetNftItemsByAddresses) | **Post** /v2/nfts/_bulk | 



## GetAccountNftHistory

> AccountEvents GetAccountNftHistory(ctx, accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetAccountNftHistory(context.Background(), accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetAccountNftHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNftHistory`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetAccountNftHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNftHistoryRequest struct via the builder pattern


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


## GetItemsFromCollection

> NftItems GetItemsFromCollection(ctx, accountId).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetItemsFromCollection(context.Background(), accountId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetItemsFromCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetItemsFromCollection`: NftItems
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetItemsFromCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetItemsFromCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 1000]
 **offset** | **int32** |  | [default to 0]

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


## GetNftCollection

> NftCollection GetNftCollection(ctx, accountId).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetNftCollection(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetNftCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftCollection`: NftCollection
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetNftCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNftCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NftCollection**](NftCollection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftCollections

> NftCollections GetNftCollections(ctx).Limit(limit).Offset(offset).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetNftCollections(context.Background()).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetNftCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftCollections`: NftCollections
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetNftCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]
 **offset** | **int32** |  | [default to 0]

### Return type

[**NftCollections**](NftCollections.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftHistoryByID

> AccountEvents GetNftHistoryByID(ctx, accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetNftHistoryByID(context.Background(), accountId).Limit(limit).AcceptLanguage(acceptLanguage).BeforeLt(beforeLt).StartDate(startDate).EndDate(endDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetNftHistoryByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftHistoryByID`: AccountEvents
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetNftHistoryByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNftHistoryByIDRequest struct via the builder pattern


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


## GetNftItemByAddress

> NftItem GetNftItemByAddress(ctx, accountId).Execute()





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
	resp, r, err := apiClient.NFTAPI.GetNftItemByAddress(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetNftItemByAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftItemByAddress`: NftItem
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetNftItemByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNftItemByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NftItem**](NftItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftItemsByAddresses

> NftItems GetNftItemsByAddresses(ctx).GetAccountsRequest(getAccountsRequest).Execute()





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
	getAccountsRequest := *openapiclient.NewGetAccountsRequest([]string{"0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621"}) // GetAccountsRequest | a list of account ids (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NFTAPI.GetNftItemsByAddresses(context.Background()).GetAccountsRequest(getAccountsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NFTAPI.GetNftItemsByAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNftItemsByAddresses`: NftItems
	fmt.Fprintf(os.Stdout, "Response from `NFTAPI.GetNftItemsByAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftItemsByAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getAccountsRequest** | [**GetAccountsRequest**](GetAccountsRequest.md) | a list of account ids | 

### Return type

[**NftItems**](NftItems.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

