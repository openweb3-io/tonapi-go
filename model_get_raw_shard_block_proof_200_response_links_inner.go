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

// checks if the GetRawShardBlockProof200ResponseLinksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawShardBlockProof200ResponseLinksInner{}

// GetRawShardBlockProof200ResponseLinksInner struct for GetRawShardBlockProof200ResponseLinksInner
type GetRawShardBlockProof200ResponseLinksInner struct {
	Id BlockRaw `json:"id"`
	Proof string `json:"proof"`
}

type _GetRawShardBlockProof200ResponseLinksInner GetRawShardBlockProof200ResponseLinksInner

// NewGetRawShardBlockProof200ResponseLinksInner instantiates a new GetRawShardBlockProof200ResponseLinksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawShardBlockProof200ResponseLinksInner(id BlockRaw, proof string) *GetRawShardBlockProof200ResponseLinksInner {
	this := GetRawShardBlockProof200ResponseLinksInner{}
	this.Id = id
	this.Proof = proof
	return &this
}

// NewGetRawShardBlockProof200ResponseLinksInnerWithDefaults instantiates a new GetRawShardBlockProof200ResponseLinksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawShardBlockProof200ResponseLinksInnerWithDefaults() *GetRawShardBlockProof200ResponseLinksInner {
	this := GetRawShardBlockProof200ResponseLinksInner{}
	return &this
}

// GetId returns the Id field value
func (o *GetRawShardBlockProof200ResponseLinksInner) GetId() BlockRaw {
	if o == nil {
		var ret BlockRaw
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetRawShardBlockProof200ResponseLinksInner) GetIdOk() (*BlockRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetRawShardBlockProof200ResponseLinksInner) SetId(v BlockRaw) {
	o.Id = v
}

// GetProof returns the Proof field value
func (o *GetRawShardBlockProof200ResponseLinksInner) GetProof() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Proof
}

// GetProofOk returns a tuple with the Proof field value
// and a boolean to check if the value has been set.
func (o *GetRawShardBlockProof200ResponseLinksInner) GetProofOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Proof, true
}

// SetProof sets field value
func (o *GetRawShardBlockProof200ResponseLinksInner) SetProof(v string) {
	o.Proof = v
}

func (o GetRawShardBlockProof200ResponseLinksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawShardBlockProof200ResponseLinksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["proof"] = o.Proof
	return toSerialize, nil
}

func (o *GetRawShardBlockProof200ResponseLinksInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"proof",
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

	varGetRawShardBlockProof200ResponseLinksInner := _GetRawShardBlockProof200ResponseLinksInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawShardBlockProof200ResponseLinksInner)

	if err != nil {
		return err
	}

	*o = GetRawShardBlockProof200ResponseLinksInner(varGetRawShardBlockProof200ResponseLinksInner)

	return err
}

type NullableGetRawShardBlockProof200ResponseLinksInner struct {
	value *GetRawShardBlockProof200ResponseLinksInner
	isSet bool
}

func (v NullableGetRawShardBlockProof200ResponseLinksInner) Get() *GetRawShardBlockProof200ResponseLinksInner {
	return v.value
}

func (v *NullableGetRawShardBlockProof200ResponseLinksInner) Set(val *GetRawShardBlockProof200ResponseLinksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawShardBlockProof200ResponseLinksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawShardBlockProof200ResponseLinksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawShardBlockProof200ResponseLinksInner(val *GetRawShardBlockProof200ResponseLinksInner) *NullableGetRawShardBlockProof200ResponseLinksInner {
	return &NullableGetRawShardBlockProof200ResponseLinksInner{value: val, isSet: true}
}

func (v NullableGetRawShardBlockProof200ResponseLinksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawShardBlockProof200ResponseLinksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

