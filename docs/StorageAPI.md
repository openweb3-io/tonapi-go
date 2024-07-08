# \StorageAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStorageProviders**](StorageAPI.md#GetStorageProviders) | **Get** /v2/storage/providers | 



## GetStorageProviders

> GetStorageProviders200Response GetStorageProviders(ctx).Execute()





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
	resp, r, err := apiClient.StorageAPI.GetStorageProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StorageAPI.GetStorageProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStorageProviders`: GetStorageProviders200Response
	fmt.Fprintf(os.Stdout, "Response from `StorageAPI.GetStorageProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageProvidersRequest struct via the builder pattern


### Return type

[**GetStorageProviders200Response**](GetStorageProviders200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

