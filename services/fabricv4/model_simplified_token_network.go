/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the SimplifiedTokenNetwork type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplifiedTokenNetwork{}

// SimplifiedTokenNetwork struct for SimplifiedTokenNetwork
type SimplifiedTokenNetwork struct {
	// url to entity
	Href *string `json:"href,omitempty"`
	// Network Identifier
	Uuid *string                     `json:"uuid,omitempty"`
	Type *SimplifiedTokenNetworkType `json:"type,omitempty"`
	// Network Name
	Name                 *string                      `json:"name,omitempty"`
	Scope                *SimplifiedTokenNetworkScope `json:"scope,omitempty"`
	Location             *SimplifiedLocation          `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SimplifiedTokenNetwork SimplifiedTokenNetwork

// NewSimplifiedTokenNetwork instantiates a new SimplifiedTokenNetwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplifiedTokenNetwork() *SimplifiedTokenNetwork {
	this := SimplifiedTokenNetwork{}
	return &this
}

// NewSimplifiedTokenNetworkWithDefaults instantiates a new SimplifiedTokenNetwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplifiedTokenNetworkWithDefaults() *SimplifiedTokenNetwork {
	this := SimplifiedTokenNetwork{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SimplifiedTokenNetwork) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SimplifiedTokenNetwork) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetType() SimplifiedTokenNetworkType {
	if o == nil || IsNil(o.Type) {
		var ret SimplifiedTokenNetworkType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetTypeOk() (*SimplifiedTokenNetworkType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given SimplifiedTokenNetworkType and assigns it to the Type field.
func (o *SimplifiedTokenNetwork) SetType(v SimplifiedTokenNetworkType) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SimplifiedTokenNetwork) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetScope() SimplifiedTokenNetworkScope {
	if o == nil || IsNil(o.Scope) {
		var ret SimplifiedTokenNetworkScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetScopeOk() (*SimplifiedTokenNetworkScope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given SimplifiedTokenNetworkScope and assigns it to the Scope field.
func (o *SimplifiedTokenNetwork) SetScope(v SimplifiedTokenNetworkScope) {
	o.Scope = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *SimplifiedTokenNetwork) GetLocation() SimplifiedLocation {
	if o == nil || IsNil(o.Location) {
		var ret SimplifiedLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplifiedTokenNetwork) GetLocationOk() (*SimplifiedLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *SimplifiedTokenNetwork) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given SimplifiedLocation and assigns it to the Location field.
func (o *SimplifiedTokenNetwork) SetLocation(v SimplifiedLocation) {
	o.Location = &v
}

func (o SimplifiedTokenNetwork) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplifiedTokenNetwork) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SimplifiedTokenNetwork) UnmarshalJSON(data []byte) (err error) {
	varSimplifiedTokenNetwork := _SimplifiedTokenNetwork{}

	err = json.Unmarshal(data, &varSimplifiedTokenNetwork)

	if err != nil {
		return err
	}

	*o = SimplifiedTokenNetwork(varSimplifiedTokenNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSimplifiedTokenNetwork struct {
	value *SimplifiedTokenNetwork
	isSet bool
}

func (v NullableSimplifiedTokenNetwork) Get() *SimplifiedTokenNetwork {
	return v.value
}

func (v *NullableSimplifiedTokenNetwork) Set(val *SimplifiedTokenNetwork) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplifiedTokenNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplifiedTokenNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplifiedTokenNetwork(val *SimplifiedTokenNetwork) *NullableSimplifiedTokenNetwork {
	return &NullableSimplifiedTokenNetwork{value: val, isSet: true}
}

func (v NullableSimplifiedTokenNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplifiedTokenNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}