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

// checks if the ServiceProfileSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceProfileSortCriteria{}

// ServiceProfileSortCriteria struct for ServiceProfileSortCriteria
type ServiceProfileSortCriteria struct {
	Direction            *ServiceProfileSortDirection `json:"direction,omitempty"`
	Property             *ServiceProfileSortBy        `json:"property,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ServiceProfileSortCriteria ServiceProfileSortCriteria

// NewServiceProfileSortCriteria instantiates a new ServiceProfileSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProfileSortCriteria() *ServiceProfileSortCriteria {
	this := ServiceProfileSortCriteria{}
	var direction ServiceProfileSortDirection = SERVICEPROFILESORTDIRECTION_DESC
	this.Direction = &direction
	var property ServiceProfileSortBy = SERVICEPROFILESORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// NewServiceProfileSortCriteriaWithDefaults instantiates a new ServiceProfileSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProfileSortCriteriaWithDefaults() *ServiceProfileSortCriteria {
	this := ServiceProfileSortCriteria{}
	var direction ServiceProfileSortDirection = SERVICEPROFILESORTDIRECTION_DESC
	this.Direction = &direction
	var property ServiceProfileSortBy = SERVICEPROFILESORTBY_CHANGE_LOG_UPDATED_DATE_TIME
	this.Property = &property
	return &this
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ServiceProfileSortCriteria) GetDirection() ServiceProfileSortDirection {
	if o == nil || IsNil(o.Direction) {
		var ret ServiceProfileSortDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileSortCriteria) GetDirectionOk() (*ServiceProfileSortDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ServiceProfileSortCriteria) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given ServiceProfileSortDirection and assigns it to the Direction field.
func (o *ServiceProfileSortCriteria) SetDirection(v ServiceProfileSortDirection) {
	o.Direction = &v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *ServiceProfileSortCriteria) GetProperty() ServiceProfileSortBy {
	if o == nil || IsNil(o.Property) {
		var ret ServiceProfileSortBy
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProfileSortCriteria) GetPropertyOk() (*ServiceProfileSortBy, bool) {
	if o == nil || IsNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *ServiceProfileSortCriteria) HasProperty() bool {
	if o != nil && !IsNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given ServiceProfileSortBy and assigns it to the Property field.
func (o *ServiceProfileSortCriteria) SetProperty(v ServiceProfileSortBy) {
	o.Property = &v
}

func (o ServiceProfileSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceProfileSortCriteria) ToMap() (map[string]interface{}, error) {
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

func (o *ServiceProfileSortCriteria) UnmarshalJSON(data []byte) (err error) {
	varServiceProfileSortCriteria := _ServiceProfileSortCriteria{}

	err = json.Unmarshal(data, &varServiceProfileSortCriteria)

	if err != nil {
		return err
	}

	*o = ServiceProfileSortCriteria(varServiceProfileSortCriteria)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "direction")
		delete(additionalProperties, "property")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceProfileSortCriteria struct {
	value *ServiceProfileSortCriteria
	isSet bool
}

func (v NullableServiceProfileSortCriteria) Get() *ServiceProfileSortCriteria {
	return v.value
}

func (v *NullableServiceProfileSortCriteria) Set(val *ServiceProfileSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProfileSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProfileSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProfileSortCriteria(val *ServiceProfileSortCriteria) *NullableServiceProfileSortCriteria {
	return &NullableServiceProfileSortCriteria{value: val, isSet: true}
}

func (v NullableServiceProfileSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProfileSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}