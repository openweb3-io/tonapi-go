# BlockchainBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxQuantity** | **int32** |  | 
**ValueFlow** | [**BlockValueFlow**](BlockValueFlow.md) |  | 
**WorkchainId** | **int32** |  | 
**Shard** | **string** |  | 
**Seqno** | **int32** |  | 
**RootHash** | **string** |  | 
**FileHash** | **string** |  | 
**GlobalId** | **int32** |  | 
**Version** | **int32** |  | 
**AfterMerge** | **bool** |  | 
**BeforeSplit** | **bool** |  | 
**AfterSplit** | **bool** |  | 
**WantSplit** | **bool** |  | 
**WantMerge** | **bool** |  | 
**KeyBlock** | **bool** |  | 
**GenUtime** | **int64** |  | 
**StartLt** | **int64** |  | 
**EndLt** | **int64** |  | 
**VertSeqno** | **int32** |  | 
**GenCatchainSeqno** | **int32** |  | 
**MinRefMcSeqno** | **int32** |  | 
**PrevKeyBlockSeqno** | **int32** |  | 
**GenSoftwareVersion** | Pointer to **int32** |  | [optional] 
**GenSoftwareCapabilities** | Pointer to **int64** |  | [optional] 
**MasterRef** | Pointer to **string** |  | [optional] 
**PrevRefs** | **[]string** |  | 
**InMsgDescrLength** | **int64** |  | 
**OutMsgDescrLength** | **int64** |  | 
**RandSeed** | **string** |  | 
**CreatedBy** | **string** |  | 

## Methods

### NewBlockchainBlock

`func NewBlockchainBlock(txQuantity int32, valueFlow BlockValueFlow, workchainId int32, shard string, seqno int32, rootHash string, fileHash string, globalId int32, version int32, afterMerge bool, beforeSplit bool, afterSplit bool, wantSplit bool, wantMerge bool, keyBlock bool, genUtime int64, startLt int64, endLt int64, vertSeqno int32, genCatchainSeqno int32, minRefMcSeqno int32, prevKeyBlockSeqno int32, prevRefs []string, inMsgDescrLength int64, outMsgDescrLength int64, randSeed string, createdBy string, ) *BlockchainBlock`

NewBlockchainBlock instantiates a new BlockchainBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainBlockWithDefaults

`func NewBlockchainBlockWithDefaults() *BlockchainBlock`

NewBlockchainBlockWithDefaults instantiates a new BlockchainBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxQuantity

`func (o *BlockchainBlock) GetTxQuantity() int32`

GetTxQuantity returns the TxQuantity field if non-nil, zero value otherwise.

### GetTxQuantityOk

`func (o *BlockchainBlock) GetTxQuantityOk() (*int32, bool)`

GetTxQuantityOk returns a tuple with the TxQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxQuantity

`func (o *BlockchainBlock) SetTxQuantity(v int32)`

SetTxQuantity sets TxQuantity field to given value.


### GetValueFlow

`func (o *BlockchainBlock) GetValueFlow() BlockValueFlow`

GetValueFlow returns the ValueFlow field if non-nil, zero value otherwise.

### GetValueFlowOk

`func (o *BlockchainBlock) GetValueFlowOk() (*BlockValueFlow, bool)`

GetValueFlowOk returns a tuple with the ValueFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFlow

`func (o *BlockchainBlock) SetValueFlow(v BlockValueFlow)`

SetValueFlow sets ValueFlow field to given value.


### GetWorkchainId

`func (o *BlockchainBlock) GetWorkchainId() int32`

GetWorkchainId returns the WorkchainId field if non-nil, zero value otherwise.

### GetWorkchainIdOk

`func (o *BlockchainBlock) GetWorkchainIdOk() (*int32, bool)`

GetWorkchainIdOk returns a tuple with the WorkchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkchainId

`func (o *BlockchainBlock) SetWorkchainId(v int32)`

SetWorkchainId sets WorkchainId field to given value.


### GetShard

`func (o *BlockchainBlock) GetShard() string`

GetShard returns the Shard field if non-nil, zero value otherwise.

### GetShardOk

`func (o *BlockchainBlock) GetShardOk() (*string, bool)`

