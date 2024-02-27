# RouteFilterRulesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Filter Rules URI | [optional] 
**Type** | Pointer to [**RouteFilterRulesDataType**](RouteFilterRulesDataType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Filter Rule identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided Route Filter Rule description | [optional] 
**State** | Pointer to [**RouteFilterRuleState**](RouteFilterRuleState.md) |  | [optional] 
**PrefixMatch** | Pointer to **string** | prefix matching operator | [optional] [default to "orlonger"]
**Change** | Pointer to [**RouteFilterRulesChange**](RouteFilterRulesChange.md) |  | [optional] 
**Action** | Pointer to [**RouteFilterRulesDataAction**](RouteFilterRulesDataAction.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRouteFilterRulesData

`func NewRouteFilterRulesData() *RouteFilterRulesData`

NewRouteFilterRulesData instantiates a new RouteFilterRulesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesDataWithDefaults

`func NewRouteFilterRulesDataWithDefaults() *RouteFilterRulesData`

NewRouteFilterRulesDataWithDefaults instantiates a new RouteFilterRulesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteFilterRulesData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFilterRulesData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFilterRulesData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFilterRulesData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteFilterRulesData) GetType() RouteFilterRulesDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterRulesData) GetTypeOk() (*RouteFilterRulesDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterRulesData) SetType(v RouteFilterRulesDataType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteFilterRulesData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteFilterRulesData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFilterRulesData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFilterRulesData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteFilterRulesData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteFilterRulesData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteFilterRulesData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteFilterRulesData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteFilterRulesData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteFilterRulesData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteFilterRulesData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteFilterRulesData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteFilterRulesData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *RouteFilterRulesData) GetState() RouteFilterRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteFilterRulesData) GetStateOk() (*RouteFilterRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteFilterRulesData) SetState(v RouteFilterRuleState)`

SetState sets State field to given value.

### HasState

`func (o *RouteFilterRulesData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPrefixMatch

`func (o *RouteFilterRulesData) GetPrefixMatch() string`

GetPrefixMatch returns the PrefixMatch field if non-nil, zero value otherwise.

### GetPrefixMatchOk

`func (o *RouteFilterRulesData) GetPrefixMatchOk() (*string, bool)`

GetPrefixMatchOk returns a tuple with the PrefixMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixMatch

`func (o *RouteFilterRulesData) SetPrefixMatch(v string)`

SetPrefixMatch sets PrefixMatch field to given value.

### HasPrefixMatch

`func (o *RouteFilterRulesData) HasPrefixMatch() bool`

HasPrefixMatch returns a boolean if a field has been set.

### GetChange

`func (o *RouteFilterRulesData) GetChange() RouteFilterRulesChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RouteFilterRulesData) GetChangeOk() (*RouteFilterRulesChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RouteFilterRulesData) SetChange(v RouteFilterRulesChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RouteFilterRulesData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetAction

`func (o *RouteFilterRulesData) GetAction() RouteFilterRulesDataAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RouteFilterRulesData) GetActionOk() (*RouteFilterRulesDataAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RouteFilterRulesData) SetAction(v RouteFilterRulesDataAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *RouteFilterRulesData) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetPrefix

`func (o *RouteFilterRulesData) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteFilterRulesData) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteFilterRulesData) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RouteFilterRulesData) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetChangelog

`func (o *RouteFilterRulesData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *RouteFilterRulesData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *RouteFilterRulesData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *RouteFilterRulesData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


