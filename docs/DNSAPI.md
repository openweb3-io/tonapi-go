# \DNSAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsResolve**](DNSAPI.md#DnsResolve) | **Get** /v2/dns/{domain_name}/resolve | 
[**GetAllAuctions**](DNSAPI.md#GetAllAuctions) | **Get** /v2/dns/auctions | 
[**GetDnsInfo**](DNSAPI.md#GetDnsInfo) | **Get** /v2/dns/{domain_name} | 
[**GetDomainBids**](DNSAPI.md#GetDomainBids) | **Get** /v2/dns/{domain_name}/bids | 



## DnsResolve

> DnsRecord DnsResolve(ctx, domainName).Execute()





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
	domainName := "wallet.ton" // string | domain name with .ton or .t.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.DnsResolve(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.DnsResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DnsResolve`: DnsRecord
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.DnsResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | domain name with .ton or .t.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiDnsResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DnsRecord**](DnsRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAuctions

> Auctions GetAllAuctions(ctx).Tld(tld).Execute()





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
	tld := "ton" // string | domain filter for current auctions \"ton\" or \"t.me\" (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetAllAuctions(context.Background()).Tld(tld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetAllAuctions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAuctions`: Auctions
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetAllAuctions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAuctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tld** | **string** | domain filter for current auctions \&quot;ton\&quot; or \&quot;t.me\&quot; | 

### Return type

[**Auctions**](Auctions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDnsInfo

> DomainInfo GetDnsInfo(ctx, domainName).Execute()





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
	domainName := "wallet.ton" // string | domain name with .ton or .t.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetDnsInfo(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDnsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDnsInfo`: DomainInfo
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDnsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | domain name with .ton or .t.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDnsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainInfo**](DomainInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDomainBids

> DomainBids GetDomainBids(ctx, domainName).Execute()





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
	domainName := "wallet.ton" // string | domain name with .ton or .t.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DNSAPI.GetDomainBids(context.Background(), domainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DNSAPI.GetDomainBids``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDomainBids`: DomainBids
	fmt.Fprintf(os.Stdout, "Response from `DNSAPI.GetDomainBids`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | domain name with .ton or .t.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainBidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DomainBids**](DomainBids.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

