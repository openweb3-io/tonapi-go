# GetRawShardBlockProof200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterchainId** | [**BlockRaw**](BlockRaw.md) |  | 
**Links** | [**[]GetRawShardBlockProof200ResponseLinksInner**](GetRawShardBlockProof200ResponseLinksInner.md) |  | 

## Methods

### NewGetRawShardBlockProof200Response

`func NewGetRawShardBlockProof200Response(masterchainId BlockRaw, links []GetRawShardBlockProof200ResponseLinksInner, ) *GetRawShardBlockProof200Response`

NewGetRawShardBlockProof200Response instantiates a new GetRawShardBlockProof200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawShardBlockProof200ResponseWithDefaults

`func NewGetRawShardBlockProof200ResponseWithDefaults() *GetRawShardBlockProof200Response`

NewGetRawShardBlockProof200ResponseWithDefaults instantiates a new GetRawShardBlockProof200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterchainId

`func (o *GetRawShardBlockProof200Response) GetMasterchainId() BlockRaw`

GetMasterchainId returns the MasterchainId field if non-nil, zero value otherwise.

### GetMasterchainIdOk

`func (o *GetRawShardBlockProof200Response) GetMasterchainIdOk() (*BlockRaw, bool)`

GetMasterchainIdOk returns a tuple with the MasterchainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterchainId

`func (o *GetRawShardBlockProof200Response) SetMasterchainId(v BlockRaw)`

SetMasterchainId sets MasterchainId field to given value.


### GetLinks

`func (o *GetRawShardBlockProof200Response) GetLinks() []GetRawShardBlockProof200ResponseLinksInner`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GetRawShardBlockProof200Response) GetLinksOk() (*[]GetRawShardBlockProof200ResponseLinksInner, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GetRawShardBlockProof200Response) SetLinks(v []GetRawShardBlockProof200ResponseLinksInner)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


