# GetRawBlockchainBlockHeader200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**BlockRaw**](BlockRaw.md) |  | 
**Mode** | **int32** |  | 
**HeaderProof** | **string** |  | 

## Methods

### NewGetRawBlockchainBlockHeader200Response

`func NewGetRawBlockchainBlockHeader200Response(id BlockRaw, mode int32, headerProof string, ) *GetRawBlockchainBlockHeader200Response`

NewGetRawBlockchainBlockHeader200Response instantiates a new GetRawBlockchainBlockHeader200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawBlockchainBlockHeader200ResponseWithDefaults

`func NewGetRawBlockchainBlockHeader200ResponseWithDefaults() *GetRawBlockchainBlockHeader200Response`

NewGetRawBlockchainBlockHeader200ResponseWithDefaults instantiates a new GetRawBlockchainBlockHeader200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRawBlockchainBlockHeader200Response) GetId() BlockRaw`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRawBlockchainBlockHeader200Response) GetIdOk() (*BlockRaw, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRawBlockchainBlockHeader200Response) SetId(v BlockRaw)`

SetId sets Id field to given value.


### GetMode

`func (o *GetRawBlockchainBlockHeader200Response) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetRawBlockchainBlockHeader200Response) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetRawBlockchainBlockHeader200Response) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetHeaderProof

`func (o *GetRawBlockchainBlockHeader200Response) GetHeaderProof() string`

GetHeaderProof returns the HeaderProof field if non-nil, zero value otherwise.

### GetHeaderProofOk

`func (o *GetRawBlockchainBlockHeader200Response) GetHeaderProofOk() (*string, bool)`

GetHeaderProofOk returns a tuple with the HeaderProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderProof

`func (o *GetRawBlockchainBlockHeader200Response) SetHeaderProof(v string)`

SetHeaderProof sets HeaderProof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


