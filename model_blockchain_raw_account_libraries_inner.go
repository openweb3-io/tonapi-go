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

// checks if the BlockchainRawAccountLibrariesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainRawAccountLibrariesInner{}

// BlockchainRawAccountLibrariesInner struct for BlockchainRawAccountLibrariesInner
type BlockchainRawAccountLibrariesInner struct {
	Public bool `json:"public"`
	Root string `json:"root"`
}

type _BlockchainRawAccountLibrariesInner BlockchainRawAccountLibrariesInner

// NewBlockchainRawAccountLibrariesInner instantiates a new BlockchainRawAccountLibrariesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainRawAccountLibrariesInner(public bool, root string) *BlockchainRawAccountLibrariesInner {
	this := BlockchainRawAccountLibrariesInner{}
	this.Public = public
	this.Root = root
	return &this
}

// NewBlockchainRawAccountLibrariesInnerWithDefaults instantiates a new BlockchainRawAccountLibrariesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainRawAccountLibrariesInnerWithDefaults() *BlockchainRawAccountLibrariesInner {
	this := BlockchainRawAccountLibrariesInner{}
	return &this
}

// GetPublic returns the Public field value
func (o *BlockchainRawAccountLibrariesInner) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *BlockchainRawAccountLibrariesInner) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *BlockchainRawAccountLibrariesInner) SetPublic(v bool) {
	o.Public = v
}

// GetRoot returns the Root field value
func (o *BlockchainRawAccountLibrariesInner) GetRoot() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Root
}

// GetRootOk returns a tuple with the Root field value
// and a boolean to check if the value has been set.
func (o *BlockchainRawAccountLibrariesInner) GetRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Root, true
}

// SetRoot sets field value
func (o *BlockchainRawAccountLibrariesInner) SetRoot(v string) {
	o.Root = v
}

func (o BlockchainRawAccountLibrariesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainRawAccountLibrariesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public"] = o.Public
	toSerialize["root"] = o.Root
	return toSerialize, nil
}

func (o *BlockchainRawAccountLibrariesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"public",
		"root",
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

	varBlockchainRawAccountLibrariesInner := _BlockchainRawAccountLibrariesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockchainRawAccountLibrariesInner)

	if err != nil {
		return err
	}

	*o = BlockchainRawAccountLibrariesInner(varBlockchainRawAccountLibrariesInner)

	return err
}

type NullableBlockchainRawAccountLibrariesInner struct {
	value *BlockchainRawAccountLibrariesInner
	isSet bool
}

func (v NullableBlockchainRawAccountLibrariesInner) Get() *BlockchainRawAccountLibrariesInner {
	return v.value
}

func (v *NullableBlockchainRawAccountLibrariesInner) Set(val *BlockchainRawAccountLibrariesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainRawAccountLibrariesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainRawAccountLibrariesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainRawAccountLibrariesInner(val *BlockchainRawAccountLibrariesInner) *NullableBlockchainRawAccountLibrariesInner {
	return &NullableBlockchainRawAccountLibrariesInner{value: val, isSet: true}
}

func (v NullableBlockchainRawAccountLibrariesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainRawAccountLibrariesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


