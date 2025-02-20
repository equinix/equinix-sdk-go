# AssetsListPayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]Category**](Category.md) |  | [optional] 
**Classification** | Pointer to **string** | asset classification for the Electrical and Mechanical | [optional] 

## Methods

### NewAssetsListPayLoad

`func NewAssetsListPayLoad() *AssetsListPayLoad`

NewAssetsListPayLoad instantiates a new AssetsListPayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsListPayLoadWithDefaults

`func NewAssetsListPayLoadWithDefaults() *AssetsListPayLoad`

NewAssetsListPayLoadWithDefaults instantiates a new AssetsListPayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *AssetsListPayLoad) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AssetsListPayLoad) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AssetsListPayLoad) SetCategories(v []Category)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AssetsListPayLoad) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetClassification

`func (o *AssetsListPayLoad) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *AssetsListPayLoad) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *AssetsListPayLoad) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *AssetsListPayLoad) HasClassification() bool`

HasClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


