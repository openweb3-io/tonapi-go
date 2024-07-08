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

// checks if the DomainRenewAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DomainRenewAction{}

// DomainRenewAction struct for DomainRenewAction
type DomainRenewAction struct {
	Domain string `json:"domain"`
	ContractAddress string `json:"contract_address"`
	Renewer AccountAddress `json:"renewer"`
}

type _DomainRenewAction DomainRenewAction

// NewDomainRenewAction instantiates a new DomainRenewAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainRenewAction(domain string, contractAddress string, renewer AccountAddress) *DomainRenewAction {
	this := DomainRenewAction{}
	this.Domain = domain
	this.ContractAddress = contractAddress
	this.Renewer = renewer
	return &this
}

// NewDomainRenewActionWithDefaults instantiates a new DomainRenewAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainRenewActionWithDefaults() *DomainRenewAction {
	this := DomainRenewAction{}
	return &this
}

// GetDomain returns the Domain field value
func (o *DomainRenewAction) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainRenewAction) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainRenewAction) SetDomain(v string) {
	o.Domain = v
}

// GetContractAddress returns the ContractAddress field value
func (o *DomainRenewAction) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *DomainRenewAction) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *DomainRenewAction) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetRenewer returns the Renewer field value
func (o *DomainRenewAction) GetRenewer() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Renewer
}

// GetRenewerOk returns a tuple with the Renewer field value
// and a boolean to check if the value has been set.
func (o *DomainRenewAction) GetRenewerOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Renewer, true
}

// SetRenewer sets field value
func (o *DomainRenewAction) SetRenewer(v AccountAddress) {
	o.Renewer = v
}

func (o DomainRenewAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DomainRenewAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["renewer"] = o.Renewer
	return toSerialize, nil
}

func (o *DomainRenewAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"contract_address",
		"renewer",
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

	varDomainRenewAction := _DomainRenewAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDomainRenewAction)

	if err != nil {
		return err
	}

	*o = DomainRenewAction(varDomainRenewAction)

	return err
}

type NullableDomainRenewAction struct {
	value *DomainRenewAction
	isSet bool
}

func (v NullableDomainRenewAction) Get() *DomainRenewAction {
	return v.value
}

func (v *NullableDomainRenewAction) Set(val *DomainRenewAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainRenewAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainRenewAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainRenewAction(val *DomainRenewAction) *NullableDomainRenewAction {
	return &NullableDomainRenewAction{value: val, isSet: true}
}

func (v NullableDomainRenewAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainRenewAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


