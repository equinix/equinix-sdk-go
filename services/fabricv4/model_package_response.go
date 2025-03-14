/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the PackageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackageResponse{}

// PackageResponse struct for PackageResponse
type PackageResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []CloudRouterPackage `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PackageResponse PackageResponse

// NewPackageResponse instantiates a new PackageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackageResponse() *PackageResponse {
	this := PackageResponse{}
	return &this
}

// NewPackageResponseWithDefaults instantiates a new PackageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageResponseWithDefaults() *PackageResponse {
	this := PackageResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PackageResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PackageResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *PackageResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PackageResponse) GetData() []CloudRouterPackage {
	if o == nil || IsNil(o.Data) {
		var ret []CloudRouterPackage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackageResponse) GetDataOk() ([]CloudRouterPackage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PackageResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CloudRouterPackage and assigns it to the Data field.
func (o *PackageResponse) SetData(v []CloudRouterPackage) {
	o.Data = v
}

func (o PackageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PackageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PackageResponse) UnmarshalJSON(data []byte) (err error) {
	varPackageResponse := _PackageResponse{}

	err = json.Unmarshal(data, &varPackageResponse)

	if err != nil {
		return err
	}

	*o = PackageResponse(varPackageResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePackageResponse struct {
	value *PackageResponse
	isSet bool
}

func (v NullablePackageResponse) Get() *PackageResponse {
	return v.value
}

func (v *NullablePackageResponse) Set(val *PackageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePackageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePackageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackageResponse(val *PackageResponse) *NullablePackageResponse {
	return &NullablePackageResponse{value: val, isSet: true}
}

func (v NullablePackageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
