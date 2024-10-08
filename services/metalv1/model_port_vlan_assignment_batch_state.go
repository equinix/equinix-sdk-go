/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// PortVlanAssignmentBatchState the model 'PortVlanAssignmentBatchState'
type PortVlanAssignmentBatchState string

// List of PortVlanAssignmentBatch_state
const (
	PORTVLANASSIGNMENTBATCHSTATE_QUEUED      PortVlanAssignmentBatchState = "queued"
	PORTVLANASSIGNMENTBATCHSTATE_IN_PROGRESS PortVlanAssignmentBatchState = "in_progress"
	PORTVLANASSIGNMENTBATCHSTATE_COMPLETED   PortVlanAssignmentBatchState = "completed"
	PORTVLANASSIGNMENTBATCHSTATE_FAILED      PortVlanAssignmentBatchState = "failed"
)

// All allowed values of PortVlanAssignmentBatchState enum
var AllowedPortVlanAssignmentBatchStateEnumValues = []PortVlanAssignmentBatchState{
	"queued",
	"in_progress",
	"completed",
	"failed",
}

func (v *PortVlanAssignmentBatchState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortVlanAssignmentBatchState(value)
	for _, existing := range AllowedPortVlanAssignmentBatchStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortVlanAssignmentBatchState", value)
}

// NewPortVlanAssignmentBatchStateFromValue returns a pointer to a valid PortVlanAssignmentBatchState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortVlanAssignmentBatchStateFromValue(v string) (*PortVlanAssignmentBatchState, error) {
	ev := PortVlanAssignmentBatchState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortVlanAssignmentBatchState: valid values are %v", v, AllowedPortVlanAssignmentBatchStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortVlanAssignmentBatchState) IsValid() bool {
	for _, existing := range AllowedPortVlanAssignmentBatchStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortVlanAssignmentBatch_state value
func (v PortVlanAssignmentBatchState) Ptr() *PortVlanAssignmentBatchState {
	return &v
}

type NullablePortVlanAssignmentBatchState struct {
	value *PortVlanAssignmentBatchState
	isSet bool
}

func (v NullablePortVlanAssignmentBatchState) Get() *PortVlanAssignmentBatchState {
	return v.value
}

func (v *NullablePortVlanAssignmentBatchState) Set(val *PortVlanAssignmentBatchState) {
	v.value = val
	v.isSet = true
}

func (v NullablePortVlanAssignmentBatchState) IsSet() bool {
	return v.isSet
}

func (v *NullablePortVlanAssignmentBatchState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortVlanAssignmentBatchState(val *PortVlanAssignmentBatchState) *NullablePortVlanAssignmentBatchState {
	return &NullablePortVlanAssignmentBatchState{value: val, isSet: true}
}

func (v NullablePortVlanAssignmentBatchState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortVlanAssignmentBatchState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
