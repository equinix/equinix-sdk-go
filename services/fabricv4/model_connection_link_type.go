/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionLinkType the model 'ConnectionLinkType'
type ConnectionLinkType string

// List of connectionLink_type
const (
	CONNECTIONLINKTYPE_EVPL_VC ConnectionLinkType = "EVPL_VC"
)

// All allowed values of ConnectionLinkType enum
var AllowedConnectionLinkTypeEnumValues = []ConnectionLinkType{
	"EVPL_VC",
}

func (v *ConnectionLinkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionLinkType(value)
	for _, existing := range AllowedConnectionLinkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionLinkType", value)
}

// NewConnectionLinkTypeFromValue returns a pointer to a valid ConnectionLinkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionLinkTypeFromValue(v string) (*ConnectionLinkType, error) {
	ev := ConnectionLinkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionLinkType: valid values are %v", v, AllowedConnectionLinkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionLinkType) IsValid() bool {
	for _, existing := range AllowedConnectionLinkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to connectionLink_type value
func (v ConnectionLinkType) Ptr() *ConnectionLinkType {
	return &v
}

type NullableConnectionLinkType struct {
	value *ConnectionLinkType
	isSet bool
}

func (v NullableConnectionLinkType) Get() *ConnectionLinkType {
	return v.value
}

func (v *NullableConnectionLinkType) Set(val *ConnectionLinkType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionLinkType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionLinkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionLinkType(val *ConnectionLinkType) *NullableConnectionLinkType {
	return &NullableConnectionLinkType{value: val, isSet: true}
}

func (v NullableConnectionLinkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionLinkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}