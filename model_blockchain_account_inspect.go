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

// checks if the BlockchainAccountInspect type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockchainAccountInspect{}

// BlockchainAccountInspect struct for BlockchainAccountInspect
type BlockchainAccountInspect struct {
	Code string `json:"code"`
	CodeHash string `json:"code_hash"`
	Methods []BlockchainAccountInspectMethodsInner `json:"methods"`
	Compiler *string `json:"compiler,omitempty"`
}

type _BlockchainAccountInspect BlockchainAccountInspect

// NewBlockchainAccountInspect instantiates a new BlockchainAccountInspect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockchainAccountInspect(code string, codeHash string, methods []BlockchainAccountInspectMethodsInner) *BlockchainAccountInspect {
	this := BlockchainAccountInspect{}
	this.Code = code
	this.CodeHash = codeHash
	this.Methods = methods
	return &this
}

// NewBlockchainAccountInspectWithDefaults instantiates a new BlockchainAccountInspect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockchainAccountInspectWithDefaults() *BlockchainAccountInspect {
	this := BlockchainAccountInspect{}
	return &this
}

// GetCode returns the Code field value
func (o *BlockchainAccountInspect) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BlockchainAccountInspect) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BlockchainAccountInspect) SetCode(v string) {
	o.Code = v
}

// GetCodeHash returns the CodeHash field value
func (o *BlockchainAccountInspect) GetCodeHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeHash
}

// GetCodeHashOk returns a tuple with the CodeHash field value
// and a boolean to check if the value has been set.
func (o *BlockchainAccountInspect) GetCodeHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeHash, true
}

// SetCodeHash sets field value
func (o *BlockchainAccountInspect) SetCodeHash(v string) {
	o.CodeHash = v
}

// GetMethods returns the Methods field value
func (o *BlockchainAccountInspect) GetMethods() []BlockchainAccountInspectMethodsInner {
	if o == nil {
		var ret []BlockchainAccountInspectMethodsInner
		return ret
	}

	return o.Methods
}

// GetMethodsOk returns a tuple with the Methods field value
// and a boolean to check if the value has been set.
func (o *BlockchainAccountInspect) GetMethodsOk() ([]BlockchainAccountInspectMethodsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Methods, true
}

// SetMethods sets field value
func (o *BlockchainAccountInspect) SetMethods(v []BlockchainAccountInspectMethodsInner) {
	o.Methods = v
}

// GetCompiler returns the Compiler field value if set, zero value otherwise.
func (o *BlockchainAccountInspect) GetCompiler() string {
	if o == nil || IsNil(o.Compiler) {
		var ret string
		return ret
	}
	return *o.Compiler
}

// GetCompilerOk returns a tuple with the Compiler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockchainAccountInspect) GetCompilerOk() (*string, bool) {
	if o == nil || IsNil(o.Compiler) {
		return nil, false
	}
	return o.Compiler, true
}

// HasCompiler returns a boolean if a field has been set.
func (o *BlockchainAccountInspect) HasCompiler() bool {
	if o != nil && !IsNil(o.Compiler) {
		return true
	}

	return false
}

// SetCompiler gets a reference to the given string and assigns it to the Compiler field.
func (o *BlockchainAccountInspect) SetCompiler(v string) {
	o.Compiler = &v
}

func (o BlockchainAccountInspect) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockchainAccountInspect) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["code_hash"] = o.CodeHash
	toSerialize["methods"] = o.Methods
	if !IsNil(o.Compiler) {
		toSerialize["compiler"] = o.Compiler
	}
	return toSerialize, nil
}

func (o *BlockchainAccountInspect) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"code_hash",
		"methods",
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

	varBlockchainAccountInspect := _BlockchainAccountInspect{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockchainAccountInspect)

	if err != nil {
		return err
	}

	*o = BlockchainAccountInspect(varBlockchainAccountInspect)

	return err
}

type NullableBlockchainAccountInspect struct {
	value *BlockchainAccountInspect
	isSet bool
}

func (v NullableBlockchainAccountInspect) Get() *BlockchainAccountInspect {
	return v.value
}

func (v *NullableBlockchainAccountInspect) Set(val *BlockchainAccountInspect) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockchainAccountInspect) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockchainAccountInspect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockchainAccountInspect(val *BlockchainAccountInspect) *NullableBlockchainAccountInspect {
	return &NullableBlockchainAccountInspect{value: val, isSet: true}
}

func (v NullableBlockchainAccountInspect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockchainAccountInspect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


