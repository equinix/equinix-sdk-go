# ProjectCreateFromRootInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | **string** | The name of the project. Cannot contain characters encoded in greater than 3 bytes such as emojis. | 
**OrganizationId** | Pointer to **string** |  | [optional] 
**PaymentMethodId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ProjectCreateFromRootInputType**](ProjectCreateFromRootInputType.md) |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProjectCreateFromRootInput

`func NewProjectCreateFromRootInput(name string, ) *ProjectCreateFromRootInput`

NewProjectCreateFromRootInput instantiates a new ProjectCreateFromRootInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateFromRootInputWithDefaults

`func NewProjectCreateFromRootInputWithDefaults() *ProjectCreateFromRootInput`

NewProjectCreateFromRootInputWithDefaults instantiates a new ProjectCreateFromRootInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomdata

`func (o *ProjectCreateFromRootInput) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *ProjectCreateFromRootInput) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *ProjectCreateFromRootInput) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *ProjectCreateFromRootInput) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetName

`func (o *ProjectCreateFromRootInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectCreateFromRootInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectCreateFromRootInput) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *ProjectCreateFromRootInput) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ProjectCreateFromRootInput) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ProjectCreateFromRootInput) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ProjectCreateFromRootInput) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetPaymentMethodId

`func (o *ProjectCreateFromRootInput) GetPaymentMethodId() string`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *ProjectCreateFromRootInput) GetPaymentMethodIdOk() (*string, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *ProjectCreateFromRootInput) SetPaymentMethodId(v string)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *ProjectCreateFromRootInput) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetType

`func (o *ProjectCreateFromRootInput) GetType() ProjectCreateFromRootInputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProjectCreateFromRootInput) GetTypeOk() (*ProjectCreateFromRootInputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProjectCreateFromRootInput) SetType(v ProjectCreateFromRootInputType)`

SetType sets Type field to given value.

### HasType

`func (o *ProjectCreateFromRootInput) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTags

`func (o *ProjectCreateFromRootInput) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ProjectCreateFromRootInput) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ProjectCreateFromRootInput) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ProjectCreateFromRootInput) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


