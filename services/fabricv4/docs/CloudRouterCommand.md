# CloudRouterCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**CloudRouterCommandType**](CloudRouterCommandType.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | Customer-provided Cloud Router name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**CloudRouterCommandState**](CloudRouterCommandState.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Request** | Pointer to [**CloudRouterCommandRequest**](CloudRouterCommandRequest.md) |  | [optional] 
**Response** | Pointer to [**CloudRouterCommandResponse**](CloudRouterCommandResponse.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewCloudRouterCommand

`func NewCloudRouterCommand() *CloudRouterCommand`

NewCloudRouterCommand instantiates a new CloudRouterCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterCommandWithDefaults

`func NewCloudRouterCommandWithDefaults() *CloudRouterCommand`

NewCloudRouterCommandWithDefaults instantiates a new CloudRouterCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CloudRouterCommand) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CloudRouterCommand) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CloudRouterCommand) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CloudRouterCommand) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *CloudRouterCommand) GetType() CloudRouterCommandType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterCommand) GetTypeOk() (*CloudRouterCommandType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterCommand) SetType(v CloudRouterCommandType)`

SetType sets Type field to given value.

### HasType

`func (o *CloudRouterCommand) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *CloudRouterCommand) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudRouterCommand) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudRouterCommand) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CloudRouterCommand) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *CloudRouterCommand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudRouterCommand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudRouterCommand) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudRouterCommand) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CloudRouterCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRouterCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRouterCommand) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRouterCommand) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *CloudRouterCommand) GetState() CloudRouterCommandState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudRouterCommand) GetStateOk() (*CloudRouterCommandState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudRouterCommand) SetState(v CloudRouterCommandState)`

SetState sets State field to given value.

### HasState

`func (o *CloudRouterCommand) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *CloudRouterCommand) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *CloudRouterCommand) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *CloudRouterCommand) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *CloudRouterCommand) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRequest

`func (o *CloudRouterCommand) GetRequest() CloudRouterCommandRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CloudRouterCommand) GetRequestOk() (*CloudRouterCommandRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CloudRouterCommand) SetRequest(v CloudRouterCommandRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *CloudRouterCommand) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *CloudRouterCommand) GetResponse() CloudRouterCommandResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *CloudRouterCommand) GetResponseOk() (*CloudRouterCommandResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *CloudRouterCommand) SetResponse(v CloudRouterCommandResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *CloudRouterCommand) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetChangeLog

`func (o *CloudRouterCommand) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CloudRouterCommand) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CloudRouterCommand) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CloudRouterCommand) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


