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

// checks if the TonTransferAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TonTransferAction{}

// TonTransferAction struct for TonTransferAction
type TonTransferAction struct {
	Sender AccountAddress `json:"sender"`
	Recipient AccountAddress `json:"recipient"`
	// amount in nanotons
	Amount int64 `json:"amount"`
	Comment *string `json:"comment,omitempty"`
	EncryptedComment *EncryptedComment `json:"encrypted_comment,omitempty"`
	Refund *Refund `json:"refund,omitempty"`
}

type _TonTransferAction TonTransferAction

// NewTonTransferAction instantiates a new TonTransferAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTonTransferAction(sender AccountAddress, recipient AccountAddress, amount int64) *TonTransferAction {
	this := TonTransferAction{}
	this.Sender = sender
	this.Recipient = recipient
	this.Amount = amount
	return &this
}

// NewTonTransferActionWithDefaults instantiates a new TonTransferAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTonTransferActionWithDefaults() *TonTransferAction {
	this := TonTransferAction{}
	return &this
}

// GetSender returns the Sender field value
func (o *TonTransferAction) GetSender() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetSenderOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *TonTransferAction) SetSender(v AccountAddress) {
	o.Sender = v
}

// GetRecipient returns the Recipient field value
func (o *TonTransferAction) GetRecipient() AccountAddress {
	if o == nil {
		var ret AccountAddress
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetRecipientOk() (*AccountAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *TonTransferAction) SetRecipient(v AccountAddress) {
	o.Recipient = v
}

// GetAmount returns the Amount field value
func (o *TonTransferAction) GetAmount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TonTransferAction) SetAmount(v int64) {
	o.Amount = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *TonTransferAction) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *TonTransferAction) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *TonTransferAction) SetComment(v string) {
	o.Comment = &v
}

// GetEncryptedComment returns the EncryptedComment field value if set, zero value otherwise.
func (o *TonTransferAction) GetEncryptedComment() EncryptedComment {
	if o == nil || IsNil(o.EncryptedComment) {
		var ret EncryptedComment
		return ret
	}
	return *o.EncryptedComment
}

// GetEncryptedCommentOk returns a tuple with the EncryptedComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetEncryptedCommentOk() (*EncryptedComment, bool) {
	if o == nil || IsNil(o.EncryptedComment) {
		return nil, false
	}
	return o.EncryptedComment, true
}

// HasEncryptedComment returns a boolean if a field has been set.
func (o *TonTransferAction) HasEncryptedComment() bool {
	if o != nil && !IsNil(o.EncryptedComment) {
		return true
	}

	return false
}

// SetEncryptedComment gets a reference to the given EncryptedComment and assigns it to the EncryptedComment field.
func (o *TonTransferAction) SetEncryptedComment(v EncryptedComment) {
	o.EncryptedComment = &v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *TonTransferAction) GetRefund() Refund {
	if o == nil || IsNil(o.Refund) {
		var ret Refund
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TonTransferAction) GetRefundOk() (*Refund, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *TonTransferAction) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given Refund and assigns it to the Refund field.
func (o *TonTransferAction) SetRefund(v Refund) {
	o.Refund = &v
}

func (o TonTransferAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TonTransferAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sender"] = o.Sender
	toSerialize["recipient"] = o.Recipient
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.EncryptedComment) {
		toSerialize["encrypted_comment"] = o.EncryptedComment
	}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	return toSerialize, nil
}

func (o *TonTransferAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sender",
		"recipient",
		"amount",
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

	varTonTransferAction := _TonTransferAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTonTransferAction)

	if err != nil {
		return err
	}

	*o = TonTransferAction(varTonTransferAction)

	return err
}

type NullableTonTransferAction struct {
	value *TonTransferAction
	isSet bool
}

func (v NullableTonTransferAction) Get() *TonTransferAction {
	return v.value
}

func (v *NullableTonTransferAction) Set(val *TonTransferAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTonTransferAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTonTransferAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTonTransferAction(val *TonTransferAction) *NullableTonTransferAction {
	return &NullableTonTransferAction{value: val, isSet: true}
}

func (v NullableTonTransferAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTonTransferAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

