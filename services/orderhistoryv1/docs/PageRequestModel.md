# PageRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **int32** | Page number indexed from 0. | [optional] 
**Size** | Pointer to **int32** | Page Size. | [optional] 

## Methods

### NewPageRequestModel

`func NewPageRequestModel() *PageRequestModel`

NewPageRequestModel instantiates a new PageRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRequestModelWithDefaults

`func NewPageRequestModelWithDefaults() *PageRequestModel`

NewPageRequestModelWithDefaults instantiates a new PageRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *PageRequestModel) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PageRequestModel) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PageRequestModel) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *PageRequestModel) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetSize

`func (o *PageRequestModel) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageRequestModel) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageRequestModel) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PageRequestModel) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


