/*
Equinix Security Token Service

Exchange ID tokens for Equinix access tokens according to managed trust relationships. STS is an alpha service and is not generally available to customers.

API version: 1.0.0-alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stsv1alpha

import (
	"encoding/json"
	"fmt"
)

// UseTokenPostRequestGrantType the model 'UseTokenPostRequestGrantType'
type UseTokenPostRequestGrantType string

// List of _use_token_post_request_grant_type
const (
	USETOKENPOSTREQUESTGRANTTYPE_CLIENT_CREDENTIALS                              UseTokenPostRequestGrantType = "client_credentials"
	USETOKENPOSTREQUESTGRANTTYPE_URN_IETF_PARAMS_OAUTH_GRANT_TYPE_TOKEN_EXCHANGE UseTokenPostRequestGrantType = "urn:ietf:params:oauth:grant-type:token-exchange"
)

// All allowed values of UseTokenPostRequestGrantType enum
var AllowedUseTokenPostRequestGrantTypeEnumValues = []UseTokenPostRequestGrantType{
	"client_credentials",
	"urn:ietf:params:oauth:grant-type:token-exchange",
}

func (v *UseTokenPostRequestGrantType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UseTokenPostRequestGrantType(value)
	for _, existing := range AllowedUseTokenPostRequestGrantTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UseTokenPostRequestGrantType", value)
}

// NewUseTokenPostRequestGrantTypeFromValue returns a pointer to a valid UseTokenPostRequestGrantType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUseTokenPostRequestGrantTypeFromValue(v string) (*UseTokenPostRequestGrantType, error) {
	ev := UseTokenPostRequestGrantType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UseTokenPostRequestGrantType: valid values are %v", v, AllowedUseTokenPostRequestGrantTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UseTokenPostRequestGrantType) IsValid() bool {
	for _, existing := range AllowedUseTokenPostRequestGrantTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to _use_token_post_request_grant_type value
func (v UseTokenPostRequestGrantType) Ptr() *UseTokenPostRequestGrantType {
	return &v
}

type NullableUseTokenPostRequestGrantType struct {
	value *UseTokenPostRequestGrantType
	isSet bool
}

func (v NullableUseTokenPostRequestGrantType) Get() *UseTokenPostRequestGrantType {
	return v.value
}

func (v *NullableUseTokenPostRequestGrantType) Set(val *UseTokenPostRequestGrantType) {
	v.value = val
	v.isSet = true
}

func (v NullableUseTokenPostRequestGrantType) IsSet() bool {
	return v.isSet
}

func (v *NullableUseTokenPostRequestGrantType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUseTokenPostRequestGrantType(val *UseTokenPostRequestGrantType) *NullableUseTokenPostRequestGrantType {
	return &NullableUseTokenPostRequestGrantType{value: val, isSet: true}
}

func (v NullableUseTokenPostRequestGrantType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUseTokenPostRequestGrantType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
