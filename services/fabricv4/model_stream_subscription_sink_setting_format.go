/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// StreamSubscriptionSinkSettingFormat webhook message format
type StreamSubscriptionSinkSettingFormat string

// List of StreamSubscriptionSinkSetting_format
const (
	STREAMSUBSCRIPTIONSINKSETTINGFORMAT_CLOUDEVENT    StreamSubscriptionSinkSettingFormat = "CLOUDEVENT"
	STREAMSUBSCRIPTIONSINKSETTINGFORMAT_OPENTELEMETRY StreamSubscriptionSinkSettingFormat = "OPENTELEMETRY"
)

// All allowed values of StreamSubscriptionSinkSettingFormat enum
var AllowedStreamSubscriptionSinkSettingFormatEnumValues = []StreamSubscriptionSinkSettingFormat{
	"CLOUDEVENT",
	"OPENTELEMETRY",
}

func (v *StreamSubscriptionSinkSettingFormat) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StreamSubscriptionSinkSettingFormat(value)
	for _, existing := range AllowedStreamSubscriptionSinkSettingFormatEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StreamSubscriptionSinkSettingFormat", value)
}

// NewStreamSubscriptionSinkSettingFormatFromValue returns a pointer to a valid StreamSubscriptionSinkSettingFormat
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStreamSubscriptionSinkSettingFormatFromValue(v string) (*StreamSubscriptionSinkSettingFormat, error) {
	ev := StreamSubscriptionSinkSettingFormat(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StreamSubscriptionSinkSettingFormat: valid values are %v", v, AllowedStreamSubscriptionSinkSettingFormatEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StreamSubscriptionSinkSettingFormat) IsValid() bool {
	for _, existing := range AllowedStreamSubscriptionSinkSettingFormatEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StreamSubscriptionSinkSetting_format value
func (v StreamSubscriptionSinkSettingFormat) Ptr() *StreamSubscriptionSinkSettingFormat {
	return &v
}

type NullableStreamSubscriptionSinkSettingFormat struct {
	value *StreamSubscriptionSinkSettingFormat
	isSet bool
}

func (v NullableStreamSubscriptionSinkSettingFormat) Get() *StreamSubscriptionSinkSettingFormat {
	return v.value
}

func (v *NullableStreamSubscriptionSinkSettingFormat) Set(val *StreamSubscriptionSinkSettingFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamSubscriptionSinkSettingFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamSubscriptionSinkSettingFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamSubscriptionSinkSettingFormat(val *StreamSubscriptionSinkSettingFormat) *NullableStreamSubscriptionSinkSettingFormat {
	return &NullableStreamSubscriptionSinkSettingFormat{value: val, isSet: true}
}

func (v NullableStreamSubscriptionSinkSettingFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamSubscriptionSinkSettingFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
