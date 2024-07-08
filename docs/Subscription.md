# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**WalletAddress** | **string** |  | 
**BeneficiaryAddress** | **string** |  | 
**Amount** | **int64** |  | 
**Period** | **int64** |  | 
**StartTime** | **int64** |  | 
**Timeout** | **int64** |  | 
**LastPaymentTime** | **int64** |  | 
**LastRequestTime** | **int64** |  | 
**SubscriptionId** | **int64** |  | 
**FailedAttempts** | **int32** |  | 

## Methods

### NewSubscription

`func NewSubscription(address string, walletAddress string, beneficiaryAddress string, amount int64, period int64, startTime int64, timeout int64, lastPaymentTime int64, lastRequestTime int64, subscriptionId int64, failedAttempts int32, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Subscription) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Subscription) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Subscription) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetWalletAddress

`func (o *Subscription) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *Subscription) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *Subscription) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.


### GetBeneficiaryAddress

`func (o *Subscription) GetBeneficiaryAddress() string`

GetBeneficiaryAddress returns the BeneficiaryAddress field if non-nil, zero value otherwise.

### GetBeneficiaryAddressOk

`func (o *Subscription) GetBeneficiaryAddressOk() (*string, bool)`

GetBeneficiaryAddressOk returns a tuple with the BeneficiaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeneficiaryAddress

`func (o *Subscription) SetBeneficiaryAddress(v string)`

SetBeneficiaryAddress sets BeneficiaryAddress field to given value.


### GetAmount

`func (o *Subscription) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Subscription) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Subscription) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetPeriod

`func (o *Subscription) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *Subscription) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *Subscription) SetPeriod(v int64)`

SetPeriod sets Period field to given value.


### GetStartTime

`func (o *Subscription) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Subscription) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Subscription) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetTimeout

`func (o *Subscription) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Subscription) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Subscription) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.


### GetLastPaymentTime

`func (o *Subscription) GetLastPaymentTime() int64`

GetLastPaymentTime returns the LastPaymentTime field if non-nil, zero value otherwise.

### GetLastPaymentTimeOk

`func (o *Subscription) GetLastPaymentTimeOk() (*int64, bool)`

GetLastPaymentTimeOk returns a tuple with the LastPaymentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPaymentTime

`func (o *Subscription) SetLastPaymentTime(v int64)`

SetLastPaymentTime sets LastPaymentTime field to given value.


### GetLastRequestTime

`func (o *Subscription) GetLastRequestTime() int64`

GetLastRequestTime returns the LastRequestTime field if non-nil, zero value otherwise.

### GetLastRequestTimeOk

`func (o *Subscription) GetLastRequestTimeOk() (*int64, bool)`

GetLastRequestTimeOk returns a tuple with the LastRequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRequestTime

`func (o *Subscription) SetLastRequestTime(v int64)`

SetLastRequestTime sets LastRequestTime field to given value.


### GetSubscriptionId

`func (o *Subscription) GetSubscriptionId() int64`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *Subscription) GetSubscriptionIdOk() (*int64, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *Subscription) SetSubscriptionId(v int64)`

SetSubscriptionId sets SubscriptionId field to given value.


### GetFailedAttempts

`func (o *Subscription) GetFailedAttempts() int32`

GetFailedAttempts returns the FailedAttempts field if non-nil, zero value otherwise.

### GetFailedAttemptsOk

`func (o *Subscription) GetFailedAttemptsOk() (*int32, bool)`

GetFailedAttemptsOk returns a tuple with the FailedAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAttempts

`func (o *Subscription) SetFailedAttempts(v int32)`

SetFailedAttempts sets FailedAttempts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


