# JettonHolders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | [**[]JettonHoldersAddressesInner**](JettonHoldersAddressesInner.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewJettonHolders

`func NewJettonHolders(addresses []JettonHoldersAddressesInner, total int64, ) *JettonHolders`

NewJettonHolders instantiates a new JettonHolders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonHoldersWithDefaults

`func NewJettonHoldersWithDefaults() *JettonHolders`

NewJettonHoldersWithDefaults instantiates a new JettonHolders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *JettonHolders) GetAddresses() []JettonHoldersAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *JettonHolders) GetAddressesOk() (*[]JettonHoldersAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *JettonHolders) SetAddresses(v []JettonHoldersAddressesInner)`

SetAddresses sets Addresses field to given value.


### GetTotal

`func (o *JettonHolders) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *JettonHolders) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *JettonHolders) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


