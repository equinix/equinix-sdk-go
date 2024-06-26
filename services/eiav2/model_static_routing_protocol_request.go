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

// checks if the StaticRoutingProtocolRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StaticRoutingProtocolRequest{}

// StaticRoutingProtocolRequest struct for StaticRoutingProtocolRequest
type StaticRoutingProtocolRequest struct {
	Tags []string            `json:"tags,omitempty"`
	Type RoutingProtocolType `json:"type"`
	// Name of the routing protocol instance.
	Name *string `json:"name,omitempty"`
	// Description of the routing protocol instance
	Description          *string                     `json:"description,omitempty"`
	Ipv4                 *RoutingProtocolIpv4Request `json:"ipv4,omitempty"`
	Ipv6                 *RoutingProtocolIpv6Request `json:"ipv6,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StaticRoutingProtocolRequest StaticRoutingProtocolRequest

// NewStaticRoutingProtocolRequest instantiates a new StaticRoutingProtocolRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStaticRoutingProtocolRequest(type_ RoutingProtocolType) *StaticRoutingProtocolRequest {
	this := StaticRoutingProtocolRequest{}
	this.Type = type_
	return &this
}

// NewStaticRoutingProtocolRequestWithDefaults instantiates a new StaticRoutingProtocolRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStaticRoutingProtocolRequestWithDefaults() *StaticRoutingProtocolRequest {
	this := StaticRoutingProtocolRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *StaticRoutingProtocolRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *StaticRoutingProtocolRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *StaticRoutingProtocolRequest) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *StaticRoutingProtocolRequest) GetType() RoutingProtocolType {
	if o == nil {
		var ret RoutingProtocolType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetTypeOk() (*RoutingProtocolType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *StaticRoutingProtocolRequest) SetType(v RoutingProtocolType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StaticRoutingProtocolRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StaticRoutingProtocolRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StaticRoutingProtocolRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StaticRoutingProtocolRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StaticRoutingProtocolRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StaticRoutingProtocolRequest) SetDescription(v string) {
	o.Description = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *StaticRoutingProtocolRequest) GetIpv4() RoutingProtocolIpv4Request {
	if o == nil || IsNil(o.Ipv4) {
		var ret RoutingProtocolIpv4Request
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *StaticRoutingProtocolRequest) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given RoutingProtocolIpv4Request and assigns it to the Ipv4 field.
func (o *StaticRoutingProtocolRequest) SetIpv4(v RoutingProtocolIpv4Request) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise.
func (o *StaticRoutingProtocolRequest) GetIpv6() RoutingProtocolIpv6Request {
	if o == nil || IsNil(o.Ipv6) {
		var ret RoutingProtocolIpv6Request
		return ret
	}
	return *o.Ipv6
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StaticRoutingProtocolRequest) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool) {
	if o == nil || IsNil(o.Ipv6) {
		return nil, false
	}
	return o.Ipv6, true
}

// HasIpv6 returns a boolean if a field has been set.
func (o *StaticRoutingProtocolRequest) HasIpv6() bool {
	if o != nil && !IsNil(o.Ipv6) {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given RoutingProtocolIpv6Request and assigns it to the Ipv6 field.
func (o *StaticRoutingProtocolRequest) SetIpv6(v RoutingProtocolIpv6Request) {
	o.Ipv6 = &v
}

func (o StaticRoutingProtocolRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StaticRoutingProtocolRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
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

func (o *StaticRoutingProtocolRequest) UnmarshalJSON(data []byte) (err error) {
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

	varStaticRoutingProtocolRequest := _StaticRoutingProtocolRequest{}

	err = json.Unmarshal(data, &varStaticRoutingProtocolRequest)

	if err != nil {
		return err
	}

	*o = StaticRoutingProtocolRequest(varStaticRoutingProtocolRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tags")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "ipv4")
		delete(additionalProperties, "ipv6")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableStaticRoutingProtocolRequest struct {
	value *StaticRoutingProtocolRequest
	isSet bool
}

func (v NullableStaticRoutingProtocolRequest) Get() *StaticRoutingProtocolRequest {
	return v.value
}

func (v *NullableStaticRoutingProtocolRequest) Set(val *StaticRoutingProtocolRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStaticRoutingProtocolRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStaticRoutingProtocolRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStaticRoutingProtocolRequest(val *StaticRoutingProtocolRequest) *NullableStaticRoutingProtocolRequest {
	return &NullableStaticRoutingProtocolRequest{value: val, isSet: true}
}

func (v NullableStaticRoutingProtocolRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStaticRoutingProtocolRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
