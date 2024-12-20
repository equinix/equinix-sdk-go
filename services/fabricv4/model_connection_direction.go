/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionDirection Connection directionality from the requester point of view
type ConnectionDirection string

// List of ConnectionDirection
const (
	CONNECTIONDIRECTION_INTERNAL ConnectionDirection = "INTERNAL"
	CONNECTIONDIRECTION_INCOMING ConnectionDirection = "INCOMING"
	CONNECTIONDIRECTION_OUTGOING ConnectionDirection = "OUTGOING"
)

// All allowed values of ConnectionDirection enum
var AllowedConnectionDirectionEnumValues = []ConnectionDirection{
	"INTERNAL",
	"INCOMING",
	"OUTGOING",
}

func (v *ConnectionDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionDirection(value)
	for _, existing := range AllowedConnectionDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionDirection", value)
}

// NewConnectionDirectionFromValue returns a pointer to a valid ConnectionDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionDirectionFromValue(v string) (*ConnectionDirection, error) {
	ev := ConnectionDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionDirection: valid values are %v", v, AllowedConnectionDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionDirection) IsValid() bool {
	for _, existing := range AllowedConnectionDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionDirection value
func (v ConnectionDirection) Ptr() *ConnectionDirection {
	return &v
}

type NullableConnectionDirection struct {
	value *ConnectionDirection
	isSet bool
}

func (v NullableConnectionDirection) Get() *ConnectionDirection {
	return v.value
}

func (v *NullableConnectionDirection) Set(val *ConnectionDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDirection(val *ConnectionDirection) *NullableConnectionDirection {
	return &NullableConnectionDirection{value: val, isSet: true}
}

func (v NullableConnectionDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
