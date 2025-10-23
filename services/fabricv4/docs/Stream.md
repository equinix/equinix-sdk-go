# Stream

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to [**StreamType**](StreamType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-provided stream name | [optional] 
**Description** | Pointer to **string** | Customer-provided stream description | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to [**StreamState**](StreamState.md) |  | [optional] 
**AssetsCount** | Pointer to **int32** | Stream assets count | [optional] 
**StreamSubscriptionsCount** | Pointer to **int32** | Stream subscriptions count | [optional] 
**AlertRulesCount** | Pointer to **int32** | Stream alert rules count | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewStream

`func NewStream() *Stream`

NewStream instantiates a new Stream object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamWithDefaults

`func NewStreamWithDefaults() *Stream`

NewStreamWithDefaults instantiates a new Stream object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Stream) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Stream) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Stream) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Stream) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Stream) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Stream) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Stream) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Stream) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *Stream) GetType() StreamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Stream) GetTypeOk() (*StreamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Stream) SetType(v StreamType)`

SetType sets Type field to given value.

### HasType

`func (o *Stream) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *Stream) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Stream) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Stream) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Stream) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Stream) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Stream) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Stream) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Stream) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProject

`func (o *Stream) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Stream) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Stream) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Stream) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *Stream) GetState() StreamState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Stream) GetStateOk() (*StreamState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Stream) SetState(v StreamState)`

SetState sets State field to given value.

### HasState

`func (o *Stream) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAssetsCount

`func (o *Stream) GetAssetsCount() int32`

GetAssetsCount returns the AssetsCount field if non-nil, zero value otherwise.

### GetAssetsCountOk

`func (o *Stream) GetAssetsCountOk() (*int32, bool)`

GetAssetsCountOk returns a tuple with the AssetsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsCount

`func (o *Stream) SetAssetsCount(v int32)`

SetAssetsCount sets AssetsCount field to given value.

### HasAssetsCount

`func (o *Stream) HasAssetsCount() bool`

HasAssetsCount returns a boolean if a field has been set.

### GetStreamSubscriptionsCount

`func (o *Stream) GetStreamSubscriptionsCount() int32`

GetStreamSubscriptionsCount returns the StreamSubscriptionsCount field if non-nil, zero value otherwise.

### GetStreamSubscriptionsCountOk

`func (o *Stream) GetStreamSubscriptionsCountOk() (*int32, bool)`

GetStreamSubscriptionsCountOk returns a tuple with the StreamSubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamSubscriptionsCount

`func (o *Stream) SetStreamSubscriptionsCount(v int32)`

SetStreamSubscriptionsCount sets StreamSubscriptionsCount field to given value.

### HasStreamSubscriptionsCount

`func (o *Stream) HasStreamSubscriptionsCount() bool`

HasStreamSubscriptionsCount returns a boolean if a field has been set.

### GetAlertRulesCount

`func (o *Stream) GetAlertRulesCount() int32`

GetAlertRulesCount returns the AlertRulesCount field if non-nil, zero value otherwise.

### GetAlertRulesCountOk

`func (o *Stream) GetAlertRulesCountOk() (*int32, bool)`

GetAlertRulesCountOk returns a tuple with the AlertRulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertRulesCount

`func (o *Stream) SetAlertRulesCount(v int32)`

SetAlertRulesCount sets AlertRulesCount field to given value.

### HasAlertRulesCount

`func (o *Stream) HasAlertRulesCount() bool`

HasAlertRulesCount returns a boolean if a field has been set.

### GetChangeLog

`func (o *Stream) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *Stream) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *Stream) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *Stream) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


