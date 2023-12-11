# ProjectUpdateInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackendTransferEnabled** | Pointer to **bool** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** | The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis. | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectUpdateInput

`func NewProjectUpdateInput() *ProjectUpdateInput`

NewProjectUpdateInput instantiates a new ProjectUpdateInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateInputWithDefaults

`func NewProjectUpdateInputWithDefaults() *ProjectUpdateInput`

NewProjectUpdateInputWithDefaults instantiates a new ProjectUpdateInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackendTransferEnabled

`func (o *ProjectUpdateInput) GetBackendTransferEnabled() bool`

GetBackendTransferEnabled returns the BackendTransferEnabled field if non-nil, zero value otherwise.

### GetBackendTransferEnabledOk

`func (o *ProjectUpdateInput) GetBackendTransferEnabledOk() (*bool, bool)`

GetBackendTransferEnabledOk returns a tuple with the BackendTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendTransferEnabled

`func (o *ProjectUpdateInput) SetBackendTransferEnabled(v bool)`

SetBackendTransferEnabled sets BackendTransferEnabled field to given value.

### HasBackendTransferEnabled

`func (o *ProjectUpdateInput) HasBackendTransferEnabled() bool`

HasBackendTransferEnabled returns a boolean if a field has been set.

### GetCustomdata

`func (o *ProjectUpdateInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *ProjectUpdateInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *ProjectUpdateInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *ProjectUpdateInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *ProjectUpdateInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectUpdateInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectUpdateInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectUpdateInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *ProjectUpdateInput) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ProjectUpdateInput) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ProjectUpdateInput) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *ProjectUpdateInput) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetTags

`func (o *ProjectUpdateInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectUpdateInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectUpdateInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectUpdateInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


