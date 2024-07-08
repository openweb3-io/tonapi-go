# TonConnectProofRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Proof** | [**TonConnectProofRequestProof**](TonConnectProofRequestProof.md) |  | 

## Methods

### NewTonConnectProofRequest

`func NewTonConnectProofRequest(address string, proof TonConnectProofRequestProof, ) *TonConnectProofRequest`

NewTonConnectProofRequest instantiates a new TonConnectProofRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTonConnectProofRequestWithDefaults

`func NewTonConnectProofRequestWithDefaults() *TonConnectProofRequest`

NewTonConnectProofRequestWithDefaults instantiates a new TonConnectProofRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *TonConnectProofRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *TonConnectProofRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *TonConnectProofRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetProof

`func (o *TonConnectProofRequest) GetProof() TonConnectProofRequestProof`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *TonConnectProofRequest) GetProofOk() (*TonConnectProofRequestProof, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *TonConnectProofRequest) SetProof(v TonConnectProofRequestProof)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


