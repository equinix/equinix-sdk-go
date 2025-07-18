/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// AWSProviderType the model 'AWSProviderType'
type AWSProviderType string

// List of AWSProvider_type
const (
	AWSPROVIDERTYPE_AWS_PROVIDER AWSProviderType = "AWS_PROVIDER"
)

// All allowed values of AWSProviderType enum
var AllowedAWSProviderTypeEnumValues = []AWSProviderType{
	"AWS_PROVIDER",
}

func (v *AWSProviderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AWSProviderType(value)
	for _, existing := range AllowedAWSProviderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AWSProviderType", value)
}

// NewAWSProviderTypeFromValue returns a pointer to a valid AWSProviderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAWSProviderTypeFromValue(v string) (*AWSProviderType, error) {
	ev := AWSProviderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AWSProviderType: valid values are %v", v, AllowedAWSProviderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AWSProviderType) IsValid() bool {
	for _, existing := range AllowedAWSProviderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AWSProvider_type value
func (v AWSProviderType) Ptr() *AWSProviderType {
	return &v
}

type NullableAWSProviderType struct {
	value *AWSProviderType
	isSet bool
}

func (v NullableAWSProviderType) Get() *AWSProviderType {
	return v.value
}

func (v *NullableAWSProviderType) Set(val *AWSProviderType) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSProviderType) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSProviderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSProviderType(val *AWSProviderType) *NullableAWSProviderType {
	return &NullableAWSProviderType{value: val, isSet: true}
}

func (v NullableAWSProviderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSProviderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
