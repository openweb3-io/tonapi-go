# MessageConsequences

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trace** | [**Trace**](Trace.md) |  | 
**Risk** | [**Risk**](Risk.md) |  | 
**Event** | [**AccountEvent**](AccountEvent.md) |  | 

## Methods

### NewMessageConsequences

`func NewMessageConsequences(trace Trace, risk Risk, event AccountEvent, ) *MessageConsequences`

NewMessageConsequences instantiates a new MessageConsequences object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageConsequencesWithDefaults

`func NewMessageConsequencesWithDefaults() *MessageConsequences`

NewMessageConsequencesWithDefaults instantiates a new MessageConsequences object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrace

`func (o *MessageConsequences) GetTrace() Trace`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *MessageConsequences) GetTraceOk() (*Trace, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *MessageConsequences) SetTrace(v Trace)`

SetTrace sets Trace field to given value.


### GetRisk

`func (o *MessageConsequences) GetRisk() Risk`

GetRisk returns the Risk field if non-nil, zero value otherwise.

### GetRiskOk

`func (o *MessageConsequences) GetRiskOk() (*Risk, bool)`

GetRiskOk returns a tuple with the Risk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRisk

`func (o *MessageConsequences) SetRisk(v Risk)`

SetRisk sets Risk field to given value.


### GetEvent

`func (o *MessageConsequences) GetEvent() AccountEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MessageConsequences) GetEventOk() (*AccountEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MessageConsequences) SetEvent(v AccountEvent)`

SetEvent sets Event field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


