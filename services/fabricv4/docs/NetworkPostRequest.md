# NetworkPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NetworkType**](NetworkType.md) |  | 
**Name** | **string** | Customer-provided network name | 
**Scope** | [**NetworkScope**](NetworkScope.md) |  | 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**Notifications** | [**[]SimplifiedNotification**](SimplifiedNotification.md) | Preferences for notifications on network configuration or status changes | 

## Methods

### NewNetworkPostRequest

`func NewNetworkPostRequest(type_ NetworkType, name string, scope NetworkScope, notifications []SimplifiedNotification, ) *NetworkPostRequest`

NewNetworkPostRequest instantiates a new NetworkPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPostRequestWithDefaults

`func NewNetworkPostRequestWithDefaults() *NetworkPostRequest`

NewNetworkPostRequestWithDefaults instantiates a new NetworkPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkPostRequest) GetType() NetworkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPostRequest) GetTypeOk() (*NetworkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPostRequest) SetType(v NetworkType)`

SetType sets Type field to given value.


### GetName

`func (o *NetworkPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *NetworkPostRequest) GetScope() NetworkScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *NetworkPostRequest) GetScopeOk() (*NetworkScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *NetworkPostRequest) SetScope(v NetworkScope)`

SetScope sets Scope field to given value.


### GetLocation

`func (o *NetworkPostRequest) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkPostRequest) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkPostRequest) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NetworkPostRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetProject

`func (o *NetworkPostRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NetworkPostRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NetworkPostRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *NetworkPostRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetNotifications

`func (o *NetworkPostRequest) GetNotifications() []SimplifiedNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *NetworkPostRequest) GetNotificationsOk() (*[]SimplifiedNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *NetworkPostRequest) SetNotifications(v []SimplifiedNotification)`

SetNotifications sets Notifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


