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

// checks if the UniqueEntityReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UniqueEntityReference{}

// UniqueEntityReference struct for UniqueEntityReference
type UniqueEntityReference struct {
	Href                 string `json:"href"`
	Uuid                 string `json:"uuid"`
	AdditionalProperties map[string]interface{}
}

type _UniqueEntityReference UniqueEntityReference

// NewUniqueEntityReference instantiates a new UniqueEntityReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUniqueEntityReference(href string, uuid string) *UniqueEntityReference {
	this := UniqueEntityReference{}
	this.Href = href
	this.Uuid = uuid
	return &this
}

// NewUniqueEntityReferenceWithDefaults instantiates a new UniqueEntityReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUniqueEntityReferenceWithDefaults() *UniqueEntityReference {
	this := UniqueEntityReference{}
	return &this
}

// GetHref returns the Href field value
func (o *UniqueEntityReference) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *UniqueEntityReference) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *UniqueEntityReference) SetHref(v string) {
	o.Href = v
}

// GetUuid returns the Uuid field value
func (o *UniqueEntityReference) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *UniqueEntityReference) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *UniqueEntityReference) SetUuid(v string) {
	o.Uuid = v
}

func (o UniqueEntityReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UniqueEntityReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["uuid"] = o.Uuid

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UniqueEntityReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"uuid",
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

	varUniqueEntityReference := _UniqueEntityReference{}

	err = json.Unmarshal(data, &varUniqueEntityReference)

	if err != nil {
		return err
	}

	*o = UniqueEntityReference(varUniqueEntityReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUniqueEntityReference struct {
	value *UniqueEntityReference
	isSet bool
}

func (v NullableUniqueEntityReference) Get() *UniqueEntityReference {
	return v.value
}

func (v *NullableUniqueEntityReference) Set(val *UniqueEntityReference) {
	v.value = val
	v.isSet = true
}

func (v NullableUniqueEntityReference) IsSet() bool {
	return v.isSet
}

func (v *NullableUniqueEntityReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUniqueEntityReference(val *UniqueEntityReference) *NullableUniqueEntityReference {
	return &NullableUniqueEntityReference{value: val, isSet: true}
}

func (v NullableUniqueEntityReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUniqueEntityReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
