/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ServiceProfileAccessPointVDType the model 'ServiceProfileAccessPointVDType'
type ServiceProfileAccessPointVDType string

// List of ServiceProfileAccessPointVD_type
const (
	SERVICEPROFILEACCESSPOINTVDTYPE_VD ServiceProfileAccessPointVDType = "VD"
)

// All allowed values of ServiceProfileAccessPointVDType enum
var AllowedServiceProfileAccessPointVDTypeEnumValues = []ServiceProfileAccessPointVDType{
	"VD",
}

func (v *ServiceProfileAccessPointVDType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceProfileAccessPointVDType(value)
	for _, existing := range AllowedServiceProfileAccessPointVDTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceProfileAccessPointVDType", value)
}

// NewServiceProfileAccessPointVDTypeFromValue returns a pointer to a valid ServiceProfileAccessPointVDType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceProfileAccessPointVDTypeFromValue(v string) (*ServiceProfileAccessPointVDType, error) {
	ev := ServiceProfileAccessPointVDType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceProfileAccessPointVDType: valid values are %v", v, AllowedServiceProfileAccessPointVDTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceProfileAccessPointVDType) IsValid() bool {
	for _, existing := range AllowedServiceProfileAccessPointVDTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceProfileAccessPointVD_type value
func (v ServiceProfileAccessPointVDType) Ptr() *ServiceProfileAccessPointVDType {
	return &v
}

type NullableServiceProfileAccessPointVDType struct {
	value *ServiceProfileAccessPointVDType
	isSet bool
}

func (v NullableServiceProfileAccessPointVDType) Get() *ServiceProfileAccessPointVDType {
	return v.value
}

func (v *NullableServiceProfileAccessPointVDType) Set(val *ServiceProfileAccessPointVDType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointVDType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointVDType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointVDType(val *ServiceProfileAccessPointVDType) *NullableServiceProfileAccessPointVDType {
	return &NullableServiceProfileAccessPointVDType{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointVDType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointVDType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
