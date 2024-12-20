/*
Equinix Fabric API v4

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fabricv4

import (
	"encoding/json"
	"fmt"
)

// checks if the MetroError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetroError{}

// MetroError Error with details
type MetroError struct {
	ErrorCode            MetroErrorErrorCode        `json:"errorCode"`
	ErrorMessage         MetroErrorErrorMessage     `json:"errorMessage"`
	CorrelationId        *string                    `json:"correlationId,omitempty"`
	Details              *string                    `json:"details,omitempty"`
	Help                 *string                    `json:"help,omitempty"`
	AdditionalInfo       []PriceErrorAdditionalInfo `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MetroError MetroError

// NewMetroError instantiates a new MetroError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetroError(errorCode MetroErrorErrorCode, errorMessage MetroErrorErrorMessage) *MetroError {
	this := MetroError{}
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	return &this
}

// NewMetroErrorWithDefaults instantiates a new MetroError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetroErrorWithDefaults() *MetroError {
	this := MetroError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *MetroError) GetErrorCode() MetroErrorErrorCode {
	if o == nil {
		var ret MetroErrorErrorCode
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *MetroError) GetErrorCodeOk() (*MetroErrorErrorCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *MetroError) SetErrorCode(v MetroErrorErrorCode) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *MetroError) GetErrorMessage() MetroErrorErrorMessage {
	if o == nil {
		var ret MetroErrorErrorMessage
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *MetroError) GetErrorMessageOk() (*MetroErrorErrorMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *MetroError) SetErrorMessage(v MetroErrorErrorMessage) {
	o.ErrorMessage = v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *MetroError) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroError) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *MetroError) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *MetroError) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MetroError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MetroError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *MetroError) SetDetails(v string) {
	o.Details = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *MetroError) GetHelp() string {
	if o == nil || IsNil(o.Help) {
		var ret string
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroError) GetHelpOk() (*string, bool) {
	if o == nil || IsNil(o.Help) {
		return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *MetroError) HasHelp() bool {
	if o != nil && !IsNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given string and assigns it to the Help field.
func (o *MetroError) SetHelp(v string) {
	o.Help = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *MetroError) GetAdditionalInfo() []PriceErrorAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []PriceErrorAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetroError) GetAdditionalInfoOk() ([]PriceErrorAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *MetroError) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []PriceErrorAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *MetroError) SetAdditionalInfo(v []PriceErrorAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o MetroError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetroError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["errorCode"] = o.ErrorCode
	toSerialize["errorMessage"] = o.ErrorMessage
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Help) {
		toSerialize["help"] = o.Help
	}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MetroError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"errorCode",
		"errorMessage",
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

	varMetroError := _MetroError{}

	err = json.Unmarshal(data, &varMetroError)

	if err != nil {
		return err
	}

	*o = MetroError(varMetroError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "errorCode")
		delete(additionalProperties, "errorMessage")
		delete(additionalProperties, "correlationId")
		delete(additionalProperties, "details")
		delete(additionalProperties, "help")
		delete(additionalProperties, "additionalInfo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMetroError struct {
	value *MetroError
	isSet bool
}

func (v NullableMetroError) Get() *MetroError {
	return v.value
}

func (v *NullableMetroError) Set(val *MetroError) {
	v.value = val
	v.isSet = true
}

func (v NullableMetroError) IsSet() bool {
	return v.isSet
}

func (v *NullableMetroError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetroError(val *MetroError) *NullableMetroError {
	return &NullableMetroError{value: val, isSet: true}
}

func (v NullableMetroError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetroError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
