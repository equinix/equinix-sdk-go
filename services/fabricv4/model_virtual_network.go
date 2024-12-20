/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the VirtualNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VirtualNetwork{}

// VirtualNetwork Virtual Network Information
type VirtualNetwork struct {
	// The Canonical URL at which the resource resides.
	Href *string `json:"href,omitempty"`
	// Equinix-assigned Virtual Network identifier
	Uuid                 *string `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VirtualNetwork VirtualNetwork

// NewVirtualNetwork instantiates a new VirtualNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualNetwork() *VirtualNetwork {
	this := VirtualNetwork{}
	return &this
}

// NewVirtualNetworkWithDefaults instantiates a new VirtualNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualNetworkWithDefaults() *VirtualNetwork {
	this := VirtualNetwork{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *VirtualNetwork) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *VirtualNetwork) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *VirtualNetwork) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *VirtualNetwork) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualNetwork) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *VirtualNetwork) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *VirtualNetwork) SetUuid(v string) {
	o.Uuid = &v
}

func (o VirtualNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VirtualNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VirtualNetwork) UnmarshalJSON(data []byte) (err error) {
	varVirtualNetwork := _VirtualNetwork{}

	err = json.Unmarshal(data, &varVirtualNetwork)

	if err != nil {
		return err
	}

	*o = VirtualNetwork(varVirtualNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVirtualNetwork struct {
	value *VirtualNetwork
	isSet bool
}

func (v NullableVirtualNetwork) Get() *VirtualNetwork {
	return v.value
}

func (v *NullableVirtualNetwork) Set(val *VirtualNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualNetwork(val *VirtualNetwork) *NullableVirtualNetwork {
	return &NullableVirtualNetwork{value: val, isSet: true}
}

func (v NullableVirtualNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
