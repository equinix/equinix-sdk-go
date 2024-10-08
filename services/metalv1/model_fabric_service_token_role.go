/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// FabricServiceTokenRole Either primary or secondary, depending on which interconnection the service token is associated to.
type FabricServiceTokenRole string

// List of FabricServiceToken_role
const (
	FABRICSERVICETOKENROLE_PRIMARY   FabricServiceTokenRole = "primary"
	FABRICSERVICETOKENROLE_SECONDARY FabricServiceTokenRole = "secondary"
)

// All allowed values of FabricServiceTokenRole enum
var AllowedFabricServiceTokenRoleEnumValues = []FabricServiceTokenRole{
	"primary",
	"secondary",
}

func (v *FabricServiceTokenRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FabricServiceTokenRole(value)
	for _, existing := range AllowedFabricServiceTokenRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FabricServiceTokenRole", value)
}

// NewFabricServiceTokenRoleFromValue returns a pointer to a valid FabricServiceTokenRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFabricServiceTokenRoleFromValue(v string) (*FabricServiceTokenRole, error) {
	ev := FabricServiceTokenRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FabricServiceTokenRole: valid values are %v", v, AllowedFabricServiceTokenRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FabricServiceTokenRole) IsValid() bool {
	for _, existing := range AllowedFabricServiceTokenRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FabricServiceToken_role value
func (v FabricServiceTokenRole) Ptr() *FabricServiceTokenRole {
	return &v
}

type NullableFabricServiceTokenRole struct {
	value *FabricServiceTokenRole
	isSet bool
}

func (v NullableFabricServiceTokenRole) Get() *FabricServiceTokenRole {
	return v.value
}

func (v *NullableFabricServiceTokenRole) Set(val *FabricServiceTokenRole) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricServiceTokenRole) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricServiceTokenRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricServiceTokenRole(val *FabricServiceTokenRole) *NullableFabricServiceTokenRole {
	return &NullableFabricServiceTokenRole{value: val, isSet: true}
}

func (v NullableFabricServiceTokenRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricServiceTokenRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
