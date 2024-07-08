# DecodedMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**AccountAddress**](AccountAddress.md) |  | 
**DestinationWalletVersion** | **string** |  | 
**ExtInMsgDecoded** | Pointer to [**DecodedMessageExtInMsgDecoded**](DecodedMessageExtInMsgDecoded.md) |  | [optional] 

## Methods

### NewDecodedMessage

`func NewDecodedMessage(destination AccountAddress, destinationWalletVersion string, ) *DecodedMessage`

NewDecodedMessage instantiates a new DecodedMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodedMessageWithDefaults

`func NewDecodedMessageWithDefaults() *DecodedMessage`

NewDecodedMessageWithDefaults instantiates a new DecodedMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *DecodedMessage) GetDestination() AccountAddress`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *DecodedMessage) GetDestinationOk() (*AccountAddress, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *DecodedMessage) SetDestination(v AccountAddress)`

SetDestination sets Destination field to given value.


### GetDestinationWalletVersion

`func (o *DecodedMessage) GetDestinationWalletVersion() string`

GetDestinationWalletVersion returns the DestinationWalletVersion field if non-nil, zero value otherwise.

### GetDestinationWalletVersionOk

`func (o *DecodedMessage) GetDestinationWalletVersionOk() (*string, bool)`

GetDestinationWalletVersionOk returns a tuple with the DestinationWalletVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationWalletVersion

`func (o *DecodedMessage) SetDestinationWalletVersion(v string)`

SetDestinationWalletVersion sets DestinationWalletVersion field to given value.


### GetExtInMsgDecoded

`func (o *DecodedMessage) GetExtInMsgDecoded() DecodedMessageExtInMsgDecoded`

GetExtInMsgDecoded returns the ExtInMsgDecoded field if non-nil, zero value otherwise.

### GetExtInMsgDecodedOk

`func (o *DecodedMessage) GetExtInMsgDecodedOk() (*DecodedMessageExtInMsgDecoded, bool)`

GetExtInMsgDecodedOk returns a tuple with the ExtInMsgDecoded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtInMsgDecoded

`func (o *DecodedMessage) SetExtInMsgDecoded(v DecodedMessageExtInMsgDecoded)`

SetExtInMsgDecoded sets ExtInMsgDecoded field to given value.

### HasExtInMsgDecoded

`func (o *DecodedMessage) HasExtInMsgDecoded() bool`

HasExtInMsgDecoded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


