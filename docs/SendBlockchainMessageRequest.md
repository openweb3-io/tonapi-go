# SendBlockchainMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boc** | Pointer to **string** |  | [optional] 
**Batch** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSendBlockchainMessageRequest

`func NewSendBlockchainMessageRequest() *SendBlockchainMessageRequest`

NewSendBlockchainMessageRequest instantiates a new SendBlockchainMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendBlockchainMessageRequestWithDefaults

`func NewSendBlockchainMessageRequestWithDefaults() *SendBlockchainMessageRequest`

NewSendBlockchainMessageRequestWithDefaults instantiates a new SendBlockchainMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoc

`func (o *SendBlockchainMessageRequest) GetBoc() string`

GetBoc returns the Boc field if non-nil, zero value otherwise.

### GetBocOk

`func (o *SendBlockchainMessageRequest) GetBocOk() (*string, bool)`

GetBocOk returns a tuple with the Boc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoc

`func (o *SendBlockchainMessageRequest) SetBoc(v string)`

SetBoc sets Boc field to given value.

### HasBoc

`func (o *SendBlockchainMessageRequest) HasBoc() bool`

HasBoc returns a boolean if a field has been set.

### GetBatch

`func (o *SendBlockchainMessageRequest) GetBatch() []string`

GetBatch returns the Batch field if non-nil, zero value otherwise.

### GetBatchOk

`func (o *SendBlockchainMessageRequest) GetBatchOk() (*[]string, bool)`

GetBatchOk returns a tuple with the Batch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatch

`func (o *SendBlockchainMessageRequest) SetBatch(v []string)`

SetBatch sets Batch field to given value.

### HasBatch

`func (o *SendBlockchainMessageRequest) HasBatch() bool`

HasBatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


