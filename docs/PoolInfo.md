# PoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Name** | **string** |  | 
**TotalAmount** | **int64** |  | 
**Implementation** | [**PoolImplementationType**](PoolImplementationType.md) |  | 
**Apy** | **float32** | APY in percent | 
**MinStake** | **int64** |  | 
**CycleStart** | **int64** | current nomination cycle beginning timestamp | 
**CycleEnd** | **int64** | current nomination cycle ending timestamp | 
**Verified** | **bool** | this pool has verified source code or managed by trusted company | 
**CurrentNominators** | **int32** | current number of nominators | 
**MaxNominators** | **int32** | maximum number of nominators | 
**LiquidJettonMaster** | Pointer to **string** | for liquid staking master account of jetton | [optional] 
**NominatorsStake** | **int64** | total stake of all nominators | 
**ValidatorStake** | **int64** | stake of validator | 
**CycleLength** | Pointer to **int64** |  | [optional] 

## Methods

### NewPoolInfo

`func NewPoolInfo(address string, name string, totalAmount int64, implementation PoolImplementationType, apy float32, minStake int64, cycleStart int64, cycleEnd int64, verified bool, currentNominators int32, maxNominators int32, nominatorsStake int64, validatorStake int64, ) *PoolInfo`

NewPoolInfo instantiates a new PoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolInfoWithDefaults

`func NewPoolInfoWithDefaults() *PoolInfo`

NewPoolInfoWithDefaults instantiates a new PoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PoolInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PoolInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PoolInfo) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *PoolInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolInfo) SetName(v string)`

SetName sets Name field to given value.


### GetTotalAmount

`func (o *PoolInfo) GetTotalAmount() int64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *PoolInfo) GetTotalAmountOk() (*int64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *PoolInfo) SetTotalAmount(v int64)`

SetTotalAmount sets TotalAmount field to given value.


### GetImplementation

`func (o *PoolInfo) GetImplementation() PoolImplementationType`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *PoolInfo) GetImplementationOk() (*PoolImplementationType, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *PoolInfo) SetImplementation(v PoolImplementationType)`

SetImplementation sets Implementation field to given value.


### GetApy

`func (o *PoolInfo) GetApy() float32`

GetApy returns the Apy field if non-nil, zero value otherwise.

### GetApyOk

`func (o *PoolInfo) GetApyOk() (*float32, bool)`

GetApyOk returns a tuple with the Apy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApy

`func (o *PoolInfo) SetApy(v float32)`

SetApy sets Apy field to given value.


### GetMinStake

`func (o *PoolInfo) GetMinStake() int64`

GetMinStake returns the MinStake field if non-nil, zero value otherwise.

### GetMinStakeOk

`func (o *PoolInfo) GetMinStakeOk() (*int64, bool)`

GetMinStakeOk returns a tuple with the MinStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStake

`func (o *PoolInfo) SetMinStake(v int64)`

SetMinStake sets MinStake field to given value.


### GetCycleStart

`func (o *PoolInfo) GetCycleStart() int64`

GetCycleStart returns the CycleStart field if non-nil, zero value otherwise.

### GetCycleStartOk

`func (o *PoolInfo) GetCycleStartOk() (*int64, bool)`

GetCycleStartOk returns a tuple with the CycleStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleStart

`func (o *PoolInfo) SetCycleStart(v int64)`

SetCycleStart sets CycleStart field to given value.


### GetCycleEnd

`func (o *PoolInfo) GetCycleEnd() int64`

GetCycleEnd returns the CycleEnd field if non-nil, zero value otherwise.

### GetCycleEndOk

`func (o *PoolInfo) GetCycleEndOk() (*int64, bool)`

GetCycleEndOk returns a tuple with the CycleEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleEnd

`func (o *PoolInfo) SetCycleEnd(v int64)`

SetCycleEnd sets CycleEnd field to given value.


### GetVerified

`func (o *PoolInfo) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *PoolInfo) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *PoolInfo) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetCurrentNominators

`func (o *PoolInfo) GetCurrentNominators() int32`

GetCurrentNominators returns the CurrentNominators field if non-nil, zero value otherwise.

### GetCurrentNominatorsOk

`func (o *PoolInfo) GetCurrentNominatorsOk() (*int32, bool)`

GetCurrentNominatorsOk returns a tuple with the CurrentNominators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNominators

`func (o *PoolInfo) SetCurrentNominators(v int32)`

SetCurrentNominators sets CurrentNominators field to given value.


### GetMaxNominators

`func (o *PoolInfo) GetMaxNominators() int32`

GetMaxNominators returns the MaxNominators field if non-nil, zero value otherwise.

### GetMaxNominatorsOk

`func (o *PoolInfo) GetMaxNominatorsOk() (*int32, bool)`

GetMaxNominatorsOk returns a tuple with the MaxNominators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNominators

`func (o *PoolInfo) SetMaxNominators(v int32)`

SetMaxNominators sets MaxNominators field to given value.


### GetLiquidJettonMaster

`func (o *PoolInfo) GetLiquidJettonMaster() string`

GetLiquidJettonMaster returns the LiquidJettonMaster field if non-nil, zero value otherwise.

### GetLiquidJettonMasterOk

`func (o *PoolInfo) GetLiquidJettonMasterOk() (*string, bool)`

GetLiquidJettonMasterOk returns a tuple with the LiquidJettonMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidJettonMaster

`func (o *PoolInfo) SetLiquidJettonMaster(v string)`

SetLiquidJettonMaster sets LiquidJettonMaster field to given value.

### HasLiquidJettonMaster

`func (o *PoolInfo) HasLiquidJettonMaster() bool`

HasLiquidJettonMaster returns a boolean if a field has been set.

### GetNominatorsStake

`func (o *PoolInfo) GetNominatorsStake() int64`

GetNominatorsStake returns the NominatorsStake field if non-nil, zero value otherwise.

### GetNominatorsStakeOk

`func (o *PoolInfo) GetNominatorsStakeOk() (*int64, bool)`

GetNominatorsStakeOk returns a tuple with the NominatorsStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNominatorsStake

`func (o *PoolInfo) SetNominatorsStake(v int64)`

SetNominatorsStake sets NominatorsStake field to given value.


### GetValidatorStake

`func (o *PoolInfo) GetValidatorStake() int64`

GetValidatorStake returns the ValidatorStake field if non-nil, zero value otherwise.

### GetValidatorStakeOk

`func (o *PoolInfo) GetValidatorStakeOk() (*int64, bool)`

GetValidatorStakeOk returns a tuple with the ValidatorStake field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatorStake

`func (o *PoolInfo) SetValidatorStake(v int64)`

SetValidatorStake sets ValidatorStake field to given value.


### GetCycleLength

`func (o *PoolInfo) GetCycleLength() int64`

GetCycleLength returns the CycleLength field if non-nil, zero value otherwise.

### GetCycleLengthOk

`func (o *PoolInfo) GetCycleLengthOk() (*int64, bool)`

GetCycleLengthOk returns a tuple with the CycleLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleLength

`func (o *PoolInfo) SetCycleLength(v int64)`

SetCycleLength sets CycleLength field to given value.

### HasCycleLength

`func (o *PoolInfo) HasCycleLength() bool`

HasCycleLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


