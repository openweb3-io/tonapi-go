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

// checks if the GetStorageProviders200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStorageProviders200Response{}

// GetStorageProviders200Response struct for GetStorageProviders200Response
type GetStorageProviders200Response struct {
	Providers []StorageProvider `json:"providers"`
}

type _GetStorageProviders200Response GetStorageProviders200Response

// NewGetStorageProviders200Response instantiates a new GetStorageProviders200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStorageProviders200Response(providers []StorageProvider) *GetStorageProviders200Response {
	this := GetStorageProviders200Response{}
	this.Providers = providers
	return &this
}

// NewGetStorageProviders200ResponseWithDefaults instantiates a new GetStorageProviders200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStorageProviders200ResponseWithDefaults() *GetStorageProviders200Response {
	this := GetStorageProviders200Response{}
	return &this
}

// GetProviders returns the Providers field value
func (o *GetStorageProviders200Response) GetProviders() []StorageProvider {
	if o == nil {
		var ret []StorageProvider
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
func (o *GetStorageProviders200Response) GetProvidersOk() ([]StorageProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *GetStorageProviders200Response) SetProviders(v []StorageProvider) {
	o.Providers = v
}

func (o GetStorageProviders200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStorageProviders200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["providers"] = o.Providers
	return toSerialize, nil
}

func (o *GetStorageProviders200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"providers",
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

	varGetStorageProviders200Response := _GetStorageProviders200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStorageProviders200Response)

	if err != nil {
		return err
	}

	*o = GetStorageProviders200Response(varGetStorageProviders200Response)

	return err
}

type NullableGetStorageProviders200Response struct {
	value *GetStorageProviders200Response
	isSet bool
}

func (v NullableGetStorageProviders200Response) Get() *GetStorageProviders200Response {
	return v.value
}

func (v *NullableGetStorageProviders200Response) Set(val *GetStorageProviders200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStorageProviders200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStorageProviders200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStorageProviders200Response(val *GetStorageProviders200Response) *NullableGetStorageProviders200Response {
	return &NullableGetStorageProviders200Response{value: val, isSet: true}
}

func (v NullableGetStorageProviders200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStorageProviders200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

