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

// checks if the NftItemTransferAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NftItemTransferAction{}

// NftItemTransferAction struct for NftItemTransferAction
type NftItemTransferAction struct {
	Sender *AccountAddress `json:"sender,omitempty"`
	Recipient *AccountAddress `json:"recipient,omitempty"`
	Nft string `json:"nft"`
	Comment *string `json:"comment,omitempty"`
	EncryptedComment *EncryptedComment `json:"encrypted_comment,omitempty"`
	// raw hex encoded payload
	Payload *string `json:"payload,omitempty"`
	Refund *Refund `json:"refund,omitempty"`
}

type _NftItemTransferAction NftItemTransferAction

// NewNftItemTransferAction instantiates a new NftItemTransferAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftItemTransferAction(nft string) *NftItemTransferAction {
	this := NftItemTransferAction{}
	this.Nft = nft
	return &this
}

// NewNftItemTransferActionWithDefaults instantiates a new NftItemTransferAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftItemTransferActionWithDefaults() *NftItemTransferAction {
	this := NftItemTransferAction{}
	return &this
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetSender() AccountAddress {
	if o == nil || IsNil(o.Sender) {
		var ret AccountAddress
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetSenderOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.Sender) {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasSender() bool {
	if o != nil && !IsNil(o.Sender) {
		return true
	}

	return false
}

// SetSender gets a reference to the given AccountAddress and assigns it to the Sender field.
func (o *NftItemTransferAction) SetSender(v AccountAddress) {
	o.Sender = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetRecipient() AccountAddress {
	if o == nil || IsNil(o.Recipient) {
		var ret AccountAddress
		return ret
	}
	return *o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetRecipientOk() (*AccountAddress, bool) {
	if o == nil || IsNil(o.Recipient) {
		return nil, false
	}
	return o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasRecipient() bool {
	if o != nil && !IsNil(o.Recipient) {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given AccountAddress and assigns it to the Recipient field.
func (o *NftItemTransferAction) SetRecipient(v AccountAddress) {
	o.Recipient = &v
}

// GetNft returns the Nft field value
func (o *NftItemTransferAction) GetNft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nft
}

// GetNftOk returns a tuple with the Nft field value
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetNftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nft, true
}

// SetNft sets field value
func (o *NftItemTransferAction) SetNft(v string) {
	o.Nft = v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *NftItemTransferAction) SetComment(v string) {
	o.Comment = &v
}

// GetEncryptedComment returns the EncryptedComment field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetEncryptedComment() EncryptedComment {
	if o == nil || IsNil(o.EncryptedComment) {
		var ret EncryptedComment
		return ret
	}
	return *o.EncryptedComment
}

// GetEncryptedCommentOk returns a tuple with the EncryptedComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetEncryptedCommentOk() (*EncryptedComment, bool) {
	if o == nil || IsNil(o.EncryptedComment) {
		return nil, false
	}
	return o.EncryptedComment, true
}

// HasEncryptedComment returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasEncryptedComment() bool {
	if o != nil && !IsNil(o.EncryptedComment) {
		return true
	}

	return false
}

// SetEncryptedComment gets a reference to the given EncryptedComment and assigns it to the EncryptedComment field.
func (o *NftItemTransferAction) SetEncryptedComment(v EncryptedComment) {
	o.EncryptedComment = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *NftItemTransferAction) SetPayload(v string) {
	o.Payload = &v
}

// GetRefund returns the Refund field value if set, zero value otherwise.
func (o *NftItemTransferAction) GetRefund() Refund {
	if o == nil || IsNil(o.Refund) {
		var ret Refund
		return ret
	}
	return *o.Refund
}

// GetRefundOk returns a tuple with the Refund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftItemTransferAction) GetRefundOk() (*Refund, bool) {
	if o == nil || IsNil(o.Refund) {
		return nil, false
	}
	return o.Refund, true
}

// HasRefund returns a boolean if a field has been set.
func (o *NftItemTransferAction) HasRefund() bool {
	if o != nil && !IsNil(o.Refund) {
		return true
	}

	return false
}

// SetRefund gets a reference to the given Refund and assigns it to the Refund field.
func (o *NftItemTransferAction) SetRefund(v Refund) {
	o.Refund = &v
}

func (o NftItemTransferAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NftItemTransferAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sender) {
		toSerialize["sender"] = o.Sender
	}
	if !IsNil(o.Recipient) {
		toSerialize["recipient"] = o.Recipient
	}
	toSerialize["nft"] = o.Nft
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.EncryptedComment) {
		toSerialize["encrypted_comment"] = o.EncryptedComment
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Refund) {
		toSerialize["refund"] = o.Refund
	}
	return toSerialize, nil
}

func (o *NftItemTransferAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nft",
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

	varNftItemTransferAction := _NftItemTransferAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNftItemTransferAction)

	if err != nil {
		return err
	}

	*o = NftItemTransferAction(varNftItemTransferAction)

	return err
}

type NullableNftItemTransferAction struct {
	value *NftItemTransferAction
	isSet bool
}

func (v NullableNftItemTransferAction) Get() *NftItemTransferAction {
	return v.value
}

func (v *NullableNftItemTransferAction) Set(val *NftItemTransferAction) {
	v.value = val
	v.isSet = true
}

func (v NullableNftItemTransferAction) IsSet() bool {
	return v.isSet
}

func (v *NullableNftItemTransferAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftItemTransferAction(val *NftItemTransferAction) *NullableNftItemTransferAction {
	return &NullableNftItemTransferAction{value: val, isSet: true}
}

func (v NullableNftItemTransferAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftItemTransferAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

