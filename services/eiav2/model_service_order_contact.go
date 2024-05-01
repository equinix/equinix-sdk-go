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

// checks if the ServiceOrderContact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceOrderContact{}

// ServiceOrderContact struct for ServiceOrderContact
type ServiceOrderContact struct {
	Type         ContactType                 `json:"type"`
	FirstName    *string                     `json:"firstName,omitempty"`
	LastName     *string                     `json:"lastName,omitempty"`
	Timezone     *string                     `json:"timezone,omitempty"`
	Notes        *string                     `json:"notes,omitempty"`
	Availability *ContactRequestAvailability `json:"availability,omitempty"`
	Details      []ContactRequestDetails     `json:"details,omitempty"`
	// Identifies (e.g., userName) a registered user. If a registered user is specified, then firstName/lastName need not be provided
	RegisteredUser       *string `json:"registeredUser,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceOrderContact ServiceOrderContact

// NewServiceOrderContact instantiates a new ServiceOrderContact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceOrderContact(type_ ContactType) *ServiceOrderContact {
	this := ServiceOrderContact{}
	this.Type = type_
	return &this
}

// NewServiceOrderContactWithDefaults instantiates a new ServiceOrderContact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceOrderContactWithDefaults() *ServiceOrderContact {
	this := ServiceOrderContact{}
	return &this
}

// GetType returns the Type field value
func (o *ServiceOrderContact) GetType() ContactType {
	if o == nil {
		var ret ContactType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetTypeOk() (*ContactType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceOrderContact) SetType(v ContactType) {
	o.Type = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ServiceOrderContact) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ServiceOrderContact) SetLastName(v string) {
	o.LastName = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ServiceOrderContact) SetTimezone(v string) {
	o.Timezone = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *ServiceOrderContact) SetNotes(v string) {
	o.Notes = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetAvailability() ContactRequestAvailability {
	if o == nil || IsNil(o.Availability) {
		var ret ContactRequestAvailability
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetAvailabilityOk() (*ContactRequestAvailability, bool) {
	if o == nil || IsNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasAvailability() bool {
	if o != nil && !IsNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given ContactRequestAvailability and assigns it to the Availability field.
func (o *ServiceOrderContact) SetAvailability(v ContactRequestAvailability) {
	o.Availability = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetDetails() []ContactRequestDetails {
	if o == nil || IsNil(o.Details) {
		var ret []ContactRequestDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetDetailsOk() ([]ContactRequestDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ContactRequestDetails and assigns it to the Details field.
func (o *ServiceOrderContact) SetDetails(v []ContactRequestDetails) {
	o.Details = v
}

// GetRegisteredUser returns the RegisteredUser field value if set, zero value otherwise.
func (o *ServiceOrderContact) GetRegisteredUser() string {
	if o == nil || IsNil(o.RegisteredUser) {
		var ret string
		return ret
	}
	return *o.RegisteredUser
}

// GetRegisteredUserOk returns a tuple with the RegisteredUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceOrderContact) GetRegisteredUserOk() (*string, bool) {
	if o == nil || IsNil(o.RegisteredUser) {
		return nil, false
	}
	return o.RegisteredUser, true
}

// HasRegisteredUser returns a boolean if a field has been set.
func (o *ServiceOrderContact) HasRegisteredUser() bool {
	if o != nil && !IsNil(o.RegisteredUser) {
		return true
	}

	return false
}

// SetRegisteredUser gets a reference to the given string and assigns it to the RegisteredUser field.
func (o *ServiceOrderContact) SetRegisteredUser(v string) {
	o.RegisteredUser = &v
}

func (o ServiceOrderContact) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceOrderContact) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.RegisteredUser) {
		toSerialize["registeredUser"] = o.RegisteredUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ServiceOrderContact) UnmarshalJSON(data []byte) (err error) {
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

	varServiceOrderContact := _ServiceOrderContact{}

	err = json.Unmarshal(data, &varServiceOrderContact)

	if err != nil {
		return err
	}

	*o = ServiceOrderContact(varServiceOrderContact)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "timezone")
		delete(additionalProperties, "notes")
		delete(additionalProperties, "availability")
		delete(additionalProperties, "details")
		delete(additionalProperties, "registeredUser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceOrderContact struct {
	value *ServiceOrderContact
	isSet bool
}

func (v NullableServiceOrderContact) Get() *ServiceOrderContact {
	return v.value
}

func (v *NullableServiceOrderContact) Set(val *ServiceOrderContact) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceOrderContact) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceOrderContact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceOrderContact(val *ServiceOrderContact) *NullableServiceOrderContact {
	return &NullableServiceOrderContact{value: val, isSet: true}
}

func (v NullableServiceOrderContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceOrderContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}