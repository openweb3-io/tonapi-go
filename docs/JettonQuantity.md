# JettonQuantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quantity** | **string** |  | 
**WalletAddress** | [**AccountAddress**](AccountAddress.md) |  | 
**Jetton** | [**JettonPreview**](JettonPreview.md) |  | 

## Methods

### NewJettonQuantity

`func NewJettonQuantity(quantity string, walletAddress AccountAddress, jetton JettonPreview, ) *JettonQuantity`

NewJettonQuantity instantiates a new JettonQuantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJettonQuantityWithDefaults

`func NewJettonQuantityWithDefaults() *JettonQuantity`

NewJettonQuantityWithDefaults instantiates a new JettonQuantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuantity

`func (o *JettonQuantity) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *JettonQuantity) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *JettonQuantity) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetWalletAddress

`func (o *JettonQuantity) GetWalletAddress() AccountAddress`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *JettonQuantity) GetWalletAddressOk() (*AccountAddress, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *JettonQuantity) SetWalletAddress(v AccountAddress)`

SetWalletAddress sets WalletAddress field to given value.


### GetJetton

`func (o *JettonQuantity) GetJetton() JettonPreview`

GetJetton returns the Jetton field if non-nil, zero value otherwise.

### GetJettonOk

`func (o *JettonQuantity) GetJettonOk() (*JettonPreview, bool)`

GetJettonOk returns a tuple with the Jetton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJetton

`func (o *JettonQuantity) SetJetton(v JettonPreview)`

SetJetton sets Jetton field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


