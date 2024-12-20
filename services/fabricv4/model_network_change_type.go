/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// NetworkChangeType Type of change
type NetworkChangeType string

// List of NetworkChangeType
const (
	NETWORKCHANGETYPE_CREATION NetworkChangeType = "NETWORK_CREATION"
	NETWORKCHANGETYPE_UPDATE   NetworkChangeType = "NETWORK_UPDATE"
	NETWORKCHANGETYPE_DELETION NetworkChangeType = "NETWORK_DELETION"
)

// All allowed values of NetworkChangeType enum
var AllowedNetworkChangeTypeEnumValues = []NetworkChangeType{
	"NETWORK_CREATION",
	"NETWORK_UPDATE",
	"NETWORK_DELETION",
}

func (v *NetworkChangeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkChangeType(value)
	for _, existing := range AllowedNetworkChangeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkChangeType", value)
}

// NewNetworkChangeTypeFromValue returns a pointer to a valid NetworkChangeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkChangeTypeFromValue(v string) (*NetworkChangeType, error) {
	ev := NetworkChangeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkChangeType: valid values are %v", v, AllowedNetworkChangeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkChangeType) IsValid() bool {
	for _, existing := range AllowedNetworkChangeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkChangeType value
func (v NetworkChangeType) Ptr() *NetworkChangeType {
	return &v
}

type NullableNetworkChangeType struct {
	value *NetworkChangeType
	isSet bool
}

func (v NullableNetworkChangeType) Get() *NetworkChangeType {
	return v.value
}

func (v *NullableNetworkChangeType) Set(val *NetworkChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkChangeType(val *NetworkChangeType) *NullableNetworkChangeType {
	return &NullableNetworkChangeType{value: val, isSet: true}
}

func (v NullableNetworkChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
