# StoragePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeesCollected** | **int64** |  | 
**FeesDue** | Pointer to **int64** |  | [optional] 
**StatusChange** | [**AccStatusChange**](AccStatusChange.md) |  | 

## Methods

### NewStoragePhase

`func NewStoragePhase(feesCollected int64, statusChange AccStatusChange, ) *StoragePhase`

NewStoragePhase instantiates a new StoragePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoragePhaseWithDefaults

`func NewStoragePhaseWithDefaults() *StoragePhase`

NewStoragePhaseWithDefaults instantiates a new StoragePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeesCollected

`func (o *StoragePhase) GetFeesCollected() int64`

GetFeesCollected returns the FeesCollected field if non-nil, zero value otherwise.

### GetFeesCollectedOk

`func (o *StoragePhase) GetFeesCollectedOk() (*int64, bool)`

GetFeesCollectedOk returns a tuple with the FeesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesCollected

`func (o *StoragePhase) SetFeesCollected(v int64)`

SetFeesCollected sets FeesCollected field to given value.


### GetFeesDue

`func (o *StoragePhase) GetFeesDue() int64`

GetFeesDue returns the FeesDue field if non-nil, zero value otherwise.

### GetFeesDueOk

`func (o *StoragePhase) GetFeesDueOk() (*int64, bool)`

GetFeesDueOk returns a tuple with the FeesDue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesDue

`func (o *StoragePhase) SetFeesDue(v int64)`

SetFeesDue sets FeesDue field to given value.

### HasFeesDue

`func (o *StoragePhase) HasFeesDue() bool`

HasFeesDue returns a boolean if a field has been set.

### GetStatusChange

`func (o *StoragePhase) GetStatusChange() AccStatusChange`

GetStatusChange returns the StatusChange field if non-nil, zero value otherwise.

### GetStatusChangeOk

`func (o *StoragePhase) GetStatusChangeOk() (*AccStatusChange, bool)`

GetStatusChangeOk returns a tuple with the StatusChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChange

`func (o *StoragePhase) SetStatusChange(v AccStatusChange)`

SetStatusChange sets StatusChange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


