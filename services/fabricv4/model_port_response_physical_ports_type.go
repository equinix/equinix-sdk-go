/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortResponsePhysicalPortsType Physical Ports Type
type PortResponsePhysicalPortsType string

// List of PortResponse_physicalPortsType
const (
	PORTRESPONSEPHYSICALPORTSTYPE__1000_BASE_LX  PortResponsePhysicalPortsType = "1000BASE_LX"
	PORTRESPONSEPHYSICALPORTSTYPE__10_GBASE_LR   PortResponsePhysicalPortsType = "10GBASE_LR"
	PORTRESPONSEPHYSICALPORTSTYPE__100_GBASE_LR4 PortResponsePhysicalPortsType = "100GBASE_LR4"
	PORTRESPONSEPHYSICALPORTSTYPE__10_GBASE_ER   PortResponsePhysicalPortsType = "10GBASE_ER"
	PORTRESPONSEPHYSICALPORTSTYPE__1000_BASE_SX  PortResponsePhysicalPortsType = "1000BASE_SX"
)

// All allowed values of PortResponsePhysicalPortsType enum
var AllowedPortResponsePhysicalPortsTypeEnumValues = []PortResponsePhysicalPortsType{
	"1000BASE_LX",
	"10GBASE_LR",
	"100GBASE_LR4",
	"10GBASE_ER",
	"1000BASE_SX",
}

func (v *PortResponsePhysicalPortsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortResponsePhysicalPortsType(value)
	for _, existing := range AllowedPortResponsePhysicalPortsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortResponsePhysicalPortsType", value)
}

// NewPortResponsePhysicalPortsTypeFromValue returns a pointer to a valid PortResponsePhysicalPortsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortResponsePhysicalPortsTypeFromValue(v string) (*PortResponsePhysicalPortsType, error) {
	ev := PortResponsePhysicalPortsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortResponsePhysicalPortsType: valid values are %v", v, AllowedPortResponsePhysicalPortsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortResponsePhysicalPortsType) IsValid() bool {
	for _, existing := range AllowedPortResponsePhysicalPortsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortResponse_physicalPortsType value
func (v PortResponsePhysicalPortsType) Ptr() *PortResponsePhysicalPortsType {
	return &v
}

type NullablePortResponsePhysicalPortsType struct {
	value *PortResponsePhysicalPortsType
	isSet bool
}

func (v NullablePortResponsePhysicalPortsType) Get() *PortResponsePhysicalPortsType {
	return v.value
}

func (v *NullablePortResponsePhysicalPortsType) Set(val *PortResponsePhysicalPortsType) {
	v.value = val
	v.isSet = true
}

func (v NullablePortResponsePhysicalPortsType) IsSet() bool {
	return v.isSet
}

func (v *NullablePortResponsePhysicalPortsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortResponsePhysicalPortsType(val *PortResponsePhysicalPortsType) *NullablePortResponsePhysicalPortsType {
	return &NullablePortResponsePhysicalPortsType{value: val, isSet: true}
}

func (v NullablePortResponsePhysicalPortsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortResponsePhysicalPortsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
