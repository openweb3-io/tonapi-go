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

// checks if the InscriptionMintAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InscriptionMintAction{}

// InscriptionMintAction struct for InscriptionMintAction
type InscriptionMintAction struct {
	Recipient AccountAddress `json:"recipient"`
	// amount in minimal particles
	Amount string `json:"amount"`
	Type string `json:"type"`
	Ticker string `json:"ticker"`
	Decimals int32 `json:"decimals"`
}

type _InscriptionMintAction InscriptionMintAction

// NewInscriptionMintAction instantiates a new InscriptionMintAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInscriptionMintAction(recipient AccountAddress, amount string, type_ string, ticker string, decimals int32) *InscriptionMintAction {
	this := InscriptionMintAction{}
	this.Recipient = recipient
	this.Amount = amount
	this.Type = type_
	this.Ticker = ticker
	this.Decimals = decimals
	return &this
}

// NewInscriptionMintActionWithDefaults instantiates a new InscriptionMintAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInscriptionMintActionWithDefaults() *InscriptionMintAction {
	this := InscriptionMintAction{}
	return &this
}

// GetRecipient returns the Recipient field value
func (o *InscriptionMintAction) GetRecipient() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *InscriptionMintAction) GetRecipientOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *InscriptionMintAction) SetRecipient(v AccountAddress) {
	o.Recipient = v
}

// GetAmount returns the Amount field value
func (o *InscriptionMintAction) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *InscriptionMintAction) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *InscriptionMintAction) SetAmount(v string) {
	o.Amount = v
}

// GetType returns the Type field value
func (o *InscriptionMintAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InscriptionMintAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InscriptionMintAction) SetType(v string) {
	o.Type = v
}

// GetTicker returns the Ticker field value
func (o *InscriptionMintAction) GetTicker() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ticker
}

// GetTickerOk returns a tuple with the Ticker field value
// and a boolean to check if the value has been set.
func (o *InscriptionMintAction) GetTickerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ticker, true
}

// SetTicker sets field value
func (o *InscriptionMintAction) SetTicker(v string) {
	o.Ticker = v
}

// GetDecimals returns the Decimals field value
func (o *InscriptionMintAction) GetDecimals() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *InscriptionMintAction) GetDecimalsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *InscriptionMintAction) SetDecimals(v int32) {
	o.Decimals = v
}

func (o InscriptionMintAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InscriptionMintAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recipient"] = o.Recipient
	toSerialize["amount"] = o.Amount
	toSerialize["type"] = o.Type
	toSerialize["ticker"] = o.Ticker
	toSerialize["decimals"] = o.Decimals
	return toSerialize, nil
}

func (o *InscriptionMintAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"recipient",
		"amount",
		"type",
		"ticker",
		"decimals",
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

	varInscriptionMintAction := _InscriptionMintAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInscriptionMintAction)

	if err != nil {
		return err
	}

	*o = InscriptionMintAction(varInscriptionMintAction)

	return err
}

type NullableInscriptionMintAction struct {
	value *InscriptionMintAction
	isSet bool
}

func (v NullableInscriptionMintAction) Get() *InscriptionMintAction {
	return v.value
}

func (v *NullableInscriptionMintAction) Set(val *InscriptionMintAction) {
	v.value = val
	v.isSet = true
}

func (v NullableInscriptionMintAction) IsSet() bool {
	return v.isSet
}

func (v *NullableInscriptionMintAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInscriptionMintAction(val *InscriptionMintAction) *NullableInscriptionMintAction {
	return &NullableInscriptionMintAction{value: val, isSet: true}
}

func (v NullableInscriptionMintAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInscriptionMintAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

