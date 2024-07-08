# ComputePhase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Skipped** | **bool** |  | 
**SkipReason** | Pointer to [**ComputeSkipReason**](ComputeSkipReason.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**GasFees** | Pointer to **int64** |  | [optional] 
**GasUsed** | Pointer to **int64** |  | [optional] 
**VmSteps** | Pointer to **int32** |  | [optional] 
**ExitCode** | Pointer to **int32** |  | [optional] 
**ExitCodeDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewComputePhase

`func NewComputePhase(skipped bool, ) *ComputePhase`

NewComputePhase instantiates a new ComputePhase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComputePhaseWithDefaults

`func NewComputePhaseWithDefaults() *ComputePhase`

NewComputePhaseWithDefaults instantiates a new ComputePhase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSkipped

`func (o *ComputePhase) GetSkipped() bool`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *ComputePhase) GetSkippedOk() (*bool, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *ComputePhase) SetSkipped(v bool)`

SetSkipped sets Skipped field to given value.


### GetSkipReason

`func (o *ComputePhase) GetSkipReason() ComputeSkipReason`

GetSkipReason returns the SkipReason field if non-nil, zero value otherwise.

### GetSkipReasonOk

`func (o *ComputePhase) GetSkipReasonOk() (*ComputeSkipReason, bool)`

GetSkipReasonOk returns a tuple with the SkipReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipReason

`func (o *ComputePhase) SetSkipReason(v ComputeSkipReason)`

SetSkipReason sets SkipReason field to given value.

### HasSkipReason

`func (o *ComputePhase) HasSkipReason() bool`

HasSkipReason returns a boolean if a field has been set.

### GetSuccess

`func (o *ComputePhase) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ComputePhase) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ComputePhase) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ComputePhase) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetGasFees

`func (o *ComputePhase) GetGasFees() int64`

GetGasFees returns the GasFees field if non-nil, zero value otherwise.

### GetGasFeesOk

`func (o *ComputePhase) GetGasFeesOk() (*int64, bool)`

GetGasFeesOk returns a tuple with the GasFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasFees

`func (o *ComputePhase) SetGasFees(v int64)`

SetGasFees sets GasFees field to given value.

### HasGasFees

`func (o *ComputePhase) HasGasFees() bool`

HasGasFees returns a boolean if a field has been set.

### GetGasUsed

`func (o *ComputePhase) GetGasUsed() int64`

GetGasUsed returns the GasUsed field if non-nil, zero value otherwise.

### GetGasUsedOk

`func (o *ComputePhase) GetGasUsedOk() (*int64, bool)`

GetGasUsedOk returns a tuple with the GasUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasUsed

`func (o *ComputePhase) SetGasUsed(v int64)`

SetGasUsed sets GasUsed field to given value.

### HasGasUsed

`func (o *ComputePhase) HasGasUsed() bool`

HasGasUsed returns a boolean if a field has been set.

### GetVmSteps

`func (o *ComputePhase) GetVmSteps() int32`

GetVmSteps returns the VmSteps field if non-nil, zero value otherwise.

### GetVmStepsOk

`func (o *ComputePhase) GetVmStepsOk() (*int32, bool)`

GetVmStepsOk returns a tuple with the VmSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmSteps

`func (o *ComputePhase) SetVmSteps(v int32)`

SetVmSteps sets VmSteps field to given value.

### HasVmSteps

`func (o *ComputePhase) HasVmSteps() bool`

HasVmSteps returns a boolean if a field has been set.

### GetExitCode

`func (o *ComputePhase) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *ComputePhase) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *ComputePhase) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.

### HasExitCode

`func (o *ComputePhase) HasExitCode() bool`

HasExitCode returns a boolean if a field has been set.

### GetExitCodeDescription

`func (o *ComputePhase) GetExitCodeDescription() string`

GetExitCodeDescription returns the ExitCodeDescription field if non-nil, zero value otherwise.

### GetExitCodeDescriptionOk

`func (o *ComputePhase) GetExitCodeDescriptionOk() (*string, bool)`

GetExitCodeDescriptionOk returns a tuple with the ExitCodeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCodeDescription

`func (o *ComputePhase) SetExitCodeDescription(v string)`

SetExitCodeDescription sets ExitCodeDescription field to given value.

### HasExitCodeDescription

`func (o *ComputePhase) HasExitCodeDescription() bool`

HasExitCodeDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


