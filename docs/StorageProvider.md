# StorageProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**AcceptNewContracts** | **bool** |  | 
**RatePerMbDay** | **int64** |  | 
**MaxSpan** | **int64** |  | 
**MinimalFileSize** | **int64** |  | 
**MaximalFileSize** | **int64** |  | 

## Methods

### NewStorageProvider

`func NewStorageProvider(address string, acceptNewContracts bool, ratePerMbDay int64, maxSpan int64, minimalFileSize int64, maximalFileSize int64, ) *StorageProvider`

NewStorageProvider instantiates a new StorageProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageProviderWithDefaults

`func NewStorageProviderWithDefaults() *StorageProvider`

NewStorageProviderWithDefaults instantiates a new StorageProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *StorageProvider) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StorageProvider) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StorageProvider) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAcceptNewContracts

`func (o *StorageProvider) GetAcceptNewContracts() bool`

GetAcceptNewContracts returns the AcceptNewContracts field if non-nil, zero value otherwise.

### GetAcceptNewContractsOk

`func (o *StorageProvider) GetAcceptNewContractsOk() (*bool, bool)`

GetAcceptNewContractsOk returns a tuple with the AcceptNewContracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptNewContracts

`func (o *StorageProvider) SetAcceptNewContracts(v bool)`

SetAcceptNewContracts sets AcceptNewContracts field to given value.


### GetRatePerMbDay

`func (o *StorageProvider) GetRatePerMbDay() int64`

GetRatePerMbDay returns the RatePerMbDay field if non-nil, zero value otherwise.

### GetRatePerMbDayOk

`func (o *StorageProvider) GetRatePerMbDayOk() (*int64, bool)`

GetRatePerMbDayOk returns a tuple with the RatePerMbDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatePerMbDay

`func (o *StorageProvider) SetRatePerMbDay(v int64)`

SetRatePerMbDay sets RatePerMbDay field to given value.


### GetMaxSpan

`func (o *StorageProvider) GetMaxSpan() int64`

GetMaxSpan returns the MaxSpan field if non-nil, zero value otherwise.

### GetMaxSpanOk

`func (o *StorageProvider) GetMaxSpanOk() (*int64, bool)`

GetMaxSpanOk returns a tuple with the MaxSpan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpan

`func (o *StorageProvider) SetMaxSpan(v int64)`

SetMaxSpan sets MaxSpan field to given value.


### GetMinimalFileSize

`func (o *StorageProvider) GetMinimalFileSize() int64`

GetMinimalFileSize returns the MinimalFileSize field if non-nil, zero value otherwise.

### GetMinimalFileSizeOk

`func (o *StorageProvider) GetMinimalFileSizeOk() (*int64, bool)`

GetMinimalFileSizeOk returns a tuple with the MinimalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimalFileSize

`func (o *StorageProvider) SetMinimalFileSize(v int64)`

SetMinimalFileSize sets MinimalFileSize field to given value.


### GetMaximalFileSize

`func (o *StorageProvider) GetMaximalFileSize() int64`

GetMaximalFileSize returns the MaximalFileSize field if non-nil, zero value otherwise.

### GetMaximalFileSizeOk

`func (o *StorageProvider) GetMaximalFileSizeOk() (*int64, bool)`

GetMaximalFileSizeOk returns a tuple with the MaximalFileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximalFileSize

`func (o *StorageProvider) SetMaximalFileSize(v int64)`

SetMaximalFileSize sets MaximalFileSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


