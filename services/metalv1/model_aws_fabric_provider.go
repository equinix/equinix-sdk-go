/*
Metal API

Contact: support@equinixmetal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metalv1

import (
	"encoding/json"
	"fmt"
)

// checks if the AWSFabricProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AWSFabricProvider{}

// AWSFabricProvider struct for AWSFabricProvider
type AWSFabricProvider struct {
	Type AWSFabricProviderType `json:"type"`
	// AWS Account ID
	AccountId            string  `json:"account_id" validate:"regexp=^\\\\d{12}$"`
	Location             *string `json:"location,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AWSFabricProvider AWSFabricProvider

// NewAWSFabricProvider instantiates a new AWSFabricProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAWSFabricProvider(type_ AWSFabricProviderType, accountId string) *AWSFabricProvider {
	this := AWSFabricProvider{}
	this.Type = type_
	this.AccountId = accountId
	return &this
}

// NewAWSFabricProviderWithDefaults instantiates a new AWSFabricProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAWSFabricProviderWithDefaults() *AWSFabricProvider {
	this := AWSFabricProvider{}
	return &this
}

// GetType returns the Type field value
func (o *AWSFabricProvider) GetType() AWSFabricProviderType {
	if o == nil {
		var ret AWSFabricProviderType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AWSFabricProvider) GetTypeOk() (*AWSFabricProviderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AWSFabricProvider) SetType(v AWSFabricProviderType) {
	o.Type = v
}

// GetAccountId returns the AccountId field value
func (o *AWSFabricProvider) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSFabricProvider) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AWSFabricProvider) SetAccountId(v string) {
	o.AccountId = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *AWSFabricProvider) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSFabricProvider) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *AWSFabricProvider) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *AWSFabricProvider) SetLocation(v string) {
	o.Location = &v
}

func (o AWSFabricProvider) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AWSFabricProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AWSFabricProvider) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"account_id",
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

	varAWSFabricProvider := _AWSFabricProvider{}

	err = json.Unmarshal(data, &varAWSFabricProvider)

	if err != nil {
		return err
	}

	*o = AWSFabricProvider(varAWSFabricProvider)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "location")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAWSFabricProvider struct {
	value *AWSFabricProvider
	isSet bool
}

func (v NullableAWSFabricProvider) Get() *AWSFabricProvider {
	return v.value
}

func (v *NullableAWSFabricProvider) Set(val *AWSFabricProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableAWSFabricProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableAWSFabricProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAWSFabricProvider(val *AWSFabricProvider) *NullableAWSFabricProvider {
	return &NullableAWSFabricProvider{value: val, isSet: true}
}

func (v NullableAWSFabricProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAWSFabricProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
