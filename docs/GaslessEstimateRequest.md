# GaslessEstimateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletAddress** | **string** |  | 
**WalletPublicKey** | **string** |  | 
**Messages** | [**[]GaslessEstimateRequestMessagesInner**](GaslessEstimateRequestMessagesInner.md) |  | 

## Methods

### NewGaslessEstimateRequest

`func NewGaslessEstimateRequest(walletAddress string, walletPublicKey string, messages []GaslessEstimateRequestMessagesInner, ) *GaslessEstimateRequest`

NewGaslessEstimateRequest instantiates a new GaslessEstimateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaslessEstimateRequestWithDefaults

`func NewGaslessEstimateRequestWithDefaults() *GaslessEstimateRequest`

NewGaslessEstimateRequestWithDefaults instantiates a new GaslessEstimateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletAddress

`func (o *GaslessEstimateRequest) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *GaslessEstimateRequest) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *GaslessEstimateRequest) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetWalletPublicKey

`func (o *GaslessEstimateRequest) GetWalletPublicKey() string`

GetWalletPublicKey returns the WalletPublicKey field if non-nil, zero value otherwise.

### GetWalletPublicKeyOk

`func (o *GaslessEstimateRequest) GetWalletPublicKeyOk() (*string, bool)`

GetWalletPublicKeyOk returns a tuple with the WalletPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletPublicKey

`func (o *GaslessEstimateRequest) SetWalletPublicKey(v string)`

SetWalletPublicKey sets WalletPublicKey field to given value.


### GetMessages

`func (o *GaslessEstimateRequest) GetMessages() []GaslessEstimateRequestMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GaslessEstimateRequest) GetMessagesOk() (*[]GaslessEstimateRequestMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GaslessEstimateRequest) SetMessages(v []GaslessEstimateRequestMessagesInner)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


