/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the CloudRouterChangeOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterChangeOperation{}

// CloudRouterChangeOperation Fabric Cloud Router change operation data
type CloudRouterChangeOperation struct {
	Op PrecisionTimeChangeOperationOp `json:"op"`
	// path inside document leading to updated parameter
	Path string `json:"path"`
	// new value for updated parameter
	Value                interface{} `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterChangeOperation CloudRouterChangeOperation

// NewCloudRouterChangeOperation instantiates a new CloudRouterChangeOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterChangeOperation(op PrecisionTimeChangeOperationOp, path string, value interface{}) *CloudRouterChangeOperation {
	this := CloudRouterChangeOperation{}
	this.Op = op
	this.Path = path
	this.Value = value
	return &this
}

// NewCloudRouterChangeOperationWithDefaults instantiates a new CloudRouterChangeOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterChangeOperationWithDefaults() *CloudRouterChangeOperation {
	this := CloudRouterChangeOperation{}
	return &this
}

// GetOp returns the Op field value
func (o *CloudRouterChangeOperation) GetOp() PrecisionTimeChangeOperationOp {
	if o == nil {
		var ret PrecisionTimeChangeOperationOp
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *CloudRouterChangeOperation) GetOpOk() (*PrecisionTimeChangeOperationOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *CloudRouterChangeOperation) SetOp(v PrecisionTimeChangeOperationOp) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *CloudRouterChangeOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CloudRouterChangeOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CloudRouterChangeOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CloudRouterChangeOperation) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudRouterChangeOperation) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CloudRouterChangeOperation) SetValue(v interface{}) {
	o.Value = v
}

func (o CloudRouterChangeOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterChangeOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterChangeOperation) UnmarshalJSON(data []byte) (err error) {
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

	varCloudRouterChangeOperation := _CloudRouterChangeOperation{}

	err = json.Unmarshal(data, &varCloudRouterChangeOperation)

	if err != nil {
		return err
	}

	*o = CloudRouterChangeOperation(varCloudRouterChangeOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "op")
		delete(additionalProperties, "path")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterChangeOperation struct {
	value *CloudRouterChangeOperation
	isSet bool
}

func (v NullableCloudRouterChangeOperation) Get() *CloudRouterChangeOperation {
	return v.value
}

func (v *NullableCloudRouterChangeOperation) Set(val *CloudRouterChangeOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterChangeOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterChangeOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterChangeOperation(val *CloudRouterChangeOperation) *NullableCloudRouterChangeOperation {
	return &NullableCloudRouterChangeOperation{value: val, isSet: true}
}

func (v NullableCloudRouterChangeOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterChangeOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
