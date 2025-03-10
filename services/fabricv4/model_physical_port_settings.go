/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PhysicalPortSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhysicalPortSettings{}

// PhysicalPortSettings Physical Port configuration settings
type PhysicalPortSettings struct {
	ErrorMessage *string `json:"errorMessage,omitempty"`
	// Deprecated
	PackageType          *string `json:"packageType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PhysicalPortSettings PhysicalPortSettings

// NewPhysicalPortSettings instantiates a new PhysicalPortSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalPortSettings() *PhysicalPortSettings {
	this := PhysicalPortSettings{}
	return &this
}

// NewPhysicalPortSettingsWithDefaults instantiates a new PhysicalPortSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalPortSettingsWithDefaults() *PhysicalPortSettings {
	this := PhysicalPortSettings{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *PhysicalPortSettings) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalPortSettings) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *PhysicalPortSettings) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *PhysicalPortSettings) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetPackageType returns the PackageType field value if set, zero value otherwise.
// Deprecated
func (o *PhysicalPortSettings) GetPackageType() string {
	if o == nil || IsNil(o.PackageType) {
		var ret string
		return ret
	}
	return *o.PackageType
}

// GetPackageTypeOk returns a tuple with the PackageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PhysicalPortSettings) GetPackageTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PackageType) {
		return nil, false
	}
	return o.PackageType, true
}

// HasPackageType returns a boolean if a field has been set.
func (o *PhysicalPortSettings) HasPackageType() bool {
	if o != nil && !IsNil(o.PackageType) {
		return true
	}

	return false
}

// SetPackageType gets a reference to the given string and assigns it to the PackageType field.
// Deprecated
func (o *PhysicalPortSettings) SetPackageType(v string) {
	o.PackageType = &v
}

func (o PhysicalPortSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhysicalPortSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.PackageType) {
		toSerialize["packageType"] = o.PackageType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PhysicalPortSettings) UnmarshalJSON(data []byte) (err error) {
	varPhysicalPortSettings := _PhysicalPortSettings{}

	err = json.Unmarshal(data, &varPhysicalPortSettings)

	if err != nil {
		return err
	}

	*o = PhysicalPortSettings(varPhysicalPortSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "packageType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePhysicalPortSettings struct {
	value *PhysicalPortSettings
	isSet bool
}

func (v NullablePhysicalPortSettings) Get() *PhysicalPortSettings {
	return v.value
}

func (v *NullablePhysicalPortSettings) Set(val *PhysicalPortSettings) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalPortSettings) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalPortSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalPortSettings(val *PhysicalPortSettings) *NullablePhysicalPortSettings {
	return &NullablePhysicalPortSettings{value: val, isSet: true}
}

func (v NullablePhysicalPortSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalPortSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
