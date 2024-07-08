# \LiteServerAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllRawShardsInfo**](LiteServerAPI.md#GetAllRawShardsInfo) | **Get** /v2/liteserver/get_all_shards_info/{block_id} | 
[**GetOutMsgQueueSizes**](LiteServerAPI.md#GetOutMsgQueueSizes) | **Get** /v2/liteserver/get_out_msg_queue_sizes | 
[**GetRawAccountState**](LiteServerAPI.md#GetRawAccountState) | **Get** /v2/liteserver/get_account_state/{account_id} | 
[**GetRawBlockProof**](LiteServerAPI.md#GetRawBlockProof) | **Get** /v2/liteserver/get_block_proof | 
[**GetRawBlockchainBlock**](LiteServerAPI.md#GetRawBlockchainBlock) | **Get** /v2/liteserver/get_block/{block_id} | 
[**GetRawBlockchainBlockHeader**](LiteServerAPI.md#GetRawBlockchainBlockHeader) | **Get** /v2/liteserver/get_block_header/{block_id} | 
[**GetRawBlockchainBlockState**](LiteServerAPI.md#GetRawBlockchainBlockState) | **Get** /v2/liteserver/get_state/{block_id} | 
[**GetRawConfig**](LiteServerAPI.md#GetRawConfig) | **Get** /v2/liteserver/get_config_all/{block_id} | 
[**GetRawListBlockTransactions**](LiteServerAPI.md#GetRawListBlockTransactions) | **Get** /v2/liteserver/list_block_transactions/{block_id} | 
[**GetRawMasterchainInfo**](LiteServerAPI.md#GetRawMasterchainInfo) | **Get** /v2/liteserver/get_masterchain_info | 
[**GetRawMasterchainInfoExt**](LiteServerAPI.md#GetRawMasterchainInfoExt) | **Get** /v2/liteserver/get_masterchain_info_ext | 
[**GetRawShardBlockProof**](LiteServerAPI.md#GetRawShardBlockProof) | **Get** /v2/liteserver/get_shard_block_proof/{block_id} | 
[**GetRawShardInfo**](LiteServerAPI.md#GetRawShardInfo) | **Get** /v2/liteserver/get_shard_info/{block_id} | 
[**GetRawTime**](LiteServerAPI.md#GetRawTime) | **Get** /v2/liteserver/get_time | 
[**GetRawTransactions**](LiteServerAPI.md#GetRawTransactions) | **Get** /v2/liteserver/get_transactions/{account_id} | 
[**SendRawMessage**](LiteServerAPI.md#SendRawMessage) | **Post** /v2/liteserver/send_message | 



## GetAllRawShardsInfo

> GetAllRawShardsInfo200Response GetAllRawShardsInfo(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetAllRawShardsInfo(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetAllRawShardsInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllRawShardsInfo`: GetAllRawShardsInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetAllRawShardsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRawShardsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllRawShardsInfo200Response**](GetAllRawShardsInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOutMsgQueueSizes

> GetOutMsgQueueSizes200Response GetOutMsgQueueSizes(ctx).Execute()





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
	resp, r, err := apiClient.LiteServerAPI.GetOutMsgQueueSizes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetOutMsgQueueSizes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOutMsgQueueSizes`: GetOutMsgQueueSizes200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetOutMsgQueueSizes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOutMsgQueueSizesRequest struct via the builder pattern


### Return type

[**GetOutMsgQueueSizes200Response**](GetOutMsgQueueSizes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawAccountState

> GetRawAccountState200Response GetRawAccountState(ctx, accountId).TargetBlock(targetBlock).Execute()





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
	targetBlock := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | target block: (workchain,shard,seqno,root_hash,file_hash) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawAccountState(context.Background(), accountId).TargetBlock(targetBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawAccountState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawAccountState`: GetRawAccountState200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawAccountState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawAccountStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **targetBlock** | **string** | target block: (workchain,shard,seqno,root_hash,file_hash) | 

### Return type

[**GetRawAccountState200Response**](GetRawAccountState200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockProof

> GetRawBlockProof200Response GetRawBlockProof(ctx).KnownBlock(knownBlock).Mode(mode).TargetBlock(targetBlock).Execute()





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
	knownBlock := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | known block: (workchain,shard,seqno,root_hash,file_hash)
	mode := int32(0) // int32 | mode
	targetBlock := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | target block: (workchain,shard,seqno,root_hash,file_hash) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawBlockProof(context.Background()).KnownBlock(knownBlock).Mode(mode).TargetBlock(targetBlock).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawBlockProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockProof`: GetRawBlockProof200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawBlockProof`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **knownBlock** | **string** | known block: (workchain,shard,seqno,root_hash,file_hash) | 
 **mode** | **int32** | mode | 
 **targetBlock** | **string** | target block: (workchain,shard,seqno,root_hash,file_hash) | 

### Return type

[**GetRawBlockProof200Response**](GetRawBlockProof200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockchainBlock

> GetRawBlockchainBlock200Response GetRawBlockchainBlock(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawBlockchainBlock(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawBlockchainBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockchainBlock`: GetRawBlockchainBlock200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawBlockchainBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockchainBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRawBlockchainBlock200Response**](GetRawBlockchainBlock200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockchainBlockHeader

> GetRawBlockchainBlockHeader200Response GetRawBlockchainBlockHeader(ctx, blockId).Mode(mode).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)
	mode := int32(0) // int32 | mode

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawBlockchainBlockHeader(context.Background(), blockId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawBlockchainBlockHeader``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockchainBlockHeader`: GetRawBlockchainBlockHeader200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawBlockchainBlockHeader`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockchainBlockHeaderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **int32** | mode | 

### Return type

[**GetRawBlockchainBlockHeader200Response**](GetRawBlockchainBlockHeader200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockchainBlockState

> GetRawBlockchainBlockState200Response GetRawBlockchainBlockState(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawBlockchainBlockState(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawBlockchainBlockState``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockchainBlockState`: GetRawBlockchainBlockState200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawBlockchainBlockState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockchainBlockStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRawBlockchainBlockState200Response**](GetRawBlockchainBlockState200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawConfig

> GetRawConfig200Response GetRawConfig(ctx, blockId).Mode(mode).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)
	mode := int32(0) // int32 | mode

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawConfig(context.Background(), blockId).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawConfig`: GetRawConfig200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **int32** | mode | 

### Return type

[**GetRawConfig200Response**](GetRawConfig200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawListBlockTransactions

> GetRawListBlockTransactions200Response GetRawListBlockTransactions(ctx, blockId).Mode(mode).Count(count).AccountId(accountId).Lt(lt).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)
	mode := int32(0) // int32 | mode
	count := int32(100) // int32 | count
	accountId := "0:97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | account ID (optional)
	lt := int64(23814011000000) // int64 | lt (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawListBlockTransactions(context.Background(), blockId).Mode(mode).Count(count).AccountId(accountId).Lt(lt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawListBlockTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawListBlockTransactions`: GetRawListBlockTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawListBlockTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawListBlockTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mode** | **int32** | mode | 
 **count** | **int32** | count | 
 **accountId** | **string** | account ID | 
 **lt** | **int64** | lt | 

### Return type

[**GetRawListBlockTransactions200Response**](GetRawListBlockTransactions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawMasterchainInfo

> GetRawMasterchainInfo200Response GetRawMasterchainInfo(ctx).Execute()





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
	resp, r, err := apiClient.LiteServerAPI.GetRawMasterchainInfo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawMasterchainInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawMasterchainInfo`: GetRawMasterchainInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawMasterchainInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawMasterchainInfoRequest struct via the builder pattern


### Return type

[**GetRawMasterchainInfo200Response**](GetRawMasterchainInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawMasterchainInfoExt

> GetRawMasterchainInfoExt200Response GetRawMasterchainInfoExt(ctx).Mode(mode).Execute()





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
	mode := int32(0) // int32 | mode

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawMasterchainInfoExt(context.Background()).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawMasterchainInfoExt``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawMasterchainInfoExt`: GetRawMasterchainInfoExt200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawMasterchainInfoExt`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRawMasterchainInfoExtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **int32** | mode | 

### Return type

[**GetRawMasterchainInfoExt200Response**](GetRawMasterchainInfoExt200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawShardBlockProof

> GetRawShardBlockProof200Response GetRawShardBlockProof(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawShardBlockProof(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawShardBlockProof``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawShardBlockProof`: GetRawShardBlockProof200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawShardBlockProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawShardBlockProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetRawShardBlockProof200Response**](GetRawShardBlockProof200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawShardInfo

> GetRawShardInfo200Response GetRawShardInfo(ctx, blockId).Workchain(workchain).Shard(shard).Exact(exact).Execute()





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
	blockId := "(-1,8000000000000000,4234234,3E575DAB1D25...90D8,47192E5C46C...BB29)" // string | block ID: (workchain,shard,seqno,root_hash,file_hash)
	workchain := int32(1) // int32 | workchain
	shard := int64(1) // int64 | shard
	exact := false // bool | exact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawShardInfo(context.Background(), blockId).Workchain(workchain).Shard(shard).Exact(exact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawShardInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawShardInfo`: GetRawShardInfo200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawShardInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID: (workchain,shard,seqno,root_hash,file_hash) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawShardInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workchain** | **int32** | workchain | 
 **shard** | **int64** | shard | 
 **exact** | **bool** | exact | 

### Return type

[**GetRawShardInfo200Response**](GetRawShardInfo200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTime

> GetRawTime200Response GetRawTime(ctx).Execute()





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
	resp, r, err := apiClient.LiteServerAPI.GetRawTime(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawTime``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawTime`: GetRawTime200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawTimeRequest struct via the builder pattern


### Return type

[**GetRawTime200Response**](GetRawTime200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawTransactions

> GetRawTransactions200Response GetRawTransactions(ctx, accountId).Count(count).Lt(lt).Hash(hash).Execute()





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
	count := int32(100) // int32 | count
	lt := int64(23814011000000) // int64 | lt
	hash := "131D0C65055F04E9C19D687B51BC70F952FD9CA6F02C2801D3B89964A779DF85" // string | hash

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.GetRawTransactions(context.Background(), accountId).Count(count).Lt(lt).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.GetRawTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawTransactions`: GetRawTransactions200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.GetRawTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **count** | **int32** | count | 
 **lt** | **int64** | lt | 
 **hash** | **string** | hash | 

### Return type

[**GetRawTransactions200Response**](GetRawTransactions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendRawMessage

> SendRawMessage200Response SendRawMessage(ctx).SendRawMessageRequest(sendRawMessageRequest).Execute()





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
	sendRawMessageRequest := *openapiclient.NewSendRawMessageRequest("Body_example") // SendRawMessageRequest | Data that is expected

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LiteServerAPI.SendRawMessage(context.Background()).SendRawMessageRequest(sendRawMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LiteServerAPI.SendRawMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SendRawMessage`: SendRawMessage200Response
	fmt.Fprintf(os.Stdout, "Response from `LiteServerAPI.SendRawMessage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendRawMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendRawMessageRequest** | [**SendRawMessageRequest**](SendRawMessageRequest.md) | Data that is expected | 

### Return type

[**SendRawMessage200Response**](SendRawMessage200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

