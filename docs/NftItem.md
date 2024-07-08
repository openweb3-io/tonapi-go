# NftItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**Index** | **int64** |  | 
**Owner** | Pointer to [**AccountAddress**](AccountAddress.md) |  | [optional] 
**Collection** | Pointer to [**NftItemCollection**](NftItemCollection.md) |  | [optional] 
**Verified** | **bool** |  | 
**Metadata** | **map[string]interface{}** |  | 
**Sale** | Pointer to [**Sale**](Sale.md) |  | [optional] 
**Previews** | Pointer to [**[]ImagePreview**](ImagePreview.md) |  | [optional] 
**Dns** | Pointer to **string** |  | [optional] 
**ApprovedBy** | **[]string** |  | 
**IncludeCnft** | Pointer to **bool** |  | [optional] 
**Trust** | [**TrustType**](TrustType.md) |  | 

## Methods

### NewNftItem

`func NewNftItem(address string, index int64, verified bool, metadata map[string]interface{}, approvedBy []string, trust TrustType, ) *NftItem`

NewNftItem instantiates a new NftItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftItemWithDefaults

`func NewNftItemWithDefaults() *NftItem`

NewNftItemWithDefaults instantiates a new NftItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *NftItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NftItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NftItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetIndex

`func (o *NftItem) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *NftItem) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *NftItem) SetIndex(v int64)`

SetIndex sets Index field to given value.


### GetOwner

`func (o *NftItem) GetOwner() AccountAddress`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NftItem) GetOwnerOk() (*AccountAddress, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NftItem) SetOwner(v AccountAddress)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NftItem) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCollection

`func (o *NftItem) GetCollection() NftItemCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *NftItem) GetCollectionOk() (*NftItemCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *NftItem) SetCollection(v NftItemCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *NftItem) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetVerified

`func (o *NftItem) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *NftItem) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *NftItem) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetMetadata

`func (o *NftItem) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *NftItem) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *NftItem) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetSale

`func (o *NftItem) GetSale() Sale`

GetSale returns the Sale field if non-nil, zero value otherwise.

### GetSaleOk

`func (o *NftItem) GetSaleOk() (*Sale, bool)`

GetSaleOk returns a tuple with the Sale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSale

`func (o *NftItem) SetSale(v Sale)`

SetSale sets Sale field to given value.

### HasSale

`func (o *NftItem) HasSale() bool`

HasSale returns a boolean if a field has been set.

### GetPreviews

`func (o *NftItem) GetPreviews() []ImagePreview`

GetPreviews returns the Previews field if non-nil, zero value otherwise.

### GetPreviewsOk

`func (o *NftItem) GetPreviewsOk() (*[]ImagePreview, bool)`

GetPreviewsOk returns a tuple with the Previews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviews

`func (o *NftItem) SetPreviews(v []ImagePreview)`

SetPreviews sets Previews field to given value.

### HasPreviews

`func (o *NftItem) HasPreviews() bool`

HasPreviews returns a boolean if a field has been set.

### GetDns

`func (o *NftItem) GetDns() string`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *NftItem) GetDnsOk() (*string, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *NftItem) SetDns(v string)`

SetDns sets Dns field to given value.

### HasDns

`func (o *NftItem) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetApprovedBy

`func (o *NftItem) GetApprovedBy() []string`

GetApprovedBy returns the ApprovedBy field if non-nil, zero value otherwise.

### GetApprovedByOk

`func (o *NftItem) GetApprovedByOk() (*[]string, bool)`

GetApprovedByOk returns a tuple with the ApprovedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedBy

`func (o *NftItem) SetApprovedBy(v []string)`

SetApprovedBy sets ApprovedBy field to given value.


### GetIncludeCnft

`func (o *NftItem) GetIncludeCnft() bool`

GetIncludeCnft returns the IncludeCnft field if non-nil, zero value otherwise.

### GetIncludeCnftOk

`func (o *NftItem) GetIncludeCnftOk() (*bool, bool)`

GetIncludeCnftOk returns a tuple with the IncludeCnft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCnft

`func (o *NftItem) SetIncludeCnft(v bool)`

SetIncludeCnft sets IncludeCnft field to given value.

### HasIncludeCnft

`func (o *NftItem) HasIncludeCnft() bool`

HasIncludeCnft returns a boolean if a field has been set.

### GetTrust

`func (o *NftItem) GetTrust() TrustType`

GetTrust returns the Trust field if non-nil, zero value otherwise.

### GetTrustOk

`func (o *NftItem) GetTrustOk() (*TrustType, bool)`

GetTrustOk returns a tuple with the Trust field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrust

`func (o *NftItem) SetTrust(v TrustType)`

SetTrust sets Trust field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


