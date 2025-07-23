# ValidationErrorAdditionalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Property** | Pointer to **string** | Property at which the validation violation occurs. Can be empty when validation happens at the object top level. | [optional] 
**Reason** | **string** | Fallback error message in plain English. | 
**ValidationRuleTag** | **string** | Violated validation rule tag. | 
**ValidationRoot** | Pointer to **string** | Name of the object at which the validation violation occurs. Can be empty when validation happens inline at the request params. | [optional] 
**ValidationParameters** | Pointer to **map[string]string** | A map containing additional violation parameters, these can be used to render localized error messages. The list of available parameters is different for each validationRuleTag. | [optional] 

## Methods

### NewValidationErrorAdditionalInfo

`func NewValidationErrorAdditionalInfo(reason string, validationRuleTag string, ) *ValidationErrorAdditionalInfo`

NewValidationErrorAdditionalInfo instantiates a new ValidationErrorAdditionalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorAdditionalInfoWithDefaults

`func NewValidationErrorAdditionalInfoWithDefaults() *ValidationErrorAdditionalInfo`

NewValidationErrorAdditionalInfoWithDefaults instantiates a new ValidationErrorAdditionalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperty

`func (o *ValidationErrorAdditionalInfo) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ValidationErrorAdditionalInfo) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ValidationErrorAdditionalInfo) SetProperty(v string)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ValidationErrorAdditionalInfo) HasProperty() bool`

HasProperty returns a boolean if a field has been set.

### GetReason

`func (o *ValidationErrorAdditionalInfo) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ValidationErrorAdditionalInfo) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ValidationErrorAdditionalInfo) SetReason(v string)`

SetReason sets Reason field to given value.


### GetValidationRuleTag

`func (o *ValidationErrorAdditionalInfo) GetValidationRuleTag() string`

GetValidationRuleTag returns the ValidationRuleTag field if non-nil, zero value otherwise.

### GetValidationRuleTagOk

`func (o *ValidationErrorAdditionalInfo) GetValidationRuleTagOk() (*string, bool)`

GetValidationRuleTagOk returns a tuple with the ValidationRuleTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleTag

`func (o *ValidationErrorAdditionalInfo) SetValidationRuleTag(v string)`

SetValidationRuleTag sets ValidationRuleTag field to given value.


### GetValidationRoot

`func (o *ValidationErrorAdditionalInfo) GetValidationRoot() string`

GetValidationRoot returns the ValidationRoot field if non-nil, zero value otherwise.

### GetValidationRootOk

`func (o *ValidationErrorAdditionalInfo) GetValidationRootOk() (*string, bool)`

GetValidationRootOk returns a tuple with the ValidationRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRoot

`func (o *ValidationErrorAdditionalInfo) SetValidationRoot(v string)`

SetValidationRoot sets ValidationRoot field to given value.

### HasValidationRoot

`func (o *ValidationErrorAdditionalInfo) HasValidationRoot() bool`

HasValidationRoot returns a boolean if a field has been set.

### GetValidationParameters

`func (o *ValidationErrorAdditionalInfo) GetValidationParameters() map[string]string`

GetValidationParameters returns the ValidationParameters field if non-nil, zero value otherwise.

### GetValidationParametersOk

`func (o *ValidationErrorAdditionalInfo) GetValidationParametersOk() (*map[string]string, bool)`

GetValidationParametersOk returns a tuple with the ValidationParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationParameters

`func (o *ValidationErrorAdditionalInfo) SetValidationParameters(v map[string]string)`

SetValidationParameters sets ValidationParameters field to given value.

### HasValidationParameters

`func (o *ValidationErrorAdditionalInfo) HasValidationParameters() bool`

HasValidationParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


