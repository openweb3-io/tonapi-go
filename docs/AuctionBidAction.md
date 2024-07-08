# AuctionBidAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionType** | **string** |  | 
**Amount** | [**Price**](Price.md) |  | 
**Nft** | Pointer to [**NftItem**](NftItem.md) |  | [optional] 
**Bidder** | [**AccountAddress**](AccountAddress.md) |  | 
**Auction** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewAuctionBidAction

`func NewAuctionBidAction(auctionType string, amount Price, bidder AccountAddress, auction AccountAddress, ) *AuctionBidAction`

NewAuctionBidAction instantiates a new AuctionBidAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuctionBidActionWithDefaults

`func NewAuctionBidActionWithDefaults() *AuctionBidAction`

NewAuctionBidActionWithDefaults instantiates a new AuctionBidAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionType

`func (o *AuctionBidAction) GetAuctionType() string`

GetAuctionType returns the AuctionType field if non-nil, zero value otherwise.

### GetAuctionTypeOk

`func (o *AuctionBidAction) GetAuctionTypeOk() (*string, bool)`

GetAuctionTypeOk returns a tuple with the AuctionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionType

`func (o *AuctionBidAction) SetAuctionType(v string)`

SetAuctionType sets AuctionType field to given value.


### GetAmount

`func (o *AuctionBidAction) GetAmount() Price`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AuctionBidAction) GetAmountOk() (*Price, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AuctionBidAction) SetAmount(v Price)`

SetAmount sets Amount field to given value.


### GetNft

`func (o *AuctionBidAction) GetNft() NftItem`

GetNft returns the Nft field if non-nil, zero value otherwise.

### GetNftOk

`func (o *AuctionBidAction) GetNftOk() (*NftItem, bool)`

GetNftOk returns a tuple with the Nft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNft

`func (o *AuctionBidAction) SetNft(v NftItem)`

SetNft sets Nft field to given value.

### HasNft

`func (o *AuctionBidAction) HasNft() bool`

HasNft returns a boolean if a field has been set.

### GetBidder

`func (o *AuctionBidAction) GetBidder() AccountAddress`

GetBidder returns the Bidder field if non-nil, zero value otherwise.

### GetBidderOk

`func (o *AuctionBidAction) GetBidderOk() (*AccountAddress, bool)`

GetBidderOk returns a tuple with the Bidder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidder

`func (o *AuctionBidAction) SetBidder(v AccountAddress)`

SetBidder sets Bidder field to given value.


### GetAuction

`func (o *AuctionBidAction) GetAuction() AccountAddress`

GetAuction returns the Auction field if non-nil, zero value otherwise.

### GetAuctionOk

`func (o *AuctionBidAction) GetAuctionOk() (*AccountAddress, bool)`

GetAuctionOk returns a tuple with the Auction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuction

`func (o *AuctionBidAction) SetAuction(v AccountAddress)`

SetAuction sets Auction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


