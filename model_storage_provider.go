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

// checks if the StorageProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageProvider{}

// StorageProvider struct for StorageProvider
type StorageProvider struct {
	Address string `json:"address"`
	AcceptNewContracts bool `json:"accept_new_contracts"`
	RatePerMbDay int64 `json:"rate_per_mb_day"`
	MaxSpan int64 `json:"max_span"`
	MinimalFileSize int64 `json:"minimal_file_size"`
	MaximalFileSize int64 `json:"maximal_file_size"`
}

type _StorageProvider StorageProvider

// NewStorageProvider instantiates a new StorageProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageProvider(address string, acceptNewContracts bool, ratePerMbDay int64, maxSpan int64, minimalFileSize int64, maximalFileSize int64) *StorageProvider {
	this := StorageProvider{}
	this.Address = address
	this.AcceptNewContracts = acceptNewContracts
	this.RatePerMbDay = ratePerMbDay
	this.MaxSpan = maxSpan
	this.MinimalFileSize = minimalFileSize
	this.MaximalFileSize = maximalFileSize
	return &this
}

// NewStorageProviderWithDefaults instantiates a new StorageProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageProviderWithDefaults() *StorageProvider {
	this := StorageProvider{}
	return &this
}

// GetAddress returns the Address field value
func (o *StorageProvider) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *StorageProvider) SetAddress(v string) {
	o.Address = v
}

// GetAcceptNewContracts returns the AcceptNewContracts field value
func (o *StorageProvider) GetAcceptNewContracts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AcceptNewContracts
}

// GetAcceptNewContractsOk returns a tuple with the AcceptNewContracts field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetAcceptNewContractsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcceptNewContracts, true
}

// SetAcceptNewContracts sets field value
func (o *StorageProvider) SetAcceptNewContracts(v bool) {
	o.AcceptNewContracts = v
}

// GetRatePerMbDay returns the RatePerMbDay field value
func (o *StorageProvider) GetRatePerMbDay() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.RatePerMbDay
}

// GetRatePerMbDayOk returns a tuple with the RatePerMbDay field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetRatePerMbDayOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatePerMbDay, true
}

// SetRatePerMbDay sets field value
func (o *StorageProvider) SetRatePerMbDay(v int64) {
	o.RatePerMbDay = v
}

// GetMaxSpan returns the MaxSpan field value
func (o *StorageProvider) GetMaxSpan() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaxSpan
}

// GetMaxSpanOk returns a tuple with the MaxSpan field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetMaxSpanOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSpan, true
}

// SetMaxSpan sets field value
func (o *StorageProvider) SetMaxSpan(v int64) {
	o.MaxSpan = v
}

// GetMinimalFileSize returns the MinimalFileSize field value
func (o *StorageProvider) GetMinimalFileSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MinimalFileSize
}

// GetMinimalFileSizeOk returns a tuple with the MinimalFileSize field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetMinimalFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinimalFileSize, true
}

// SetMinimalFileSize sets field value
func (o *StorageProvider) SetMinimalFileSize(v int64) {
	o.MinimalFileSize = v
}

// GetMaximalFileSize returns the MaximalFileSize field value
func (o *StorageProvider) GetMaximalFileSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MaximalFileSize
}

// GetMaximalFileSizeOk returns a tuple with the MaximalFileSize field value
// and a boolean to check if the value has been set.
func (o *StorageProvider) GetMaximalFileSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaximalFileSize, true
}

// SetMaximalFileSize sets field value
func (o *StorageProvider) SetMaximalFileSize(v int64) {
	o.MaximalFileSize = v
}

func (o StorageProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["accept_new_contracts"] = o.AcceptNewContracts
	toSerialize["rate_per_mb_day"] = o.RatePerMbDay
	toSerialize["max_span"] = o.MaxSpan
	toSerialize["minimal_file_size"] = o.MinimalFileSize
	toSerialize["maximal_file_size"] = o.MaximalFileSize
	return toSerialize, nil
}

func (o *StorageProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"accept_new_contracts",
		"rate_per_mb_day",
		"max_span",
		"minimal_file_size",
		"maximal_file_size",
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

	varStorageProvider := _StorageProvider{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageProvider)

	if err != nil {
		return err
	}

	*o = StorageProvider(varStorageProvider)

	return err
}

type NullableStorageProvider struct {
	value *StorageProvider
	isSet bool
}

func (v NullableStorageProvider) Get() *StorageProvider {
	return v.value
}

func (v *NullableStorageProvider) Set(val *StorageProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageProvider(val *StorageProvider) *NullableStorageProvider {
	return &NullableStorageProvider{value: val, isSet: true}
}

func (v NullableStorageProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

