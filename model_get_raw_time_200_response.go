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

// checks if the GetRawTime200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawTime200Response{}

// GetRawTime200Response struct for GetRawTime200Response
type GetRawTime200Response struct {
	Time int32 `json:"time"`
}

type _GetRawTime200Response GetRawTime200Response

// NewGetRawTime200Response instantiates a new GetRawTime200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawTime200Response(time int32) *GetRawTime200Response {
	this := GetRawTime200Response{}
	this.Time = time
	return &this
}

// NewGetRawTime200ResponseWithDefaults instantiates a new GetRawTime200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawTime200ResponseWithDefaults() *GetRawTime200Response {
	this := GetRawTime200Response{}
	return &this
}

// GetTime returns the Time field value
func (o *GetRawTime200Response) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *GetRawTime200Response) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *GetRawTime200Response) SetTime(v int32) {
	o.Time = v
}

func (o GetRawTime200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawTime200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["time"] = o.Time
	return toSerialize, nil
}

func (o *GetRawTime200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"time",
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

	varGetRawTime200Response := _GetRawTime200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawTime200Response)

	if err != nil {
		return err
	}

	*o = GetRawTime200Response(varGetRawTime200Response)

	return err
}

type NullableGetRawTime200Response struct {
	value *GetRawTime200Response
	isSet bool
}

func (v NullableGetRawTime200Response) Get() *GetRawTime200Response {
	return v.value
}

func (v *NullableGetRawTime200Response) Set(val *GetRawTime200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawTime200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawTime200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawTime200Response(val *GetRawTime200Response) *NullableGetRawTime200Response {
	return &NullableGetRawTime200Response{value: val, isSet: true}
}

func (v NullableGetRawTime200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawTime200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


