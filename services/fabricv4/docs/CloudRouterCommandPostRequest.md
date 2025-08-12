# CloudRouterCommandPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CloudRouterCommandType**](CloudRouterCommandType.md) |  | 
**Name** | Pointer to **string** | Customer-provided Cloud Router Command name | [optional] 
**Description** | Pointer to **string** | Customer-provided Cloud Router Command description | [optional] 
**Project** | [**Project**](Project.md) |  | 
**Request** | [**CloudRouterCommandRequestPayload**](CloudRouterCommandRequestPayload.md) |  | 

## Methods

### NewCloudRouterCommandPostRequest

`func NewCloudRouterCommandPostRequest(type_ CloudRouterCommandType, project Project, request CloudRouterCommandRequestPayload, ) *CloudRouterCommandPostRequest`

NewCloudRouterCommandPostRequest instantiates a new CloudRouterCommandPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandPostRequestWithDefaults

`func NewCloudRouterCommandPostRequestWithDefaults() *CloudRouterCommandPostRequest`

NewCloudRouterCommandPostRequestWithDefaults instantiates a new CloudRouterCommandPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterCommandPostRequest) GetType() CloudRouterCommandType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterCommandPostRequest) GetTypeOk() (*CloudRouterCommandType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterCommandPostRequest) SetType(v CloudRouterCommandType)`

SetType sets Type field to given value.


### GetName

`func (o *CloudRouterCommandPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouterCommandPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouterCommandPostRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouterCommandPostRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CloudRouterCommandPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRouterCommandPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRouterCommandPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRouterCommandPostRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouterCommandPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouterCommandPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouterCommandPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.


### GetRequest

`func (o *CloudRouterCommandPostRequest) GetRequest() CloudRouterCommandRequestPayload`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CloudRouterCommandPostRequest) GetRequestOk() (*CloudRouterCommandRequestPayload, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CloudRouterCommandPostRequest) SetRequest(v CloudRouterCommandRequestPayload)`

SetRequest sets Request field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


