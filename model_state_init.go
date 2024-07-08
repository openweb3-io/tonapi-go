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

// checks if the StateInit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StateInit{}

// StateInit struct for StateInit
type StateInit struct {
	Boc string `json:"boc"`
}

type _StateInit StateInit

// NewStateInit instantiates a new StateInit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateInit(boc string) *StateInit {
	this := StateInit{}
	this.Boc = boc
	return &this
}

// NewStateInitWithDefaults instantiates a new StateInit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateInitWithDefaults() *StateInit {
	this := StateInit{}
	return &this
}

// GetBoc returns the Boc field value
func (o *StateInit) GetBoc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Boc
}

// GetBocOk returns a tuple with the Boc field value
// and a boolean to check if the value has been set.
func (o *StateInit) GetBocOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Boc, true
}

// SetBoc sets field value
func (o *StateInit) SetBoc(v string) {
	o.Boc = v
}

func (o StateInit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StateInit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boc"] = o.Boc
	return toSerialize, nil
}

func (o *StateInit) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"boc",
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

	varStateInit := _StateInit{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStateInit)

	if err != nil {
		return err
	}

	*o = StateInit(varStateInit)

	return err
}

type NullableStateInit struct {
	value *StateInit
	isSet bool
}

func (v NullableStateInit) Get() *StateInit {
	return v.value
}

func (v *NullableStateInit) Set(val *StateInit) {
	v.value = val
	v.isSet = true
}

func (v NullableStateInit) IsSet() bool {
	return v.isSet
}

func (v *NullableStateInit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateInit(val *StateInit) *NullableStateInit {
	return &NullableStateInit{value: val, isSet: true}
}

func (v NullableStateInit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateInit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

