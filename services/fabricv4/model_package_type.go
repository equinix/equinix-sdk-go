/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PackageType Port service Type
type PackageType string

// List of Package_type
const (
	PACKAGETYPE_EPL PackageType = "EPL"
	PACKAGETYPE_MSP PackageType = "MSP"
)

// All allowed values of PackageType enum
var AllowedPackageTypeEnumValues = []PackageType{
	"EPL",
	"MSP",
}

func (v *PackageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PackageType(value)
	for _, existing := range AllowedPackageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PackageType", value)
}

// NewPackageTypeFromValue returns a pointer to a valid PackageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPackageTypeFromValue(v string) (*PackageType, error) {
	ev := PackageType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PackageType: valid values are %v", v, AllowedPackageTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PackageType) IsValid() bool {
	for _, existing := range AllowedPackageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Package_type value
func (v PackageType) Ptr() *PackageType {
	return &v
}

type NullablePackageType struct {
	value *PackageType
	isSet bool
}

func (v NullablePackageType) Get() *PackageType {
	return v.value
}

func (v *NullablePackageType) Set(val *PackageType) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageType) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageType(val *PackageType) *NullablePackageType {
	return &NullablePackageType{value: val, isSet: true}
}

func (v NullablePackageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}