# AccountAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Name** | Pointer to **string** | Display name. Data collected from different sources like moderation lists, dns, collections names and over. | [optional] 
**IsScam** | **bool** | Is this account was marked as part of scammers activity | 
**Icon** | Pointer to **string** |  | [optional] 
**IsWallet** | **bool** |  | 

## Methods

### NewAccountAddress

`func NewAccountAddress(address string, isScam bool, isWallet bool, ) *AccountAddress`

NewAccountAddress instantiates a new AccountAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAddressWithDefaults

`func NewAccountAddressWithDefaults() *AccountAddress`

NewAccountAddressWithDefaults instantiates a new AccountAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetName

`func (o *AccountAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsScam

`func (o *AccountAddress) GetIsScam() bool`

GetIsScam returns the IsScam field if non-nil, zero value otherwise.

### GetIsScamOk

`func (o *AccountAddress) GetIsScamOk() (*bool, bool)`

GetIsScamOk returns a tuple with the IsScam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScam

`func (o *AccountAddress) SetIsScam(v bool)`

SetIsScam sets IsScam field to given value.


### GetIcon

`func (o *AccountAddress) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *AccountAddress) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *AccountAddress) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *AccountAddress) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsWallet

`func (o *AccountAddress) GetIsWallet() bool`

GetIsWallet returns the IsWallet field if non-nil, zero value otherwise.

### GetIsWalletOk

`func (o *AccountAddress) GetIsWalletOk() (*bool, bool)`

GetIsWalletOk returns a tuple with the IsWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWallet

`func (o *AccountAddress) SetIsWallet(v bool)`

SetIsWallet sets IsWallet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


