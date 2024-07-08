# GetRawAccountState200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**BlockRaw**](BlockRaw.md) |  | 
**Shardblk** | [**BlockRaw**](BlockRaw.md) |  | 
**ShardProof** | **string** |  | 
**Proof** | **string** |  | 
**State** | **string** |  | 

## Methods

### NewGetRawAccountState200Response

`func NewGetRawAccountState200Response(id BlockRaw, shardblk BlockRaw, shardProof string, proof string, state string, ) *GetRawAccountState200Response`

NewGetRawAccountState200Response instantiates a new GetRawAccountState200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawAccountState200ResponseWithDefaults

`func NewGetRawAccountState200ResponseWithDefaults() *GetRawAccountState200Response`

NewGetRawAccountState200ResponseWithDefaults instantiates a new GetRawAccountState200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRawAccountState200Response) GetId() BlockRaw`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRawAccountState200Response) GetIdOk() (*BlockRaw, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRawAccountState200Response) SetId(v BlockRaw)`

SetId sets Id field to given value.


### GetShardblk

`func (o *GetRawAccountState200Response) GetShardblk() BlockRaw`

GetShardblk returns the Shardblk field if non-nil, zero value otherwise.

### GetShardblkOk

`func (o *GetRawAccountState200Response) GetShardblkOk() (*BlockRaw, bool)`

GetShardblkOk returns a tuple with the Shardblk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardblk

`func (o *GetRawAccountState200Response) SetShardblk(v BlockRaw)`

SetShardblk sets Shardblk field to given value.


### GetShardProof

`func (o *GetRawAccountState200Response) GetShardProof() string`

GetShardProof returns the ShardProof field if non-nil, zero value otherwise.

### GetShardProofOk

`func (o *GetRawAccountState200Response) GetShardProofOk() (*string, bool)`

GetShardProofOk returns a tuple with the ShardProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShardProof

`func (o *GetRawAccountState200Response) SetShardProof(v string)`

SetShardProof sets ShardProof field to given value.


### GetProof

`func (o *GetRawAccountState200Response) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *GetRawAccountState200Response) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *GetRawAccountState200Response) SetProof(v string)`

SetProof sets Proof field to given value.


### GetState

`func (o *GetRawAccountState200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetRawAccountState200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetRawAccountState200Response) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


