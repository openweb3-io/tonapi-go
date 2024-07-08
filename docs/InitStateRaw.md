# InitStateRaw

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Workchain** | **int32** |  | 
**RootHash** | **string** |  | 
**FileHash** | **string** |  | 

## Methods

### NewInitStateRaw

`func NewInitStateRaw(workchain int32, rootHash string, fileHash string, ) *InitStateRaw`

NewInitStateRaw instantiates a new InitStateRaw object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitStateRawWithDefaults

`func NewInitStateRawWithDefaults() *InitStateRaw`

NewInitStateRawWithDefaults instantiates a new InitStateRaw object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkchain

`func (o *InitStateRaw) GetWorkchain() int32`

GetWorkchain returns the Workchain field if non-nil, zero value otherwise.

### GetWorkchainOk

`func (o *InitStateRaw) GetWorkchainOk() (*int32, bool)`

GetWorkchainOk returns a tuple with the Workchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkchain

`func (o *InitStateRaw) SetWorkchain(v int32)`

SetWorkchain sets Workchain field to given value.


### GetRootHash

`func (o *InitStateRaw) GetRootHash() string`

GetRootHash returns the RootHash field if non-nil, zero value otherwise.

### GetRootHashOk

`func (o *InitStateRaw) GetRootHashOk() (*string, bool)`

GetRootHashOk returns a tuple with the RootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootHash

`func (o *InitStateRaw) SetRootHash(v string)`

SetRootHash sets RootHash field to given value.


### GetFileHash

`func (o *InitStateRaw) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *InitStateRaw) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *InitStateRaw) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


