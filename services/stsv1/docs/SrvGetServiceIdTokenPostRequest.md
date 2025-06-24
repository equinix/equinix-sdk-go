# SrvGetServiceIdTokenPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | Globally unique identifier of a project. | 
**Audiences** | **[]string** | array of case-sensitive audiences to include in the &#x60;aud&#x60; claim of the ID token | 

## Methods

### NewSrvGetServiceIdTokenPostRequest

`func NewSrvGetServiceIdTokenPostRequest(projectId string, audiences []string, ) *SrvGetServiceIdTokenPostRequest`

NewSrvGetServiceIdTokenPostRequest instantiates a new SrvGetServiceIdTokenPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSrvGetServiceIdTokenPostRequestWithDefaults

`func NewSrvGetServiceIdTokenPostRequestWithDefaults() *SrvGetServiceIdTokenPostRequest`

NewSrvGetServiceIdTokenPostRequestWithDefaults instantiates a new SrvGetServiceIdTokenPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SrvGetServiceIdTokenPostRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SrvGetServiceIdTokenPostRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SrvGetServiceIdTokenPostRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetAudiences

`func (o *SrvGetServiceIdTokenPostRequest) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *SrvGetServiceIdTokenPostRequest) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *SrvGetServiceIdTokenPostRequest) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


