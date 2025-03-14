/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PtpAdvanceConfigurationTimeScale Time Scale value, ARB denotes Arbitrary and PTP denotes Precision Time Protocol.
type PtpAdvanceConfigurationTimeScale string

// List of ptpAdvanceConfiguration_timeScale
const (
	PTPADVANCECONFIGURATIONTIMESCALE_ARB PtpAdvanceConfigurationTimeScale = "ARB"
	PTPADVANCECONFIGURATIONTIMESCALE_PTP PtpAdvanceConfigurationTimeScale = "PTP"
)

// All allowed values of PtpAdvanceConfigurationTimeScale enum
var AllowedPtpAdvanceConfigurationTimeScaleEnumValues = []PtpAdvanceConfigurationTimeScale{
	"ARB",
	"PTP",
}

func (v *PtpAdvanceConfigurationTimeScale) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PtpAdvanceConfigurationTimeScale(value)
	for _, existing := range AllowedPtpAdvanceConfigurationTimeScaleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PtpAdvanceConfigurationTimeScale", value)
}

// NewPtpAdvanceConfigurationTimeScaleFromValue returns a pointer to a valid PtpAdvanceConfigurationTimeScale
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPtpAdvanceConfigurationTimeScaleFromValue(v string) (*PtpAdvanceConfigurationTimeScale, error) {
	ev := PtpAdvanceConfigurationTimeScale(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PtpAdvanceConfigurationTimeScale: valid values are %v", v, AllowedPtpAdvanceConfigurationTimeScaleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PtpAdvanceConfigurationTimeScale) IsValid() bool {
	for _, existing := range AllowedPtpAdvanceConfigurationTimeScaleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ptpAdvanceConfiguration_timeScale value
func (v PtpAdvanceConfigurationTimeScale) Ptr() *PtpAdvanceConfigurationTimeScale {
	return &v
}

type NullablePtpAdvanceConfigurationTimeScale struct {
	value *PtpAdvanceConfigurationTimeScale
	isSet bool
}

func (v NullablePtpAdvanceConfigurationTimeScale) Get() *PtpAdvanceConfigurationTimeScale {
	return v.value
}

func (v *NullablePtpAdvanceConfigurationTimeScale) Set(val *PtpAdvanceConfigurationTimeScale) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpAdvanceConfigurationTimeScale) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpAdvanceConfigurationTimeScale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpAdvanceConfigurationTimeScale(val *PtpAdvanceConfigurationTimeScale) *NullablePtpAdvanceConfigurationTimeScale {
	return &NullablePtpAdvanceConfigurationTimeScale{value: val, isSet: true}
}

func (v NullablePtpAdvanceConfigurationTimeScale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpAdvanceConfigurationTimeScale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
