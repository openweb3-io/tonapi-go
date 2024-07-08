# ValidatorsSetListInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | **string** |  | 
**Weight** | **int64** |  | 
**AdnlAddr** | Pointer to **string** |  | [optional] 

## Methods

### NewValidatorsSetListInner

`func NewValidatorsSetListInner(publicKey string, weight int64, ) *ValidatorsSetListInner`

NewValidatorsSetListInner instantiates a new ValidatorsSetListInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatorsSetListInnerWithDefaults

`func NewValidatorsSetListInnerWithDefaults() *ValidatorsSetListInner`

NewValidatorsSetListInnerWithDefaults instantiates a new ValidatorsSetListInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *ValidatorsSetListInner) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ValidatorsSetListInner) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ValidatorsSetListInner) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetWeight

`func (o *ValidatorsSetListInner) GetWeight() int64`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ValidatorsSetListInner) GetWeightOk() (*int64, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ValidatorsSetListInner) SetWeight(v int64)`

SetWeight sets Weight field to given value.


### GetAdnlAddr

`func (o *ValidatorsSetListInner) GetAdnlAddr() string`

GetAdnlAddr returns the AdnlAddr field if non-nil, zero value otherwise.

### GetAdnlAddrOk

`func (o *ValidatorsSetListInner) GetAdnlAddrOk() (*string, bool)`

GetAdnlAddrOk returns a tuple with the AdnlAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdnlAddr

`func (o *ValidatorsSetListInner) SetAdnlAddr(v string)`

SetAdnlAddr sets AdnlAddr field to given value.

### HasAdnlAddr

`func (o *ValidatorsSetListInner) HasAdnlAddr() bool`

HasAdnlAddr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


