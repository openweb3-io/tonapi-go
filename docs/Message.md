# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MsgType** | **string** |  | 
**CreatedLt** | **int64** |  | 
**IhrDisabled** | **bool** |  | 
**Bounce** | **bool** |  | 
**Bounced** | **bool** |  | 
**Value** | **int64** |  | 
**FwdFee** | **int64** |  | 
**IhrFee** | **int64** |  | 
**Destination** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Source** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**ImportFee** | **int64** |  | 
**CreatedAt** | **int64** |  | 
**OpCode** | Pointer to **string** |  | [optional] 
**Init** | Pointer to [**StateInit**](StateInit.md) |  | [optional] 
**Hash** | **string** |  | 
**RawBody** | Pointer to **string** | hex-encoded BoC with raw message body | [optional] 
**DecodedOpName** | Pointer to **string** |  | [optional] 
**DecodedBody** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewMessage

`func NewMessage(msgType string, createdLt int64, ihrDisabled bool, bounce bool, bounced bool, value int64, fwdFee int64, ihrFee int64, importFee int64, createdAt int64, hash string, ) *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsgType

`func (o *Message) GetMsgType() string`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *Message) GetMsgTypeOk() (*string, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *Message) SetMsgType(v string)`

SetMsgType sets MsgType field to given value.


### GetCreatedLt

`func (o *Message) GetCreatedLt() int64`

GetCreatedLt returns the CreatedLt field if non-nil, zero value otherwise.

### GetCreatedLtOk

`func (o *Message) GetCreatedLtOk() (*int64, bool)`

GetCreatedLtOk returns a tuple with the CreatedLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedLt

`func (o *Message) SetCreatedLt(v int64)`

SetCreatedLt sets CreatedLt field to given value.


### GetIhrDisabled

`func (o *Message) GetIhrDisabled() bool`

GetIhrDisabled returns the IhrDisabled field if non-nil, zero value otherwise.

### GetIhrDisabledOk

`func (o *Message) GetIhrDisabledOk() (*bool, bool)`

GetIhrDisabledOk returns a tuple with the IhrDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIhrDisabled

`func (o *Message) SetIhrDisabled(v bool)`

SetIhrDisabled sets IhrDisabled field to given value.


### GetBounce

`func (o *Message) GetBounce() bool`

GetBounce returns the Bounce field if non-nil, zero value otherwise.

### GetBounceOk

`func (o *Message) GetBounceOk() (*bool, bool)`

GetBounceOk returns a tuple with the Bounce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounce

`func (o *Message) SetBounce(v bool)`

SetBounce sets Bounce field to given value.


### GetBounced

`func (o *Message) GetBounced() bool`

GetBounced returns the Bounced field if non-nil, zero value otherwise.

### GetBouncedOk

`func (o *Message) GetBouncedOk() (*bool, bool)`

GetBouncedOk returns a tuple with the Bounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounced

`func (o *Message) SetBounced(v bool)`

SetBounced sets Bounced field to given value.


### GetValue

`func (o *Message) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Message) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Message) SetValue(v int64)`

SetValue sets Value field to given value.


### GetFwdFee

`func (o *Message) GetFwdFee() int64`

GetFwdFee returns the FwdFee field if non-nil, zero value otherwise.

### GetFwdFeeOk

`func (o *Message) GetFwdFeeOk() (*int64, bool)`

GetFwdFeeOk returns a tuple with the FwdFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFwdFee

`func (o *Message) SetFwdFee(v int64)`

SetFwdFee sets FwdFee field to given value.


### GetIhrFee

`func (o *Message) GetIhrFee() int64`

GetIhrFee returns the IhrFee field if non-nil, zero value otherwise.

### GetIhrFeeOk

`func (o *Message) GetIhrFeeOk() (*int64, bool)`

GetIhrFeeOk returns a tuple with the IhrFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIhrFee

`func (o *Message) SetIhrFee(v int64)`

SetIhrFee sets IhrFee field to given value.


### GetDestination

