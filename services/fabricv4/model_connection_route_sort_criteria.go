/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ConnectionRouteSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionRouteSortCriteria{}

// ConnectionRouteSortCriteria struct for ConnectionRouteSortCriteria
type ConnectionRouteSortCriteria struct {
	Direction            *ConnectionRouteEntrySortDirection `json:"direction,omitempty"`
	Property             *ConnectionRouteEntrySortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionRouteSortCriteria ConnectionRouteSortCriteria

// NewConnectionRouteSortCriteria instantiates a new ConnectionRouteSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionRouteSortCriteria() *ConnectionRouteSortCriteria {
	this := ConnectionRouteSortCriteria{}
	var direction ConnectionRouteEntrySortDirection = CONNECTIONROUTEENTRYSORTDIRECTION_DESC
	this.Direction = &direction
	var property ConnectionRouteEntrySortBy = CONNECTIONROUTEENTRYSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewConnectionRouteSortCriteriaWithDefaults instantiates a new ConnectionRouteSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionRouteSortCriteriaWithDefaults() *ConnectionRouteSortCriteria {
	this := ConnectionRouteSortCriteria{}
	var direction ConnectionRouteEntrySortDirection = CONNECTIONROUTEENTRYSORTDIRECTION_DESC
	this.Direction = &direction
	var property ConnectionRouteEntrySortBy = CONNECTIONROUTEENTRYSORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ConnectionRouteSortCriteria) GetDirection() ConnectionRouteEntrySortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret ConnectionRouteEntrySortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteSortCriteria) GetDirectionOk() (*ConnectionRouteEntrySortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ConnectionRouteSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given ConnectionRouteEntrySortDirection and assigns it to the Direction field.
func (o *ConnectionRouteSortCriteria) SetDirection(v ConnectionRouteEntrySortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ConnectionRouteSortCriteria) GetProperty() ConnectionRouteEntrySortBy {
	if o == nil || IsNil(o.Property) {
		var ret ConnectionRouteEntrySortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionRouteSortCriteria) GetPropertyOk() (*ConnectionRouteEntrySortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ConnectionRouteSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given ConnectionRouteEntrySortBy and assigns it to the Property field.
func (o *ConnectionRouteSortCriteria) SetProperty(v ConnectionRouteEntrySortBy) {
	o.Property = &v
}

func (o ConnectionRouteSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionRouteSortCriteria) ToMap() (map[string]interface{}, error) {
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

func (o *ConnectionRouteSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varConnectionRouteSortCriteria := _ConnectionRouteSortCriteria{}

	err = json.Unmarshal(data, &varConnectionRouteSortCriteria)

	if err != nil {
		return err
	}

	*o = ConnectionRouteSortCriteria(varConnectionRouteSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionRouteSortCriteria struct {
	value *ConnectionRouteSortCriteria
	isSet bool
}

func (v NullableConnectionRouteSortCriteria) Get() *ConnectionRouteSortCriteria {
	return v.value
}

func (v *NullableConnectionRouteSortCriteria) Set(val *ConnectionRouteSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteSortCriteria(val *ConnectionRouteSortCriteria) *NullableConnectionRouteSortCriteria {
	return &NullableConnectionRouteSortCriteria{value: val, isSet: true}
}

func (v NullableConnectionRouteSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
