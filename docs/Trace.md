# Trace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transaction** | [**Transaction**](Transaction.md) |  | 
**Interfaces** | **[]string** |  | 
**Children** | Pointer to [**[]Trace**](Trace.md) |  | [optional] 
**Emulated** | Pointer to **bool** |  | [optional] 

## Methods

### NewTrace

`func NewTrace(transaction Transaction, interfaces []string, ) *Trace`

NewTrace instantiates a new Trace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceWithDefaults

`func NewTraceWithDefaults() *Trace`

NewTraceWithDefaults instantiates a new Trace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransaction

`func (o *Trace) GetTransaction() Transaction`

GetTransaction returns the Transaction field if non-nil, zero value otherwise.

### GetTransactionOk

`func (o *Trace) GetTransactionOk() (*Transaction, bool)`

GetTransactionOk returns a tuple with the Transaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransaction

`func (o *Trace) SetTransaction(v Transaction)`

SetTransaction sets Transaction field to given value.


### GetInterfaces

`func (o *Trace) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Trace) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Trace) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.


### GetChildren

`func (o *Trace) GetChildren() []Trace`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Trace) GetChildrenOk() (*[]Trace, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Trace) SetChildren(v []Trace)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *Trace) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetEmulated

`func (o *Trace) GetEmulated() bool`

GetEmulated returns the Emulated field if non-nil, zero value otherwise.

### GetEmulatedOk

`func (o *Trace) GetEmulatedOk() (*bool, bool)`

GetEmulatedOk returns a tuple with the Emulated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmulated

`func (o *Trace) SetEmulated(v bool)`

SetEmulated sets Emulated field to given value.

### HasEmulated

`func (o *Trace) HasEmulated() bool`

HasEmulated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


