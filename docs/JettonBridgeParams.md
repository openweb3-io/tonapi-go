# JettonBridgeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeAddress** | **string** |  | 
**OraclesAddress** | **string** |  | 
**StateFlags** | **int32** |  | 
**BurnBridgeFee** | Pointer to **int64** |  | [optional] 
**Oracles** | [**[]Oracle**](Oracle.md) |  | 
**ExternalChainAddress** | Pointer to **string** |  | [optional] 
**Prices** | Pointer to [**JettonBridgePrices**](JettonBridgePrices.md) |  | [optional] 

## Methods

### NewJettonBridgeParams

`func NewJettonBridgeParams(bridgeAddress string, oraclesAddress string, stateFlags int32, oracles []Oracle, ) *JettonBridgeParams`

NewJettonBridgeParams instantiates a new JettonBridgeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonBridgeParamsWithDefaults

`func NewJettonBridgeParamsWithDefaults() *JettonBridgeParams`

NewJettonBridgeParamsWithDefaults instantiates a new JettonBridgeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeAddress

`func (o *JettonBridgeParams) GetBridgeAddress() string`

GetBridgeAddress returns the BridgeAddress field if non-nil, zero value otherwise.

### GetBridgeAddressOk

`func (o *JettonBridgeParams) GetBridgeAddressOk() (*string, bool)`

GetBridgeAddressOk returns a tuple with the BridgeAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeAddress

`func (o *JettonBridgeParams) SetBridgeAddress(v string)`

SetBridgeAddress sets BridgeAddress field to given value.


### GetOraclesAddress

`func (o *JettonBridgeParams) GetOraclesAddress() string`

GetOraclesAddress returns the OraclesAddress field if non-nil, zero value otherwise.

### GetOraclesAddressOk

`func (o *JettonBridgeParams) GetOraclesAddressOk() (*string, bool)`

GetOraclesAddressOk returns a tuple with the OraclesAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOraclesAddress

`func (o *JettonBridgeParams) SetOraclesAddress(v string)`

SetOraclesAddress sets OraclesAddress field to given value.


### GetStateFlags

`func (o *JettonBridgeParams) GetStateFlags() int32`

GetStateFlags returns the StateFlags field if non-nil, zero value otherwise.

### GetStateFlagsOk

`func (o *JettonBridgeParams) GetStateFlagsOk() (*int32, bool)`

GetStateFlagsOk returns a tuple with the StateFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFlags

`func (o *JettonBridgeParams) SetStateFlags(v int32)`

SetStateFlags sets StateFlags field to given value.


### GetBurnBridgeFee

`func (o *JettonBridgeParams) GetBurnBridgeFee() int64`

GetBurnBridgeFee returns the BurnBridgeFee field if non-nil, zero value otherwise.

### GetBurnBridgeFeeOk

`func (o *JettonBridgeParams) GetBurnBridgeFeeOk() (*int64, bool)`

GetBurnBridgeFeeOk returns a tuple with the BurnBridgeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnBridgeFee

`func (o *JettonBridgeParams) SetBurnBridgeFee(v int64)`

SetBurnBridgeFee sets BurnBridgeFee field to given value.

### HasBurnBridgeFee

`func (o *JettonBridgeParams) HasBurnBridgeFee() bool`

HasBurnBridgeFee returns a boolean if a field has been set.

### GetOracles

`func (o *JettonBridgeParams) GetOracles() []Oracle`

GetOracles returns the Oracles field if non-nil, zero value otherwise.

### GetOraclesOk

`func (o *JettonBridgeParams) GetOraclesOk() (*[]Oracle, bool)`

GetOraclesOk returns a tuple with the Oracles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracles

`func (o *JettonBridgeParams) SetOracles(v []Oracle)`

SetOracles sets Oracles field to given value.


### GetExternalChainAddress

`func (o *JettonBridgeParams) GetExternalChainAddress() string`

GetExternalChainAddress returns the ExternalChainAddress field if non-nil, zero value otherwise.

### GetExternalChainAddressOk

`func (o *JettonBridgeParams) GetExternalChainAddressOk() (*string, bool)`

GetExternalChainAddressOk returns a tuple with the ExternalChainAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChainAddress

`func (o *JettonBridgeParams) SetExternalChainAddress(v string)`

SetExternalChainAddress sets ExternalChainAddress field to given value.

### HasExternalChainAddress

`func (o *JettonBridgeParams) HasExternalChainAddress() bool`

HasExternalChainAddress returns a boolean if a field has been set.

### GetPrices

`func (o *JettonBridgeParams) GetPrices() JettonBridgePrices`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *JettonBridgeParams) GetPricesOk() (*JettonBridgePrices, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *JettonBridgeParams) SetPrices(v JettonBridgePrices)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *JettonBridgeParams) HasPrices() bool`

HasPrices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


