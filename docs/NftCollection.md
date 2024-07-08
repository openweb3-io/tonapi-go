# NftCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**NextItemIndex** | **int64** |  | 
**Owner** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**RawCollectionContent** | **string** |  | 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Previews** | Pointer to [**[]ImagePreview**](ImagePreview.md) |  | [optional] 
**ApprovedBy** | **[]string** |  | 

## Methods

### NewNftCollection

`func NewNftCollection(address string, nextItemIndex int64, rawCollectionContent string, approvedBy []string, ) *NftCollection`

NewNftCollection instantiates a new NftCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftCollectionWithDefaults

`func NewNftCollectionWithDefaults() *NftCollection`

NewNftCollectionWithDefaults instantiates a new NftCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NftCollection) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NftCollection) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NftCollection) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetNextItemIndex

`func (o *NftCollection) GetNextItemIndex() int64`

GetNextItemIndex returns the NextItemIndex field if non-nil, zero value otherwise.

### GetNextItemIndexOk

`func (o *NftCollection) GetNextItemIndexOk() (*int64, bool)`

GetNextItemIndexOk returns a tuple with the NextItemIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextItemIndex

`func (o *NftCollection) SetNextItemIndex(v int64)`

SetNextItemIndex sets NextItemIndex field to given value.


### GetOwner

`func (o *NftCollection) GetOwner() AccountAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftCollection) GetOwnerOk() (*AccountAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftCollection) SetOwner(v AccountAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NftCollection) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRawCollectionContent

`func (o *NftCollection) GetRawCollectionContent() string`

GetRawCollectionContent returns the RawCollectionContent field if non-nil, zero value otherwise.

### GetRawCollectionContentOk

`func (o *NftCollection) GetRawCollectionContentOk() (*string, bool)`

GetRawCollectionContentOk returns a tuple with the RawCollectionContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawCollectionContent

`func (o *NftCollection) SetRawCollectionContent(v string)`

SetRawCollectionContent sets RawCollectionContent field to given value.


### GetMetadata

`func (o *NftCollection) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NftCollection) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NftCollection) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *NftCollection) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPreviews

`func (o *NftCollection) GetPreviews() []ImagePreview`

GetPreviews returns the Previews field if non-nil, zero value otherwise.

### GetPreviewsOk

`func (o *NftCollection) GetPreviewsOk() (*[]ImagePreview, bool)`

GetPreviewsOk returns a tuple with the Previews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviews

`func (o *NftCollection) SetPreviews(v []ImagePreview)`

SetPreviews sets Previews field to given value.

### HasPreviews

`func (o *NftCollection) HasPreviews() bool`

HasPreviews returns a boolean if a field has been set.

### GetApprovedBy

`func (o *NftCollection) GetApprovedBy() []string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *NftCollection) GetApprovedByOk() (*[]string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *NftCollection) SetApprovedBy(v []string)`

SetApprovedBy sets ApprovedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


