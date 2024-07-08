# JettonBurnAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | [**AccountAddress**](AccountAddress.md) |  | 
**SendersWallet** | **string** |  | 
**Amount** | **string** | amount in quanta of tokens | 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 

## Methods

### NewJettonBurnAction

`func NewJettonBurnAction(sender AccountAddress, sendersWallet string, amount string, jetton JettonPreview, ) *JettonBurnAction`

NewJettonBurnAction instantiates a new JettonBurnAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonBurnActionWithDefaults

`func NewJettonBurnActionWithDefaults() *JettonBurnAction`

NewJettonBurnActionWithDefaults instantiates a new JettonBurnAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *JettonBurnAction) GetSender() AccountAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *JettonBurnAction) GetSenderOk() (*AccountAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *JettonBurnAction) SetSender(v AccountAddress)`

SetSender sets Sender field to given value.


### GetSendersWallet

`func (o *JettonBurnAction) GetSendersWallet() string`

GetSendersWallet returns the SendersWallet field if non-nil, zero value otherwise.

### GetSendersWalletOk

`func (o *JettonBurnAction) GetSendersWalletOk() (*string, bool)`

GetSendersWalletOk returns a tuple with the SendersWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendersWallet

`func (o *JettonBurnAction) SetSendersWallet(v string)`

SetSendersWallet sets SendersWallet field to given value.


### GetAmount

`func (o *JettonBurnAction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JettonBurnAction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JettonBurnAction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetJetton

`func (o *JettonBurnAction) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *JettonBurnAction) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *JettonBurnAction) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


