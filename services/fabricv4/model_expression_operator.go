/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// ExpressionOperator the model 'ExpressionOperator'
type ExpressionOperator string

// List of Expression_operator
const (
	EXPRESSIONOPERATOR_EQUAL                    ExpressionOperator = "="
	EXPRESSIONOPERATOR_NOT_EQUAL                ExpressionOperator = "!="
	EXPRESSIONOPERATOR_GREATER_THAN             ExpressionOperator = ">"
	EXPRESSIONOPERATOR_GREATER_THAN_OR_EQUAL_TO ExpressionOperator = ">="
	EXPRESSIONOPERATOR_LESS_THAN                ExpressionOperator = "<"
	EXPRESSIONOPERATOR_LESS_THAN_OR_EQUAL_TO    ExpressionOperator = "<="
	EXPRESSIONOPERATOR_BETWEEN                  ExpressionOperator = "BETWEEN"
	EXPRESSIONOPERATOR_NOT_BETWEEN              ExpressionOperator = "NOT BETWEEN"
	EXPRESSIONOPERATOR_LIKE                     ExpressionOperator = "LIKE"
	EXPRESSIONOPERATOR_NOT_LIKE                 ExpressionOperator = "NOT LIKE"
	EXPRESSIONOPERATOR_IN                       ExpressionOperator = "IN"
	EXPRESSIONOPERATOR_NOT_IN                   ExpressionOperator = "NOT IN"
	EXPRESSIONOPERATOR_IS_NOT_NULL              ExpressionOperator = "IS NOT NULL"
	EXPRESSIONOPERATOR_IS_NULL                  ExpressionOperator = "IS NULL"
)

// All allowed values of ExpressionOperator enum
var AllowedExpressionOperatorEnumValues = []ExpressionOperator{
	"=",
	"!=",
	">",
	">=",
	"<",
	"<=",
	"BETWEEN",
	"NOT BETWEEN",
	"LIKE",
	"NOT LIKE",
	"IN",
	"NOT IN",
	"IS NOT NULL",
	"IS NULL",
}

func (v *ExpressionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpressionOperator(value)
	for _, existing := range AllowedExpressionOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpressionOperator", value)
}

// NewExpressionOperatorFromValue returns a pointer to a valid ExpressionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpressionOperatorFromValue(v string) (*ExpressionOperator, error) {
	ev := ExpressionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpressionOperator: valid values are %v", v, AllowedExpressionOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpressionOperator) IsValid() bool {
	for _, existing := range AllowedExpressionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Expression_operator value
func (v ExpressionOperator) Ptr() *ExpressionOperator {
	return &v
}

type NullableExpressionOperator struct {
	value *ExpressionOperator
	isSet bool
}

func (v NullableExpressionOperator) Get() *ExpressionOperator {
	return v.value
}

func (v *NullableExpressionOperator) Set(val *ExpressionOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableExpressionOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableExpressionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpressionOperator(val *ExpressionOperator) *NullableExpressionOperator {
	return &NullableExpressionOperator{value: val, isSet: true}
}

func (v NullableExpressionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpressionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}