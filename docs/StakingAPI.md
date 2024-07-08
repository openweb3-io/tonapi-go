# \StakingAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountNominatorsPools**](StakingAPI.md#GetAccountNominatorsPools) | **Get** /v2/staking/nominator/{account_id}/pools | 
[**GetStakingPoolHistory**](StakingAPI.md#GetStakingPoolHistory) | **Get** /v2/staking/pool/{account_id}/history | 
[**GetStakingPoolInfo**](StakingAPI.md#GetStakingPoolInfo) | **Get** /v2/staking/pool/{account_id} | 
[**GetStakingPools**](StakingAPI.md#GetStakingPools) | **Get** /v2/staking/pools | 



## GetAccountNominatorsPools

> AccountStaking GetAccountNominatorsPools(ctx, accountId).Execute()





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
	resp, r, err := apiClient.StakingAPI.GetAccountNominatorsPools(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetAccountNominatorsPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNominatorsPools`: AccountStaking
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetAccountNominatorsPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNominatorsPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccountStaking**](AccountStaking.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingPoolHistory

> GetStakingPoolHistory200Response GetStakingPoolHistory(ctx, accountId).Execute()





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
	resp, r, err := apiClient.StakingAPI.GetStakingPoolHistory(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetStakingPoolHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingPoolHistory`: GetStakingPoolHistory200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetStakingPoolHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingPoolHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetStakingPoolHistory200Response**](GetStakingPoolHistory200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingPoolInfo

> GetStakingPoolInfo200Response GetStakingPoolInfo(ctx, accountId).AcceptLanguage(acceptLanguage).Execute()





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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetStakingPoolInfo(context.Background(), accountId).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetStakingPoolInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingPoolInfo`: GetStakingPoolInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetStakingPoolInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingPoolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]

### Return type

[**GetStakingPoolInfo200Response**](GetStakingPoolInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStakingPools

> GetStakingPools200Response GetStakingPools(ctx).AvailableFor(availableFor).IncludeUnverified(includeUnverified).AcceptLanguage(acceptLanguage).Execute()





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
	availableFor := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID (optional)
	includeUnverified := false // bool | return also pools not from white list - just compatible by interfaces (maybe dangerous!) (optional)
	acceptLanguage := "ru-RU,ru;q=0.5" // string |  (optional) (default to "en")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StakingAPI.GetStakingPools(context.Background()).AvailableFor(availableFor).IncludeUnverified(includeUnverified).AcceptLanguage(acceptLanguage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StakingAPI.GetStakingPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStakingPools`: GetStakingPools200Response
	fmt.Fprintf(os.Stdout, "Response from `StakingAPI.GetStakingPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStakingPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **availableFor** | **string** | account ID | 
 **includeUnverified** | **bool** | return also pools not from white list - just compatible by interfaces (maybe dangerous!) | 
 **acceptLanguage** | **string** |  | [default to &quot;en&quot;]

### Return type

[**GetStakingPools200Response**](GetStakingPools200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

