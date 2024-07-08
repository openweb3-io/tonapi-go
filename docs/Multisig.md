# Multisig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Seqno** | **int64** |  | 
**Threshold** | **int32** |  | 
**Signers** | **[]string** |  | 
**Proposers** | **[]string** |  | 
**Orders** | [**[]MultisigOrder**](MultisigOrder.md) |  | 

## Methods

### NewMultisig

`func NewMultisig(address string, seqno int64, threshold int32, signers []string, proposers []string, orders []MultisigOrder, ) *Multisig`

NewMultisig instantiates a new Multisig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultisigWithDefaults

`func NewMultisigWithDefaults() *Multisig`

NewMultisigWithDefaults instantiates a new Multisig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Multisig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Multisig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Multisig) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSeqno

`func (o *Multisig) GetSeqno() int64`

GetSeqno returns the Seqno field if non-nil, zero value otherwise.

### GetSeqnoOk

`func (o *Multisig) GetSeqnoOk() (*int64, bool)`

GetSeqnoOk returns a tuple with the Seqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeqno

`func (o *Multisig) SetSeqno(v int64)`

SetSeqno sets Seqno field to given value.


### GetThreshold

`func (o *Multisig) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *Multisig) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *Multisig) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetSigners

`func (o *Multisig) GetSigners() []string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *Multisig) GetSignersOk() (*[]string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *Multisig) SetSigners(v []string)`

SetSigners sets Signers field to given value.


### GetProposers

`func (o *Multisig) GetProposers() []string`

GetProposers returns the Proposers field if non-nil, zero value otherwise.

### GetProposersOk

`func (o *Multisig) GetProposersOk() (*[]string, bool)`

GetProposersOk returns a tuple with the Proposers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposers

`func (o *Multisig) SetProposers(v []string)`

SetProposers sets Proposers field to given value.


### GetOrders

`func (o *Multisig) GetOrders() []MultisigOrder`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Multisig) GetOrdersOk() (*[]MultisigOrder, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Multisig) SetOrders(v []MultisigOrder)`

SetOrders sets Orders field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


