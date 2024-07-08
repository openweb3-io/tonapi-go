# DomainBid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | [default to false]
**Value** | **int64** |  | 
**TxTime** | **int64** |  | 
**TxHash** | **string** |  | 
**Bidder** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewDomainBid

`func NewDomainBid(success bool, value int64, txTime int64, txHash string, bidder AccountAddress, ) *DomainBid`

NewDomainBid instantiates a new DomainBid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainBidWithDefaults

`func NewDomainBidWithDefaults() *DomainBid`

NewDomainBidWithDefaults instantiates a new DomainBid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DomainBid) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DomainBid) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DomainBid) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetValue

`func (o *DomainBid) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DomainBid) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DomainBid) SetValue(v int64)`

SetValue sets Value field to given value.


### GetTxTime

`func (o *DomainBid) GetTxTime() int64`

GetTxTime returns the TxTime field if non-nil, zero value otherwise.

### GetTxTimeOk

`func (o *DomainBid) GetTxTimeOk() (*int64, bool)`

GetTxTimeOk returns a tuple with the TxTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxTime

`func (o *DomainBid) SetTxTime(v int64)`

SetTxTime sets TxTime field to given value.


### GetTxHash

`func (o *DomainBid) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *DomainBid) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *DomainBid) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.


### GetBidder

`func (o *DomainBid) GetBidder() AccountAddress`

GetBidder returns the Bidder field if non-nil, zero value otherwise.

### GetBidderOk

`func (o *DomainBid) GetBidderOk() (*AccountAddress, bool)`

GetBidderOk returns a tuple with the Bidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidder

`func (o *DomainBid) SetBidder(v AccountAddress)`

SetBidder sets Bidder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


