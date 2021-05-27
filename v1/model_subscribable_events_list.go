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

// SubscribableEventsList struct for SubscribableEventsList
type SubscribableEventsList struct {
	SubscribableEvents *[]SubscribableEvent `json:"subscribable_events,omitempty"`
	Meta *Meta `json:"meta,omitempty"`
}

// NewSubscribableEventsList instantiates a new SubscribableEventsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribableEventsList() *SubscribableEventsList {
	this := SubscribableEventsList{}
	return &this
}

// NewSubscribableEventsListWithDefaults instantiates a new SubscribableEventsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribableEventsListWithDefaults() *SubscribableEventsList {
	this := SubscribableEventsList{}
	return &this
}

// GetSubscribableEvents returns the SubscribableEvents field value if set, zero value otherwise.
func (o *SubscribableEventsList) GetSubscribableEvents() []SubscribableEvent {
	if o == nil || o.SubscribableEvents == nil {
		var ret []SubscribableEvent
		return ret
	}
	return *o.SubscribableEvents
}

// GetSubscribableEventsOk returns a tuple with the SubscribableEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribableEventsList) GetSubscribableEventsOk() (*[]SubscribableEvent, bool) {
	if o == nil || o.SubscribableEvents == nil {
		return nil, false
	}
	return o.SubscribableEvents, true
}

// HasSubscribableEvents returns a boolean if a field has been set.
func (o *SubscribableEventsList) HasSubscribableEvents() bool {
	if o != nil && o.SubscribableEvents != nil {
		return true
	}

	return false
}

// SetSubscribableEvents gets a reference to the given []SubscribableEvent and assigns it to the SubscribableEvents field.
func (o *SubscribableEventsList) SetSubscribableEvents(v []SubscribableEvent) {
	o.SubscribableEvents = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *SubscribableEventsList) GetMeta() Meta {
	if o == nil || o.Meta == nil {
		var ret Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscribableEventsList) GetMetaOk() (*Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *SubscribableEventsList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Meta and assigns it to the Meta field.
func (o *SubscribableEventsList) SetMeta(v Meta) {
	o.Meta = &v
}

func (o SubscribableEventsList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SubscribableEvents != nil {
		toSerialize["subscribable_events"] = o.SubscribableEvents
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	return json.Marshal(toSerialize)
}

type NullableSubscribableEventsList struct {
	value *SubscribableEventsList
	isSet bool
}

func (v NullableSubscribableEventsList) Get() *SubscribableEventsList {
	return v.value
}

func (v *NullableSubscribableEventsList) Set(val *SubscribableEventsList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribableEventsList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribableEventsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribableEventsList(val *SubscribableEventsList) *NullableSubscribableEventsList {
	return &NullableSubscribableEventsList{value: val, isSet: true}
}

func (v NullableSubscribableEventsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribableEventsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

