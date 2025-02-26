/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PrecisionTimeChangeOperationOp Handy shortcut for operation name
type PrecisionTimeChangeOperationOp string

// List of precisionTimeChangeOperation_op
const (
	PRECISIONTIMECHANGEOPERATIONOP_REPLACE PrecisionTimeChangeOperationOp = "replace"
	PRECISIONTIMECHANGEOPERATIONOP_ADD     PrecisionTimeChangeOperationOp = "add"
	PRECISIONTIMECHANGEOPERATIONOP_REMOVE  PrecisionTimeChangeOperationOp = "remove"
)

// All allowed values of PrecisionTimeChangeOperationOp enum
var AllowedPrecisionTimeChangeOperationOpEnumValues = []PrecisionTimeChangeOperationOp{
	"replace",
	"add",
	"remove",
}

func (v *PrecisionTimeChangeOperationOp) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrecisionTimeChangeOperationOp(value)
	for _, existing := range AllowedPrecisionTimeChangeOperationOpEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrecisionTimeChangeOperationOp", value)
}

// NewPrecisionTimeChangeOperationOpFromValue returns a pointer to a valid PrecisionTimeChangeOperationOp
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrecisionTimeChangeOperationOpFromValue(v string) (*PrecisionTimeChangeOperationOp, error) {
	ev := PrecisionTimeChangeOperationOp(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrecisionTimeChangeOperationOp: valid values are %v", v, AllowedPrecisionTimeChangeOperationOpEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrecisionTimeChangeOperationOp) IsValid() bool {
	for _, existing := range AllowedPrecisionTimeChangeOperationOpEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to precisionTimeChangeOperation_op value
func (v PrecisionTimeChangeOperationOp) Ptr() *PrecisionTimeChangeOperationOp {
	return &v
}

type NullablePrecisionTimeChangeOperationOp struct {
	value *PrecisionTimeChangeOperationOp
	isSet bool
}

func (v NullablePrecisionTimeChangeOperationOp) Get() *PrecisionTimeChangeOperationOp {
	return v.value
}

func (v *NullablePrecisionTimeChangeOperationOp) Set(val *PrecisionTimeChangeOperationOp) {
	v.value = val
	v.isSet = true
}

func (v NullablePrecisionTimeChangeOperationOp) IsSet() bool {
	return v.isSet
}

func (v *NullablePrecisionTimeChangeOperationOp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrecisionTimeChangeOperationOp(val *PrecisionTimeChangeOperationOp) *NullablePrecisionTimeChangeOperationOp {
	return &NullablePrecisionTimeChangeOperationOp{value: val, isSet: true}
}

func (v NullablePrecisionTimeChangeOperationOp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrecisionTimeChangeOperationOp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
