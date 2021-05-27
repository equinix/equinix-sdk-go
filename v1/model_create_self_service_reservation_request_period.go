/*
 * Metal API
 *
 * This is the API for Equinix Metal. The API allows you to programmatically interact with all of your Equinix Metal resources, including devices, networks, addresses, organizations, projects, and your user account.  The official API docs are hosted at <https://metal.equinix.com/developers/api>. 
 *
 * API version: 1.0.0
 * Contact: support@equinixmetal.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

// CreateSelfServiceReservationRequestPeriod struct for CreateSelfServiceReservationRequestPeriod
type CreateSelfServiceReservationRequestPeriod struct {
	Unit *string `json:"unit,omitempty"`
	Count *float32 `json:"count,omitempty"`
}

// NewCreateSelfServiceReservationRequestPeriod instantiates a new CreateSelfServiceReservationRequestPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSelfServiceReservationRequestPeriod() *CreateSelfServiceReservationRequestPeriod {
	this := CreateSelfServiceReservationRequestPeriod{}
	return &this
}

// NewCreateSelfServiceReservationRequestPeriodWithDefaults instantiates a new CreateSelfServiceReservationRequestPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSelfServiceReservationRequestPeriodWithDefaults() *CreateSelfServiceReservationRequestPeriod {
	this := CreateSelfServiceReservationRequestPeriod{}
	return &this
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequestPeriod) GetUnit() string {
	if o == nil || o.Unit == nil {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequestPeriod) GetUnitOk() (*string, bool) {
	if o == nil || o.Unit == nil {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequestPeriod) HasUnit() bool {
	if o != nil && o.Unit != nil {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *CreateSelfServiceReservationRequestPeriod) SetUnit(v string) {
	o.Unit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CreateSelfServiceReservationRequestPeriod) GetCount() float32 {
	if o == nil || o.Count == nil {
		var ret float32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSelfServiceReservationRequestPeriod) GetCountOk() (*float32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CreateSelfServiceReservationRequestPeriod) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given float32 and assigns it to the Count field.
func (o *CreateSelfServiceReservationRequestPeriod) SetCount(v float32) {
	o.Count = &v
}

func (o CreateSelfServiceReservationRequestPeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Unit != nil {
		toSerialize["unit"] = o.Unit
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableCreateSelfServiceReservationRequestPeriod struct {
	value *CreateSelfServiceReservationRequestPeriod
	isSet bool
}

func (v NullableCreateSelfServiceReservationRequestPeriod) Get() *CreateSelfServiceReservationRequestPeriod {
	return v.value
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) Set(val *CreateSelfServiceReservationRequestPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSelfServiceReservationRequestPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSelfServiceReservationRequestPeriod(val *CreateSelfServiceReservationRequestPeriod) *NullableCreateSelfServiceReservationRequestPeriod {
	return &NullableCreateSelfServiceReservationRequestPeriod{value: val, isSet: true}
}

func (v NullableCreateSelfServiceReservationRequestPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSelfServiceReservationRequestPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

