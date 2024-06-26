/*
Equinix Fabric API v4

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"time"
)

// checks if the PortOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortOperation{}

// PortOperation Operational specifications for ports.
type PortOperation struct {
	OperationalStatus *PortOperationOperationalStatus `json:"operationalStatus,omitempty"`
	// Total number of connections.
	ConnectionCount *int32 `json:"connectionCount,omitempty"`
	// Date and time at which port availability changed.
	OpStatusChangedAt    *time.Time `json:"opStatusChangedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PortOperation PortOperation

// NewPortOperation instantiates a new PortOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortOperation() *PortOperation {
	this := PortOperation{}
	return &this
}

// NewPortOperationWithDefaults instantiates a new PortOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortOperationWithDefaults() *PortOperation {
	this := PortOperation{}
	return &this
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *PortOperation) GetOperationalStatus() PortOperationOperationalStatus {
	if o == nil || IsNil(o.OperationalStatus) {
		var ret PortOperationOperationalStatus
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOperation) GetOperationalStatusOk() (*PortOperationOperationalStatus, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *PortOperation) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given PortOperationOperationalStatus and assigns it to the OperationalStatus field.
func (o *PortOperation) SetOperationalStatus(v PortOperationOperationalStatus) {
	o.OperationalStatus = &v
}

// GetConnectionCount returns the ConnectionCount field value if set, zero value otherwise.
func (o *PortOperation) GetConnectionCount() int32 {
	if o == nil || IsNil(o.ConnectionCount) {
		var ret int32
		return ret
	}
	return *o.ConnectionCount
}

// GetConnectionCountOk returns a tuple with the ConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOperation) GetConnectionCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ConnectionCount) {
		return nil, false
	}
	return o.ConnectionCount, true
}

// HasConnectionCount returns a boolean if a field has been set.
func (o *PortOperation) HasConnectionCount() bool {
	if o != nil && !IsNil(o.ConnectionCount) {
		return true
	}

	return false
}

// SetConnectionCount gets a reference to the given int32 and assigns it to the ConnectionCount field.
func (o *PortOperation) SetConnectionCount(v int32) {
	o.ConnectionCount = &v
}

// GetOpStatusChangedAt returns the OpStatusChangedAt field value if set, zero value otherwise.
func (o *PortOperation) GetOpStatusChangedAt() time.Time {
	if o == nil || IsNil(o.OpStatusChangedAt) {
		var ret time.Time
		return ret
	}
	return *o.OpStatusChangedAt
}

// GetOpStatusChangedAtOk returns a tuple with the OpStatusChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortOperation) GetOpStatusChangedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpStatusChangedAt) {
		return nil, false
	}
	return o.OpStatusChangedAt, true
}

// HasOpStatusChangedAt returns a boolean if a field has been set.
func (o *PortOperation) HasOpStatusChangedAt() bool {
	if o != nil && !IsNil(o.OpStatusChangedAt) {
		return true
	}

	return false
}

// SetOpStatusChangedAt gets a reference to the given time.Time and assigns it to the OpStatusChangedAt field.
func (o *PortOperation) SetOpStatusChangedAt(v time.Time) {
	o.OpStatusChangedAt = &v
}

func (o PortOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OperationalStatus) {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.ConnectionCount) {
		toSerialize["connectionCount"] = o.ConnectionCount
	}
	if !IsNil(o.OpStatusChangedAt) {
		toSerialize["opStatusChangedAt"] = o.OpStatusChangedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PortOperation) UnmarshalJSON(data []byte) (err error) {
	varPortOperation := _PortOperation{}

	err = json.Unmarshal(data, &varPortOperation)

	if err != nil {
		return err
	}

	*o = PortOperation(varPortOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "operationalStatus")
		delete(additionalProperties, "connectionCount")
		delete(additionalProperties, "opStatusChangedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePortOperation struct {
	value *PortOperation
	isSet bool
}

func (v NullablePortOperation) Get() *PortOperation {
	return v.value
}

func (v *NullablePortOperation) Set(val *PortOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePortOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePortOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortOperation(val *PortOperation) *NullablePortOperation {
	return &NullablePortOperation{value: val, isSet: true}
}

func (v NullablePortOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
