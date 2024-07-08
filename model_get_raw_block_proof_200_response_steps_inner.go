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

// checks if the GetRawBlockProof200ResponseStepsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawBlockProof200ResponseStepsInner{}

// GetRawBlockProof200ResponseStepsInner struct for GetRawBlockProof200ResponseStepsInner
type GetRawBlockProof200ResponseStepsInner struct {
	LiteServerBlockLinkBack GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack `json:"lite_server_block_link_back"`
	LiteServerBlockLinkForward GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward `json:"lite_server_block_link_forward"`
}

type _GetRawBlockProof200ResponseStepsInner GetRawBlockProof200ResponseStepsInner

// NewGetRawBlockProof200ResponseStepsInner instantiates a new GetRawBlockProof200ResponseStepsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawBlockProof200ResponseStepsInner(liteServerBlockLinkBack GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack, liteServerBlockLinkForward GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward) *GetRawBlockProof200ResponseStepsInner {
	this := GetRawBlockProof200ResponseStepsInner{}
	this.LiteServerBlockLinkBack = liteServerBlockLinkBack
	this.LiteServerBlockLinkForward = liteServerBlockLinkForward
	return &this
}

// NewGetRawBlockProof200ResponseStepsInnerWithDefaults instantiates a new GetRawBlockProof200ResponseStepsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawBlockProof200ResponseStepsInnerWithDefaults() *GetRawBlockProof200ResponseStepsInner {
	this := GetRawBlockProof200ResponseStepsInner{}
	return &this
}

// GetLiteServerBlockLinkBack returns the LiteServerBlockLinkBack field value
func (o *GetRawBlockProof200ResponseStepsInner) GetLiteServerBlockLinkBack() GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack {
	if o == nil {
		var ret GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack
		return ret
	}

	return o.LiteServerBlockLinkBack
}

// GetLiteServerBlockLinkBackOk returns a tuple with the LiteServerBlockLinkBack field value
// and a boolean to check if the value has been set.
func (o *GetRawBlockProof200ResponseStepsInner) GetLiteServerBlockLinkBackOk() (*GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiteServerBlockLinkBack, true
}

// SetLiteServerBlockLinkBack sets field value
func (o *GetRawBlockProof200ResponseStepsInner) SetLiteServerBlockLinkBack(v GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkBack) {
	o.LiteServerBlockLinkBack = v
}

// GetLiteServerBlockLinkForward returns the LiteServerBlockLinkForward field value
func (o *GetRawBlockProof200ResponseStepsInner) GetLiteServerBlockLinkForward() GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward {
	if o == nil {
		var ret GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward
		return ret
	}

	return o.LiteServerBlockLinkForward
}

// GetLiteServerBlockLinkForwardOk returns a tuple with the LiteServerBlockLinkForward field value
// and a boolean to check if the value has been set.
func (o *GetRawBlockProof200ResponseStepsInner) GetLiteServerBlockLinkForwardOk() (*GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LiteServerBlockLinkForward, true
}

// SetLiteServerBlockLinkForward sets field value
func (o *GetRawBlockProof200ResponseStepsInner) SetLiteServerBlockLinkForward(v GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForward) {
	o.LiteServerBlockLinkForward = v
}

func (o GetRawBlockProof200ResponseStepsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawBlockProof200ResponseStepsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lite_server_block_link_back"] = o.LiteServerBlockLinkBack
	toSerialize["lite_server_block_link_forward"] = o.LiteServerBlockLinkForward
	return toSerialize, nil
}

func (o *GetRawBlockProof200ResponseStepsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"lite_server_block_link_back",
		"lite_server_block_link_forward",
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

	varGetRawBlockProof200ResponseStepsInner := _GetRawBlockProof200ResponseStepsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawBlockProof200ResponseStepsInner)

	if err != nil {
		return err
	}

	*o = GetRawBlockProof200ResponseStepsInner(varGetRawBlockProof200ResponseStepsInner)

	return err
}

type NullableGetRawBlockProof200ResponseStepsInner struct {
	value *GetRawBlockProof200ResponseStepsInner
	isSet bool
}

func (v NullableGetRawBlockProof200ResponseStepsInner) Get() *GetRawBlockProof200ResponseStepsInner {
	return v.value
}

func (v *NullableGetRawBlockProof200ResponseStepsInner) Set(val *GetRawBlockProof200ResponseStepsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawBlockProof200ResponseStepsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawBlockProof200ResponseStepsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawBlockProof200ResponseStepsInner(val *GetRawBlockProof200ResponseStepsInner) *NullableGetRawBlockProof200ResponseStepsInner {
	return &NullableGetRawBlockProof200ResponseStepsInner{value: val, isSet: true}
}

func (v NullableGetRawBlockProof200ResponseStepsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawBlockProof200ResponseStepsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

