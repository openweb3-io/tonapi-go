# JettonBalanceLock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** |  | 
**Till** | **int64** |  | 

## Methods

### NewJettonBalanceLock

`func NewJettonBalanceLock(amount string, till int64, ) *JettonBalanceLock`

NewJettonBalanceLock instantiates a new JettonBalanceLock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonBalanceLockWithDefaults

`func NewJettonBalanceLockWithDefaults() *JettonBalanceLock`

NewJettonBalanceLockWithDefaults instantiates a new JettonBalanceLock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *JettonBalanceLock) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JettonBalanceLock) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JettonBalanceLock) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetTill

`func (o *JettonBalanceLock) GetTill() int64`

GetTill returns the Till field if non-nil, zero value otherwise.

### GetTillOk

`func (o *JettonBalanceLock) GetTillOk() (*int64, bool)`

GetTillOk returns a tuple with the Till field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTill

`func (o *JettonBalanceLock) SetTill(v int64)`

SetTill sets Till field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


