/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the ReplaceOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceOperation{}

// ReplaceOperation Replace attribute value or sub-resource in the existing model
type ReplaceOperation struct {
	Op OpEnum `json:"op"`
	// A JSON Pointer path.
	Path string `json:"path"`
	// value to replace with
	Value                map[string]interface{} `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _ReplaceOperation ReplaceOperation

// NewReplaceOperation instantiates a new ReplaceOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceOperation(op OpEnum, path string, value map[string]interface{}) *ReplaceOperation {
	this := ReplaceOperation{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewReplaceOperationWithDefaults instantiates a new ReplaceOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceOperationWithDefaults() *ReplaceOperation {
	this := ReplaceOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *ReplaceOperation) GetOp() OpEnum {
	if o == nil {
		var ret OpEnum
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *ReplaceOperation) GetOpOk() (*OpEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *ReplaceOperation) SetOp(v OpEnum) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *ReplaceOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ReplaceOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ReplaceOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
func (o *ReplaceOperation) GetValue() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ReplaceOperation) GetValueOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *ReplaceOperation) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o ReplaceOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ReplaceOperation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"op",
		"path",
		"value",
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

	varReplaceOperation := _ReplaceOperation{}

	err = json.Unmarshal(data, &varReplaceOperation)

	if err != nil {
		return err
	}

	*o = ReplaceOperation(varReplaceOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReplaceOperation struct {
	value *ReplaceOperation
	isSet bool
}

func (v NullableReplaceOperation) Get() *ReplaceOperation {
	return v.value
}

func (v *NullableReplaceOperation) Set(val *ReplaceOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceOperation(val *ReplaceOperation) *NullableReplaceOperation {
	return &NullableReplaceOperation{value: val, isSet: true}
}

func (v NullableReplaceOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}