# AppSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | [**AppSubscriptionType**](AppSubscriptionType.md) |  | [default to APPSUBSCRIPTIONTYPE_APP_SUBSCRIPTION]
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Source** | [**AppSubscriptionSource**](AppSubscriptionSource.md) |  | 
**Target** | [**AppSubscriptionTarget**](AppSubscriptionTarget.md) |  | 
**State** | Pointer to [**AppSubscriptionState**](AppSubscriptionState.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**Change** | Pointer to [**AppSubscriptionChange**](AppSubscriptionChange.md) |  | [optional] 

## Methods

### NewAppSubscription

`func NewAppSubscription(type_ AppSubscriptionType, source AppSubscriptionSource, target AppSubscriptionTarget, project Project, ) *AppSubscription`

NewAppSubscription instantiates a new AppSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppSubscriptionWithDefaults

`func NewAppSubscriptionWithDefaults() *AppSubscription`

NewAppSubscriptionWithDefaults instantiates a new AppSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppSubscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppSubscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppSubscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppSubscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppSubscription) GetType() AppSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppSubscription) GetTypeOk() (*AppSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppSubscription) SetType(v AppSubscriptionType)`

SetType sets Type field to given value.


### GetUuid

`func (o *AppSubscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppSubscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppSubscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppSubscription) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSource

`func (o *AppSubscription) GetSource() AppSubscriptionSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AppSubscription) GetSourceOk() (*AppSubscriptionSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AppSubscription) SetSource(v AppSubscriptionSource)`

SetSource sets Source field to given value.


### GetTarget

`func (o *AppSubscription) GetTarget() AppSubscriptionTarget`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AppSubscription) GetTargetOk() (*AppSubscriptionTarget, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AppSubscription) SetTarget(v AppSubscriptionTarget)`

SetTarget sets Target field to given value.


### GetState

`func (o *AppSubscription) GetState() AppSubscriptionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AppSubscription) GetStateOk() (*AppSubscriptionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AppSubscription) SetState(v AppSubscriptionState)`

SetState sets State field to given value.

### HasState

`func (o *AppSubscription) HasState() bool`

HasState returns a boolean if a field has been set.

### GetProject

`func (o *AppSubscription) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AppSubscription) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AppSubscription) SetProject(v Project)`

SetProject sets Project field to given value.


### GetChangeLog

`func (o *AppSubscription) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppSubscription) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppSubscription) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppSubscription) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetChange

`func (o *AppSubscription) GetChange() AppSubscriptionChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *AppSubscription) GetChangeOk() (*AppSubscriptionChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *AppSubscription) SetChange(v AppSubscriptionChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *AppSubscription) HasChange() bool`

HasChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


