# TraceID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Utime** | **int64** |  | 

## Methods

### NewTraceID

`func NewTraceID(id string, utime int64, ) *TraceID`

NewTraceID instantiates a new TraceID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceIDWithDefaults

`func NewTraceIDWithDefaults() *TraceID`

NewTraceIDWithDefaults instantiates a new TraceID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TraceID) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TraceID) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TraceID) SetId(v string)`

SetId sets Id field to given value.


### GetUtime

`func (o *TraceID) GetUtime() int64`

GetUtime returns the Utime field if non-nil, zero value otherwise.

### GetUtimeOk

`func (o *TraceID) GetUtimeOk() (*int64, bool)`

GetUtimeOk returns a tuple with the Utime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtime

`func (o *TraceID) SetUtime(v int64)`

SetUtime sets Utime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


