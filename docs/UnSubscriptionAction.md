# UnSubscriptionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subscriber** | [**AccountAddress**](AccountAddress.md) |  | 
**Subscription** | **string** |  | 
**Beneficiary** | [**AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewUnSubscriptionAction

`func NewUnSubscriptionAction(subscriber AccountAddress, subscription string, beneficiary AccountAddress, ) *UnSubscriptionAction`

NewUnSubscriptionAction instantiates a new UnSubscriptionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnSubscriptionActionWithDefaults

`func NewUnSubscriptionActionWithDefaults() *UnSubscriptionAction`

NewUnSubscriptionActionWithDefaults instantiates a new UnSubscriptionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubscriber

`func (o *UnSubscriptionAction) GetSubscriber() AccountAddress`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *UnSubscriptionAction) GetSubscriberOk() (*AccountAddress, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *UnSubscriptionAction) SetSubscriber(v AccountAddress)`

SetSubscriber sets Subscriber field to given value.


### GetSubscription

`func (o *UnSubscriptionAction) GetSubscription() string`

GetSubscription returns the Subscription field if non-nil, zero value otherwise.

### GetSubscriptionOk

`func (o *UnSubscriptionAction) GetSubscriptionOk() (*string, bool)`

GetSubscriptionOk returns a tuple with the Subscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscription

`func (o *UnSubscriptionAction) SetSubscription(v string)`

SetSubscription sets Subscription field to given value.


### GetBeneficiary

`func (o *UnSubscriptionAction) GetBeneficiary() AccountAddress`

GetBeneficiary returns the Beneficiary field if non-nil, zero value otherwise.

### GetBeneficiaryOk

`func (o *UnSubscriptionAction) GetBeneficiaryOk() (*AccountAddress, bool)`

GetBeneficiaryOk returns a tuple with the Beneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiary

`func (o *UnSubscriptionAction) SetBeneficiary(v AccountAddress)`

SetBeneficiary sets Beneficiary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


