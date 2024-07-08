# JettonMintAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | [**AccountAddress**](AccountAddress.md) |  | 
**RecipientsWallet** | **string** |  | 
**Amount** | **string** | amount in quanta of tokens | 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 

## Methods

### NewJettonMintAction

`func NewJettonMintAction(recipient AccountAddress, recipientsWallet string, amount string, jetton JettonPreview, ) *JettonMintAction`

NewJettonMintAction instantiates a new JettonMintAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonMintActionWithDefaults

`func NewJettonMintActionWithDefaults() *JettonMintAction`

NewJettonMintActionWithDefaults instantiates a new JettonMintAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *JettonMintAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *JettonMintAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *JettonMintAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.


### GetRecipientsWallet

`func (o *JettonMintAction) GetRecipientsWallet() string`

GetRecipientsWallet returns the RecipientsWallet field if non-nil, zero value otherwise.

### GetRecipientsWalletOk

`func (o *JettonMintAction) GetRecipientsWalletOk() (*string, bool)`

GetRecipientsWalletOk returns a tuple with the RecipientsWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsWallet

`func (o *JettonMintAction) SetRecipientsWallet(v string)`

SetRecipientsWallet sets RecipientsWallet field to given value.


### GetAmount

`func (o *JettonMintAction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JettonMintAction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JettonMintAction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetJetton

`func (o *JettonMintAction) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *JettonMintAction) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *JettonMintAction) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


