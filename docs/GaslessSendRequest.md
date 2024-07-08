# GaslessSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletPublicKey** | **string** | hex encoded public key | 
**Boc** | **string** |  | 

## Methods

### NewGaslessSendRequest

`func NewGaslessSendRequest(walletPublicKey string, boc string, ) *GaslessSendRequest`

NewGaslessSendRequest instantiates a new GaslessSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGaslessSendRequestWithDefaults

`func NewGaslessSendRequestWithDefaults() *GaslessSendRequest`

NewGaslessSendRequestWithDefaults instantiates a new GaslessSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletPublicKey

`func (o *GaslessSendRequest) GetWalletPublicKey() string`

GetWalletPublicKey returns the WalletPublicKey field if non-nil, zero value otherwise.

### GetWalletPublicKeyOk

`func (o *GaslessSendRequest) GetWalletPublicKeyOk() (*string, bool)`

GetWalletPublicKeyOk returns a tuple with the WalletPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletPublicKey

`func (o *GaslessSendRequest) SetWalletPublicKey(v string)`

SetWalletPublicKey sets WalletPublicKey field to given value.


### GetBoc

`func (o *GaslessSendRequest) GetBoc() string`

GetBoc returns the Boc field if non-nil, zero value otherwise.

### GetBocOk

`func (o *GaslessSendRequest) GetBocOk() (*string, bool)`

GetBocOk returns a tuple with the Boc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoc

`func (o *GaslessSendRequest) SetBoc(v string)`

SetBoc sets Boc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


