/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// SubscriptionAssetType the model 'SubscriptionAssetType'
type SubscriptionAssetType string

// List of SubscriptionAssetType
const (
	SUBSCRIPTIONASSETTYPE_XF_ROUTER SubscriptionAssetType = "XF_ROUTER"
	SUBSCRIPTIONASSETTYPE_IP_VC     SubscriptionAssetType = "IP_VC"
)

// All allowed values of SubscriptionAssetType enum
var AllowedSubscriptionAssetTypeEnumValues = []SubscriptionAssetType{
	"XF_ROUTER",
	"IP_VC",
}

func (v *SubscriptionAssetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionAssetType(value)
	for _, existing := range AllowedSubscriptionAssetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionAssetType", value)
}

// NewSubscriptionAssetTypeFromValue returns a pointer to a valid SubscriptionAssetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionAssetTypeFromValue(v string) (*SubscriptionAssetType, error) {
	ev := SubscriptionAssetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionAssetType: valid values are %v", v, AllowedSubscriptionAssetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionAssetType) IsValid() bool {
	for _, existing := range AllowedSubscriptionAssetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionAssetType value
func (v SubscriptionAssetType) Ptr() *SubscriptionAssetType {
	return &v
}

type NullableSubscriptionAssetType struct {
	value *SubscriptionAssetType
	isSet bool
}

func (v NullableSubscriptionAssetType) Get() *SubscriptionAssetType {
	return v.value
}

func (v *NullableSubscriptionAssetType) Set(val *SubscriptionAssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAssetType(val *SubscriptionAssetType) *NullableSubscriptionAssetType {
	return &NullableSubscriptionAssetType{value: val, isSet: true}
}

func (v NullableSubscriptionAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
