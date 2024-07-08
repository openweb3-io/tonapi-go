# EmulateMessageToWalletRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boc** | **string** |  | 
**Params** | Pointer to [**[]EmulateMessageToWalletRequestParamsInner**](EmulateMessageToWalletRequestParamsInner.md) | additional per account configuration | [optional] 

## Methods

### NewEmulateMessageToWalletRequest

`func NewEmulateMessageToWalletRequest(boc string, ) *EmulateMessageToWalletRequest`

NewEmulateMessageToWalletRequest instantiates a new EmulateMessageToWalletRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmulateMessageToWalletRequestWithDefaults

`func NewEmulateMessageToWalletRequestWithDefaults() *EmulateMessageToWalletRequest`

NewEmulateMessageToWalletRequestWithDefaults instantiates a new EmulateMessageToWalletRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoc

`func (o *EmulateMessageToWalletRequest) GetBoc() string`

GetBoc returns the Boc field if non-nil, zero value otherwise.

### GetBocOk

`func (o *EmulateMessageToWalletRequest) GetBocOk() (*string, bool)`

GetBocOk returns a tuple with the Boc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoc

`func (o *EmulateMessageToWalletRequest) SetBoc(v string)`

SetBoc sets Boc field to given value.


### GetParams

`func (o *EmulateMessageToWalletRequest) GetParams() []EmulateMessageToWalletRequestParamsInner`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *EmulateMessageToWalletRequest) GetParamsOk() (*[]EmulateMessageToWalletRequestParamsInner, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *EmulateMessageToWalletRequest) SetParams(v []EmulateMessageToWalletRequestParamsInner)`

SetParams sets Params field to given value.

### HasParams

`func (o *EmulateMessageToWalletRequest) HasParams() bool`

HasParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


