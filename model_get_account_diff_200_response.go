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

// checks if the GetAccountDiff200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountDiff200Response{}

// GetAccountDiff200Response struct for GetAccountDiff200Response
type GetAccountDiff200Response struct {
	BalanceChange int64 `json:"balance_change"`
}

type _GetAccountDiff200Response GetAccountDiff200Response

// NewGetAccountDiff200Response instantiates a new GetAccountDiff200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountDiff200Response(balanceChange int64) *GetAccountDiff200Response {
	this := GetAccountDiff200Response{}
	this.BalanceChange = balanceChange
	return &this
}

// NewGetAccountDiff200ResponseWithDefaults instantiates a new GetAccountDiff200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountDiff200ResponseWithDefaults() *GetAccountDiff200Response {
	this := GetAccountDiff200Response{}
	return &this
}

// GetBalanceChange returns the BalanceChange field value
func (o *GetAccountDiff200Response) GetBalanceChange() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BalanceChange
}

// GetBalanceChangeOk returns a tuple with the BalanceChange field value
// and a boolean to check if the value has been set.
func (o *GetAccountDiff200Response) GetBalanceChangeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BalanceChange, true
}

// SetBalanceChange sets field value
func (o *GetAccountDiff200Response) SetBalanceChange(v int64) {
	o.BalanceChange = v
}

func (o GetAccountDiff200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountDiff200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["balance_change"] = o.BalanceChange
	return toSerialize, nil
}

func (o *GetAccountDiff200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"balance_change",
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

	varGetAccountDiff200Response := _GetAccountDiff200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAccountDiff200Response)

	if err != nil {
		return err
	}

	*o = GetAccountDiff200Response(varGetAccountDiff200Response)

	return err
}

type NullableGetAccountDiff200Response struct {
	value *GetAccountDiff200Response
	isSet bool
}

func (v NullableGetAccountDiff200Response) Get() *GetAccountDiff200Response {
	return v.value
}

func (v *NullableGetAccountDiff200Response) Set(val *GetAccountDiff200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountDiff200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountDiff200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountDiff200Response(val *GetAccountDiff200Response) *NullableGetAccountDiff200Response {
	return &NullableGetAccountDiff200Response{value: val, isSet: true}
}

func (v NullableGetAccountDiff200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountDiff200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


