/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionRouteAggregationDataType Route Aggregation type
type ConnectionRouteAggregationDataType string

// List of ConnectionRouteAggregationData_type
const (
	CONNECTIONROUTEAGGREGATIONDATATYPE_BGP_IPV4_PREFIX_AGGREGATION ConnectionRouteAggregationDataType = "BGP_IPv4_PREFIX_AGGREGATION"
)

// All allowed values of ConnectionRouteAggregationDataType enum
var AllowedConnectionRouteAggregationDataTypeEnumValues = []ConnectionRouteAggregationDataType{
	"BGP_IPv4_PREFIX_AGGREGATION",
}

func (v *ConnectionRouteAggregationDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionRouteAggregationDataType(value)
	for _, existing := range AllowedConnectionRouteAggregationDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionRouteAggregationDataType", value)
}

// NewConnectionRouteAggregationDataTypeFromValue returns a pointer to a valid ConnectionRouteAggregationDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionRouteAggregationDataTypeFromValue(v string) (*ConnectionRouteAggregationDataType, error) {
	ev := ConnectionRouteAggregationDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionRouteAggregationDataType: valid values are %v", v, AllowedConnectionRouteAggregationDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionRouteAggregationDataType) IsValid() bool {
	for _, existing := range AllowedConnectionRouteAggregationDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionRouteAggregationData_type value
func (v ConnectionRouteAggregationDataType) Ptr() *ConnectionRouteAggregationDataType {
	return &v
}

type NullableConnectionRouteAggregationDataType struct {
	value *ConnectionRouteAggregationDataType
	isSet bool
}

func (v NullableConnectionRouteAggregationDataType) Get() *ConnectionRouteAggregationDataType {
	return v.value
}

func (v *NullableConnectionRouteAggregationDataType) Set(val *ConnectionRouteAggregationDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteAggregationDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteAggregationDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteAggregationDataType(val *ConnectionRouteAggregationDataType) *NullableConnectionRouteAggregationDataType {
	return &NullableConnectionRouteAggregationDataType{value: val, isSet: true}
}

func (v NullableConnectionRouteAggregationDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteAggregationDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}