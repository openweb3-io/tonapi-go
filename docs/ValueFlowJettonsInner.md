# ValueFlowJettonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**AccountAddress**](AccountAddress.md) |  | 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 
**Quantity** | **int64** |  | 

## Methods

### NewValueFlowJettonsInner

`func NewValueFlowJettonsInner(account AccountAddress, jetton JettonPreview, quantity int64, ) *ValueFlowJettonsInner`

NewValueFlowJettonsInner instantiates a new ValueFlowJettonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValueFlowJettonsInnerWithDefaults

`func NewValueFlowJettonsInnerWithDefaults() *ValueFlowJettonsInner`

NewValueFlowJettonsInnerWithDefaults instantiates a new ValueFlowJettonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ValueFlowJettonsInner) GetAccount() AccountAddress`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ValueFlowJettonsInner) GetAccountOk() (*AccountAddress, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ValueFlowJettonsInner) SetAccount(v AccountAddress)`

SetAccount sets Account field to given value.


### GetJetton

`func (o *ValueFlowJettonsInner) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *ValueFlowJettonsInner) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *ValueFlowJettonsInner) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.


### GetQuantity

`func (o *ValueFlowJettonsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ValueFlowJettonsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ValueFlowJettonsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


