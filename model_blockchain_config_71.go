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

// checks if the BlockchainConfig71 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainConfig71{}

// BlockchainConfig71 Bridge parameters for wrapping TON in other networks.
type BlockchainConfig71 struct {
	OracleBridgeParams OracleBridgeParams `json:"oracle_bridge_params"`
}

type _BlockchainConfig71 BlockchainConfig71

// NewBlockchainConfig71 instantiates a new BlockchainConfig71 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainConfig71(oracleBridgeParams OracleBridgeParams) *BlockchainConfig71 {
	this := BlockchainConfig71{}
	this.OracleBridgeParams = oracleBridgeParams
	return &this
}

// NewBlockchainConfig71WithDefaults instantiates a new BlockchainConfig71 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainConfig71WithDefaults() *BlockchainConfig71 {
	this := BlockchainConfig71{}
	return &this
}

// GetOracleBridgeParams returns the OracleBridgeParams field value
func (o *BlockchainConfig71) GetOracleBridgeParams() OracleBridgeParams {
	if o == nil {
		var ret OracleBridgeParams
		return ret
	}

	return o.OracleBridgeParams
}

// GetOracleBridgeParamsOk returns a tuple with the OracleBridgeParams field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig71) GetOracleBridgeParamsOk() (*OracleBridgeParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OracleBridgeParams, true
}

// SetOracleBridgeParams sets field value
func (o *BlockchainConfig71) SetOracleBridgeParams(v OracleBridgeParams) {
	o.OracleBridgeParams = v
}

func (o BlockchainConfig71) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainConfig71) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oracle_bridge_params"] = o.OracleBridgeParams
	return toSerialize, nil
}

func (o *BlockchainConfig71) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"oracle_bridge_params",
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

	varBlockchainConfig71 := _BlockchainConfig71{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockchainConfig71)

	if err != nil {
		return err
	}

	*o = BlockchainConfig71(varBlockchainConfig71)

	return err
}

type NullableBlockchainConfig71 struct {
	value *BlockchainConfig71
	isSet bool
}

func (v NullableBlockchainConfig71) Get() *BlockchainConfig71 {
	return v.value
}

func (v *NullableBlockchainConfig71) Set(val *BlockchainConfig71) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainConfig71) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainConfig71) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainConfig71(val *BlockchainConfig71) *NullableBlockchainConfig71 {
	return &NullableBlockchainConfig71{value: val, isSet: true}
}

func (v NullableBlockchainConfig71) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainConfig71) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


