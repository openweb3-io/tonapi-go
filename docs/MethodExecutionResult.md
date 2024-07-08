# MethodExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**ExitCode** | **int32** | tvm exit code | 
**Stack** | [**[]TvmStackRecord**](TvmStackRecord.md) |  | 
**Decoded** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMethodExecutionResult

`func NewMethodExecutionResult(success bool, exitCode int32, stack []TvmStackRecord, ) *MethodExecutionResult`

NewMethodExecutionResult instantiates a new MethodExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMethodExecutionResultWithDefaults

`func NewMethodExecutionResultWithDefaults() *MethodExecutionResult`

NewMethodExecutionResultWithDefaults instantiates a new MethodExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MethodExecutionResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MethodExecutionResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MethodExecutionResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetExitCode

`func (o *MethodExecutionResult) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *MethodExecutionResult) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *MethodExecutionResult) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetStack

`func (o *MethodExecutionResult) GetStack() []TvmStackRecord`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *MethodExecutionResult) GetStackOk() (*[]TvmStackRecord, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *MethodExecutionResult) SetStack(v []TvmStackRecord)`

SetStack sets Stack field to given value.


### GetDecoded

`func (o *MethodExecutionResult) GetDecoded() interface{}`

GetDecoded returns the Decoded field if non-nil, zero value otherwise.

### GetDecodedOk

`func (o *MethodExecutionResult) GetDecodedOk() (*interface{}, bool)`

GetDecodedOk returns a tuple with the Decoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoded

`func (o *MethodExecutionResult) SetDecoded(v interface{})`

SetDecoded sets Decoded field to given value.

### HasDecoded

`func (o *MethodExecutionResult) HasDecoded() bool`

HasDecoded returns a boolean if a field has been set.

### SetDecodedNil

`func (o *MethodExecutionResult) SetDecodedNil(b bool)`

 SetDecodedNil sets the value for Decoded to be an explicit nil

### UnsetDecoded
`func (o *MethodExecutionResult) UnsetDecoded()`

UnsetDecoded ensures that no value is present for Decoded, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


