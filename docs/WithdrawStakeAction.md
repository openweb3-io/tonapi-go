# WithdrawStakeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int64** |  | 
**Staker** | [**AccountAddress**](AccountAddress.md) |  | 
**Pool** | [**AccountAddress**](AccountAddress.md) |  | 
**Implementation** | [**PoolImplementationType**](PoolImplementationType.md) |  | 

## Methods

### NewWithdrawStakeAction

`func NewWithdrawStakeAction(amount int64, staker AccountAddress, pool AccountAddress, implementation PoolImplementationType, ) *WithdrawStakeAction`

NewWithdrawStakeAction instantiates a new WithdrawStakeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawStakeActionWithDefaults

`func NewWithdrawStakeActionWithDefaults() *WithdrawStakeAction`

NewWithdrawStakeActionWithDefaults instantiates a new WithdrawStakeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WithdrawStakeAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawStakeAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawStakeAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetStaker

`func (o *WithdrawStakeAction) GetStaker() AccountAddress`

GetStaker returns the Staker field if non-nil, zero value otherwise.

### GetStakerOk

`func (o *WithdrawStakeAction) GetStakerOk() (*AccountAddress, bool)`

GetStakerOk returns a tuple with the Staker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaker

`func (o *WithdrawStakeAction) SetStaker(v AccountAddress)`

SetStaker sets Staker field to given value.


### GetPool

`func (o *WithdrawStakeAction) GetPool() AccountAddress`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *WithdrawStakeAction) GetPoolOk() (*AccountAddress, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *WithdrawStakeAction) SetPool(v AccountAddress)`

SetPool sets Pool field to given value.


### GetImplementation

`func (o *WithdrawStakeAction) GetImplementation() PoolImplementationType`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *WithdrawStakeAction) GetImplementationOk() (*PoolImplementationType, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *WithdrawStakeAction) SetImplementation(v PoolImplementationType)`

SetImplementation sets Implementation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


