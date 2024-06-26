/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// IpBlockQuestionsAnswer the model 'IpBlockQuestionsAnswer'
type IpBlockQuestionsAnswer string

// List of IpBlockQuestions_answer
const (
	IPBLOCKQUESTIONSANSWER_TRUE  IpBlockQuestionsAnswer = "true"
	IPBLOCKQUESTIONSANSWER_FALSE IpBlockQuestionsAnswer = "false"
)

// All allowed values of IpBlockQuestionsAnswer enum
var AllowedIpBlockQuestionsAnswerEnumValues = []IpBlockQuestionsAnswer{
	"true",
	"false",
}

func (v *IpBlockQuestionsAnswer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IpBlockQuestionsAnswer(value)
	for _, existing := range AllowedIpBlockQuestionsAnswerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IpBlockQuestionsAnswer", value)
}

// NewIpBlockQuestionsAnswerFromValue returns a pointer to a valid IpBlockQuestionsAnswer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIpBlockQuestionsAnswerFromValue(v string) (*IpBlockQuestionsAnswer, error) {
	ev := IpBlockQuestionsAnswer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IpBlockQuestionsAnswer: valid values are %v", v, AllowedIpBlockQuestionsAnswerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IpBlockQuestionsAnswer) IsValid() bool {
	for _, existing := range AllowedIpBlockQuestionsAnswerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IpBlockQuestions_answer value
func (v IpBlockQuestionsAnswer) Ptr() *IpBlockQuestionsAnswer {
	return &v
}

type NullableIpBlockQuestionsAnswer struct {
	value *IpBlockQuestionsAnswer
	isSet bool
}

func (v NullableIpBlockQuestionsAnswer) Get() *IpBlockQuestionsAnswer {
	return v.value
}

func (v *NullableIpBlockQuestionsAnswer) Set(val *IpBlockQuestionsAnswer) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockQuestionsAnswer) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockQuestionsAnswer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockQuestionsAnswer(val *IpBlockQuestionsAnswer) *NullableIpBlockQuestionsAnswer {
	return &NullableIpBlockQuestionsAnswer{value: val, isSet: true}
}

func (v NullableIpBlockQuestionsAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockQuestionsAnswer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
