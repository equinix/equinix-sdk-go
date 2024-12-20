/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the Interface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Interface{}

// Interface Interface Information
type Interface struct {
	// Interface URI
	Href *string `json:"href,omitempty"`
	// Equinix-assigned Interface identifier
	Uuid *string `json:"uuid,omitempty"`
	// Interface id
	Id   *int32         `json:"id,omitempty"`
	Type *InterfaceType `json:"type,omitempty"`
	// Interface Project ID
	ProjectId            *string `json:"projectId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Interface Interface

// NewInterface instantiates a new Interface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterface() *Interface {
	this := Interface{}
	return &this
}

// NewInterfaceWithDefaults instantiates a new Interface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceWithDefaults() *Interface {
	this := Interface{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *Interface) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *Interface) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *Interface) SetHref(v string) {
	o.Href = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Interface) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Interface) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Interface) SetUuid(v string) {
	o.Uuid = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Interface) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Interface) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Interface) SetId(v int32) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Interface) GetType() InterfaceType {
	if o == nil || IsNil(o.Type) {
		var ret InterfaceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetTypeOk() (*InterfaceType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Interface) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given InterfaceType and assigns it to the Type field.
func (o *Interface) SetType(v InterfaceType) {
	o.Type = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *Interface) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Interface) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *Interface) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *Interface) SetProjectId(v string) {
	o.ProjectId = &v
}

func (o Interface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Interface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Interface) UnmarshalJSON(data []byte) (err error) {
	varInterface := _Interface{}

	err = json.Unmarshal(data, &varInterface)

	if err != nil {
		return err
	}

	*o = Interface(varInterface)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "projectId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterface struct {
	value *Interface
	isSet bool
}

func (v NullableInterface) Get() *Interface {
	return v.value
}

func (v *NullableInterface) Set(val *Interface) {
	v.value = val
	v.isSet = true
}

func (v NullableInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterface(val *Interface) *NullableInterface {
	return &NullableInterface{value: val, isSet: true}
}

func (v NullableInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
