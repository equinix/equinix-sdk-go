/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// GetTimeServicesPackageByCodePackageCodeParameter the model 'GetTimeServicesPackageByCodePackageCodeParameter'
type GetTimeServicesPackageByCodePackageCodeParameter string

// List of getTimeServicesPackageByCode_packageCode_parameter
const (
	GETTIMESERVICESPACKAGEBYCODEPACKAGECODEPARAMETER_NTP_STANDARD   GetTimeServicesPackageByCodePackageCodeParameter = "NTP_STANDARD"
	GETTIMESERVICESPACKAGEBYCODEPACKAGECODEPARAMETER_NTP_ENTERPRISE GetTimeServicesPackageByCodePackageCodeParameter = "NTP_ENTERPRISE"
	GETTIMESERVICESPACKAGEBYCODEPACKAGECODEPARAMETER_PTP_STANDARD   GetTimeServicesPackageByCodePackageCodeParameter = "PTP_STANDARD"
	GETTIMESERVICESPACKAGEBYCODEPACKAGECODEPARAMETER_PTP_ENTERPRISE GetTimeServicesPackageByCodePackageCodeParameter = "PTP_ENTERPRISE"
)

// All allowed values of GetTimeServicesPackageByCodePackageCodeParameter enum
var AllowedGetTimeServicesPackageByCodePackageCodeParameterEnumValues = []GetTimeServicesPackageByCodePackageCodeParameter{
	"NTP_STANDARD",
	"NTP_ENTERPRISE",
	"PTP_STANDARD",
	"PTP_ENTERPRISE",
}

func (v *GetTimeServicesPackageByCodePackageCodeParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GetTimeServicesPackageByCodePackageCodeParameter(value)
	for _, existing := range AllowedGetTimeServicesPackageByCodePackageCodeParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GetTimeServicesPackageByCodePackageCodeParameter", value)
}

// NewGetTimeServicesPackageByCodePackageCodeParameterFromValue returns a pointer to a valid GetTimeServicesPackageByCodePackageCodeParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGetTimeServicesPackageByCodePackageCodeParameterFromValue(v string) (*GetTimeServicesPackageByCodePackageCodeParameter, error) {
	ev := GetTimeServicesPackageByCodePackageCodeParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GetTimeServicesPackageByCodePackageCodeParameter: valid values are %v", v, AllowedGetTimeServicesPackageByCodePackageCodeParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GetTimeServicesPackageByCodePackageCodeParameter) IsValid() bool {
	for _, existing := range AllowedGetTimeServicesPackageByCodePackageCodeParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to getTimeServicesPackageByCode_packageCode_parameter value
func (v GetTimeServicesPackageByCodePackageCodeParameter) Ptr() *GetTimeServicesPackageByCodePackageCodeParameter {
	return &v
}

type NullableGetTimeServicesPackageByCodePackageCodeParameter struct {
	value *GetTimeServicesPackageByCodePackageCodeParameter
	isSet bool
}

func (v NullableGetTimeServicesPackageByCodePackageCodeParameter) Get() *GetTimeServicesPackageByCodePackageCodeParameter {
	return v.value
}

func (v *NullableGetTimeServicesPackageByCodePackageCodeParameter) Set(val *GetTimeServicesPackageByCodePackageCodeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTimeServicesPackageByCodePackageCodeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTimeServicesPackageByCodePackageCodeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTimeServicesPackageByCodePackageCodeParameter(val *GetTimeServicesPackageByCodePackageCodeParameter) *NullableGetTimeServicesPackageByCodePackageCodeParameter {
	return &NullableGetTimeServicesPackageByCodePackageCodeParameter{value: val, isSet: true}
}

func (v NullableGetTimeServicesPackageByCodePackageCodeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTimeServicesPackageByCodePackageCodeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
