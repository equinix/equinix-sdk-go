/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the BulkPhysicalPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkPhysicalPort{}

// BulkPhysicalPort Add to Lag request
type BulkPhysicalPort struct {
	// add physical ports to virtual port
	Data                 []PhysicalPort `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BulkPhysicalPort BulkPhysicalPort

// NewBulkPhysicalPort instantiates a new BulkPhysicalPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkPhysicalPort() *BulkPhysicalPort {
	this := BulkPhysicalPort{}
	return &this
}

// NewBulkPhysicalPortWithDefaults instantiates a new BulkPhysicalPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkPhysicalPortWithDefaults() *BulkPhysicalPort {
	this := BulkPhysicalPort{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BulkPhysicalPort) GetData() []PhysicalPort {
	if o == nil || IsNil(o.Data) {
		var ret []PhysicalPort
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkPhysicalPort) GetDataOk() ([]PhysicalPort, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BulkPhysicalPort) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PhysicalPort and assigns it to the Data field.
func (o *BulkPhysicalPort) SetData(v []PhysicalPort) {
	o.Data = v
}

func (o BulkPhysicalPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkPhysicalPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BulkPhysicalPort) UnmarshalJSON(data []byte) (err error) {
	varBulkPhysicalPort := _BulkPhysicalPort{}

	err = json.Unmarshal(data, &varBulkPhysicalPort)

	if err != nil {
		return err
	}

	*o = BulkPhysicalPort(varBulkPhysicalPort)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBulkPhysicalPort struct {
	value *BulkPhysicalPort
	isSet bool
}

func (v NullableBulkPhysicalPort) Get() *BulkPhysicalPort {
	return v.value
}

func (v *NullableBulkPhysicalPort) Set(val *BulkPhysicalPort) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkPhysicalPort) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkPhysicalPort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkPhysicalPort(val *BulkPhysicalPort) *NullableBulkPhysicalPort {
	return &NullableBulkPhysicalPort{value: val, isSet: true}
}

func (v NullableBulkPhysicalPort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkPhysicalPort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
