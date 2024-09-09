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

// SubscriptionResponseMarketplace the model 'SubscriptionResponseMarketplace'
type SubscriptionResponseMarketplace string

// List of SubscriptionResponse_marketplace
const (
	SUBSCRIPTIONRESPONSEMARKETPLACE_AWS    SubscriptionResponseMarketplace = "AWS"
	SUBSCRIPTIONRESPONSEMARKETPLACE_GCP    SubscriptionResponseMarketplace = "GCP"
	SUBSCRIPTIONRESPONSEMARKETPLACE_AZURE  SubscriptionResponseMarketplace = "AZURE"
	SUBSCRIPTIONRESPONSEMARKETPLACE_REDHAT SubscriptionResponseMarketplace = "REDHAT"
)

// All allowed values of SubscriptionResponseMarketplace enum
var AllowedSubscriptionResponseMarketplaceEnumValues = []SubscriptionResponseMarketplace{
	"AWS",
	"GCP",
	"AZURE",
	"REDHAT",
}

func (v *SubscriptionResponseMarketplace) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscriptionResponseMarketplace(value)
	for _, existing := range AllowedSubscriptionResponseMarketplaceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscriptionResponseMarketplace", value)
}

// NewSubscriptionResponseMarketplaceFromValue returns a pointer to a valid SubscriptionResponseMarketplace
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscriptionResponseMarketplaceFromValue(v string) (*SubscriptionResponseMarketplace, error) {
	ev := SubscriptionResponseMarketplace(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscriptionResponseMarketplace: valid values are %v", v, AllowedSubscriptionResponseMarketplaceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscriptionResponseMarketplace) IsValid() bool {
	for _, existing := range AllowedSubscriptionResponseMarketplaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscriptionResponse_marketplace value
func (v SubscriptionResponseMarketplace) Ptr() *SubscriptionResponseMarketplace {
	return &v
}

type NullableSubscriptionResponseMarketplace struct {
	value *SubscriptionResponseMarketplace
	isSet bool
}

func (v NullableSubscriptionResponseMarketplace) Get() *SubscriptionResponseMarketplace {
	return v.value
}

func (v *NullableSubscriptionResponseMarketplace) Set(val *SubscriptionResponseMarketplace) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionResponseMarketplace) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionResponseMarketplace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionResponseMarketplace(val *SubscriptionResponseMarketplace) *NullableSubscriptionResponseMarketplace {
	return &NullableSubscriptionResponseMarketplace{value: val, isSet: true}
}

func (v NullableSubscriptionResponseMarketplace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionResponseMarketplace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}