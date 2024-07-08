# GetRawShardInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**BlockRaw**](BlockRaw.md) |  | 
**Shardblk** | [**BlockRaw**](BlockRaw.md) |  | 
**ShardProof** | **string** |  | 
**ShardDescr** | **string** |  | 

## Methods

### NewGetRawShardInfo200Response

`func NewGetRawShardInfo200Response(id BlockRaw, shardblk BlockRaw, shardProof string, shardDescr string, ) *GetRawShardInfo200Response`

NewGetRawShardInfo200Response instantiates a new GetRawShardInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawShardInfo200ResponseWithDefaults

`func NewGetRawShardInfo200ResponseWithDefaults() *GetRawShardInfo200Response`

NewGetRawShardInfo200ResponseWithDefaults instantiates a new GetRawShardInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRawShardInfo200Response) GetId() BlockRaw`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRawShardInfo200Response) GetIdOk() (*BlockRaw, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRawShardInfo200Response) SetId(v BlockRaw)`

SetId sets Id field to given value.


### GetShardblk

`func (o *GetRawShardInfo200Response) GetShardblk() BlockRaw`

GetShardblk returns the Shardblk field if non-nil, zero value otherwise.

### GetShardblkOk

`func (o *GetRawShardInfo200Response) GetShardblkOk() (*BlockRaw, bool)`

GetShardblkOk returns a tuple with the Shardblk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardblk

`func (o *GetRawShardInfo200Response) SetShardblk(v BlockRaw)`

SetShardblk sets Shardblk field to given value.


### GetShardProof

`func (o *GetRawShardInfo200Response) GetShardProof() string`

GetShardProof returns the ShardProof field if non-nil, zero value otherwise.

### GetShardProofOk

`func (o *GetRawShardInfo200Response) GetShardProofOk() (*string, bool)`

GetShardProofOk returns a tuple with the ShardProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardProof

`func (o *GetRawShardInfo200Response) SetShardProof(v string)`

SetShardProof sets ShardProof field to given value.


### GetShardDescr

`func (o *GetRawShardInfo200Response) GetShardDescr() string`

GetShardDescr returns the ShardDescr field if non-nil, zero value otherwise.

### GetShardDescrOk

`func (o *GetRawShardInfo200Response) GetShardDescrOk() (*string, bool)`

GetShardDescrOk returns a tuple with the ShardDescr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardDescr

`func (o *GetRawShardInfo200Response) SetShardDescr(v string)`

SetShardDescr sets ShardDescr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


