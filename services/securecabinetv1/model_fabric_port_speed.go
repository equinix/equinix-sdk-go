/*
Secure Cabinet API

Secure Cabinet API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package securecabinetv1

import (
	"encoding/json"
	"fmt"
)

// FabricPortSpeed Port speed
type FabricPortSpeed string

// List of FabricPortSpeed
const (
	FABRICPORTSPEED__1_GBPS  FabricPortSpeed = "SPEED_1_GBPS"
	FABRICPORTSPEED__10_GBPS FabricPortSpeed = "SPEED_10_GBPS"
)

// All allowed values of FabricPortSpeed enum
var AllowedFabricPortSpeedEnumValues = []FabricPortSpeed{
	"SPEED_1_GBPS",
	"SPEED_10_GBPS",
}

func (v *FabricPortSpeed) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FabricPortSpeed(value)
	for _, existing := range AllowedFabricPortSpeedEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FabricPortSpeed", value)
}

// NewFabricPortSpeedFromValue returns a pointer to a valid FabricPortSpeed
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFabricPortSpeedFromValue(v string) (*FabricPortSpeed, error) {
	ev := FabricPortSpeed(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FabricPortSpeed: valid values are %v", v, AllowedFabricPortSpeedEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FabricPortSpeed) IsValid() bool {
	for _, existing := range AllowedFabricPortSpeedEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FabricPortSpeed value
func (v FabricPortSpeed) Ptr() *FabricPortSpeed {
	return &v
}

type NullableFabricPortSpeed struct {
	value *FabricPortSpeed
	isSet bool
}

func (v NullableFabricPortSpeed) Get() *FabricPortSpeed {
	return v.value
}

func (v *NullableFabricPortSpeed) Set(val *FabricPortSpeed) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricPortSpeed) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricPortSpeed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricPortSpeed(val *FabricPortSpeed) *NullableFabricPortSpeed {
	return &NullableFabricPortSpeed{value: val, isSet: true}
}

func (v NullableFabricPortSpeed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricPortSpeed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
