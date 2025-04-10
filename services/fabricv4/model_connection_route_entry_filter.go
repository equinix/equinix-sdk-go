/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ConnectionRouteEntryFilter struct for ConnectionRouteEntryFilter
type ConnectionRouteEntryFilter struct {
	ConnectionRouteEntryOrFilter         *ConnectionRouteEntryOrFilter
	ConnectionRouteEntrySimpleExpression *ConnectionRouteEntrySimpleExpression
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ConnectionRouteEntryFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ConnectionRouteEntryOrFilter
	err = json.Unmarshal(data, &dst.ConnectionRouteEntryOrFilter)
	if err == nil {
		jsonConnectionRouteEntryOrFilter, _ := json.Marshal(dst.ConnectionRouteEntryOrFilter)
		if string(jsonConnectionRouteEntryOrFilter) == "{}" { // empty struct
			dst.ConnectionRouteEntryOrFilter = nil
		} else {
			return nil // data stored in dst.ConnectionRouteEntryOrFilter, return on the first match
		}
	} else {
		dst.ConnectionRouteEntryOrFilter = nil
	}

	// try to unmarshal JSON data into ConnectionRouteEntrySimpleExpression
	err = json.Unmarshal(data, &dst.ConnectionRouteEntrySimpleExpression)
	if err == nil {
		jsonConnectionRouteEntrySimpleExpression, _ := json.Marshal(dst.ConnectionRouteEntrySimpleExpression)
		if string(jsonConnectionRouteEntrySimpleExpression) == "{}" { // empty struct
			dst.ConnectionRouteEntrySimpleExpression = nil
		} else {
			return nil // data stored in dst.ConnectionRouteEntrySimpleExpression, return on the first match
		}
	} else {
		dst.ConnectionRouteEntrySimpleExpression = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ConnectionRouteEntryFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConnectionRouteEntryFilter) MarshalJSON() ([]byte, error) {
	if src.ConnectionRouteEntryOrFilter != nil {
		return json.Marshal(&src.ConnectionRouteEntryOrFilter)
	}

	if src.ConnectionRouteEntrySimpleExpression != nil {
		return json.Marshal(&src.ConnectionRouteEntrySimpleExpression)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConnectionRouteEntryFilter struct {
	value *ConnectionRouteEntryFilter
	isSet bool
}

func (v NullableConnectionRouteEntryFilter) Get() *ConnectionRouteEntryFilter {
	return v.value
}

func (v *NullableConnectionRouteEntryFilter) Set(val *ConnectionRouteEntryFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionRouteEntryFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionRouteEntryFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionRouteEntryFilter(val *ConnectionRouteEntryFilter) *NullableConnectionRouteEntryFilter {
	return &NullableConnectionRouteEntryFilter{value: val, isSet: true}
}

func (v NullableConnectionRouteEntryFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionRouteEntryFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
