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
)

// checks if the DecodedMessageExtInMsgDecoded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecodedMessageExtInMsgDecoded{}

// DecodedMessageExtInMsgDecoded struct for DecodedMessageExtInMsgDecoded
type DecodedMessageExtInMsgDecoded struct {
	WalletV3 *DecodedMessageExtInMsgDecodedWalletV3 `json:"wallet_v3,omitempty"`
	WalletV4 *DecodedMessageExtInMsgDecodedWalletV4 `json:"wallet_v4,omitempty"`
	WalletHighloadV2 *DecodedMessageExtInMsgDecodedWalletHighloadV2 `json:"wallet_highload_v2,omitempty"`
}

// NewDecodedMessageExtInMsgDecoded instantiates a new DecodedMessageExtInMsgDecoded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodedMessageExtInMsgDecoded() *DecodedMessageExtInMsgDecoded {
	this := DecodedMessageExtInMsgDecoded{}
	return &this
}

// NewDecodedMessageExtInMsgDecodedWithDefaults instantiates a new DecodedMessageExtInMsgDecoded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodedMessageExtInMsgDecodedWithDefaults() *DecodedMessageExtInMsgDecoded {
	this := DecodedMessageExtInMsgDecoded{}
	return &this
}

// GetWalletV3 returns the WalletV3 field value if set, zero value otherwise.
func (o *DecodedMessageExtInMsgDecoded) GetWalletV3() DecodedMessageExtInMsgDecodedWalletV3 {
	if o == nil || IsNil(o.WalletV3) {
		var ret DecodedMessageExtInMsgDecodedWalletV3
		return ret
	}
	return *o.WalletV3
}

// GetWalletV3Ok returns a tuple with the WalletV3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecoded) GetWalletV3Ok() (*DecodedMessageExtInMsgDecodedWalletV3, bool) {
	if o == nil || IsNil(o.WalletV3) {
		return nil, false
	}
	return o.WalletV3, true
}

// HasWalletV3 returns a boolean if a field has been set.
func (o *DecodedMessageExtInMsgDecoded) HasWalletV3() bool {
	if o != nil && !IsNil(o.WalletV3) {
		return true
	}

	return false
}

// SetWalletV3 gets a reference to the given DecodedMessageExtInMsgDecodedWalletV3 and assigns it to the WalletV3 field.
func (o *DecodedMessageExtInMsgDecoded) SetWalletV3(v DecodedMessageExtInMsgDecodedWalletV3) {
	o.WalletV3 = &v
}

// GetWalletV4 returns the WalletV4 field value if set, zero value otherwise.
func (o *DecodedMessageExtInMsgDecoded) GetWalletV4() DecodedMessageExtInMsgDecodedWalletV4 {
	if o == nil || IsNil(o.WalletV4) {
		var ret DecodedMessageExtInMsgDecodedWalletV4
		return ret
	}
	return *o.WalletV4
}

// GetWalletV4Ok returns a tuple with the WalletV4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecoded) GetWalletV4Ok() (*DecodedMessageExtInMsgDecodedWalletV4, bool) {
	if o == nil || IsNil(o.WalletV4) {
		return nil, false
	}
	return o.WalletV4, true
}

// HasWalletV4 returns a boolean if a field has been set.
func (o *DecodedMessageExtInMsgDecoded) HasWalletV4() bool {
	if o != nil && !IsNil(o.WalletV4) {
		return true
	}

	return false
}

// SetWalletV4 gets a reference to the given DecodedMessageExtInMsgDecodedWalletV4 and assigns it to the WalletV4 field.
func (o *DecodedMessageExtInMsgDecoded) SetWalletV4(v DecodedMessageExtInMsgDecodedWalletV4) {
	o.WalletV4 = &v
}

// GetWalletHighloadV2 returns the WalletHighloadV2 field value if set, zero value otherwise.
func (o *DecodedMessageExtInMsgDecoded) GetWalletHighloadV2() DecodedMessageExtInMsgDecodedWalletHighloadV2 {
	if o == nil || IsNil(o.WalletHighloadV2) {
		var ret DecodedMessageExtInMsgDecodedWalletHighloadV2
		return ret
	}
	return *o.WalletHighloadV2
}

// GetWalletHighloadV2Ok returns a tuple with the WalletHighloadV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecoded) GetWalletHighloadV2Ok() (*DecodedMessageExtInMsgDecodedWalletHighloadV2, bool) {
	if o == nil || IsNil(o.WalletHighloadV2) {
		return nil, false
	}
	return o.WalletHighloadV2, true
}

// HasWalletHighloadV2 returns a boolean if a field has been set.
func (o *DecodedMessageExtInMsgDecoded) HasWalletHighloadV2() bool {
	if o != nil && !IsNil(o.WalletHighloadV2) {
		return true
	}

	return false
}

// SetWalletHighloadV2 gets a reference to the given DecodedMessageExtInMsgDecodedWalletHighloadV2 and assigns it to the WalletHighloadV2 field.
func (o *DecodedMessageExtInMsgDecoded) SetWalletHighloadV2(v DecodedMessageExtInMsgDecodedWalletHighloadV2) {
	o.WalletHighloadV2 = &v
}

func (o DecodedMessageExtInMsgDecoded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecodedMessageExtInMsgDecoded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WalletV3) {
		toSerialize["wallet_v3"] = o.WalletV3
	}
	if !IsNil(o.WalletV4) {
		toSerialize["wallet_v4"] = o.WalletV4
	}
	if !IsNil(o.WalletHighloadV2) {
		toSerialize["wallet_highload_v2"] = o.WalletHighloadV2
	}
	return toSerialize, nil
}

type NullableDecodedMessageExtInMsgDecoded struct {
	value *DecodedMessageExtInMsgDecoded
	isSet bool
}

func (v NullableDecodedMessageExtInMsgDecoded) Get() *DecodedMessageExtInMsgDecoded {
	return v.value
}

func (v *NullableDecodedMessageExtInMsgDecoded) Set(val *DecodedMessageExtInMsgDecoded) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodedMessageExtInMsgDecoded) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodedMessageExtInMsgDecoded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodedMessageExtInMsgDecoded(val *DecodedMessageExtInMsgDecoded) *NullableDecodedMessageExtInMsgDecoded {
	return &NullableDecodedMessageExtInMsgDecoded{value: val, isSet: true}
}

func (v NullableDecodedMessageExtInMsgDecoded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodedMessageExtInMsgDecoded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


