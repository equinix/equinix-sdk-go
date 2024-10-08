/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the VlanVirtualCircuitCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VlanVirtualCircuitCreateInput{}

// VlanVirtualCircuitCreateInput struct for VlanVirtualCircuitCreateInput
type VlanVirtualCircuitCreateInput struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
	NniVlan     *int32  `json:"nni_vlan,omitempty"`
	ProjectId   string  `json:"project_id"`
	// speed can be passed as integer number representing bps speed or string (e.g. '52m' or '100g' or '4 gbps')
	Speed *string  `json:"speed,omitempty"`
	Tags  []string `json:"tags,omitempty"`
	// A Virtual Network record UUID or the VNID of a Metro Virtual Network in your project (sent as integer).
	Vnid                 *string `json:"vnid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VlanVirtualCircuitCreateInput VlanVirtualCircuitCreateInput

// NewVlanVirtualCircuitCreateInput instantiates a new VlanVirtualCircuitCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVirtualCircuitCreateInput(projectId string) *VlanVirtualCircuitCreateInput {
	this := VlanVirtualCircuitCreateInput{}
	this.ProjectId = projectId
	return &this
}

// NewVlanVirtualCircuitCreateInputWithDefaults instantiates a new VlanVirtualCircuitCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVirtualCircuitCreateInputWithDefaults() *VlanVirtualCircuitCreateInput {
	this := VlanVirtualCircuitCreateInput{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VlanVirtualCircuitCreateInput) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VlanVirtualCircuitCreateInput) SetName(v string) {
	o.Name = &v
}

// GetNniVlan returns the NniVlan field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetNniVlan() int32 {
	if o == nil || IsNil(o.NniVlan) {
		var ret int32
		return ret
	}
	return *o.NniVlan
}

// GetNniVlanOk returns a tuple with the NniVlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetNniVlanOk() (*int32, bool) {
	if o == nil || IsNil(o.NniVlan) {
		return nil, false
	}
	return o.NniVlan, true
}

// HasNniVlan returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasNniVlan() bool {
	if o != nil && !IsNil(o.NniVlan) {
		return true
	}

	return false
}

// SetNniVlan gets a reference to the given int32 and assigns it to the NniVlan field.
func (o *VlanVirtualCircuitCreateInput) SetNniVlan(v int32) {
	o.NniVlan = &v
}

// GetProjectId returns the ProjectId field value
func (o *VlanVirtualCircuitCreateInput) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *VlanVirtualCircuitCreateInput) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *VlanVirtualCircuitCreateInput) SetSpeed(v string) {
	o.Speed = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *VlanVirtualCircuitCreateInput) SetTags(v []string) {
	o.Tags = v
}

// GetVnid returns the Vnid field value if set, zero value otherwise.
func (o *VlanVirtualCircuitCreateInput) GetVnid() string {
	if o == nil || IsNil(o.Vnid) {
		var ret string
		return ret
	}
	return *o.Vnid
}

// GetVnidOk returns a tuple with the Vnid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVirtualCircuitCreateInput) GetVnidOk() (*string, bool) {
	if o == nil || IsNil(o.Vnid) {
		return nil, false
	}
	return o.Vnid, true
}

// HasVnid returns a boolean if a field has been set.
func (o *VlanVirtualCircuitCreateInput) HasVnid() bool {
	if o != nil && !IsNil(o.Vnid) {
		return true
	}

	return false
}

// SetVnid gets a reference to the given string and assigns it to the Vnid field.
func (o *VlanVirtualCircuitCreateInput) SetVnid(v string) {
	o.Vnid = &v
}

func (o VlanVirtualCircuitCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VlanVirtualCircuitCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.NniVlan) {
		toSerialize["nni_vlan"] = o.NniVlan
	}
	toSerialize["project_id"] = o.ProjectId
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Vnid) {
		toSerialize["vnid"] = o.Vnid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *VlanVirtualCircuitCreateInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"project_id",
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

	varVlanVirtualCircuitCreateInput := _VlanVirtualCircuitCreateInput{}

	err = json.Unmarshal(data, &varVlanVirtualCircuitCreateInput)

	if err != nil {
		return err
	}

	*o = VlanVirtualCircuitCreateInput(varVlanVirtualCircuitCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nni_vlan")
		delete(additionalProperties, "project_id")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "vnid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableVlanVirtualCircuitCreateInput struct {
	value *VlanVirtualCircuitCreateInput
	isSet bool
}

func (v NullableVlanVirtualCircuitCreateInput) Get() *VlanVirtualCircuitCreateInput {
	return v.value
}

func (v *NullableVlanVirtualCircuitCreateInput) Set(val *VlanVirtualCircuitCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVirtualCircuitCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVirtualCircuitCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVirtualCircuitCreateInput(val *VlanVirtualCircuitCreateInput) *NullableVlanVirtualCircuitCreateInput {
	return &NullableVlanVirtualCircuitCreateInput{value: val, isSet: true}
}

func (v NullableVlanVirtualCircuitCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVirtualCircuitCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