GetShardOk returns a tuple with the Shard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShard

`func (o *BlockchainBlock) SetShard(v string)`

SetShard sets Shard field to given value.


### GetSeqno

`func (o *BlockchainBlock) GetSeqno() int32`

GetSeqno returns the Seqno field if non-nil, zero value otherwise.

### GetSeqnoOk

`func (o *BlockchainBlock) GetSeqnoOk() (*int32, bool)`

GetSeqnoOk returns a tuple with the Seqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqno

`func (o *BlockchainBlock) SetSeqno(v int32)`

SetSeqno sets Seqno field to given value.


### GetRootHash

`func (o *BlockchainBlock) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *BlockchainBlock) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *BlockchainBlock) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.


### GetFileHash

`func (o *BlockchainBlock) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *BlockchainBlock) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *BlockchainBlock) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.


### GetGlobalId

`func (o *BlockchainBlock) GetGlobalId() int32`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *BlockchainBlock) GetGlobalIdOk() (*int32, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *BlockchainBlock) SetGlobalId(v int32)`

SetGlobalId sets GlobalId field to given value.


### GetVersion

`func (o *BlockchainBlock) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BlockchainBlock) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BlockchainBlock) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetAfterMerge

`func (o *BlockchainBlock) GetAfterMerge() bool`

GetAfterMerge returns the AfterMerge field if non-nil, zero value otherwise.

### GetAfterMergeOk

`func (o *BlockchainBlock) GetAfterMergeOk() (*bool, bool)`

GetAfterMergeOk returns a tuple with the AfterMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterMerge

`func (o *BlockchainBlock) SetAfterMerge(v bool)`

SetAfterMerge sets AfterMerge field to given value.


### GetBeforeSplit

`func (o *BlockchainBlock) GetBeforeSplit() bool`

GetBeforeSplit returns the BeforeSplit field if non-nil, zero value otherwise.

### GetBeforeSplitOk

`func (o *BlockchainBlock) GetBeforeSplitOk() (*bool, bool)`

GetBeforeSplitOk returns a tuple with the BeforeSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeSplit

`func (o *BlockchainBlock) SetBeforeSplit(v bool)`

SetBeforeSplit sets BeforeSplit field to given value.


### GetAfterSplit

`func (o *BlockchainBlock) GetAfterSplit() bool`

GetAfterSplit returns the AfterSplit field if non-nil, zero value otherwise.

### GetAfterSplitOk

`func (o *BlockchainBlock) GetAfterSplitOk() (*bool, bool)`

GetAfterSplitOk returns a tuple with the AfterSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterSplit

`func (o *BlockchainBlock) SetAfterSplit(v bool)`

SetAfterSplit sets AfterSplit field to given value.


### GetWantSplit

`func (o *BlockchainBlock) GetWantSplit() bool`

GetWantSplit returns the WantSplit field if non-nil, zero value otherwise.

### GetWantSplitOk

`func (o *BlockchainBlock) GetWantSplitOk() (*bool, bool)`

GetWantSplitOk returns a tuple with the WantSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantSplit

`func (o *BlockchainBlock) SetWantSplit(v bool)`

SetWantSplit sets WantSplit field to given value.


### GetWantMerge

`func (o *BlockchainBlock) GetWantMerge() bool`

GetWantMerge returns the WantMerge field if non-nil, zero value otherwise.

### GetWantMergeOk

`func (o *BlockchainBlock) GetWantMergeOk() (*bool, bool)`

GetWantMergeOk returns a tuple with the WantMerge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWantMerge

`func (o *BlockchainBlock) SetWantMerge(v bool)`

SetWantMerge sets WantMerge field to given value.


### GetKeyBlock

`func (o *BlockchainBlock) GetKeyBlock() bool`

GetKeyBlock returns the KeyBlock field if non-nil, zero value otherwise.

### GetKeyBlockOk

`func (o *BlockchainBlock) GetKeyBlockOk() (*bool, bool)`

GetKeyBlockOk returns a tuple with the KeyBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBlock

`func (o *BlockchainBlock) SetKeyBlock(v bool)`

SetKeyBlock sets KeyBlock field to given value.


### GetGenUtime

