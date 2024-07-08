# DnsRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallet** | Pointer to [**WalletDNS**](WalletDNS.md) |  | [optional] 
**NextResolver** | Pointer to **string** |  | [optional] 
**Sites** | **[]string** |  | 
**Storage** | Pointer to **string** | tonstorage bag id | [optional] 

## Methods

### NewDnsRecord

`func NewDnsRecord(sites []string, ) *DnsRecord`

NewDnsRecord instantiates a new DnsRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRecordWithDefaults

`func NewDnsRecordWithDefaults() *DnsRecord`

NewDnsRecordWithDefaults instantiates a new DnsRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallet

`func (o *DnsRecord) GetWallet() WalletDNS`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *DnsRecord) GetWalletOk() (*WalletDNS, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *DnsRecord) SetWallet(v WalletDNS)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *DnsRecord) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### GetNextResolver

`func (o *DnsRecord) GetNextResolver() string`

GetNextResolver returns the NextResolver field if non-nil, zero value otherwise.

### GetNextResolverOk

`func (o *DnsRecord) GetNextResolverOk() (*string, bool)`

GetNextResolverOk returns a tuple with the NextResolver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextResolver

`func (o *DnsRecord) SetNextResolver(v string)`

SetNextResolver sets NextResolver field to given value.

### HasNextResolver

`func (o *DnsRecord) HasNextResolver() bool`

HasNextResolver returns a boolean if a field has been set.

### GetSites

`func (o *DnsRecord) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *DnsRecord) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *DnsRecord) SetSites(v []string)`

SetSites sets Sites field to given value.


### GetStorage

`func (o *DnsRecord) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *DnsRecord) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *DnsRecord) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *DnsRecord) HasStorage() bool`

HasStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


