/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// RegionsList struct for RegionsList
type RegionsList struct {
	Regions *[]Region `json:"regions,omitempty"`
}

// NewRegionsList instantiates a new RegionsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionsList() *RegionsList {
	this := RegionsList{}
	return &this
}

// NewRegionsListWithDefaults instantiates a new RegionsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionsListWithDefaults() *RegionsList {
	this := RegionsList{}
	return &this
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *RegionsList) GetRegions() []Region {
	if o == nil || o.Regions == nil {
		var ret []Region
		return ret
	}
	return *o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegionsList) GetRegionsOk() (*[]Region, bool) {
	if o == nil || o.Regions == nil {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *RegionsList) HasRegions() bool {
	if o != nil && o.Regions != nil {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []Region and assigns it to the Regions field.
func (o *RegionsList) SetRegions(v []Region) {
	o.Regions = &v
}

func (o RegionsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Regions != nil {
		toSerialize["regions"] = o.Regions
	}
	return json.Marshal(toSerialize)
}

type NullableRegionsList struct {
	value *RegionsList
	isSet bool
}

func (v NullableRegionsList) Get() *RegionsList {
	return v.value
}

func (v *NullableRegionsList) Set(val *RegionsList) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionsList) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionsList(val *RegionsList) *NullableRegionsList {
	return &NullableRegionsList{value: val, isSet: true}
}

func (v NullableRegionsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

