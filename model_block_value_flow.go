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

// checks if the BlockValueFlow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockValueFlow{}

// BlockValueFlow struct for BlockValueFlow
type BlockValueFlow struct {
	FromPrevBlk BlockCurrencyCollection `json:"from_prev_blk"`
	ToNextBlk BlockCurrencyCollection `json:"to_next_blk"`
	Imported BlockCurrencyCollection `json:"imported"`
	Exported BlockCurrencyCollection `json:"exported"`
	FeesCollected BlockCurrencyCollection `json:"fees_collected"`
	Burned *BlockCurrencyCollection `json:"burned,omitempty"`
	FeesImported BlockCurrencyCollection `json:"fees_imported"`
	Recovered BlockCurrencyCollection `json:"recovered"`
	Created BlockCurrencyCollection `json:"created"`
	Minted BlockCurrencyCollection `json:"minted"`
}

type _BlockValueFlow BlockValueFlow

// NewBlockValueFlow instantiates a new BlockValueFlow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockValueFlow(fromPrevBlk BlockCurrencyCollection, toNextBlk BlockCurrencyCollection, imported BlockCurrencyCollection, exported BlockCurrencyCollection, feesCollected BlockCurrencyCollection, feesImported BlockCurrencyCollection, recovered BlockCurrencyCollection, created BlockCurrencyCollection, minted BlockCurrencyCollection) *BlockValueFlow {
	this := BlockValueFlow{}
	this.FromPrevBlk = fromPrevBlk
	this.ToNextBlk = toNextBlk
	this.Imported = imported
	this.Exported = exported
	this.FeesCollected = feesCollected
	this.FeesImported = feesImported
	this.Recovered = recovered
	this.Created = created
	this.Minted = minted
	return &this
}

// NewBlockValueFlowWithDefaults instantiates a new BlockValueFlow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockValueFlowWithDefaults() *BlockValueFlow {
	this := BlockValueFlow{}
	return &this
}

// GetFromPrevBlk returns the FromPrevBlk field value
func (o *BlockValueFlow) GetFromPrevBlk() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.FromPrevBlk
}

// GetFromPrevBlkOk returns a tuple with the FromPrevBlk field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetFromPrevBlkOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromPrevBlk, true
}

// SetFromPrevBlk sets field value
func (o *BlockValueFlow) SetFromPrevBlk(v BlockCurrencyCollection) {
	o.FromPrevBlk = v
}

// GetToNextBlk returns the ToNextBlk field value
func (o *BlockValueFlow) GetToNextBlk() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.ToNextBlk
}

// GetToNextBlkOk returns a tuple with the ToNextBlk field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetToNextBlkOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToNextBlk, true
}

// SetToNextBlk sets field value
func (o *BlockValueFlow) SetToNextBlk(v BlockCurrencyCollection) {
	o.ToNextBlk = v
}

// GetImported returns the Imported field value
func (o *BlockValueFlow) GetImported() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.Imported
}

// GetImportedOk returns a tuple with the Imported field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetImportedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Imported, true
}

// SetImported sets field value
func (o *BlockValueFlow) SetImported(v BlockCurrencyCollection) {
	o.Imported = v
}

// GetExported returns the Exported field value
func (o *BlockValueFlow) GetExported() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.Exported
}

// GetExportedOk returns a tuple with the Exported field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetExportedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exported, true
}

// SetExported sets field value
func (o *BlockValueFlow) SetExported(v BlockCurrencyCollection) {
	o.Exported = v
}

// GetFeesCollected returns the FeesCollected field value
func (o *BlockValueFlow) GetFeesCollected() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.FeesCollected
}

// GetFeesCollectedOk returns a tuple with the FeesCollected field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetFeesCollectedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeesCollected, true
}

// SetFeesCollected sets field value
func (o *BlockValueFlow) SetFeesCollected(v BlockCurrencyCollection) {
	o.FeesCollected = v
}

