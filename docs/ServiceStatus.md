# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestOnline** | **bool** |  | [default to true]
**IndexingLatency** | **int32** |  | 

## Methods

### NewServiceStatus

`func NewServiceStatus(restOnline bool, indexingLatency int32, ) *ServiceStatus`

NewServiceStatus instantiates a new ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusWithDefaults

`func NewServiceStatusWithDefaults() *ServiceStatus`

NewServiceStatusWithDefaults instantiates a new ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestOnline

`func (o *ServiceStatus) GetRestOnline() bool`

GetRestOnline returns the RestOnline field if non-nil, zero value otherwise.

### GetRestOnlineOk

`func (o *ServiceStatus) GetRestOnlineOk() (*bool, bool)`

GetRestOnlineOk returns a tuple with the RestOnline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestOnline

`func (o *ServiceStatus) SetRestOnline(v bool)`

SetRestOnline sets RestOnline field to given value.


### GetIndexingLatency

`func (o *ServiceStatus) GetIndexingLatency() int32`

GetIndexingLatency returns the IndexingLatency field if non-nil, zero value otherwise.

### GetIndexingLatencyOk

`func (o *ServiceStatus) GetIndexingLatencyOk() (*int32, bool)`

GetIndexingLatencyOk returns a tuple with the IndexingLatency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexingLatency

`func (o *ServiceStatus) SetIndexingLatency(v int32)`

SetIndexingLatency sets IndexingLatency field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


