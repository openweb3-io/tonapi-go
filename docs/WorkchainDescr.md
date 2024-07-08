# WorkchainDescr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workchain** | **int32** |  | 
**EnabledSince** | **int64** |  | 
**ActualMinSplit** | **int32** |  | 
**MinSplit** | **int32** |  | 
**MaxSplit** | **int32** |  | 
**Basic** | **int32** |  | 
**Active** | **bool** |  | 
**AcceptMsgs** | **bool** |  | 
**Flags** | **int32** |  | 
**ZerostateRootHash** | **string** |  | 
**ZerostateFileHash** | **string** |  | 
**Version** | **int64** |  | 

## Methods

### NewWorkchainDescr

`func NewWorkchainDescr(workchain int32, enabledSince int64, actualMinSplit int32, minSplit int32, maxSplit int32, basic int32, active bool, acceptMsgs bool, flags int32, zerostateRootHash string, zerostateFileHash string, version int64, ) *WorkchainDescr`

NewWorkchainDescr instantiates a new WorkchainDescr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkchainDescrWithDefaults

`func NewWorkchainDescrWithDefaults() *WorkchainDescr`

NewWorkchainDescrWithDefaults instantiates a new WorkchainDescr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkchain

`func (o *WorkchainDescr) GetWorkchain() int32`

GetWorkchain returns the Workchain field if non-nil, zero value otherwise.

### GetWorkchainOk

`func (o *WorkchainDescr) GetWorkchainOk() (*int32, bool)`

GetWorkchainOk returns a tuple with the Workchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkchain

`func (o *WorkchainDescr) SetWorkchain(v int32)`

SetWorkchain sets Workchain field to given value.


### GetEnabledSince

`func (o *WorkchainDescr) GetEnabledSince() int64`

GetEnabledSince returns the EnabledSince field if non-nil, zero value otherwise.

### GetEnabledSinceOk

`func (o *WorkchainDescr) GetEnabledSinceOk() (*int64, bool)`

GetEnabledSinceOk returns a tuple with the EnabledSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledSince

`func (o *WorkchainDescr) SetEnabledSince(v int64)`

SetEnabledSince sets EnabledSince field to given value.


### GetActualMinSplit

`func (o *WorkchainDescr) GetActualMinSplit() int32`

GetActualMinSplit returns the ActualMinSplit field if non-nil, zero value otherwise.

### GetActualMinSplitOk

`func (o *WorkchainDescr) GetActualMinSplitOk() (*int32, bool)`

GetActualMinSplitOk returns a tuple with the ActualMinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMinSplit

`func (o *WorkchainDescr) SetActualMinSplit(v int32)`

SetActualMinSplit sets ActualMinSplit field to given value.


### GetMinSplit

`func (o *WorkchainDescr) GetMinSplit() int32`

GetMinSplit returns the MinSplit field if non-nil, zero value otherwise.

### GetMinSplitOk

`func (o *WorkchainDescr) GetMinSplitOk() (*int32, bool)`

GetMinSplitOk returns a tuple with the MinSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSplit

`func (o *WorkchainDescr) SetMinSplit(v int32)`

SetMinSplit sets MinSplit field to given value.


### GetMaxSplit

`func (o *WorkchainDescr) GetMaxSplit() int32`

GetMaxSplit returns the MaxSplit field if non-nil, zero value otherwise.

### GetMaxSplitOk

`func (o *WorkchainDescr) GetMaxSplitOk() (*int32, bool)`

GetMaxSplitOk returns a tuple with the MaxSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSplit

`func (o *WorkchainDescr) SetMaxSplit(v int32)`

SetMaxSplit sets MaxSplit field to given value.


### GetBasic

`func (o *WorkchainDescr) GetBasic() int32`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *WorkchainDescr) GetBasicOk() (*int32, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *WorkchainDescr) SetBasic(v int32)`

SetBasic sets Basic field to given value.


### GetActive

`func (o *WorkchainDescr) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *WorkchainDescr) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *WorkchainDescr) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAcceptMsgs

`func (o *WorkchainDescr) GetAcceptMsgs() bool`

GetAcceptMsgs returns the AcceptMsgs field if non-nil, zero value otherwise.

### GetAcceptMsgsOk

`func (o *WorkchainDescr) GetAcceptMsgsOk() (*bool, bool)`

GetAcceptMsgsOk returns a tuple with the AcceptMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptMsgs

`func (o *WorkchainDescr) SetAcceptMsgs(v bool)`

SetAcceptMsgs sets AcceptMsgs field to given value.


### GetFlags

`func (o *WorkchainDescr) GetFlags() int32`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *WorkchainDescr) GetFlagsOk() (*int32, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *WorkchainDescr) SetFlags(v int32)`

SetFlags sets Flags field to given value.


### GetZerostateRootHash

`func (o *WorkchainDescr) GetZerostateRootHash() string`

GetZerostateRootHash returns the ZerostateRootHash field if non-nil, zero value otherwise.

### GetZerostateRootHashOk

`func (o *WorkchainDescr) GetZerostateRootHashOk() (*string, bool)`

GetZerostateRootHashOk returns a tuple with the ZerostateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZerostateRootHash

`func (o *WorkchainDescr) SetZerostateRootHash(v string)`

SetZerostateRootHash sets ZerostateRootHash field to given value.


### GetZerostateFileHash

`func (o *WorkchainDescr) GetZerostateFileHash() string`

GetZerostateFileHash returns the ZerostateFileHash field if non-nil, zero value otherwise.

### GetZerostateFileHashOk

`func (o *WorkchainDescr) GetZerostateFileHashOk() (*string, bool)`

GetZerostateFileHashOk returns a tuple with the ZerostateFileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZerostateFileHash

`func (o *WorkchainDescr) SetZerostateFileHash(v string)`

SetZerostateFileHash sets ZerostateFileHash field to given value.


### GetVersion

`func (o *WorkchainDescr) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkchainDescr) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkchainDescr) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


