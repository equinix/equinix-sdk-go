/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// CloudRouterPostRequestType the model 'CloudRouterPostRequestType'
type CloudRouterPostRequestType string

// List of CloudRouterPostRequest_type
const (
	CLOUDROUTERPOSTREQUESTTYPE_XF_ROUTER CloudRouterPostRequestType = "XF_ROUTER"
)

// All allowed values of CloudRouterPostRequestType enum
var AllowedCloudRouterPostRequestTypeEnumValues = []CloudRouterPostRequestType{
	"XF_ROUTER",
}

func (v *CloudRouterPostRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudRouterPostRequestType(value)
	for _, existing := range AllowedCloudRouterPostRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudRouterPostRequestType", value)
}

// NewCloudRouterPostRequestTypeFromValue returns a pointer to a valid CloudRouterPostRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudRouterPostRequestTypeFromValue(v string) (*CloudRouterPostRequestType, error) {
	ev := CloudRouterPostRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudRouterPostRequestType: valid values are %v", v, AllowedCloudRouterPostRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudRouterPostRequestType) IsValid() bool {
	for _, existing := range AllowedCloudRouterPostRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudRouterPostRequest_type value
func (v CloudRouterPostRequestType) Ptr() *CloudRouterPostRequestType {
	return &v
}

type NullableCloudRouterPostRequestType struct {
	value *CloudRouterPostRequestType
	isSet bool
}

func (v NullableCloudRouterPostRequestType) Get() *CloudRouterPostRequestType {
	return v.value
}

func (v *NullableCloudRouterPostRequestType) Set(val *CloudRouterPostRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterPostRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterPostRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterPostRequestType(val *CloudRouterPostRequestType) *NullableCloudRouterPostRequestType {
	return &NullableCloudRouterPostRequestType{value: val, isSet: true}
}

func (v NullableCloudRouterPostRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterPostRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
