# DnsExpiringItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiringAt** | **int64** |  | 
**Name** | **string** |  | 
**DnsItem** | Pointer to [**NftItem**](NftItem.md) |  | [optional] 

## Methods

### NewDnsExpiringItemsInner

`func NewDnsExpiringItemsInner(expiringAt int64, name string, ) *DnsExpiringItemsInner`

NewDnsExpiringItemsInner instantiates a new DnsExpiringItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsExpiringItemsInnerWithDefaults

`func NewDnsExpiringItemsInnerWithDefaults() *DnsExpiringItemsInner`

NewDnsExpiringItemsInnerWithDefaults instantiates a new DnsExpiringItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiringAt

`func (o *DnsExpiringItemsInner) GetExpiringAt() int64`

GetExpiringAt returns the ExpiringAt field if non-nil, zero value otherwise.

### GetExpiringAtOk

`func (o *DnsExpiringItemsInner) GetExpiringAtOk() (*int64, bool)`

GetExpiringAtOk returns a tuple with the ExpiringAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiringAt

`func (o *DnsExpiringItemsInner) SetExpiringAt(v int64)`

SetExpiringAt sets ExpiringAt field to given value.


### GetName

`func (o *DnsExpiringItemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DnsExpiringItemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DnsExpiringItemsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDnsItem

`func (o *DnsExpiringItemsInner) GetDnsItem() NftItem`

GetDnsItem returns the DnsItem field if non-nil, zero value otherwise.

### GetDnsItemOk

`func (o *DnsExpiringItemsInner) GetDnsItemOk() (*NftItem, bool)`

GetDnsItemOk returns a tuple with the DnsItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsItem

`func (o *DnsExpiringItemsInner) SetDnsItem(v NftItem)`

SetDnsItem sets DnsItem field to given value.

### HasDnsItem

`func (o *DnsExpiringItemsInner) HasDnsItem() bool`

HasDnsItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


