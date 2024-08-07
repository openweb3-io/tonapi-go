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

// checks if the GetRawShardBlockProof200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawShardBlockProof200Response{}

// GetRawShardBlockProof200Response struct for GetRawShardBlockProof200Response
type GetRawShardBlockProof200Response struct {
	MasterchainId BlockRaw `json:"masterchain_id"`
	Links []GetRawShardBlockProof200ResponseLinksInner `json:"links"`
}

type _GetRawShardBlockProof200Response GetRawShardBlockProof200Response

// NewGetRawShardBlockProof200Response instantiates a new GetRawShardBlockProof200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawShardBlockProof200Response(masterchainId BlockRaw, links []GetRawShardBlockProof200ResponseLinksInner) *GetRawShardBlockProof200Response {
	this := GetRawShardBlockProof200Response{}
	this.MasterchainId = masterchainId
	this.Links = links
	return &this
}

// NewGetRawShardBlockProof200ResponseWithDefaults instantiates a new GetRawShardBlockProof200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawShardBlockProof200ResponseWithDefaults() *GetRawShardBlockProof200Response {
	this := GetRawShardBlockProof200Response{}
	return &this
}

// GetMasterchainId returns the MasterchainId field value
func (o *GetRawShardBlockProof200Response) GetMasterchainId() BlockRaw {
	if o == nil {
		var ret BlockRaw
		return ret
	}

	return o.MasterchainId
}

// GetMasterchainIdOk returns a tuple with the MasterchainId field value
// and a boolean to check if the value has been set.
func (o *GetRawShardBlockProof200Response) GetMasterchainIdOk() (*BlockRaw, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MasterchainId, true
}

// SetMasterchainId sets field value
func (o *GetRawShardBlockProof200Response) SetMasterchainId(v BlockRaw) {
	o.MasterchainId = v
}

// GetLinks returns the Links field value
func (o *GetRawShardBlockProof200Response) GetLinks() []GetRawShardBlockProof200ResponseLinksInner {
	if o == nil {
		var ret []GetRawShardBlockProof200ResponseLinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetRawShardBlockProof200Response) GetLinksOk() ([]GetRawShardBlockProof200ResponseLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *GetRawShardBlockProof200Response) SetLinks(v []GetRawShardBlockProof200ResponseLinksInner) {
	o.Links = v
}

func (o GetRawShardBlockProof200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawShardBlockProof200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["masterchain_id"] = o.MasterchainId
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

func (o *GetRawShardBlockProof200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"masterchain_id",
		"links",
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

	varGetRawShardBlockProof200Response := _GetRawShardBlockProof200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawShardBlockProof200Response)

	if err != nil {
		return err
	}

	*o = GetRawShardBlockProof200Response(varGetRawShardBlockProof200Response)

	return err
}

type NullableGetRawShardBlockProof200Response struct {
	value *GetRawShardBlockProof200Response
	isSet bool
}

func (v NullableGetRawShardBlockProof200Response) Get() *GetRawShardBlockProof200Response {
	return v.value
}

func (v *NullableGetRawShardBlockProof200Response) Set(val *GetRawShardBlockProof200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawShardBlockProof200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawShardBlockProof200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawShardBlockProof200Response(val *GetRawShardBlockProof200Response) *NullableGetRawShardBlockProof200Response {
	return &NullableGetRawShardBlockProof200Response{value: val, isSet: true}
}

func (v NullableGetRawShardBlockProof200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawShardBlockProof200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


