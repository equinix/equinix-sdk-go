/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PriceTermLength In months. No value or value of 1 means on-demand
type PriceTermLength int32

// List of Price_termLength
const (
	PRICETERMLENGTH__1  PriceTermLength = 1
	PRICETERMLENGTH__12 PriceTermLength = 12
	PRICETERMLENGTH__24 PriceTermLength = 24
	PRICETERMLENGTH__36 PriceTermLength = 36
)

// All allowed values of PriceTermLength enum
var AllowedPriceTermLengthEnumValues = []PriceTermLength{
	1,
	12,
	24,
	36,
}

func (v *PriceTermLength) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriceTermLength(value)
	for _, existing := range AllowedPriceTermLengthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriceTermLength", value)
}

// NewPriceTermLengthFromValue returns a pointer to a valid PriceTermLength
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriceTermLengthFromValue(v int32) (*PriceTermLength, error) {
	ev := PriceTermLength(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriceTermLength: valid values are %v", v, AllowedPriceTermLengthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriceTermLength) IsValid() bool {
	for _, existing := range AllowedPriceTermLengthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Price_termLength value
func (v PriceTermLength) Ptr() *PriceTermLength {
	return &v
}

type NullablePriceTermLength struct {
	value *PriceTermLength
	isSet bool
}

func (v NullablePriceTermLength) Get() *PriceTermLength {
	return v.value
}

func (v *NullablePriceTermLength) Set(val *PriceTermLength) {
	v.value = val
	v.isSet = true
}

func (v NullablePriceTermLength) IsSet() bool {
	return v.isSet
}

func (v *NullablePriceTermLength) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriceTermLength(val *PriceTermLength) *NullablePriceTermLength {
	return &NullablePriceTermLength{value: val, isSet: true}
}

func (v NullablePriceTermLength) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriceTermLength) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
