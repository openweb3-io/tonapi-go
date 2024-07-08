# GetRawMasterchainInfo200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Last** | [**BlockRaw**](BlockRaw.md) |  | 
**StateRootHash** | **string** |  | 
**Init** | [**InitStateRaw**](InitStateRaw.md) |  | 

## Methods

### NewGetRawMasterchainInfo200Response

`func NewGetRawMasterchainInfo200Response(last BlockRaw, stateRootHash string, init InitStateRaw, ) *GetRawMasterchainInfo200Response`

NewGetRawMasterchainInfo200Response instantiates a new GetRawMasterchainInfo200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawMasterchainInfo200ResponseWithDefaults

`func NewGetRawMasterchainInfo200ResponseWithDefaults() *GetRawMasterchainInfo200Response`

NewGetRawMasterchainInfo200ResponseWithDefaults instantiates a new GetRawMasterchainInfo200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLast

`func (o *GetRawMasterchainInfo200Response) GetLast() BlockRaw`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetRawMasterchainInfo200Response) GetLastOk() (*BlockRaw, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetRawMasterchainInfo200Response) SetLast(v BlockRaw)`

SetLast sets Last field to given value.


### GetStateRootHash

`func (o *GetRawMasterchainInfo200Response) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *GetRawMasterchainInfo200Response) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *GetRawMasterchainInfo200Response) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetInit

`func (o *GetRawMasterchainInfo200Response) GetInit() InitStateRaw`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *GetRawMasterchainInfo200Response) GetInitOk() (*InitStateRaw, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *GetRawMasterchainInfo200Response) SetInit(v InitStateRaw)`

SetInit sets Init field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


