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

// checks if the RouteAggregationChangeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteAggregationChangeData{}

// RouteAggregationChangeData Current state of latest Route Aggregation change
type RouteAggregationChangeData struct {
	// Current outcome of the change flow
	Status *string `json:"status,omitempty"`
	// Created by User Key
	CreatedBy *string `json:"createdBy,omitempty"`
	// Set when change flow starts
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Updated by User Key
	UpdatedBy *string `json:"updatedBy,omitempty"`
	// Set when change object is updated
	UpdatedDateTime *time.Time `json:"updatedDateTime,omitempty"`
	// Additional information
	Information *string                           `json:"information,omitempty"`
	Data        *RouteAggregationsChangeOperation `json:"data,omitempty"`
	// Uniquely identifies a change
	Uuid string                      `json:"uuid"`
	Type RouteAggregationsChangeType `json:"type"`
	// Route AGGREGATION Change URI
	Href                 *string `json:"href,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RouteAggregationChangeData RouteAggregationChangeData

// NewRouteAggregationChangeData instantiates a new RouteAggregationChangeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteAggregationChangeData(uuid string, type_ RouteAggregationsChangeType) *RouteAggregationChangeData {
	this := RouteAggregationChangeData{}
	this.Uuid = uuid
	this.Type = type_
	return &this
}

// NewRouteAggregationChangeDataWithDefaults instantiates a new RouteAggregationChangeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteAggregationChangeDataWithDefaults() *RouteAggregationChangeData {
	this := RouteAggregationChangeData{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RouteAggregationChangeData) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *RouteAggregationChangeData) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetCreatedDateTime() time.Time {
	if o == nil || IsNil(o.CreatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDateTime) {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasCreatedDateTime() bool {
	if o != nil && !IsNil(o.CreatedDateTime) {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *RouteAggregationChangeData) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetUpdatedBy() string {
	if o == nil || IsNil(o.UpdatedBy) {
		var ret string
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedBy) {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasUpdatedBy() bool {
	if o != nil && !IsNil(o.UpdatedBy) {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given string and assigns it to the UpdatedBy field.
func (o *RouteAggregationChangeData) SetUpdatedBy(v string) {
	o.UpdatedBy = &v
}

// GetUpdatedDateTime returns the UpdatedDateTime field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetUpdatedDateTime() time.Time {
	if o == nil || IsNil(o.UpdatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDateTime
}

// GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDateTime) {
		return nil, false
	}
	return o.UpdatedDateTime, true
}

// HasUpdatedDateTime returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasUpdatedDateTime() bool {
	if o != nil && !IsNil(o.UpdatedDateTime) {
		return true
	}

	return false
}

// SetUpdatedDateTime gets a reference to the given time.Time and assigns it to the UpdatedDateTime field.
func (o *RouteAggregationChangeData) SetUpdatedDateTime(v time.Time) {
	o.UpdatedDateTime = &v
}

// GetInformation returns the Information field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetInformation() string {
	if o == nil || IsNil(o.Information) {
		var ret string
		return ret
	}
	return *o.Information
}

// GetInformationOk returns a tuple with the Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetInformationOk() (*string, bool) {
	if o == nil || IsNil(o.Information) {
		return nil, false
	}
	return o.Information, true
}

// HasInformation returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasInformation() bool {
	if o != nil && !IsNil(o.Information) {
		return true
	}

	return false
}

// SetInformation gets a reference to the given string and assigns it to the Information field.
func (o *RouteAggregationChangeData) SetInformation(v string) {
	o.Information = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetData() RouteAggregationsChangeOperation {
	if o == nil || IsNil(o.Data) {
		var ret RouteAggregationsChangeOperation
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetDataOk() (*RouteAggregationsChangeOperation, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RouteAggregationsChangeOperation and assigns it to the Data field.
func (o *RouteAggregationChangeData) SetData(v RouteAggregationsChangeOperation) {
	o.Data = &v
}

// GetUuid returns the Uuid field value
func (o *RouteAggregationChangeData) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *RouteAggregationChangeData) SetUuid(v string) {
	o.Uuid = v
}

// GetType returns the Type field value
func (o *RouteAggregationChangeData) GetType() RouteAggregationsChangeType {
	if o == nil {
		var ret RouteAggregationsChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetTypeOk() (*RouteAggregationsChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RouteAggregationChangeData) SetType(v RouteAggregationsChangeType) {
	o.Type = v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *RouteAggregationChangeData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteAggregationChangeData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *RouteAggregationChangeData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *RouteAggregationChangeData) SetHref(v string) {
	o.Href = &v
}

func (o RouteAggregationChangeData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteAggregationChangeData) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Information) {
		toSerialize["information"] = o.Information
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

func (o *RouteAggregationChangeData) UnmarshalJSON(data []byte) (err error) {
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

	varRouteAggregationChangeData := _RouteAggregationChangeData{}

	err = json.Unmarshal(data, &varRouteAggregationChangeData)

	if err != nil {
		return err
	}

	*o = RouteAggregationChangeData(varRouteAggregationChangeData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "createdDateTime")
		delete(additionalProperties, "updatedBy")
		delete(additionalProperties, "updatedDateTime")
		delete(additionalProperties, "information")
		delete(additionalProperties, "data")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "href")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRouteAggregationChangeData struct {
	value *RouteAggregationChangeData
	isSet bool
}

func (v NullableRouteAggregationChangeData) Get() *RouteAggregationChangeData {
	return v.value
}

func (v *NullableRouteAggregationChangeData) Set(val *RouteAggregationChangeData) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteAggregationChangeData) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteAggregationChangeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteAggregationChangeData(val *RouteAggregationChangeData) *NullableRouteAggregationChangeData {
	return &NullableRouteAggregationChangeData{value: val, isSet: true}
}

func (v NullableRouteAggregationChangeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteAggregationChangeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
