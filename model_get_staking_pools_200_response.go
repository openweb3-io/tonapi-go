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

// checks if the GetStakingPools200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStakingPools200Response{}

// GetStakingPools200Response struct for GetStakingPools200Response
type GetStakingPools200Response struct {
	Pools []PoolInfo `json:"pools"`
	Implementations map[string]PoolImplementation `json:"implementations"`
}

type _GetStakingPools200Response GetStakingPools200Response

// NewGetStakingPools200Response instantiates a new GetStakingPools200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStakingPools200Response(pools []PoolInfo, implementations map[string]PoolImplementation) *GetStakingPools200Response {
	this := GetStakingPools200Response{}
	this.Pools = pools
	this.Implementations = implementations
	return &this
}

// NewGetStakingPools200ResponseWithDefaults instantiates a new GetStakingPools200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStakingPools200ResponseWithDefaults() *GetStakingPools200Response {
	this := GetStakingPools200Response{}
	return &this
}

// GetPools returns the Pools field value
func (o *GetStakingPools200Response) GetPools() []PoolInfo {
	if o == nil {
		var ret []PoolInfo
		return ret
	}

	return o.Pools
}

// GetPoolsOk returns a tuple with the Pools field value
// and a boolean to check if the value has been set.
func (o *GetStakingPools200Response) GetPoolsOk() ([]PoolInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pools, true
}

// SetPools sets field value
func (o *GetStakingPools200Response) SetPools(v []PoolInfo) {
	o.Pools = v
}

// GetImplementations returns the Implementations field value
func (o *GetStakingPools200Response) GetImplementations() map[string]PoolImplementation {
	if o == nil {
		var ret map[string]PoolImplementation
		return ret
	}

	return o.Implementations
}

// GetImplementationsOk returns a tuple with the Implementations field value
// and a boolean to check if the value has been set.
func (o *GetStakingPools200Response) GetImplementationsOk() (*map[string]PoolImplementation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Implementations, true
}

// SetImplementations sets field value
func (o *GetStakingPools200Response) SetImplementations(v map[string]PoolImplementation) {
	o.Implementations = v
}

func (o GetStakingPools200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStakingPools200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pools"] = o.Pools
	toSerialize["implementations"] = o.Implementations
	return toSerialize, nil
}

func (o *GetStakingPools200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pools",
		"implementations",
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

	varGetStakingPools200Response := _GetStakingPools200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStakingPools200Response)

	if err != nil {
		return err
	}

	*o = GetStakingPools200Response(varGetStakingPools200Response)

	return err
}

type NullableGetStakingPools200Response struct {
	value *GetStakingPools200Response
	isSet bool
}

func (v NullableGetStakingPools200Response) Get() *GetStakingPools200Response {
	return v.value
}

func (v *NullableGetStakingPools200Response) Set(val *GetStakingPools200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStakingPools200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStakingPools200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStakingPools200Response(val *GetStakingPools200Response) *NullableGetStakingPools200Response {
	return &NullableGetStakingPools200Response{value: val, isSet: true}
}

func (v NullableGetStakingPools200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStakingPools200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

