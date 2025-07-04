/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the TopologyProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopologyProperties{}

// TopologyProperties TopologyProperties is a schema that defines the properties of a topology in the orchestrator. It includes the element ID and any dependencies that the topology may have.
type TopologyProperties struct {
	ElementId            string   `json:"elementId"`
	DependsOn            []string `json:"dependsOn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopologyProperties TopologyProperties

// NewTopologyProperties instantiates a new TopologyProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopologyProperties(elementId string) *TopologyProperties {
	this := TopologyProperties{}
	this.ElementId = elementId
	return &this
}

// NewTopologyPropertiesWithDefaults instantiates a new TopologyProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyPropertiesWithDefaults() *TopologyProperties {
	this := TopologyProperties{}
	return &this
}

// GetElementId returns the ElementId field value
func (o *TopologyProperties) GetElementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value
// and a boolean to check if the value has been set.
func (o *TopologyProperties) GetElementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElementId, true
}

// SetElementId sets field value
func (o *TopologyProperties) SetElementId(v string) {
	o.ElementId = v
}

// GetDependsOn returns the DependsOn field value if set, zero value otherwise.
func (o *TopologyProperties) GetDependsOn() []string {
	if o == nil || IsNil(o.DependsOn) {
		var ret []string
		return ret
	}
	return o.DependsOn
}

// GetDependsOnOk returns a tuple with the DependsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopologyProperties) GetDependsOnOk() ([]string, bool) {
	if o == nil || IsNil(o.DependsOn) {
		return nil, false
	}
	return o.DependsOn, true
}

// HasDependsOn returns a boolean if a field has been set.
func (o *TopologyProperties) HasDependsOn() bool {
	if o != nil && !IsNil(o.DependsOn) {
		return true
	}

	return false
}

// SetDependsOn gets a reference to the given []string and assigns it to the DependsOn field.
func (o *TopologyProperties) SetDependsOn(v []string) {
	o.DependsOn = v
}

func (o TopologyProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologyProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["elementId"] = o.ElementId
	if !IsNil(o.DependsOn) {
		toSerialize["dependsOn"] = o.DependsOn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopologyProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"elementId",
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

	varTopologyProperties := _TopologyProperties{}

	err = json.Unmarshal(data, &varTopologyProperties)

	if err != nil {
		return err
	}

	*o = TopologyProperties(varTopologyProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "elementId")
		delete(additionalProperties, "dependsOn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopologyProperties struct {
	value *TopologyProperties
	isSet bool
}

func (v NullableTopologyProperties) Get() *TopologyProperties {
	return v.value
}

func (v *NullableTopologyProperties) Set(val *TopologyProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTopologyProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyProperties(val *TopologyProperties) *NullableTopologyProperties {
	return &NullableTopologyProperties{value: val, isSet: true}
}

func (v NullableTopologyProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
