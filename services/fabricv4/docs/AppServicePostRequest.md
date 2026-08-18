# AppServicePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AppServiceType**](AppServiceType.md) |  | [default to APPSERVICETYPE_APP_SERVICE]
**Name** | **string** | Customer-provided App Service name | 
**Description** | Pointer to **string** | Customer-provided App Service description | [optional] 
**Endpoint** | **string** | Accessible endpoint through this service | 
**SourceDomains** | Pointer to **[]string** | List of source domains from where traffic is allowed | [optional] 
**Project** | [**Project**](Project.md) |  | 

## Methods

### NewAppServicePostRequest

`func NewAppServicePostRequest(type_ AppServiceType, name string, endpoint string, project Project, ) *AppServicePostRequest`

NewAppServicePostRequest instantiates a new AppServicePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServicePostRequestWithDefaults

`func NewAppServicePostRequestWithDefaults() *AppServicePostRequest`

NewAppServicePostRequestWithDefaults instantiates a new AppServicePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppServicePostRequest) GetType() AppServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppServicePostRequest) GetTypeOk() (*AppServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppServicePostRequest) SetType(v AppServiceType)`

SetType sets Type field to given value.


### GetName

`func (o *AppServicePostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppServicePostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppServicePostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppServicePostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppServicePostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppServicePostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppServicePostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *AppServicePostRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AppServicePostRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AppServicePostRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetSourceDomains

`func (o *AppServicePostRequest) GetSourceDomains() []string`

GetSourceDomains returns the SourceDomains field if non-nil, zero value otherwise.

### GetSourceDomainsOk

`func (o *AppServicePostRequest) GetSourceDomainsOk() (*[]string, bool)`

GetSourceDomainsOk returns a tuple with the SourceDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDomains

`func (o *AppServicePostRequest) SetSourceDomains(v []string)`

SetSourceDomains sets SourceDomains field to given value.

### HasSourceDomains

`func (o *AppServicePostRequest) HasSourceDomains() bool`

HasSourceDomains returns a boolean if a field has been set.

### GetProject

`func (o *AppServicePostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppServicePostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppServicePostRequest) SetProject(v Project)`

SetProject sets Project field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


