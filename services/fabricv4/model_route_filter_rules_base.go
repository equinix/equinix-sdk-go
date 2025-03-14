/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RouteFilterRulesBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteFilterRulesBase{}

// RouteFilterRulesBase struct for RouteFilterRulesBase
type RouteFilterRulesBase struct {
	Name *string `json:"name,omitempty"`
	// Customer-provided Route Filter Rule description
	Description          *string `json:"description,omitempty"`
	Prefix               string  `json:"prefix"`
	PrefixMatch          *string `json:"prefixMatch,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteFilterRulesBase RouteFilterRulesBase

// NewRouteFilterRulesBase instantiates a new RouteFilterRulesBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteFilterRulesBase(prefix string) *RouteFilterRulesBase {
	this := RouteFilterRulesBase{}
	this.Prefix = prefix
	var prefixMatch string = "orlonger"
	this.PrefixMatch = &prefixMatch
	return &this
}

// NewRouteFilterRulesBaseWithDefaults instantiates a new RouteFilterRulesBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteFilterRulesBaseWithDefaults() *RouteFilterRulesBase {
	this := RouteFilterRulesBase{}
	var prefixMatch string = "orlonger"
	this.PrefixMatch = &prefixMatch
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RouteFilterRulesBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RouteFilterRulesBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RouteFilterRulesBase) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RouteFilterRulesBase) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RouteFilterRulesBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RouteFilterRulesBase) SetDescription(v string) {
	o.Description = &v
}

// GetPrefix returns the Prefix field value
func (o *RouteFilterRulesBase) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesBase) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *RouteFilterRulesBase) SetPrefix(v string) {
	o.Prefix = v
}

// GetPrefixMatch returns the PrefixMatch field value if set, zero value otherwise.
func (o *RouteFilterRulesBase) GetPrefixMatch() string {
	if o == nil || IsNil(o.PrefixMatch) {
		var ret string
		return ret
	}
	return *o.PrefixMatch
}

// GetPrefixMatchOk returns a tuple with the PrefixMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteFilterRulesBase) GetPrefixMatchOk() (*string, bool) {
	if o == nil || IsNil(o.PrefixMatch) {
		return nil, false
	}
	return o.PrefixMatch, true
}

// HasPrefixMatch returns a boolean if a field has been set.
func (o *RouteFilterRulesBase) HasPrefixMatch() bool {
	if o != nil && !IsNil(o.PrefixMatch) {
		return true
	}

	return false
}

// SetPrefixMatch gets a reference to the given string and assigns it to the PrefixMatch field.
func (o *RouteFilterRulesBase) SetPrefixMatch(v string) {
	o.PrefixMatch = &v
}

func (o RouteFilterRulesBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteFilterRulesBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["prefix"] = o.Prefix
	if !IsNil(o.PrefixMatch) {
		toSerialize["prefixMatch"] = o.PrefixMatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteFilterRulesBase) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefix",
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

	varRouteFilterRulesBase := _RouteFilterRulesBase{}

	err = json.Unmarshal(data, &varRouteFilterRulesBase)

	if err != nil {
		return err
	}

	*o = RouteFilterRulesBase(varRouteFilterRulesBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "prefix")
		delete(additionalProperties, "prefixMatch")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteFilterRulesBase struct {
	value *RouteFilterRulesBase
	isSet bool
}

func (v NullableRouteFilterRulesBase) Get() *RouteFilterRulesBase {
	return v.value
}

func (v *NullableRouteFilterRulesBase) Set(val *RouteFilterRulesBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteFilterRulesBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteFilterRulesBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteFilterRulesBase(val *RouteFilterRulesBase) *NullableRouteFilterRulesBase {
	return &NullableRouteFilterRulesBase{value: val, isSet: true}
}

func (v NullableRouteFilterRulesBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteFilterRulesBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
