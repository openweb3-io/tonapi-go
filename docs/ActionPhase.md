# ActionPhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**ResultCode** | **int32** |  | 
**TotalActions** | **int32** |  | 
**SkippedActions** | **int32** |  | 
**FwdFees** | **int64** |  | 
**TotalFees** | **int64** |  | 
**ResultCodeDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewActionPhase

`func NewActionPhase(success bool, resultCode int32, totalActions int32, skippedActions int32, fwdFees int64, totalFees int64, ) *ActionPhase`

NewActionPhase instantiates a new ActionPhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPhaseWithDefaults

`func NewActionPhaseWithDefaults() *ActionPhase`

NewActionPhaseWithDefaults instantiates a new ActionPhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ActionPhase) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ActionPhase) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ActionPhase) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetResultCode

`func (o *ActionPhase) GetResultCode() int32`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *ActionPhase) GetResultCodeOk() (*int32, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *ActionPhase) SetResultCode(v int32)`

SetResultCode sets ResultCode field to given value.


### GetTotalActions

`func (o *ActionPhase) GetTotalActions() int32`

GetTotalActions returns the TotalActions field if non-nil, zero value otherwise.

### GetTotalActionsOk

`func (o *ActionPhase) GetTotalActionsOk() (*int32, bool)`

GetTotalActionsOk returns a tuple with the TotalActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalActions

`func (o *ActionPhase) SetTotalActions(v int32)`

SetTotalActions sets TotalActions field to given value.


### GetSkippedActions

`func (o *ActionPhase) GetSkippedActions() int32`

GetSkippedActions returns the SkippedActions field if non-nil, zero value otherwise.

### GetSkippedActionsOk

`func (o *ActionPhase) GetSkippedActionsOk() (*int32, bool)`

GetSkippedActionsOk returns a tuple with the SkippedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedActions

`func (o *ActionPhase) SetSkippedActions(v int32)`

SetSkippedActions sets SkippedActions field to given value.


### GetFwdFees

`func (o *ActionPhase) GetFwdFees() int64`

GetFwdFees returns the FwdFees field if non-nil, zero value otherwise.

### GetFwdFeesOk

`func (o *ActionPhase) GetFwdFeesOk() (*int64, bool)`

GetFwdFeesOk returns a tuple with the FwdFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdFees

`func (o *ActionPhase) SetFwdFees(v int64)`

SetFwdFees sets FwdFees field to given value.


### GetTotalFees

`func (o *ActionPhase) GetTotalFees() int64`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *ActionPhase) GetTotalFeesOk() (*int64, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *ActionPhase) SetTotalFees(v int64)`

SetTotalFees sets TotalFees field to given value.


### GetResultCodeDescription

`func (o *ActionPhase) GetResultCodeDescription() string`

GetResultCodeDescription returns the ResultCodeDescription field if non-nil, zero value otherwise.

### GetResultCodeDescriptionOk

`func (o *ActionPhase) GetResultCodeDescriptionOk() (*string, bool)`

GetResultCodeDescriptionOk returns a tuple with the ResultCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCodeDescription

`func (o *ActionPhase) SetResultCodeDescription(v string)`

SetResultCodeDescription sets ResultCodeDescription field to given value.

### HasResultCodeDescription

`func (o *ActionPhase) HasResultCodeDescription() bool`

HasResultCodeDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


