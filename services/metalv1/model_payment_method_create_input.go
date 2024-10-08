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

// checks if the PaymentMethodCreateInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentMethodCreateInput{}

// PaymentMethodCreateInput struct for PaymentMethodCreateInput
type PaymentMethodCreateInput struct {
	Default              *bool  `json:"default,omitempty"`
	Name                 string `json:"name"`
	Nonce                string `json:"nonce"`
	AdditionalProperties map[string]interface{}
}

type _PaymentMethodCreateInput PaymentMethodCreateInput

// NewPaymentMethodCreateInput instantiates a new PaymentMethodCreateInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentMethodCreateInput(name string, nonce string) *PaymentMethodCreateInput {
	this := PaymentMethodCreateInput{}
	this.Name = name
	this.Nonce = nonce
	return &this
}

// NewPaymentMethodCreateInputWithDefaults instantiates a new PaymentMethodCreateInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentMethodCreateInputWithDefaults() *PaymentMethodCreateInput {
	this := PaymentMethodCreateInput{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PaymentMethodCreateInput) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateInput) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PaymentMethodCreateInput) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PaymentMethodCreateInput) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value
func (o *PaymentMethodCreateInput) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateInput) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentMethodCreateInput) SetName(v string) {
	o.Name = v
}

// GetNonce returns the Nonce field value
func (o *PaymentMethodCreateInput) GetNonce() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *PaymentMethodCreateInput) GetNonceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *PaymentMethodCreateInput) SetNonce(v string) {
	o.Nonce = v
}

func (o PaymentMethodCreateInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentMethodCreateInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["name"] = o.Name
	toSerialize["nonce"] = o.Nonce

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaymentMethodCreateInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"nonce",
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

	varPaymentMethodCreateInput := _PaymentMethodCreateInput{}

	err = json.Unmarshal(data, &varPaymentMethodCreateInput)

	if err != nil {
		return err
	}

	*o = PaymentMethodCreateInput(varPaymentMethodCreateInput)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "name")
		delete(additionalProperties, "nonce")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentMethodCreateInput struct {
	value *PaymentMethodCreateInput
	isSet bool
}

func (v NullablePaymentMethodCreateInput) Get() *PaymentMethodCreateInput {
	return v.value
}

func (v *NullablePaymentMethodCreateInput) Set(val *PaymentMethodCreateInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentMethodCreateInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentMethodCreateInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentMethodCreateInput(val *PaymentMethodCreateInput) *NullablePaymentMethodCreateInput {
	return &NullablePaymentMethodCreateInput{value: val, isSet: true}
}

func (v NullablePaymentMethodCreateInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentMethodCreateInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
