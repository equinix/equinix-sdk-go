/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
)

// checks if the Taggable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Taggable{}

// Taggable struct for Taggable
type Taggable struct {
	Tags                 []string `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Taggable Taggable

// NewTaggable instantiates a new Taggable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaggable() *Taggable {
	this := Taggable{}
	return &this
}

// NewTaggableWithDefaults instantiates a new Taggable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaggableWithDefaults() *Taggable {
	this := Taggable{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Taggable) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Taggable) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Taggable) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Taggable) SetTags(v []string) {
	o.Tags = v
}

func (o Taggable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Taggable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Taggable) UnmarshalJSON(data []byte) (err error) {
	varTaggable := _Taggable{}

	err = json.Unmarshal(data, &varTaggable)

	if err != nil {
		return err
	}

	*o = Taggable(varTaggable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTaggable struct {
	value *Taggable
	isSet bool
}

func (v NullableTaggable) Get() *Taggable {
	return v.value
}

func (v *NullableTaggable) Set(val *Taggable) {
	v.value = val
	v.isSet = true
}

func (v NullableTaggable) IsSet() bool {
	return v.isSet
}

func (v *NullableTaggable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaggable(val *Taggable) *NullableTaggable {
	return &NullableTaggable{value: val, isSet: true}
}

func (v NullableTaggable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaggable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
