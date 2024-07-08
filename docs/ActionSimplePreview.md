# ActionSimplePreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | **string** |  | 
**ActionImage** | Pointer to **string** | a link to an image for this particular action. | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**ValueImage** | Pointer to **string** | a link to an image that depicts this action&#39;s asset. | [optional] 
**Accounts** | [**[]AccountAddress**](AccountAddress.md) |  | 

## Methods

### NewActionSimplePreview

`func NewActionSimplePreview(name string, description string, accounts []AccountAddress, ) *ActionSimplePreview`

NewActionSimplePreview instantiates a new ActionSimplePreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionSimplePreviewWithDefaults

`func NewActionSimplePreviewWithDefaults() *ActionSimplePreview`

NewActionSimplePreviewWithDefaults instantiates a new ActionSimplePreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ActionSimplePreview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionSimplePreview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionSimplePreview) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ActionSimplePreview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ActionSimplePreview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ActionSimplePreview) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetActionImage

`func (o *ActionSimplePreview) GetActionImage() string`

GetActionImage returns the ActionImage field if non-nil, zero value otherwise.

### GetActionImageOk

`func (o *ActionSimplePreview) GetActionImageOk() (*string, bool)`

GetActionImageOk returns a tuple with the ActionImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionImage

`func (o *ActionSimplePreview) SetActionImage(v string)`

SetActionImage sets ActionImage field to given value.

### HasActionImage

`func (o *ActionSimplePreview) HasActionImage() bool`

HasActionImage returns a boolean if a field has been set.

### GetValue

`func (o *ActionSimplePreview) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActionSimplePreview) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActionSimplePreview) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActionSimplePreview) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueImage

`func (o *ActionSimplePreview) GetValueImage() string`

GetValueImage returns the ValueImage field if non-nil, zero value otherwise.

### GetValueImageOk

`func (o *ActionSimplePreview) GetValueImageOk() (*string, bool)`

GetValueImageOk returns a tuple with the ValueImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueImage

`func (o *ActionSimplePreview) SetValueImage(v string)`

SetValueImage sets ValueImage field to given value.

### HasValueImage

`func (o *ActionSimplePreview) HasValueImage() bool`

HasValueImage returns a boolean if a field has been set.

### GetAccounts

`func (o *ActionSimplePreview) GetAccounts() []AccountAddress`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ActionSimplePreview) GetAccountsOk() (*[]AccountAddress, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ActionSimplePreview) SetAccounts(v []AccountAddress)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


