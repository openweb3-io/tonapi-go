# AccountEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Account** | [**AccountAddress**](AccountAddress.md) |  | 
**Timestamp** | **int64** |  | 
**Actions** | [**[]Action**](Action.md) |  | 
**IsScam** | **bool** | scam | 
**Lt** | **int64** |  | 
**InProgress** | **bool** | Event is not finished yet. Transactions still happening | 
**Extra** | **int64** | TODO | 

## Methods

### NewAccountEvent

`func NewAccountEvent(eventId string, account AccountAddress, timestamp int64, actions []Action, isScam bool, lt int64, inProgress bool, extra int64, ) *AccountEvent`

NewAccountEvent instantiates a new AccountEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEventWithDefaults

`func NewAccountEventWithDefaults() *AccountEvent`

NewAccountEventWithDefaults instantiates a new AccountEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *AccountEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *AccountEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *AccountEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetAccount

`func (o *AccountEvent) GetAccount() AccountAddress`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AccountEvent) GetAccountOk() (*AccountAddress, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AccountEvent) SetAccount(v AccountAddress)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *AccountEvent) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AccountEvent) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AccountEvent) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetActions

`func (o *AccountEvent) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AccountEvent) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AccountEvent) SetActions(v []Action)`

SetActions sets Actions field to given value.


### GetIsScam

`func (o *AccountEvent) GetIsScam() bool`

GetIsScam returns the IsScam field if non-nil, zero value otherwise.

### GetIsScamOk

`func (o *AccountEvent) GetIsScamOk() (*bool, bool)`

GetIsScamOk returns a tuple with the IsScam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScam

`func (o *AccountEvent) SetIsScam(v bool)`

SetIsScam sets IsScam field to given value.


### GetLt

`func (o *AccountEvent) GetLt() int64`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *AccountEvent) GetLtOk() (*int64, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *AccountEvent) SetLt(v int64)`

SetLt sets Lt field to given value.


### GetInProgress

`func (o *AccountEvent) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *AccountEvent) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *AccountEvent) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.


### GetExtra

`func (o *AccountEvent) GetExtra() int64`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *AccountEvent) GetExtraOk() (*int64, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *AccountEvent) SetExtra(v int64)`

SetExtra sets Extra field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


