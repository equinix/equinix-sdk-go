/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the Ipv4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv4{}

// Ipv4 EPT service network information
type Ipv4 struct {
	// Primary Timing Server IP Address
	Primary string `json:"primary"`
	// Secondary Timing Server IP Address
	Secondary string `json:"secondary"`
	// Network Mask
	NetworkMask string `json:"networkMask"`
	// Gateway Interface IP address
	DefaultGateway       *string `json:"defaultGateway,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Ipv4 Ipv4

// NewIpv4 instantiates a new Ipv4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv4(primary string, secondary string, networkMask string) *Ipv4 {
	this := Ipv4{}
	this.Primary = primary
	this.Secondary = secondary
	this.NetworkMask = networkMask
	return &this
}

// NewIpv4WithDefaults instantiates a new Ipv4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv4WithDefaults() *Ipv4 {
	this := Ipv4{}
	return &this
}

// GetPrimary returns the Primary field value
func (o *Ipv4) GetPrimary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value
// and a boolean to check if the value has been set.
func (o *Ipv4) GetPrimaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Primary, true
}

// SetPrimary sets field value
func (o *Ipv4) SetPrimary(v string) {
	o.Primary = v
}

// GetSecondary returns the Secondary field value
func (o *Ipv4) GetSecondary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secondary
}

// GetSecondaryOk returns a tuple with the Secondary field value
// and a boolean to check if the value has been set.
func (o *Ipv4) GetSecondaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secondary, true
}

// SetSecondary sets field value
func (o *Ipv4) SetSecondary(v string) {
	o.Secondary = v
}

// GetNetworkMask returns the NetworkMask field value
func (o *Ipv4) GetNetworkMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkMask
}

// GetNetworkMaskOk returns a tuple with the NetworkMask field value
// and a boolean to check if the value has been set.
func (o *Ipv4) GetNetworkMaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkMask, true
}

// SetNetworkMask sets field value
func (o *Ipv4) SetNetworkMask(v string) {
	o.NetworkMask = v
}

// GetDefaultGateway returns the DefaultGateway field value if set, zero value otherwise.
func (o *Ipv4) GetDefaultGateway() string {
	if o == nil || IsNil(o.DefaultGateway) {
		var ret string
		return ret
	}
	return *o.DefaultGateway
}

// GetDefaultGatewayOk returns a tuple with the DefaultGateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ipv4) GetDefaultGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultGateway) {
		return nil, false
	}
	return o.DefaultGateway, true
}

// HasDefaultGateway returns a boolean if a field has been set.
func (o *Ipv4) HasDefaultGateway() bool {
	if o != nil && !IsNil(o.DefaultGateway) {
		return true
	}

	return false
}

// SetDefaultGateway gets a reference to the given string and assigns it to the DefaultGateway field.
func (o *Ipv4) SetDefaultGateway(v string) {
	o.DefaultGateway = &v
}

func (o Ipv4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["primary"] = o.Primary
	toSerialize["secondary"] = o.Secondary
	toSerialize["networkMask"] = o.NetworkMask
	if !IsNil(o.DefaultGateway) {
		toSerialize["defaultGateway"] = o.DefaultGateway
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Ipv4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"primary",
		"secondary",
		"networkMask",
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

	varIpv4 := _Ipv4{}

	err = json.Unmarshal(data, &varIpv4)

	if err != nil {
		return err
	}

	*o = Ipv4(varIpv4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "primary")
		delete(additionalProperties, "secondary")
		delete(additionalProperties, "networkMask")
		delete(additionalProperties, "defaultGateway")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIpv4 struct {
	value *Ipv4
	isSet bool
}

func (v NullableIpv4) Get() *Ipv4 {
	return v.value
}

func (v *NullableIpv4) Set(val *Ipv4) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv4) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv4(val *Ipv4) *NullableIpv4 {
	return &NullableIpv4{value: val, isSet: true}
}

func (v NullableIpv4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
