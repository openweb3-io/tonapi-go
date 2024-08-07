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

// checks if the GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures{}

// GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures struct for GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures
type GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures struct {
	ValidatorSetHash int64 `json:"validator_set_hash"`
	CatchainSeqno int32 `json:"catchain_seqno"`
	Signatures []GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner `json:"signatures"`
}

type _GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures

// NewGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures instantiates a new GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures(validatorSetHash int64, catchainSeqno int32, signatures []GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner) *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures {
	this := GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures{}
	this.ValidatorSetHash = validatorSetHash
	this.CatchainSeqno = catchainSeqno
	this.Signatures = signatures
	return &this
}

// NewGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesWithDefaults instantiates a new GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesWithDefaults() *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures {
	this := GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures{}
	return &this
}

// GetValidatorSetHash returns the ValidatorSetHash field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetValidatorSetHash() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidatorSetHash
}

// GetValidatorSetHashOk returns a tuple with the ValidatorSetHash field value
// and a boolean to check if the value has been set.
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetValidatorSetHashOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidatorSetHash, true
}

// SetValidatorSetHash sets field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) SetValidatorSetHash(v int64) {
	o.ValidatorSetHash = v
}

// GetCatchainSeqno returns the CatchainSeqno field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetCatchainSeqno() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CatchainSeqno
}

// GetCatchainSeqnoOk returns a tuple with the CatchainSeqno field value
// and a boolean to check if the value has been set.
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetCatchainSeqnoOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CatchainSeqno, true
}

// SetCatchainSeqno sets field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) SetCatchainSeqno(v int32) {
	o.CatchainSeqno = v
}

// GetSignatures returns the Signatures field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetSignatures() []GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner {
	if o == nil {
		var ret []GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner
		return ret
	}

	return o.Signatures
}

// GetSignaturesOk returns a tuple with the Signatures field value
// and a boolean to check if the value has been set.
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) GetSignaturesOk() ([]GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signatures, true
}

// SetSignatures sets field value
func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) SetSignatures(v []GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignaturesSignaturesInner) {
	o.Signatures = v
}

func (o GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["validator_set_hash"] = o.ValidatorSetHash
	toSerialize["catchain_seqno"] = o.CatchainSeqno
	toSerialize["signatures"] = o.Signatures
	return toSerialize, nil
}

func (o *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"validator_set_hash",
		"catchain_seqno",
		"signatures",
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

	varGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures := _GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures)

	if err != nil {
		return err
	}

	*o = GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures(varGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures)

	return err
}

type NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures struct {
	value *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures
	isSet bool
}

func (v NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) Get() *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures {
	return v.value
}

func (v *NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) Set(val *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures(val *GetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) *NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures {
	return &NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures{value: val, isSet: true}
}

func (v NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawBlockProof200ResponseStepsInnerLiteServerBlockLinkForwardSignatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


