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

// checks if the GetAccountPublicKey200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccountPublicKey200Response{}

// GetAccountPublicKey200Response struct for GetAccountPublicKey200Response
type GetAccountPublicKey200Response struct {
	PublicKey string `json:"public_key"`
}

type _GetAccountPublicKey200Response GetAccountPublicKey200Response

// NewGetAccountPublicKey200Response instantiates a new GetAccountPublicKey200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccountPublicKey200Response(publicKey string) *GetAccountPublicKey200Response {
	this := GetAccountPublicKey200Response{}
	this.PublicKey = publicKey
	return &this
}

// NewGetAccountPublicKey200ResponseWithDefaults instantiates a new GetAccountPublicKey200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccountPublicKey200ResponseWithDefaults() *GetAccountPublicKey200Response {
	this := GetAccountPublicKey200Response{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *GetAccountPublicKey200Response) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *GetAccountPublicKey200Response) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *GetAccountPublicKey200Response) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o GetAccountPublicKey200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccountPublicKey200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["public_key"] = o.PublicKey
	return toSerialize, nil
}

func (o *GetAccountPublicKey200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"public_key",
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

	varGetAccountPublicKey200Response := _GetAccountPublicKey200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetAccountPublicKey200Response)

	if err != nil {
		return err
	}

	*o = GetAccountPublicKey200Response(varGetAccountPublicKey200Response)

	return err
}

type NullableGetAccountPublicKey200Response struct {
	value *GetAccountPublicKey200Response
	isSet bool
}

func (v NullableGetAccountPublicKey200Response) Get() *GetAccountPublicKey200Response {
	return v.value
}

func (v *NullableGetAccountPublicKey200Response) Set(val *GetAccountPublicKey200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccountPublicKey200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccountPublicKey200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccountPublicKey200Response(val *GetAccountPublicKey200Response) *NullableGetAccountPublicKey200Response {
	return &NullableGetAccountPublicKey200Response{value: val, isSet: true}
}

func (v NullableGetAccountPublicKey200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccountPublicKey200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


