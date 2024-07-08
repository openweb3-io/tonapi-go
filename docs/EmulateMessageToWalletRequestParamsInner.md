# EmulateMessageToWalletRequestParamsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Balance** | Pointer to **int64** |  | [optional] 

## Methods

### NewEmulateMessageToWalletRequestParamsInner

`func NewEmulateMessageToWalletRequestParamsInner(address string, ) *EmulateMessageToWalletRequestParamsInner`

NewEmulateMessageToWalletRequestParamsInner instantiates a new EmulateMessageToWalletRequestParamsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmulateMessageToWalletRequestParamsInnerWithDefaults

`func NewEmulateMessageToWalletRequestParamsInnerWithDefaults() *EmulateMessageToWalletRequestParamsInner`

NewEmulateMessageToWalletRequestParamsInnerWithDefaults instantiates a new EmulateMessageToWalletRequestParamsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *EmulateMessageToWalletRequestParamsInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *EmulateMessageToWalletRequestParamsInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *EmulateMessageToWalletRequestParamsInner) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *EmulateMessageToWalletRequestParamsInner) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *EmulateMessageToWalletRequestParamsInner) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *EmulateMessageToWalletRequestParamsInner) SetBalance(v int64)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *EmulateMessageToWalletRequestParamsInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


