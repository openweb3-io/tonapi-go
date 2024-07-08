# ReducedBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkchainId** | **int32** |  | 
**Shard** | **string** |  | 
**Seqno** | **int32** |  | 
**MasterRef** | Pointer to **string** |  | [optional] 
**TxQuantity** | **int32** |  | 
**Utime** | **int64** |  | 
**ShardsBlocks** | **[]string** |  | 

## Methods

### NewReducedBlock

`func NewReducedBlock(workchainId int32, shard string, seqno int32, txQuantity int32, utime int64, shardsBlocks []string, ) *ReducedBlock`

NewReducedBlock instantiates a new ReducedBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReducedBlockWithDefaults

`func NewReducedBlockWithDefaults() *ReducedBlock`

NewReducedBlockWithDefaults instantiates a new ReducedBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkchainId

`func (o *ReducedBlock) GetWorkchainId() int32`

GetWorkchainId returns the WorkchainId field if non-nil, zero value otherwise.

### GetWorkchainIdOk

`func (o *ReducedBlock) GetWorkchainIdOk() (*int32, bool)`

GetWorkchainIdOk returns a tuple with the WorkchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkchainId

`func (o *ReducedBlock) SetWorkchainId(v int32)`

SetWorkchainId sets WorkchainId field to given value.


### GetShard

`func (o *ReducedBlock) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *ReducedBlock) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *ReducedBlock) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeqno

`func (o *ReducedBlock) GetSeqno() int32`

GetSeqno returns the Seqno field if non-nil, zero value otherwise.

### GetSeqnoOk

`func (o *ReducedBlock) GetSeqnoOk() (*int32, bool)`

GetSeqnoOk returns a tuple with the Seqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqno

`func (o *ReducedBlock) SetSeqno(v int32)`

SetSeqno sets Seqno field to given value.


### GetMasterRef

`func (o *ReducedBlock) GetMasterRef() string`

GetMasterRef returns the MasterRef field if non-nil, zero value otherwise.

### GetMasterRefOk

`func (o *ReducedBlock) GetMasterRefOk() (*string, bool)`

GetMasterRefOk returns a tuple with the MasterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRef

`func (o *ReducedBlock) SetMasterRef(v string)`

SetMasterRef sets MasterRef field to given value.

### HasMasterRef

`func (o *ReducedBlock) HasMasterRef() bool`

HasMasterRef returns a boolean if a field has been set.

### GetTxQuantity

`func (o *ReducedBlock) GetTxQuantity() int32`

GetTxQuantity returns the TxQuantity field if non-nil, zero value otherwise.

### GetTxQuantityOk

`func (o *ReducedBlock) GetTxQuantityOk() (*int32, bool)`

GetTxQuantityOk returns a tuple with the TxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQuantity

`func (o *ReducedBlock) SetTxQuantity(v int32)`

SetTxQuantity sets TxQuantity field to given value.


### GetUtime

`func (o *ReducedBlock) GetUtime() int64`

GetUtime returns the Utime field if non-nil, zero value otherwise.

### GetUtimeOk

`func (o *ReducedBlock) GetUtimeOk() (*int64, bool)`

GetUtimeOk returns a tuple with the Utime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtime

`func (o *ReducedBlock) SetUtime(v int64)`

SetUtime sets Utime field to given value.


### GetShardsBlocks

`func (o *ReducedBlock) GetShardsBlocks() []string`

GetShardsBlocks returns the ShardsBlocks field if non-nil, zero value otherwise.

### GetShardsBlocksOk

`func (o *ReducedBlock) GetShardsBlocksOk() (*[]string, bool)`

GetShardsBlocksOk returns a tuple with the ShardsBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardsBlocks

`func (o *ReducedBlock) SetShardsBlocks(v []string)`

SetShardsBlocks sets ShardsBlocks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


