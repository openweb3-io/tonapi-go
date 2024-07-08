/*
REST api to TON blockchain explorer

Provide access to indexed TON blockchain

API version: 2.0.0
Contact: support@tonkeeper.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tonapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetChartRates200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChartRates200Response{}

// GetChartRates200Response struct for GetChartRates200Response
type GetChartRates200Response struct {
	Points map[string]interface{} `json:"points"`
}

type _GetChartRates200Response GetChartRates200Response

// NewGetChartRates200Response instantiates a new GetChartRates200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChartRates200Response(points map[string]interface{}) *GetChartRates200Response {
	this := GetChartRates200Response{}
	this.Points = points
	return &this
}

// NewGetChartRates200ResponseWithDefaults instantiates a new GetChartRates200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChartRates200ResponseWithDefaults() *GetChartRates200Response {
	this := GetChartRates200Response{}
	return &this
}

// GetPoints returns the Points field value
func (o *GetChartRates200Response) GetPoints() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *GetChartRates200Response) GetPointsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Points, true
}

// SetPoints sets field value
func (o *GetChartRates200Response) SetPoints(v map[string]interface{}) {
	o.Points = v
}

func (o GetChartRates200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChartRates200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["points"] = o.Points
	return toSerialize, nil
}

func (o *GetChartRates200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"points",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetChartRates200Response := _GetChartRates200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetChartRates200Response)

	if err != nil {
		return err
	}

	*o = GetChartRates200Response(varGetChartRates200Response)

	return err
}

type NullableGetChartRates200Response struct {
	value *GetChartRates200Response
	isSet bool
}

func (v NullableGetChartRates200Response) Get() *GetChartRates200Response {
	return v.value
}

func (v *NullableGetChartRates200Response) Set(val *GetChartRates200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChartRates200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChartRates200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChartRates200Response(val *GetChartRates200Response) *NullableGetChartRates200Response {
	return &NullableGetChartRates200Response{value: val, isSet: true}
}

func (v NullableGetChartRates200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChartRates200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

