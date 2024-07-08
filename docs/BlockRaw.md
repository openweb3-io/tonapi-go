# BlockRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workchain** | **int32** |  | 
**Shard** | **string** |  | 
**Seqno** | **int32** |  | 
**RootHash** | **string** |  | 
**FileHash** | **string** |  | 

## Methods

### NewBlockRaw

`func NewBlockRaw(workchain int32, shard string, seqno int32, rootHash string, fileHash string, ) *BlockRaw`

NewBlockRaw instantiates a new BlockRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockRawWithDefaults

`func NewBlockRawWithDefaults() *BlockRaw`

NewBlockRawWithDefaults instantiates a new BlockRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkchain

`func (o *BlockRaw) GetWorkchain() int32`

GetWorkchain returns the Workchain field if non-nil, zero value otherwise.

### GetWorkchainOk

`func (o *BlockRaw) GetWorkchainOk() (*int32, bool)`

GetWorkchainOk returns a tuple with the Workchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkchain

`func (o *BlockRaw) SetWorkchain(v int32)`

SetWorkchain sets Workchain field to given value.


### GetShard

`func (o *BlockRaw) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *BlockRaw) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *BlockRaw) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeqno

`func (o *BlockRaw) GetSeqno() int32`

GetSeqno returns the Seqno field if non-nil, zero value otherwise.

### GetSeqnoOk

`func (o *BlockRaw) GetSeqnoOk() (*int32, bool)`

GetSeqnoOk returns a tuple with the Seqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqno

`func (o *BlockRaw) SetSeqno(v int32)`

SetSeqno sets Seqno field to given value.


### GetRootHash

`func (o *BlockRaw) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *BlockRaw) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *BlockRaw) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.


### GetFileHash

`func (o *BlockRaw) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *BlockRaw) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *BlockRaw) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


