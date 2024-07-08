# GetStakingPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pools** | [**[]PoolInfo**](PoolInfo.md) |  | 
**Implementations** | [**map[string]PoolImplementation**](PoolImplementation.md) |  | 

## Methods

### NewGetStakingPools200Response

`func NewGetStakingPools200Response(pools []PoolInfo, implementations map[string]PoolImplementation, ) *GetStakingPools200Response`

NewGetStakingPools200Response instantiates a new GetStakingPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStakingPools200ResponseWithDefaults

`func NewGetStakingPools200ResponseWithDefaults() *GetStakingPools200Response`

NewGetStakingPools200ResponseWithDefaults instantiates a new GetStakingPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPools

`func (o *GetStakingPools200Response) GetPools() []PoolInfo`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *GetStakingPools200Response) GetPoolsOk() (*[]PoolInfo, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *GetStakingPools200Response) SetPools(v []PoolInfo)`

SetPools sets Pools field to given value.


### GetImplementations

`func (o *GetStakingPools200Response) GetImplementations() map[string]PoolImplementation`

GetImplementations returns the Implementations field if non-nil, zero value otherwise.

### GetImplementationsOk

`func (o *GetStakingPools200Response) GetImplementationsOk() (*map[string]PoolImplementation, bool)`

GetImplementationsOk returns a tuple with the Implementations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementations

`func (o *GetStakingPools200Response) SetImplementations(v map[string]PoolImplementation)`

SetImplementations sets Implementations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


