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

// checks if the DeviceActionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceActionInput{}

// DeviceActionInput struct for DeviceActionInput
type DeviceActionInput struct {
	Type DeviceActionInputType `json:"type"`
	// May be required to perform actions under certain conditions
	ForceDelete *bool `json:"force_delete,omitempty"`
	// When type is `reinstall`, enabling fast deprovisioning will bypass full disk wiping.
	DeprovisionFast *bool `json:"deprovision_fast,omitempty"`
	// When type is `reinstall`, preserve the existing data on all disks except the operating-system disk.
	PreserveData *bool `json:"preserve_data,omitempty"`
	// When type is `reinstall`, use this `operating_system` (defaults to the current `operating system`)
	OperatingSystem *string `json:"operating_system,omitempty"`
	// When type is `reinstall`, use this `ipxe_script_url` (`operating_system` must be `custom_ipxe`, defaults to the current `ipxe_script_url`)
	IpxeScriptUrl        *string `json:"ipxe_script_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeviceActionInput DeviceActionInput

// NewDeviceActionInput instantiates a new DeviceActionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceActionInput(type_ DeviceActionInputType) *DeviceActionInput {
	this := DeviceActionInput{}
	this.Type = type_
	return &this
}

// NewDeviceActionInputWithDefaults instantiates a new DeviceActionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceActionInputWithDefaults() *DeviceActionInput {
	this := DeviceActionInput{}
	return &this
}

// GetType returns the Type field value
func (o *DeviceActionInput) GetType() DeviceActionInputType {
	if o == nil {
		var ret DeviceActionInputType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetTypeOk() (*DeviceActionInputType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DeviceActionInput) SetType(v DeviceActionInputType) {
	o.Type = v
}

// GetForceDelete returns the ForceDelete field value if set, zero value otherwise.
func (o *DeviceActionInput) GetForceDelete() bool {
	if o == nil || IsNil(o.ForceDelete) {
		var ret bool
		return ret
	}
	return *o.ForceDelete
}

// GetForceDeleteOk returns a tuple with the ForceDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetForceDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceDelete) {
		return nil, false
	}
	return o.ForceDelete, true
}

// HasForceDelete returns a boolean if a field has been set.
func (o *DeviceActionInput) HasForceDelete() bool {
	if o != nil && !IsNil(o.ForceDelete) {
		return true
	}

	return false
}

// SetForceDelete gets a reference to the given bool and assigns it to the ForceDelete field.
func (o *DeviceActionInput) SetForceDelete(v bool) {
	o.ForceDelete = &v
}

// GetDeprovisionFast returns the DeprovisionFast field value if set, zero value otherwise.
func (o *DeviceActionInput) GetDeprovisionFast() bool {
	if o == nil || IsNil(o.DeprovisionFast) {
		var ret bool
		return ret
	}
	return *o.DeprovisionFast
}

// GetDeprovisionFastOk returns a tuple with the DeprovisionFast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetDeprovisionFastOk() (*bool, bool) {
	if o == nil || IsNil(o.DeprovisionFast) {
		return nil, false
	}
	return o.DeprovisionFast, true
}

// HasDeprovisionFast returns a boolean if a field has been set.
func (o *DeviceActionInput) HasDeprovisionFast() bool {
	if o != nil && !IsNil(o.DeprovisionFast) {
		return true
	}

	return false
}

// SetDeprovisionFast gets a reference to the given bool and assigns it to the DeprovisionFast field.
func (o *DeviceActionInput) SetDeprovisionFast(v bool) {
	o.DeprovisionFast = &v
}

// GetPreserveData returns the PreserveData field value if set, zero value otherwise.
func (o *DeviceActionInput) GetPreserveData() bool {
	if o == nil || IsNil(o.PreserveData) {
		var ret bool
		return ret
	}
	return *o.PreserveData
}

// GetPreserveDataOk returns a tuple with the PreserveData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetPreserveDataOk() (*bool, bool) {
	if o == nil || IsNil(o.PreserveData) {
		return nil, false
	}
	return o.PreserveData, true
}

// HasPreserveData returns a boolean if a field has been set.
func (o *DeviceActionInput) HasPreserveData() bool {
	if o != nil && !IsNil(o.PreserveData) {
		return true
	}

	return false
}

// SetPreserveData gets a reference to the given bool and assigns it to the PreserveData field.
func (o *DeviceActionInput) SetPreserveData(v bool) {
	o.PreserveData = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *DeviceActionInput) GetOperatingSystem() string {
	if o == nil || IsNil(o.OperatingSystem) {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetOperatingSystemOk() (*string, bool) {
	if o == nil || IsNil(o.OperatingSystem) {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *DeviceActionInput) HasOperatingSystem() bool {
	if o != nil && !IsNil(o.OperatingSystem) {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *DeviceActionInput) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetIpxeScriptUrl returns the IpxeScriptUrl field value if set, zero value otherwise.
func (o *DeviceActionInput) GetIpxeScriptUrl() string {
	if o == nil || IsNil(o.IpxeScriptUrl) {
		var ret string
		return ret
	}
	return *o.IpxeScriptUrl
}

// GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActionInput) GetIpxeScriptUrlOk() (*string, bool) {
	if o == nil || IsNil(o.IpxeScriptUrl) {
		return nil, false
	}
	return o.IpxeScriptUrl, true
}

// HasIpxeScriptUrl returns a boolean if a field has been set.
func (o *DeviceActionInput) HasIpxeScriptUrl() bool {
	if o != nil && !IsNil(o.IpxeScriptUrl) {
		return true
	}

	return false
}

// SetIpxeScriptUrl gets a reference to the given string and assigns it to the IpxeScriptUrl field.
func (o *DeviceActionInput) SetIpxeScriptUrl(v string) {
	o.IpxeScriptUrl = &v
}

func (o DeviceActionInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceActionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.ForceDelete) {
		toSerialize["force_delete"] = o.ForceDelete
	}
	if !IsNil(o.DeprovisionFast) {
		toSerialize["deprovision_fast"] = o.DeprovisionFast
	}
	if !IsNil(o.PreserveData) {
		toSerialize["preserve_data"] = o.PreserveData
	}
	if !IsNil(o.OperatingSystem) {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if !IsNil(o.IpxeScriptUrl) {
		toSerialize["ipxe_script_url"] = o.IpxeScriptUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeviceActionInput) UnmarshalJSON(data []byte) (err error) {
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

	varDeviceActionInput := _DeviceActionInput{}

	err = json.Unmarshal(data, &varDeviceActionInput)

	if err != nil {
		return err
	}

	*o = DeviceActionInput(varDeviceActionInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "force_delete")
		delete(additionalProperties, "deprovision_fast")
		delete(additionalProperties, "preserve_data")
		delete(additionalProperties, "operating_system")
		delete(additionalProperties, "ipxe_script_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeviceActionInput struct {
	value *DeviceActionInput
	isSet bool
}

func (v NullableDeviceActionInput) Get() *DeviceActionInput {
	return v.value
}

func (v *NullableDeviceActionInput) Set(val *DeviceActionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceActionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceActionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceActionInput(val *DeviceActionInput) *NullableDeviceActionInput {
	return &NullableDeviceActionInput{value: val, isSet: true}
}

func (v NullableDeviceActionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceActionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