`func (o *BlockchainBlock) GetGenUtime() int64`

GetGenUtime returns the GenUtime field if non-nil, zero value otherwise.

### GetGenUtimeOk

`func (o *BlockchainBlock) GetGenUtimeOk() (*int64, bool)`

GetGenUtimeOk returns a tuple with the GenUtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenUtime

`func (o *BlockchainBlock) SetGenUtime(v int64)`

SetGenUtime sets GenUtime field to given value.


### GetStartLt

`func (o *BlockchainBlock) GetStartLt() int64`

GetStartLt returns the StartLt field if non-nil, zero value otherwise.

### GetStartLtOk

`func (o *BlockchainBlock) GetStartLtOk() (*int64, bool)`

GetStartLtOk returns a tuple with the StartLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLt

`func (o *BlockchainBlock) SetStartLt(v int64)`

SetStartLt sets StartLt field to given value.


### GetEndLt

`func (o *BlockchainBlock) GetEndLt() int64`

GetEndLt returns the EndLt field if non-nil, zero value otherwise.

### GetEndLtOk

`func (o *BlockchainBlock) GetEndLtOk() (*int64, bool)`

GetEndLtOk returns a tuple with the EndLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLt

`func (o *BlockchainBlock) SetEndLt(v int64)`

SetEndLt sets EndLt field to given value.


### GetVertSeqno

`func (o *BlockchainBlock) GetVertSeqno() int32`

GetVertSeqno returns the VertSeqno field if non-nil, zero value otherwise.

### GetVertSeqnoOk

`func (o *BlockchainBlock) GetVertSeqnoOk() (*int32, bool)`

GetVertSeqnoOk returns a tuple with the VertSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertSeqno

`func (o *BlockchainBlock) SetVertSeqno(v int32)`

SetVertSeqno sets VertSeqno field to given value.


### GetGenCatchainSeqno

`func (o *BlockchainBlock) GetGenCatchainSeqno() int32`

GetGenCatchainSeqno returns the GenCatchainSeqno field if non-nil, zero value otherwise.

### GetGenCatchainSeqnoOk

`func (o *BlockchainBlock) GetGenCatchainSeqnoOk() (*int32, bool)`

GetGenCatchainSeqnoOk returns a tuple with the GenCatchainSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenCatchainSeqno

`func (o *BlockchainBlock) SetGenCatchainSeqno(v int32)`

SetGenCatchainSeqno sets GenCatchainSeqno field to given value.


### GetMinRefMcSeqno

`func (o *BlockchainBlock) GetMinRefMcSeqno() int32`

GetMinRefMcSeqno returns the MinRefMcSeqno field if non-nil, zero value otherwise.

### GetMinRefMcSeqnoOk

`func (o *BlockchainBlock) GetMinRefMcSeqnoOk() (*int32, bool)`

GetMinRefMcSeqnoOk returns a tuple with the MinRefMcSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRefMcSeqno

`func (o *BlockchainBlock) SetMinRefMcSeqno(v int32)`

SetMinRefMcSeqno sets MinRefMcSeqno field to given value.


### GetPrevKeyBlockSeqno

`func (o *BlockchainBlock) GetPrevKeyBlockSeqno() int32`

GetPrevKeyBlockSeqno returns the PrevKeyBlockSeqno field if non-nil, zero value otherwise.

### GetPrevKeyBlockSeqnoOk

`func (o *BlockchainBlock) GetPrevKeyBlockSeqnoOk() (*int32, bool)`

GetPrevKeyBlockSeqnoOk returns a tuple with the PrevKeyBlockSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevKeyBlockSeqno

`func (o *BlockchainBlock) SetPrevKeyBlockSeqno(v int32)`

SetPrevKeyBlockSeqno sets PrevKeyBlockSeqno field to given value.


### GetGenSoftwareVersion

`func (o *BlockchainBlock) GetGenSoftwareVersion() int32`

GetGenSoftwareVersion returns the GenSoftwareVersion field if non-nil, zero value otherwise.

### GetGenSoftwareVersionOk

`func (o *BlockchainBlock) GetGenSoftwareVersionOk() (*int32, bool)`

