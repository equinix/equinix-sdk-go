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
	"time"
)

// checks if the Change type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Change{}

// Change Current state of latest connection change
type Change struct {
	// Uniquely identifies a change
	Uuid   *string       `json:"uuid,omitempty"`
	Type   ChangeType    `json:"type"`
	Status *ChangeStatus `json:"status,omitempty"`
	// Set when change flow starts
	CreatedDateTime time.Time `json:"createdDateTime"`
	// Set when change object is updated
	UpdatedDateTime *time.Time `json:"updatedDateTime,omitempty"`
	// Additional information
	Information          *string                    `json:"information,omitempty"`
	Data                 *ConnectionChangeOperation `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Change Change

// NewChange instantiates a new Change object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChange(type_ ChangeType, createdDateTime time.Time) *Change {
	this := Change{}
	this.Type = type_
	this.CreatedDateTime = createdDateTime
	return &this
}

// NewChangeWithDefaults instantiates a new Change object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeWithDefaults() *Change {
	this := Change{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Change) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Change) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Change) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *Change) GetType() ChangeType {
	if o == nil {
		var ret ChangeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Change) GetTypeOk() (*ChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Change) SetType(v ChangeType) {
	o.Type = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Change) GetStatus() ChangeStatus {
	if o == nil || IsNil(o.Status) {
		var ret ChangeStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetStatusOk() (*ChangeStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Change) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ChangeStatus and assigns it to the Status field.
func (o *Change) SetStatus(v ChangeStatus) {
	o.Status = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value
func (o *Change) GetCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *Change) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedDateTime, true
}

// SetCreatedDateTime sets field value
func (o *Change) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = v
}

// GetUpdatedDateTime returns the UpdatedDateTime field value if set, zero value otherwise.
func (o *Change) GetUpdatedDateTime() time.Time {
	if o == nil || IsNil(o.UpdatedDateTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedDateTime
}

// GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedDateTime) {
		return nil, false
	}
	return o.UpdatedDateTime, true
}

// HasUpdatedDateTime returns a boolean if a field has been set.
func (o *Change) HasUpdatedDateTime() bool {
	if o != nil && !IsNil(o.UpdatedDateTime) {
		return true
	}

	return false
}

// SetUpdatedDateTime gets a reference to the given time.Time and assigns it to the UpdatedDateTime field.
func (o *Change) SetUpdatedDateTime(v time.Time) {
	o.UpdatedDateTime = &v
}

// GetInformation returns the Information field value if set, zero value otherwise.
func (o *Change) GetInformation() string {
	if o == nil || IsNil(o.Information) {
		var ret string
		return ret
	}
	return *o.Information
}

// GetInformationOk returns a tuple with the Information field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetInformationOk() (*string, bool) {
	if o == nil || IsNil(o.Information) {
		return nil, false
	}
	return o.Information, true
}

// HasInformation returns a boolean if a field has been set.
func (o *Change) HasInformation() bool {
	if o != nil && !IsNil(o.Information) {
		return true
	}

	return false
}

// SetInformation gets a reference to the given string and assigns it to the Information field.
func (o *Change) SetInformation(v string) {
	o.Information = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Change) GetData() ConnectionChangeOperation {
	if o == nil || IsNil(o.Data) {
		var ret ConnectionChangeOperation
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Change) GetDataOk() (*ConnectionChangeOperation, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Change) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ConnectionChangeOperation and assigns it to the Data field.
func (o *Change) SetData(v ConnectionChangeOperation) {
	o.Data = &v
}

func (o Change) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Change) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["createdDateTime"] = o.CreatedDateTime
	if !IsNil(o.UpdatedDateTime) {
		toSerialize["updatedDateTime"] = o.UpdatedDateTime
	}
	if !IsNil(o.Information) {
		toSerialize["information"] = o.Information
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Change) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"createdDateTime",
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

	varChange := _Change{}

	err = json.Unmarshal(data, &varChange)

	if err != nil {
		return err
	}

	*o = Change(varChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "type")
		delete(additionalProperties, "status")
		delete(additionalProperties, "createdDateTime")
		delete(additionalProperties, "updatedDateTime")
		delete(additionalProperties, "information")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChange struct {
	value *Change
	isSet bool
}

func (v NullableChange) Get() *Change {
	return v.value
}

func (v *NullableChange) Set(val *Change) {
	v.value = val
	v.isSet = true
}

func (v NullableChange) IsSet() bool {
	return v.isSet
}

func (v *NullableChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChange(val *Change) *NullableChange {
	return &NullableChange{value: val, isSet: true}
}

func (v NullableChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}