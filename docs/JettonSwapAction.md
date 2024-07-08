# JettonSwapAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dex** | **string** |  | 
**AmountIn** | **string** |  | 
**AmountOut** | **string** |  | 
**TonIn** | Pointer to **int64** |  | [optional] 
**TonOut** | Pointer to **int64** |  | [optional] 
**UserWallet** | [**AccountAddress**](AccountAddress.md) |  | 
**Router** | [**AccountAddress**](AccountAddress.md) |  | 
**JettonMasterIn** | Pointer to [**JettonPreview**](JettonPreview.md) |  | [optional] 
**JettonMasterOut** | Pointer to [**JettonPreview**](JettonPreview.md) |  | [optional] 

## Methods

### NewJettonSwapAction

`func NewJettonSwapAction(dex string, amountIn string, amountOut string, userWallet AccountAddress, router AccountAddress, ) *JettonSwapAction`

NewJettonSwapAction instantiates a new JettonSwapAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonSwapActionWithDefaults

`func NewJettonSwapActionWithDefaults() *JettonSwapAction`

NewJettonSwapActionWithDefaults instantiates a new JettonSwapAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDex

`func (o *JettonSwapAction) GetDex() string`

GetDex returns the Dex field if non-nil, zero value otherwise.

### GetDexOk

`func (o *JettonSwapAction) GetDexOk() (*string, bool)`

GetDexOk returns a tuple with the Dex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDex

`func (o *JettonSwapAction) SetDex(v string)`

SetDex sets Dex field to given value.


### GetAmountIn

`func (o *JettonSwapAction) GetAmountIn() string`

GetAmountIn returns the AmountIn field if non-nil, zero value otherwise.

### GetAmountInOk

`func (o *JettonSwapAction) GetAmountInOk() (*string, bool)`

GetAmountInOk returns a tuple with the AmountIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIn

`func (o *JettonSwapAction) SetAmountIn(v string)`

SetAmountIn sets AmountIn field to given value.


### GetAmountOut

`func (o *JettonSwapAction) GetAmountOut() string`

GetAmountOut returns the AmountOut field if non-nil, zero value otherwise.

### GetAmountOutOk

`func (o *JettonSwapAction) GetAmountOutOk() (*string, bool)`

GetAmountOutOk returns a tuple with the AmountOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOut

`func (o *JettonSwapAction) SetAmountOut(v string)`

SetAmountOut sets AmountOut field to given value.


### GetTonIn

`func (o *JettonSwapAction) GetTonIn() int64`

GetTonIn returns the TonIn field if non-nil, zero value otherwise.

### GetTonInOk

`func (o *JettonSwapAction) GetTonInOk() (*int64, bool)`

GetTonInOk returns a tuple with the TonIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonIn

`func (o *JettonSwapAction) SetTonIn(v int64)`

SetTonIn sets TonIn field to given value.

### HasTonIn

`func (o *JettonSwapAction) HasTonIn() bool`

HasTonIn returns a boolean if a field has been set.

### GetTonOut

`func (o *JettonSwapAction) GetTonOut() int64`

GetTonOut returns the TonOut field if non-nil, zero value otherwise.

### GetTonOutOk

`func (o *JettonSwapAction) GetTonOutOk() (*int64, bool)`

GetTonOutOk returns a tuple with the TonOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonOut

`func (o *JettonSwapAction) SetTonOut(v int64)`

SetTonOut sets TonOut field to given value.

### HasTonOut

`func (o *JettonSwapAction) HasTonOut() bool`

HasTonOut returns a boolean if a field has been set.

### GetUserWallet

`func (o *JettonSwapAction) GetUserWallet() AccountAddress`

GetUserWallet returns the UserWallet field if non-nil, zero value otherwise.

### GetUserWalletOk

`func (o *JettonSwapAction) GetUserWalletOk() (*AccountAddress, bool)`

GetUserWalletOk returns a tuple with the UserWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWallet

`func (o *JettonSwapAction) SetUserWallet(v AccountAddress)`

SetUserWallet sets UserWallet field to given value.


### GetRouter

`func (o *JettonSwapAction) GetRouter() AccountAddress`

GetRouter returns the Router field if non-nil, zero value otherwise.

### GetRouterOk

`func (o *JettonSwapAction) GetRouterOk() (*AccountAddress, bool)`

GetRouterOk returns a tuple with the Router field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouter

`func (o *JettonSwapAction) SetRouter(v AccountAddress)`

SetRouter sets Router field to given value.


### GetJettonMasterIn

`func (o *JettonSwapAction) GetJettonMasterIn() JettonPreview`

GetJettonMasterIn returns the JettonMasterIn field if non-nil, zero value otherwise.

### GetJettonMasterInOk

`func (o *JettonSwapAction) GetJettonMasterInOk() (*JettonPreview, bool)`

GetJettonMasterInOk returns a tuple with the JettonMasterIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonMasterIn

`func (o *JettonSwapAction) SetJettonMasterIn(v JettonPreview)`

SetJettonMasterIn sets JettonMasterIn field to given value.

### HasJettonMasterIn

`func (o *JettonSwapAction) HasJettonMasterIn() bool`

HasJettonMasterIn returns a boolean if a field has been set.

### GetJettonMasterOut

`func (o *JettonSwapAction) GetJettonMasterOut() JettonPreview`

GetJettonMasterOut returns the JettonMasterOut field if non-nil, zero value otherwise.

### GetJettonMasterOutOk

`func (o *JettonSwapAction) GetJettonMasterOutOk() (*JettonPreview, bool)`

GetJettonMasterOutOk returns a tuple with the JettonMasterOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJettonMasterOut

`func (o *JettonSwapAction) SetJettonMasterOut(v JettonPreview)`

SetJettonMasterOut sets JettonMasterOut field to given value.

### HasJettonMasterOut

`func (o *JettonSwapAction) HasJettonMasterOut() bool`

HasJettonMasterOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


