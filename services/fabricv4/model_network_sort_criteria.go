/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the NetworkSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSortCriteria{}

// NetworkSortCriteria struct for NetworkSortCriteria
type NetworkSortCriteria struct {
	Direction            *NetworkSortDirection `json:"direction,omitempty"`
	Property             *NetworkSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSortCriteria NetworkSortCriteria

// NewNetworkSortCriteria instantiates a new NetworkSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSortCriteria() *NetworkSortCriteria {
	this := NetworkSortCriteria{}
	var direction NetworkSortDirection = NETWORKSORTDIRECTION_DESC
	this.Direction = &direction
	var property NetworkSortBy = NETWORKSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewNetworkSortCriteriaWithDefaults instantiates a new NetworkSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSortCriteriaWithDefaults() *NetworkSortCriteria {
	this := NetworkSortCriteria{}
	var direction NetworkSortDirection = NETWORKSORTDIRECTION_DESC
	this.Direction = &direction
	var property NetworkSortBy = NETWORKSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *NetworkSortCriteria) GetDirection() NetworkSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret NetworkSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSortCriteria) GetDirectionOk() (*NetworkSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *NetworkSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NetworkSortDirection and assigns it to the Direction field.
func (o *NetworkSortCriteria) SetDirection(v NetworkSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *NetworkSortCriteria) GetProperty() NetworkSortBy {
	if o == nil || IsNil(o.Property) {
		var ret NetworkSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSortCriteria) GetPropertyOk() (*NetworkSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *NetworkSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given NetworkSortBy and assigns it to the Property field.
func (o *NetworkSortCriteria) SetProperty(v NetworkSortBy) {
	o.Property = &v
}

func (o NetworkSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSortCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varNetworkSortCriteria := _NetworkSortCriteria{}

	err = json.Unmarshal(data, &varNetworkSortCriteria)

	if err != nil {
		return err
	}

	*o = NetworkSortCriteria(varNetworkSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSortCriteria struct {
	value *NetworkSortCriteria
	isSet bool
}

func (v NullableNetworkSortCriteria) Get() *NetworkSortCriteria {
	return v.value
}

func (v *NullableNetworkSortCriteria) Set(val *NetworkSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSortCriteria(val *NetworkSortCriteria) *NullableNetworkSortCriteria {
	return &NullableNetworkSortCriteria{value: val, isSet: true}
}

func (v NullableNetworkSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