GetGenSoftwareVersionOk returns a tuple with the GenSoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenSoftwareVersion

`func (o *BlockchainBlock) SetGenSoftwareVersion(v int32)`

SetGenSoftwareVersion sets GenSoftwareVersion field to given value.

### HasGenSoftwareVersion

`func (o *BlockchainBlock) HasGenSoftwareVersion() bool`

HasGenSoftwareVersion returns a boolean if a field has been set.

### GetGenSoftwareCapabilities

`func (o *BlockchainBlock) GetGenSoftwareCapabilities() int64`

GetGenSoftwareCapabilities returns the GenSoftwareCapabilities field if non-nil, zero value otherwise.

### GetGenSoftwareCapabilitiesOk

`func (o *BlockchainBlock) GetGenSoftwareCapabilitiesOk() (*int64, bool)`

GetGenSoftwareCapabilitiesOk returns a tuple with the GenSoftwareCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenSoftwareCapabilities

`func (o *BlockchainBlock) SetGenSoftwareCapabilities(v int64)`

SetGenSoftwareCapabilities sets GenSoftwareCapabilities field to given value.

### HasGenSoftwareCapabilities

`func (o *BlockchainBlock) HasGenSoftwareCapabilities() bool`

HasGenSoftwareCapabilities returns a boolean if a field has been set.

### GetMasterRef

`func (o *BlockchainBlock) GetMasterRef() string`

GetMasterRef returns the MasterRef field if non-nil, zero value otherwise.

### GetMasterRefOk

`func (o *BlockchainBlock) GetMasterRefOk() (*string, bool)`

GetMasterRefOk returns a tuple with the MasterRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterRef

`func (o *BlockchainBlock) SetMasterRef(v string)`

SetMasterRef sets MasterRef field to given value.

### HasMasterRef

`func (o *BlockchainBlock) HasMasterRef() bool`

HasMasterRef returns a boolean if a field has been set.

### GetPrevRefs

`func (o *BlockchainBlock) GetPrevRefs() []string`

GetPrevRefs returns the PrevRefs field if non-nil, zero value otherwise.

### GetPrevRefsOk

`func (o *BlockchainBlock) GetPrevRefsOk() (*[]string, bool)`

GetPrevRefsOk returns a tuple with the PrevRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevRefs

`func (o *BlockchainBlock) SetPrevRefs(v []string)`

SetPrevRefs sets PrevRefs field to given value.


### GetInMsgDescrLength

`func (o *BlockchainBlock) GetInMsgDescrLength() int64`

GetInMsgDescrLength returns the InMsgDescrLength field if non-nil, zero value otherwise.

### GetInMsgDescrLengthOk

`func (o *BlockchainBlock) GetInMsgDescrLengthOk() (*int64, bool)`

GetInMsgDescrLengthOk returns a tuple with the InMsgDescrLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsgDescrLength

`func (o *BlockchainBlock) SetInMsgDescrLength(v int64)`

SetInMsgDescrLength sets InMsgDescrLength field to given value.


### GetOutMsgDescrLength

`func (o *BlockchainBlock) GetOutMsgDescrLength() int64`

GetOutMsgDescrLength returns the OutMsgDescrLength field if non-nil, zero value otherwise.

### GetOutMsgDescrLengthOk

`func (o *BlockchainBlock) GetOutMsgDescrLengthOk() (*int64, bool)`

GetOutMsgDescrLengthOk returns a tuple with the OutMsgDescrLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgDescrLength

`func (o *BlockchainBlock) SetOutMsgDescrLength(v int64)`

SetOutMsgDescrLength sets OutMsgDescrLength field to given value.


### GetRandSeed

`func (o *BlockchainBlock) GetRandSeed() string`

GetRandSeed returns the RandSeed field if non-nil, zero value otherwise.

### GetRandSeedOk

`func (o *BlockchainBlock) GetRandSeedOk() (*string, bool)`

GetRandSeedOk returns a tuple with the RandSeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandSeed

`func (o *BlockchainBlock) SetRandSeed(v string)`

SetRandSeed sets RandSeed field to given value.


### GetCreatedBy

`func (o *BlockchainBlock) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *BlockchainBlock) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *BlockchainBlock) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


