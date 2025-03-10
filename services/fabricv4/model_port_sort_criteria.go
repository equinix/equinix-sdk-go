/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PortSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortSortCriteria{}

// PortSortCriteria struct for PortSortCriteria
type PortSortCriteria struct {
	Direction            *PortSortDirection `json:"direction,omitempty"`
	Property             *PortSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortSortCriteria PortSortCriteria

// NewPortSortCriteria instantiates a new PortSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortSortCriteria() *PortSortCriteria {
	this := PortSortCriteria{}
	var direction PortSortDirection = PORTSORTDIRECTION_DESC
	this.Direction = &direction
	var property PortSortBy = PORTSORTBY_DEVICE_NAME
	this.Property = &property
	return &this
}

// NewPortSortCriteriaWithDefaults instantiates a new PortSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortSortCriteriaWithDefaults() *PortSortCriteria {
	this := PortSortCriteria{}
	var direction PortSortDirection = PORTSORTDIRECTION_DESC
	this.Direction = &direction
	var property PortSortBy = PORTSORTBY_DEVICE_NAME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *PortSortCriteria) GetDirection() PortSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret PortSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSortCriteria) GetDirectionOk() (*PortSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *PortSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given PortSortDirection and assigns it to the Direction field.
func (o *PortSortCriteria) SetDirection(v PortSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *PortSortCriteria) GetProperty() PortSortBy {
	if o == nil || IsNil(o.Property) {
		var ret PortSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortSortCriteria) GetPropertyOk() (*PortSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *PortSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given PortSortBy and assigns it to the Property field.
func (o *PortSortCriteria) SetProperty(v PortSortBy) {
	o.Property = &v
}

func (o PortSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortSortCriteria) ToMap() (map[string]interface{}, error) {
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

func (o *PortSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varPortSortCriteria := _PortSortCriteria{}

	err = json.Unmarshal(data, &varPortSortCriteria)

	if err != nil {
		return err
	}

	*o = PortSortCriteria(varPortSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortSortCriteria struct {
	value *PortSortCriteria
	isSet bool
}

func (v NullablePortSortCriteria) Get() *PortSortCriteria {
	return v.value
}

func (v *NullablePortSortCriteria) Set(val *PortSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullablePortSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullablePortSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortSortCriteria(val *PortSortCriteria) *NullablePortSortCriteria {
	return &NullablePortSortCriteria{value: val, isSet: true}
}

func (v NullablePortSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
