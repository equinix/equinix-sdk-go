/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RouteAggregationRulesBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteAggregationRulesBase{}

// RouteAggregationRulesBase struct for RouteAggregationRulesBase
type RouteAggregationRulesBase struct {
	Name *string `json:"name,omitempty"`
	// Customer-provided Route Aggregation Rule description
	Description          *string `json:"description,omitempty"`
	Prefix               string  `json:"prefix"`
	AdditionalProperties map[string]interface{}
}

type _RouteAggregationRulesBase RouteAggregationRulesBase

// NewRouteAggregationRulesBase instantiates a new RouteAggregationRulesBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteAggregationRulesBase(prefix string) *RouteAggregationRulesBase {
	this := RouteAggregationRulesBase{}
	this.Prefix = prefix
	return &this
}

// NewRouteAggregationRulesBaseWithDefaults instantiates a new RouteAggregationRulesBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteAggregationRulesBaseWithDefaults() *RouteAggregationRulesBase {
	this := RouteAggregationRulesBase{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RouteAggregationRulesBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RouteAggregationRulesBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RouteAggregationRulesBase) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RouteAggregationRulesBase) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesBase) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RouteAggregationRulesBase) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RouteAggregationRulesBase) SetDescription(v string) {
	o.Description = &v
}

// GetPrefix returns the Prefix field value
func (o *RouteAggregationRulesBase) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesBase) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *RouteAggregationRulesBase) SetPrefix(v string) {
	o.Prefix = v
}

func (o RouteAggregationRulesBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteAggregationRulesBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["prefix"] = o.Prefix

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteAggregationRulesBase) UnmarshalJSON(data []byte) (err error) {
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

	varRouteAggregationRulesBase := _RouteAggregationRulesBase{}

	err = json.Unmarshal(data, &varRouteAggregationRulesBase)

	if err != nil {
		return err
	}

	*o = RouteAggregationRulesBase(varRouteAggregationRulesBase)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteAggregationRulesBase struct {
	value *RouteAggregationRulesBase
	isSet bool
}

func (v NullableRouteAggregationRulesBase) Get() *RouteAggregationRulesBase {
	return v.value
}

func (v *NullableRouteAggregationRulesBase) Set(val *RouteAggregationRulesBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteAggregationRulesBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteAggregationRulesBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteAggregationRulesBase(val *RouteAggregationRulesBase) *NullableRouteAggregationRulesBase {
	return &NullableRouteAggregationRulesBase{value: val, isSet: true}
}

func (v NullableRouteAggregationRulesBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteAggregationRulesBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
