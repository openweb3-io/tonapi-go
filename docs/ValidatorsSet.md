# ValidatorsSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UtimeSince** | **int32** |  | 
**UtimeUntil** | **int32** |  | 
**Total** | **int32** |  | 
**Main** | **int32** |  | 
**TotalWeight** | Pointer to **string** |  | [optional] 
**List** | [**[]ValidatorsSetListInner**](ValidatorsSetListInner.md) |  | 

## Methods

### NewValidatorsSet

`func NewValidatorsSet(utimeSince int32, utimeUntil int32, total int32, main int32, list []ValidatorsSetListInner, ) *ValidatorsSet`

NewValidatorsSet instantiates a new ValidatorsSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorsSetWithDefaults

`func NewValidatorsSetWithDefaults() *ValidatorsSet`

NewValidatorsSetWithDefaults instantiates a new ValidatorsSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUtimeSince

`func (o *ValidatorsSet) GetUtimeSince() int32`

GetUtimeSince returns the UtimeSince field if non-nil, zero value otherwise.

### GetUtimeSinceOk

`func (o *ValidatorsSet) GetUtimeSinceOk() (*int32, bool)`

GetUtimeSinceOk returns a tuple with the UtimeSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtimeSince

`func (o *ValidatorsSet) SetUtimeSince(v int32)`

SetUtimeSince sets UtimeSince field to given value.


### GetUtimeUntil

`func (o *ValidatorsSet) GetUtimeUntil() int32`

GetUtimeUntil returns the UtimeUntil field if non-nil, zero value otherwise.

### GetUtimeUntilOk

`func (o *ValidatorsSet) GetUtimeUntilOk() (*int32, bool)`

GetUtimeUntilOk returns a tuple with the UtimeUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtimeUntil

`func (o *ValidatorsSet) SetUtimeUntil(v int32)`

SetUtimeUntil sets UtimeUntil field to given value.


### GetTotal

`func (o *ValidatorsSet) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ValidatorsSet) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ValidatorsSet) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetMain

`func (o *ValidatorsSet) GetMain() int32`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *ValidatorsSet) GetMainOk() (*int32, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *ValidatorsSet) SetMain(v int32)`

SetMain sets Main field to given value.


### GetTotalWeight

`func (o *ValidatorsSet) GetTotalWeight() string`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *ValidatorsSet) GetTotalWeightOk() (*string, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *ValidatorsSet) SetTotalWeight(v string)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *ValidatorsSet) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetList

`func (o *ValidatorsSet) GetList() []ValidatorsSetListInner`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *ValidatorsSet) GetListOk() (*[]ValidatorsSetListInner, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *ValidatorsSet) SetList(v []ValidatorsSetListInner)`

SetList sets List field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


