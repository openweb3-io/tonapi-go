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
	"fmt"
)

// TrustType the model 'TrustType'
type TrustType string

// List of TrustType
const (
	TRUSTTYPE_WHITELIST TrustType = "whitelist"
	TRUSTTYPE_GRAYLIST TrustType = "graylist"
	TRUSTTYPE_BLACKLIST TrustType = "blacklist"
	TRUSTTYPE_NONE TrustType = "none"
)

// All allowed values of TrustType enum
var AllowedTrustTypeEnumValues = []TrustType{
	"whitelist",
	"graylist",
	"blacklist",
	"none",
}

func (v *TrustType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrustType(value)
	for _, existing := range AllowedTrustTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrustType", value)
}

// NewTrustTypeFromValue returns a pointer to a valid TrustType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrustTypeFromValue(v string) (*TrustType, error) {
	ev := TrustType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrustType: valid values are %v", v, AllowedTrustTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrustType) IsValid() bool {
	for _, existing := range AllowedTrustTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrustType value
func (v TrustType) Ptr() *TrustType {
	return &v
}

type NullableTrustType struct {
	value *TrustType
	isSet bool
}

func (v NullableTrustType) Get() *TrustType {
	return v.value
}

func (v *NullableTrustType) Set(val *TrustType) {
	v.value = val
	v.isSet = true
}

func (v NullableTrustType) IsSet() bool {
	return v.isSet
}

func (v *NullableTrustType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrustType(val *TrustType) *NullableTrustType {
	return &NullableTrustType{value: val, isSet: true}
}

func (v NullableTrustType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrustType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
