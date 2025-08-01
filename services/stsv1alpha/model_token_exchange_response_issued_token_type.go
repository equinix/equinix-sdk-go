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

// TokenExchangeResponseIssuedTokenType the model 'TokenExchangeResponseIssuedTokenType'
type TokenExchangeResponseIssuedTokenType string

// List of TokenExchangeResponse_issued_token_type
const (
	TOKENEXCHANGERESPONSEISSUEDTOKENTYPE_URN_IETF_PARAMS_OAUTH_TOKEN_TYPE_ACCESS_TOKEN TokenExchangeResponseIssuedTokenType = "urn:ietf:params:oauth:token-type:access_token"
)

// All allowed values of TokenExchangeResponseIssuedTokenType enum
var AllowedTokenExchangeResponseIssuedTokenTypeEnumValues = []TokenExchangeResponseIssuedTokenType{
	"urn:ietf:params:oauth:token-type:access_token",
}

func (v *TokenExchangeResponseIssuedTokenType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TokenExchangeResponseIssuedTokenType(value)
	for _, existing := range AllowedTokenExchangeResponseIssuedTokenTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TokenExchangeResponseIssuedTokenType", value)
}

// NewTokenExchangeResponseIssuedTokenTypeFromValue returns a pointer to a valid TokenExchangeResponseIssuedTokenType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTokenExchangeResponseIssuedTokenTypeFromValue(v string) (*TokenExchangeResponseIssuedTokenType, error) {
	ev := TokenExchangeResponseIssuedTokenType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TokenExchangeResponseIssuedTokenType: valid values are %v", v, AllowedTokenExchangeResponseIssuedTokenTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TokenExchangeResponseIssuedTokenType) IsValid() bool {
	for _, existing := range AllowedTokenExchangeResponseIssuedTokenTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TokenExchangeResponse_issued_token_type value
func (v TokenExchangeResponseIssuedTokenType) Ptr() *TokenExchangeResponseIssuedTokenType {
	return &v
}

type NullableTokenExchangeResponseIssuedTokenType struct {
	value *TokenExchangeResponseIssuedTokenType
	isSet bool
}

func (v NullableTokenExchangeResponseIssuedTokenType) Get() *TokenExchangeResponseIssuedTokenType {
	return v.value
}

func (v *NullableTokenExchangeResponseIssuedTokenType) Set(val *TokenExchangeResponseIssuedTokenType) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenExchangeResponseIssuedTokenType) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenExchangeResponseIssuedTokenType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenExchangeResponseIssuedTokenType(val *TokenExchangeResponseIssuedTokenType) *NullableTokenExchangeResponseIssuedTokenType {
	return &NullableTokenExchangeResponseIssuedTokenType{value: val, isSet: true}
}

func (v NullableTokenExchangeResponseIssuedTokenType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenExchangeResponseIssuedTokenType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
