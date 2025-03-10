/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// NetworkSortDirectionResponse Sorting direction
type NetworkSortDirectionResponse string

// List of NetworkSortDirectionResponse
const (
	NETWORKSORTDIRECTIONRESPONSE_DESC NetworkSortDirectionResponse = "DESC"
	NETWORKSORTDIRECTIONRESPONSE_ASC  NetworkSortDirectionResponse = "ASC"
)

// All allowed values of NetworkSortDirectionResponse enum
var AllowedNetworkSortDirectionResponseEnumValues = []NetworkSortDirectionResponse{
	"DESC",
	"ASC",
}

func (v *NetworkSortDirectionResponse) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkSortDirectionResponse(value)
	for _, existing := range AllowedNetworkSortDirectionResponseEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkSortDirectionResponse", value)
}

// NewNetworkSortDirectionResponseFromValue returns a pointer to a valid NetworkSortDirectionResponse
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkSortDirectionResponseFromValue(v string) (*NetworkSortDirectionResponse, error) {
	ev := NetworkSortDirectionResponse(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NetworkSortDirectionResponse: valid values are %v", v, AllowedNetworkSortDirectionResponseEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkSortDirectionResponse) IsValid() bool {
	for _, existing := range AllowedNetworkSortDirectionResponseEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NetworkSortDirectionResponse value
func (v NetworkSortDirectionResponse) Ptr() *NetworkSortDirectionResponse {
	return &v
}

type NullableNetworkSortDirectionResponse struct {
	value *NetworkSortDirectionResponse
	isSet bool
}

func (v NullableNetworkSortDirectionResponse) Get() *NetworkSortDirectionResponse {
	return v.value
}

func (v *NullableNetworkSortDirectionResponse) Set(val *NetworkSortDirectionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSortDirectionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSortDirectionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSortDirectionResponse(val *NetworkSortDirectionResponse) *NullableNetworkSortDirectionResponse {
	return &NullableNetworkSortDirectionResponse{value: val, isSet: true}
}

func (v NullableNetworkSortDirectionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSortDirectionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
