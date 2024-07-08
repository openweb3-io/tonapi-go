# GetRates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rates** | [**map[string]TokenRates**](TokenRates.md) |  | 

## Methods

### NewGetRates200Response

`func NewGetRates200Response(rates map[string]TokenRates, ) *GetRates200Response`

NewGetRates200Response instantiates a new GetRates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRates200ResponseWithDefaults

`func NewGetRates200ResponseWithDefaults() *GetRates200Response`

NewGetRates200ResponseWithDefaults instantiates a new GetRates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRates

`func (o *GetRates200Response) GetRates() map[string]TokenRates`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *GetRates200Response) GetRatesOk() (*map[string]TokenRates, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *GetRates200Response) SetRates(v map[string]TokenRates)`

SetRates sets Rates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


