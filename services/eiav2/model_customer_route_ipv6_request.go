/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
)

// checks if the CustomerRouteIpv6Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerRouteIpv6Request{}

// CustomerRouteIpv6Request struct for CustomerRouteIpv6Request
type CustomerRouteIpv6Request struct {
	IpBlock *IpBlockIpv6Request `json:"ipBlock,omitempty"`
	// Subnet prefix
	Prefix               *string `json:"prefix,omitempty" validate:"regexp=^s*((([0-9A-Fa-f]{1,4}:){7}([0-9A-Fa-f]{1,4}|:))|(([0-9A-Fa-f]{1,4}:){6}(:[0-9A-Fa-f]{1,4}|((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){5}(((:[0-9A-Fa-f]{1,4}){1,2})|:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3})|:))|(([0-9A-Fa-f]{1,4}:){4}(((:[0-9A-Fa-f]{1,4}){1,3})|((:[0-9A-Fa-f]{1,4})?:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){3}(((:[0-9A-Fa-f]{1,4}){1,4})|((:[0-9A-Fa-f]{1,4}){0,2}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){2}(((:[0-9A-Fa-f]{1,4}){1,5})|((:[0-9A-Fa-f]{1,4}){0,3}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(([0-9A-Fa-f]{1,4}:){1}(((:[0-9A-Fa-f]{1,4}){1,6})|((:[0-9A-Fa-f]{1,4}){0,4}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:))|(:(((:[0-9A-Fa-f]{1,4}){1,7})|((:[0-9A-Fa-f]{1,4}){0,5}:((25[0-5]|2[0-4]d|1dd|[1-9]?d)(.(25[0-5]|2[0-4]d|1dd|[1-9]?d)){3}))|:)))(%.+)?s*(\\/([0-9]|[1-9][0-9]|1[0-1][0-9]|12[0-8]))?$"`
	AdditionalProperties map[string]interface{}
}

type _CustomerRouteIpv6Request CustomerRouteIpv6Request

// NewCustomerRouteIpv6Request instantiates a new CustomerRouteIpv6Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerRouteIpv6Request() *CustomerRouteIpv6Request {
	this := CustomerRouteIpv6Request{}
	return &this
}

// NewCustomerRouteIpv6RequestWithDefaults instantiates a new CustomerRouteIpv6Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerRouteIpv6RequestWithDefaults() *CustomerRouteIpv6Request {
	this := CustomerRouteIpv6Request{}
	return &this
}

// GetIpBlock returns the IpBlock field value if set, zero value otherwise.
func (o *CustomerRouteIpv6Request) GetIpBlock() IpBlockIpv6Request {
	if o == nil || IsNil(o.IpBlock) {
		var ret IpBlockIpv6Request
		return ret
	}
	return *o.IpBlock
}

// GetIpBlockOk returns a tuple with the IpBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRouteIpv6Request) GetIpBlockOk() (*IpBlockIpv6Request, bool) {
	if o == nil || IsNil(o.IpBlock) {
		return nil, false
	}
	return o.IpBlock, true
}

// HasIpBlock returns a boolean if a field has been set.
func (o *CustomerRouteIpv6Request) HasIpBlock() bool {
	if o != nil && !IsNil(o.IpBlock) {
		return true
	}

	return false
}

// SetIpBlock gets a reference to the given IpBlockIpv6Request and assigns it to the IpBlock field.
func (o *CustomerRouteIpv6Request) SetIpBlock(v IpBlockIpv6Request) {
	o.IpBlock = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *CustomerRouteIpv6Request) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerRouteIpv6Request) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *CustomerRouteIpv6Request) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *CustomerRouteIpv6Request) SetPrefix(v string) {
	o.Prefix = &v
}

func (o CustomerRouteIpv6Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerRouteIpv6Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpBlock) {
		toSerialize["ipBlock"] = o.IpBlock
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomerRouteIpv6Request) UnmarshalJSON(data []byte) (err error) {
	varCustomerRouteIpv6Request := _CustomerRouteIpv6Request{}

	err = json.Unmarshal(data, &varCustomerRouteIpv6Request)

	if err != nil {
		return err
	}

	*o = CustomerRouteIpv6Request(varCustomerRouteIpv6Request)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ipBlock")
		delete(additionalProperties, "prefix")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomerRouteIpv6Request struct {
	value *CustomerRouteIpv6Request
	isSet bool
}

func (v NullableCustomerRouteIpv6Request) Get() *CustomerRouteIpv6Request {
	return v.value
}

func (v *NullableCustomerRouteIpv6Request) Set(val *CustomerRouteIpv6Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerRouteIpv6Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerRouteIpv6Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerRouteIpv6Request(val *CustomerRouteIpv6Request) *NullableCustomerRouteIpv6Request {
	return &NullableCustomerRouteIpv6Request{value: val, isSet: true}
}

func (v NullableCustomerRouteIpv6Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerRouteIpv6Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
