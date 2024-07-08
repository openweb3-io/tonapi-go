# InscriptionMintAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | [**AccountAddress**](AccountAddress.md) |  | 
**Amount** | **string** | amount in minimal particles | 
**Type** | **string** |  | 
**Ticker** | **string** |  | 
**Decimals** | **int32** |  | 

## Methods

### NewInscriptionMintAction

`func NewInscriptionMintAction(recipient AccountAddress, amount string, type_ string, ticker string, decimals int32, ) *InscriptionMintAction`

NewInscriptionMintAction instantiates a new InscriptionMintAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInscriptionMintActionWithDefaults

`func NewInscriptionMintActionWithDefaults() *InscriptionMintAction`

NewInscriptionMintActionWithDefaults instantiates a new InscriptionMintAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipient

`func (o *InscriptionMintAction) GetRecipient() AccountAddress`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *InscriptionMintAction) GetRecipientOk() (*AccountAddress, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *InscriptionMintAction) SetRecipient(v AccountAddress)`

SetRecipient sets Recipient field to given value.


### GetAmount

`func (o *InscriptionMintAction) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InscriptionMintAction) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InscriptionMintAction) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetType

`func (o *InscriptionMintAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InscriptionMintAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InscriptionMintAction) SetType(v string)`

SetType sets Type field to given value.


### GetTicker

`func (o *InscriptionMintAction) GetTicker() string`

GetTicker returns the Ticker field if non-nil, zero value otherwise.

### GetTickerOk

`func (o *InscriptionMintAction) GetTickerOk() (*string, bool)`

GetTickerOk returns a tuple with the Ticker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicker

`func (o *InscriptionMintAction) SetTicker(v string)`

SetTicker sets Ticker field to given value.


### GetDecimals

`func (o *InscriptionMintAction) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *InscriptionMintAction) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *InscriptionMintAction) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


