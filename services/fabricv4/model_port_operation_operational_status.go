/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortOperationOperationalStatus Availability of a given physical port.
type PortOperationOperationalStatus string

// List of PortOperation_operationalStatus
const (
	PORTOPERATIONOPERATIONALSTATUS_UP      PortOperationOperationalStatus = "UP"
	PORTOPERATIONOPERATIONALSTATUS_DOWN    PortOperationOperationalStatus = "DOWN"
	PORTOPERATIONOPERATIONALSTATUS_PARTIAL PortOperationOperationalStatus = "PARTIAL"
)

// All allowed values of PortOperationOperationalStatus enum
var AllowedPortOperationOperationalStatusEnumValues = []PortOperationOperationalStatus{
	"UP",
	"DOWN",
	"PARTIAL",
}

func (v *PortOperationOperationalStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortOperationOperationalStatus(value)
	for _, existing := range AllowedPortOperationOperationalStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortOperationOperationalStatus", value)
}

// NewPortOperationOperationalStatusFromValue returns a pointer to a valid PortOperationOperationalStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortOperationOperationalStatusFromValue(v string) (*PortOperationOperationalStatus, error) {
	ev := PortOperationOperationalStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortOperationOperationalStatus: valid values are %v", v, AllowedPortOperationOperationalStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortOperationOperationalStatus) IsValid() bool {
	for _, existing := range AllowedPortOperationOperationalStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortOperation_operationalStatus value
func (v PortOperationOperationalStatus) Ptr() *PortOperationOperationalStatus {
	return &v
}

type NullablePortOperationOperationalStatus struct {
	value *PortOperationOperationalStatus
	isSet bool
}

func (v NullablePortOperationOperationalStatus) Get() *PortOperationOperationalStatus {
	return v.value
}

func (v *NullablePortOperationOperationalStatus) Set(val *PortOperationOperationalStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePortOperationOperationalStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePortOperationOperationalStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortOperationOperationalStatus(val *PortOperationOperationalStatus) *NullablePortOperationOperationalStatus {
	return &NullablePortOperationOperationalStatus{value: val, isSet: true}
}

func (v NullablePortOperationOperationalStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortOperationOperationalStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
