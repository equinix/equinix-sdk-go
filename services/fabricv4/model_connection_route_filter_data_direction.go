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

// ConnectionRouteFilterDataDirection the model 'ConnectionRouteFilterDataDirection'
type ConnectionRouteFilterDataDirection string

// List of ConnectionRouteFilterData_direction
const (
	CONNECTIONROUTEFILTERDATADIRECTION_INBOUND  ConnectionRouteFilterDataDirection = "INBOUND"
	CONNECTIONROUTEFILTERDATADIRECTION_OUTBOUND ConnectionRouteFilterDataDirection = "OUTBOUND"
)

// All allowed values of ConnectionRouteFilterDataDirection enum
var AllowedConnectionRouteFilterDataDirectionEnumValues = []ConnectionRouteFilterDataDirection{
	"INBOUND",
	"OUTBOUND",
}

func (v *ConnectionRouteFilterDataDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionRouteFilterDataDirection(value)
	for _, existing := range AllowedConnectionRouteFilterDataDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionRouteFilterDataDirection", value)
}

// NewConnectionRouteFilterDataDirectionFromValue returns a pointer to a valid ConnectionRouteFilterDataDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionRouteFilterDataDirectionFromValue(v string) (*ConnectionRouteFilterDataDirection, error) {
	ev := ConnectionRouteFilterDataDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionRouteFilterDataDirection: valid values are %v", v, AllowedConnectionRouteFilterDataDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionRouteFilterDataDirection) IsValid() bool {
	for _, existing := range AllowedConnectionRouteFilterDataDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionRouteFilterData_direction value
func (v ConnectionRouteFilterDataDirection) Ptr() *ConnectionRouteFilterDataDirection {
	return &v
}

type NullableConnectionRouteFilterDataDirection struct {
	value *ConnectionRouteFilterDataDirection
	isSet bool
}

func (v NullableConnectionRouteFilterDataDirection) Get() *ConnectionRouteFilterDataDirection {
	return v.value
}

func (v *NullableConnectionRouteFilterDataDirection) Set(val *ConnectionRouteFilterDataDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteFilterDataDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteFilterDataDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteFilterDataDirection(val *ConnectionRouteFilterDataDirection) *NullableConnectionRouteFilterDataDirection {
	return &NullableConnectionRouteFilterDataDirection{value: val, isSet: true}
}

func (v NullableConnectionRouteFilterDataDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteFilterDataDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}