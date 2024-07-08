# BlockchainRawAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Balance** | **int64** |  | 
**ExtraBalance** | Pointer to **map[string]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **string** |  | [optional] 
**LastTransactionLt** | **int64** |  | 
**LastTransactionHash** | Pointer to **string** |  | [optional] 
**FrozenHash** | Pointer to **string** |  | [optional] 
**Status** | [**AccountStatus**](AccountStatus.md) |  | 
**Storage** | [**AccountStorageInfo**](AccountStorageInfo.md) |  | 
**Libraries** | Pointer to [**[]BlockchainRawAccountLibrariesInner**](BlockchainRawAccountLibrariesInner.md) |  | [optional] 

## Methods

### NewBlockchainRawAccount

`func NewBlockchainRawAccount(address string, balance int64, lastTransactionLt int64, status AccountStatus, storage AccountStorageInfo, ) *BlockchainRawAccount`

NewBlockchainRawAccount instantiates a new BlockchainRawAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockchainRawAccountWithDefaults

`func NewBlockchainRawAccountWithDefaults() *BlockchainRawAccount`

NewBlockchainRawAccountWithDefaults instantiates a new BlockchainRawAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *BlockchainRawAccount) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BlockchainRawAccount) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BlockchainRawAccount) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBalance

`func (o *BlockchainRawAccount) GetBalance() int64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *BlockchainRawAccount) GetBalanceOk() (*int64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *BlockchainRawAccount) SetBalance(v int64)`

SetBalance sets Balance field to given value.


### GetExtraBalance

`func (o *BlockchainRawAccount) GetExtraBalance() map[string]string`

GetExtraBalance returns the ExtraBalance field if non-nil, zero value otherwise.

### GetExtraBalanceOk

`func (o *BlockchainRawAccount) GetExtraBalanceOk() (*map[string]string, bool)`

GetExtraBalanceOk returns a tuple with the ExtraBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraBalance

`func (o *BlockchainRawAccount) SetExtraBalance(v map[string]string)`

SetExtraBalance sets ExtraBalance field to given value.

### HasExtraBalance

`func (o *BlockchainRawAccount) HasExtraBalance() bool`

HasExtraBalance returns a boolean if a field has been set.

### GetCode

`func (o *BlockchainRawAccount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BlockchainRawAccount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BlockchainRawAccount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *BlockchainRawAccount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *BlockchainRawAccount) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BlockchainRawAccount) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BlockchainRawAccount) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *BlockchainRawAccount) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLastTransactionLt

`func (o *BlockchainRawAccount) GetLastTransactionLt() int64`

GetLastTransactionLt returns the LastTransactionLt field if non-nil, zero value otherwise.

### GetLastTransactionLtOk

`func (o *BlockchainRawAccount) GetLastTransactionLtOk() (*int64, bool)`

GetLastTransactionLtOk returns a tuple with the LastTransactionLt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionLt

`func (o *BlockchainRawAccount) SetLastTransactionLt(v int64)`

SetLastTransactionLt sets LastTransactionLt field to given value.


### GetLastTransactionHash

`func (o *BlockchainRawAccount) GetLastTransactionHash() string`

GetLastTransactionHash returns the LastTransactionHash field if non-nil, zero value otherwise.

### GetLastTransactionHashOk

`func (o *BlockchainRawAccount) GetLastTransactionHashOk() (*string, bool)`

GetLastTransactionHashOk returns a tuple with the LastTransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTransactionHash

`func (o *BlockchainRawAccount) SetLastTransactionHash(v string)`

SetLastTransactionHash sets LastTransactionHash field to given value.

### HasLastTransactionHash

`func (o *BlockchainRawAccount) HasLastTransactionHash() bool`

HasLastTransactionHash returns a boolean if a field has been set.

### GetFrozenHash

`func (o *BlockchainRawAccount) GetFrozenHash() string`

GetFrozenHash returns the FrozenHash field if non-nil, zero value otherwise.

### GetFrozenHashOk

`func (o *BlockchainRawAccount) GetFrozenHashOk() (*string, bool)`

GetFrozenHashOk returns a tuple with the FrozenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrozenHash

`func (o *BlockchainRawAccount) SetFrozenHash(v string)`

SetFrozenHash sets FrozenHash field to given value.

### HasFrozenHash

`func (o *BlockchainRawAccount) HasFrozenHash() bool`

HasFrozenHash returns a boolean if a field has been set.

### GetStatus

`func (o *BlockchainRawAccount) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BlockchainRawAccount) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BlockchainRawAccount) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.


### GetStorage

`func (o *BlockchainRawAccount) GetStorage() AccountStorageInfo`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *BlockchainRawAccount) GetStorageOk() (*AccountStorageInfo, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *BlockchainRawAccount) SetStorage(v AccountStorageInfo)`

SetStorage sets Storage field to given value.


### GetLibraries

`func (o *BlockchainRawAccount) GetLibraries() []BlockchainRawAccountLibrariesInner`

GetLibraries returns the Libraries field if non-nil, zero value otherwise.

### GetLibrariesOk

`func (o *BlockchainRawAccount) GetLibrariesOk() (*[]BlockchainRawAccountLibrariesInner, bool)`

GetLibrariesOk returns a tuple with the Libraries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraries

`func (o *BlockchainRawAccount) SetLibraries(v []BlockchainRawAccountLibrariesInner)`

SetLibraries sets Libraries field to given value.

### HasLibraries

`func (o *BlockchainRawAccount) HasLibraries() bool`

HasLibraries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


