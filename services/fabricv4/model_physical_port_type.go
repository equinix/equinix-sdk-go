/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PhysicalPortType Type of Port
type PhysicalPortType string

// List of PhysicalPortType
const (
	PHYSICALPORTTYPE_XF_PHYSICAL_PORT PhysicalPortType = "XF_PHYSICAL_PORT"
)

// All allowed values of PhysicalPortType enum
var AllowedPhysicalPortTypeEnumValues = []PhysicalPortType{
	"XF_PHYSICAL_PORT",
}

func (v *PhysicalPortType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhysicalPortType(value)
	for _, existing := range AllowedPhysicalPortTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhysicalPortType", value)
}

// NewPhysicalPortTypeFromValue returns a pointer to a valid PhysicalPortType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhysicalPortTypeFromValue(v string) (*PhysicalPortType, error) {
	ev := PhysicalPortType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PhysicalPortType: valid values are %v", v, AllowedPhysicalPortTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhysicalPortType) IsValid() bool {
	for _, existing := range AllowedPhysicalPortTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhysicalPortType value
func (v PhysicalPortType) Ptr() *PhysicalPortType {
	return &v
}

type NullablePhysicalPortType struct {
	value *PhysicalPortType
	isSet bool
}

func (v NullablePhysicalPortType) Get() *PhysicalPortType {
	return v.value
}

func (v *NullablePhysicalPortType) Set(val *PhysicalPortType) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalPortType) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalPortType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalPortType(val *PhysicalPortType) *NullablePhysicalPortType {
	return &NullablePhysicalPortType{value: val, isSet: true}
}

func (v NullablePhysicalPortType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalPortType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
