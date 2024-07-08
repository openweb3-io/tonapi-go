# \MultisigAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMultisigAccount**](MultisigAPI.md#GetMultisigAccount) | **Get** /v2/multisig/{account_id} | 



## GetMultisigAccount

> Multisig GetMultisigAccount(ctx, accountId).Execute()





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
	resp, r, err := apiClient.MultisigAPI.GetMultisigAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultisigAPI.GetMultisigAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMultisigAccount`: Multisig
	fmt.Fprintf(os.Stdout, "Response from `MultisigAPI.GetMultisigAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultisigAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Multisig**](Multisig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

