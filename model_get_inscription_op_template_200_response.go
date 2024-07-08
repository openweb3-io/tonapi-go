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

// checks if the GetInscriptionOpTemplate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInscriptionOpTemplate200Response{}

// GetInscriptionOpTemplate200Response struct for GetInscriptionOpTemplate200Response
type GetInscriptionOpTemplate200Response struct {
	Comment string `json:"comment"`
	Destination string `json:"destination"`
}

type _GetInscriptionOpTemplate200Response GetInscriptionOpTemplate200Response

// NewGetInscriptionOpTemplate200Response instantiates a new GetInscriptionOpTemplate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInscriptionOpTemplate200Response(comment string, destination string) *GetInscriptionOpTemplate200Response {
	this := GetInscriptionOpTemplate200Response{}
	this.Comment = comment
	this.Destination = destination
	return &this
}

// NewGetInscriptionOpTemplate200ResponseWithDefaults instantiates a new GetInscriptionOpTemplate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInscriptionOpTemplate200ResponseWithDefaults() *GetInscriptionOpTemplate200Response {
	this := GetInscriptionOpTemplate200Response{}
	return &this
}

// GetComment returns the Comment field value
func (o *GetInscriptionOpTemplate200Response) GetComment() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *GetInscriptionOpTemplate200Response) GetCommentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comment, true
}

// SetComment sets field value
func (o *GetInscriptionOpTemplate200Response) SetComment(v string) {
	o.Comment = v
}

// GetDestination returns the Destination field value
func (o *GetInscriptionOpTemplate200Response) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *GetInscriptionOpTemplate200Response) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *GetInscriptionOpTemplate200Response) SetDestination(v string) {
	o.Destination = v
}

func (o GetInscriptionOpTemplate200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInscriptionOpTemplate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comment"] = o.Comment
	toSerialize["destination"] = o.Destination
	return toSerialize, nil
}

func (o *GetInscriptionOpTemplate200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"comment",
		"destination",
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

	varGetInscriptionOpTemplate200Response := _GetInscriptionOpTemplate200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInscriptionOpTemplate200Response)

	if err != nil {
		return err
	}

	*o = GetInscriptionOpTemplate200Response(varGetInscriptionOpTemplate200Response)

	return err
}

type NullableGetInscriptionOpTemplate200Response struct {
	value *GetInscriptionOpTemplate200Response
	isSet bool
}

func (v NullableGetInscriptionOpTemplate200Response) Get() *GetInscriptionOpTemplate200Response {
	return v.value
}

func (v *NullableGetInscriptionOpTemplate200Response) Set(val *GetInscriptionOpTemplate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInscriptionOpTemplate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInscriptionOpTemplate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInscriptionOpTemplate200Response(val *GetInscriptionOpTemplate200Response) *NullableGetInscriptionOpTemplate200Response {
	return &NullableGetInscriptionOpTemplate200Response{value: val, isSet: true}
}

func (v NullableGetInscriptionOpTemplate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInscriptionOpTemplate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


