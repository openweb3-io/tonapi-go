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

// checks if the GetRawTransactions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRawTransactions200Response{}

// GetRawTransactions200Response struct for GetRawTransactions200Response
type GetRawTransactions200Response struct {
	Ids []BlockRaw `json:"ids"`
	Transactions string `json:"transactions"`
}

type _GetRawTransactions200Response GetRawTransactions200Response

// NewGetRawTransactions200Response instantiates a new GetRawTransactions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRawTransactions200Response(ids []BlockRaw, transactions string) *GetRawTransactions200Response {
	this := GetRawTransactions200Response{}
	this.Ids = ids
	this.Transactions = transactions
	return &this
}

// NewGetRawTransactions200ResponseWithDefaults instantiates a new GetRawTransactions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRawTransactions200ResponseWithDefaults() *GetRawTransactions200Response {
	this := GetRawTransactions200Response{}
	return &this
}

// GetIds returns the Ids field value
func (o *GetRawTransactions200Response) GetIds() []BlockRaw {
	if o == nil {
		var ret []BlockRaw
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *GetRawTransactions200Response) GetIdsOk() ([]BlockRaw, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ids, true
}

// SetIds sets field value
func (o *GetRawTransactions200Response) SetIds(v []BlockRaw) {
	o.Ids = v
}

// GetTransactions returns the Transactions field value
func (o *GetRawTransactions200Response) GetTransactions() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Transactions
}

// GetTransactionsOk returns a tuple with the Transactions field value
// and a boolean to check if the value has been set.
func (o *GetRawTransactions200Response) GetTransactionsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transactions, true
}

// SetTransactions sets field value
func (o *GetRawTransactions200Response) SetTransactions(v string) {
	o.Transactions = v
}

func (o GetRawTransactions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRawTransactions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ids"] = o.Ids
	toSerialize["transactions"] = o.Transactions
	return toSerialize, nil
}

func (o *GetRawTransactions200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ids",
		"transactions",
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

	varGetRawTransactions200Response := _GetRawTransactions200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRawTransactions200Response)

	if err != nil {
		return err
	}

	*o = GetRawTransactions200Response(varGetRawTransactions200Response)

	return err
}

type NullableGetRawTransactions200Response struct {
	value *GetRawTransactions200Response
	isSet bool
}

func (v NullableGetRawTransactions200Response) Get() *GetRawTransactions200Response {
	return v.value
}

func (v *NullableGetRawTransactions200Response) Set(val *GetRawTransactions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRawTransactions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRawTransactions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRawTransactions200Response(val *GetRawTransactions200Response) *NullableGetRawTransactions200Response {
	return &NullableGetRawTransactions200Response{value: val, isSet: true}
}

func (v NullableGetRawTransactions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRawTransactions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

