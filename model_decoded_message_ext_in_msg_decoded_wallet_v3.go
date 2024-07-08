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

// checks if the DecodedMessageExtInMsgDecodedWalletV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DecodedMessageExtInMsgDecodedWalletV3{}

// DecodedMessageExtInMsgDecodedWalletV3 struct for DecodedMessageExtInMsgDecodedWalletV3
type DecodedMessageExtInMsgDecodedWalletV3 struct {
	SubwalletId int64 `json:"subwallet_id"`
	ValidUntil int64 `json:"valid_until"`
	Seqno int64 `json:"seqno"`
	RawMessages []DecodedRawMessage `json:"raw_messages"`
}

type _DecodedMessageExtInMsgDecodedWalletV3 DecodedMessageExtInMsgDecodedWalletV3

// NewDecodedMessageExtInMsgDecodedWalletV3 instantiates a new DecodedMessageExtInMsgDecodedWalletV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecodedMessageExtInMsgDecodedWalletV3(subwalletId int64, validUntil int64, seqno int64, rawMessages []DecodedRawMessage) *DecodedMessageExtInMsgDecodedWalletV3 {
	this := DecodedMessageExtInMsgDecodedWalletV3{}
	this.SubwalletId = subwalletId
	this.ValidUntil = validUntil
	this.Seqno = seqno
	this.RawMessages = rawMessages
	return &this
}

// NewDecodedMessageExtInMsgDecodedWalletV3WithDefaults instantiates a new DecodedMessageExtInMsgDecodedWalletV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecodedMessageExtInMsgDecodedWalletV3WithDefaults() *DecodedMessageExtInMsgDecodedWalletV3 {
	this := DecodedMessageExtInMsgDecodedWalletV3{}
	return &this
}

// GetSubwalletId returns the SubwalletId field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetSubwalletId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SubwalletId
}

// GetSubwalletIdOk returns a tuple with the SubwalletId field value
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetSubwalletIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubwalletId, true
}

// SetSubwalletId sets field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) SetSubwalletId(v int64) {
	o.SubwalletId = v
}

// GetValidUntil returns the ValidUntil field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetValidUntil() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ValidUntil
}

// GetValidUntilOk returns a tuple with the ValidUntil field value
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetValidUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidUntil, true
}

// SetValidUntil sets field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) SetValidUntil(v int64) {
	o.ValidUntil = v
}

// GetSeqno returns the Seqno field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetSeqno() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Seqno
}

// GetSeqnoOk returns a tuple with the Seqno field value
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetSeqnoOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seqno, true
}

// SetSeqno sets field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) SetSeqno(v int64) {
	o.Seqno = v
}

// GetRawMessages returns the RawMessages field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetRawMessages() []DecodedRawMessage {
	if o == nil {
		var ret []DecodedRawMessage
		return ret
	}

	return o.RawMessages
}

// GetRawMessagesOk returns a tuple with the RawMessages field value
// and a boolean to check if the value has been set.
func (o *DecodedMessageExtInMsgDecodedWalletV3) GetRawMessagesOk() ([]DecodedRawMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.RawMessages, true
}

// SetRawMessages sets field value
func (o *DecodedMessageExtInMsgDecodedWalletV3) SetRawMessages(v []DecodedRawMessage) {
	o.RawMessages = v
}

func (o DecodedMessageExtInMsgDecodedWalletV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DecodedMessageExtInMsgDecodedWalletV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subwallet_id"] = o.SubwalletId
	toSerialize["valid_until"] = o.ValidUntil
	toSerialize["seqno"] = o.Seqno
	toSerialize["raw_messages"] = o.RawMessages
	return toSerialize, nil
}

func (o *DecodedMessageExtInMsgDecodedWalletV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subwallet_id",
		"valid_until",
		"seqno",
		"raw_messages",
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

	varDecodedMessageExtInMsgDecodedWalletV3 := _DecodedMessageExtInMsgDecodedWalletV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDecodedMessageExtInMsgDecodedWalletV3)

	if err != nil {
		return err
	}

	*o = DecodedMessageExtInMsgDecodedWalletV3(varDecodedMessageExtInMsgDecodedWalletV3)

	return err
}

type NullableDecodedMessageExtInMsgDecodedWalletV3 struct {
	value *DecodedMessageExtInMsgDecodedWalletV3
	isSet bool
}

func (v NullableDecodedMessageExtInMsgDecodedWalletV3) Get() *DecodedMessageExtInMsgDecodedWalletV3 {
	return v.value
}

func (v *NullableDecodedMessageExtInMsgDecodedWalletV3) Set(val *DecodedMessageExtInMsgDecodedWalletV3) {
	v.value = val
	v.isSet = true
}

func (v NullableDecodedMessageExtInMsgDecodedWalletV3) IsSet() bool {
	return v.isSet
}

func (v *NullableDecodedMessageExtInMsgDecodedWalletV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecodedMessageExtInMsgDecodedWalletV3(val *DecodedMessageExtInMsgDecodedWalletV3) *NullableDecodedMessageExtInMsgDecodedWalletV3 {
	return &NullableDecodedMessageExtInMsgDecodedWalletV3{value: val, isSet: true}
}

func (v NullableDecodedMessageExtInMsgDecodedWalletV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecodedMessageExtInMsgDecodedWalletV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

