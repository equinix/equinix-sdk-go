/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionRouteFilterDataAttachmentStatus the model 'ConnectionRouteFilterDataAttachmentStatus'
type ConnectionRouteFilterDataAttachmentStatus string

// List of ConnectionRouteFilterData_attachmentStatus
const (
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_ATTACHING                 ConnectionRouteFilterDataAttachmentStatus = "ATTACHING"
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_ATTACHED                  ConnectionRouteFilterDataAttachmentStatus = "ATTACHED"
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_DETACHED                  ConnectionRouteFilterDataAttachmentStatus = "DETACHED"
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_DETACHING                 ConnectionRouteFilterDataAttachmentStatus = "DETACHING"
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_FAILED                    ConnectionRouteFilterDataAttachmentStatus = "FAILED"
	CONNECTIONROUTEFILTERDATAATTACHMENTSTATUS_PENDING_BGP_CONFIGURATION ConnectionRouteFilterDataAttachmentStatus = "PENDING_BGP_CONFIGURATION"
)

// All allowed values of ConnectionRouteFilterDataAttachmentStatus enum
var AllowedConnectionRouteFilterDataAttachmentStatusEnumValues = []ConnectionRouteFilterDataAttachmentStatus{
	"ATTACHING",
	"ATTACHED",
	"DETACHED",
	"DETACHING",
	"FAILED",
	"PENDING_BGP_CONFIGURATION",
}

func (v *ConnectionRouteFilterDataAttachmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionRouteFilterDataAttachmentStatus(value)
	for _, existing := range AllowedConnectionRouteFilterDataAttachmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionRouteFilterDataAttachmentStatus", value)
}

// NewConnectionRouteFilterDataAttachmentStatusFromValue returns a pointer to a valid ConnectionRouteFilterDataAttachmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionRouteFilterDataAttachmentStatusFromValue(v string) (*ConnectionRouteFilterDataAttachmentStatus, error) {
	ev := ConnectionRouteFilterDataAttachmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionRouteFilterDataAttachmentStatus: valid values are %v", v, AllowedConnectionRouteFilterDataAttachmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionRouteFilterDataAttachmentStatus) IsValid() bool {
	for _, existing := range AllowedConnectionRouteFilterDataAttachmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionRouteFilterData_attachmentStatus value
func (v ConnectionRouteFilterDataAttachmentStatus) Ptr() *ConnectionRouteFilterDataAttachmentStatus {
	return &v
}

type NullableConnectionRouteFilterDataAttachmentStatus struct {
	value *ConnectionRouteFilterDataAttachmentStatus
	isSet bool
}

func (v NullableConnectionRouteFilterDataAttachmentStatus) Get() *ConnectionRouteFilterDataAttachmentStatus {
	return v.value
}

func (v *NullableConnectionRouteFilterDataAttachmentStatus) Set(val *ConnectionRouteFilterDataAttachmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteFilterDataAttachmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteFilterDataAttachmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteFilterDataAttachmentStatus(val *ConnectionRouteFilterDataAttachmentStatus) *NullableConnectionRouteFilterDataAttachmentStatus {
	return &NullableConnectionRouteFilterDataAttachmentStatus{value: val, isSet: true}
}

func (v NullableConnectionRouteFilterDataAttachmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteFilterDataAttachmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
