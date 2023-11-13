# ProjectCreateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis. | 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ProjectCreateFromRootInputType**](ProjectCreateFromRootInputType.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectCreateInput

`func NewProjectCreateInput(name string, ) *ProjectCreateInput`

NewProjectCreateInput instantiates a new ProjectCreateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateInputWithDefaults

`func NewProjectCreateInputWithDefaults() *ProjectCreateInput`

NewProjectCreateInputWithDefaults instantiates a new ProjectCreateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomdata

`func (o *ProjectCreateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *ProjectCreateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *ProjectCreateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *ProjectCreateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *ProjectCreateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreateInput) SetName(v string)`

SetName sets Name field to given value.


### GetPaymentMethodId

`func (o *ProjectCreateInput) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ProjectCreateInput) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ProjectCreateInput) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *ProjectCreateInput) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *ProjectCreateInput) GetType() ProjectCreateFromRootInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectCreateInput) GetTypeOk() (*ProjectCreateFromRootInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectCreateInput) SetType(v ProjectCreateFromRootInputType)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectCreateInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *ProjectCreateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectCreateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectCreateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectCreateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


