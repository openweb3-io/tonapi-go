# DecodedRawMessageMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boc** | **string** |  | 
**DecodedOpName** | Pointer to **string** |  | [optional] 
**OpCode** | Pointer to **string** |  | [optional] 
**DecodedBody** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewDecodedRawMessageMessage

`func NewDecodedRawMessageMessage(boc string, ) *DecodedRawMessageMessage`

NewDecodedRawMessageMessage instantiates a new DecodedRawMessageMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecodedRawMessageMessageWithDefaults

`func NewDecodedRawMessageMessageWithDefaults() *DecodedRawMessageMessage`

NewDecodedRawMessageMessageWithDefaults instantiates a new DecodedRawMessageMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoc

`func (o *DecodedRawMessageMessage) GetBoc() string`

GetBoc returns the Boc field if non-nil, zero value otherwise.

### GetBocOk

`func (o *DecodedRawMessageMessage) GetBocOk() (*string, bool)`

GetBocOk returns a tuple with the Boc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoc

`func (o *DecodedRawMessageMessage) SetBoc(v string)`

SetBoc sets Boc field to given value.


### GetDecodedOpName

`func (o *DecodedRawMessageMessage) GetDecodedOpName() string`

GetDecodedOpName returns the DecodedOpName field if non-nil, zero value otherwise.

### GetDecodedOpNameOk

`func (o *DecodedRawMessageMessage) GetDecodedOpNameOk() (*string, bool)`

GetDecodedOpNameOk returns a tuple with the DecodedOpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedOpName

`func (o *DecodedRawMessageMessage) SetDecodedOpName(v string)`

SetDecodedOpName sets DecodedOpName field to given value.

### HasDecodedOpName

`func (o *DecodedRawMessageMessage) HasDecodedOpName() bool`

HasDecodedOpName returns a boolean if a field has been set.

### GetOpCode

`func (o *DecodedRawMessageMessage) GetOpCode() string`

GetOpCode returns the OpCode field if non-nil, zero value otherwise.

### GetOpCodeOk

`func (o *DecodedRawMessageMessage) GetOpCodeOk() (*string, bool)`

GetOpCodeOk returns a tuple with the OpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpCode

`func (o *DecodedRawMessageMessage) SetOpCode(v string)`

SetOpCode sets OpCode field to given value.

### HasOpCode

`func (o *DecodedRawMessageMessage) HasOpCode() bool`

HasOpCode returns a boolean if a field has been set.

### GetDecodedBody

`func (o *DecodedRawMessageMessage) GetDecodedBody() interface{}`

GetDecodedBody returns the DecodedBody field if non-nil, zero value otherwise.

### GetDecodedBodyOk

`func (o *DecodedRawMessageMessage) GetDecodedBodyOk() (*interface{}, bool)`

GetDecodedBodyOk returns a tuple with the DecodedBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecodedBody

`func (o *DecodedRawMessageMessage) SetDecodedBody(v interface{})`

SetDecodedBody sets DecodedBody field to given value.

### HasDecodedBody

`func (o *DecodedRawMessageMessage) HasDecodedBody() bool`

HasDecodedBody returns a boolean if a field has been set.

### SetDecodedBodyNil

`func (o *DecodedRawMessageMessage) SetDecodedBodyNil(b bool)`

 SetDecodedBodyNil sets the value for DecodedBody to be an explicit nil

### UnsetDecodedBody
`func (o *DecodedRawMessageMessage) UnsetDecodedBody()`

UnsetDecodedBody ensures that no value is present for DecodedBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


