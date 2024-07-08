# ElectionsDepositStakeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**Staker** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewElectionsDepositStakeAction

`func NewElectionsDepositStakeAction(amount int64, staker AccountAddress, ) *ElectionsDepositStakeAction`

NewElectionsDepositStakeAction instantiates a new ElectionsDepositStakeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElectionsDepositStakeActionWithDefaults

`func NewElectionsDepositStakeActionWithDefaults() *ElectionsDepositStakeAction`

NewElectionsDepositStakeActionWithDefaults instantiates a new ElectionsDepositStakeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *ElectionsDepositStakeAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ElectionsDepositStakeAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ElectionsDepositStakeAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetStaker

`func (o *ElectionsDepositStakeAction) GetStaker() AccountAddress`

GetStaker returns the Staker field if non-nil, zero value otherwise.

### GetStakerOk

`func (o *ElectionsDepositStakeAction) GetStakerOk() (*AccountAddress, bool)`

GetStakerOk returns a tuple with the Staker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaker

`func (o *ElectionsDepositStakeAction) SetStaker(v AccountAddress)`

SetStaker sets Staker field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


