/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
	"time"
)

// checks if the RouteAggregationRulesChangeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteAggregationRulesChangeData{}

// RouteAggregationRulesChangeData Current state of latest Route Aggregation Rules change
type RouteAggregationRulesChangeData struct {
	// Current outcome of the change flow
	Status *string `json:"status,omitempty"`
	// Created by User Key
	CreatedBy *string `json:"createdBy,omitempty"`
	// Set when change flow starts
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Updated by User Key
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// Set when change object is updated
	UpdatedDateTime *time.Time                            `json:"updatedDateTime,omitempty"`
	Data            *RouteAggregationRulesChangeOperation `json:"data,omitempty"`
	// Uniquely identifies a change
	Uuid string                          `json:"uuid"`
	Type RouteAggregationRulesChangeType `json:"type"`
	// Route Aggregation Change URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteAggregationRulesChangeData RouteAggregationRulesChangeData

// NewRouteAggregationRulesChangeData instantiates a new RouteAggregationRulesChangeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteAggregationRulesChangeData(uuid string, type_ RouteAggregationRulesChangeType) *RouteAggregationRulesChangeData {
	this := RouteAggregationRulesChangeData{}
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewRouteAggregationRulesChangeDataWithDefaults instantiates a new RouteAggregationRulesChangeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteAggregationRulesChangeDataWithDefaults() *RouteAggregationRulesChangeData {
	this := RouteAggregationRulesChangeData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RouteAggregationRulesChangeData) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RouteAggregationRulesChangeData) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetCreatedDateTime() time.Time {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *RouteAggregationRulesChangeData) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RouteAggregationRulesChangeData) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedDateTime returns the UpdatedDateTime field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetUpdatedDateTime() time.Time {
	if o == nil || IsNil(o.UpdatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDateTime
}

// GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDateTime) {
		return nil, false
	}
	return o.UpdatedDateTime, true
}

// HasUpdatedDateTime returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasUpdatedDateTime() bool {
	if o != nil && !IsNil(o.UpdatedDateTime) {
		return true
	}

	return false
}

// SetUpdatedDateTime gets a reference to the given time.Time and assigns it to the UpdatedDateTime field.
func (o *RouteAggregationRulesChangeData) SetUpdatedDateTime(v time.Time) {
	o.UpdatedDateTime = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetData() RouteAggregationRulesChangeOperation {
	if o == nil || IsNil(o.Data) {
		var ret RouteAggregationRulesChangeOperation
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetDataOk() (*RouteAggregationRulesChangeOperation, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RouteAggregationRulesChangeOperation and assigns it to the Data field.
func (o *RouteAggregationRulesChangeData) SetData(v RouteAggregationRulesChangeOperation) {
	o.Data = &v
}

// GetUuid returns the Uuid field value
func (o *RouteAggregationRulesChangeData) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RouteAggregationRulesChangeData) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *RouteAggregationRulesChangeData) GetType() RouteAggregationRulesChangeType {
	if o == nil {
		var ret RouteAggregationRulesChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetTypeOk() (*RouteAggregationRulesChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteAggregationRulesChangeData) SetType(v RouteAggregationRulesChangeType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteAggregationRulesChangeData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationRulesChangeData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteAggregationRulesChangeData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteAggregationRulesChangeData) SetHref(v string) {
	o.Href = &v
}

func (o RouteAggregationRulesChangeData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteAggregationRulesChangeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.CreatedDateTime) {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if !IsNil(o.UpdatedBy) {
		toSerialize["updatedBy"] = o.UpdatedBy
	}
	if !IsNil(o.UpdatedDateTime) {
		toSerialize["updatedDateTime"] = o.UpdatedDateTime
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["uuid"] = o.Uuid
	toSerialize["type"] = o.Type
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RouteAggregationRulesChangeData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varRouteAggregationRulesChangeData := _RouteAggregationRulesChangeData{}

	err = json.Unmarshal(data, &varRouteAggregationRulesChangeData)

	if err != nil {
		return err
	}

	*o = RouteAggregationRulesChangeData(varRouteAggregationRulesChangeData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdDateTime")
		delete(additionalProperties, "updatedBy")
		delete(additionalProperties, "updatedDateTime")
		delete(additionalProperties, "data")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteAggregationRulesChangeData struct {
	value *RouteAggregationRulesChangeData
	isSet bool
}

func (v NullableRouteAggregationRulesChangeData) Get() *RouteAggregationRulesChangeData {
	return v.value
}

func (v *NullableRouteAggregationRulesChangeData) Set(val *RouteAggregationRulesChangeData) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteAggregationRulesChangeData) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteAggregationRulesChangeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteAggregationRulesChangeData(val *RouteAggregationRulesChangeData) *NullableRouteAggregationRulesChangeData {
	return &NullableRouteAggregationRulesChangeData{value: val, isSet: true}
}

func (v NullableRouteAggregationRulesChangeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteAggregationRulesChangeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}