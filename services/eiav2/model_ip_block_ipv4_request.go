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

// checks if the IpBlockIpv4Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpBlockIpv4Request{}

// IpBlockIpv4Request struct for IpBlockIpv4Request
type IpBlockIpv4Request struct {
	Uuid *string `json:"uuid,omitempty"`
	// Collection of addressing plans
	AddressingPlans []IpBlockAddressingPlans `json:"addressingPlans,omitempty"`
	// Connection of questions
	Questions []IpBlockQuestions `json:"questions,omitempty"`
	// Length of the IP block, number after the / (slash)
	PrefixLength         int32 `json:"prefixLength"`
	AdditionalProperties map[string]interface{}
}

type _IpBlockIpv4Request IpBlockIpv4Request

// NewIpBlockIpv4Request instantiates a new IpBlockIpv4Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlockIpv4Request(prefixLength int32) *IpBlockIpv4Request {
	this := IpBlockIpv4Request{}
	this.PrefixLength = prefixLength
	return &this
}

// NewIpBlockIpv4RequestWithDefaults instantiates a new IpBlockIpv4Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlockIpv4RequestWithDefaults() *IpBlockIpv4Request {
	this := IpBlockIpv4Request{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *IpBlockIpv4Request) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlockIpv4Request) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *IpBlockIpv4Request) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *IpBlockIpv4Request) SetUuid(v string) {
	o.Uuid = &v
}

// GetAddressingPlans returns the AddressingPlans field value if set, zero value otherwise.
func (o *IpBlockIpv4Request) GetAddressingPlans() []IpBlockAddressingPlans {
	if o == nil || IsNil(o.AddressingPlans) {
		var ret []IpBlockAddressingPlans
		return ret
	}
	return o.AddressingPlans
}

// GetAddressingPlansOk returns a tuple with the AddressingPlans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlockIpv4Request) GetAddressingPlansOk() ([]IpBlockAddressingPlans, bool) {
	if o == nil || IsNil(o.AddressingPlans) {
		return nil, false
	}
	return o.AddressingPlans, true
}

// HasAddressingPlans returns a boolean if a field has been set.
func (o *IpBlockIpv4Request) HasAddressingPlans() bool {
	if o != nil && !IsNil(o.AddressingPlans) {
		return true
	}

	return false
}

// SetAddressingPlans gets a reference to the given []IpBlockAddressingPlans and assigns it to the AddressingPlans field.
func (o *IpBlockIpv4Request) SetAddressingPlans(v []IpBlockAddressingPlans) {
	o.AddressingPlans = v
}

// GetQuestions returns the Questions field value if set, zero value otherwise.
func (o *IpBlockIpv4Request) GetQuestions() []IpBlockQuestions {
	if o == nil || IsNil(o.Questions) {
		var ret []IpBlockQuestions
		return ret
	}
	return o.Questions
}

// GetQuestionsOk returns a tuple with the Questions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlockIpv4Request) GetQuestionsOk() ([]IpBlockQuestions, bool) {
	if o == nil || IsNil(o.Questions) {
		return nil, false
	}
	return o.Questions, true
}

// HasQuestions returns a boolean if a field has been set.
func (o *IpBlockIpv4Request) HasQuestions() bool {
	if o != nil && !IsNil(o.Questions) {
		return true
	}

	return false
}

// SetQuestions gets a reference to the given []IpBlockQuestions and assigns it to the Questions field.
func (o *IpBlockIpv4Request) SetQuestions(v []IpBlockQuestions) {
	o.Questions = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *IpBlockIpv4Request) GetPrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *IpBlockIpv4Request) GetPrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *IpBlockIpv4Request) SetPrefixLength(v int32) {
	o.PrefixLength = v
}

func (o IpBlockIpv4Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpBlockIpv4Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.AddressingPlans) {
		toSerialize["addressingPlans"] = o.AddressingPlans
	}
	if !IsNil(o.Questions) {
		toSerialize["questions"] = o.Questions
	}
	toSerialize["prefixLength"] = o.PrefixLength

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IpBlockIpv4Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefixLength",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varIpBlockIpv4Request := _IpBlockIpv4Request{}

	err = json.Unmarshal(data, &varIpBlockIpv4Request)

	if err != nil {
		return err
	}

	*o = IpBlockIpv4Request(varIpBlockIpv4Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "addressingPlans")
		delete(additionalProperties, "questions")
		delete(additionalProperties, "prefixLength")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpBlockIpv4Request struct {
	value *IpBlockIpv4Request
	isSet bool
}

func (v NullableIpBlockIpv4Request) Get() *IpBlockIpv4Request {
	return v.value
}

func (v *NullableIpBlockIpv4Request) Set(val *IpBlockIpv4Request) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlockIpv4Request) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlockIpv4Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlockIpv4Request(val *IpBlockIpv4Request) *NullableIpBlockIpv4Request {
	return &NullableIpBlockIpv4Request{value: val, isSet: true}
}

func (v NullableIpBlockIpv4Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlockIpv4Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
