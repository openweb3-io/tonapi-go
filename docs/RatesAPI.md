# \RatesAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetChartRates**](RatesAPI.md#GetChartRates) | **Get** /v2/rates/chart | 
[**GetMarketsRates**](RatesAPI.md#GetMarketsRates) | **Get** /v2/rates/markets | 
[**GetRates**](RatesAPI.md#GetRates) | **Get** /v2/rates | 



## GetChartRates

> GetChartRates200Response GetChartRates(ctx).Token(token).Currency(currency).StartDate(startDate).EndDate(endDate).PointsCount(pointsCount).Execute()





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
	token := "token_example" // string | accept jetton master address
	currency := "usd" // string |  (optional)
	startDate := int64(1668436763) // int64 |  (optional)
	endDate := int64(1668436763) // int64 |  (optional)
	pointsCount := int32(56) // int32 |  (optional) (default to 200)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.GetChartRates(context.Background()).Token(token).Currency(currency).StartDate(startDate).EndDate(endDate).PointsCount(pointsCount).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetChartRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChartRates`: GetChartRates200Response
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetChartRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChartRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | accept jetton master address | 
 **currency** | **string** |  | 
 **startDate** | **int64** |  | 
 **endDate** | **int64** |  | 
 **pointsCount** | **int32** |  | [default to 200]

### Return type

[**GetChartRates200Response**](GetChartRates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMarketsRates

> GetMarketsRates200Response GetMarketsRates(ctx).Execute()





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
	resp, r, err := apiClient.RatesAPI.GetMarketsRates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetMarketsRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMarketsRates`: GetMarketsRates200Response
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetMarketsRates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketsRatesRequest struct via the builder pattern


### Return type

[**GetMarketsRates200Response**](GetMarketsRates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRates

> GetRates200Response GetRates(ctx).Tokens(tokens).Currencies(currencies).Execute()





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
	tokens := []string{"Inner_example"} // []string | accept ton and jetton master addresses, separated by commas
	currencies := []string{"Inner_example"} // []string | accept ton and all possible fiat currencies, separated by commas

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RatesAPI.GetRates(context.Background()).Tokens(tokens).Currencies(currencies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RatesAPI.GetRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRates`: GetRates200Response
	fmt.Fprintf(os.Stdout, "Response from `RatesAPI.GetRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokens** | **[]string** | accept ton and jetton master addresses, separated by commas | 
 **currencies** | **[]string** | accept ton and all possible fiat currencies, separated by commas | 

### Return type

[**GetRates200Response**](GetRates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

