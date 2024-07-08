# WithdrawStakeRequestAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int64** |  | [optional] 
**Staker** | [**AccountAddress**](AccountAddress.md) |  | 
**Pool** | [**AccountAddress**](AccountAddress.md) |  | 
**Implementation** | [**PoolImplementationType**](PoolImplementationType.md) |  | 

## Methods

### NewWithdrawStakeRequestAction

`func NewWithdrawStakeRequestAction(staker AccountAddress, pool AccountAddress, implementation PoolImplementationType, ) *WithdrawStakeRequestAction`

NewWithdrawStakeRequestAction instantiates a new WithdrawStakeRequestAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawStakeRequestActionWithDefaults

`func NewWithdrawStakeRequestActionWithDefaults() *WithdrawStakeRequestAction`

NewWithdrawStakeRequestActionWithDefaults instantiates a new WithdrawStakeRequestAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WithdrawStakeRequestAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WithdrawStakeRequestAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WithdrawStakeRequestAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WithdrawStakeRequestAction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetStaker

`func (o *WithdrawStakeRequestAction) GetStaker() AccountAddress`

GetStaker returns the Staker field if non-nil, zero value otherwise.

### GetStakerOk

`func (o *WithdrawStakeRequestAction) GetStakerOk() (*AccountAddress, bool)`

GetStakerOk returns a tuple with the Staker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaker

`func (o *WithdrawStakeRequestAction) SetStaker(v AccountAddress)`

SetStaker sets Staker field to given value.


### GetPool

`func (o *WithdrawStakeRequestAction) GetPool() AccountAddress`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *WithdrawStakeRequestAction) GetPoolOk() (*AccountAddress, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *WithdrawStakeRequestAction) SetPool(v AccountAddress)`

SetPool sets Pool field to given value.


### GetImplementation

`func (o *WithdrawStakeRequestAction) GetImplementation() PoolImplementationType`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *WithdrawStakeRequestAction) GetImplementationOk() (*PoolImplementationType, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *WithdrawStakeRequestAction) SetImplementation(v PoolImplementationType)`

SetImplementation sets Implementation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


