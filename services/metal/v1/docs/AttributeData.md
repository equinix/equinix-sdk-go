# AttributeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latest** | Pointer to **bool** | Boolean flag to know if the firmware set is the latest for the model and vendor | [optional] [readonly] 
**Model** | Pointer to **string** | Model on which this firmware set can be applied | [optional] [readonly] 
**Vendor** | Pointer to **string** | Vendor on which this firmware set can be applied | [optional] [readonly] 
**Plan** | Pointer to **string** | Plan where the firmware set can be applied | [optional] [readonly] 

## Methods

### NewAttributeData

`func NewAttributeData() *AttributeData`

NewAttributeData instantiates a new AttributeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributeDataWithDefaults

`func NewAttributeDataWithDefaults() *AttributeData`

NewAttributeDataWithDefaults instantiates a new AttributeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatest

`func (o *AttributeData) GetLatest() bool`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *AttributeData) GetLatestOk() (*bool, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *AttributeData) SetLatest(v bool)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *AttributeData) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetModel

`func (o *AttributeData) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AttributeData) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AttributeData) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AttributeData) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetVendor

`func (o *AttributeData) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AttributeData) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AttributeData) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AttributeData) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetPlan

`func (o *AttributeData) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AttributeData) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AttributeData) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AttributeData) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


