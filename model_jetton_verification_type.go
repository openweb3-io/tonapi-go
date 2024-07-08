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

// JettonVerificationType the model 'JettonVerificationType'
type JettonVerificationType string

// List of JettonVerificationType
const (
	JETTONVERIFICATIONTYPE_WHITELIST JettonVerificationType = "whitelist"
	JETTONVERIFICATIONTYPE_BLACKLIST JettonVerificationType = "blacklist"
	JETTONVERIFICATIONTYPE_NONE JettonVerificationType = "none"
)

// All allowed values of JettonVerificationType enum
var AllowedJettonVerificationTypeEnumValues = []JettonVerificationType{
	"whitelist",
	"blacklist",
	"none",
}

func (v *JettonVerificationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JettonVerificationType(value)
	for _, existing := range AllowedJettonVerificationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JettonVerificationType", value)
}

// NewJettonVerificationTypeFromValue returns a pointer to a valid JettonVerificationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJettonVerificationTypeFromValue(v string) (*JettonVerificationType, error) {
	ev := JettonVerificationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JettonVerificationType: valid values are %v", v, AllowedJettonVerificationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JettonVerificationType) IsValid() bool {
	for _, existing := range AllowedJettonVerificationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to JettonVerificationType value
func (v JettonVerificationType) Ptr() *JettonVerificationType {
	return &v
}

type NullableJettonVerificationType struct {
	value *JettonVerificationType
	isSet bool
}

func (v NullableJettonVerificationType) Get() *JettonVerificationType {
	return v.value
}

func (v *NullableJettonVerificationType) Set(val *JettonVerificationType) {
	v.value = val
	v.isSet = true
}

func (v NullableJettonVerificationType) IsSet() bool {
	return v.isSet
}

func (v *NullableJettonVerificationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJettonVerificationType(val *JettonVerificationType) *NullableJettonVerificationType {
	return &NullableJettonVerificationType{value: val, isSet: true}
}

func (v NullableJettonVerificationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJettonVerificationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

