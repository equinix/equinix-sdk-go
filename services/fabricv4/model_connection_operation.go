/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"time"
)

// checks if the ConnectionOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionOperation{}

// ConnectionOperation Connection type-specific operational data
type ConnectionOperation struct {
	ProviderStatus    *ProviderStatus                       `json:"providerStatus,omitempty"`
	EquinixStatus     *EquinixStatus                        `json:"equinixStatus,omitempty"`
	OperationalStatus *ConnectionOperationOperationalStatus `json:"operationalStatus,omitempty"`
	Errors            []Error                               `json:"errors,omitempty"`
	// When connection transitioned into current operational status
	OpStatusChangedAt    *time.Time `json:"opStatusChangedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionOperation ConnectionOperation

// NewConnectionOperation instantiates a new ConnectionOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionOperation() *ConnectionOperation {
	this := ConnectionOperation{}
	return &this
}

// NewConnectionOperationWithDefaults instantiates a new ConnectionOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionOperationWithDefaults() *ConnectionOperation {
	this := ConnectionOperation{}
	return &this
}

// GetProviderStatus returns the ProviderStatus field value if set, zero value otherwise.
func (o *ConnectionOperation) GetProviderStatus() ProviderStatus {
	if o == nil || IsNil(o.ProviderStatus) {
		var ret ProviderStatus
		return ret
	}
	return *o.ProviderStatus
}

// GetProviderStatusOk returns a tuple with the ProviderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOperation) GetProviderStatusOk() (*ProviderStatus, bool) {
	if o == nil || IsNil(o.ProviderStatus) {
		return nil, false
	}
	return o.ProviderStatus, true
}

// HasProviderStatus returns a boolean if a field has been set.
func (o *ConnectionOperation) HasProviderStatus() bool {
	if o != nil && !IsNil(o.ProviderStatus) {
		return true
	}

	return false
}

// SetProviderStatus gets a reference to the given ProviderStatus and assigns it to the ProviderStatus field.
func (o *ConnectionOperation) SetProviderStatus(v ProviderStatus) {
	o.ProviderStatus = &v
}

// GetEquinixStatus returns the EquinixStatus field value if set, zero value otherwise.
func (o *ConnectionOperation) GetEquinixStatus() EquinixStatus {
	if o == nil || IsNil(o.EquinixStatus) {
		var ret EquinixStatus
		return ret
	}
	return *o.EquinixStatus
}

// GetEquinixStatusOk returns a tuple with the EquinixStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOperation) GetEquinixStatusOk() (*EquinixStatus, bool) {
	if o == nil || IsNil(o.EquinixStatus) {
		return nil, false
	}
	return o.EquinixStatus, true
}

// HasEquinixStatus returns a boolean if a field has been set.
func (o *ConnectionOperation) HasEquinixStatus() bool {
	if o != nil && !IsNil(o.EquinixStatus) {
		return true
	}

	return false
}

// SetEquinixStatus gets a reference to the given EquinixStatus and assigns it to the EquinixStatus field.
func (o *ConnectionOperation) SetEquinixStatus(v EquinixStatus) {
	o.EquinixStatus = &v
}

// GetOperationalStatus returns the OperationalStatus field value if set, zero value otherwise.
func (o *ConnectionOperation) GetOperationalStatus() ConnectionOperationOperationalStatus {
	if o == nil || IsNil(o.OperationalStatus) {
		var ret ConnectionOperationOperationalStatus
		return ret
	}
	return *o.OperationalStatus
}

// GetOperationalStatusOk returns a tuple with the OperationalStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOperation) GetOperationalStatusOk() (*ConnectionOperationOperationalStatus, bool) {
	if o == nil || IsNil(o.OperationalStatus) {
		return nil, false
	}
	return o.OperationalStatus, true
}

// HasOperationalStatus returns a boolean if a field has been set.
func (o *ConnectionOperation) HasOperationalStatus() bool {
	if o != nil && !IsNil(o.OperationalStatus) {
		return true
	}

	return false
}

// SetOperationalStatus gets a reference to the given ConnectionOperationOperationalStatus and assigns it to the OperationalStatus field.
func (o *ConnectionOperation) SetOperationalStatus(v ConnectionOperationOperationalStatus) {
	o.OperationalStatus = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ConnectionOperation) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOperation) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ConnectionOperation) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ConnectionOperation) SetErrors(v []Error) {
	o.Errors = v
}

// GetOpStatusChangedAt returns the OpStatusChangedAt field value if set, zero value otherwise.
func (o *ConnectionOperation) GetOpStatusChangedAt() time.Time {
	if o == nil || IsNil(o.OpStatusChangedAt) {
		var ret time.Time
		return ret
	}
	return *o.OpStatusChangedAt
}

// GetOpStatusChangedAtOk returns a tuple with the OpStatusChangedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionOperation) GetOpStatusChangedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OpStatusChangedAt) {
		return nil, false
	}
	return o.OpStatusChangedAt, true
}

// HasOpStatusChangedAt returns a boolean if a field has been set.
func (o *ConnectionOperation) HasOpStatusChangedAt() bool {
	if o != nil && !IsNil(o.OpStatusChangedAt) {
		return true
	}

	return false
}

// SetOpStatusChangedAt gets a reference to the given time.Time and assigns it to the OpStatusChangedAt field.
func (o *ConnectionOperation) SetOpStatusChangedAt(v time.Time) {
	o.OpStatusChangedAt = &v
}

func (o ConnectionOperation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProviderStatus) {
		toSerialize["providerStatus"] = o.ProviderStatus
	}
	if !IsNil(o.EquinixStatus) {
		toSerialize["equinixStatus"] = o.EquinixStatus
	}
	if !IsNil(o.OperationalStatus) {
		toSerialize["operationalStatus"] = o.OperationalStatus
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.OpStatusChangedAt) {
		toSerialize["opStatusChangedAt"] = o.OpStatusChangedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionOperation) UnmarshalJSON(data []byte) (err error) {
	varConnectionOperation := _ConnectionOperation{}

	err = json.Unmarshal(data, &varConnectionOperation)

	if err != nil {
		return err
	}

	*o = ConnectionOperation(varConnectionOperation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "providerStatus")
		delete(additionalProperties, "equinixStatus")
		delete(additionalProperties, "operationalStatus")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "opStatusChangedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionOperation struct {
	value *ConnectionOperation
	isSet bool
}

func (v NullableConnectionOperation) Get() *ConnectionOperation {
	return v.value
}

func (v *NullableConnectionOperation) Set(val *ConnectionOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionOperation(val *ConnectionOperation) *NullableConnectionOperation {
	return &NullableConnectionOperation{value: val, isSet: true}
}

func (v NullableConnectionOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
