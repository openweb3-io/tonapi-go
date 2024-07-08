# GetRawConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** |  | 
**Id** | [**BlockRaw**](BlockRaw.md) |  | 
**StateProof** | **string** |  | 
**ConfigProof** | **string** |  | 

## Methods

### NewGetRawConfig200Response

`func NewGetRawConfig200Response(mode int32, id BlockRaw, stateProof string, configProof string, ) *GetRawConfig200Response`

NewGetRawConfig200Response instantiates a new GetRawConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawConfig200ResponseWithDefaults

`func NewGetRawConfig200ResponseWithDefaults() *GetRawConfig200Response`

NewGetRawConfig200ResponseWithDefaults instantiates a new GetRawConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetRawConfig200Response) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetRawConfig200Response) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetRawConfig200Response) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetId

`func (o *GetRawConfig200Response) GetId() BlockRaw`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRawConfig200Response) GetIdOk() (*BlockRaw, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRawConfig200Response) SetId(v BlockRaw)`

SetId sets Id field to given value.


### GetStateProof

`func (o *GetRawConfig200Response) GetStateProof() string`

GetStateProof returns the StateProof field if non-nil, zero value otherwise.

### GetStateProofOk

`func (o *GetRawConfig200Response) GetStateProofOk() (*string, bool)`

GetStateProofOk returns a tuple with the StateProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProof

`func (o *GetRawConfig200Response) SetStateProof(v string)`

SetStateProof sets StateProof field to given value.


### GetConfigProof

`func (o *GetRawConfig200Response) GetConfigProof() string`

GetConfigProof returns the ConfigProof field if non-nil, zero value otherwise.

### GetConfigProofOk

`func (o *GetRawConfig200Response) GetConfigProofOk() (*string, bool)`

GetConfigProofOk returns a tuple with the ConfigProof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProof

`func (o *GetRawConfig200Response) SetConfigProof(v string)`

SetConfigProof sets ConfigProof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


