# GetRawListBlockTransactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**BlockRaw**](BlockRaw.md) |  | 
**ReqCount** | **int32** |  | 
**Incomplete** | **bool** |  | 
**Ids** | [**[]GetRawListBlockTransactions200ResponseIdsInner**](GetRawListBlockTransactions200ResponseIdsInner.md) |  | 
**Proof** | **string** |  | 

## Methods

### NewGetRawListBlockTransactions200Response

`func NewGetRawListBlockTransactions200Response(id BlockRaw, reqCount int32, incomplete bool, ids []GetRawListBlockTransactions200ResponseIdsInner, proof string, ) *GetRawListBlockTransactions200Response`

NewGetRawListBlockTransactions200Response instantiates a new GetRawListBlockTransactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRawListBlockTransactions200ResponseWithDefaults

`func NewGetRawListBlockTransactions200ResponseWithDefaults() *GetRawListBlockTransactions200Response`

NewGetRawListBlockTransactions200ResponseWithDefaults instantiates a new GetRawListBlockTransactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRawListBlockTransactions200Response) GetId() BlockRaw`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRawListBlockTransactions200Response) GetIdOk() (*BlockRaw, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRawListBlockTransactions200Response) SetId(v BlockRaw)`

SetId sets Id field to given value.


### GetReqCount

`func (o *GetRawListBlockTransactions200Response) GetReqCount() int32`

GetReqCount returns the ReqCount field if non-nil, zero value otherwise.

### GetReqCountOk

`func (o *GetRawListBlockTransactions200Response) GetReqCountOk() (*int32, bool)`

GetReqCountOk returns a tuple with the ReqCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqCount

`func (o *GetRawListBlockTransactions200Response) SetReqCount(v int32)`

SetReqCount sets ReqCount field to given value.


### GetIncomplete

`func (o *GetRawListBlockTransactions200Response) GetIncomplete() bool`

GetIncomplete returns the Incomplete field if non-nil, zero value otherwise.

### GetIncompleteOk

`func (o *GetRawListBlockTransactions200Response) GetIncompleteOk() (*bool, bool)`

GetIncompleteOk returns a tuple with the Incomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomplete

`func (o *GetRawListBlockTransactions200Response) SetIncomplete(v bool)`

SetIncomplete sets Incomplete field to given value.


### GetIds

`func (o *GetRawListBlockTransactions200Response) GetIds() []GetRawListBlockTransactions200ResponseIdsInner`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *GetRawListBlockTransactions200Response) GetIdsOk() (*[]GetRawListBlockTransactions200ResponseIdsInner, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *GetRawListBlockTransactions200Response) SetIds(v []GetRawListBlockTransactions200ResponseIdsInner)`

SetIds sets Ids field to given value.


### GetProof

`func (o *GetRawListBlockTransactions200Response) GetProof() string`

GetProof returns the Proof field if non-nil, zero value otherwise.

### GetProofOk

`func (o *GetRawListBlockTransactions200Response) GetProofOk() (*string, bool)`

GetProofOk returns a tuple with the Proof field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProof

`func (o *GetRawListBlockTransactions200Response) SetProof(v string)`

SetProof sets Proof field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


