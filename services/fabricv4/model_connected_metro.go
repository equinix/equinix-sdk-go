/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the ConnectedMetro type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectedMetro{}

// ConnectedMetro Arrays of objects containing latency data for the specified metros
type ConnectedMetro struct {
	// The Canonical URL at which the resource resides.
	Href *string `json:"href,omitempty"`
	// Code assigned to an Equinix International Business Exchange (IBX) data center in a specified metropolitan area.
	Code *string `json:"code,omitempty"`
	// Average latency (in milliseconds[ms]) between two specified metros.
	AvgLatency *float32 `json:"avgLatency,omitempty"`
	// This field holds the Max Connection speed with connected metros
	RemoteVCBandwidthMax *int64 `json:"remoteVCBandwidthMax,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectedMetro ConnectedMetro

// NewConnectedMetro instantiates a new ConnectedMetro object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectedMetro() *ConnectedMetro {
	this := ConnectedMetro{}
	return &this
}

// NewConnectedMetroWithDefaults instantiates a new ConnectedMetro object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectedMetroWithDefaults() *ConnectedMetro {
	this := ConnectedMetro{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *ConnectedMetro) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedMetro) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *ConnectedMetro) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *ConnectedMetro) SetHref(v string) {
	o.Href = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ConnectedMetro) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedMetro) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ConnectedMetro) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ConnectedMetro) SetCode(v string) {
	o.Code = &v
}

// GetAvgLatency returns the AvgLatency field value if set, zero value otherwise.
func (o *ConnectedMetro) GetAvgLatency() float32 {
	if o == nil || IsNil(o.AvgLatency) {
		var ret float32
		return ret
	}
	return *o.AvgLatency
}

// GetAvgLatencyOk returns a tuple with the AvgLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedMetro) GetAvgLatencyOk() (*float32, bool) {
	if o == nil || IsNil(o.AvgLatency) {
		return nil, false
	}
	return o.AvgLatency, true
}

// HasAvgLatency returns a boolean if a field has been set.
func (o *ConnectedMetro) HasAvgLatency() bool {
	if o != nil && !IsNil(o.AvgLatency) {
		return true
	}

	return false
}

// SetAvgLatency gets a reference to the given float32 and assigns it to the AvgLatency field.
func (o *ConnectedMetro) SetAvgLatency(v float32) {
	o.AvgLatency = &v
}

// GetRemoteVCBandwidthMax returns the RemoteVCBandwidthMax field value if set, zero value otherwise.
func (o *ConnectedMetro) GetRemoteVCBandwidthMax() int64 {
	if o == nil || IsNil(o.RemoteVCBandwidthMax) {
		var ret int64
		return ret
	}
	return *o.RemoteVCBandwidthMax
}

// GetRemoteVCBandwidthMaxOk returns a tuple with the RemoteVCBandwidthMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectedMetro) GetRemoteVCBandwidthMaxOk() (*int64, bool) {
	if o == nil || IsNil(o.RemoteVCBandwidthMax) {
		return nil, false
	}
	return o.RemoteVCBandwidthMax, true
}

// HasRemoteVCBandwidthMax returns a boolean if a field has been set.
func (o *ConnectedMetro) HasRemoteVCBandwidthMax() bool {
	if o != nil && !IsNil(o.RemoteVCBandwidthMax) {
		return true
	}

	return false
}

// SetRemoteVCBandwidthMax gets a reference to the given int64 and assigns it to the RemoteVCBandwidthMax field.
func (o *ConnectedMetro) SetRemoteVCBandwidthMax(v int64) {
	o.RemoteVCBandwidthMax = &v
}

func (o ConnectedMetro) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectedMetro) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.AvgLatency) {
		toSerialize["avgLatency"] = o.AvgLatency
	}
	if !IsNil(o.RemoteVCBandwidthMax) {
		toSerialize["remoteVCBandwidthMax"] = o.RemoteVCBandwidthMax
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectedMetro) UnmarshalJSON(data []byte) (err error) {
	varConnectedMetro := _ConnectedMetro{}

	err = json.Unmarshal(data, &varConnectedMetro)

	if err != nil {
		return err
	}

	*o = ConnectedMetro(varConnectedMetro)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "code")
		delete(additionalProperties, "avgLatency")
		delete(additionalProperties, "remoteVCBandwidthMax")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectedMetro struct {
	value *ConnectedMetro
	isSet bool
}

func (v NullableConnectedMetro) Get() *ConnectedMetro {
	return v.value
}

func (v *NullableConnectedMetro) Set(val *ConnectedMetro) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectedMetro) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectedMetro) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectedMetro(val *ConnectedMetro) *NullableConnectedMetro {
	return &NullableConnectedMetro{value: val, isSet: true}
}

func (v NullableConnectedMetro) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectedMetro) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
