# Auction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Owner** | **string** |  | 
**Price** | **int64** |  | 
**Bids** | **int64** |  | 
**Date** | **int64** |  | 

## Methods

### NewAuction

`func NewAuction(domain string, owner string, price int64, bids int64, date int64, ) *Auction`

NewAuction instantiates a new Auction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuctionWithDefaults

`func NewAuctionWithDefaults() *Auction`

NewAuctionWithDefaults instantiates a new Auction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *Auction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Auction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Auction) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetOwner

`func (o *Auction) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Auction) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Auction) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPrice

`func (o *Auction) GetPrice() int64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Auction) GetPriceOk() (*int64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Auction) SetPrice(v int64)`

SetPrice sets Price field to given value.


### GetBids

`func (o *Auction) GetBids() int64`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *Auction) GetBidsOk() (*int64, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *Auction) SetBids(v int64)`

SetBids sets Bids field to given value.


### GetDate

`func (o *Auction) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Auction) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Auction) SetDate(v int64)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


