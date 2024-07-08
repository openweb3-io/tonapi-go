# SubscriptionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriber** | [**AccountAddress**](AccountAddress.md) |  | 
**Subscription** | **string** |  | 
**Beneficiary** | [**AccountAddress**](AccountAddress.md) |  | 
**Amount** | **int64** |  | 
**Initial** | **bool** |  | 

## Methods

### NewSubscriptionAction

`func NewSubscriptionAction(subscriber AccountAddress, subscription string, beneficiary AccountAddress, amount int64, initial bool, ) *SubscriptionAction`

NewSubscriptionAction instantiates a new SubscriptionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionActionWithDefaults

`func NewSubscriptionActionWithDefaults() *SubscriptionAction`

NewSubscriptionActionWithDefaults instantiates a new SubscriptionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriber

`func (o *SubscriptionAction) GetSubscriber() AccountAddress`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *SubscriptionAction) GetSubscriberOk() (*AccountAddress, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *SubscriptionAction) SetSubscriber(v AccountAddress)`

SetSubscriber sets Subscriber field to given value.


### GetSubscription

`func (o *SubscriptionAction) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *SubscriptionAction) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *SubscriptionAction) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetBeneficiary

`func (o *SubscriptionAction) GetBeneficiary() AccountAddress`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *SubscriptionAction) GetBeneficiaryOk() (*AccountAddress, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *SubscriptionAction) SetBeneficiary(v AccountAddress)`

SetBeneficiary sets Beneficiary field to given value.


### GetAmount

`func (o *SubscriptionAction) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionAction) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionAction) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetInitial

`func (o *SubscriptionAction) GetInitial() bool`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *SubscriptionAction) GetInitialOk() (*bool, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *SubscriptionAction) SetInitial(v bool)`

SetInitial sets Initial field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


