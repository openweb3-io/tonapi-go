# Validator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AdnlAddress** | **string** |  | 
**Stake** | **int64** |  | 
**MaxFactor** | **int64** |  | 

## Methods

### NewValidator

`func NewValidator(address string, adnlAddress string, stake int64, maxFactor int64, ) *Validator`

NewValidator instantiates a new Validator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorWithDefaults

`func NewValidatorWithDefaults() *Validator`

NewValidatorWithDefaults instantiates a new Validator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Validator) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Validator) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Validator) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAdnlAddress

`func (o *Validator) GetAdnlAddress() string`

GetAdnlAddress returns the AdnlAddress field if non-nil, zero value otherwise.

### GetAdnlAddressOk

`func (o *Validator) GetAdnlAddressOk() (*string, bool)`

GetAdnlAddressOk returns a tuple with the AdnlAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdnlAddress

`func (o *Validator) SetAdnlAddress(v string)`

SetAdnlAddress sets AdnlAddress field to given value.


### GetStake

`func (o *Validator) GetStake() int64`

GetStake returns the Stake field if non-nil, zero value otherwise.

### GetStakeOk

`func (o *Validator) GetStakeOk() (*int64, bool)`

GetStakeOk returns a tuple with the Stake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStake

`func (o *Validator) SetStake(v int64)`

SetStake sets Stake field to given value.


### GetMaxFactor

`func (o *Validator) GetMaxFactor() int64`

GetMaxFactor returns the MaxFactor field if non-nil, zero value otherwise.

### GetMaxFactorOk

`func (o *Validator) GetMaxFactorOk() (*int64, bool)`

GetMaxFactorOk returns a tuple with the MaxFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFactor

`func (o *Validator) SetMaxFactor(v int64)`

SetMaxFactor sets MaxFactor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


