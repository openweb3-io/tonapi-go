# ValueFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountAddress**](AccountAddress.md) |  | 
**Ton** | **int64** |  | 
**Fees** | **int64** |  | 
**Jettons** | Pointer to [**[]ValueFlowJettonsInner**](ValueFlowJettonsInner.md) |  | [optional] 

## Methods

### NewValueFlow

`func NewValueFlow(account AccountAddress, ton int64, fees int64, ) *ValueFlow`

NewValueFlow instantiates a new ValueFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueFlowWithDefaults

`func NewValueFlowWithDefaults() *ValueFlow`

NewValueFlowWithDefaults instantiates a new ValueFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ValueFlow) GetAccount() AccountAddress`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ValueFlow) GetAccountOk() (*AccountAddress, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ValueFlow) SetAccount(v AccountAddress)`

SetAccount sets Account field to given value.


### GetTon

`func (o *ValueFlow) GetTon() int64`

GetTon returns the Ton field if non-nil, zero value otherwise.

### GetTonOk

`func (o *ValueFlow) GetTonOk() (*int64, bool)`

GetTonOk returns a tuple with the Ton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTon

`func (o *ValueFlow) SetTon(v int64)`

SetTon sets Ton field to given value.


### GetFees

`func (o *ValueFlow) GetFees() int64`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *ValueFlow) GetFeesOk() (*int64, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *ValueFlow) SetFees(v int64)`

SetFees sets Fees field to given value.


### GetJettons

`func (o *ValueFlow) GetJettons() []ValueFlowJettonsInner`

GetJettons returns the Jettons field if non-nil, zero value otherwise.

### GetJettonsOk

`func (o *ValueFlow) GetJettonsOk() (*[]ValueFlowJettonsInner, bool)`

GetJettonsOk returns a tuple with the Jettons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettons

`func (o *ValueFlow) SetJettons(v []ValueFlowJettonsInner)`

SetJettons sets Jettons field to given value.

### HasJettons

`func (o *ValueFlow) HasJettons() bool`

HasJettons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


