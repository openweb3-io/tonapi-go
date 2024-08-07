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

// checks if the PoolImplementation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolImplementation{}

// PoolImplementation struct for PoolImplementation
type PoolImplementation struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Url string `json:"url"`
	Socials []string `json:"socials"`
}

type _PoolImplementation PoolImplementation

// NewPoolImplementation instantiates a new PoolImplementation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolImplementation(name string, description string, url string, socials []string) *PoolImplementation {
	this := PoolImplementation{}
	this.Name = name
	this.Description = description
	this.Url = url
	this.Socials = socials
	return &this
}

// NewPoolImplementationWithDefaults instantiates a new PoolImplementation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolImplementationWithDefaults() *PoolImplementation {
	this := PoolImplementation{}
	return &this
}

// GetName returns the Name field value
func (o *PoolImplementation) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PoolImplementation) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PoolImplementation) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *PoolImplementation) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PoolImplementation) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PoolImplementation) SetDescription(v string) {
	o.Description = v
}

// GetUrl returns the Url field value
func (o *PoolImplementation) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PoolImplementation) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PoolImplementation) SetUrl(v string) {
	o.Url = v
}

// GetSocials returns the Socials field value
func (o *PoolImplementation) GetSocials() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Socials
}

// GetSocialsOk returns a tuple with the Socials field value
// and a boolean to check if the value has been set.
func (o *PoolImplementation) GetSocialsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Socials, true
}

// SetSocials sets field value
func (o *PoolImplementation) SetSocials(v []string) {
	o.Socials = v
}

func (o PoolImplementation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolImplementation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["url"] = o.Url
	toSerialize["socials"] = o.Socials
	return toSerialize, nil
}

func (o *PoolImplementation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"description",
		"url",
		"socials",
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

	varPoolImplementation := _PoolImplementation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolImplementation)

	if err != nil {
		return err
	}

	*o = PoolImplementation(varPoolImplementation)

	return err
}

type NullablePoolImplementation struct {
	value *PoolImplementation
	isSet bool
}

func (v NullablePoolImplementation) Get() *PoolImplementation {
	return v.value
}

func (v *NullablePoolImplementation) Set(val *PoolImplementation) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolImplementation) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolImplementation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolImplementation(val *PoolImplementation) *NullablePoolImplementation {
	return &NullablePoolImplementation{value: val, isSet: true}
}

func (v NullablePoolImplementation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolImplementation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


