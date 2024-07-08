# MultisigOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**OrderSeqno** | **int64** |  | 
**Threshold** | **int32** |  | 
**SentForExecution** | **bool** |  | 
**Signers** | **[]string** |  | 
**ApprovalsNum** | **int32** |  | 
**ExpirationDate** | **int64** |  | 

## Methods

### NewMultisigOrder

`func NewMultisigOrder(address string, orderSeqno int64, threshold int32, sentForExecution bool, signers []string, approvalsNum int32, expirationDate int64, ) *MultisigOrder`

NewMultisigOrder instantiates a new MultisigOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultisigOrderWithDefaults

`func NewMultisigOrderWithDefaults() *MultisigOrder`

NewMultisigOrderWithDefaults instantiates a new MultisigOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MultisigOrder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MultisigOrder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MultisigOrder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetOrderSeqno

`func (o *MultisigOrder) GetOrderSeqno() int64`

GetOrderSeqno returns the OrderSeqno field if non-nil, zero value otherwise.

### GetOrderSeqnoOk

`func (o *MultisigOrder) GetOrderSeqnoOk() (*int64, bool)`

GetOrderSeqnoOk returns a tuple with the OrderSeqno field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderSeqno

`func (o *MultisigOrder) SetOrderSeqno(v int64)`

SetOrderSeqno sets OrderSeqno field to given value.


### GetThreshold

`func (o *MultisigOrder) GetThreshold() int32`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MultisigOrder) GetThresholdOk() (*int32, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MultisigOrder) SetThreshold(v int32)`

SetThreshold sets Threshold field to given value.


### GetSentForExecution

`func (o *MultisigOrder) GetSentForExecution() bool`

GetSentForExecution returns the SentForExecution field if non-nil, zero value otherwise.

### GetSentForExecutionOk

`func (o *MultisigOrder) GetSentForExecutionOk() (*bool, bool)`

GetSentForExecutionOk returns a tuple with the SentForExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentForExecution

`func (o *MultisigOrder) SetSentForExecution(v bool)`

SetSentForExecution sets SentForExecution field to given value.


### GetSigners

`func (o *MultisigOrder) GetSigners() []string`

GetSigners returns the Signers field if non-nil, zero value otherwise.

### GetSignersOk

`func (o *MultisigOrder) GetSignersOk() (*[]string, bool)`

GetSignersOk returns a tuple with the Signers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigners

`func (o *MultisigOrder) SetSigners(v []string)`

SetSigners sets Signers field to given value.


### GetApprovalsNum

`func (o *MultisigOrder) GetApprovalsNum() int32`

GetApprovalsNum returns the ApprovalsNum field if non-nil, zero value otherwise.

### GetApprovalsNumOk

`func (o *MultisigOrder) GetApprovalsNumOk() (*int32, bool)`

GetApprovalsNumOk returns a tuple with the ApprovalsNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalsNum

`func (o *MultisigOrder) SetApprovalsNum(v int32)`

SetApprovalsNum sets ApprovalsNum field to given value.


### GetExpirationDate

`func (o *MultisigOrder) GetExpirationDate() int64`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *MultisigOrder) GetExpirationDateOk() (*int64, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *MultisigOrder) SetExpirationDate(v int64)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


