/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the GetAllConnectionRouteAggregationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllConnectionRouteAggregationsResponse{}

// GetAllConnectionRouteAggregationsResponse struct for GetAllConnectionRouteAggregationsResponse
type GetAllConnectionRouteAggregationsResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// List of Route Aggregations attached to a Connection
	Data                 []ConnectionRouteAggregationData `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GetAllConnectionRouteAggregationsResponse GetAllConnectionRouteAggregationsResponse

// NewGetAllConnectionRouteAggregationsResponse instantiates a new GetAllConnectionRouteAggregationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllConnectionRouteAggregationsResponse() *GetAllConnectionRouteAggregationsResponse {
	this := GetAllConnectionRouteAggregationsResponse{}
	return &this
}

// NewGetAllConnectionRouteAggregationsResponseWithDefaults instantiates a new GetAllConnectionRouteAggregationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllConnectionRouteAggregationsResponseWithDefaults() *GetAllConnectionRouteAggregationsResponse {
	this := GetAllConnectionRouteAggregationsResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *GetAllConnectionRouteAggregationsResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllConnectionRouteAggregationsResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *GetAllConnectionRouteAggregationsResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *GetAllConnectionRouteAggregationsResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllConnectionRouteAggregationsResponse) GetData() []ConnectionRouteAggregationData {
	if o == nil || IsNil(o.Data) {
		var ret []ConnectionRouteAggregationData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllConnectionRouteAggregationsResponse) GetDataOk() ([]ConnectionRouteAggregationData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllConnectionRouteAggregationsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ConnectionRouteAggregationData and assigns it to the Data field.
func (o *GetAllConnectionRouteAggregationsResponse) SetData(v []ConnectionRouteAggregationData) {
	o.Data = v
}

func (o GetAllConnectionRouteAggregationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllConnectionRouteAggregationsResponse) ToMap() (map[string]interface{}, error) {
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

func (o *GetAllConnectionRouteAggregationsResponse) UnmarshalJSON(data []byte) (err error) {
	varGetAllConnectionRouteAggregationsResponse := _GetAllConnectionRouteAggregationsResponse{}

	err = json.Unmarshal(data, &varGetAllConnectionRouteAggregationsResponse)

	if err != nil {
		return err
	}

	*o = GetAllConnectionRouteAggregationsResponse(varGetAllConnectionRouteAggregationsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGetAllConnectionRouteAggregationsResponse struct {
	value *GetAllConnectionRouteAggregationsResponse
	isSet bool
}

func (v NullableGetAllConnectionRouteAggregationsResponse) Get() *GetAllConnectionRouteAggregationsResponse {
	return v.value
}

func (v *NullableGetAllConnectionRouteAggregationsResponse) Set(val *GetAllConnectionRouteAggregationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllConnectionRouteAggregationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllConnectionRouteAggregationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllConnectionRouteAggregationsResponse(val *GetAllConnectionRouteAggregationsResponse) *NullableGetAllConnectionRouteAggregationsResponse {
	return &NullableGetAllConnectionRouteAggregationsResponse{value: val, isSet: true}
}

func (v NullableGetAllConnectionRouteAggregationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllConnectionRouteAggregationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}