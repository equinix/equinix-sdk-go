# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Description** | **string** |  | 
**Required** | Pointer to **bool** |  | [optional] 
**DataType** | [**CustomFieldDataType**](CustomFieldDataType.md) |  | 
**Options** | Pointer to **[]string** |  | [optional] 
**CaptureInEmail** | Pointer to **bool** | capture this field as a part of email notification | [optional] 

## Methods

### NewCustomField

`func NewCustomField(label string, description string, dataType CustomFieldDataType, ) *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *CustomField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CustomField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CustomField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *CustomField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomField) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRequired

`func (o *CustomField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *CustomField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *CustomField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *CustomField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *CustomField) GetDataType() CustomFieldDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CustomField) GetDataTypeOk() (*CustomFieldDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CustomField) SetDataType(v CustomFieldDataType)`

SetDataType sets DataType field to given value.


### GetOptions

`func (o *CustomField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CustomField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CustomField) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CustomField) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetCaptureInEmail

`func (o *CustomField) GetCaptureInEmail() bool`

GetCaptureInEmail returns the CaptureInEmail field if non-nil, zero value otherwise.

### GetCaptureInEmailOk

`func (o *CustomField) GetCaptureInEmailOk() (*bool, bool)`

GetCaptureInEmailOk returns a tuple with the CaptureInEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureInEmail

`func (o *CustomField) SetCaptureInEmail(v bool)`

SetCaptureInEmail sets CaptureInEmail field to given value.

### HasCaptureInEmail

`func (o *CustomField) HasCaptureInEmail() bool`

HasCaptureInEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