`func (o *Message) GetDestination() AccountAddress`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Message) GetDestinationOk() (*AccountAddress, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Message) SetDestination(v AccountAddress)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Message) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetSource

`func (o *Message) GetSource() AccountAddress`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Message) GetSourceOk() (*AccountAddress, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Message) SetSource(v AccountAddress)`

SetSource sets Source field to given value.

### HasSource

`func (o *Message) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetImportFee

`func (o *Message) GetImportFee() int64`

GetImportFee returns the ImportFee field if non-nil, zero value otherwise.

### GetImportFeeOk

`func (o *Message) GetImportFeeOk() (*int64, bool)`

GetImportFeeOk returns a tuple with the ImportFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFee

`func (o *Message) SetImportFee(v int64)`

SetImportFee sets ImportFee field to given value.


### GetCreatedAt

`func (o *Message) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Message) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Message) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetOpCode

`func (o *Message) GetOpCode() string`

GetOpCode returns the OpCode field if non-nil, zero value otherwise.

### GetOpCodeOk

`func (o *Message) GetOpCodeOk() (*string, bool)`

GetOpCodeOk returns a tuple with the OpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpCode

`func (o *Message) SetOpCode(v string)`

SetOpCode sets OpCode field to given value.

### HasOpCode

`func (o *Message) HasOpCode() bool`

HasOpCode returns a boolean if a field has been set.

### GetInit

`func (o *Message) GetInit() StateInit`

GetInit returns the Init field if non-nil, zero value otherwise.

### GetInitOk

`func (o *Message) GetInitOk() (*StateInit, bool)`

GetInitOk returns a tuple with the Init field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInit

`func (o *Message) SetInit(v StateInit)`

SetInit sets Init field to given value.

### HasInit

`func (o *Message) HasInit() bool`

HasInit returns a boolean if a field has been set.

### GetHash

`func (o *Message) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Message) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Message) SetHash(v string)`

SetHash sets Hash field to given value.


### GetRawBody

`func (o *Message) GetRawBody() string`

GetRawBody returns the RawBody field if non-nil, zero value otherwise.

### GetRawBodyOk

`func (o *Message) GetRawBodyOk() (*string, bool)`

GetRawBodyOk returns a tuple with the RawBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawBody

`func (o *Message) SetRawBody(v string)`

SetRawBody sets RawBody field to given value.

### HasRawBody

`func (o *Message) HasRawBody() bool`

HasRawBody returns a boolean if a field has been set.

### GetDecodedOpName

`func (o *Message) GetDecodedOpName() string`

GetDecodedOpName returns the DecodedOpName field if non-nil, zero value otherwise.

### GetDecodedOpNameOk

`func (o *Message) GetDecodedOpNameOk() (*string, bool)`

GetDecodedOpNameOk returns a tuple with the DecodedOpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedOpName

`func (o *Message) SetDecodedOpName(v string)`

SetDecodedOpName sets DecodedOpName field to given value.

### HasDecodedOpName

`func (o *Message) HasDecodedOpName() bool`

HasDecodedOpName returns a boolean if a field has been set.

### GetDecodedBody

`func (o *Message) GetDecodedBody() interface{}`

GetDecodedBody returns the DecodedBody field if non-nil, zero value otherwise.

### GetDecodedBodyOk

`func (o *Message) GetDecodedBodyOk() (*interface{}, bool)`

GetDecodedBodyOk returns a tuple with the DecodedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedBody

`func (o *Message) SetDecodedBody(v interface{})`

SetDecodedBody sets DecodedBody field to given value.

### HasDecodedBody

`func (o *Message) HasDecodedBody() bool`

HasDecodedBody returns a boolean if a field has been set.

### SetDecodedBodyNil

`func (o *Message) SetDecodedBodyNil(b bool)`

 SetDecodedBodyNil sets the value for DecodedBody to be an explicit nil

### UnsetDecodedBody
`func (o *Message) UnsetDecodedBody()`

UnsetDecodedBody ensures that no value is present for DecodedBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


