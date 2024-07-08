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

// checks if the AccountAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAddress{}

// AccountAddress struct for AccountAddress
type AccountAddress struct {
	Address string `json:"address"`
	// Display name. Data collected from different sources like moderation lists, dns, collections names and over.
	Name *string `json:"name,omitempty"`
	// Is this account was marked as part of scammers activity
	IsScam bool `json:"is_scam"`
	Icon *string `json:"icon,omitempty"`
	IsWallet bool `json:"is_wallet"`
}

type _AccountAddress AccountAddress

// NewAccountAddress instantiates a new AccountAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAddress(address string, isScam bool, isWallet bool) *AccountAddress {
	this := AccountAddress{}
	this.Address = address
	this.IsScam = isScam
	this.IsWallet = isWallet
	return &this
}

// NewAccountAddressWithDefaults instantiates a new AccountAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAddressWithDefaults() *AccountAddress {
	this := AccountAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *AccountAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *AccountAddress) SetAddress(v string) {
	o.Address = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountAddress) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountAddress) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountAddress) SetName(v string) {
	o.Name = &v
}

// GetIsScam returns the IsScam field value
func (o *AccountAddress) GetIsScam() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsScam
}

// GetIsScamOk returns a tuple with the IsScam field value
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetIsScamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsScam, true
}

// SetIsScam sets field value
func (o *AccountAddress) SetIsScam(v bool) {
	o.IsScam = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *AccountAddress) GetIcon() string {
	if o == nil || IsNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetIconOk() (*string, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *AccountAddress) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *AccountAddress) SetIcon(v string) {
	o.Icon = &v
}

// GetIsWallet returns the IsWallet field value
func (o *AccountAddress) GetIsWallet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWallet
}

// GetIsWalletOk returns a tuple with the IsWallet field value
// and a boolean to check if the value has been set.
func (o *AccountAddress) GetIsWalletOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWallet, true
}

// SetIsWallet sets field value
func (o *AccountAddress) SetIsWallet(v bool) {
	o.IsWallet = v
}

func (o AccountAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["is_scam"] = o.IsScam
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	toSerialize["is_wallet"] = o.IsWallet
	return toSerialize, nil
}

func (o *AccountAddress) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"is_scam",
		"is_wallet",
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

	varAccountAddress := _AccountAddress{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountAddress)

	if err != nil {
		return err
	}

	*o = AccountAddress(varAccountAddress)

	return err
}

type NullableAccountAddress struct {
	value *AccountAddress
	isSet bool
}

func (v NullableAccountAddress) Get() *AccountAddress {
	return v.value
}

func (v *NullableAccountAddress) Set(val *AccountAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAddress(val *AccountAddress) *NullableAccountAddress {
	return &NullableAccountAddress{value: val, isSet: true}
}

func (v NullableAccountAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


