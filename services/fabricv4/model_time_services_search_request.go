/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the TimeServicesSearchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeServicesSearchRequest{}

// TimeServicesSearchRequest Search requests containing criteria
type TimeServicesSearchRequest struct {
	Filter               *TimeServiceFilters       `json:"filter,omitempty"`
	Pagination           *PaginationRequest        `json:"pagination,omitempty"`
	Sort                 []TimeServiceSortCriteria `json:"sort,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TimeServicesSearchRequest TimeServicesSearchRequest

// NewTimeServicesSearchRequest instantiates a new TimeServicesSearchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeServicesSearchRequest() *TimeServicesSearchRequest {
	this := TimeServicesSearchRequest{}
	return &this
}

// NewTimeServicesSearchRequestWithDefaults instantiates a new TimeServicesSearchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeServicesSearchRequestWithDefaults() *TimeServicesSearchRequest {
	this := TimeServicesSearchRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *TimeServicesSearchRequest) GetFilter() TimeServiceFilters {
	if o == nil || IsNil(o.Filter) {
		var ret TimeServiceFilters
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServicesSearchRequest) GetFilterOk() (*TimeServiceFilters, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *TimeServicesSearchRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given TimeServiceFilters and assigns it to the Filter field.
func (o *TimeServicesSearchRequest) SetFilter(v TimeServiceFilters) {
	o.Filter = &v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *TimeServicesSearchRequest) GetPagination() PaginationRequest {
	if o == nil || IsNil(o.Pagination) {
		var ret PaginationRequest
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServicesSearchRequest) GetPaginationOk() (*PaginationRequest, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *TimeServicesSearchRequest) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given PaginationRequest and assigns it to the Pagination field.
func (o *TimeServicesSearchRequest) SetPagination(v PaginationRequest) {
	o.Pagination = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *TimeServicesSearchRequest) GetSort() []TimeServiceSortCriteria {
	if o == nil || IsNil(o.Sort) {
		var ret []TimeServiceSortCriteria
		return ret
	}
	return o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeServicesSearchRequest) GetSortOk() ([]TimeServiceSortCriteria, bool) {
	if o == nil || IsNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *TimeServicesSearchRequest) HasSort() bool {
	if o != nil && !IsNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given []TimeServiceSortCriteria and assigns it to the Sort field.
func (o *TimeServicesSearchRequest) SetSort(v []TimeServiceSortCriteria) {
	o.Sort = v
}

func (o TimeServicesSearchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeServicesSearchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TimeServicesSearchRequest) UnmarshalJSON(data []byte) (err error) {
	varTimeServicesSearchRequest := _TimeServicesSearchRequest{}

	err = json.Unmarshal(data, &varTimeServicesSearchRequest)

	if err != nil {
		return err
	}

	*o = TimeServicesSearchRequest(varTimeServicesSearchRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "filter")
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTimeServicesSearchRequest struct {
	value *TimeServicesSearchRequest
	isSet bool
}

func (v NullableTimeServicesSearchRequest) Get() *TimeServicesSearchRequest {
	return v.value
}

func (v *NullableTimeServicesSearchRequest) Set(val *TimeServicesSearchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeServicesSearchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeServicesSearchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeServicesSearchRequest(val *TimeServicesSearchRequest) *NullableTimeServicesSearchRequest {
	return &NullableTimeServicesSearchRequest{value: val, isSet: true}
}

func (v NullableTimeServicesSearchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeServicesSearchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
