/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the CustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomField{}

// CustomField Define Custom Attributes
type CustomField struct {
	Label       string              `json:"label"`
	Description *string             `json:"description,omitempty"`
	Required    *bool               `json:"required,omitempty"`
	DataType    CustomFieldDataType `json:"dataType"`
	Options     []string            `json:"options,omitempty"`
	// capture this field as a part of email notification
	CaptureInEmail       *bool `json:"captureInEmail,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomField CustomField

// NewCustomField instantiates a new CustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomField(label string, dataType CustomFieldDataType) *CustomField {
	this := CustomField{}
	this.Label = label
	this.DataType = dataType
	return &this
}

// NewCustomFieldWithDefaults instantiates a new CustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFieldWithDefaults() *CustomField {
	this := CustomField{}
	return &this
}

// GetLabel returns the Label field value
func (o *CustomField) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *CustomField) SetLabel(v string) {
	o.Label = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CustomField) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CustomField) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CustomField) SetDescription(v string) {
	o.Description = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *CustomField) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *CustomField) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *CustomField) SetRequired(v bool) {
	o.Required = &v
}

// GetDataType returns the DataType field value
func (o *CustomField) GetDataType() CustomFieldDataType {
	if o == nil {
		var ret CustomFieldDataType
		return ret
	}

	return o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value
// and a boolean to check if the value has been set.
func (o *CustomField) GetDataTypeOk() (*CustomFieldDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataType, true
}

// SetDataType sets field value
func (o *CustomField) SetDataType(v CustomFieldDataType) {
	o.DataType = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *CustomField) GetOptions() []string {
	if o == nil || IsNil(o.Options) {
		var ret []string
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetOptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *CustomField) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []string and assigns it to the Options field.
func (o *CustomField) SetOptions(v []string) {
	o.Options = v
}

// GetCaptureInEmail returns the CaptureInEmail field value if set, zero value otherwise.
func (o *CustomField) GetCaptureInEmail() bool {
	if o == nil || IsNil(o.CaptureInEmail) {
		var ret bool
		return ret
	}
	return *o.CaptureInEmail
}

// GetCaptureInEmailOk returns a tuple with the CaptureInEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomField) GetCaptureInEmailOk() (*bool, bool) {
	if o == nil || IsNil(o.CaptureInEmail) {
		return nil, false
	}
	return o.CaptureInEmail, true
}

// HasCaptureInEmail returns a boolean if a field has been set.
func (o *CustomField) HasCaptureInEmail() bool {
	if o != nil && !IsNil(o.CaptureInEmail) {
		return true
	}

	return false
}

// SetCaptureInEmail gets a reference to the given bool and assigns it to the CaptureInEmail field.
func (o *CustomField) SetCaptureInEmail(v bool) {
	o.CaptureInEmail = &v
}

func (o CustomField) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	toSerialize["dataType"] = o.DataType
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.CaptureInEmail) {
		toSerialize["captureInEmail"] = o.CaptureInEmail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomField) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"label",
		"dataType",
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

	varCustomField := _CustomField{}

	err = json.Unmarshal(data, &varCustomField)

	if err != nil {
		return err
	}

	*o = CustomField(varCustomField)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		delete(additionalProperties, "required")
		delete(additionalProperties, "dataType")
		delete(additionalProperties, "options")
		delete(additionalProperties, "captureInEmail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomField struct {
	value *CustomField
	isSet bool
}

func (v NullableCustomField) Get() *CustomField {
	return v.value
}

func (v *NullableCustomField) Set(val *CustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomField(val *CustomField) *NullableCustomField {
	return &NullableCustomField{value: val, isSet: true}
}

func (v NullableCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
