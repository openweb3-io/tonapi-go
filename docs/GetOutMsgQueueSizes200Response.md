# GetOutMsgQueueSizes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtMsgQueueSizeLimit** | **int32** |  | 
**Shards** | [**[]GetOutMsgQueueSizes200ResponseShardsInner**](GetOutMsgQueueSizes200ResponseShardsInner.md) |  | 

## Methods

### NewGetOutMsgQueueSizes200Response

`func NewGetOutMsgQueueSizes200Response(extMsgQueueSizeLimit int32, shards []GetOutMsgQueueSizes200ResponseShardsInner, ) *GetOutMsgQueueSizes200Response`

NewGetOutMsgQueueSizes200Response instantiates a new GetOutMsgQueueSizes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOutMsgQueueSizes200ResponseWithDefaults

`func NewGetOutMsgQueueSizes200ResponseWithDefaults() *GetOutMsgQueueSizes200Response`

NewGetOutMsgQueueSizes200ResponseWithDefaults instantiates a new GetOutMsgQueueSizes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtMsgQueueSizeLimit

`func (o *GetOutMsgQueueSizes200Response) GetExtMsgQueueSizeLimit() int32`

GetExtMsgQueueSizeLimit returns the ExtMsgQueueSizeLimit field if non-nil, zero value otherwise.

### GetExtMsgQueueSizeLimitOk

`func (o *GetOutMsgQueueSizes200Response) GetExtMsgQueueSizeLimitOk() (*int32, bool)`

GetExtMsgQueueSizeLimitOk returns a tuple with the ExtMsgQueueSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtMsgQueueSizeLimit

`func (o *GetOutMsgQueueSizes200Response) SetExtMsgQueueSizeLimit(v int32)`

SetExtMsgQueueSizeLimit sets ExtMsgQueueSizeLimit field to given value.


### GetShards

`func (o *GetOutMsgQueueSizes200Response) GetShards() []GetOutMsgQueueSizes200ResponseShardsInner`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *GetOutMsgQueueSizes200Response) GetShardsOk() (*[]GetOutMsgQueueSizes200ResponseShardsInner, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *GetOutMsgQueueSizes200Response) SetShards(v []GetOutMsgQueueSizes200ResponseShardsInner)`

SetShards sets Shards field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


