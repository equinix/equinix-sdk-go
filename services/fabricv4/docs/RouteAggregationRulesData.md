# RouteAggregationRulesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Aggregation Rules URI | [optional] 
**Type** | Pointer to [**RouteAggregationRulesDataType**](RouteAggregationRulesDataType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Aggregation Rule identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided Route Aggregation Rule description | [optional] 
**State** | Pointer to [**RouteAggregationRuleState**](RouteAggregationRuleState.md) |  | [optional] 
**Change** | Pointer to [**RouteAggregationRulesChange**](RouteAggregationRulesChange.md) |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRouteAggregationRulesData

`func NewRouteAggregationRulesData() *RouteAggregationRulesData`

NewRouteAggregationRulesData instantiates a new RouteAggregationRulesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesDataWithDefaults

`func NewRouteAggregationRulesDataWithDefaults() *RouteAggregationRulesData`

NewRouteAggregationRulesDataWithDefaults instantiates a new RouteAggregationRulesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteAggregationRulesData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationRulesData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationRulesData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationRulesData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteAggregationRulesData) GetType() RouteAggregationRulesDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationRulesData) GetTypeOk() (*RouteAggregationRulesDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationRulesData) SetType(v RouteAggregationRulesDataType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteAggregationRulesData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteAggregationRulesData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationRulesData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationRulesData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteAggregationRulesData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteAggregationRulesData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteAggregationRulesData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteAggregationRulesData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteAggregationRulesData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteAggregationRulesData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteAggregationRulesData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteAggregationRulesData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteAggregationRulesData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *RouteAggregationRulesData) GetState() RouteAggregationRuleState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteAggregationRulesData) GetStateOk() (*RouteAggregationRuleState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteAggregationRulesData) SetState(v RouteAggregationRuleState)`

SetState sets State field to given value.

### HasState

`func (o *RouteAggregationRulesData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChange

`func (o *RouteAggregationRulesData) GetChange() RouteAggregationRulesChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RouteAggregationRulesData) GetChangeOk() (*RouteAggregationRulesChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RouteAggregationRulesData) SetChange(v RouteAggregationRulesChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RouteAggregationRulesData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetPrefix

`func (o *RouteAggregationRulesData) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RouteAggregationRulesData) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RouteAggregationRulesData) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RouteAggregationRulesData) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetChangeLog

`func (o *RouteAggregationRulesData) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RouteAggregationRulesData) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RouteAggregationRulesData) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *RouteAggregationRulesData) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


