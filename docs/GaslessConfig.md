# GaslessConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayAddress** | **string** | sending excess to this address decreases the commission of a gasless transfer | 
**GasJettons** | [**[]GaslessConfigGasJettonsInner**](GaslessConfigGasJettonsInner.md) | list of jettons, any of them can be used to pay for gas | 

## Methods

### NewGaslessConfig

`func NewGaslessConfig(relayAddress string, gasJettons []GaslessConfigGasJettonsInner, ) *GaslessConfig`

NewGaslessConfig instantiates a new GaslessConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaslessConfigWithDefaults

`func NewGaslessConfigWithDefaults() *GaslessConfig`

NewGaslessConfigWithDefaults instantiates a new GaslessConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayAddress

`func (o *GaslessConfig) GetRelayAddress() string`

GetRelayAddress returns the RelayAddress field if non-nil, zero value otherwise.

### GetRelayAddressOk

`func (o *GaslessConfig) GetRelayAddressOk() (*string, bool)`

GetRelayAddressOk returns a tuple with the RelayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAddress

`func (o *GaslessConfig) SetRelayAddress(v string)`

SetRelayAddress sets RelayAddress field to given value.


### GetGasJettons

`func (o *GaslessConfig) GetGasJettons() []GaslessConfigGasJettonsInner`

GetGasJettons returns the GasJettons field if non-nil, zero value otherwise.

### GetGasJettonsOk

`func (o *GaslessConfig) GetGasJettonsOk() (*[]GaslessConfigGasJettonsInner, bool)`

GetGasJettonsOk returns a tuple with the GasJettons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasJettons

`func (o *GaslessConfig) SetGasJettons(v []GaslessConfigGasJettonsInner)`

SetGasJettons sets GasJettons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


