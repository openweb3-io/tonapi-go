# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** |  | 
**Timestamp** | **int64** |  | 
**Actions** | [**[]Action**](Action.md) |  | 
**ValueFlow** | [**[]ValueFlow**](ValueFlow.md) |  | 
**IsScam** | **bool** | scam | 
**Lt** | **int64** |  | 
**InProgress** | **bool** | Event is not finished yet. Transactions still happening | 

## Methods

### NewEvent

`func NewEvent(eventId string, timestamp int64, actions []Action, valueFlow []ValueFlow, isScam bool, lt int64, inProgress bool, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *Event) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Event) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Event) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetTimestamp

`func (o *Event) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Event) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Event) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetActions

`func (o *Event) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Event) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Event) SetActions(v []Action)`

SetActions sets Actions field to given value.


### GetValueFlow

`func (o *Event) GetValueFlow() []ValueFlow`

GetValueFlow returns the ValueFlow field if non-nil, zero value otherwise.

### GetValueFlowOk

`func (o *Event) GetValueFlowOk() (*[]ValueFlow, bool)`

GetValueFlowOk returns a tuple with the ValueFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFlow

`func (o *Event) SetValueFlow(v []ValueFlow)`

SetValueFlow sets ValueFlow field to given value.


### GetIsScam

`func (o *Event) GetIsScam() bool`

GetIsScam returns the IsScam field if non-nil, zero value otherwise.

### GetIsScamOk

`func (o *Event) GetIsScamOk() (*bool, bool)`

GetIsScamOk returns a tuple with the IsScam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScam

`func (o *Event) SetIsScam(v bool)`

SetIsScam sets IsScam field to given value.


### GetLt

`func (o *Event) GetLt() int64`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *Event) GetLtOk() (*int64, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *Event) SetLt(v int64)`

SetLt sets Lt field to given value.


### GetInProgress

`func (o *Event) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *Event) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *Event) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


