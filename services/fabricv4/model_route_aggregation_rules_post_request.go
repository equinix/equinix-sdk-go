/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the RouteAggregationRulesPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteAggregationRulesPostRequest{}

// RouteAggregationRulesPostRequest Create Route Aggregation Rule POST request
type RouteAggregationRulesPostRequest struct {
	// Route Aggregation Rule configuration
	Data                 []RouteAggregationRulesBase `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteAggregationRulesPostRequest RouteAggregationRulesPostRequest

// NewRouteAggregationRulesPostRequest instantiates a new RouteAggregationRulesPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteAggregationRulesPostRequest() *RouteAggregationRulesPostRequest {
	this := RouteAggregationRulesPostRequest{}
	return &this
}

// NewRouteAggregationRulesPostRequestWithDefaults instantiates a new RouteAggregationRulesPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteAggregationRulesPostRequestWithDefaults() *RouteAggregationRulesPostRequest {
	this := RouteAggregationRulesPostRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RouteAggregationRulesPostRequest) GetData() []RouteAggregationRulesBase {
	if o == nil || IsNil(o.Data) {
		var ret []RouteAggregationRulesBase
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesPostRequest) GetDataOk() ([]RouteAggregationRulesBase, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RouteAggregationRulesPostRequest) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RouteAggregationRulesBase and assigns it to the Data field.
func (o *RouteAggregationRulesPostRequest) SetData(v []RouteAggregationRulesBase) {
	o.Data = v
}

func (o RouteAggregationRulesPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteAggregationRulesPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteAggregationRulesPostRequest) UnmarshalJSON(data []byte) (err error) {
	varRouteAggregationRulesPostRequest := _RouteAggregationRulesPostRequest{}

	err = json.Unmarshal(data, &varRouteAggregationRulesPostRequest)

	if err != nil {
		return err
	}

	*o = RouteAggregationRulesPostRequest(varRouteAggregationRulesPostRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteAggregationRulesPostRequest struct {
	value *RouteAggregationRulesPostRequest
	isSet bool
}

func (v NullableRouteAggregationRulesPostRequest) Get() *RouteAggregationRulesPostRequest {
	return v.value
}

func (v *NullableRouteAggregationRulesPostRequest) Set(val *RouteAggregationRulesPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteAggregationRulesPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteAggregationRulesPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteAggregationRulesPostRequest(val *RouteAggregationRulesPostRequest) *NullableRouteAggregationRulesPostRequest {
	return &NullableRouteAggregationRulesPostRequest{value: val, isSet: true}
}

func (v NullableRouteAggregationRulesPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteAggregationRulesPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}