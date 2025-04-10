/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// ServiceProfileAccessPointType - Access Point Type
type ServiceProfileAccessPointType struct {
	ServiceProfileAccessPointTypeCOLO *ServiceProfileAccessPointTypeCOLO
	ServiceProfileAccessPointTypeVD   *ServiceProfileAccessPointTypeVD
}

// ServiceProfileAccessPointTypeCOLOAsServiceProfileAccessPointType is a convenience function that returns ServiceProfileAccessPointTypeCOLO wrapped in ServiceProfileAccessPointType
func ServiceProfileAccessPointTypeCOLOAsServiceProfileAccessPointType(v *ServiceProfileAccessPointTypeCOLO) ServiceProfileAccessPointType {
	return ServiceProfileAccessPointType{
		ServiceProfileAccessPointTypeCOLO: v,
	}
}

// ServiceProfileAccessPointTypeVDAsServiceProfileAccessPointType is a convenience function that returns ServiceProfileAccessPointTypeVD wrapped in ServiceProfileAccessPointType
func ServiceProfileAccessPointTypeVDAsServiceProfileAccessPointType(v *ServiceProfileAccessPointTypeVD) ServiceProfileAccessPointType {
	return ServiceProfileAccessPointType{
		ServiceProfileAccessPointTypeVD: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ServiceProfileAccessPointType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ServiceProfileAccessPointTypeCOLO
	err = newStrictDecoder(data).Decode(&dst.ServiceProfileAccessPointTypeCOLO)
	if err == nil {
		jsonServiceProfileAccessPointTypeCOLO, _ := json.Marshal(dst.ServiceProfileAccessPointTypeCOLO)
		if string(jsonServiceProfileAccessPointTypeCOLO) == "{}" { // empty struct
			dst.ServiceProfileAccessPointTypeCOLO = nil
		} else {
			if err = validator.Validate(dst.ServiceProfileAccessPointTypeCOLO); err != nil {
				dst.ServiceProfileAccessPointTypeCOLO = nil
			} else {
				match++
			}
		}
	} else {
		dst.ServiceProfileAccessPointTypeCOLO = nil
	}

	// try to unmarshal data into ServiceProfileAccessPointTypeVD
	err = newStrictDecoder(data).Decode(&dst.ServiceProfileAccessPointTypeVD)
	if err == nil {
		jsonServiceProfileAccessPointTypeVD, _ := json.Marshal(dst.ServiceProfileAccessPointTypeVD)
		if string(jsonServiceProfileAccessPointTypeVD) == "{}" { // empty struct
			dst.ServiceProfileAccessPointTypeVD = nil
		} else {
			if err = validator.Validate(dst.ServiceProfileAccessPointTypeVD); err != nil {
				dst.ServiceProfileAccessPointTypeVD = nil
			} else {
				match++
			}
		}
	} else {
		dst.ServiceProfileAccessPointTypeVD = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ServiceProfileAccessPointTypeCOLO = nil
		dst.ServiceProfileAccessPointTypeVD = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ServiceProfileAccessPointType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ServiceProfileAccessPointType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ServiceProfileAccessPointType) MarshalJSON() ([]byte, error) {
	if src.ServiceProfileAccessPointTypeCOLO != nil {
		return json.Marshal(&src.ServiceProfileAccessPointTypeCOLO)
	}

	if src.ServiceProfileAccessPointTypeVD != nil {
		return json.Marshal(&src.ServiceProfileAccessPointTypeVD)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ServiceProfileAccessPointType) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ServiceProfileAccessPointTypeCOLO != nil {
		return obj.ServiceProfileAccessPointTypeCOLO
	}

	if obj.ServiceProfileAccessPointTypeVD != nil {
		return obj.ServiceProfileAccessPointTypeVD
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ServiceProfileAccessPointType) GetActualInstanceValue() interface{} {
	if obj.ServiceProfileAccessPointTypeCOLO != nil {
		return *obj.ServiceProfileAccessPointTypeCOLO
	}

	if obj.ServiceProfileAccessPointTypeVD != nil {
		return *obj.ServiceProfileAccessPointTypeVD
	}

	// all schemas are nil
	return nil
}

type NullableServiceProfileAccessPointType struct {
	value *ServiceProfileAccessPointType
	isSet bool
}

func (v NullableServiceProfileAccessPointType) Get() *ServiceProfileAccessPointType {
	return v.value
}

func (v *NullableServiceProfileAccessPointType) Set(val *ServiceProfileAccessPointType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointType(val *ServiceProfileAccessPointType) *NullableServiceProfileAccessPointType {
	return &NullableServiceProfileAccessPointType{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
