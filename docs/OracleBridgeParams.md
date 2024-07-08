# OracleBridgeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BridgeAddr** | **string** |  | 
**OracleMultisigAddress** | **string** |  | 
**ExternalChainAddress** | **string** |  | 
**Oracles** | [**[]Oracle**](Oracle.md) |  | 

## Methods

### NewOracleBridgeParams

`func NewOracleBridgeParams(bridgeAddr string, oracleMultisigAddress string, externalChainAddress string, oracles []Oracle, ) *OracleBridgeParams`

NewOracleBridgeParams instantiates a new OracleBridgeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleBridgeParamsWithDefaults

`func NewOracleBridgeParamsWithDefaults() *OracleBridgeParams`

NewOracleBridgeParamsWithDefaults instantiates a new OracleBridgeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridgeAddr

`func (o *OracleBridgeParams) GetBridgeAddr() string`

GetBridgeAddr returns the BridgeAddr field if non-nil, zero value otherwise.

### GetBridgeAddrOk

`func (o *OracleBridgeParams) GetBridgeAddrOk() (*string, bool)`

GetBridgeAddrOk returns a tuple with the BridgeAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeAddr

`func (o *OracleBridgeParams) SetBridgeAddr(v string)`

SetBridgeAddr sets BridgeAddr field to given value.


### GetOracleMultisigAddress

`func (o *OracleBridgeParams) GetOracleMultisigAddress() string`

GetOracleMultisigAddress returns the OracleMultisigAddress field if non-nil, zero value otherwise.

### GetOracleMultisigAddressOk

`func (o *OracleBridgeParams) GetOracleMultisigAddressOk() (*string, bool)`

GetOracleMultisigAddressOk returns a tuple with the OracleMultisigAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracleMultisigAddress

`func (o *OracleBridgeParams) SetOracleMultisigAddress(v string)`

SetOracleMultisigAddress sets OracleMultisigAddress field to given value.


### GetExternalChainAddress

`func (o *OracleBridgeParams) GetExternalChainAddress() string`

GetExternalChainAddress returns the ExternalChainAddress field if non-nil, zero value otherwise.

### GetExternalChainAddressOk

`func (o *OracleBridgeParams) GetExternalChainAddressOk() (*string, bool)`

GetExternalChainAddressOk returns a tuple with the ExternalChainAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalChainAddress

`func (o *OracleBridgeParams) SetExternalChainAddress(v string)`

SetExternalChainAddress sets ExternalChainAddress field to given value.


### GetOracles

`func (o *OracleBridgeParams) GetOracles() []Oracle`

GetOracles returns the Oracles field if non-nil, zero value otherwise.

### GetOraclesOk

`func (o *OracleBridgeParams) GetOraclesOk() (*[]Oracle, bool)`

GetOraclesOk returns a tuple with the Oracles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOracles

`func (o *OracleBridgeParams) SetOracles(v []Oracle)`

SetOracles sets Oracles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


