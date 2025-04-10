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

// RoutingProtocolBase - struct for RoutingProtocolBase
type RoutingProtocolBase struct {
	RoutingProtocolBGPType    *RoutingProtocolBGPType
	RoutingProtocolDirectType *RoutingProtocolDirectType
}

// RoutingProtocolBGPTypeAsRoutingProtocolBase is a convenience function that returns RoutingProtocolBGPType wrapped in RoutingProtocolBase
func RoutingProtocolBGPTypeAsRoutingProtocolBase(v *RoutingProtocolBGPType) RoutingProtocolBase {
	return RoutingProtocolBase{
		RoutingProtocolBGPType: v,
	}
}

// RoutingProtocolDirectTypeAsRoutingProtocolBase is a convenience function that returns RoutingProtocolDirectType wrapped in RoutingProtocolBase
func RoutingProtocolDirectTypeAsRoutingProtocolBase(v *RoutingProtocolDirectType) RoutingProtocolBase {
	return RoutingProtocolBase{
		RoutingProtocolDirectType: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RoutingProtocolBase) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RoutingProtocolBGPType
	err = newStrictDecoder(data).Decode(&dst.RoutingProtocolBGPType)
	if err == nil {
		jsonRoutingProtocolBGPType, _ := json.Marshal(dst.RoutingProtocolBGPType)
		if string(jsonRoutingProtocolBGPType) == "{}" { // empty struct
			dst.RoutingProtocolBGPType = nil
		} else {
			if err = validator.Validate(dst.RoutingProtocolBGPType); err != nil {
				dst.RoutingProtocolBGPType = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoutingProtocolBGPType = nil
	}

	// try to unmarshal data into RoutingProtocolDirectType
	err = newStrictDecoder(data).Decode(&dst.RoutingProtocolDirectType)
	if err == nil {
		jsonRoutingProtocolDirectType, _ := json.Marshal(dst.RoutingProtocolDirectType)
		if string(jsonRoutingProtocolDirectType) == "{}" { // empty struct
			dst.RoutingProtocolDirectType = nil
		} else {
			if err = validator.Validate(dst.RoutingProtocolDirectType); err != nil {
				dst.RoutingProtocolDirectType = nil
			} else {
				match++
			}
		}
	} else {
		dst.RoutingProtocolDirectType = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RoutingProtocolBGPType = nil
		dst.RoutingProtocolDirectType = nil

		return fmt.Errorf("data matches more than one schema in oneOf(RoutingProtocolBase)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RoutingProtocolBase)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RoutingProtocolBase) MarshalJSON() ([]byte, error) {
	if src.RoutingProtocolBGPType != nil {
		return json.Marshal(&src.RoutingProtocolBGPType)
	}

	if src.RoutingProtocolDirectType != nil {
		return json.Marshal(&src.RoutingProtocolDirectType)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RoutingProtocolBase) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.RoutingProtocolBGPType != nil {
		return obj.RoutingProtocolBGPType
	}

	if obj.RoutingProtocolDirectType != nil {
		return obj.RoutingProtocolDirectType
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj RoutingProtocolBase) GetActualInstanceValue() interface{} {
	if obj.RoutingProtocolBGPType != nil {
		return *obj.RoutingProtocolBGPType
	}

	if obj.RoutingProtocolDirectType != nil {
		return *obj.RoutingProtocolDirectType
	}

	// all schemas are nil
	return nil
}

type NullableRoutingProtocolBase struct {
	value *RoutingProtocolBase
	isSet bool
}

func (v NullableRoutingProtocolBase) Get() *RoutingProtocolBase {
	return v.value
}

func (v *NullableRoutingProtocolBase) Set(val *RoutingProtocolBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolBase(val *RoutingProtocolBase) *NullableRoutingProtocolBase {
	return &NullableRoutingProtocolBase{value: val, isSet: true}
}

func (v NullableRoutingProtocolBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
