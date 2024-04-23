# RouteFiltersData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Filter URI | [optional] 
**Type** | Pointer to [**ConnectionRouteFilterDataType**](ConnectionRouteFilterDataType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Filter identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**State** | Pointer to [**RouteFilterState**](RouteFilterState.md) |  | [optional] 
**Change** | Pointer to [**RouteFiltersChange**](RouteFiltersChange.md) |  | [optional] 
**NotMatchedRuleAction** | Pointer to [**RouteFiltersDataNotMatchedRuleAction**](RouteFiltersDataNotMatchedRuleAction.md) |  | [optional] 
**ConnectionsCount** | Pointer to **int32** |  | [optional] 
**RulesCount** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**RouteFiltersDataProject**](RouteFiltersDataProject.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRouteFiltersData

`func NewRouteFiltersData() *RouteFiltersData`

NewRouteFiltersData instantiates a new RouteFiltersData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersDataWithDefaults

`func NewRouteFiltersDataWithDefaults() *RouteFiltersData`

NewRouteFiltersDataWithDefaults instantiates a new RouteFiltersData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteFiltersData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFiltersData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFiltersData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFiltersData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteFiltersData) GetType() ConnectionRouteFilterDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFiltersData) GetTypeOk() (*ConnectionRouteFilterDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFiltersData) SetType(v ConnectionRouteFilterDataType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteFiltersData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteFiltersData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFiltersData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFiltersData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteFiltersData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteFiltersData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFiltersData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFiltersData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteFiltersData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteFiltersData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFiltersData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFiltersData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFiltersData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *RouteFiltersData) GetState() RouteFilterState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteFiltersData) GetStateOk() (*RouteFilterState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteFiltersData) SetState(v RouteFilterState)`

SetState sets State field to given value.

### HasState

`func (o *RouteFiltersData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChange

`func (o *RouteFiltersData) GetChange() RouteFiltersChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RouteFiltersData) GetChangeOk() (*RouteFiltersChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RouteFiltersData) SetChange(v RouteFiltersChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RouteFiltersData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetNotMatchedRuleAction

`func (o *RouteFiltersData) GetNotMatchedRuleAction() RouteFiltersDataNotMatchedRuleAction`

GetNotMatchedRuleAction returns the NotMatchedRuleAction field if non-nil, zero value otherwise.

### GetNotMatchedRuleActionOk

`func (o *RouteFiltersData) GetNotMatchedRuleActionOk() (*RouteFiltersDataNotMatchedRuleAction, bool)`

GetNotMatchedRuleActionOk returns a tuple with the NotMatchedRuleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotMatchedRuleAction

`func (o *RouteFiltersData) SetNotMatchedRuleAction(v RouteFiltersDataNotMatchedRuleAction)`

SetNotMatchedRuleAction sets NotMatchedRuleAction field to given value.

### HasNotMatchedRuleAction

`func (o *RouteFiltersData) HasNotMatchedRuleAction() bool`

HasNotMatchedRuleAction returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *RouteFiltersData) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *RouteFiltersData) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *RouteFiltersData) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *RouteFiltersData) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetRulesCount

`func (o *RouteFiltersData) GetRulesCount() int32`

GetRulesCount returns the RulesCount field if non-nil, zero value otherwise.

### GetRulesCountOk

`func (o *RouteFiltersData) GetRulesCountOk() (*int32, bool)`

GetRulesCountOk returns a tuple with the RulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesCount

`func (o *RouteFiltersData) SetRulesCount(v int32)`

SetRulesCount sets RulesCount field to given value.

### HasRulesCount

`func (o *RouteFiltersData) HasRulesCount() bool`

HasRulesCount returns a boolean if a field has been set.

### GetProject

`func (o *RouteFiltersData) GetProject() RouteFiltersDataProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteFiltersData) GetProjectOk() (*RouteFiltersDataProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteFiltersData) SetProject(v RouteFiltersDataProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *RouteFiltersData) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetChangelog

`func (o *RouteFiltersData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *RouteFiltersData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *RouteFiltersData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *RouteFiltersData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


