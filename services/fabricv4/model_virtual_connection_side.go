/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualConnectionSide type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualConnectionSide{}

// VirtualConnectionSide Fabric Connection access point object.
type VirtualConnectionSide struct {
	AccessPoint          *AccessPoint `json:"accessPoint,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualConnectionSide VirtualConnectionSide

// NewVirtualConnectionSide instantiates a new VirtualConnectionSide object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualConnectionSide() *VirtualConnectionSide {
	this := VirtualConnectionSide{}
	return &this
}

// NewVirtualConnectionSideWithDefaults instantiates a new VirtualConnectionSide object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualConnectionSideWithDefaults() *VirtualConnectionSide {
	this := VirtualConnectionSide{}
	return &this
}

// GetAccessPoint returns the AccessPoint field value if set, zero value otherwise.
func (o *VirtualConnectionSide) GetAccessPoint() AccessPoint {
	if o == nil || IsNil(o.AccessPoint) {
		var ret AccessPoint
		return ret
	}
	return *o.AccessPoint
}

// GetAccessPointOk returns a tuple with the AccessPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualConnectionSide) GetAccessPointOk() (*AccessPoint, bool) {
	if o == nil || IsNil(o.AccessPoint) {
		return nil, false
	}
	return o.AccessPoint, true
}

// HasAccessPoint returns a boolean if a field has been set.
func (o *VirtualConnectionSide) HasAccessPoint() bool {
	if o != nil && !IsNil(o.AccessPoint) {
		return true
	}

	return false
}

// SetAccessPoint gets a reference to the given AccessPoint and assigns it to the AccessPoint field.
func (o *VirtualConnectionSide) SetAccessPoint(v AccessPoint) {
	o.AccessPoint = &v
}

func (o VirtualConnectionSide) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualConnectionSide) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccessPoint) {
		toSerialize["accessPoint"] = o.AccessPoint
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualConnectionSide) UnmarshalJSON(data []byte) (err error) {
	varVirtualConnectionSide := _VirtualConnectionSide{}

	err = json.Unmarshal(data, &varVirtualConnectionSide)

	if err != nil {
		return err
	}

	*o = VirtualConnectionSide(varVirtualConnectionSide)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "accessPoint")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualConnectionSide struct {
	value *VirtualConnectionSide
	isSet bool
}

func (v NullableVirtualConnectionSide) Get() *VirtualConnectionSide {
	return v.value
}

func (v *NullableVirtualConnectionSide) Set(val *VirtualConnectionSide) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualConnectionSide) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualConnectionSide) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualConnectionSide(val *VirtualConnectionSide) *NullableVirtualConnectionSide {
	return &NullableVirtualConnectionSide{value: val, isSet: true}
}

func (v NullableVirtualConnectionSide) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualConnectionSide) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}