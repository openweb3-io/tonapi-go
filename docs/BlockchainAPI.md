# \BlockchainAPI

All URIs are relative to *https://tonapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BlockchainAccountInspect**](BlockchainAPI.md#BlockchainAccountInspect) | **Get** /v2/blockchain/accounts/{account_id}/inspect | 
[**ExecGetMethodForBlockchainAccount**](BlockchainAPI.md#ExecGetMethodForBlockchainAccount) | **Get** /v2/blockchain/accounts/{account_id}/methods/{method_name} | 
[**GetBlockchainAccountTransactions**](BlockchainAPI.md#GetBlockchainAccountTransactions) | **Get** /v2/blockchain/accounts/{account_id}/transactions | 
[**GetBlockchainBlock**](BlockchainAPI.md#GetBlockchainBlock) | **Get** /v2/blockchain/blocks/{block_id} | 
[**GetBlockchainBlockTransactions**](BlockchainAPI.md#GetBlockchainBlockTransactions) | **Get** /v2/blockchain/blocks/{block_id}/transactions | 
[**GetBlockchainConfig**](BlockchainAPI.md#GetBlockchainConfig) | **Get** /v2/blockchain/config | 
[**GetBlockchainConfigFromBlock**](BlockchainAPI.md#GetBlockchainConfigFromBlock) | **Get** /v2/blockchain/masterchain/{masterchain_seqno}/config | 
[**GetBlockchainMasterchainBlocks**](BlockchainAPI.md#GetBlockchainMasterchainBlocks) | **Get** /v2/blockchain/masterchain/{masterchain_seqno}/blocks | 
[**GetBlockchainMasterchainHead**](BlockchainAPI.md#GetBlockchainMasterchainHead) | **Get** /v2/blockchain/masterchain-head | 
[**GetBlockchainMasterchainShards**](BlockchainAPI.md#GetBlockchainMasterchainShards) | **Get** /v2/blockchain/masterchain/{masterchain_seqno}/shards | 
[**GetBlockchainMasterchainTransactions**](BlockchainAPI.md#GetBlockchainMasterchainTransactions) | **Get** /v2/blockchain/masterchain/{masterchain_seqno}/transactions | 
[**GetBlockchainRawAccount**](BlockchainAPI.md#GetBlockchainRawAccount) | **Get** /v2/blockchain/accounts/{account_id} | 
[**GetBlockchainTransaction**](BlockchainAPI.md#GetBlockchainTransaction) | **Get** /v2/blockchain/transactions/{transaction_id} | 
[**GetBlockchainTransactionByMessageHash**](BlockchainAPI.md#GetBlockchainTransactionByMessageHash) | **Get** /v2/blockchain/messages/{msg_id}/transaction | 
[**GetBlockchainValidators**](BlockchainAPI.md#GetBlockchainValidators) | **Get** /v2/blockchain/validators | 
[**GetRawBlockchainConfig**](BlockchainAPI.md#GetRawBlockchainConfig) | **Get** /v2/blockchain/config/raw | 
[**GetRawBlockchainConfigFromBlock**](BlockchainAPI.md#GetRawBlockchainConfigFromBlock) | **Get** /v2/blockchain/masterchain/{masterchain_seqno}/config/raw | 
[**GetReducedBlockchainBlocks**](BlockchainAPI.md#GetReducedBlockchainBlocks) | **Get** /v2/blockchain/reduced/blocks | 
[**SendBlockchainMessage**](BlockchainAPI.md#SendBlockchainMessage) | **Post** /v2/blockchain/message | 
[**Status**](BlockchainAPI.md#Status) | **Get** /v2/status | 



## BlockchainAccountInspect

> BlockchainAccountInspect BlockchainAccountInspect(ctx, accountId).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.BlockchainAccountInspect(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.BlockchainAccountInspect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BlockchainAccountInspect`: BlockchainAccountInspect
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.BlockchainAccountInspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBlockchainAccountInspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainAccountInspect**](BlockchainAccountInspect.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecGetMethodForBlockchainAccount

> MethodExecutionResult ExecGetMethodForBlockchainAccount(ctx, accountId, methodName).Args(args).Execute()





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
	methodName := "get_wallet_address" // string | contract get method name
	args := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.ExecGetMethodForBlockchainAccount(context.Background(), accountId, methodName).Args(args).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.ExecGetMethodForBlockchainAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExecGetMethodForBlockchainAccount`: MethodExecutionResult
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.ExecGetMethodForBlockchainAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 
**methodName** | **string** | contract get method name | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecGetMethodForBlockchainAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **args** | **[]string** |  | 

### Return type

[**MethodExecutionResult**](MethodExecutionResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainAccountTransactions

> Transactions GetBlockchainAccountTransactions(ctx, accountId).AfterLt(afterLt).BeforeLt(beforeLt).Limit(limit).SortOrder(sortOrder).Execute()





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
	afterLt := int64(39787624000003) // int64 | omit this parameter to get last transactions (optional)
	beforeLt := int64(39787624000003) // int64 | omit this parameter to get last transactions (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	sortOrder := "sortOrder_example" // string |  (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainAccountTransactions(context.Background(), accountId).AfterLt(afterLt).BeforeLt(beforeLt).Limit(limit).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainAccountTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainAccountTransactions`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainAccountTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainAccountTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **afterLt** | **int64** | omit this parameter to get last transactions | 
 **beforeLt** | **int64** | omit this parameter to get last transactions | 
 **limit** | **int32** |  | [default to 100]
 **sortOrder** | **string** |  | [default to &quot;desc&quot;]

### Return type

[**Transactions**](Transactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainBlock

> BlockchainBlock GetBlockchainBlock(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234)" // string | block ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainBlock(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainBlock`: BlockchainBlock
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainBlock**](BlockchainBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainBlockTransactions

> Transactions GetBlockchainBlockTransactions(ctx, blockId).Execute()





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
	blockId := "(-1,8000000000000000,4234234)" // string | block ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainBlockTransactions(context.Background(), blockId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainBlockTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainBlockTransactions`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainBlockTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**blockId** | **string** | block ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainBlockTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transactions**](Transactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainConfig

> BlockchainConfig GetBlockchainConfig(ctx).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainConfig`: BlockchainConfig
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainConfigRequest struct via the builder pattern


### Return type

[**BlockchainConfig**](BlockchainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainConfigFromBlock

> BlockchainConfig GetBlockchainConfigFromBlock(ctx, masterchainSeqno).Execute()





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
	masterchainSeqno := int32(123456) // int32 | masterchain block seqno

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainConfigFromBlock(context.Background(), masterchainSeqno).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainConfigFromBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainConfigFromBlock`: BlockchainConfig
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainConfigFromBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterchainSeqno** | **int32** | masterchain block seqno | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainConfigFromBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainConfig**](BlockchainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainMasterchainBlocks

> BlockchainBlocks GetBlockchainMasterchainBlocks(ctx, masterchainSeqno).Execute()





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
	masterchainSeqno := int32(123456) // int32 | masterchain block seqno

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainMasterchainBlocks(context.Background(), masterchainSeqno).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainMasterchainBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainMasterchainBlocks`: BlockchainBlocks
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainMasterchainBlocks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterchainSeqno** | **int32** | masterchain block seqno | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainMasterchainBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainBlocks**](BlockchainBlocks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainMasterchainHead

> BlockchainBlock GetBlockchainMasterchainHead(ctx).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainMasterchainHead(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainMasterchainHead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainMasterchainHead`: BlockchainBlock
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainMasterchainHead`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainMasterchainHeadRequest struct via the builder pattern


### Return type

[**BlockchainBlock**](BlockchainBlock.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainMasterchainShards

> BlockchainBlockShards GetBlockchainMasterchainShards(ctx, masterchainSeqno).Execute()





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
	masterchainSeqno := int32(123456) // int32 | masterchain block seqno

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainMasterchainShards(context.Background(), masterchainSeqno).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainMasterchainShards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainMasterchainShards`: BlockchainBlockShards
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainMasterchainShards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterchainSeqno** | **int32** | masterchain block seqno | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainMasterchainShardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainBlockShards**](BlockchainBlockShards.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainMasterchainTransactions

> Transactions GetBlockchainMasterchainTransactions(ctx, masterchainSeqno).Execute()





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
	masterchainSeqno := int32(123456) // int32 | masterchain block seqno

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainMasterchainTransactions(context.Background(), masterchainSeqno).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainMasterchainTransactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainMasterchainTransactions`: Transactions
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainMasterchainTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterchainSeqno** | **int32** | masterchain block seqno | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainMasterchainTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transactions**](Transactions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainRawAccount

> BlockchainRawAccount GetBlockchainRawAccount(ctx, accountId).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainRawAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainRawAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainRawAccount`: BlockchainRawAccount
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainRawAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | account ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainRawAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BlockchainRawAccount**](BlockchainRawAccount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainTransaction

> Transaction GetBlockchainTransaction(ctx, transactionId).Execute()





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
	transactionId := "97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | transaction ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainTransaction(context.Background(), transactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainTransaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainTransaction`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainTransactionByMessageHash

> Transaction GetBlockchainTransactionByMessageHash(ctx, msgId).Execute()





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
	msgId := "97264395BD65A255A429B11326C84128B7D70FFED7949ABAE3036D506BA38621" // string | message ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainTransactionByMessageHash(context.Background(), msgId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainTransactionByMessageHash``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainTransactionByMessageHash`: Transaction
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainTransactionByMessageHash`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msgId** | **string** | message ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainTransactionByMessageHashRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transaction**](Transaction.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlockchainValidators

> Validators GetBlockchainValidators(ctx).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.GetBlockchainValidators(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetBlockchainValidators``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlockchainValidators`: Validators
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetBlockchainValidators`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlockchainValidatorsRequest struct via the builder pattern


### Return type

[**Validators**](Validators.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockchainConfig

> RawBlockchainConfig GetRawBlockchainConfig(ctx).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.GetRawBlockchainConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetRawBlockchainConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockchainConfig`: RawBlockchainConfig
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetRawBlockchainConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockchainConfigRequest struct via the builder pattern


### Return type

[**RawBlockchainConfig**](RawBlockchainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawBlockchainConfigFromBlock

> RawBlockchainConfig GetRawBlockchainConfigFromBlock(ctx, masterchainSeqno).Execute()





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
	masterchainSeqno := int32(123456) // int32 | masterchain block seqno

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetRawBlockchainConfigFromBlock(context.Background(), masterchainSeqno).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetRawBlockchainConfigFromBlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawBlockchainConfigFromBlock`: RawBlockchainConfig
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetRawBlockchainConfigFromBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**masterchainSeqno** | **int32** | masterchain block seqno | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawBlockchainConfigFromBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RawBlockchainConfig**](RawBlockchainConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReducedBlockchainBlocks

> ReducedBlocks GetReducedBlockchainBlocks(ctx).From(from).To(to).Execute()





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
	from := int64(789) // int64 | 
	to := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlockchainAPI.GetReducedBlockchainBlocks(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.GetReducedBlockchainBlocks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReducedBlockchainBlocks`: ReducedBlocks
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.GetReducedBlockchainBlocks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetReducedBlockchainBlocksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** |  | 
 **to** | **int64** |  | 

### Return type

[**ReducedBlocks**](ReducedBlocks.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendBlockchainMessage

> SendBlockchainMessage(ctx).SendBlockchainMessageRequest(sendBlockchainMessageRequest).Execute()





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
	sendBlockchainMessageRequest := *openapiclient.NewSendBlockchainMessageRequest() // SendBlockchainMessageRequest | both a single boc and a batch of boc serialized in base64 are accepted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlockchainAPI.SendBlockchainMessage(context.Background()).SendBlockchainMessageRequest(sendBlockchainMessageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.SendBlockchainMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendBlockchainMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sendBlockchainMessageRequest** | [**SendBlockchainMessageRequest**](SendBlockchainMessageRequest.md) | both a single boc and a batch of boc serialized in base64 are accepted | 

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


## Status

> ServiceStatus Status(ctx).Execute()





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
	resp, r, err := apiClient.BlockchainAPI.Status(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlockchainAPI.Status``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Status`: ServiceStatus
	fmt.Fprintf(os.Stdout, "Response from `BlockchainAPI.Status`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusRequest struct via the builder pattern


### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

