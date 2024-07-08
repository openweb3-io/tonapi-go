# SignRawParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayAddress** | **string** |  | 
**Commission** | **string** | Commission for the transaction. In nanocoins. | 
**From** | **string** |  | 
**ValidUntil** | **int64** |  | 
**Messages** | [**[]SignRawMessage**](SignRawMessage.md) |  | 

## Methods

### NewSignRawParams

`func NewSignRawParams(relayAddress string, commission string, from string, validUntil int64, messages []SignRawMessage, ) *SignRawParams`

NewSignRawParams instantiates a new SignRawParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignRawParamsWithDefaults

`func NewSignRawParamsWithDefaults() *SignRawParams`

NewSignRawParamsWithDefaults instantiates a new SignRawParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayAddress

`func (o *SignRawParams) GetRelayAddress() string`

GetRelayAddress returns the RelayAddress field if non-nil, zero value otherwise.

### GetRelayAddressOk

`func (o *SignRawParams) GetRelayAddressOk() (*string, bool)`

GetRelayAddressOk returns a tuple with the RelayAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAddress

`func (o *SignRawParams) SetRelayAddress(v string)`

SetRelayAddress sets RelayAddress field to given value.


### GetCommission

`func (o *SignRawParams) GetCommission() string`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *SignRawParams) GetCommissionOk() (*string, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *SignRawParams) SetCommission(v string)`

SetCommission sets Commission field to given value.


### GetFrom

`func (o *SignRawParams) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *SignRawParams) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *SignRawParams) SetFrom(v string)`

SetFrom sets From field to given value.


### GetValidUntil

`func (o *SignRawParams) GetValidUntil() int64`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *SignRawParams) GetValidUntilOk() (*int64, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *SignRawParams) SetValidUntil(v int64)`

SetValidUntil sets ValidUntil field to given value.


### GetMessages

`func (o *SignRawParams) GetMessages() []SignRawMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *SignRawParams) GetMessagesOk() (*[]SignRawMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *SignRawParams) SetMessages(v []SignRawMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


