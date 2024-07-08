# JettonTransferAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Recipient** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**SendersWallet** | **string** |  | 
**RecipientsWallet** | **string** |  | 
**Amount** | **string** | amount in quanta of tokens | 
**Comment** | Pointer to **string** |  | [optional] 
**EncryptedComment** | Pointer to [**EncryptedComment**](EncryptedComment.md) |  | [optional] 
**Refund** | Pointer to [**Refund**](Refund.md) |  | [optional] 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 

## Methods

### NewJettonTransferAction

`func NewJettonTransferAction(sendersWallet string, recipientsWallet string, amount string, jetton JettonPreview, ) *JettonTransferAction`

NewJettonTransferAction instantiates a new JettonTransferAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonTransferActionWithDefaults

`func NewJettonTransferActionWithDefaults() *JettonTransferAction`

NewJettonTransferActionWithDefaults instantiates a new JettonTransferAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *JettonTransferAction) GetSender() AccountAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *JettonTransferAction) GetSenderOk() (*AccountAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *JettonTransferAction) SetSender(v AccountAddress)`

SetSender sets Sender field to given value.

### HasSender

`func (o *JettonTransferAction) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *JettonTransferAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *JettonTransferAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *JettonTransferAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *JettonTransferAction) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetSendersWallet

`func (o *JettonTransferAction) GetSendersWallet() string`

GetSendersWallet returns the SendersWallet field if non-nil, zero value otherwise.

### GetSendersWalletOk

`func (o *JettonTransferAction) GetSendersWalletOk() (*string, bool)`

GetSendersWalletOk returns a tuple with the SendersWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendersWallet

`func (o *JettonTransferAction) SetSendersWallet(v string)`

SetSendersWallet sets SendersWallet field to given value.


### GetRecipientsWallet

`func (o *JettonTransferAction) GetRecipientsWallet() string`

GetRecipientsWallet returns the RecipientsWallet field if non-nil, zero value otherwise.

### GetRecipientsWalletOk

`func (o *JettonTransferAction) GetRecipientsWalletOk() (*string, bool)`

GetRecipientsWalletOk returns a tuple with the RecipientsWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsWallet

`func (o *JettonTransferAction) SetRecipientsWallet(v string)`

SetRecipientsWallet sets RecipientsWallet field to given value.


### GetAmount

`func (o *JettonTransferAction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *JettonTransferAction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *JettonTransferAction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetComment

`func (o *JettonTransferAction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *JettonTransferAction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *JettonTransferAction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *JettonTransferAction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEncryptedComment

`func (o *JettonTransferAction) GetEncryptedComment() EncryptedComment`

GetEncryptedComment returns the EncryptedComment field if non-nil, zero value otherwise.

### GetEncryptedCommentOk

`func (o *JettonTransferAction) GetEncryptedCommentOk() (*EncryptedComment, bool)`

GetEncryptedCommentOk returns a tuple with the EncryptedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedComment

`func (o *JettonTransferAction) SetEncryptedComment(v EncryptedComment)`

SetEncryptedComment sets EncryptedComment field to given value.

### HasEncryptedComment

`func (o *JettonTransferAction) HasEncryptedComment() bool`

HasEncryptedComment returns a boolean if a field has been set.

### GetRefund

`func (o *JettonTransferAction) GetRefund() Refund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *JettonTransferAction) GetRefundOk() (*Refund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *JettonTransferAction) SetRefund(v Refund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *JettonTransferAction) HasRefund() bool`

HasRefund returns a boolean if a field has been set.

### GetJetton

`func (o *JettonTransferAction) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *JettonTransferAction) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *JettonTransferAction) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


