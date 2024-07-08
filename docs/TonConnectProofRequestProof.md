# TonConnectProofRequestProof

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | **int64** |  | 
**Domain** | [**TonConnectProofRequestProofDomain**](TonConnectProofRequestProofDomain.md) |  | 
**Signature** | **string** |  | 
**Payload** | **string** |  | 
**StateInit** | Pointer to **string** |  | [optional] 

## Methods

### NewTonConnectProofRequestProof

`func NewTonConnectProofRequestProof(timestamp int64, domain TonConnectProofRequestProofDomain, signature string, payload string, ) *TonConnectProofRequestProof`

NewTonConnectProofRequestProof instantiates a new TonConnectProofRequestProof object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTonConnectProofRequestProofWithDefaults

`func NewTonConnectProofRequestProofWithDefaults() *TonConnectProofRequestProof`

NewTonConnectProofRequestProofWithDefaults instantiates a new TonConnectProofRequestProof object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *TonConnectProofRequestProof) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *TonConnectProofRequestProof) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *TonConnectProofRequestProof) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.


### GetDomain

`func (o *TonConnectProofRequestProof) GetDomain() TonConnectProofRequestProofDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TonConnectProofRequestProof) GetDomainOk() (*TonConnectProofRequestProofDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TonConnectProofRequestProof) SetDomain(v TonConnectProofRequestProofDomain)`

SetDomain sets Domain field to given value.


### GetSignature

`func (o *TonConnectProofRequestProof) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *TonConnectProofRequestProof) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *TonConnectProofRequestProof) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetPayload

`func (o *TonConnectProofRequestProof) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *TonConnectProofRequestProof) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *TonConnectProofRequestProof) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetStateInit

`func (o *TonConnectProofRequestProof) GetStateInit() string`

GetStateInit returns the StateInit field if non-nil, zero value otherwise.

### GetStateInitOk

`func (o *TonConnectProofRequestProof) GetStateInitOk() (*string, bool)`

GetStateInitOk returns a tuple with the StateInit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateInit

`func (o *TonConnectProofRequestProof) SetStateInit(v string)`

SetStateInit sets StateInit field to given value.

### HasStateInit

`func (o *TonConnectProofRequestProof) HasStateInit() bool`

HasStateInit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


