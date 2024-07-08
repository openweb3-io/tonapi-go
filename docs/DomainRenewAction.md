# DomainRenewAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**ContractAddress** | **string** |  | 
**Renewer** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewDomainRenewAction

`func NewDomainRenewAction(domain string, contractAddress string, renewer AccountAddress, ) *DomainRenewAction`

NewDomainRenewAction instantiates a new DomainRenewAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainRenewActionWithDefaults

`func NewDomainRenewActionWithDefaults() *DomainRenewAction`

NewDomainRenewActionWithDefaults instantiates a new DomainRenewAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DomainRenewAction) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DomainRenewAction) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DomainRenewAction) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetContractAddress

`func (o *DomainRenewAction) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *DomainRenewAction) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *DomainRenewAction) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetRenewer

`func (o *DomainRenewAction) GetRenewer() AccountAddress`

GetRenewer returns the Renewer field if non-nil, zero value otherwise.

### GetRenewerOk

`func (o *DomainRenewAction) GetRenewerOk() (*AccountAddress, bool)`

GetRenewerOk returns a tuple with the Renewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewer

`func (o *DomainRenewAction) SetRenewer(v AccountAddress)`

SetRenewer sets Renewer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


