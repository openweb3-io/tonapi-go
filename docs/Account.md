# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Balance** | **int64** |  | 
**CurrenciesBalance** | Pointer to **map[string]interface{}** | {&#39;USD&#39;: 1, &#39;IDR&#39;: 1000} | [optional] 
**LastActivity** | **int64** | unix timestamp | 
**Status** | [**AccountStatus**](AccountStatus.md) |  | 
**Interfaces** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsScam** | Pointer to **bool** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**MemoRequired** | Pointer to **bool** |  | [optional] 
**GetMethods** | **[]string** |  | 
**IsSuspended** | Pointer to **bool** |  | [optional] 
**IsWallet** | **bool** |  | 

## Methods

### NewAccount

`func NewAccount(address string, balance int64, lastActivity int64, status AccountStatus, getMethods []string, isWallet bool, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Account) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Account) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Account) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *Account) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Account) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Account) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetCurrenciesBalance

`func (o *Account) GetCurrenciesBalance() map[string]interface{}`

GetCurrenciesBalance returns the CurrenciesBalance field if non-nil, zero value otherwise.

### GetCurrenciesBalanceOk

`func (o *Account) GetCurrenciesBalanceOk() (*map[string]interface{}, bool)`

GetCurrenciesBalanceOk returns a tuple with the CurrenciesBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrenciesBalance

`func (o *Account) SetCurrenciesBalance(v map[string]interface{})`

SetCurrenciesBalance sets CurrenciesBalance field to given value.

### HasCurrenciesBalance

`func (o *Account) HasCurrenciesBalance() bool`

HasCurrenciesBalance returns a boolean if a field has been set.

### GetLastActivity

`func (o *Account) GetLastActivity() int64`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *Account) GetLastActivityOk() (*int64, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *Account) SetLastActivity(v int64)`

SetLastActivity sets LastActivity field to given value.


### GetStatus

`func (o *Account) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.


### GetInterfaces

`func (o *Account) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Account) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Account) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Account) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetName

`func (o *Account) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Account) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Account) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Account) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsScam

`func (o *Account) GetIsScam() bool`

GetIsScam returns the IsScam field if non-nil, zero value otherwise.

### GetIsScamOk

`func (o *Account) GetIsScamOk() (*bool, bool)`

GetIsScamOk returns a tuple with the IsScam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScam

`func (o *Account) SetIsScam(v bool)`

SetIsScam sets IsScam field to given value.

### HasIsScam

`func (o *Account) HasIsScam() bool`

HasIsScam returns a boolean if a field has been set.

### GetIcon

`func (o *Account) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Account) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Account) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Account) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetMemoRequired

`func (o *Account) GetMemoRequired() bool`

GetMemoRequired returns the MemoRequired field if non-nil, zero value otherwise.

### GetMemoRequiredOk

`func (o *Account) GetMemoRequiredOk() (*bool, bool)`

GetMemoRequiredOk returns a tuple with the MemoRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoRequired

`func (o *Account) SetMemoRequired(v bool)`

SetMemoRequired sets MemoRequired field to given value.

### HasMemoRequired

`func (o *Account) HasMemoRequired() bool`

HasMemoRequired returns a boolean if a field has been set.

### GetGetMethods

`func (o *Account) GetGetMethods() []string`

GetGetMethods returns the GetMethods field if non-nil, zero value otherwise.

### GetGetMethodsOk

`func (o *Account) GetGetMethodsOk() (*[]string, bool)`

GetGetMethodsOk returns a tuple with the GetMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetMethods

`func (o *Account) SetGetMethods(v []string)`

SetGetMethods sets GetMethods field to given value.


### GetIsSuspended

`func (o *Account) GetIsSuspended() bool`

GetIsSuspended returns the IsSuspended field if non-nil, zero value otherwise.

### GetIsSuspendedOk

`func (o *Account) GetIsSuspendedOk() (*bool, bool)`

GetIsSuspendedOk returns a tuple with the IsSuspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuspended

`func (o *Account) SetIsSuspended(v bool)`

SetIsSuspended sets IsSuspended field to given value.

### HasIsSuspended

`func (o *Account) HasIsSuspended() bool`

HasIsSuspended returns a boolean if a field has been set.

### GetIsWallet

`func (o *Account) GetIsWallet() bool`

GetIsWallet returns the IsWallet field if non-nil, zero value otherwise.

### GetIsWalletOk

`func (o *Account) GetIsWalletOk() (*bool, bool)`

GetIsWalletOk returns a tuple with the IsWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWallet

`func (o *Account) SetIsWallet(v bool)`

SetIsWallet sets IsWallet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


