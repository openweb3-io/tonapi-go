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

// checks if the TokenRates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRates{}

// TokenRates struct for TokenRates
type TokenRates struct {
	Prices *map[string]float32 `json:"prices,omitempty"`
	Diff24h *map[string]string `json:"diff_24h,omitempty"`
	Diff7d *map[string]string `json:"diff_7d,omitempty"`
	Diff30d *map[string]string `json:"diff_30d,omitempty"`
}

// NewTokenRates instantiates a new TokenRates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRates() *TokenRates {
	this := TokenRates{}
	return &this
}

// NewTokenRatesWithDefaults instantiates a new TokenRates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRatesWithDefaults() *TokenRates {
	this := TokenRates{}
	return &this
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *TokenRates) GetPrices() map[string]float32 {
	if o == nil || IsNil(o.Prices) {
		var ret map[string]float32
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRates) GetPricesOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.Prices) {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *TokenRates) HasPrices() bool {
	if o != nil && !IsNil(o.Prices) {
		return true
	}

	return false
}

// SetPrices gets a reference to the given map[string]float32 and assigns it to the Prices field.
func (o *TokenRates) SetPrices(v map[string]float32) {
	o.Prices = &v
}

// GetDiff24h returns the Diff24h field value if set, zero value otherwise.
func (o *TokenRates) GetDiff24h() map[string]string {
	if o == nil || IsNil(o.Diff24h) {
		var ret map[string]string
		return ret
	}
	return *o.Diff24h
}

// GetDiff24hOk returns a tuple with the Diff24h field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRates) GetDiff24hOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Diff24h) {
		return nil, false
	}
	return o.Diff24h, true
}

// HasDiff24h returns a boolean if a field has been set.
func (o *TokenRates) HasDiff24h() bool {
	if o != nil && !IsNil(o.Diff24h) {
		return true
	}

	return false
}

// SetDiff24h gets a reference to the given map[string]string and assigns it to the Diff24h field.
func (o *TokenRates) SetDiff24h(v map[string]string) {
	o.Diff24h = &v
}

// GetDiff7d returns the Diff7d field value if set, zero value otherwise.
func (o *TokenRates) GetDiff7d() map[string]string {
	if o == nil || IsNil(o.Diff7d) {
		var ret map[string]string
		return ret
	}
	return *o.Diff7d
}

// GetDiff7dOk returns a tuple with the Diff7d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRates) GetDiff7dOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Diff7d) {
		return nil, false
	}
	return o.Diff7d, true
}

// HasDiff7d returns a boolean if a field has been set.
func (o *TokenRates) HasDiff7d() bool {
	if o != nil && !IsNil(o.Diff7d) {
		return true
	}

	return false
}

// SetDiff7d gets a reference to the given map[string]string and assigns it to the Diff7d field.
func (o *TokenRates) SetDiff7d(v map[string]string) {
	o.Diff7d = &v
}

// GetDiff30d returns the Diff30d field value if set, zero value otherwise.
func (o *TokenRates) GetDiff30d() map[string]string {
	if o == nil || IsNil(o.Diff30d) {
		var ret map[string]string
		return ret
	}
	return *o.Diff30d
}

// GetDiff30dOk returns a tuple with the Diff30d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRates) GetDiff30dOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Diff30d) {
		return nil, false
	}
	return o.Diff30d, true
}

// HasDiff30d returns a boolean if a field has been set.
func (o *TokenRates) HasDiff30d() bool {
	if o != nil && !IsNil(o.Diff30d) {
		return true
	}

	return false
}

// SetDiff30d gets a reference to the given map[string]string and assigns it to the Diff30d field.
func (o *TokenRates) SetDiff30d(v map[string]string) {
	o.Diff30d = &v
}

func (o TokenRates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Prices) {
		toSerialize["prices"] = o.Prices
	}
	if !IsNil(o.Diff24h) {
		toSerialize["diff_24h"] = o.Diff24h
	}
	if !IsNil(o.Diff7d) {
		toSerialize["diff_7d"] = o.Diff7d
	}
	if !IsNil(o.Diff30d) {
		toSerialize["diff_30d"] = o.Diff30d
	}
	return toSerialize, nil
}

type NullableTokenRates struct {
	value *TokenRates
	isSet bool
}

func (v NullableTokenRates) Get() *TokenRates {
	return v.value
}

func (v *NullableTokenRates) Set(val *TokenRates) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRates) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRates(val *TokenRates) *NullableTokenRates {
	return &NullableTokenRates{value: val, isSet: true}
}

func (v NullableTokenRates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

