# CloudRouterActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CloudRouterActionType**](CloudRouterActionType.md) |  | 
**Uuid** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**State** | [**CloudRouterActionState**](CloudRouterActionState.md) |  | 
**ChangeLog** | [**Changelog**](Changelog.md) |  | 

## Methods

### NewCloudRouterActionResponse

`func NewCloudRouterActionResponse(type_ CloudRouterActionType, uuid string, state CloudRouterActionState, changeLog Changelog, ) *CloudRouterActionResponse`

NewCloudRouterActionResponse instantiates a new CloudRouterActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionResponseWithDefaults

`func NewCloudRouterActionResponseWithDefaults() *CloudRouterActionResponse`

NewCloudRouterActionResponseWithDefaults instantiates a new CloudRouterActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CloudRouterActionResponse) GetType() CloudRouterActionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudRouterActionResponse) GetTypeOk() (*CloudRouterActionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudRouterActionResponse) SetType(v CloudRouterActionType)`

SetType sets Type field to given value.


### GetUuid

`func (o *CloudRouterActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CloudRouterActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CloudRouterActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDescription

`func (o *CloudRouterActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudRouterActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudRouterActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudRouterActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *CloudRouterActionResponse) GetState() CloudRouterActionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudRouterActionResponse) GetStateOk() (*CloudRouterActionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudRouterActionResponse) SetState(v CloudRouterActionState)`

SetState sets State field to given value.


### GetChangeLog

`func (o *CloudRouterActionResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CloudRouterActionResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CloudRouterActionResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


