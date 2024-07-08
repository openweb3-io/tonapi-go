# Validators

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElectAt** | **int64** |  | 
**ElectClose** | **int64** |  | 
**MinStake** | **int64** |  | 
**TotalStake** | **int64** |  | 
**Validators** | [**[]Validator**](Validator.md) |  | 

## Methods

### NewValidators

`func NewValidators(electAt int64, electClose int64, minStake int64, totalStake int64, validators []Validator, ) *Validators`

NewValidators instantiates a new Validators object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorsWithDefaults

`func NewValidatorsWithDefaults() *Validators`

NewValidatorsWithDefaults instantiates a new Validators object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElectAt

`func (o *Validators) GetElectAt() int64`

GetElectAt returns the ElectAt field if non-nil, zero value otherwise.

### GetElectAtOk

`func (o *Validators) GetElectAtOk() (*int64, bool)`

GetElectAtOk returns a tuple with the ElectAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectAt

`func (o *Validators) SetElectAt(v int64)`

SetElectAt sets ElectAt field to given value.


### GetElectClose

`func (o *Validators) GetElectClose() int64`

GetElectClose returns the ElectClose field if non-nil, zero value otherwise.

### GetElectCloseOk

`func (o *Validators) GetElectCloseOk() (*int64, bool)`

GetElectCloseOk returns a tuple with the ElectClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElectClose

`func (o *Validators) SetElectClose(v int64)`

SetElectClose sets ElectClose field to given value.


### GetMinStake

`func (o *Validators) GetMinStake() int64`

GetMinStake returns the MinStake field if non-nil, zero value otherwise.

### GetMinStakeOk

`func (o *Validators) GetMinStakeOk() (*int64, bool)`

GetMinStakeOk returns a tuple with the MinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStake

`func (o *Validators) SetMinStake(v int64)`

SetMinStake sets MinStake field to given value.


### GetTotalStake

`func (o *Validators) GetTotalStake() int64`

GetTotalStake returns the TotalStake field if non-nil, zero value otherwise.

### GetTotalStakeOk

`func (o *Validators) GetTotalStakeOk() (*int64, bool)`

GetTotalStakeOk returns a tuple with the TotalStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStake

`func (o *Validators) SetTotalStake(v int64)`

SetTotalStake sets TotalStake field to given value.


### GetValidators

`func (o *Validators) GetValidators() []Validator`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *Validators) GetValidatorsOk() (*[]Validator, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *Validators) SetValidators(v []Validator)`

SetValidators sets Validators field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


