/*
Equinix Internet Access API

Equinix Internet Access provides direct access to the Internet with scalable bandwidth options in IBX data centers.

Contact: api-support@equinix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eiav2

import (
	"encoding/json"
	"fmt"
)

// checks if the OrderSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderSignature{}

// OrderSignature struct for OrderSignature
type OrderSignature struct {
	Signatory            OrderSignatureSignatory `json:"signatory"`
	Delegate             *OrderSignatureDelegate `json:"delegate,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrderSignature OrderSignature

// NewOrderSignature instantiates a new OrderSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderSignature(signatory OrderSignatureSignatory) *OrderSignature {
	this := OrderSignature{}
	this.Signatory = signatory
	return &this
}

// NewOrderSignatureWithDefaults instantiates a new OrderSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderSignatureWithDefaults() *OrderSignature {
	this := OrderSignature{}
	return &this
}

// GetSignatory returns the Signatory field value
func (o *OrderSignature) GetSignatory() OrderSignatureSignatory {
	if o == nil {
		var ret OrderSignatureSignatory
		return ret
	}

	return o.Signatory
}

// GetSignatoryOk returns a tuple with the Signatory field value
// and a boolean to check if the value has been set.
func (o *OrderSignature) GetSignatoryOk() (*OrderSignatureSignatory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signatory, true
}

// SetSignatory sets field value
func (o *OrderSignature) SetSignatory(v OrderSignatureSignatory) {
	o.Signatory = v
}

// GetDelegate returns the Delegate field value if set, zero value otherwise.
func (o *OrderSignature) GetDelegate() OrderSignatureDelegate {
	if o == nil || IsNil(o.Delegate) {
		var ret OrderSignatureDelegate
		return ret
	}
	return *o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderSignature) GetDelegateOk() (*OrderSignatureDelegate, bool) {
	if o == nil || IsNil(o.Delegate) {
		return nil, false
	}
	return o.Delegate, true
}

// HasDelegate returns a boolean if a field has been set.
func (o *OrderSignature) HasDelegate() bool {
	if o != nil && !IsNil(o.Delegate) {
		return true
	}

	return false
}

// SetDelegate gets a reference to the given OrderSignatureDelegate and assigns it to the Delegate field.
func (o *OrderSignature) SetDelegate(v OrderSignatureDelegate) {
	o.Delegate = &v
}

func (o OrderSignature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signatory"] = o.Signatory
	if !IsNil(o.Delegate) {
		toSerialize["delegate"] = o.Delegate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrderSignature) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"signatory",
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

	varOrderSignature := _OrderSignature{}

	err = json.Unmarshal(data, &varOrderSignature)

	if err != nil {
		return err
	}

	*o = OrderSignature(varOrderSignature)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signatory")
		delete(additionalProperties, "delegate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrderSignature struct {
	value *OrderSignature
	isSet bool
}

func (v NullableOrderSignature) Get() *OrderSignature {
	return v.value
}

func (v *NullableOrderSignature) Set(val *OrderSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderSignature(val *OrderSignature) *NullableOrderSignature {
	return &NullableOrderSignature{value: val, isSet: true}
}

func (v NullableOrderSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
