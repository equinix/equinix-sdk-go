/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionRouteFiltersBaseDirection Route Filter direction to attach to a connection
type ConnectionRouteFiltersBaseDirection string

// List of ConnectionRouteFiltersBase_direction
const (
	CONNECTIONROUTEFILTERSBASEDIRECTION_INBOUND  ConnectionRouteFiltersBaseDirection = "INBOUND"
	CONNECTIONROUTEFILTERSBASEDIRECTION_OUTBOUND ConnectionRouteFiltersBaseDirection = "OUTBOUND"
)

// All allowed values of ConnectionRouteFiltersBaseDirection enum
var AllowedConnectionRouteFiltersBaseDirectionEnumValues = []ConnectionRouteFiltersBaseDirection{
	"INBOUND",
	"OUTBOUND",
}

func (v *ConnectionRouteFiltersBaseDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionRouteFiltersBaseDirection(value)
	for _, existing := range AllowedConnectionRouteFiltersBaseDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionRouteFiltersBaseDirection", value)
}

// NewConnectionRouteFiltersBaseDirectionFromValue returns a pointer to a valid ConnectionRouteFiltersBaseDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionRouteFiltersBaseDirectionFromValue(v string) (*ConnectionRouteFiltersBaseDirection, error) {
	ev := ConnectionRouteFiltersBaseDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionRouteFiltersBaseDirection: valid values are %v", v, AllowedConnectionRouteFiltersBaseDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionRouteFiltersBaseDirection) IsValid() bool {
	for _, existing := range AllowedConnectionRouteFiltersBaseDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionRouteFiltersBase_direction value
func (v ConnectionRouteFiltersBaseDirection) Ptr() *ConnectionRouteFiltersBaseDirection {
	return &v
}

type NullableConnectionRouteFiltersBaseDirection struct {
	value *ConnectionRouteFiltersBaseDirection
	isSet bool
}

func (v NullableConnectionRouteFiltersBaseDirection) Get() *ConnectionRouteFiltersBaseDirection {
	return v.value
}

func (v *NullableConnectionRouteFiltersBaseDirection) Set(val *ConnectionRouteFiltersBaseDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteFiltersBaseDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteFiltersBaseDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteFiltersBaseDirection(val *ConnectionRouteFiltersBaseDirection) *NullableConnectionRouteFiltersBaseDirection {
	return &NullableConnectionRouteFiltersBaseDirection{value: val, isSet: true}
}

func (v NullableConnectionRouteFiltersBaseDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteFiltersBaseDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
