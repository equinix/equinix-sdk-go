/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the BGPActionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BGPActionData{}

// BGPActionData struct for BGPActionData
type BGPActionData struct {
	// Routing Protocol URI
	Href *string `json:"href,omitempty"`
	// Routing protocol identifier
	Uuid *string     `json:"uuid,omitempty"`
	Type *BGPActions `json:"type,omitempty"`
	// BGP action description
	Description          *string          `json:"description,omitempty"`
	State                *BGPActionStates `json:"state,omitempty"`
	Changelog            *Changelog       `json:"changelog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BGPActionData BGPActionData

// NewBGPActionData instantiates a new BGPActionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBGPActionData() *BGPActionData {
	this := BGPActionData{}
	return &this
}

// NewBGPActionDataWithDefaults instantiates a new BGPActionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBGPActionDataWithDefaults() *BGPActionData {
	this := BGPActionData{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BGPActionData) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BGPActionData) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BGPActionData) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *BGPActionData) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *BGPActionData) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *BGPActionData) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BGPActionData) GetType() BGPActions {
	if o == nil || IsNil(o.Type) {
		var ret BGPActions
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetTypeOk() (*BGPActions, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BGPActionData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given BGPActions and assigns it to the Type field.
func (o *BGPActionData) SetType(v BGPActions) {
	o.Type = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BGPActionData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BGPActionData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BGPActionData) SetDescription(v string) {
	o.Description = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BGPActionData) GetState() BGPActionStates {
	if o == nil || IsNil(o.State) {
		var ret BGPActionStates
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetStateOk() (*BGPActionStates, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BGPActionData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BGPActionStates and assigns it to the State field.
func (o *BGPActionData) SetState(v BGPActionStates) {
	o.State = &v
}

// GetChangelog returns the Changelog field value if set, zero value otherwise.
func (o *BGPActionData) GetChangelog() Changelog {
	if o == nil || IsNil(o.Changelog) {
		var ret Changelog
		return ret
	}
	return *o.Changelog
}

// GetChangelogOk returns a tuple with the Changelog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BGPActionData) GetChangelogOk() (*Changelog, bool) {
	if o == nil || IsNil(o.Changelog) {
		return nil, false
	}
	return o.Changelog, true
}

// HasChangelog returns a boolean if a field has been set.
func (o *BGPActionData) HasChangelog() bool {
	if o != nil && !IsNil(o.Changelog) {
		return true
	}

	return false
}

// SetChangelog gets a reference to the given Changelog and assigns it to the Changelog field.
func (o *BGPActionData) SetChangelog(v Changelog) {
	o.Changelog = &v
}

func (o BGPActionData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BGPActionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Changelog) {
		toSerialize["changelog"] = o.Changelog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BGPActionData) UnmarshalJSON(data []byte) (err error) {
	varBGPActionData := _BGPActionData{}

	err = json.Unmarshal(data, &varBGPActionData)

	if err != nil {
		return err
	}

	*o = BGPActionData(varBGPActionData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "description")
		delete(additionalProperties, "state")
		delete(additionalProperties, "changelog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBGPActionData struct {
	value *BGPActionData
	isSet bool
}

func (v NullableBGPActionData) Get() *BGPActionData {
	return v.value
}

func (v *NullableBGPActionData) Set(val *BGPActionData) {
	v.value = val
	v.isSet = true
}

func (v NullableBGPActionData) IsSet() bool {
	return v.isSet
}

func (v *NullableBGPActionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBGPActionData(val *BGPActionData) *NullableBGPActionData {
	return &NullableBGPActionData{value: val, isSet: true}
}

func (v NullableBGPActionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBGPActionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
