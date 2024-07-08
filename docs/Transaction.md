# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** |  | 
**Lt** | **int64** |  | 
**Account** | [**AccountAddress**](AccountAddress.md) |  | 
**Success** | **bool** |  | 
**Utime** | **int64** |  | 
**OrigStatus** | [**AccountStatus**](AccountStatus.md) |  | 
**EndStatus** | [**AccountStatus**](AccountStatus.md) |  | 
**TotalFees** | **int64** |  | 
**EndBalance** | **int64** |  | 
**TransactionType** | [**TransactionType**](TransactionType.md) |  | 
**StateUpdateOld** | **string** |  | 
**StateUpdateNew** | **string** |  | 
**InMsg** | Pointer to [**Message**](Message.md) |  | [optional] 
**OutMsgs** | [**[]Message**](Message.md) |  | 
**Block** | **string** |  | 
**PrevTransHash** | Pointer to **string** |  | [optional] 
**PrevTransLt** | Pointer to **int64** |  | [optional] 
**ComputePhase** | Pointer to [**ComputePhase**](ComputePhase.md) |  | [optional] 
**StoragePhase** | Pointer to [**StoragePhase**](StoragePhase.md) |  | [optional] 
**CreditPhase** | Pointer to [**CreditPhase**](CreditPhase.md) |  | [optional] 
**ActionPhase** | Pointer to [**ActionPhase**](ActionPhase.md) |  | [optional] 
**BouncePhase** | Pointer to [**BouncePhaseType**](BouncePhaseType.md) |  | [optional] 
**Aborted** | **bool** |  | 
**Destroyed** | **bool** |  | 
**Raw** | **string** | hex encoded boc with raw transaction | 

## Methods

### NewTransaction

`func NewTransaction(hash string, lt int64, account AccountAddress, success bool, utime int64, origStatus AccountStatus, endStatus AccountStatus, totalFees int64, endBalance int64, transactionType TransactionType, stateUpdateOld string, stateUpdateNew string, outMsgs []Message, block string, aborted bool, destroyed bool, raw string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *Transaction) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Transaction) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Transaction) SetHash(v string)`

SetHash sets Hash field to given value.


### GetLt

`func (o *Transaction) GetLt() int64`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Transaction) GetLtOk() (*int64, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Transaction) SetLt(v int64)`

SetLt sets Lt field to given value.


### GetAccount

`func (o *Transaction) GetAccount() AccountAddress`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Transaction) GetAccountOk() (*AccountAddress, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Transaction) SetAccount(v AccountAddress)`

SetAccount sets Account field to given value.


### GetSuccess

`func (o *Transaction) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Transaction) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Transaction) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetUtime

`func (o *Transaction) GetUtime() int64`

GetUtime returns the Utime field if non-nil, zero value otherwise.

### GetUtimeOk

`func (o *Transaction) GetUtimeOk() (*int64, bool)`

GetUtimeOk returns a tuple with the Utime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtime

`func (o *Transaction) SetUtime(v int64)`

SetUtime sets Utime field to given value.


### GetOrigStatus

`func (o *Transaction) GetOrigStatus() AccountStatus`

GetOrigStatus returns the OrigStatus field if non-nil, zero value otherwise.

### GetOrigStatusOk

`func (o *Transaction) GetOrigStatusOk() (*AccountStatus, bool)`

GetOrigStatusOk returns a tuple with the OrigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigStatus

`func (o *Transaction) SetOrigStatus(v AccountStatus)`

SetOrigStatus sets OrigStatus field to given value.


### GetEndStatus

`func (o *Transaction) GetEndStatus() AccountStatus`

GetEndStatus returns the EndStatus field if non-nil, zero value otherwise.

### GetEndStatusOk

`func (o *Transaction) GetEndStatusOk() (*AccountStatus, bool)`

GetEndStatusOk returns a tuple with the EndStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndStatus

`func (o *Transaction) SetEndStatus(v AccountStatus)`

SetEndStatus sets EndStatus field to given value.


### GetTotalFees

`func (o *Transaction) GetTotalFees() int64`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *Transaction) GetTotalFeesOk() (*int64, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *Transaction) SetTotalFees(v int64)`

