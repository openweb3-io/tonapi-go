# WalletDNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Account** | [**AccountAddress**](AccountAddress.md) |  | 
**IsWallet** | **bool** |  | 
**HasMethodPubkey** | **bool** |  | 
**HasMethodSeqno** | **bool** |  | 
**Names** | **[]string** |  | 

## Methods

### NewWalletDNS

`func NewWalletDNS(address string, account AccountAddress, isWallet bool, hasMethodPubkey bool, hasMethodSeqno bool, names []string, ) *WalletDNS`

NewWalletDNS instantiates a new WalletDNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletDNSWithDefaults

`func NewWalletDNSWithDefaults() *WalletDNS`

NewWalletDNSWithDefaults instantiates a new WalletDNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *WalletDNS) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *WalletDNS) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *WalletDNS) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAccount

`func (o *WalletDNS) GetAccount() AccountAddress`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WalletDNS) GetAccountOk() (*AccountAddress, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WalletDNS) SetAccount(v AccountAddress)`

SetAccount sets Account field to given value.


### GetIsWallet

`func (o *WalletDNS) GetIsWallet() bool`

GetIsWallet returns the IsWallet field if non-nil, zero value otherwise.

### GetIsWalletOk

`func (o *WalletDNS) GetIsWalletOk() (*bool, bool)`

GetIsWalletOk returns a tuple with the IsWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWallet

`func (o *WalletDNS) SetIsWallet(v bool)`

SetIsWallet sets IsWallet field to given value.


### GetHasMethodPubkey

`func (o *WalletDNS) GetHasMethodPubkey() bool`

GetHasMethodPubkey returns the HasMethodPubkey field if non-nil, zero value otherwise.

### GetHasMethodPubkeyOk

`func (o *WalletDNS) GetHasMethodPubkeyOk() (*bool, bool)`

GetHasMethodPubkeyOk returns a tuple with the HasMethodPubkey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMethodPubkey

`func (o *WalletDNS) SetHasMethodPubkey(v bool)`

SetHasMethodPubkey sets HasMethodPubkey field to given value.


### GetHasMethodSeqno

`func (o *WalletDNS) GetHasMethodSeqno() bool`

GetHasMethodSeqno returns the HasMethodSeqno field if non-nil, zero value otherwise.

### GetHasMethodSeqnoOk

`func (o *WalletDNS) GetHasMethodSeqnoOk() (*bool, bool)`

GetHasMethodSeqnoOk returns a tuple with the HasMethodSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMethodSeqno

`func (o *WalletDNS) SetHasMethodSeqno(v bool)`

SetHasMethodSeqno sets HasMethodSeqno field to given value.


### GetNames

`func (o *WalletDNS) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *WalletDNS) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *WalletDNS) SetNames(v []string)`

SetNames sets Names field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


