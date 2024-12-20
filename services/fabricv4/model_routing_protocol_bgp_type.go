/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the RoutingProtocolBGPType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingProtocolBGPType{}

// RoutingProtocolBGPType struct for RoutingProtocolBGPType
type RoutingProtocolBGPType struct {
	Type    RoutingProtocolBGPTypeType `json:"type"`
	Name    *string                    `json:"name,omitempty"`
	BgpIpv4 *BGPConnectionIpv4         `json:"bgpIpv4,omitempty"`
	BgpIpv6 *BGPConnectionIpv6         `json:"bgpIpv6,omitempty"`
	// Customer asn
	CustomerAsn *int64 `json:"customerAsn,omitempty"`
	// Equinix asn
	EquinixAsn *int64 `json:"equinixAsn,omitempty"`
	// BGP authorization key
	BgpAuthKey *string `json:"bgpAuthKey,omitempty"`
	// Enable AS number override
	AsOverrideEnabled    *bool               `json:"asOverrideEnabled,omitempty"`
	Bfd                  *RoutingProtocolBFD `json:"bfd,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RoutingProtocolBGPType RoutingProtocolBGPType

// NewRoutingProtocolBGPType instantiates a new RoutingProtocolBGPType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingProtocolBGPType(type_ RoutingProtocolBGPTypeType) *RoutingProtocolBGPType {
	this := RoutingProtocolBGPType{}
	this.Type = type_
	return &this
}

// NewRoutingProtocolBGPTypeWithDefaults instantiates a new RoutingProtocolBGPType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingProtocolBGPTypeWithDefaults() *RoutingProtocolBGPType {
	this := RoutingProtocolBGPType{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingProtocolBGPType) GetType() RoutingProtocolBGPTypeType {
	if o == nil {
		var ret RoutingProtocolBGPTypeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetTypeOk() (*RoutingProtocolBGPTypeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingProtocolBGPType) SetType(v RoutingProtocolBGPTypeType) {
	o.Type = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoutingProtocolBGPType) SetName(v string) {
	o.Name = &v
}

// GetBgpIpv4 returns the BgpIpv4 field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetBgpIpv4() BGPConnectionIpv4 {
	if o == nil || IsNil(o.BgpIpv4) {
		var ret BGPConnectionIpv4
		return ret
	}
	return *o.BgpIpv4
}

// GetBgpIpv4Ok returns a tuple with the BgpIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetBgpIpv4Ok() (*BGPConnectionIpv4, bool) {
	if o == nil || IsNil(o.BgpIpv4) {
		return nil, false
	}
	return o.BgpIpv4, true
}

// HasBgpIpv4 returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasBgpIpv4() bool {
	if o != nil && !IsNil(o.BgpIpv4) {
		return true
	}

	return false
}

// SetBgpIpv4 gets a reference to the given BGPConnectionIpv4 and assigns it to the BgpIpv4 field.
func (o *RoutingProtocolBGPType) SetBgpIpv4(v BGPConnectionIpv4) {
	o.BgpIpv4 = &v
}

// GetBgpIpv6 returns the BgpIpv6 field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetBgpIpv6() BGPConnectionIpv6 {
	if o == nil || IsNil(o.BgpIpv6) {
		var ret BGPConnectionIpv6
		return ret
	}
	return *o.BgpIpv6
}

// GetBgpIpv6Ok returns a tuple with the BgpIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetBgpIpv6Ok() (*BGPConnectionIpv6, bool) {
	if o == nil || IsNil(o.BgpIpv6) {
		return nil, false
	}
	return o.BgpIpv6, true
}

// HasBgpIpv6 returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasBgpIpv6() bool {
	if o != nil && !IsNil(o.BgpIpv6) {
		return true
	}

	return false
}

// SetBgpIpv6 gets a reference to the given BGPConnectionIpv6 and assigns it to the BgpIpv6 field.
func (o *RoutingProtocolBGPType) SetBgpIpv6(v BGPConnectionIpv6) {
	o.BgpIpv6 = &v
}

// GetCustomerAsn returns the CustomerAsn field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetCustomerAsn() int64 {
	if o == nil || IsNil(o.CustomerAsn) {
		var ret int64
		return ret
	}
	return *o.CustomerAsn
}

// GetCustomerAsnOk returns a tuple with the CustomerAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetCustomerAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.CustomerAsn) {
		return nil, false
	}
	return o.CustomerAsn, true
}

// HasCustomerAsn returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasCustomerAsn() bool {
	if o != nil && !IsNil(o.CustomerAsn) {
		return true
	}

	return false
}

// SetCustomerAsn gets a reference to the given int64 and assigns it to the CustomerAsn field.
func (o *RoutingProtocolBGPType) SetCustomerAsn(v int64) {
	o.CustomerAsn = &v
}

// GetEquinixAsn returns the EquinixAsn field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetEquinixAsn() int64 {
	if o == nil || IsNil(o.EquinixAsn) {
		var ret int64
		return ret
	}
	return *o.EquinixAsn
}

// GetEquinixAsnOk returns a tuple with the EquinixAsn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetEquinixAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.EquinixAsn) {
		return nil, false
	}
	return o.EquinixAsn, true
}

// HasEquinixAsn returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasEquinixAsn() bool {
	if o != nil && !IsNil(o.EquinixAsn) {
		return true
	}

	return false
}

// SetEquinixAsn gets a reference to the given int64 and assigns it to the EquinixAsn field.
func (o *RoutingProtocolBGPType) SetEquinixAsn(v int64) {
	o.EquinixAsn = &v
}

// GetBgpAuthKey returns the BgpAuthKey field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetBgpAuthKey() string {
	if o == nil || IsNil(o.BgpAuthKey) {
		var ret string
		return ret
	}
	return *o.BgpAuthKey
}

// GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetBgpAuthKeyOk() (*string, bool) {
	if o == nil || IsNil(o.BgpAuthKey) {
		return nil, false
	}
	return o.BgpAuthKey, true
}

// HasBgpAuthKey returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasBgpAuthKey() bool {
	if o != nil && !IsNil(o.BgpAuthKey) {
		return true
	}

	return false
}

// SetBgpAuthKey gets a reference to the given string and assigns it to the BgpAuthKey field.
func (o *RoutingProtocolBGPType) SetBgpAuthKey(v string) {
	o.BgpAuthKey = &v
}

// GetAsOverrideEnabled returns the AsOverrideEnabled field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetAsOverrideEnabled() bool {
	if o == nil || IsNil(o.AsOverrideEnabled) {
		var ret bool
		return ret
	}
	return *o.AsOverrideEnabled
}

// GetAsOverrideEnabledOk returns a tuple with the AsOverrideEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetAsOverrideEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AsOverrideEnabled) {
		return nil, false
	}
	return o.AsOverrideEnabled, true
}

// HasAsOverrideEnabled returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasAsOverrideEnabled() bool {
	if o != nil && !IsNil(o.AsOverrideEnabled) {
		return true
	}

	return false
}

// SetAsOverrideEnabled gets a reference to the given bool and assigns it to the AsOverrideEnabled field.
func (o *RoutingProtocolBGPType) SetAsOverrideEnabled(v bool) {
	o.AsOverrideEnabled = &v
}

// GetBfd returns the Bfd field value if set, zero value otherwise.
func (o *RoutingProtocolBGPType) GetBfd() RoutingProtocolBFD {
	if o == nil || IsNil(o.Bfd) {
		var ret RoutingProtocolBFD
		return ret
	}
	return *o.Bfd
}

// GetBfdOk returns a tuple with the Bfd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingProtocolBGPType) GetBfdOk() (*RoutingProtocolBFD, bool) {
	if o == nil || IsNil(o.Bfd) {
		return nil, false
	}
	return o.Bfd, true
}

// HasBfd returns a boolean if a field has been set.
func (o *RoutingProtocolBGPType) HasBfd() bool {
	if o != nil && !IsNil(o.Bfd) {
		return true
	}

	return false
}

// SetBfd gets a reference to the given RoutingProtocolBFD and assigns it to the Bfd field.
func (o *RoutingProtocolBGPType) SetBfd(v RoutingProtocolBFD) {
	o.Bfd = &v
}

func (o RoutingProtocolBGPType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingProtocolBGPType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BgpIpv4) {
		toSerialize["bgpIpv4"] = o.BgpIpv4
	}
	if !IsNil(o.BgpIpv6) {
		toSerialize["bgpIpv6"] = o.BgpIpv6
	}
	if !IsNil(o.CustomerAsn) {
		toSerialize["customerAsn"] = o.CustomerAsn
	}
	if !IsNil(o.EquinixAsn) {
		toSerialize["equinixAsn"] = o.EquinixAsn
	}
	if !IsNil(o.BgpAuthKey) {
		toSerialize["bgpAuthKey"] = o.BgpAuthKey
	}
	if !IsNil(o.AsOverrideEnabled) {
		toSerialize["asOverrideEnabled"] = o.AsOverrideEnabled
	}
	if !IsNil(o.Bfd) {
		toSerialize["bfd"] = o.Bfd
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RoutingProtocolBGPType) UnmarshalJSON(data []byte) (err error) {
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

	varRoutingProtocolBGPType := _RoutingProtocolBGPType{}

	err = json.Unmarshal(data, &varRoutingProtocolBGPType)

	if err != nil {
		return err
	}

	*o = RoutingProtocolBGPType(varRoutingProtocolBGPType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "bgpIpv4")
		delete(additionalProperties, "bgpIpv6")
		delete(additionalProperties, "customerAsn")
		delete(additionalProperties, "equinixAsn")
		delete(additionalProperties, "bgpAuthKey")
		delete(additionalProperties, "asOverrideEnabled")
		delete(additionalProperties, "bfd")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRoutingProtocolBGPType struct {
	value *RoutingProtocolBGPType
	isSet bool
}

func (v NullableRoutingProtocolBGPType) Get() *RoutingProtocolBGPType {
	return v.value
}

func (v *NullableRoutingProtocolBGPType) Set(val *RoutingProtocolBGPType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingProtocolBGPType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingProtocolBGPType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingProtocolBGPType(val *RoutingProtocolBGPType) *NullableRoutingProtocolBGPType {
	return &NullableRoutingProtocolBGPType{value: val, isSet: true}
}

func (v NullableRoutingProtocolBGPType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingProtocolBGPType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
