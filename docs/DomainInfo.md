# DomainInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExpiringAt** | Pointer to **int64** | date of expiring. optional. not all domain in ton has expiration date | [optional] 
**Item** | Pointer to [**NftItem**](NftItem.md) |  | [optional] 

## Methods

### NewDomainInfo

`func NewDomainInfo(name string, ) *DomainInfo`

NewDomainInfo instantiates a new DomainInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInfoWithDefaults

`func NewDomainInfoWithDefaults() *DomainInfo`

NewDomainInfoWithDefaults instantiates a new DomainInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DomainInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainInfo) SetName(v string)`

SetName sets Name field to given value.


### GetExpiringAt

`func (o *DomainInfo) GetExpiringAt() int64`

GetExpiringAt returns the ExpiringAt field if non-nil, zero value otherwise.

### GetExpiringAtOk

`func (o *DomainInfo) GetExpiringAtOk() (*int64, bool)`

GetExpiringAtOk returns a tuple with the ExpiringAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringAt

`func (o *DomainInfo) SetExpiringAt(v int64)`

SetExpiringAt sets ExpiringAt field to given value.

### HasExpiringAt

`func (o *DomainInfo) HasExpiringAt() bool`

HasExpiringAt returns a boolean if a field has been set.

### GetItem

`func (o *DomainInfo) GetItem() NftItem`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *DomainInfo) GetItemOk() (*NftItem, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *DomainInfo) SetItem(v NftItem)`

SetItem sets Item field to given value.

### HasItem

`func (o *DomainInfo) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


