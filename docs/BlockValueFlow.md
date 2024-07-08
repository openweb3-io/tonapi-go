# BlockValueFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPrevBlk** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**ToNextBlk** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Imported** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Exported** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**FeesCollected** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Burned** | Pointer to [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | [optional] 
**FeesImported** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Recovered** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Created** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 
**Minted** | [**BlockCurrencyCollection**](BlockCurrencyCollection.md) |  | 

## Methods

### NewBlockValueFlow

`func NewBlockValueFlow(fromPrevBlk BlockCurrencyCollection, toNextBlk BlockCurrencyCollection, imported BlockCurrencyCollection, exported BlockCurrencyCollection, feesCollected BlockCurrencyCollection, feesImported BlockCurrencyCollection, recovered BlockCurrencyCollection, created BlockCurrencyCollection, minted BlockCurrencyCollection, ) *BlockValueFlow`

NewBlockValueFlow instantiates a new BlockValueFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockValueFlowWithDefaults

`func NewBlockValueFlowWithDefaults() *BlockValueFlow`

NewBlockValueFlowWithDefaults instantiates a new BlockValueFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPrevBlk

`func (o *BlockValueFlow) GetFromPrevBlk() BlockCurrencyCollection`

GetFromPrevBlk returns the FromPrevBlk field if non-nil, zero value otherwise.

### GetFromPrevBlkOk

`func (o *BlockValueFlow) GetFromPrevBlkOk() (*BlockCurrencyCollection, bool)`

GetFromPrevBlkOk returns a tuple with the FromPrevBlk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPrevBlk

`func (o *BlockValueFlow) SetFromPrevBlk(v BlockCurrencyCollection)`

SetFromPrevBlk sets FromPrevBlk field to given value.


### GetToNextBlk

`func (o *BlockValueFlow) GetToNextBlk() BlockCurrencyCollection`

GetToNextBlk returns the ToNextBlk field if non-nil, zero value otherwise.

### GetToNextBlkOk

`func (o *BlockValueFlow) GetToNextBlkOk() (*BlockCurrencyCollection, bool)`

GetToNextBlkOk returns a tuple with the ToNextBlk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNextBlk

`func (o *BlockValueFlow) SetToNextBlk(v BlockCurrencyCollection)`

SetToNextBlk sets ToNextBlk field to given value.


### GetImported

`func (o *BlockValueFlow) GetImported() BlockCurrencyCollection`

GetImported returns the Imported field if non-nil, zero value otherwise.

### GetImportedOk

`func (o *BlockValueFlow) GetImportedOk() (*BlockCurrencyCollection, bool)`

GetImportedOk returns a tuple with the Imported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImported

`func (o *BlockValueFlow) SetImported(v BlockCurrencyCollection)`

SetImported sets Imported field to given value.


### GetExported

`func (o *BlockValueFlow) GetExported() BlockCurrencyCollection`

GetExported returns the Exported field if non-nil, zero value otherwise.

### GetExportedOk

`func (o *BlockValueFlow) GetExportedOk() (*BlockCurrencyCollection, bool)`

GetExportedOk returns a tuple with the Exported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExported

`func (o *BlockValueFlow) SetExported(v BlockCurrencyCollection)`

SetExported sets Exported field to given value.


### GetFeesCollected

`func (o *BlockValueFlow) GetFeesCollected() BlockCurrencyCollection`

GetFeesCollected returns the FeesCollected field if non-nil, zero value otherwise.

### GetFeesCollectedOk

`func (o *BlockValueFlow) GetFeesCollectedOk() (*BlockCurrencyCollection, bool)`

GetFeesCollectedOk returns a tuple with the FeesCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesCollected

`func (o *BlockValueFlow) SetFeesCollected(v BlockCurrencyCollection)`

SetFeesCollected sets FeesCollected field to given value.


### GetBurned

`func (o *BlockValueFlow) GetBurned() BlockCurrencyCollection`

GetBurned returns the Burned field if non-nil, zero value otherwise.

### GetBurnedOk

`func (o *BlockValueFlow) GetBurnedOk() (*BlockCurrencyCollection, bool)`

GetBurnedOk returns a tuple with the Burned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurned

`func (o *BlockValueFlow) SetBurned(v BlockCurrencyCollection)`

SetBurned sets Burned field to given value.

### HasBurned

`func (o *BlockValueFlow) HasBurned() bool`

HasBurned returns a boolean if a field has been set.

### GetFeesImported

`func (o *BlockValueFlow) GetFeesImported() BlockCurrencyCollection`

GetFeesImported returns the FeesImported field if non-nil, zero value otherwise.

### GetFeesImportedOk

`func (o *BlockValueFlow) GetFeesImportedOk() (*BlockCurrencyCollection, bool)`

GetFeesImportedOk returns a tuple with the FeesImported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesImported

`func (o *BlockValueFlow) SetFeesImported(v BlockCurrencyCollection)`

SetFeesImported sets FeesImported field to given value.


### GetRecovered

`func (o *BlockValueFlow) GetRecovered() BlockCurrencyCollection`

GetRecovered returns the Recovered field if non-nil, zero value otherwise.

### GetRecoveredOk

`func (o *BlockValueFlow) GetRecoveredOk() (*BlockCurrencyCollection, bool)`

GetRecoveredOk returns a tuple with the Recovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovered

`func (o *BlockValueFlow) SetRecovered(v BlockCurrencyCollection)`

SetRecovered sets Recovered field to given value.


### GetCreated

`func (o *BlockValueFlow) GetCreated() BlockCurrencyCollection`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BlockValueFlow) GetCreatedOk() (*BlockCurrencyCollection, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BlockValueFlow) SetCreated(v BlockCurrencyCollection)`

SetCreated sets Created field to given value.


### GetMinted

`func (o *BlockValueFlow) GetMinted() BlockCurrencyCollection`

GetMinted returns the Minted field if non-nil, zero value otherwise.

### GetMintedOk

`func (o *BlockValueFlow) GetMintedOk() (*BlockCurrencyCollection, bool)`

GetMintedOk returns a tuple with the Minted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinted

`func (o *BlockValueFlow) SetMinted(v BlockCurrencyCollection)`

SetMinted sets Minted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


