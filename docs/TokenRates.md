# TokenRates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prices** | Pointer to **map[string]float32** |  | [optional] 
**Diff24h** | Pointer to **map[string]string** |  | [optional] 
**Diff7d** | Pointer to **map[string]string** |  | [optional] 
**Diff30d** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTokenRates

`func NewTokenRates() *TokenRates`

NewTokenRates instantiates a new TokenRates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRatesWithDefaults

`func NewTokenRatesWithDefaults() *TokenRates`

NewTokenRatesWithDefaults instantiates a new TokenRates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrices

`func (o *TokenRates) GetPrices() map[string]float32`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *TokenRates) GetPricesOk() (*map[string]float32, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *TokenRates) SetPrices(v map[string]float32)`

SetPrices sets Prices field to given value.

### HasPrices

`func (o *TokenRates) HasPrices() bool`

HasPrices returns a boolean if a field has been set.

### GetDiff24h

`func (o *TokenRates) GetDiff24h() map[string]string`

GetDiff24h returns the Diff24h field if non-nil, zero value otherwise.

### GetDiff24hOk

`func (o *TokenRates) GetDiff24hOk() (*map[string]string, bool)`

GetDiff24hOk returns a tuple with the Diff24h field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff24h

`func (o *TokenRates) SetDiff24h(v map[string]string)`

SetDiff24h sets Diff24h field to given value.

### HasDiff24h

`func (o *TokenRates) HasDiff24h() bool`

HasDiff24h returns a boolean if a field has been set.

### GetDiff7d

`func (o *TokenRates) GetDiff7d() map[string]string`

GetDiff7d returns the Diff7d field if non-nil, zero value otherwise.

### GetDiff7dOk

`func (o *TokenRates) GetDiff7dOk() (*map[string]string, bool)`

GetDiff7dOk returns a tuple with the Diff7d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff7d

`func (o *TokenRates) SetDiff7d(v map[string]string)`

SetDiff7d sets Diff7d field to given value.

### HasDiff7d

`func (o *TokenRates) HasDiff7d() bool`

HasDiff7d returns a boolean if a field has been set.

### GetDiff30d

`func (o *TokenRates) GetDiff30d() map[string]string`

GetDiff30d returns the Diff30d field if non-nil, zero value otherwise.

### GetDiff30dOk

`func (o *TokenRates) GetDiff30dOk() (*map[string]string, bool)`

GetDiff30dOk returns a tuple with the Diff30d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiff30d

`func (o *TokenRates) SetDiff30d(v map[string]string)`

SetDiff30d sets Diff30d field to given value.

### HasDiff30d

`func (o *TokenRates) HasDiff30d() bool`

HasDiff30d returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


