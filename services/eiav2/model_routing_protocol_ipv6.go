/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolIpv6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolIpv6{}

// RoutingProtocolIpv6 struct for RoutingProtocolIpv6
type RoutingProtocolIpv6 struct {
	CustomerRoutes       []RoutingProtocolCustomerRouteIpv6 `json:"customerRoutes"`
	Peerings             []RoutingProtocolPeeringIpv6       `json:"peerings"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolIpv6 RoutingProtocolIpv6

// NewRoutingProtocolIpv6 instantiates a new RoutingProtocolIpv6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolIpv6(customerRoutes []RoutingProtocolCustomerRouteIpv6, peerings []RoutingProtocolPeeringIpv6) *RoutingProtocolIpv6 {
	this := RoutingProtocolIpv6{}
	this.CustomerRoutes = customerRoutes
	this.Peerings = peerings
	return &this
}

// NewRoutingProtocolIpv6WithDefaults instantiates a new RoutingProtocolIpv6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolIpv6WithDefaults() *RoutingProtocolIpv6 {
	this := RoutingProtocolIpv6{}
	return &this
}

// GetCustomerRoutes returns the CustomerRoutes field value
func (o *RoutingProtocolIpv6) GetCustomerRoutes() []RoutingProtocolCustomerRouteIpv6 {
	if o == nil {
		var ret []RoutingProtocolCustomerRouteIpv6
		return ret
	}

	return o.CustomerRoutes
}

// GetCustomerRoutesOk returns a tuple with the CustomerRoutes field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolIpv6) GetCustomerRoutesOk() ([]RoutingProtocolCustomerRouteIpv6, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerRoutes, true
}

// SetCustomerRoutes sets field value
func (o *RoutingProtocolIpv6) SetCustomerRoutes(v []RoutingProtocolCustomerRouteIpv6) {
	o.CustomerRoutes = v
}

// GetPeerings returns the Peerings field value
func (o *RoutingProtocolIpv6) GetPeerings() []RoutingProtocolPeeringIpv6 {
	if o == nil {
		var ret []RoutingProtocolPeeringIpv6
		return ret
	}

	return o.Peerings
}

// GetPeeringsOk returns a tuple with the Peerings field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolIpv6) GetPeeringsOk() ([]RoutingProtocolPeeringIpv6, bool) {
	if o == nil {
		return nil, false
	}
	return o.Peerings, true
}

// SetPeerings sets field value
func (o *RoutingProtocolIpv6) SetPeerings(v []RoutingProtocolPeeringIpv6) {
	o.Peerings = v
}

func (o RoutingProtocolIpv6) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolIpv6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customerRoutes"] = o.CustomerRoutes
	toSerialize["peerings"] = o.Peerings

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolIpv6) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"customerRoutes",
		"peerings",
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

	varRoutingProtocolIpv6 := _RoutingProtocolIpv6{}

	err = json.Unmarshal(data, &varRoutingProtocolIpv6)

	if err != nil {
		return err
	}

	*o = RoutingProtocolIpv6(varRoutingProtocolIpv6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "customerRoutes")
		delete(additionalProperties, "peerings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolIpv6 struct {
	value *RoutingProtocolIpv6
	isSet bool
}

func (v NullableRoutingProtocolIpv6) Get() *RoutingProtocolIpv6 {
	return v.value
}

func (v *NullableRoutingProtocolIpv6) Set(val *RoutingProtocolIpv6) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolIpv6) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolIpv6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolIpv6(val *RoutingProtocolIpv6) *NullableRoutingProtocolIpv6 {
	return &NullableRoutingProtocolIpv6{value: val, isSet: true}
}

func (v NullableRoutingProtocolIpv6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolIpv6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
