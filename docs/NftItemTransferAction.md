# NftItemTransferAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Recipient** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Nft** | **string** |  | 
**Comment** | Pointer to **string** |  | [optional] 
**EncryptedComment** | Pointer to [**EncryptedComment**](EncryptedComment.md) |  | [optional] 
**Payload** | Pointer to **string** | raw hex encoded payload | [optional] 
**Refund** | Pointer to [**Refund**](Refund.md) |  | [optional] 

## Methods

### NewNftItemTransferAction

`func NewNftItemTransferAction(nft string, ) *NftItemTransferAction`

NewNftItemTransferAction instantiates a new NftItemTransferAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftItemTransferActionWithDefaults

`func NewNftItemTransferActionWithDefaults() *NftItemTransferAction`

NewNftItemTransferActionWithDefaults instantiates a new NftItemTransferAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *NftItemTransferAction) GetSender() AccountAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *NftItemTransferAction) GetSenderOk() (*AccountAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *NftItemTransferAction) SetSender(v AccountAddress)`

SetSender sets Sender field to given value.

### HasSender

`func (o *NftItemTransferAction) HasSender() bool`

HasSender returns a boolean if a field has been set.

### GetRecipient

`func (o *NftItemTransferAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *NftItemTransferAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *NftItemTransferAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *NftItemTransferAction) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetNft

`func (o *NftItemTransferAction) GetNft() string`

GetNft returns the Nft field if non-nil, zero value otherwise.

### GetNftOk

`func (o *NftItemTransferAction) GetNftOk() (*string, bool)`

GetNftOk returns a tuple with the Nft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNft

`func (o *NftItemTransferAction) SetNft(v string)`

SetNft sets Nft field to given value.


### GetComment

`func (o *NftItemTransferAction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NftItemTransferAction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NftItemTransferAction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NftItemTransferAction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEncryptedComment

`func (o *NftItemTransferAction) GetEncryptedComment() EncryptedComment`

GetEncryptedComment returns the EncryptedComment field if non-nil, zero value otherwise.

### GetEncryptedCommentOk

`func (o *NftItemTransferAction) GetEncryptedCommentOk() (*EncryptedComment, bool)`

GetEncryptedCommentOk returns a tuple with the EncryptedComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedComment

`func (o *NftItemTransferAction) SetEncryptedComment(v EncryptedComment)`

SetEncryptedComment sets EncryptedComment field to given value.

### HasEncryptedComment

`func (o *NftItemTransferAction) HasEncryptedComment() bool`

HasEncryptedComment returns a boolean if a field has been set.

### GetPayload

`func (o *NftItemTransferAction) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *NftItemTransferAction) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *NftItemTransferAction) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *NftItemTransferAction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRefund

`func (o *NftItemTransferAction) GetRefund() Refund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *NftItemTransferAction) GetRefundOk() (*Refund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *NftItemTransferAction) SetRefund(v Refund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *NftItemTransferAction) HasRefund() bool`

HasRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


