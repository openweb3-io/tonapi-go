# JettonInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mintable** | **bool** |  | 
**TotalSupply** | **string** |  | 
**Admin** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Metadata** | [**JettonMetadata**](JettonMetadata.md) |  | 
**Verification** | [**JettonVerificationType**](JettonVerificationType.md) |  | 
**HoldersCount** | **int32** |  | 

## Methods

### NewJettonInfo

`func NewJettonInfo(mintable bool, totalSupply string, metadata JettonMetadata, verification JettonVerificationType, holdersCount int32, ) *JettonInfo`

NewJettonInfo instantiates a new JettonInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonInfoWithDefaults

`func NewJettonInfoWithDefaults() *JettonInfo`

NewJettonInfoWithDefaults instantiates a new JettonInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMintable

`func (o *JettonInfo) GetMintable() bool`

GetMintable returns the Mintable field if non-nil, zero value otherwise.

### GetMintableOk

`func (o *JettonInfo) GetMintableOk() (*bool, bool)`

GetMintableOk returns a tuple with the Mintable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintable

`func (o *JettonInfo) SetMintable(v bool)`

SetMintable sets Mintable field to given value.


### GetTotalSupply

`func (o *JettonInfo) GetTotalSupply() string`

GetTotalSupply returns the TotalSupply field if non-nil, zero value otherwise.

### GetTotalSupplyOk

`func (o *JettonInfo) GetTotalSupplyOk() (*string, bool)`

GetTotalSupplyOk returns a tuple with the TotalSupply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSupply

`func (o *JettonInfo) SetTotalSupply(v string)`

SetTotalSupply sets TotalSupply field to given value.


### GetAdmin

`func (o *JettonInfo) GetAdmin() AccountAddress`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *JettonInfo) GetAdminOk() (*AccountAddress, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *JettonInfo) SetAdmin(v AccountAddress)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *JettonInfo) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetMetadata

`func (o *JettonInfo) GetMetadata() JettonMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *JettonInfo) GetMetadataOk() (*JettonMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *JettonInfo) SetMetadata(v JettonMetadata)`

SetMetadata sets Metadata field to given value.


### GetVerification

`func (o *JettonInfo) GetVerification() JettonVerificationType`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *JettonInfo) GetVerificationOk() (*JettonVerificationType, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *JettonInfo) SetVerification(v JettonVerificationType)`

SetVerification sets Verification field to given value.


### GetHoldersCount

`func (o *JettonInfo) GetHoldersCount() int32`

GetHoldersCount returns the HoldersCount field if non-nil, zero value otherwise.

### GetHoldersCountOk

`func (o *JettonInfo) GetHoldersCountOk() (*int32, bool)`

GetHoldersCountOk returns a tuple with the HoldersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoldersCount

`func (o *JettonInfo) SetHoldersCount(v int32)`

SetHoldersCount sets HoldersCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


