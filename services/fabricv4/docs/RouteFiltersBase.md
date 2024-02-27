# RouteFiltersBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**ConnectionRouteFilterDataType**](ConnectionRouteFilterDataType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Notifications** | Pointer to [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on route filter configuration or status changes | [optional] 

## Methods

### NewRouteFiltersBase

`func NewRouteFiltersBase() *RouteFiltersBase`

NewRouteFiltersBase instantiates a new RouteFiltersBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersBaseWithDefaults

`func NewRouteFiltersBaseWithDefaults() *RouteFiltersBase`

NewRouteFiltersBaseWithDefaults instantiates a new RouteFiltersBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RouteFiltersBase) GetType() ConnectionRouteFilterDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFiltersBase) GetTypeOk() (*ConnectionRouteFilterDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFiltersBase) SetType(v ConnectionRouteFilterDataType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteFiltersBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *RouteFiltersBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFiltersBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFiltersBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteFiltersBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteFiltersBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFiltersBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFiltersBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFiltersBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *RouteFiltersBase) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteFiltersBase) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteFiltersBase) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *RouteFiltersBase) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNotifications

`func (o *RouteFiltersBase) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *RouteFiltersBase) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *RouteFiltersBase) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *RouteFiltersBase) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


