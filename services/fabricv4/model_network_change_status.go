/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// NetworkChangeStatus Current outcome of the change flow
type NetworkChangeStatus string

// List of NetworkChangeStatus
const (
	NETWORKCHANGESTATUS_APPROVED               NetworkChangeStatus = "APPROVED"
	NETWORKCHANGESTATUS_COMPLETED              NetworkChangeStatus = "COMPLETED"
	NETWORKCHANGESTATUS_FAILED                 NetworkChangeStatus = "FAILED"
	NETWORKCHANGESTATUS_REJECTED               NetworkChangeStatus = "REJECTED"
	NETWORKCHANGESTATUS_REQUESTED              NetworkChangeStatus = "REQUESTED"
	NETWORKCHANGESTATUS_SUBMITTED_FOR_APPROVAL NetworkChangeStatus = "SUBMITTED_FOR_APPROVAL"
)

// All allowed values of NetworkChangeStatus enum
var AllowedNetworkChangeStatusEnumValues = []NetworkChangeStatus{
	"APPROVED",
	"COMPLETED",
	"FAILED",
	"REJECTED",
	"REQUESTED",
	"SUBMITTED_FOR_APPROVAL",
}

func (v *NetworkChangeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkChangeStatus(value)
	for _, existing := range AllowedNetworkChangeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkChangeStatus", value)
}

// NewNetworkChangeStatusFromValue returns a pointer to a valid NetworkChangeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkChangeStatusFromValue(v string) (*NetworkChangeStatus, error) {
	ev := NetworkChangeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkChangeStatus: valid values are %v", v, AllowedNetworkChangeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkChangeStatus) IsValid() bool {
	for _, existing := range AllowedNetworkChangeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkChangeStatus value
func (v NetworkChangeStatus) Ptr() *NetworkChangeStatus {
	return &v
}

type NullableNetworkChangeStatus struct {
	value *NetworkChangeStatus
	isSet bool
}

func (v NullableNetworkChangeStatus) Get() *NetworkChangeStatus {
	return v.value
}

func (v *NullableNetworkChangeStatus) Set(val *NetworkChangeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkChangeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkChangeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkChangeStatus(val *NetworkChangeStatus) *NullableNetworkChangeStatus {
	return &NullableNetworkChangeStatus{value: val, isSet: true}
}

func (v NullableNetworkChangeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkChangeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}