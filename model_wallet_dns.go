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

// checks if the WalletDNS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletDNS{}

// WalletDNS struct for WalletDNS
type WalletDNS struct {
	Address string `json:"address"`
	Account AccountAddress `json:"account"`
	IsWallet bool `json:"is_wallet"`
	HasMethodPubkey bool `json:"has_method_pubkey"`
	HasMethodSeqno bool `json:"has_method_seqno"`
	Names []string `json:"names"`
}

type _WalletDNS WalletDNS

// NewWalletDNS instantiates a new WalletDNS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletDNS(address string, account AccountAddress, isWallet bool, hasMethodPubkey bool, hasMethodSeqno bool, names []string) *WalletDNS {
	this := WalletDNS{}
	this.Address = address
	this.Account = account
	this.IsWallet = isWallet
	this.HasMethodPubkey = hasMethodPubkey
	this.HasMethodSeqno = hasMethodSeqno
	this.Names = names
	return &this
}

// NewWalletDNSWithDefaults instantiates a new WalletDNS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletDNSWithDefaults() *WalletDNS {
	this := WalletDNS{}
	return &this
}

// GetAddress returns the Address field value
func (o *WalletDNS) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *WalletDNS) SetAddress(v string) {
	o.Address = v
}

// GetAccount returns the Account field value
func (o *WalletDNS) GetAccount() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetAccountOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *WalletDNS) SetAccount(v AccountAddress) {
	o.Account = v
}

// GetIsWallet returns the IsWallet field value
func (o *WalletDNS) GetIsWallet() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsWallet
}

// GetIsWalletOk returns a tuple with the IsWallet field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetIsWalletOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsWallet, true
}

// SetIsWallet sets field value
func (o *WalletDNS) SetIsWallet(v bool) {
	o.IsWallet = v
}

// GetHasMethodPubkey returns the HasMethodPubkey field value
func (o *WalletDNS) GetHasMethodPubkey() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMethodPubkey
}

// GetHasMethodPubkeyOk returns a tuple with the HasMethodPubkey field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetHasMethodPubkeyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMethodPubkey, true
}

// SetHasMethodPubkey sets field value
func (o *WalletDNS) SetHasMethodPubkey(v bool) {
	o.HasMethodPubkey = v
}

// GetHasMethodSeqno returns the HasMethodSeqno field value
func (o *WalletDNS) GetHasMethodSeqno() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMethodSeqno
}

// GetHasMethodSeqnoOk returns a tuple with the HasMethodSeqno field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetHasMethodSeqnoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMethodSeqno, true
}

// SetHasMethodSeqno sets field value
func (o *WalletDNS) SetHasMethodSeqno(v bool) {
	o.HasMethodSeqno = v
}

// GetNames returns the Names field value
func (o *WalletDNS) GetNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *WalletDNS) GetNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Names, true
}

// SetNames sets field value
func (o *WalletDNS) SetNames(v []string) {
	o.Names = v
}

func (o WalletDNS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletDNS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["account"] = o.Account
	toSerialize["is_wallet"] = o.IsWallet
	toSerialize["has_method_pubkey"] = o.HasMethodPubkey
	toSerialize["has_method_seqno"] = o.HasMethodSeqno
	toSerialize["names"] = o.Names
	return toSerialize, nil
}

func (o *WalletDNS) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address",
		"account",
		"is_wallet",
		"has_method_pubkey",
		"has_method_seqno",
		"names",
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

	varWalletDNS := _WalletDNS{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWalletDNS)

	if err != nil {
		return err
	}

	*o = WalletDNS(varWalletDNS)

	return err
}

type NullableWalletDNS struct {
	value *WalletDNS
	isSet bool
}

func (v NullableWalletDNS) Get() *WalletDNS {
	return v.value
}

func (v *NullableWalletDNS) Set(val *WalletDNS) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletDNS) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletDNS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletDNS(val *WalletDNS) *NullableWalletDNS {
	return &NullableWalletDNS{value: val, isSet: true}
}

func (v NullableWalletDNS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletDNS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


