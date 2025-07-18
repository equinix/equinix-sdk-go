/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the FabricBGPConnectionIpv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FabricBGPConnectionIpv4{}

// FabricBGPConnectionIpv4 Defines the structure for a BGP IPv4 connection, including customer and Equinix peering IP addresses.
type FabricBGPConnectionIpv4 struct {
	// Customer side peering ip
	CustomerIp *string `json:"customerIp,omitempty"`
	// Equinix side peering ip
	EquinixIp            string `json:"equinixIp"`
	AdditionalProperties map[string]interface{}
}

type _FabricBGPConnectionIpv4 FabricBGPConnectionIpv4

// NewFabricBGPConnectionIpv4 instantiates a new FabricBGPConnectionIpv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFabricBGPConnectionIpv4(equinixIp string) *FabricBGPConnectionIpv4 {
	this := FabricBGPConnectionIpv4{}
	this.EquinixIp = equinixIp
	return &this
}

// NewFabricBGPConnectionIpv4WithDefaults instantiates a new FabricBGPConnectionIpv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFabricBGPConnectionIpv4WithDefaults() *FabricBGPConnectionIpv4 {
	this := FabricBGPConnectionIpv4{}
	return &this
}

// GetCustomerIp returns the CustomerIp field value if set, zero value otherwise.
func (o *FabricBGPConnectionIpv4) GetCustomerIp() string {
	if o == nil || IsNil(o.CustomerIp) {
		var ret string
		return ret
	}
	return *o.CustomerIp
}

// GetCustomerIpOk returns a tuple with the CustomerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FabricBGPConnectionIpv4) GetCustomerIpOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerIp) {
		return nil, false
	}
	return o.CustomerIp, true
}

// HasCustomerIp returns a boolean if a field has been set.
func (o *FabricBGPConnectionIpv4) HasCustomerIp() bool {
	if o != nil && !IsNil(o.CustomerIp) {
		return true
	}

	return false
}

// SetCustomerIp gets a reference to the given string and assigns it to the CustomerIp field.
func (o *FabricBGPConnectionIpv4) SetCustomerIp(v string) {
	o.CustomerIp = &v
}

// GetEquinixIp returns the EquinixIp field value
func (o *FabricBGPConnectionIpv4) GetEquinixIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EquinixIp
}

// GetEquinixIpOk returns a tuple with the EquinixIp field value
// and a boolean to check if the value has been set.
func (o *FabricBGPConnectionIpv4) GetEquinixIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EquinixIp, true
}

// SetEquinixIp sets field value
func (o *FabricBGPConnectionIpv4) SetEquinixIp(v string) {
	o.EquinixIp = v
}

func (o FabricBGPConnectionIpv4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FabricBGPConnectionIpv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomerIp) {
		toSerialize["customerIp"] = o.CustomerIp
	}
	toSerialize["equinixIp"] = o.EquinixIp

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FabricBGPConnectionIpv4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"equinixIp",
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

	varFabricBGPConnectionIpv4 := _FabricBGPConnectionIpv4{}

	err = json.Unmarshal(data, &varFabricBGPConnectionIpv4)

	if err != nil {
		return err
	}

	*o = FabricBGPConnectionIpv4(varFabricBGPConnectionIpv4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerIp")
		delete(additionalProperties, "equinixIp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFabricBGPConnectionIpv4 struct {
	value *FabricBGPConnectionIpv4
	isSet bool
}

func (v NullableFabricBGPConnectionIpv4) Get() *FabricBGPConnectionIpv4 {
	return v.value
}

func (v *NullableFabricBGPConnectionIpv4) Set(val *FabricBGPConnectionIpv4) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricBGPConnectionIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricBGPConnectionIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricBGPConnectionIpv4(val *FabricBGPConnectionIpv4) *NullableFabricBGPConnectionIpv4 {
	return &NullableFabricBGPConnectionIpv4{value: val, isSet: true}
}

func (v NullableFabricBGPConnectionIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricBGPConnectionIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
