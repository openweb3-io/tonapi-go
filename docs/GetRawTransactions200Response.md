# GetRawTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | [**[]BlockRaw**](BlockRaw.md) |  | 
**Transactions** | **string** |  | 

## Methods

### NewGetRawTransactions200Response

`func NewGetRawTransactions200Response(ids []BlockRaw, transactions string, ) *GetRawTransactions200Response`

NewGetRawTransactions200Response instantiates a new GetRawTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawTransactions200ResponseWithDefaults

`func NewGetRawTransactions200ResponseWithDefaults() *GetRawTransactions200Response`

NewGetRawTransactions200ResponseWithDefaults instantiates a new GetRawTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *GetRawTransactions200Response) GetIds() []BlockRaw`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GetRawTransactions200Response) GetIdsOk() (*[]BlockRaw, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GetRawTransactions200Response) SetIds(v []BlockRaw)`

SetIds sets Ids field to given value.


### GetTransactions

`func (o *GetRawTransactions200Response) GetTransactions() string`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *GetRawTransactions200Response) GetTransactionsOk() (*string, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *GetRawTransactions200Response) SetTransactions(v string)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


