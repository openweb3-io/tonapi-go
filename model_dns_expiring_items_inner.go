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

// checks if the DnsExpiringItemsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsExpiringItemsInner{}

// DnsExpiringItemsInner struct for DnsExpiringItemsInner
type DnsExpiringItemsInner struct {
	ExpiringAt int64 `json:"expiring_at"`
	Name string `json:"name"`
	DnsItem *NftItem `json:"dns_item,omitempty"`
}

type _DnsExpiringItemsInner DnsExpiringItemsInner

// NewDnsExpiringItemsInner instantiates a new DnsExpiringItemsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsExpiringItemsInner(expiringAt int64, name string) *DnsExpiringItemsInner {
	this := DnsExpiringItemsInner{}
	this.ExpiringAt = expiringAt
	this.Name = name
	return &this
}

// NewDnsExpiringItemsInnerWithDefaults instantiates a new DnsExpiringItemsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsExpiringItemsInnerWithDefaults() *DnsExpiringItemsInner {
	this := DnsExpiringItemsInner{}
	return &this
}

// GetExpiringAt returns the ExpiringAt field value
func (o *DnsExpiringItemsInner) GetExpiringAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.ExpiringAt
}

// GetExpiringAtOk returns a tuple with the ExpiringAt field value
// and a boolean to check if the value has been set.
func (o *DnsExpiringItemsInner) GetExpiringAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiringAt, true
}

// SetExpiringAt sets field value
func (o *DnsExpiringItemsInner) SetExpiringAt(v int64) {
	o.ExpiringAt = v
}

// GetName returns the Name field value
func (o *DnsExpiringItemsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DnsExpiringItemsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DnsExpiringItemsInner) SetName(v string) {
	o.Name = v
}

// GetDnsItem returns the DnsItem field value if set, zero value otherwise.
func (o *DnsExpiringItemsInner) GetDnsItem() NftItem {
	if o == nil || IsNil(o.DnsItem) {
		var ret NftItem
		return ret
	}
	return *o.DnsItem
}

// GetDnsItemOk returns a tuple with the DnsItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsExpiringItemsInner) GetDnsItemOk() (*NftItem, bool) {
	if o == nil || IsNil(o.DnsItem) {
		return nil, false
	}
	return o.DnsItem, true
}

// HasDnsItem returns a boolean if a field has been set.
func (o *DnsExpiringItemsInner) HasDnsItem() bool {
	if o != nil && !IsNil(o.DnsItem) {
		return true
	}

	return false
}

// SetDnsItem gets a reference to the given NftItem and assigns it to the DnsItem field.
func (o *DnsExpiringItemsInner) SetDnsItem(v NftItem) {
	o.DnsItem = &v
}

func (o DnsExpiringItemsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsExpiringItemsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expiring_at"] = o.ExpiringAt
	toSerialize["name"] = o.Name
	if !IsNil(o.DnsItem) {
		toSerialize["dns_item"] = o.DnsItem
	}
	return toSerialize, nil
}

func (o *DnsExpiringItemsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expiring_at",
		"name",
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

	varDnsExpiringItemsInner := _DnsExpiringItemsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDnsExpiringItemsInner)

	if err != nil {
		return err
	}

	*o = DnsExpiringItemsInner(varDnsExpiringItemsInner)

	return err
}

type NullableDnsExpiringItemsInner struct {
	value *DnsExpiringItemsInner
	isSet bool
}

func (v NullableDnsExpiringItemsInner) Get() *DnsExpiringItemsInner {
	return v.value
}

func (v *NullableDnsExpiringItemsInner) Set(val *DnsExpiringItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsExpiringItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsExpiringItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsExpiringItemsInner(val *DnsExpiringItemsInner) *NullableDnsExpiringItemsInner {
	return &NullableDnsExpiringItemsInner{value: val, isSet: true}
}

func (v NullableDnsExpiringItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsExpiringItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

