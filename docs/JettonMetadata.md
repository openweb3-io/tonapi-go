# JettonMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Name** | **string** |  | 
**Symbol** | **string** |  | 
**Decimals** | **string** |  | 
**Image** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Social** | Pointer to **[]string** |  | [optional] 
**Websites** | Pointer to **[]string** |  | [optional] 
**Catalogs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJettonMetadata

`func NewJettonMetadata(address string, name string, symbol string, decimals string, ) *JettonMetadata`

NewJettonMetadata instantiates a new JettonMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonMetadataWithDefaults

`func NewJettonMetadataWithDefaults() *JettonMetadata`

NewJettonMetadataWithDefaults instantiates a new JettonMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *JettonMetadata) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *JettonMetadata) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *JettonMetadata) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *JettonMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JettonMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JettonMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetSymbol

`func (o *JettonMetadata) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *JettonMetadata) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *JettonMetadata) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetDecimals

`func (o *JettonMetadata) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *JettonMetadata) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *JettonMetadata) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.


### GetImage

`func (o *JettonMetadata) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *JettonMetadata) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *JettonMetadata) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *JettonMetadata) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetDescription

`func (o *JettonMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *JettonMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *JettonMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *JettonMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSocial

`func (o *JettonMetadata) GetSocial() []string`

GetSocial returns the Social field if non-nil, zero value otherwise.

### GetSocialOk

`func (o *JettonMetadata) GetSocialOk() (*[]string, bool)`

GetSocialOk returns a tuple with the Social field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocial

`func (o *JettonMetadata) SetSocial(v []string)`

SetSocial sets Social field to given value.

### HasSocial

`func (o *JettonMetadata) HasSocial() bool`

HasSocial returns a boolean if a field has been set.

### GetWebsites

`func (o *JettonMetadata) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *JettonMetadata) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *JettonMetadata) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *JettonMetadata) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### GetCatalogs

`func (o *JettonMetadata) GetCatalogs() []string`

GetCatalogs returns the Catalogs field if non-nil, zero value otherwise.

### GetCatalogsOk

`func (o *JettonMetadata) GetCatalogsOk() (*[]string, bool)`

GetCatalogsOk returns a tuple with the Catalogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogs

`func (o *JettonMetadata) SetCatalogs(v []string)`

SetCatalogs sets Catalogs field to given value.

### HasCatalogs

`func (o *JettonMetadata) HasCatalogs() bool`

HasCatalogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


