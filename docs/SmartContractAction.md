# SmartContractAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executor** | [**AccountAddress**](AccountAddress.md) |  | 
**Contract** | [**AccountAddress**](AccountAddress.md) |  | 
**TonAttached** | **int64** | amount in nanotons | 
**Operation** | **string** |  | 
**Payload** | Pointer to **string** |  | [optional] 
**Refund** | Pointer to [**Refund**](Refund.md) |  | [optional] 

## Methods

### NewSmartContractAction

`func NewSmartContractAction(executor AccountAddress, contract AccountAddress, tonAttached int64, operation string, ) *SmartContractAction`

NewSmartContractAction instantiates a new SmartContractAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSmartContractActionWithDefaults

`func NewSmartContractActionWithDefaults() *SmartContractAction`

NewSmartContractActionWithDefaults instantiates a new SmartContractAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutor

`func (o *SmartContractAction) GetExecutor() AccountAddress`

GetExecutor returns the Executor field if non-nil, zero value otherwise.

### GetExecutorOk

`func (o *SmartContractAction) GetExecutorOk() (*AccountAddress, bool)`

GetExecutorOk returns a tuple with the Executor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutor

`func (o *SmartContractAction) SetExecutor(v AccountAddress)`

SetExecutor sets Executor field to given value.


### GetContract

`func (o *SmartContractAction) GetContract() AccountAddress`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *SmartContractAction) GetContractOk() (*AccountAddress, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *SmartContractAction) SetContract(v AccountAddress)`

SetContract sets Contract field to given value.


### GetTonAttached

`func (o *SmartContractAction) GetTonAttached() int64`

GetTonAttached returns the TonAttached field if non-nil, zero value otherwise.

### GetTonAttachedOk

`func (o *SmartContractAction) GetTonAttachedOk() (*int64, bool)`

GetTonAttachedOk returns a tuple with the TonAttached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTonAttached

`func (o *SmartContractAction) SetTonAttached(v int64)`

SetTonAttached sets TonAttached field to given value.


### GetOperation

`func (o *SmartContractAction) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SmartContractAction) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SmartContractAction) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetPayload

`func (o *SmartContractAction) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SmartContractAction) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SmartContractAction) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SmartContractAction) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetRefund

`func (o *SmartContractAction) GetRefund() Refund`

GetRefund returns the Refund field if non-nil, zero value otherwise.

### GetRefundOk

`func (o *SmartContractAction) GetRefundOk() (*Refund, bool)`

GetRefundOk returns a tuple with the Refund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefund

`func (o *SmartContractAction) SetRefund(v Refund)`

SetRefund sets Refund field to given value.

### HasRefund

`func (o *SmartContractAction) HasRefund() bool`

HasRefund returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


