# ProcessStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Service profile custom step title | [optional] 
**SubTitle** | Pointer to **string** | Service profile custom step sub title | [optional] 
**Description** | Pointer to **string** | Service profile custom step description | [optional] 

## Methods

### NewProcessStep

`func NewProcessStep() *ProcessStep`

NewProcessStep instantiates a new ProcessStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessStepWithDefaults

`func NewProcessStepWithDefaults() *ProcessStep`

NewProcessStepWithDefaults instantiates a new ProcessStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ProcessStep) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ProcessStep) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ProcessStep) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ProcessStep) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSubTitle

`func (o *ProcessStep) GetSubTitle() string`

GetSubTitle returns the SubTitle field if non-nil, zero value otherwise.

### GetSubTitleOk

`func (o *ProcessStep) GetSubTitleOk() (*string, bool)`

GetSubTitleOk returns a tuple with the SubTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitle

`func (o *ProcessStep) SetSubTitle(v string)`

SetSubTitle sets SubTitle field to given value.

### HasSubTitle

`func (o *ProcessStep) HasSubTitle() bool`

HasSubTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ProcessStep) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessStep) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessStep) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProcessStep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


