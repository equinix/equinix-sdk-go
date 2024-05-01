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

// checks if the RoutingProtocol type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocol{}

// RoutingProtocol struct for RoutingProtocol
type RoutingProtocol struct {
	Type                 RoutingProtocolType     `json:"type"`
	Ipv4                 *RoutingProtocolIpBlock `json:"ipv4,omitempty"`
	Ipv6                 *RoutingProtocolIpBlock `json:"ipv6,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocol RoutingProtocol

// NewRoutingProtocol instantiates a new RoutingProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocol(type_ RoutingProtocolType) *RoutingProtocol {
	this := RoutingProtocol{}
	this.Type = type_
	return &this
}

// NewRoutingProtocolWithDefaults instantiates a new RoutingProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolWithDefaults() *RoutingProtocol {
	this := RoutingProtocol{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingProtocol) GetType() RoutingProtocolType {
	if o == nil {
		var ret RoutingProtocolType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocol) GetTypeOk() (*RoutingProtocolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingProtocol) SetType(v RoutingProtocolType) {
	o.Type = v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *RoutingProtocol) GetIpv4() RoutingProtocolIpBlock {
	if o == nil || IsNil(o.Ipv4) {
		var ret RoutingProtocolIpBlock
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocol) GetIpv4Ok() (*RoutingProtocolIpBlock, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *RoutingProtocol) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given RoutingProtocolIpBlock and assigns it to the Ipv4 field.
func (o *RoutingProtocol) SetIpv4(v RoutingProtocolIpBlock) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *RoutingProtocol) GetIpv6() RoutingProtocolIpBlock {
	if o == nil || IsNil(o.Ipv6) {
		var ret RoutingProtocolIpBlock
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocol) GetIpv6Ok() (*RoutingProtocolIpBlock, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *RoutingProtocol) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given RoutingProtocolIpBlock and assigns it to the Ipv6 field.
func (o *RoutingProtocol) SetIpv6(v RoutingProtocolIpBlock) {
	o.Ipv6 = &v
}

func (o RoutingProtocol) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocol) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if !IsNil(o.Ipv6) {
		toSerialize["ipv6"] = o.Ipv6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocol) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRoutingProtocol := _RoutingProtocol{}

	err = json.Unmarshal(data, &varRoutingProtocol)

	if err != nil {
		return err
	}

	*o = RoutingProtocol(varRoutingProtocol)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ipv6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocol struct {
	value *RoutingProtocol
	isSet bool
}

func (v NullableRoutingProtocol) Get() *RoutingProtocol {
	return v.value
}

func (v *NullableRoutingProtocol) Set(val *RoutingProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocol(val *RoutingProtocol) *NullableRoutingProtocol {
	return &NullableRoutingProtocol{value: val, isSet: true}
}

func (v NullableRoutingProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}