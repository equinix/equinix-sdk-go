/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RouteAggregationsDataProject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteAggregationsDataProject{}

// RouteAggregationsDataProject struct for RouteAggregationsDataProject
type RouteAggregationsDataProject struct {
	// Subscriber-assigned project ID
	ProjectId string `json:"projectId"`
	// Project URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteAggregationsDataProject RouteAggregationsDataProject

// NewRouteAggregationsDataProject instantiates a new RouteAggregationsDataProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteAggregationsDataProject(projectId string) *RouteAggregationsDataProject {
	this := RouteAggregationsDataProject{}
	this.ProjectId = projectId
	return &this
}

// NewRouteAggregationsDataProjectWithDefaults instantiates a new RouteAggregationsDataProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteAggregationsDataProjectWithDefaults() *RouteAggregationsDataProject {
	this := RouteAggregationsDataProject{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *RouteAggregationsDataProject) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationsDataProject) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *RouteAggregationsDataProject) SetProjectId(v string) {
	o.ProjectId = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteAggregationsDataProject) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationsDataProject) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteAggregationsDataProject) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteAggregationsDataProject) SetHref(v string) {
	o.Href = &v
}

func (o RouteAggregationsDataProject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteAggregationsDataProject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteAggregationsDataProject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"projectId",
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

	varRouteAggregationsDataProject := _RouteAggregationsDataProject{}

	err = json.Unmarshal(data, &varRouteAggregationsDataProject)

	if err != nil {
		return err
	}

	*o = RouteAggregationsDataProject(varRouteAggregationsDataProject)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "projectId")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteAggregationsDataProject struct {
	value *RouteAggregationsDataProject
	isSet bool
}

func (v NullableRouteAggregationsDataProject) Get() *RouteAggregationsDataProject {
	return v.value
}

func (v *NullableRouteAggregationsDataProject) Set(val *RouteAggregationsDataProject) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteAggregationsDataProject) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteAggregationsDataProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteAggregationsDataProject(val *RouteAggregationsDataProject) *NullableRouteAggregationsDataProject {
	return &NullableRouteAggregationsDataProject{value: val, isSet: true}
}

func (v NullableRouteAggregationsDataProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteAggregationsDataProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}