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

// checks if the Trace type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Trace{}

// Trace struct for Trace
type Trace struct {
	Transaction Transaction `json:"transaction"`
	Interfaces []string `json:"interfaces"`
	Children []Trace `json:"children,omitempty"`
	Emulated *bool `json:"emulated,omitempty"`
}

type _Trace Trace

// NewTrace instantiates a new Trace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrace(transaction Transaction, interfaces []string) *Trace {
	this := Trace{}
	this.Transaction = transaction
	this.Interfaces = interfaces
	return &this
}

// NewTraceWithDefaults instantiates a new Trace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTraceWithDefaults() *Trace {
	this := Trace{}
	return &this
}

// GetTransaction returns the Transaction field value
func (o *Trace) GetTransaction() Transaction {
	if o == nil {
		var ret Transaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *Trace) GetTransactionOk() (*Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *Trace) SetTransaction(v Transaction) {
	o.Transaction = v
}

// GetInterfaces returns the Interfaces field value
func (o *Trace) GetInterfaces() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value
// and a boolean to check if the value has been set.
func (o *Trace) GetInterfacesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Interfaces, true
}

// SetInterfaces sets field value
func (o *Trace) SetInterfaces(v []string) {
	o.Interfaces = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *Trace) GetChildren() []Trace {
	if o == nil || IsNil(o.Children) {
		var ret []Trace
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trace) GetChildrenOk() ([]Trace, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *Trace) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []Trace and assigns it to the Children field.
func (o *Trace) SetChildren(v []Trace) {
	o.Children = v
}

// GetEmulated returns the Emulated field value if set, zero value otherwise.
func (o *Trace) GetEmulated() bool {
	if o == nil || IsNil(o.Emulated) {
		var ret bool
		return ret
	}
	return *o.Emulated
}

// GetEmulatedOk returns a tuple with the Emulated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Trace) GetEmulatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Emulated) {
		return nil, false
	}
	return o.Emulated, true
}

// HasEmulated returns a boolean if a field has been set.
func (o *Trace) HasEmulated() bool {
	if o != nil && !IsNil(o.Emulated) {
		return true
	}

	return false
}

// SetEmulated gets a reference to the given bool and assigns it to the Emulated field.
func (o *Trace) SetEmulated(v bool) {
	o.Emulated = &v
}

func (o Trace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Trace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transaction"] = o.Transaction
	toSerialize["interfaces"] = o.Interfaces
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Emulated) {
		toSerialize["emulated"] = o.Emulated
	}
	return toSerialize, nil
}

func (o *Trace) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transaction",
		"interfaces",
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

	varTrace := _Trace{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTrace)

	if err != nil {
		return err
	}

	*o = Trace(varTrace)

	return err
}

type NullableTrace struct {
	value *Trace
	isSet bool
}

func (v NullableTrace) Get() *Trace {
	return v.value
}

func (v *NullableTrace) Set(val *Trace) {
	v.value = val
	v.isSet = true
}

func (v NullableTrace) IsSet() bool {
	return v.isSet
}

func (v *NullableTrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrace(val *Trace) *NullableTrace {
	return &NullableTrace{value: val, isSet: true}
}

func (v NullableTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


