# AccountEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]AccountEvent**](AccountEvent.md) |  | 
**NextFrom** | **int64** |  | 

## Methods

### NewAccountEvents

`func NewAccountEvents(events []AccountEvent, nextFrom int64, ) *AccountEvents`

NewAccountEvents instantiates a new AccountEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountEventsWithDefaults

`func NewAccountEventsWithDefaults() *AccountEvents`

NewAccountEventsWithDefaults instantiates a new AccountEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *AccountEvents) GetEvents() []AccountEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *AccountEvents) GetEventsOk() (*[]AccountEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *AccountEvents) SetEvents(v []AccountEvent)`

SetEvents sets Events field to given value.


### GetNextFrom

`func (o *AccountEvents) GetNextFrom() int64`

GetNextFrom returns the NextFrom field if non-nil, zero value otherwise.

### GetNextFromOk

`func (o *AccountEvents) GetNextFromOk() (*int64, bool)`

GetNextFromOk returns a tuple with the NextFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextFrom

`func (o *AccountEvents) SetNextFrom(v int64)`

SetNextFrom sets NextFrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


