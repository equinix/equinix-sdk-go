# AppDomainPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppDomainType**](AppDomainType.md) |  | [default to APPDOMAINTYPE_APP_DOMAIN]
**Name** | **string** | Customer-provided App Domain name | 
**Description** | Pointer to **string** | Customer-provided App Domain description | [optional] 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewAppDomainPostRequest

`func NewAppDomainPostRequest(type_ AppDomainType, name string, project Project, ) *AppDomainPostRequest`

NewAppDomainPostRequest instantiates a new AppDomainPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDomainPostRequestWithDefaults

`func NewAppDomainPostRequestWithDefaults() *AppDomainPostRequest`

NewAppDomainPostRequestWithDefaults instantiates a new AppDomainPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppDomainPostRequest) GetType() AppDomainType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppDomainPostRequest) GetTypeOk() (*AppDomainType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppDomainPostRequest) SetType(v AppDomainType)`

SetType sets Type field to given value.


### GetName

`func (o *AppDomainPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDomainPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDomainPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppDomainPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppDomainPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppDomainPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppDomainPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *AppDomainPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppDomainPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppDomainPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


