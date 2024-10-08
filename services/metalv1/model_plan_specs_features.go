/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
)

// checks if the PlanSpecsFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlanSpecsFeatures{}

// PlanSpecsFeatures struct for PlanSpecsFeatures
type PlanSpecsFeatures struct {
	Raid                 *bool `json:"raid,omitempty"`
	Txt                  *bool `json:"txt,omitempty"`
	Uefi                 *bool `json:"uefi,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlanSpecsFeatures PlanSpecsFeatures

// NewPlanSpecsFeatures instantiates a new PlanSpecsFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanSpecsFeatures() *PlanSpecsFeatures {
	this := PlanSpecsFeatures{}
	return &this
}

// NewPlanSpecsFeaturesWithDefaults instantiates a new PlanSpecsFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanSpecsFeaturesWithDefaults() *PlanSpecsFeatures {
	this := PlanSpecsFeatures{}
	return &this
}

// GetRaid returns the Raid field value if set, zero value otherwise.
func (o *PlanSpecsFeatures) GetRaid() bool {
	if o == nil || IsNil(o.Raid) {
		var ret bool
		return ret
	}
	return *o.Raid
}

// GetRaidOk returns a tuple with the Raid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsFeatures) GetRaidOk() (*bool, bool) {
	if o == nil || IsNil(o.Raid) {
		return nil, false
	}
	return o.Raid, true
}

// HasRaid returns a boolean if a field has been set.
func (o *PlanSpecsFeatures) HasRaid() bool {
	if o != nil && !IsNil(o.Raid) {
		return true
	}

	return false
}

// SetRaid gets a reference to the given bool and assigns it to the Raid field.
func (o *PlanSpecsFeatures) SetRaid(v bool) {
	o.Raid = &v
}

// GetTxt returns the Txt field value if set, zero value otherwise.
func (o *PlanSpecsFeatures) GetTxt() bool {
	if o == nil || IsNil(o.Txt) {
		var ret bool
		return ret
	}
	return *o.Txt
}

// GetTxtOk returns a tuple with the Txt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsFeatures) GetTxtOk() (*bool, bool) {
	if o == nil || IsNil(o.Txt) {
		return nil, false
	}
	return o.Txt, true
}

// HasTxt returns a boolean if a field has been set.
func (o *PlanSpecsFeatures) HasTxt() bool {
	if o != nil && !IsNil(o.Txt) {
		return true
	}

	return false
}

// SetTxt gets a reference to the given bool and assigns it to the Txt field.
func (o *PlanSpecsFeatures) SetTxt(v bool) {
	o.Txt = &v
}

// GetUefi returns the Uefi field value if set, zero value otherwise.
func (o *PlanSpecsFeatures) GetUefi() bool {
	if o == nil || IsNil(o.Uefi) {
		var ret bool
		return ret
	}
	return *o.Uefi
}

// GetUefiOk returns a tuple with the Uefi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanSpecsFeatures) GetUefiOk() (*bool, bool) {
	if o == nil || IsNil(o.Uefi) {
		return nil, false
	}
	return o.Uefi, true
}

// HasUefi returns a boolean if a field has been set.
func (o *PlanSpecsFeatures) HasUefi() bool {
	if o != nil && !IsNil(o.Uefi) {
		return true
	}

	return false
}

// SetUefi gets a reference to the given bool and assigns it to the Uefi field.
func (o *PlanSpecsFeatures) SetUefi(v bool) {
	o.Uefi = &v
}

func (o PlanSpecsFeatures) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlanSpecsFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Raid) {
		toSerialize["raid"] = o.Raid
	}
	if !IsNil(o.Txt) {
		toSerialize["txt"] = o.Txt
	}
	if !IsNil(o.Uefi) {
		toSerialize["uefi"] = o.Uefi
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlanSpecsFeatures) UnmarshalJSON(data []byte) (err error) {
	varPlanSpecsFeatures := _PlanSpecsFeatures{}

	err = json.Unmarshal(data, &varPlanSpecsFeatures)

	if err != nil {
		return err
	}

	*o = PlanSpecsFeatures(varPlanSpecsFeatures)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "raid")
		delete(additionalProperties, "txt")
		delete(additionalProperties, "uefi")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlanSpecsFeatures struct {
	value *PlanSpecsFeatures
	isSet bool
}

func (v NullablePlanSpecsFeatures) Get() *PlanSpecsFeatures {
	return v.value
}

func (v *NullablePlanSpecsFeatures) Set(val *PlanSpecsFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanSpecsFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanSpecsFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanSpecsFeatures(val *PlanSpecsFeatures) *NullablePlanSpecsFeatures {
	return &NullablePlanSpecsFeatures{value: val, isSet: true}
}

func (v NullablePlanSpecsFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanSpecsFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
