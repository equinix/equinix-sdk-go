/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ServiceProfileAccessPointTypeVD type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileAccessPointTypeVD{}

// ServiceProfileAccessPointTypeVD VirtualDevice Access Point Type
type ServiceProfileAccessPointTypeVD struct {
	Uuid                *string                           `json:"uuid,omitempty"`
	Type                ServiceProfileAccessPointTypeEnum `json:"type"`
	SupportedBandwidths []int32                           `json:"supportedBandwidths,omitempty"`
	// Allow remote connections to Service Profile
	AllowRemoteConnections *bool `json:"allowRemoteConnections,omitempty"`
	AllowCustomBandwidth   *bool `json:"allowCustomBandwidth,omitempty"`
}

type _ServiceProfileAccessPointTypeVD ServiceProfileAccessPointTypeVD

// NewServiceProfileAccessPointTypeVD instantiates a new ServiceProfileAccessPointTypeVD object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileAccessPointTypeVD(type_ ServiceProfileAccessPointTypeEnum) *ServiceProfileAccessPointTypeVD {
	this := ServiceProfileAccessPointTypeVD{}
	this.Type = type_
	return &this
}

// NewServiceProfileAccessPointTypeVDWithDefaults instantiates a new ServiceProfileAccessPointTypeVD object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileAccessPointTypeVDWithDefaults() *ServiceProfileAccessPointTypeVD {
	this := ServiceProfileAccessPointTypeVD{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeVD) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeVD) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeVD) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ServiceProfileAccessPointTypeVD) SetUuid(v string) {
	o.Uuid = &v
}

// GetType returns the Type field value
func (o *ServiceProfileAccessPointTypeVD) GetType() ServiceProfileAccessPointTypeEnum {
	if o == nil {
		var ret ServiceProfileAccessPointTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeVD) GetTypeOk() (*ServiceProfileAccessPointTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceProfileAccessPointTypeVD) SetType(v ServiceProfileAccessPointTypeEnum) {
	o.Type = v
}

// GetSupportedBandwidths returns the SupportedBandwidths field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeVD) GetSupportedBandwidths() []int32 {
	if o == nil || IsNil(o.SupportedBandwidths) {
		var ret []int32
		return ret
	}
	return o.SupportedBandwidths
}

// GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeVD) GetSupportedBandwidthsOk() ([]int32, bool) {
	if o == nil || IsNil(o.SupportedBandwidths) {
		return nil, false
	}
	return o.SupportedBandwidths, true
}

// HasSupportedBandwidths returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeVD) HasSupportedBandwidths() bool {
	if o != nil && !IsNil(o.SupportedBandwidths) {
		return true
	}

	return false
}

// SetSupportedBandwidths gets a reference to the given []int32 and assigns it to the SupportedBandwidths field.
func (o *ServiceProfileAccessPointTypeVD) SetSupportedBandwidths(v []int32) {
	o.SupportedBandwidths = v
}

// GetAllowRemoteConnections returns the AllowRemoteConnections field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeVD) GetAllowRemoteConnections() bool {
	if o == nil || IsNil(o.AllowRemoteConnections) {
		var ret bool
		return ret
	}
	return *o.AllowRemoteConnections
}

// GetAllowRemoteConnectionsOk returns a tuple with the AllowRemoteConnections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeVD) GetAllowRemoteConnectionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRemoteConnections) {
		return nil, false
	}
	return o.AllowRemoteConnections, true
}

// HasAllowRemoteConnections returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeVD) HasAllowRemoteConnections() bool {
	if o != nil && !IsNil(o.AllowRemoteConnections) {
		return true
	}

	return false
}

// SetAllowRemoteConnections gets a reference to the given bool and assigns it to the AllowRemoteConnections field.
func (o *ServiceProfileAccessPointTypeVD) SetAllowRemoteConnections(v bool) {
	o.AllowRemoteConnections = &v
}

// GetAllowCustomBandwidth returns the AllowCustomBandwidth field value if set, zero value otherwise.
func (o *ServiceProfileAccessPointTypeVD) GetAllowCustomBandwidth() bool {
	if o == nil || IsNil(o.AllowCustomBandwidth) {
		var ret bool
		return ret
	}
	return *o.AllowCustomBandwidth
}

// GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileAccessPointTypeVD) GetAllowCustomBandwidthOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowCustomBandwidth) {
		return nil, false
	}
	return o.AllowCustomBandwidth, true
}

// HasAllowCustomBandwidth returns a boolean if a field has been set.
func (o *ServiceProfileAccessPointTypeVD) HasAllowCustomBandwidth() bool {
	if o != nil && !IsNil(o.AllowCustomBandwidth) {
		return true
	}

	return false
}

// SetAllowCustomBandwidth gets a reference to the given bool and assigns it to the AllowCustomBandwidth field.
func (o *ServiceProfileAccessPointTypeVD) SetAllowCustomBandwidth(v bool) {
	o.AllowCustomBandwidth = &v
}

func (o ServiceProfileAccessPointTypeVD) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileAccessPointTypeVD) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.SupportedBandwidths) {
		toSerialize["supportedBandwidths"] = o.SupportedBandwidths
	}
	if !IsNil(o.AllowRemoteConnections) {
		toSerialize["allowRemoteConnections"] = o.AllowRemoteConnections
	}
	if !IsNil(o.AllowCustomBandwidth) {
		toSerialize["allowCustomBandwidth"] = o.AllowCustomBandwidth
	}
	return toSerialize, nil
}

func (o *ServiceProfileAccessPointTypeVD) UnmarshalJSON(data []byte) (err error) {
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

	varServiceProfileAccessPointTypeVD := _ServiceProfileAccessPointTypeVD{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceProfileAccessPointTypeVD)

	if err != nil {
		return err
	}

	*o = ServiceProfileAccessPointTypeVD(varServiceProfileAccessPointTypeVD)

	return err
}

type NullableServiceProfileAccessPointTypeVD struct {
	value *ServiceProfileAccessPointTypeVD
	isSet bool
}

func (v NullableServiceProfileAccessPointTypeVD) Get() *ServiceProfileAccessPointTypeVD {
	return v.value
}

func (v *NullableServiceProfileAccessPointTypeVD) Set(val *ServiceProfileAccessPointTypeVD) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileAccessPointTypeVD) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileAccessPointTypeVD) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileAccessPointTypeVD(val *ServiceProfileAccessPointTypeVD) *NullableServiceProfileAccessPointTypeVD {
	return &NullableServiceProfileAccessPointTypeVD{value: val, isSet: true}
}

func (v NullableServiceProfileAccessPointTypeVD) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileAccessPointTypeVD) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