// GetBurned returns the Burned field value if set, zero value otherwise.
func (o *BlockValueFlow) GetBurned() BlockCurrencyCollection {
	if o == nil || IsNil(o.Burned) {
		var ret BlockCurrencyCollection
		return ret
	}
	return *o.Burned
}

// GetBurnedOk returns a tuple with the Burned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetBurnedOk() (*BlockCurrencyCollection, bool) {
	if o == nil || IsNil(o.Burned) {
		return nil, false
	}
	return o.Burned, true
}

// HasBurned returns a boolean if a field has been set.
func (o *BlockValueFlow) HasBurned() bool {
	if o != nil && !IsNil(o.Burned) {
		return true
	}

	return false
}

// SetBurned gets a reference to the given BlockCurrencyCollection and assigns it to the Burned field.
func (o *BlockValueFlow) SetBurned(v BlockCurrencyCollection) {
	o.Burned = &v
}

// GetFeesImported returns the FeesImported field value
func (o *BlockValueFlow) GetFeesImported() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.FeesImported
}

// GetFeesImportedOk returns a tuple with the FeesImported field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetFeesImportedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeesImported, true
}

// SetFeesImported sets field value
func (o *BlockValueFlow) SetFeesImported(v BlockCurrencyCollection) {
	o.FeesImported = v
}

// GetRecovered returns the Recovered field value
func (o *BlockValueFlow) GetRecovered() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.Recovered
}

// GetRecoveredOk returns a tuple with the Recovered field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetRecoveredOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recovered, true
}

// SetRecovered sets field value
func (o *BlockValueFlow) SetRecovered(v BlockCurrencyCollection) {
	o.Recovered = v
}

// GetCreated returns the Created field value
func (o *BlockValueFlow) GetCreated() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetCreatedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *BlockValueFlow) SetCreated(v BlockCurrencyCollection) {
	o.Created = v
}

// GetMinted returns the Minted field value
func (o *BlockValueFlow) GetMinted() BlockCurrencyCollection {
	if o == nil {
		var ret BlockCurrencyCollection
		return ret
	}

	return o.Minted
}

// GetMintedOk returns a tuple with the Minted field value
// and a boolean to check if the value has been set.
func (o *BlockValueFlow) GetMintedOk() (*BlockCurrencyCollection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minted, true
}

// SetMinted sets field value
func (o *BlockValueFlow) SetMinted(v BlockCurrencyCollection) {
	o.Minted = v
}

func (o BlockValueFlow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockValueFlow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from_prev_blk"] = o.FromPrevBlk
	toSerialize["to_next_blk"] = o.ToNextBlk
	toSerialize["imported"] = o.Imported
	toSerialize["exported"] = o.Exported
	toSerialize["fees_collected"] = o.FeesCollected
	if !IsNil(o.Burned) {
		toSerialize["burned"] = o.Burned
	}
	toSerialize["fees_imported"] = o.FeesImported
	toSerialize["recovered"] = o.Recovered
	toSerialize["created"] = o.Created
	toSerialize["minted"] = o.Minted
	return toSerialize, nil
}

func (o *BlockValueFlow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from_prev_blk",
		"to_next_blk",
		"imported",
		"exported",
		"fees_collected",
		"fees_imported",
		"recovered",
		"created",
		"minted",
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

	varBlockValueFlow := _BlockValueFlow{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBlockValueFlow)

	if err != nil {
		return err
	}

	*o = BlockValueFlow(varBlockValueFlow)

	return err
}

type NullableBlockValueFlow struct {
	value *BlockValueFlow
	isSet bool
}

func (v NullableBlockValueFlow) Get() *BlockValueFlow {
	return v.value
}

func (v *NullableBlockValueFlow) Set(val *BlockValueFlow) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockValueFlow) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockValueFlow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockValueFlow(val *BlockValueFlow) *NullableBlockValueFlow {
	return &NullableBlockValueFlow{value: val, isSet: true}
}

func (v NullableBlockValueFlow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockValueFlow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

