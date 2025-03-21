/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PtpAdvanceConfigurationLogDelayReqInterval Logarithmic value that controls the rate of PTP DelayReq packets. Default is -4 (16 packets per second), Unit packets/second..
type PtpAdvanceConfigurationLogDelayReqInterval int32

// List of ptpAdvanceConfiguration_logDelayReqInterval
const (
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__MINUS_5 PtpAdvanceConfigurationLogDelayReqInterval = -5
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__MINUS_4 PtpAdvanceConfigurationLogDelayReqInterval = -4
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__MINUS_3 PtpAdvanceConfigurationLogDelayReqInterval = -3
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__MINUS_2 PtpAdvanceConfigurationLogDelayReqInterval = -2
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__MINUS_1 PtpAdvanceConfigurationLogDelayReqInterval = -1
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__0       PtpAdvanceConfigurationLogDelayReqInterval = 0
	PTPADVANCECONFIGURATIONLOGDELAYREQINTERVAL__1       PtpAdvanceConfigurationLogDelayReqInterval = 1
)

// All allowed values of PtpAdvanceConfigurationLogDelayReqInterval enum
var AllowedPtpAdvanceConfigurationLogDelayReqIntervalEnumValues = []PtpAdvanceConfigurationLogDelayReqInterval{
	-5,
	-4,
	-3,
	-2,
	-1,
	0,
	1,
}

func (v *PtpAdvanceConfigurationLogDelayReqInterval) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PtpAdvanceConfigurationLogDelayReqInterval(value)
	for _, existing := range AllowedPtpAdvanceConfigurationLogDelayReqIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PtpAdvanceConfigurationLogDelayReqInterval", value)
}

// NewPtpAdvanceConfigurationLogDelayReqIntervalFromValue returns a pointer to a valid PtpAdvanceConfigurationLogDelayReqInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPtpAdvanceConfigurationLogDelayReqIntervalFromValue(v int32) (*PtpAdvanceConfigurationLogDelayReqInterval, error) {
	ev := PtpAdvanceConfigurationLogDelayReqInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PtpAdvanceConfigurationLogDelayReqInterval: valid values are %v", v, AllowedPtpAdvanceConfigurationLogDelayReqIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PtpAdvanceConfigurationLogDelayReqInterval) IsValid() bool {
	for _, existing := range AllowedPtpAdvanceConfigurationLogDelayReqIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ptpAdvanceConfiguration_logDelayReqInterval value
func (v PtpAdvanceConfigurationLogDelayReqInterval) Ptr() *PtpAdvanceConfigurationLogDelayReqInterval {
	return &v
}

type NullablePtpAdvanceConfigurationLogDelayReqInterval struct {
	value *PtpAdvanceConfigurationLogDelayReqInterval
	isSet bool
}

func (v NullablePtpAdvanceConfigurationLogDelayReqInterval) Get() *PtpAdvanceConfigurationLogDelayReqInterval {
	return v.value
}

func (v *NullablePtpAdvanceConfigurationLogDelayReqInterval) Set(val *PtpAdvanceConfigurationLogDelayReqInterval) {
	v.value = val
	v.isSet = true
}

func (v NullablePtpAdvanceConfigurationLogDelayReqInterval) IsSet() bool {
	return v.isSet
}

func (v *NullablePtpAdvanceConfigurationLogDelayReqInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePtpAdvanceConfigurationLogDelayReqInterval(val *PtpAdvanceConfigurationLogDelayReqInterval) *NullablePtpAdvanceConfigurationLogDelayReqInterval {
	return &NullablePtpAdvanceConfigurationLogDelayReqInterval{value: val, isSet: true}
}

func (v NullablePtpAdvanceConfigurationLogDelayReqInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePtpAdvanceConfigurationLogDelayReqInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
