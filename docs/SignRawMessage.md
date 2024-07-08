# SignRawMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Amount** | **string** | Number of nanocoins to send. Decimal string. | 
**Payload** | Pointer to **string** | Raw one-cell BoC encoded in Base64. | [optional] 
**StateInit** | Pointer to **string** | Raw once-cell BoC encoded in Base64. | [optional] 

## Methods

### NewSignRawMessage

`func NewSignRawMessage(address string, amount string, ) *SignRawMessage`

NewSignRawMessage instantiates a new SignRawMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignRawMessageWithDefaults

`func NewSignRawMessageWithDefaults() *SignRawMessage`

NewSignRawMessageWithDefaults instantiates a new SignRawMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SignRawMessage) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SignRawMessage) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SignRawMessage) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAmount

`func (o *SignRawMessage) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SignRawMessage) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SignRawMessage) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetPayload

`func (o *SignRawMessage) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SignRawMessage) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SignRawMessage) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SignRawMessage) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetStateInit

`func (o *SignRawMessage) GetStateInit() string`

GetStateInit returns the StateInit field if non-nil, zero value otherwise.

### GetStateInitOk

`func (o *SignRawMessage) GetStateInitOk() (*string, bool)`

GetStateInitOk returns a tuple with the StateInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInit

`func (o *SignRawMessage) SetStateInit(v string)`

SetStateInit sets StateInit field to given value.

### HasStateInit

`func (o *SignRawMessage) HasStateInit() bool`

HasStateInit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


