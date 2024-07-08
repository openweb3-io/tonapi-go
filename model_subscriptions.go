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

// checks if the Subscriptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscriptions{}

// Subscriptions struct for Subscriptions
type Subscriptions struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type _Subscriptions Subscriptions

// NewSubscriptions instantiates a new Subscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptions(subscriptions []Subscription) *Subscriptions {
	this := Subscriptions{}
	this.Subscriptions = subscriptions
	return &this
}

// NewSubscriptionsWithDefaults instantiates a new Subscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionsWithDefaults() *Subscriptions {
	this := Subscriptions{}
	return &this
}

// GetSubscriptions returns the Subscriptions field value
func (o *Subscriptions) GetSubscriptions() []Subscription {
	if o == nil {
		var ret []Subscription
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *Subscriptions) GetSubscriptionsOk() ([]Subscription, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *Subscriptions) SetSubscriptions(v []Subscription) {
	o.Subscriptions = v
}

func (o Subscriptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscriptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptions"] = o.Subscriptions
	return toSerialize, nil
}

func (o *Subscriptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptions",
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

	varSubscriptions := _Subscriptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptions)

	if err != nil {
		return err
	}

	*o = Subscriptions(varSubscriptions)

	return err
}

type NullableSubscriptions struct {
	value *Subscriptions
	isSet bool
}

func (v NullableSubscriptions) Get() *Subscriptions {
	return v.value
}

func (v *NullableSubscriptions) Set(val *Subscriptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptions(val *Subscriptions) *NullableSubscriptions {
	return &NullableSubscriptions{value: val, isSet: true}
}

func (v NullableSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


