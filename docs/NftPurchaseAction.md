# NftPurchaseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuctionType** | **string** |  | 
**Amount** | [**Price**](Price.md) |  | 
**Nft** | [**NftItem**](NftItem.md) |  | 
**Seller** | [**AccountAddress**](AccountAddress.md) |  | 
**Buyer** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewNftPurchaseAction

`func NewNftPurchaseAction(auctionType string, amount Price, nft NftItem, seller AccountAddress, buyer AccountAddress, ) *NftPurchaseAction`

NewNftPurchaseAction instantiates a new NftPurchaseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftPurchaseActionWithDefaults

`func NewNftPurchaseActionWithDefaults() *NftPurchaseAction`

NewNftPurchaseActionWithDefaults instantiates a new NftPurchaseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuctionType

`func (o *NftPurchaseAction) GetAuctionType() string`

GetAuctionType returns the AuctionType field if non-nil, zero value otherwise.

### GetAuctionTypeOk

`func (o *NftPurchaseAction) GetAuctionTypeOk() (*string, bool)`

GetAuctionTypeOk returns a tuple with the AuctionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionType

`func (o *NftPurchaseAction) SetAuctionType(v string)`

SetAuctionType sets AuctionType field to given value.


### GetAmount

`func (o *NftPurchaseAction) GetAmount() Price`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NftPurchaseAction) GetAmountOk() (*Price, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NftPurchaseAction) SetAmount(v Price)`

SetAmount sets Amount field to given value.


### GetNft

`func (o *NftPurchaseAction) GetNft() NftItem`

GetNft returns the Nft field if non-nil, zero value otherwise.

### GetNftOk

`func (o *NftPurchaseAction) GetNftOk() (*NftItem, bool)`

GetNftOk returns a tuple with the Nft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNft

`func (o *NftPurchaseAction) SetNft(v NftItem)`

SetNft sets Nft field to given value.


### GetSeller

`func (o *NftPurchaseAction) GetSeller() AccountAddress`

GetSeller returns the Seller field if non-nil, zero value otherwise.

### GetSellerOk

`func (o *NftPurchaseAction) GetSellerOk() (*AccountAddress, bool)`

GetSellerOk returns a tuple with the Seller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeller

`func (o *NftPurchaseAction) SetSeller(v AccountAddress)`

SetSeller sets Seller field to given value.


### GetBuyer

`func (o *NftPurchaseAction) GetBuyer() AccountAddress`

GetBuyer returns the Buyer field if non-nil, zero value otherwise.

### GetBuyerOk

`func (o *NftPurchaseAction) GetBuyerOk() (*AccountAddress, bool)`

GetBuyerOk returns a tuple with the Buyer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyer

`func (o *NftPurchaseAction) SetBuyer(v AccountAddress)`

SetBuyer sets Buyer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


