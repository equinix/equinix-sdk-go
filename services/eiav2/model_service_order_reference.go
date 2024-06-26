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

// checks if the ServiceOrderReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOrderReference{}

// ServiceOrderReference struct for ServiceOrderReference
type ServiceOrderReference struct {
	Href                 string           `json:"href"`
	Uuid                 string           `json:"uuid"`
	Type                 ServiceOrderType `json:"type"`
	AdditionalProperties map[string]interface{}
}

type _ServiceOrderReference ServiceOrderReference

// NewServiceOrderReference instantiates a new ServiceOrderReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOrderReference(href string, uuid string, type_ ServiceOrderType) *ServiceOrderReference {
	this := ServiceOrderReference{}
	this.Href = href
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewServiceOrderReferenceWithDefaults instantiates a new ServiceOrderReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOrderReferenceWithDefaults() *ServiceOrderReference {
	this := ServiceOrderReference{}
	return &this
}

// GetHref returns the Href field value
func (o *ServiceOrderReference) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ServiceOrderReference) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ServiceOrderReference) SetHref(v string) {
	o.Href = v
}

// GetUuid returns the Uuid field value
func (o *ServiceOrderReference) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *ServiceOrderReference) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *ServiceOrderReference) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *ServiceOrderReference) GetType() ServiceOrderType {
	if o == nil {
		var ret ServiceOrderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceOrderReference) GetTypeOk() (*ServiceOrderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceOrderReference) SetType(v ServiceOrderType) {
	o.Type = v
}

func (o ServiceOrderReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOrderReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceOrderReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"uuid",
		"type",
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

	varServiceOrderReference := _ServiceOrderReference{}

	err = json.Unmarshal(data, &varServiceOrderReference)

	if err != nil {
		return err
	}

	*o = ServiceOrderReference(varServiceOrderReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceOrderReference struct {
	value *ServiceOrderReference
	isSet bool
}

func (v NullableServiceOrderReference) Get() *ServiceOrderReference {
	return v.value
}

func (v *NullableServiceOrderReference) Set(val *ServiceOrderReference) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOrderReference) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOrderReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOrderReference(val *ServiceOrderReference) *NullableServiceOrderReference {
	return &NullableServiceOrderReference{value: val, isSet: true}
}

func (v NullableServiceOrderReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOrderReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
