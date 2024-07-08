# Sale

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Market** | [**AccountAddress**](AccountAddress.md) |  | 
**Owner** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Price** | [**Price**](Price.md) |  | 

## Methods

### NewSale

`func NewSale(address string, market AccountAddress, price Price, ) *Sale`

NewSale instantiates a new Sale object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaleWithDefaults

`func NewSaleWithDefaults() *Sale`

NewSaleWithDefaults instantiates a new Sale object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Sale) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Sale) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Sale) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetMarket

`func (o *Sale) GetMarket() AccountAddress`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *Sale) GetMarketOk() (*AccountAddress, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *Sale) SetMarket(v AccountAddress)`

SetMarket sets Market field to given value.


### GetOwner

`func (o *Sale) GetOwner() AccountAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Sale) GetOwnerOk() (*AccountAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Sale) SetOwner(v AccountAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Sale) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *Sale) GetPrice() Price`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Sale) GetPriceOk() (*Price, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Sale) SetPrice(v Price)`

SetPrice sets Price field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


