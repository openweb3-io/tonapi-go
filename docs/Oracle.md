# Oracle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**SecpPubkey** | **string** |  | 

## Methods

### NewOracle

`func NewOracle(address string, secpPubkey string, ) *Oracle`

NewOracle instantiates a new Oracle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOracleWithDefaults

`func NewOracleWithDefaults() *Oracle`

NewOracleWithDefaults instantiates a new Oracle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Oracle) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Oracle) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Oracle) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetSecpPubkey

`func (o *Oracle) GetSecpPubkey() string`

GetSecpPubkey returns the SecpPubkey field if non-nil, zero value otherwise.

### GetSecpPubkeyOk

`func (o *Oracle) GetSecpPubkeyOk() (*string, bool)`

GetSecpPubkeyOk returns a tuple with the SecpPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecpPubkey

`func (o *Oracle) SetSecpPubkey(v string)`

SetSecpPubkey sets SecpPubkey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


