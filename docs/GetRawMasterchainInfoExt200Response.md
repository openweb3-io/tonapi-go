# GetRawMasterchainInfoExt200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** |  | 
**Version** | **int32** |  | 
**Capabilities** | **int64** |  | 
**Last** | [**BlockRaw**](BlockRaw.md) |  | 
**LastUtime** | **int32** |  | 
**Now** | **int32** |  | 
**StateRootHash** | **string** |  | 
**Init** | [**InitStateRaw**](InitStateRaw.md) |  | 

## Methods

### NewGetRawMasterchainInfoExt200Response

`func NewGetRawMasterchainInfoExt200Response(mode int32, version int32, capabilities int64, last BlockRaw, lastUtime int32, now int32, stateRootHash string, init InitStateRaw, ) *GetRawMasterchainInfoExt200Response`

NewGetRawMasterchainInfoExt200Response instantiates a new GetRawMasterchainInfoExt200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawMasterchainInfoExt200ResponseWithDefaults

`func NewGetRawMasterchainInfoExt200ResponseWithDefaults() *GetRawMasterchainInfoExt200Response`

NewGetRawMasterchainInfoExt200ResponseWithDefaults instantiates a new GetRawMasterchainInfoExt200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GetRawMasterchainInfoExt200Response) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GetRawMasterchainInfoExt200Response) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GetRawMasterchainInfoExt200Response) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetVersion

`func (o *GetRawMasterchainInfoExt200Response) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetRawMasterchainInfoExt200Response) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetRawMasterchainInfoExt200Response) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCapabilities

`func (o *GetRawMasterchainInfoExt200Response) GetCapabilities() int64`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetRawMasterchainInfoExt200Response) GetCapabilitiesOk() (*int64, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetRawMasterchainInfoExt200Response) SetCapabilities(v int64)`

SetCapabilities sets Capabilities field to given value.


### GetLast

`func (o *GetRawMasterchainInfoExt200Response) GetLast() BlockRaw`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *GetRawMasterchainInfoExt200Response) GetLastOk() (*BlockRaw, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *GetRawMasterchainInfoExt200Response) SetLast(v BlockRaw)`

SetLast sets Last field to given value.


### GetLastUtime

`func (o *GetRawMasterchainInfoExt200Response) GetLastUtime() int32`

GetLastUtime returns the LastUtime field if non-nil, zero value otherwise.

### GetLastUtimeOk

`func (o *GetRawMasterchainInfoExt200Response) GetLastUtimeOk() (*int32, bool)`

GetLastUtimeOk returns a tuple with the LastUtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUtime

`func (o *GetRawMasterchainInfoExt200Response) SetLastUtime(v int32)`

SetLastUtime sets LastUtime field to given value.


### GetNow

`func (o *GetRawMasterchainInfoExt200Response) GetNow() int32`

GetNow returns the Now field if non-nil, zero value otherwise.

### GetNowOk

`func (o *GetRawMasterchainInfoExt200Response) GetNowOk() (*int32, bool)`

GetNowOk returns a tuple with the Now field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNow

`func (o *GetRawMasterchainInfoExt200Response) SetNow(v int32)`

SetNow sets Now field to given value.


### GetStateRootHash

`func (o *GetRawMasterchainInfoExt200Response) GetStateRootHash() string`

GetStateRootHash returns the StateRootHash field if non-nil, zero value otherwise.

### GetStateRootHashOk

`func (o *GetRawMasterchainInfoExt200Response) GetStateRootHashOk() (*string, bool)`

GetStateRootHashOk returns a tuple with the StateRootHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateRootHash

`func (o *GetRawMasterchainInfoExt200Response) SetStateRootHash(v string)`

SetStateRootHash sets StateRootHash field to given value.


### GetInit

`func (o *GetRawMasterchainInfoExt200Response) GetInit() InitStateRaw`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *GetRawMasterchainInfoExt200Response) GetInitOk() (*InitStateRaw, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *GetRawMasterchainInfoExt200Response) SetInit(v InitStateRaw)`

SetInit sets Init field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


