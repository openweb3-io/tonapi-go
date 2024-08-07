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

// checks if the JettonPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JettonPreview{}

// JettonPreview struct for JettonPreview
type JettonPreview struct {
	Address string `json:"address"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Decimals int32 `json:"decimals"`
	Image string `json:"image"`
	Verification JettonVerificationType `json:"verification"`
}

type _JettonPreview JettonPreview

// NewJettonPreview instantiates a new JettonPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJettonPreview(address string, name string, symbol string, decimals int32, image string, verification JettonVerificationType) *JettonPreview {
	this := JettonPreview{}
	this.Address = address
	this.Name = name
	this.Symbol = symbol
	this.Decimals = decimals
	this.Image = image
	this.Verification = verification
	return &this
}

// NewJettonPreviewWithDefaults instantiates a new JettonPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJettonPreviewWithDefaults() *JettonPreview {
	this := JettonPreview{}
	return &this
}

// GetAddress returns the Address field value
func (o *JettonPreview) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *JettonPreview) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value
func (o *JettonPreview) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JettonPreview) SetName(v string) {
	o.Name = v
}

// GetSymbol returns the Symbol field value
func (o *JettonPreview) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *JettonPreview) SetSymbol(v string) {
	o.Symbol = v
}

// GetDecimals returns the Decimals field value
func (o *JettonPreview) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *JettonPreview) SetDecimals(v int32) {
	o.Decimals = v
}

// GetImage returns the Image field value
func (o *JettonPreview) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *JettonPreview) SetImage(v string) {
	o.Image = v
}

// GetVerification returns the Verification field value
func (o *JettonPreview) GetVerification() JettonVerificationType {
	if o == nil {
		var ret JettonVerificationType
		return ret
	}

	return o.Verification
}

// GetVerificationOk returns a tuple with the Verification field value
// and a boolean to check if the value has been set.
func (o *JettonPreview) GetVerificationOk() (*JettonVerificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verification, true
}

// SetVerification sets field value
func (o *JettonPreview) SetVerification(v JettonVerificationType) {
	o.Verification = v
}

func (o JettonPreview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JettonPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["name"] = o.Name
	toSerialize["symbol"] = o.Symbol
	toSerialize["decimals"] = o.Decimals
	toSerialize["image"] = o.Image
	toSerialize["verification"] = o.Verification
	return toSerialize, nil
}

func (o *JettonPreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"name",
		"symbol",
		"decimals",
		"image",
		"verification",
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

	varJettonPreview := _JettonPreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJettonPreview)

	if err != nil {
		return err
	}

	*o = JettonPreview(varJettonPreview)

	return err
}

type NullableJettonPreview struct {
	value *JettonPreview
	isSet bool
}

func (v NullableJettonPreview) Get() *JettonPreview {
	return v.value
}

func (v *NullableJettonPreview) Set(val *JettonPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableJettonPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableJettonPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJettonPreview(val *JettonPreview) *NullableJettonPreview {
	return &NullableJettonPreview{value: val, isSet: true}
}

func (v NullableJettonPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJettonPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


