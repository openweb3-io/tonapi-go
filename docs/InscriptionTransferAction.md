# InscriptionTransferAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sender** | [**AccountAddress**](AccountAddress.md) |  | 
**Recipient** | [**AccountAddress**](AccountAddress.md) |  | 
**Amount** | **string** | amount in minimal particles | 
**Comment** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Ticker** | **string** |  | 
**Decimals** | **int32** |  | 

## Methods

### NewInscriptionTransferAction

`func NewInscriptionTransferAction(sender AccountAddress, recipient AccountAddress, amount string, type_ string, ticker string, decimals int32, ) *InscriptionTransferAction`

NewInscriptionTransferAction instantiates a new InscriptionTransferAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionTransferActionWithDefaults

`func NewInscriptionTransferActionWithDefaults() *InscriptionTransferAction`

NewInscriptionTransferActionWithDefaults instantiates a new InscriptionTransferAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSender

`func (o *InscriptionTransferAction) GetSender() AccountAddress`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *InscriptionTransferAction) GetSenderOk() (*AccountAddress, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *InscriptionTransferAction) SetSender(v AccountAddress)`

SetSender sets Sender field to given value.


### GetRecipient

`func (o *InscriptionTransferAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *InscriptionTransferAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *InscriptionTransferAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.


### GetAmount

`func (o *InscriptionTransferAction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InscriptionTransferAction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InscriptionTransferAction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetComment

`func (o *InscriptionTransferAction) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *InscriptionTransferAction) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *InscriptionTransferAction) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *InscriptionTransferAction) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetType

`func (o *InscriptionTransferAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InscriptionTransferAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InscriptionTransferAction) SetType(v string)`

SetType sets Type field to given value.


### GetTicker

`func (o *InscriptionTransferAction) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *InscriptionTransferAction) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *InscriptionTransferAction) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetDecimals

`func (o *InscriptionTransferAction) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *InscriptionTransferAction) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *InscriptionTransferAction) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


