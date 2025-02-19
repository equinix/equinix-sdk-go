/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterPackage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterPackage{}

// CloudRouterPackage Fabric Cloud Router Package
type CloudRouterPackage struct {
	// Cloud Router package URI
	Href *string                 `json:"href,omitempty"`
	Type *CloudRouterPackageType `json:"type,omitempty"`
	Code *Code                   `json:"code,omitempty"`
	// Fabric Cloud Router Package description
	Description *string `json:"description,omitempty"`
	// Cloud Router package BGP IPv4 routes limit
	TotalIPv4RoutesMax *int32 `json:"totalIPv4RoutesMax,omitempty"`
	// Cloud Router package BGP IPv6 routes limit
	TotalIPv6RoutesMax *int32 `json:"totalIPv6RoutesMax,omitempty"`
	// CloudRouter package route filter support
	RouteFilterSupported *bool `json:"routeFilterSupported,omitempty"`
	// CloudRouter package Max Connection limit
	VcCountMax *int32 `json:"vcCountMax,omitempty"`
	// CloudRouter package Max CloudRouter limit
	CrCountMax *int32 `json:"crCountMax,omitempty"`
	// CloudRouter package Max Bandwidth limit
	VcBandwidthMax       *int32            `json:"vcBandwidthMax,omitempty"`
	ChangeLog            *PackageChangeLog `json:"changeLog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterPackage CloudRouterPackage

// NewCloudRouterPackage instantiates a new CloudRouterPackage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterPackage() *CloudRouterPackage {
	this := CloudRouterPackage{}
	return &this
}

// NewCloudRouterPackageWithDefaults instantiates a new CloudRouterPackage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterPackageWithDefaults() *CloudRouterPackage {
	this := CloudRouterPackage{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *CloudRouterPackage) SetHref(v string) {
	o.Href = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetType() CloudRouterPackageType {
	if o == nil || IsNil(o.Type) {
		var ret CloudRouterPackageType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetTypeOk() (*CloudRouterPackageType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given CloudRouterPackageType and assigns it to the Type field.
func (o *CloudRouterPackage) SetType(v CloudRouterPackageType) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetCode() Code {
	if o == nil || IsNil(o.Code) {
		var ret Code
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetCodeOk() (*Code, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given Code and assigns it to the Code field.
func (o *CloudRouterPackage) SetCode(v Code) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudRouterPackage) SetDescription(v string) {
	o.Description = &v
}

// GetTotalIPv4RoutesMax returns the TotalIPv4RoutesMax field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetTotalIPv4RoutesMax() int32 {
	if o == nil || IsNil(o.TotalIPv4RoutesMax) {
		var ret int32
		return ret
	}
	return *o.TotalIPv4RoutesMax
}

// GetTotalIPv4RoutesMaxOk returns a tuple with the TotalIPv4RoutesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetTotalIPv4RoutesMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalIPv4RoutesMax) {
		return nil, false
	}
	return o.TotalIPv4RoutesMax, true
}

// HasTotalIPv4RoutesMax returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasTotalIPv4RoutesMax() bool {
	if o != nil && !IsNil(o.TotalIPv4RoutesMax) {
		return true
	}

	return false
}

// SetTotalIPv4RoutesMax gets a reference to the given int32 and assigns it to the TotalIPv4RoutesMax field.
func (o *CloudRouterPackage) SetTotalIPv4RoutesMax(v int32) {
	o.TotalIPv4RoutesMax = &v
}

// GetTotalIPv6RoutesMax returns the TotalIPv6RoutesMax field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetTotalIPv6RoutesMax() int32 {
	if o == nil || IsNil(o.TotalIPv6RoutesMax) {
		var ret int32
		return ret
	}
	return *o.TotalIPv6RoutesMax
}

// GetTotalIPv6RoutesMaxOk returns a tuple with the TotalIPv6RoutesMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetTotalIPv6RoutesMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalIPv6RoutesMax) {
		return nil, false
	}
	return o.TotalIPv6RoutesMax, true
}

// HasTotalIPv6RoutesMax returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasTotalIPv6RoutesMax() bool {
	if o != nil && !IsNil(o.TotalIPv6RoutesMax) {
		return true
	}

	return false
}

// SetTotalIPv6RoutesMax gets a reference to the given int32 and assigns it to the TotalIPv6RoutesMax field.
func (o *CloudRouterPackage) SetTotalIPv6RoutesMax(v int32) {
	o.TotalIPv6RoutesMax = &v
}

// GetRouteFilterSupported returns the RouteFilterSupported field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetRouteFilterSupported() bool {
	if o == nil || IsNil(o.RouteFilterSupported) {
		var ret bool
		return ret
	}
	return *o.RouteFilterSupported
}

// GetRouteFilterSupportedOk returns a tuple with the RouteFilterSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetRouteFilterSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.RouteFilterSupported) {
		return nil, false
	}
	return o.RouteFilterSupported, true
}

// HasRouteFilterSupported returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasRouteFilterSupported() bool {
	if o != nil && !IsNil(o.RouteFilterSupported) {
		return true
	}

	return false
}

// SetRouteFilterSupported gets a reference to the given bool and assigns it to the RouteFilterSupported field.
func (o *CloudRouterPackage) SetRouteFilterSupported(v bool) {
	o.RouteFilterSupported = &v
}

// GetVcCountMax returns the VcCountMax field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetVcCountMax() int32 {
	if o == nil || IsNil(o.VcCountMax) {
		var ret int32
		return ret
	}
	return *o.VcCountMax
}

// GetVcCountMaxOk returns a tuple with the VcCountMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetVcCountMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.VcCountMax) {
		return nil, false
	}
	return o.VcCountMax, true
}

// HasVcCountMax returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasVcCountMax() bool {
	if o != nil && !IsNil(o.VcCountMax) {
		return true
	}

	return false
}

// SetVcCountMax gets a reference to the given int32 and assigns it to the VcCountMax field.
func (o *CloudRouterPackage) SetVcCountMax(v int32) {
	o.VcCountMax = &v
}

// GetCrCountMax returns the CrCountMax field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetCrCountMax() int32 {
	if o == nil || IsNil(o.CrCountMax) {
		var ret int32
		return ret
	}
	return *o.CrCountMax
}

// GetCrCountMaxOk returns a tuple with the CrCountMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetCrCountMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.CrCountMax) {
		return nil, false
	}
	return o.CrCountMax, true
}

// HasCrCountMax returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasCrCountMax() bool {
	if o != nil && !IsNil(o.CrCountMax) {
		return true
	}

	return false
}

// SetCrCountMax gets a reference to the given int32 and assigns it to the CrCountMax field.
func (o *CloudRouterPackage) SetCrCountMax(v int32) {
	o.CrCountMax = &v
}

// GetVcBandwidthMax returns the VcBandwidthMax field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetVcBandwidthMax() int32 {
	if o == nil || IsNil(o.VcBandwidthMax) {
		var ret int32
		return ret
	}
	return *o.VcBandwidthMax
}

// GetVcBandwidthMaxOk returns a tuple with the VcBandwidthMax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetVcBandwidthMaxOk() (*int32, bool) {
	if o == nil || IsNil(o.VcBandwidthMax) {
		return nil, false
	}
	return o.VcBandwidthMax, true
}

// HasVcBandwidthMax returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasVcBandwidthMax() bool {
	if o != nil && !IsNil(o.VcBandwidthMax) {
		return true
	}

	return false
}

// SetVcBandwidthMax gets a reference to the given int32 and assigns it to the VcBandwidthMax field.
func (o *CloudRouterPackage) SetVcBandwidthMax(v int32) {
	o.VcBandwidthMax = &v
}

// GetChangeLog returns the ChangeLog field value if set, zero value otherwise.
func (o *CloudRouterPackage) GetChangeLog() PackageChangeLog {
	if o == nil || IsNil(o.ChangeLog) {
		var ret PackageChangeLog
		return ret
	}
	return *o.ChangeLog
}

// GetChangeLogOk returns a tuple with the ChangeLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterPackage) GetChangeLogOk() (*PackageChangeLog, bool) {
	if o == nil || IsNil(o.ChangeLog) {
		return nil, false
	}
	return o.ChangeLog, true
}

// HasChangeLog returns a boolean if a field has been set.
func (o *CloudRouterPackage) HasChangeLog() bool {
	if o != nil && !IsNil(o.ChangeLog) {
		return true
	}

	return false
}

// SetChangeLog gets a reference to the given PackageChangeLog and assigns it to the ChangeLog field.
func (o *CloudRouterPackage) SetChangeLog(v PackageChangeLog) {
	o.ChangeLog = &v
}

func (o CloudRouterPackage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterPackage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TotalIPv4RoutesMax) {
		toSerialize["totalIPv4RoutesMax"] = o.TotalIPv4RoutesMax
	}
	if !IsNil(o.TotalIPv6RoutesMax) {
		toSerialize["totalIPv6RoutesMax"] = o.TotalIPv6RoutesMax
	}
	if !IsNil(o.RouteFilterSupported) {
		toSerialize["routeFilterSupported"] = o.RouteFilterSupported
	}
	if !IsNil(o.VcCountMax) {
		toSerialize["vcCountMax"] = o.VcCountMax
	}
	if !IsNil(o.CrCountMax) {
		toSerialize["crCountMax"] = o.CrCountMax
	}
	if !IsNil(o.VcBandwidthMax) {
		toSerialize["vcBandwidthMax"] = o.VcBandwidthMax
	}
	if !IsNil(o.ChangeLog) {
		toSerialize["changeLog"] = o.ChangeLog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CloudRouterPackage) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterPackage := _CloudRouterPackage{}

	err = json.Unmarshal(data, &varCloudRouterPackage)

	if err != nil {
		return err
	}

	*o = CloudRouterPackage(varCloudRouterPackage)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "href")
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "description")
		delete(additionalProperties, "totalIPv4RoutesMax")
		delete(additionalProperties, "totalIPv6RoutesMax")
		delete(additionalProperties, "routeFilterSupported")
		delete(additionalProperties, "vcCountMax")
		delete(additionalProperties, "crCountMax")
		delete(additionalProperties, "vcBandwidthMax")
		delete(additionalProperties, "changeLog")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterPackage struct {
	value *CloudRouterPackage
	isSet bool
}

func (v NullableCloudRouterPackage) Get() *CloudRouterPackage {
	return v.value
}

func (v *NullableCloudRouterPackage) Set(val *CloudRouterPackage) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterPackage) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterPackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterPackage(val *CloudRouterPackage) *NullableCloudRouterPackage {
	return &NullableCloudRouterPackage{value: val, isSet: true}
}

func (v NullableCloudRouterPackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterPackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
