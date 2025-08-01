/*
Secure Cabinet API

Secure Cabinet API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package securecabinetv1

import (
	"encoding/json"
	"fmt"
)

// checks if the Error type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Error{}

// Error Error responses are included with 4xx and 5xx HTTP responses from the API service. Either \"error\" or \"errors\" will be set.
type Error struct {
	// Application error code. Format :EQ-<error code>
	ErrorCode string `json:"errorCode" validate:"regexp=^EQ-70[0-9]{5}$"`
	// Application error message
	ErrorMessage string `json:"errorMessage"`
	// Correlation ID without any sensitive or meaningful information.
	CorrelationId *string `json:"correlationId,omitempty"`
	// More information on errors.
	Details *string `json:"details,omitempty"`
	// Help message or URL to a help page for the corresponding errorCode.
	Help *string `json:"help,omitempty"`
	// Contains application specific information for this error. The object inside this array can have any number of application specific properties.
	AdditionalInfo       []ErrorAdditionalInfo `json:"additionalInfo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Error Error

// NewError instantiates a new Error object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewError(errorCode string, errorMessage string) *Error {
	this := Error{}
	this.ErrorCode = errorCode
	this.ErrorMessage = errorMessage
	return &this
}

// NewErrorWithDefaults instantiates a new Error object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorWithDefaults() *Error {
	this := Error{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *Error) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *Error) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *Error) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value
func (o *Error) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *Error) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *Error) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *Error) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *Error) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *Error) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *Error) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *Error) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *Error) SetDetails(v string) {
	o.Details = &v
}

// GetHelp returns the Help field value if set, zero value otherwise.
func (o *Error) GetHelp() string {
	if o == nil || IsNil(o.Help) {
		var ret string
		return ret
	}
	return *o.Help
}

// GetHelpOk returns a tuple with the Help field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetHelpOk() (*string, bool) {
	if o == nil || IsNil(o.Help) {
		return nil, false
	}
	return o.Help, true
}

// HasHelp returns a boolean if a field has been set.
func (o *Error) HasHelp() bool {
	if o != nil && !IsNil(o.Help) {
		return true
	}

	return false
}

// SetHelp gets a reference to the given string and assigns it to the Help field.
func (o *Error) SetHelp(v string) {
	o.Help = &v
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *Error) GetAdditionalInfo() []ErrorAdditionalInfo {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret []ErrorAdditionalInfo
		return ret
	}
	return o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Error) GetAdditionalInfoOk() ([]ErrorAdditionalInfo, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *Error) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given []ErrorAdditionalInfo and assigns it to the AdditionalInfo field.
func (o *Error) SetAdditionalInfo(v []ErrorAdditionalInfo) {
	o.AdditionalInfo = v
}

func (o Error) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Error) ToMap() (map[string]interface{}, error) {
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

func (o *Error) UnmarshalJSON(data []byte) (err error) {
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

	varError := _Error{}

	err = json.Unmarshal(data, &varError)

	if err != nil {
		return err
	}

	*o = Error(varError)

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

type NullableError struct {
	value *Error
	isSet bool
}

func (v NullableError) Get() *Error {
	return v.value
}

func (v *NullableError) Set(val *Error) {
	v.value = val
	v.isSet = true
}

func (v NullableError) IsSet() bool {
	return v.isSet
}

func (v *NullableError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableError(val *Error) *NullableError {
	return &NullableError{value: val, isSet: true}
}

func (v NullableError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