SetTotalFees sets TotalFees field to given value.


### GetEndBalance

`func (o *Transaction) GetEndBalance() int64`

GetEndBalance returns the EndBalance field if non-nil, zero value otherwise.

### GetEndBalanceOk

`func (o *Transaction) GetEndBalanceOk() (*int64, bool)`

GetEndBalanceOk returns a tuple with the EndBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBalance

`func (o *Transaction) SetEndBalance(v int64)`

SetEndBalance sets EndBalance field to given value.


### GetTransactionType

`func (o *Transaction) GetTransactionType() TransactionType`

GetTransactionType returns the TransactionType field if non-nil, zero value otherwise.

### GetTransactionTypeOk

`func (o *Transaction) GetTransactionTypeOk() (*TransactionType, bool)`

GetTransactionTypeOk returns a tuple with the TransactionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionType

`func (o *Transaction) SetTransactionType(v TransactionType)`

SetTransactionType sets TransactionType field to given value.


### GetStateUpdateOld

`func (o *Transaction) GetStateUpdateOld() string`

GetStateUpdateOld returns the StateUpdateOld field if non-nil, zero value otherwise.

### GetStateUpdateOldOk

`func (o *Transaction) GetStateUpdateOldOk() (*string, bool)`

GetStateUpdateOldOk returns a tuple with the StateUpdateOld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdateOld

`func (o *Transaction) SetStateUpdateOld(v string)`

SetStateUpdateOld sets StateUpdateOld field to given value.


### GetStateUpdateNew

`func (o *Transaction) GetStateUpdateNew() string`

GetStateUpdateNew returns the StateUpdateNew field if non-nil, zero value otherwise.

### GetStateUpdateNewOk

`func (o *Transaction) GetStateUpdateNewOk() (*string, bool)`

GetStateUpdateNewOk returns a tuple with the StateUpdateNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateUpdateNew

`func (o *Transaction) SetStateUpdateNew(v string)`

SetStateUpdateNew sets StateUpdateNew field to given value.


### GetInMsg

`func (o *Transaction) GetInMsg() Message`

GetInMsg returns the InMsg field if non-nil, zero value otherwise.

### GetInMsgOk

`func (o *Transaction) GetInMsgOk() (*Message, bool)`

GetInMsgOk returns a tuple with the InMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMsg

`func (o *Transaction) SetInMsg(v Message)`

SetInMsg sets InMsg field to given value.

### HasInMsg

`func (o *Transaction) HasInMsg() bool`

HasInMsg returns a boolean if a field has been set.

### GetOutMsgs

`func (o *Transaction) GetOutMsgs() []Message`

GetOutMsgs returns the OutMsgs field if non-nil, zero value otherwise.

### GetOutMsgsOk

`func (o *Transaction) GetOutMsgsOk() (*[]Message, bool)`

GetOutMsgsOk returns a tuple with the OutMsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutMsgs

`func (o *Transaction) SetOutMsgs(v []Message)`

SetOutMsgs sets OutMsgs field to given value.


### GetBlock

