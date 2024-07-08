# SizeLimitsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMsgBits** | **int64** |  | 
**MaxMsgCells** | **int64** |  | 
**MaxLibraryCells** | **int64** |  | 
**MaxVmDataDepth** | **int32** |  | 
**MaxExtMsgSize** | **int64** |  | 
**MaxExtMsgDepth** | **int32** |  | 
**MaxAccStateCells** | Pointer to **int64** |  | [optional] 
**MaxAccStateBits** | Pointer to **int64** |  | [optional] 

## Methods

### NewSizeLimitsConfig

`func NewSizeLimitsConfig(maxMsgBits int64, maxMsgCells int64, maxLibraryCells int64, maxVmDataDepth int32, maxExtMsgSize int64, maxExtMsgDepth int32, ) *SizeLimitsConfig`

NewSizeLimitsConfig instantiates a new SizeLimitsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSizeLimitsConfigWithDefaults

`func NewSizeLimitsConfigWithDefaults() *SizeLimitsConfig`

NewSizeLimitsConfigWithDefaults instantiates a new SizeLimitsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMsgBits

`func (o *SizeLimitsConfig) GetMaxMsgBits() int64`

GetMaxMsgBits returns the MaxMsgBits field if non-nil, zero value otherwise.

### GetMaxMsgBitsOk

`func (o *SizeLimitsConfig) GetMaxMsgBitsOk() (*int64, bool)`

GetMaxMsgBitsOk returns a tuple with the MaxMsgBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgBits

`func (o *SizeLimitsConfig) SetMaxMsgBits(v int64)`

SetMaxMsgBits sets MaxMsgBits field to given value.


### GetMaxMsgCells

`func (o *SizeLimitsConfig) GetMaxMsgCells() int64`

GetMaxMsgCells returns the MaxMsgCells field if non-nil, zero value otherwise.

### GetMaxMsgCellsOk

`func (o *SizeLimitsConfig) GetMaxMsgCellsOk() (*int64, bool)`

GetMaxMsgCellsOk returns a tuple with the MaxMsgCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsgCells

`func (o *SizeLimitsConfig) SetMaxMsgCells(v int64)`

SetMaxMsgCells sets MaxMsgCells field to given value.


### GetMaxLibraryCells

`func (o *SizeLimitsConfig) GetMaxLibraryCells() int64`

GetMaxLibraryCells returns the MaxLibraryCells field if non-nil, zero value otherwise.

### GetMaxLibraryCellsOk

`func (o *SizeLimitsConfig) GetMaxLibraryCellsOk() (*int64, bool)`

GetMaxLibraryCellsOk returns a tuple with the MaxLibraryCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLibraryCells

`func (o *SizeLimitsConfig) SetMaxLibraryCells(v int64)`

SetMaxLibraryCells sets MaxLibraryCells field to given value.


### GetMaxVmDataDepth

`func (o *SizeLimitsConfig) GetMaxVmDataDepth() int32`

GetMaxVmDataDepth returns the MaxVmDataDepth field if non-nil, zero value otherwise.

### GetMaxVmDataDepthOk

`func (o *SizeLimitsConfig) GetMaxVmDataDepthOk() (*int32, bool)`

GetMaxVmDataDepthOk returns a tuple with the MaxVmDataDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVmDataDepth

`func (o *SizeLimitsConfig) SetMaxVmDataDepth(v int32)`

SetMaxVmDataDepth sets MaxVmDataDepth field to given value.


### GetMaxExtMsgSize

`func (o *SizeLimitsConfig) GetMaxExtMsgSize() int64`

GetMaxExtMsgSize returns the MaxExtMsgSize field if non-nil, zero value otherwise.

### GetMaxExtMsgSizeOk

`func (o *SizeLimitsConfig) GetMaxExtMsgSizeOk() (*int64, bool)`

GetMaxExtMsgSizeOk returns a tuple with the MaxExtMsgSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExtMsgSize

`func (o *SizeLimitsConfig) SetMaxExtMsgSize(v int64)`

SetMaxExtMsgSize sets MaxExtMsgSize field to given value.


### GetMaxExtMsgDepth

`func (o *SizeLimitsConfig) GetMaxExtMsgDepth() int32`

GetMaxExtMsgDepth returns the MaxExtMsgDepth field if non-nil, zero value otherwise.

### GetMaxExtMsgDepthOk

`func (o *SizeLimitsConfig) GetMaxExtMsgDepthOk() (*int32, bool)`

GetMaxExtMsgDepthOk returns a tuple with the MaxExtMsgDepth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExtMsgDepth

`func (o *SizeLimitsConfig) SetMaxExtMsgDepth(v int32)`

SetMaxExtMsgDepth sets MaxExtMsgDepth field to given value.


### GetMaxAccStateCells

`func (o *SizeLimitsConfig) GetMaxAccStateCells() int64`

GetMaxAccStateCells returns the MaxAccStateCells field if non-nil, zero value otherwise.

### GetMaxAccStateCellsOk

`func (o *SizeLimitsConfig) GetMaxAccStateCellsOk() (*int64, bool)`

GetMaxAccStateCellsOk returns a tuple with the MaxAccStateCells field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccStateCells

`func (o *SizeLimitsConfig) SetMaxAccStateCells(v int64)`

SetMaxAccStateCells sets MaxAccStateCells field to given value.

### HasMaxAccStateCells

`func (o *SizeLimitsConfig) HasMaxAccStateCells() bool`

HasMaxAccStateCells returns a boolean if a field has been set.

### GetMaxAccStateBits

`func (o *SizeLimitsConfig) GetMaxAccStateBits() int64`

GetMaxAccStateBits returns the MaxAccStateBits field if non-nil, zero value otherwise.

### GetMaxAccStateBitsOk

`func (o *SizeLimitsConfig) GetMaxAccStateBitsOk() (*int64, bool)`

GetMaxAccStateBitsOk returns a tuple with the MaxAccStateBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAccStateBits

`func (o *SizeLimitsConfig) SetMaxAccStateBits(v int64)`

SetMaxAccStateBits sets MaxAccStateBits field to given value.

### HasMaxAccStateBits

`func (o *SizeLimitsConfig) HasMaxAccStateBits() bool`

HasMaxAccStateBits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


