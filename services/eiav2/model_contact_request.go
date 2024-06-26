/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the ContactRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContactRequest{}

// ContactRequest struct for ContactRequest
type ContactRequest struct {
	Type                 ContactType                 `json:"type"`
	FirstName            *string                     `json:"firstName,omitempty"`
	LastName             *string                     `json:"lastName,omitempty"`
	Timezone             *string                     `json:"timezone,omitempty"`
	Notes                *string                     `json:"notes,omitempty"`
	Availability         *ContactRequestAvailability `json:"availability,omitempty"`
	Details              []ContactRequestDetails     `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ContactRequest ContactRequest

// NewContactRequest instantiates a new ContactRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContactRequest(type_ ContactType) *ContactRequest {
	this := ContactRequest{}
	this.Type = type_
	return &this
}

// NewContactRequestWithDefaults instantiates a new ContactRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContactRequestWithDefaults() *ContactRequest {
	this := ContactRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ContactRequest) GetType() ContactType {
	if o == nil {
		var ret ContactType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetTypeOk() (*ContactType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ContactRequest) SetType(v ContactType) {
	o.Type = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ContactRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ContactRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ContactRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ContactRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ContactRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ContactRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ContactRequest) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ContactRequest) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ContactRequest) SetTimezone(v string) {
	o.Timezone = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ContactRequest) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ContactRequest) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ContactRequest) SetNotes(v string) {
	o.Notes = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ContactRequest) GetAvailability() ContactRequestAvailability {
	if o == nil || IsNil(o.Availability) {
		var ret ContactRequestAvailability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetAvailabilityOk() (*ContactRequestAvailability, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ContactRequest) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given ContactRequestAvailability and assigns it to the Availability field.
func (o *ContactRequest) SetAvailability(v ContactRequestAvailability) {
	o.Availability = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ContactRequest) GetDetails() []ContactRequestDetails {
	if o == nil || IsNil(o.Details) {
		var ret []ContactRequestDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContactRequest) GetDetailsOk() ([]ContactRequestDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ContactRequest) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ContactRequestDetails and assigns it to the Details field.
func (o *ContactRequest) SetDetails(v []ContactRequestDetails) {
	o.Details = v
}

func (o ContactRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContactRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ContactRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varContactRequest := _ContactRequest{}

	err = json.Unmarshal(data, &varContactRequest)

	if err != nil {
		return err
	}

	*o = ContactRequest(varContactRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "availability")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableContactRequest struct {
	value *ContactRequest
	isSet bool
}

func (v NullableContactRequest) Get() *ContactRequest {
	return v.value
}

func (v *NullableContactRequest) Set(val *ContactRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContactRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContactRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContactRequest(val *ContactRequest) *NullableContactRequest {
	return &NullableContactRequest{value: val, isSet: true}
}

func (v NullableContactRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContactRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
