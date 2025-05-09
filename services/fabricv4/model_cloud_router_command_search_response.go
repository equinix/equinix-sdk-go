/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the CloudRouterCommandSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CloudRouterCommandSearchResponse{}

// CloudRouterCommandSearchResponse struct for CloudRouterCommandSearchResponse
type CloudRouterCommandSearchResponse struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// Data returned from the API call.
	Data                 []CloudRouterCommand `json:"data,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CloudRouterCommandSearchResponse CloudRouterCommandSearchResponse

// NewCloudRouterCommandSearchResponse instantiates a new CloudRouterCommandSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudRouterCommandSearchResponse() *CloudRouterCommandSearchResponse {
	this := CloudRouterCommandSearchResponse{}
	return &this
}

// NewCloudRouterCommandSearchResponseWithDefaults instantiates a new CloudRouterCommandSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudRouterCommandSearchResponseWithDefaults() *CloudRouterCommandSearchResponse {
	this := CloudRouterCommandSearchResponse{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *CloudRouterCommandSearchResponse) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommandSearchResponse) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *CloudRouterCommandSearchResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *CloudRouterCommandSearchResponse) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CloudRouterCommandSearchResponse) GetData() []CloudRouterCommand {
	if o == nil || IsNil(o.Data) {
		var ret []CloudRouterCommand
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudRouterCommandSearchResponse) GetDataOk() ([]CloudRouterCommand, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CloudRouterCommandSearchResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CloudRouterCommand and assigns it to the Data field.
func (o *CloudRouterCommandSearchResponse) SetData(v []CloudRouterCommand) {
	o.Data = v
}

func (o CloudRouterCommandSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CloudRouterCommandSearchResponse) ToMap() (map[string]interface{}, error) {
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

func (o *CloudRouterCommandSearchResponse) UnmarshalJSON(data []byte) (err error) {
	varCloudRouterCommandSearchResponse := _CloudRouterCommandSearchResponse{}

	err = json.Unmarshal(data, &varCloudRouterCommandSearchResponse)

	if err != nil {
		return err
	}

	*o = CloudRouterCommandSearchResponse(varCloudRouterCommandSearchResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCloudRouterCommandSearchResponse struct {
	value *CloudRouterCommandSearchResponse
	isSet bool
}

func (v NullableCloudRouterCommandSearchResponse) Get() *CloudRouterCommandSearchResponse {
	return v.value
}

func (v *NullableCloudRouterCommandSearchResponse) Set(val *CloudRouterCommandSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudRouterCommandSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudRouterCommandSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudRouterCommandSearchResponse(val *CloudRouterCommandSearchResponse) *NullableCloudRouterCommandSearchResponse {
	return &NullableCloudRouterCommandSearchResponse{value: val, isSet: true}
}

func (v NullableCloudRouterCommandSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudRouterCommandSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
