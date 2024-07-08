# TonTransferAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | [**AccountAddress**](AccountAddress.md) |  | 
**Recipient** | [**AccountAddress**](AccountAddress.md) |  | 
**Amount** | **int64** | amount in nanotons | 
**Comment** | Pointer to **string** |  | [optional] 
**EncryptedComment** | Pointer to [**EncryptedComment**](EncryptedComment.md) |  | [optional] 
**Refund** | Pointer to [**Refund**](Refund.md) |  | [optional] 

## Methods

### NewTonTransferAction

`func NewTonTransferAction(sender AccountAddress, recipient AccountAddress, amount int64, ) *TonTransferAction`

NewTonTransferAction instantiates a new TonTransferAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTonTransferActionWithDefaults

`func NewTonTransferActionWithDefaults() *TonTransferAction`

NewTonTransferActionWithDefaults instantiates a new TonTransferAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *TonTransferAction) GetSender() AccountAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TonTransferAction) GetSenderOk() (*AccountAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TonTransferAction) SetSender(v AccountAddress)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *TonTransferAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *TonTransferAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *TonTransferAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.


### GetAmount

`func (o *TonTransferAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TonTransferAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TonTransferAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetComment

`func (o *TonTransferAction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *TonTransferAction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *TonTransferAction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *TonTransferAction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEncryptedComment

`func (o *TonTransferAction) GetEncryptedComment() EncryptedComment`

GetEncryptedComment returns the EncryptedComment field if non-nil, zero value otherwise.

### GetEncryptedCommentOk

`func (o *TonTransferAction) GetEncryptedCommentOk() (*EncryptedComment, bool)`

GetEncryptedCommentOk returns a tuple with the EncryptedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedComment

`func (o *TonTransferAction) SetEncryptedComment(v EncryptedComment)`

SetEncryptedComment sets EncryptedComment field to given value.

### HasEncryptedComment

`func (o *TonTransferAction) HasEncryptedComment() bool`

HasEncryptedComment returns a boolean if a field has been set.

### GetRefund

`func (o *TonTransferAction) GetRefund() Refund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *TonTransferAction) GetRefundOk() (*Refund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *TonTransferAction) SetRefund(v Refund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *TonTransferAction) HasRefund() bool`

HasRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


