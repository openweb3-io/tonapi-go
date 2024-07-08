# JettonBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** |  | 
**Price** | Pointer to [**TokenRates**](TokenRates.md) |  | [optional] 
**WalletAddress** | [**AccountAddress**](AccountAddress.md) |  | 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 
**Lock** | Pointer to [**JettonBalanceLock**](JettonBalanceLock.md) |  | [optional] 

## Methods

### NewJettonBalance

`func NewJettonBalance(balance string, walletAddress AccountAddress, jetton JettonPreview, ) *JettonBalance`

NewJettonBalance instantiates a new JettonBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonBalanceWithDefaults

`func NewJettonBalanceWithDefaults() *JettonBalance`

NewJettonBalanceWithDefaults instantiates a new JettonBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *JettonBalance) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *JettonBalance) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *JettonBalance) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetPrice

`func (o *JettonBalance) GetPrice() TokenRates`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *JettonBalance) GetPriceOk() (*TokenRates, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *JettonBalance) SetPrice(v TokenRates)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *JettonBalance) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetWalletAddress

`func (o *JettonBalance) GetWalletAddress() AccountAddress`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *JettonBalance) GetWalletAddressOk() (*AccountAddress, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *JettonBalance) SetWalletAddress(v AccountAddress)`

SetWalletAddress sets WalletAddress field to given value.


### GetJetton

`func (o *JettonBalance) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *JettonBalance) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *JettonBalance) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.


### GetLock

`func (o *JettonBalance) GetLock() JettonBalanceLock`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *JettonBalance) GetLockOk() (*JettonBalanceLock, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *JettonBalance) SetLock(v JettonBalanceLock)`

SetLock sets Lock field to given value.

### HasLock

`func (o *JettonBalance) HasLock() bool`

HasLock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


