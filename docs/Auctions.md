# Auctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Auction**](Auction.md) |  | 
**Total** | **int64** |  | 

## Methods

### NewAuctions

`func NewAuctions(data []Auction, total int64, ) *Auctions`

NewAuctions instantiates a new Auctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuctionsWithDefaults

`func NewAuctionsWithDefaults() *Auctions`

NewAuctionsWithDefaults instantiates a new Auctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *Auctions) GetData() []Auction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Auctions) GetDataOk() (*[]Auction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Auctions) SetData(v []Auction)`

SetData sets Data field to given value.


### GetTotal

`func (o *Auctions) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Auctions) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Auctions) SetTotal(v int64)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


