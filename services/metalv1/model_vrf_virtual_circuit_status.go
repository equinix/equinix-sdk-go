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

// VrfVirtualCircuitStatus The status changes of a VRF virtual circuit are generally the same as Virtual Circuits that aren't in a VRF. However, for VRF Virtual Circuits on Fabric VCs, the status will change to 'waiting_on_peering_details' once the Fabric service token associated with the virtual circuit has been redeemed on Fabric, and Metal has found the associated Fabric connection. At this point, users can update the subnet, MD5 password, customer IP and/or metal IP accordingly. For VRF Virtual Circuits on Dedicated Ports, we require all peering details to be set on creation of a VRF Virtual Circuit. The status will change to `changing_peering_details` whenever an active VRF Virtual Circuit has any of its peering details updated.
type VrfVirtualCircuitStatus string

// List of VrfVirtualCircuit_status
const (
	VRFVIRTUALCIRCUITSTATUS_PENDING                         VrfVirtualCircuitStatus = "pending"
	VRFVIRTUALCIRCUITSTATUS_WAITING_ON_PEERING_DETAILS      VrfVirtualCircuitStatus = "waiting_on_peering_details"
	VRFVIRTUALCIRCUITSTATUS_ACTIVATING                      VrfVirtualCircuitStatus = "activating"
	VRFVIRTUALCIRCUITSTATUS_CHANGING_PEERING_DETAILS        VrfVirtualCircuitStatus = "changing_peering_details"
	VRFVIRTUALCIRCUITSTATUS_DEACTIVATING                    VrfVirtualCircuitStatus = "deactivating"
	VRFVIRTUALCIRCUITSTATUS_DELETING                        VrfVirtualCircuitStatus = "deleting"
	VRFVIRTUALCIRCUITSTATUS_ACTIVE                          VrfVirtualCircuitStatus = "active"
	VRFVIRTUALCIRCUITSTATUS_EXPIRED                         VrfVirtualCircuitStatus = "expired"
	VRFVIRTUALCIRCUITSTATUS_ACTIVATION_FAILED               VrfVirtualCircuitStatus = "activation_failed"
	VRFVIRTUALCIRCUITSTATUS_CHANGING_PEERING_DETAILS_FAILED VrfVirtualCircuitStatus = "changing_peering_details_failed"
	VRFVIRTUALCIRCUITSTATUS_DEACTIVATION_FAILED             VrfVirtualCircuitStatus = "deactivation_failed"
	VRFVIRTUALCIRCUITSTATUS_DELETE_FAILED                   VrfVirtualCircuitStatus = "delete_failed"
)

// All allowed values of VrfVirtualCircuitStatus enum
var AllowedVrfVirtualCircuitStatusEnumValues = []VrfVirtualCircuitStatus{
	"pending",
	"waiting_on_peering_details",
	"activating",
	"changing_peering_details",
	"deactivating",
	"deleting",
	"active",
	"expired",
	"activation_failed",
	"changing_peering_details_failed",
	"deactivation_failed",
	"delete_failed",
}

func (v *VrfVirtualCircuitStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VrfVirtualCircuitStatus(value)
	for _, existing := range AllowedVrfVirtualCircuitStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VrfVirtualCircuitStatus", value)
}

// NewVrfVirtualCircuitStatusFromValue returns a pointer to a valid VrfVirtualCircuitStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVrfVirtualCircuitStatusFromValue(v string) (*VrfVirtualCircuitStatus, error) {
	ev := VrfVirtualCircuitStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VrfVirtualCircuitStatus: valid values are %v", v, AllowedVrfVirtualCircuitStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VrfVirtualCircuitStatus) IsValid() bool {
	for _, existing := range AllowedVrfVirtualCircuitStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VrfVirtualCircuit_status value
func (v VrfVirtualCircuitStatus) Ptr() *VrfVirtualCircuitStatus {
	return &v
}

type NullableVrfVirtualCircuitStatus struct {
	value *VrfVirtualCircuitStatus
	isSet bool
}

func (v NullableVrfVirtualCircuitStatus) Get() *VrfVirtualCircuitStatus {
	return v.value
}

func (v *NullableVrfVirtualCircuitStatus) Set(val *VrfVirtualCircuitStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVrfVirtualCircuitStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVrfVirtualCircuitStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVrfVirtualCircuitStatus(val *VrfVirtualCircuitStatus) *NullableVrfVirtualCircuitStatus {
	return &NullableVrfVirtualCircuitStatus{value: val, isSet: true}
}

func (v NullableVrfVirtualCircuitStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVrfVirtualCircuitStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
