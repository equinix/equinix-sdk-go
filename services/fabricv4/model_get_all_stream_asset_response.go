/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the GetAllStreamAssetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllStreamAssetResponse{}

// GetAllStreamAssetResponse struct for GetAllStreamAssetResponse
type GetAllStreamAssetResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []StreamAsset `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAllStreamAssetResponse GetAllStreamAssetResponse

// NewGetAllStreamAssetResponse instantiates a new GetAllStreamAssetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllStreamAssetResponse() *GetAllStreamAssetResponse {
	this := GetAllStreamAssetResponse{}
	return &this
}

// NewGetAllStreamAssetResponseWithDefaults instantiates a new GetAllStreamAssetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllStreamAssetResponseWithDefaults() *GetAllStreamAssetResponse {
	this := GetAllStreamAssetResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetAllStreamAssetResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllStreamAssetResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetAllStreamAssetResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetAllStreamAssetResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllStreamAssetResponse) GetData() []StreamAsset {
	if o == nil || IsNil(o.Data) {
		var ret []StreamAsset
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllStreamAssetResponse) GetDataOk() ([]StreamAsset, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllStreamAssetResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []StreamAsset and assigns it to the Data field.
func (o *GetAllStreamAssetResponse) SetData(v []StreamAsset) {
	o.Data = v
}

func (o GetAllStreamAssetResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllStreamAssetResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetAllStreamAssetResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAllStreamAssetResponse := _GetAllStreamAssetResponse{}

	err = json.Unmarshal(data, &varGetAllStreamAssetResponse)

	if err != nil {
		return err
	}

	*o = GetAllStreamAssetResponse(varGetAllStreamAssetResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAllStreamAssetResponse struct {
	value *GetAllStreamAssetResponse
	isSet bool
}

func (v NullableGetAllStreamAssetResponse) Get() *GetAllStreamAssetResponse {
	return v.value
}

func (v *NullableGetAllStreamAssetResponse) Set(val *GetAllStreamAssetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllStreamAssetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllStreamAssetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllStreamAssetResponse(val *GetAllStreamAssetResponse) *NullableGetAllStreamAssetResponse {
	return &NullableGetAllStreamAssetResponse{value: val, isSet: true}
}

func (v NullableGetAllStreamAssetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllStreamAssetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
