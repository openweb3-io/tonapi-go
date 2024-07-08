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

// checks if the Transactions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transactions{}

// Transactions struct for Transactions
type Transactions struct {
	Transactions []Transaction `json:"transactions"`
}

type _Transactions Transactions

// NewTransactions instantiates a new Transactions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactions(transactions []Transaction) *Transactions {
	this := Transactions{}
	this.Transactions = transactions
	return &this
}

// NewTransactionsWithDefaults instantiates a new Transactions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsWithDefaults() *Transactions {
	this := Transactions{}
	return &this
}

// GetTransactions returns the Transactions field value
func (o *Transactions) GetTransactions() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *Transactions) GetTransactionsOk() ([]Transaction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Transactions, true
}

// SetTransactions sets field value
func (o *Transactions) SetTransactions(v []Transaction) {
	o.Transactions = v
}

func (o Transactions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transactions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

func (o *Transactions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"transactions",
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

	varTransactions := _Transactions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransactions)

	if err != nil {
		return err
	}

	*o = Transactions(varTransactions)

	return err
}

type NullableTransactions struct {
	value *Transactions
	isSet bool
}

func (v NullableTransactions) Get() *Transactions {
	return v.value
}

func (v *NullableTransactions) Set(val *Transactions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactions(val *Transactions) *NullableTransactions {
	return &NullableTransactions{value: val, isSet: true}
}

func (v NullableTransactions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


