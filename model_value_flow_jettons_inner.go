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

// checks if the ValueFlowJettonsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueFlowJettonsInner{}

// ValueFlowJettonsInner struct for ValueFlowJettonsInner
type ValueFlowJettonsInner struct {
	Account AccountAddress `json:"account"`
	Jetton JettonPreview `json:"jetton"`
	Quantity int64 `json:"quantity"`
}

type _ValueFlowJettonsInner ValueFlowJettonsInner

// NewValueFlowJettonsInner instantiates a new ValueFlowJettonsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueFlowJettonsInner(account AccountAddress, jetton JettonPreview, quantity int64) *ValueFlowJettonsInner {
	this := ValueFlowJettonsInner{}
	this.Account = account
	this.Jetton = jetton
	this.Quantity = quantity
	return &this
}

// NewValueFlowJettonsInnerWithDefaults instantiates a new ValueFlowJettonsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueFlowJettonsInnerWithDefaults() *ValueFlowJettonsInner {
	this := ValueFlowJettonsInner{}
	return &this
}

// GetAccount returns the Account field value
func (o *ValueFlowJettonsInner) GetAccount() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *ValueFlowJettonsInner) GetAccountOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *ValueFlowJettonsInner) SetAccount(v AccountAddress) {
	o.Account = v
}

// GetJetton returns the Jetton field value
func (o *ValueFlowJettonsInner) GetJetton() JettonPreview {
	if o == nil {
		var ret JettonPreview
		return ret
	}

	return o.Jetton
}

// GetJettonOk returns a tuple with the Jetton field value
// and a boolean to check if the value has been set.
func (o *ValueFlowJettonsInner) GetJettonOk() (*JettonPreview, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Jetton, true
}

// SetJetton sets field value
func (o *ValueFlowJettonsInner) SetJetton(v JettonPreview) {
	o.Jetton = v
}

// GetQuantity returns the Quantity field value
func (o *ValueFlowJettonsInner) GetQuantity() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *ValueFlowJettonsInner) GetQuantityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *ValueFlowJettonsInner) SetQuantity(v int64) {
	o.Quantity = v
}

func (o ValueFlowJettonsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueFlowJettonsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	toSerialize["jetton"] = o.Jetton
	toSerialize["quantity"] = o.Quantity
	return toSerialize, nil
}

func (o *ValueFlowJettonsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account",
		"jetton",
		"quantity",
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

	varValueFlowJettonsInner := _ValueFlowJettonsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValueFlowJettonsInner)

	if err != nil {
		return err
	}

	*o = ValueFlowJettonsInner(varValueFlowJettonsInner)

	return err
}

type NullableValueFlowJettonsInner struct {
	value *ValueFlowJettonsInner
	isSet bool
}

func (v NullableValueFlowJettonsInner) Get() *ValueFlowJettonsInner {
	return v.value
}

func (v *NullableValueFlowJettonsInner) Set(val *ValueFlowJettonsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValueFlowJettonsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValueFlowJettonsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueFlowJettonsInner(val *ValueFlowJettonsInner) *NullableValueFlowJettonsInner {
	return &NullableValueFlowJettonsInner{value: val, isSet: true}
}

func (v NullableValueFlowJettonsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueFlowJettonsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


