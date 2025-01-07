# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryName** | Pointer to **string** |  | [optional] 
**Templates** | Pointer to [**[]Template**](Template.md) |  | [optional] 

## Methods

### NewCategory

`func NewCategory() *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryName

`func (o *Category) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *Category) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *Category) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *Category) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetTemplates

`func (o *Category) GetTemplates() []Template`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Category) GetTemplatesOk() (*[]Template, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *Category) SetTemplates(v []Template)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *Category) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


