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

// checks if the BlockchainConfig13 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainConfig13{}

// BlockchainConfig13 The cost of filing complaints about incorrect operation of validators.
type BlockchainConfig13 struct {
	Deposit int64 `json:"deposit"`
	BitPrice int64 `json:"bit_price"`
	CellPrice int64 `json:"cell_price"`
}

type _BlockchainConfig13 BlockchainConfig13

// NewBlockchainConfig13 instantiates a new BlockchainConfig13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainConfig13(deposit int64, bitPrice int64, cellPrice int64) *BlockchainConfig13 {
	this := BlockchainConfig13{}
	this.Deposit = deposit
	this.BitPrice = bitPrice
	this.CellPrice = cellPrice
	return &this
}

// NewBlockchainConfig13WithDefaults instantiates a new BlockchainConfig13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainConfig13WithDefaults() *BlockchainConfig13 {
	this := BlockchainConfig13{}
	return &this
}

// GetDeposit returns the Deposit field value
func (o *BlockchainConfig13) GetDeposit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Deposit
}

// GetDepositOk returns a tuple with the Deposit field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig13) GetDepositOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deposit, true
}

// SetDeposit sets field value
func (o *BlockchainConfig13) SetDeposit(v int64) {
	o.Deposit = v
}

// GetBitPrice returns the BitPrice field value
func (o *BlockchainConfig13) GetBitPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BitPrice
}

// GetBitPriceOk returns a tuple with the BitPrice field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig13) GetBitPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BitPrice, true
}

// SetBitPrice sets field value
func (o *BlockchainConfig13) SetBitPrice(v int64) {
	o.BitPrice = v
}

// GetCellPrice returns the CellPrice field value
func (o *BlockchainConfig13) GetCellPrice() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CellPrice
}

// GetCellPriceOk returns a tuple with the CellPrice field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig13) GetCellPriceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CellPrice, true
}

// SetCellPrice sets field value
func (o *BlockchainConfig13) SetCellPrice(v int64) {
	o.CellPrice = v
}

func (o BlockchainConfig13) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainConfig13) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["deposit"] = o.Deposit
	toSerialize["bit_price"] = o.BitPrice
	toSerialize["cell_price"] = o.CellPrice
	return toSerialize, nil
}

func (o *BlockchainConfig13) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"deposit",
		"bit_price",
		"cell_price",
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

	varBlockchainConfig13 := _BlockchainConfig13{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockchainConfig13)

	if err != nil {
		return err
	}

	*o = BlockchainConfig13(varBlockchainConfig13)

	return err
}

type NullableBlockchainConfig13 struct {
	value *BlockchainConfig13
	isSet bool
}

func (v NullableBlockchainConfig13) Get() *BlockchainConfig13 {
	return v.value
}

func (v *NullableBlockchainConfig13) Set(val *BlockchainConfig13) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainConfig13) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainConfig13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainConfig13(val *BlockchainConfig13) *NullableBlockchainConfig13 {
	return &NullableBlockchainConfig13{value: val, isSet: true}
}

func (v NullableBlockchainConfig13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainConfig13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

