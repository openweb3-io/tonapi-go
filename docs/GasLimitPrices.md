# GasLimitPrices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpecialGasLimit** | Pointer to **int64** |  | [optional] 
**FlatGasLimit** | Pointer to **int64** |  | [optional] 
**FlatGasPrice** | Pointer to **int64** |  | [optional] 
**GasPrice** | **int64** |  | 
**GasLimit** | **int64** |  | 
**GasCredit** | **int64** |  | 
**BlockGasLimit** | **int64** |  | 
**FreezeDueLimit** | **int64** |  | 
**DeleteDueLimit** | **int64** |  | 

## Methods

### NewGasLimitPrices

`func NewGasLimitPrices(gasPrice int64, gasLimit int64, gasCredit int64, blockGasLimit int64, freezeDueLimit int64, deleteDueLimit int64, ) *GasLimitPrices`

NewGasLimitPrices instantiates a new GasLimitPrices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGasLimitPricesWithDefaults

`func NewGasLimitPricesWithDefaults() *GasLimitPrices`

NewGasLimitPricesWithDefaults instantiates a new GasLimitPrices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpecialGasLimit

`func (o *GasLimitPrices) GetSpecialGasLimit() int64`

GetSpecialGasLimit returns the SpecialGasLimit field if non-nil, zero value otherwise.

### GetSpecialGasLimitOk

`func (o *GasLimitPrices) GetSpecialGasLimitOk() (*int64, bool)`

GetSpecialGasLimitOk returns a tuple with the SpecialGasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialGasLimit

`func (o *GasLimitPrices) SetSpecialGasLimit(v int64)`

SetSpecialGasLimit sets SpecialGasLimit field to given value.

### HasSpecialGasLimit

`func (o *GasLimitPrices) HasSpecialGasLimit() bool`

HasSpecialGasLimit returns a boolean if a field has been set.

### GetFlatGasLimit

`func (o *GasLimitPrices) GetFlatGasLimit() int64`

GetFlatGasLimit returns the FlatGasLimit field if non-nil, zero value otherwise.

### GetFlatGasLimitOk

`func (o *GasLimitPrices) GetFlatGasLimitOk() (*int64, bool)`

GetFlatGasLimitOk returns a tuple with the FlatGasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatGasLimit

`func (o *GasLimitPrices) SetFlatGasLimit(v int64)`

SetFlatGasLimit sets FlatGasLimit field to given value.

### HasFlatGasLimit

`func (o *GasLimitPrices) HasFlatGasLimit() bool`

HasFlatGasLimit returns a boolean if a field has been set.

### GetFlatGasPrice

`func (o *GasLimitPrices) GetFlatGasPrice() int64`

GetFlatGasPrice returns the FlatGasPrice field if non-nil, zero value otherwise.

### GetFlatGasPriceOk

`func (o *GasLimitPrices) GetFlatGasPriceOk() (*int64, bool)`

GetFlatGasPriceOk returns a tuple with the FlatGasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlatGasPrice

`func (o *GasLimitPrices) SetFlatGasPrice(v int64)`

SetFlatGasPrice sets FlatGasPrice field to given value.

### HasFlatGasPrice

`func (o *GasLimitPrices) HasFlatGasPrice() bool`

HasFlatGasPrice returns a boolean if a field has been set.

### GetGasPrice

`func (o *GasLimitPrices) GetGasPrice() int64`

GetGasPrice returns the GasPrice field if non-nil, zero value otherwise.

### GetGasPriceOk

`func (o *GasLimitPrices) GetGasPriceOk() (*int64, bool)`

GetGasPriceOk returns a tuple with the GasPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasPrice

`func (o *GasLimitPrices) SetGasPrice(v int64)`

SetGasPrice sets GasPrice field to given value.


### GetGasLimit

`func (o *GasLimitPrices) GetGasLimit() int64`

GetGasLimit returns the GasLimit field if non-nil, zero value otherwise.

### GetGasLimitOk

`func (o *GasLimitPrices) GetGasLimitOk() (*int64, bool)`

GetGasLimitOk returns a tuple with the GasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasLimit

`func (o *GasLimitPrices) SetGasLimit(v int64)`

SetGasLimit sets GasLimit field to given value.


### GetGasCredit

`func (o *GasLimitPrices) GetGasCredit() int64`

GetGasCredit returns the GasCredit field if non-nil, zero value otherwise.

### GetGasCreditOk

`func (o *GasLimitPrices) GetGasCreditOk() (*int64, bool)`

GetGasCreditOk returns a tuple with the GasCredit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGasCredit

`func (o *GasLimitPrices) SetGasCredit(v int64)`

SetGasCredit sets GasCredit field to given value.


### GetBlockGasLimit

`func (o *GasLimitPrices) GetBlockGasLimit() int64`

GetBlockGasLimit returns the BlockGasLimit field if non-nil, zero value otherwise.

### GetBlockGasLimitOk

`func (o *GasLimitPrices) GetBlockGasLimitOk() (*int64, bool)`

GetBlockGasLimitOk returns a tuple with the BlockGasLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockGasLimit

`func (o *GasLimitPrices) SetBlockGasLimit(v int64)`

SetBlockGasLimit sets BlockGasLimit field to given value.


### GetFreezeDueLimit

`func (o *GasLimitPrices) GetFreezeDueLimit() int64`

GetFreezeDueLimit returns the FreezeDueLimit field if non-nil, zero value otherwise.

### GetFreezeDueLimitOk

`func (o *GasLimitPrices) GetFreezeDueLimitOk() (*int64, bool)`

GetFreezeDueLimitOk returns a tuple with the FreezeDueLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreezeDueLimit

`func (o *GasLimitPrices) SetFreezeDueLimit(v int64)`

SetFreezeDueLimit sets FreezeDueLimit field to given value.


### GetDeleteDueLimit

`func (o *GasLimitPrices) GetDeleteDueLimit() int64`

GetDeleteDueLimit returns the DeleteDueLimit field if non-nil, zero value otherwise.

### GetDeleteDueLimitOk

`func (o *GasLimitPrices) GetDeleteDueLimitOk() (*int64, bool)`

GetDeleteDueLimitOk returns a tuple with the DeleteDueLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteDueLimit

`func (o *GasLimitPrices) SetDeleteDueLimit(v int64)`

SetDeleteDueLimit sets DeleteDueLimit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


