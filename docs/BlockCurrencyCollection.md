# BlockCurrencyCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grams** | **int64** |  | 
**Other** | [**[]BlockCurrencyCollectionOtherInner**](BlockCurrencyCollectionOtherInner.md) |  | 

## Methods

### NewBlockCurrencyCollection

`func NewBlockCurrencyCollection(grams int64, other []BlockCurrencyCollectionOtherInner, ) *BlockCurrencyCollection`

NewBlockCurrencyCollection instantiates a new BlockCurrencyCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockCurrencyCollectionWithDefaults

`func NewBlockCurrencyCollectionWithDefaults() *BlockCurrencyCollection`

NewBlockCurrencyCollectionWithDefaults instantiates a new BlockCurrencyCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrams

`func (o *BlockCurrencyCollection) GetGrams() int64`

GetGrams returns the Grams field if non-nil, zero value otherwise.

### GetGramsOk

`func (o *BlockCurrencyCollection) GetGramsOk() (*int64, bool)`

GetGramsOk returns a tuple with the Grams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrams

`func (o *BlockCurrencyCollection) SetGrams(v int64)`

SetGrams sets Grams field to given value.


### GetOther

`func (o *BlockCurrencyCollection) GetOther() []BlockCurrencyCollectionOtherInner`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *BlockCurrencyCollection) GetOtherOk() (*[]BlockCurrencyCollectionOtherInner, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *BlockCurrencyCollection) SetOther(v []BlockCurrencyCollectionOtherInner)`

SetOther sets Other field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


