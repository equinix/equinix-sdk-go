/*
Equinix Fabric API v4

Equinix Fabric is an advanced software-defined interconnection solution that enables you to directly, securely and dynamically connect to distributed infrastructure and digital ecosystems on platform Equinix via a single port, Customers can use Fabric to connect to: </br> 1. Cloud Service Providers - Clouds, network and other service providers.  </br> 2. Enterprises - Other Equinix customers, vendors and partners.  </br> 3. Myself - Another customer instance deployed at Equinix. </br>

API version: 4.12
Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
)

// checks if the NetworkSortCriteriaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSortCriteriaResponse{}

// NetworkSortCriteriaResponse struct for NetworkSortCriteriaResponse
type NetworkSortCriteriaResponse struct {
	Direction            *NetworkSortDirectionResponse `json:"direction,omitempty"`
	Property             *NetworkSortByResponse        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NetworkSortCriteriaResponse NetworkSortCriteriaResponse

// NewNetworkSortCriteriaResponse instantiates a new NetworkSortCriteriaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSortCriteriaResponse() *NetworkSortCriteriaResponse {
	this := NetworkSortCriteriaResponse{}
	var direction NetworkSortDirectionResponse = NETWORKSORTDIRECTIONRESPONSE_DESC
	this.Direction = &direction
	var property NetworkSortByResponse = NETWORKSORTBYRESPONSE_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewNetworkSortCriteriaResponseWithDefaults instantiates a new NetworkSortCriteriaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSortCriteriaResponseWithDefaults() *NetworkSortCriteriaResponse {
	this := NetworkSortCriteriaResponse{}
	var direction NetworkSortDirectionResponse = NETWORKSORTDIRECTIONRESPONSE_DESC
	this.Direction = &direction
	var property NetworkSortByResponse = NETWORKSORTBYRESPONSE_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *NetworkSortCriteriaResponse) GetDirection() NetworkSortDirectionResponse {
	if o == nil || IsNil(o.Direction) {
		var ret NetworkSortDirectionResponse
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSortCriteriaResponse) GetDirectionOk() (*NetworkSortDirectionResponse, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *NetworkSortCriteriaResponse) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NetworkSortDirectionResponse and assigns it to the Direction field.
func (o *NetworkSortCriteriaResponse) SetDirection(v NetworkSortDirectionResponse) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *NetworkSortCriteriaResponse) GetProperty() NetworkSortByResponse {
	if o == nil || IsNil(o.Property) {
		var ret NetworkSortByResponse
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSortCriteriaResponse) GetPropertyOk() (*NetworkSortByResponse, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *NetworkSortCriteriaResponse) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given NetworkSortByResponse and assigns it to the Property field.
func (o *NetworkSortCriteriaResponse) SetProperty(v NetworkSortByResponse) {
	o.Property = &v
}

func (o NetworkSortCriteriaResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSortCriteriaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.Property) {
		toSerialize["property"] = o.Property
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NetworkSortCriteriaResponse) UnmarshalJSON(data []byte) (err error) {
	varNetworkSortCriteriaResponse := _NetworkSortCriteriaResponse{}

	err = json.Unmarshal(data, &varNetworkSortCriteriaResponse)

	if err != nil {
		return err
	}

	*o = NetworkSortCriteriaResponse(varNetworkSortCriteriaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetworkSortCriteriaResponse struct {
	value *NetworkSortCriteriaResponse
	isSet bool
}

func (v NullableNetworkSortCriteriaResponse) Get() *NetworkSortCriteriaResponse {
	return v.value
}

func (v *NullableNetworkSortCriteriaResponse) Set(val *NetworkSortCriteriaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSortCriteriaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSortCriteriaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSortCriteriaResponse(val *NetworkSortCriteriaResponse) *NullableNetworkSortCriteriaResponse {
	return &NullableNetworkSortCriteriaResponse{value: val, isSet: true}
}

func (v NullableNetworkSortCriteriaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSortCriteriaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}