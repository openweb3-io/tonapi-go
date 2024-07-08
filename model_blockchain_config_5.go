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

// checks if the BlockchainConfig5 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainConfig5{}

// BlockchainConfig5 struct for BlockchainConfig5
type BlockchainConfig5 struct {
	BlackholeAddr *string `json:"blackhole_addr,omitempty"`
	FeeBurnNom int64 `json:"fee_burn_nom"`
	FeeBurnDenom int64 `json:"fee_burn_denom"`
}

type _BlockchainConfig5 BlockchainConfig5

// NewBlockchainConfig5 instantiates a new BlockchainConfig5 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainConfig5(feeBurnNom int64, feeBurnDenom int64) *BlockchainConfig5 {
	this := BlockchainConfig5{}
	this.FeeBurnNom = feeBurnNom
	this.FeeBurnDenom = feeBurnDenom
	return &this
}

// NewBlockchainConfig5WithDefaults instantiates a new BlockchainConfig5 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainConfig5WithDefaults() *BlockchainConfig5 {
	this := BlockchainConfig5{}
	return &this
}

// GetBlackholeAddr returns the BlackholeAddr field value if set, zero value otherwise.
func (o *BlockchainConfig5) GetBlackholeAddr() string {
	if o == nil || IsNil(o.BlackholeAddr) {
		var ret string
		return ret
	}
	return *o.BlackholeAddr
}

// GetBlackholeAddrOk returns a tuple with the BlackholeAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainConfig5) GetBlackholeAddrOk() (*string, bool) {
	if o == nil || IsNil(o.BlackholeAddr) {
		return nil, false
	}
	return o.BlackholeAddr, true
}

// HasBlackholeAddr returns a boolean if a field has been set.
func (o *BlockchainConfig5) HasBlackholeAddr() bool {
	if o != nil && !IsNil(o.BlackholeAddr) {
		return true
	}

	return false
}

// SetBlackholeAddr gets a reference to the given string and assigns it to the BlackholeAddr field.
func (o *BlockchainConfig5) SetBlackholeAddr(v string) {
	o.BlackholeAddr = &v
}

// GetFeeBurnNom returns the FeeBurnNom field value
func (o *BlockchainConfig5) GetFeeBurnNom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FeeBurnNom
}

// GetFeeBurnNomOk returns a tuple with the FeeBurnNom field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig5) GetFeeBurnNomOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeBurnNom, true
}

// SetFeeBurnNom sets field value
func (o *BlockchainConfig5) SetFeeBurnNom(v int64) {
	o.FeeBurnNom = v
}

// GetFeeBurnDenom returns the FeeBurnDenom field value
func (o *BlockchainConfig5) GetFeeBurnDenom() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FeeBurnDenom
}

// GetFeeBurnDenomOk returns a tuple with the FeeBurnDenom field value
// and a boolean to check if the value has been set.
func (o *BlockchainConfig5) GetFeeBurnDenomOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeBurnDenom, true
}

// SetFeeBurnDenom sets field value
func (o *BlockchainConfig5) SetFeeBurnDenom(v int64) {
	o.FeeBurnDenom = v
}

func (o BlockchainConfig5) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainConfig5) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BlackholeAddr) {
		toSerialize["blackhole_addr"] = o.BlackholeAddr
	}
	toSerialize["fee_burn_nom"] = o.FeeBurnNom
	toSerialize["fee_burn_denom"] = o.FeeBurnDenom
	return toSerialize, nil
}

func (o *BlockchainConfig5) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fee_burn_nom",
		"fee_burn_denom",
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

	varBlockchainConfig5 := _BlockchainConfig5{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockchainConfig5)

	if err != nil {
		return err
	}

	*o = BlockchainConfig5(varBlockchainConfig5)

	return err
}

type NullableBlockchainConfig5 struct {
	value *BlockchainConfig5
	isSet bool
}

func (v NullableBlockchainConfig5) Get() *BlockchainConfig5 {
	return v.value
}

func (v *NullableBlockchainConfig5) Set(val *BlockchainConfig5) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainConfig5) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainConfig5) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainConfig5(val *BlockchainConfig5) *NullableBlockchainConfig5 {
	return &NullableBlockchainConfig5{value: val, isSet: true}
}

func (v NullableBlockchainConfig5) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainConfig5) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


