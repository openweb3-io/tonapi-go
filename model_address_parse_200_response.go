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

// checks if the AddressParse200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressParse200Response{}

// AddressParse200Response struct for AddressParse200Response
type AddressParse200Response struct {
	RawForm string `json:"raw_form"`
	Bounceable AddressParse200ResponseBounceable `json:"bounceable"`
	NonBounceable AddressParse200ResponseBounceable `json:"non_bounceable"`
	GivenType string `json:"given_type"`
	TestOnly bool `json:"test_only"`
}

type _AddressParse200Response AddressParse200Response

// NewAddressParse200Response instantiates a new AddressParse200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressParse200Response(rawForm string, bounceable AddressParse200ResponseBounceable, nonBounceable AddressParse200ResponseBounceable, givenType string, testOnly bool) *AddressParse200Response {
	this := AddressParse200Response{}
	this.RawForm = rawForm
	this.Bounceable = bounceable
	this.NonBounceable = nonBounceable
	this.GivenType = givenType
	this.TestOnly = testOnly
	return &this
}

// NewAddressParse200ResponseWithDefaults instantiates a new AddressParse200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressParse200ResponseWithDefaults() *AddressParse200Response {
	this := AddressParse200Response{}
	return &this
}

// GetRawForm returns the RawForm field value
func (o *AddressParse200Response) GetRawForm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RawForm
}

// GetRawFormOk returns a tuple with the RawForm field value
// and a boolean to check if the value has been set.
func (o *AddressParse200Response) GetRawFormOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RawForm, true
}

// SetRawForm sets field value
func (o *AddressParse200Response) SetRawForm(v string) {
	o.RawForm = v
}

// GetBounceable returns the Bounceable field value
func (o *AddressParse200Response) GetBounceable() AddressParse200ResponseBounceable {
	if o == nil {
		var ret AddressParse200ResponseBounceable
		return ret
	}

	return o.Bounceable
}

// GetBounceableOk returns a tuple with the Bounceable field value
// and a boolean to check if the value has been set.
func (o *AddressParse200Response) GetBounceableOk() (*AddressParse200ResponseBounceable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bounceable, true
}

// SetBounceable sets field value
func (o *AddressParse200Response) SetBounceable(v AddressParse200ResponseBounceable) {
	o.Bounceable = v
}

// GetNonBounceable returns the NonBounceable field value
func (o *AddressParse200Response) GetNonBounceable() AddressParse200ResponseBounceable {
	if o == nil {
		var ret AddressParse200ResponseBounceable
		return ret
	}

	return o.NonBounceable
}

// GetNonBounceableOk returns a tuple with the NonBounceable field value
// and a boolean to check if the value has been set.
func (o *AddressParse200Response) GetNonBounceableOk() (*AddressParse200ResponseBounceable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NonBounceable, true
}

// SetNonBounceable sets field value
func (o *AddressParse200Response) SetNonBounceable(v AddressParse200ResponseBounceable) {
	o.NonBounceable = v
}

// GetGivenType returns the GivenType field value
func (o *AddressParse200Response) GetGivenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GivenType
}

// GetGivenTypeOk returns a tuple with the GivenType field value
// and a boolean to check if the value has been set.
func (o *AddressParse200Response) GetGivenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GivenType, true
}

// SetGivenType sets field value
func (o *AddressParse200Response) SetGivenType(v string) {
	o.GivenType = v
}

// GetTestOnly returns the TestOnly field value
func (o *AddressParse200Response) GetTestOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TestOnly
}

// GetTestOnlyOk returns a tuple with the TestOnly field value
// and a boolean to check if the value has been set.
func (o *AddressParse200Response) GetTestOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestOnly, true
}

// SetTestOnly sets field value
func (o *AddressParse200Response) SetTestOnly(v bool) {
	o.TestOnly = v
}

func (o AddressParse200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressParse200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["raw_form"] = o.RawForm
	toSerialize["bounceable"] = o.Bounceable
	toSerialize["non_bounceable"] = o.NonBounceable
	toSerialize["given_type"] = o.GivenType
	toSerialize["test_only"] = o.TestOnly
	return toSerialize, nil
}

func (o *AddressParse200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"raw_form",
		"bounceable",
		"non_bounceable",
		"given_type",
		"test_only",
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

	varAddressParse200Response := _AddressParse200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddressParse200Response)

	if err != nil {
		return err
	}

	*o = AddressParse200Response(varAddressParse200Response)

	return err
}

type NullableAddressParse200Response struct {
	value *AddressParse200Response
	isSet bool
}

func (v NullableAddressParse200Response) Get() *AddressParse200Response {
	return v.value
}

func (v *NullableAddressParse200Response) Set(val *AddressParse200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressParse200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressParse200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressParse200Response(val *AddressParse200Response) *NullableAddressParse200Response {
	return &NullableAddressParse200Response{value: val, isSet: true}
}

func (v NullableAddressParse200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressParse200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