`func (o *Transaction) GetBlock() string`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *Transaction) GetBlockOk() (*string, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *Transaction) SetBlock(v string)`

SetBlock sets Block field to given value.


### GetPrevTransHash

`func (o *Transaction) GetPrevTransHash() string`

GetPrevTransHash returns the PrevTransHash field if non-nil, zero value otherwise.

### GetPrevTransHashOk

`func (o *Transaction) GetPrevTransHashOk() (*string, bool)`

GetPrevTransHashOk returns a tuple with the PrevTransHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTransHash

`func (o *Transaction) SetPrevTransHash(v string)`

SetPrevTransHash sets PrevTransHash field to given value.

### HasPrevTransHash

`func (o *Transaction) HasPrevTransHash() bool`

HasPrevTransHash returns a boolean if a field has been set.

### GetPrevTransLt

`func (o *Transaction) GetPrevTransLt() int64`

GetPrevTransLt returns the PrevTransLt field if non-nil, zero value otherwise.

### GetPrevTransLtOk

`func (o *Transaction) GetPrevTransLtOk() (*int64, bool)`

GetPrevTransLtOk returns a tuple with the PrevTransLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevTransLt

`func (o *Transaction) SetPrevTransLt(v int64)`

SetPrevTransLt sets PrevTransLt field to given value.

### HasPrevTransLt

`func (o *Transaction) HasPrevTransLt() bool`

HasPrevTransLt returns a boolean if a field has been set.

### GetComputePhase

`func (o *Transaction) GetComputePhase() ComputePhase`

GetComputePhase returns the ComputePhase field if non-nil, zero value otherwise.

### GetComputePhaseOk

`func (o *Transaction) GetComputePhaseOk() (*ComputePhase, bool)`

GetComputePhaseOk returns a tuple with the ComputePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputePhase

`func (o *Transaction) SetComputePhase(v ComputePhase)`

SetComputePhase sets ComputePhase field to given value.

### HasComputePhase

`func (o *Transaction) HasComputePhase() bool`

HasComputePhase returns a boolean if a field has been set.

### GetStoragePhase

`func (o *Transaction) GetStoragePhase() StoragePhase`

GetStoragePhase returns the StoragePhase field if non-nil, zero value otherwise.

### GetStoragePhaseOk

`func (o *Transaction) GetStoragePhaseOk() (*StoragePhase, bool)`

GetStoragePhaseOk returns a tuple with the StoragePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePhase

`func (o *Transaction) SetStoragePhase(v StoragePhase)`

SetStoragePhase sets StoragePhase field to given value.

### HasStoragePhase

`func (o *Transaction) HasStoragePhase() bool`

HasStoragePhase returns a boolean if a field has been set.

### GetCreditPhase

`func (o *Transaction) GetCreditPhase() CreditPhase`

GetCreditPhase returns the CreditPhase field if non-nil, zero value otherwise.

### GetCreditPhaseOk

`func (o *Transaction) GetCreditPhaseOk() (*CreditPhase, bool)`

GetCreditPhaseOk returns a tuple with the CreditPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditPhase

`func (o *Transaction) SetCreditPhase(v CreditPhase)`

SetCreditPhase sets CreditPhase field to given value.

### HasCreditPhase

`func (o *Transaction) HasCreditPhase() bool`

HasCreditPhase returns a boolean if a field has been set.

### GetActionPhase

`func (o *Transaction) GetActionPhase() ActionPhase`

GetActionPhase returns the ActionPhase field if non-nil, zero value otherwise.

### GetActionPhaseOk

`func (o *Transaction) GetActionPhaseOk() (*ActionPhase, bool)`

GetActionPhaseOk returns a tuple with the ActionPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPhase

`func (o *Transaction) SetActionPhase(v ActionPhase)`

SetActionPhase sets ActionPhase field to given value.

### HasActionPhase

`func (o *Transaction) HasActionPhase() bool`

HasActionPhase returns a boolean if a field has been set.

### GetBouncePhase

`func (o *Transaction) GetBouncePhase() BouncePhaseType`

GetBouncePhase returns the BouncePhase field if non-nil, zero value otherwise.

### GetBouncePhaseOk

`func (o *Transaction) GetBouncePhaseOk() (*BouncePhaseType, bool)`

GetBouncePhaseOk returns a tuple with the BouncePhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBouncePhase

`func (o *Transaction) SetBouncePhase(v BouncePhaseType)`

SetBouncePhase sets BouncePhase field to given value.

### HasBouncePhase

`func (o *Transaction) HasBouncePhase() bool`

HasBouncePhase returns a boolean if a field has been set.

### GetAborted

`func (o *Transaction) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *Transaction) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *Transaction) SetAborted(v bool)`

SetAborted sets Aborted field to given value.


### GetDestroyed

`func (o *Transaction) GetDestroyed() bool`

GetDestroyed returns the Destroyed field if non-nil, zero value otherwise.

### GetDestroyedOk

`func (o *Transaction) GetDestroyedOk() (*bool, bool)`

GetDestroyedOk returns a tuple with the Destroyed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyed

`func (o *Transaction) SetDestroyed(v bool)`

SetDestroyed sets Destroyed field to given value.


### GetRaw

`func (o *Transaction) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *Transaction) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *Transaction) SetRaw(v string)`

SetRaw sets Raw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


