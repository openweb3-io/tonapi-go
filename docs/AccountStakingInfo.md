# AccountStakingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool** | **string** |  | 
**Amount** | **int64** |  | 
**PendingDeposit** | **int64** |  | 
**PendingWithdraw** | **int64** |  | 
**ReadyWithdraw** | **int64** |  | 

## Methods

### NewAccountStakingInfo

`func NewAccountStakingInfo(pool string, amount int64, pendingDeposit int64, pendingWithdraw int64, readyWithdraw int64, ) *AccountStakingInfo`

NewAccountStakingInfo instantiates a new AccountStakingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStakingInfoWithDefaults

`func NewAccountStakingInfoWithDefaults() *AccountStakingInfo`

NewAccountStakingInfoWithDefaults instantiates a new AccountStakingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool

`func (o *AccountStakingInfo) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *AccountStakingInfo) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *AccountStakingInfo) SetPool(v string)`

SetPool sets Pool field to given value.


### GetAmount

`func (o *AccountStakingInfo) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AccountStakingInfo) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AccountStakingInfo) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetPendingDeposit

`func (o *AccountStakingInfo) GetPendingDeposit() int64`

GetPendingDeposit returns the PendingDeposit field if non-nil, zero value otherwise.

### GetPendingDepositOk

`func (o *AccountStakingInfo) GetPendingDepositOk() (*int64, bool)`

GetPendingDepositOk returns a tuple with the PendingDeposit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeposit

`func (o *AccountStakingInfo) SetPendingDeposit(v int64)`

SetPendingDeposit sets PendingDeposit field to given value.


### GetPendingWithdraw

`func (o *AccountStakingInfo) GetPendingWithdraw() int64`

GetPendingWithdraw returns the PendingWithdraw field if non-nil, zero value otherwise.

### GetPendingWithdrawOk

`func (o *AccountStakingInfo) GetPendingWithdrawOk() (*int64, bool)`

GetPendingWithdrawOk returns a tuple with the PendingWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingWithdraw

`func (o *AccountStakingInfo) SetPendingWithdraw(v int64)`

SetPendingWithdraw sets PendingWithdraw field to given value.


### GetReadyWithdraw

`func (o *AccountStakingInfo) GetReadyWithdraw() int64`

GetReadyWithdraw returns the ReadyWithdraw field if non-nil, zero value otherwise.

### GetReadyWithdrawOk

`func (o *AccountStakingInfo) GetReadyWithdrawOk() (*int64, bool)`

GetReadyWithdrawOk returns a tuple with the ReadyWithdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyWithdraw

`func (o *AccountStakingInfo) SetReadyWithdraw(v int64)`

SetReadyWithdraw sets ReadyWithdraw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


