/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the LinkAggregationGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkAggregationGroup{}

// LinkAggregationGroup Link aggregation group (LAG) preferences and settings.
type LinkAggregationGroup struct {
	// Parameter showing whether LAG configuration is mandatory. The default is false.
	Enabled              *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LinkAggregationGroup LinkAggregationGroup

// NewLinkAggregationGroup instantiates a new LinkAggregationGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkAggregationGroup() *LinkAggregationGroup {
	this := LinkAggregationGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// NewLinkAggregationGroupWithDefaults instantiates a new LinkAggregationGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkAggregationGroupWithDefaults() *LinkAggregationGroup {
	this := LinkAggregationGroup{}
	var enabled bool = false
	this.Enabled = &enabled
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *LinkAggregationGroup) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkAggregationGroup) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *LinkAggregationGroup) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *LinkAggregationGroup) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o LinkAggregationGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkAggregationGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LinkAggregationGroup) UnmarshalJSON(data []byte) (err error) {
	varLinkAggregationGroup := _LinkAggregationGroup{}

	err = json.Unmarshal(data, &varLinkAggregationGroup)

	if err != nil {
		return err
	}

	*o = LinkAggregationGroup(varLinkAggregationGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLinkAggregationGroup struct {
	value *LinkAggregationGroup
	isSet bool
}

func (v NullableLinkAggregationGroup) Get() *LinkAggregationGroup {
	return v.value
}

func (v *NullableLinkAggregationGroup) Set(val *LinkAggregationGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkAggregationGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkAggregationGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkAggregationGroup(val *LinkAggregationGroup) *NullableLinkAggregationGroup {
	return &NullableLinkAggregationGroup{value: val, isSet: true}
}

func (v NullableLinkAggregationGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkAggregationGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
