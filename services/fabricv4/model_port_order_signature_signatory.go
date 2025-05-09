/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// PortOrderSignatureSignatory Port signature Type
type PortOrderSignatureSignatory string

// List of PortOrder_signature_signatory
const (
	PORTORDERSIGNATURESIGNATORY_DELEGATE        PortOrderSignatureSignatory = "DELEGATE"
	PORTORDERSIGNATURESIGNATORY_SELF            PortOrderSignatureSignatory = "SELF"
	PORTORDERSIGNATURESIGNATORY_ACCOUNT_SUPPORT PortOrderSignatureSignatory = "ACCOUNT_SUPPORT"
)

// All allowed values of PortOrderSignatureSignatory enum
var AllowedPortOrderSignatureSignatoryEnumValues = []PortOrderSignatureSignatory{
	"DELEGATE",
	"SELF",
	"ACCOUNT_SUPPORT",
}

func (v *PortOrderSignatureSignatory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortOrderSignatureSignatory(value)
	for _, existing := range AllowedPortOrderSignatureSignatoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortOrderSignatureSignatory", value)
}

// NewPortOrderSignatureSignatoryFromValue returns a pointer to a valid PortOrderSignatureSignatory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortOrderSignatureSignatoryFromValue(v string) (*PortOrderSignatureSignatory, error) {
	ev := PortOrderSignatureSignatory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortOrderSignatureSignatory: valid values are %v", v, AllowedPortOrderSignatureSignatoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortOrderSignatureSignatory) IsValid() bool {
	for _, existing := range AllowedPortOrderSignatureSignatoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortOrder_signature_signatory value
func (v PortOrderSignatureSignatory) Ptr() *PortOrderSignatureSignatory {
	return &v
}

type NullablePortOrderSignatureSignatory struct {
	value *PortOrderSignatureSignatory
	isSet bool
}

func (v NullablePortOrderSignatureSignatory) Get() *PortOrderSignatureSignatory {
	return v.value
}

func (v *NullablePortOrderSignatureSignatory) Set(val *PortOrderSignatureSignatory) {
	v.value = val
	v.isSet = true
}

func (v NullablePortOrderSignatureSignatory) IsSet() bool {
	return v.isSet
}

func (v *NullablePortOrderSignatureSignatory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortOrderSignatureSignatory(val *PortOrderSignatureSignatory) *NullablePortOrderSignatureSignatory {
	return &NullablePortOrderSignatureSignatory{value: val, isSet: true}
}

func (v NullablePortOrderSignatureSignatory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortOrderSignatureSignatory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
