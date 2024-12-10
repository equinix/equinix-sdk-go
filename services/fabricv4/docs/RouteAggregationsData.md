# RouteAggregationsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Aggregation URI | [optional] 
**Type** | Pointer to [**RouteAggregationsBaseType**](RouteAggregationsBaseType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Aggregation identifier | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** | Customer-provided connection description | [optional] 
**State** | Pointer to [**RouteAggregationState**](RouteAggregationState.md) |  | [optional] 
**Change** | Pointer to [**RouteAggregationsChange**](RouteAggregationsChange.md) |  | [optional] 
**ConnectionsCount** | Pointer to **int32** |  | [optional] 
**RulesCount** | Pointer to **int32** |  | [optional] 
**Project** | Pointer to [**RouteAggregationsDataProject**](RouteAggregationsDataProject.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewRouteAggregationsData

`func NewRouteAggregationsData() *RouteAggregationsData`

NewRouteAggregationsData instantiates a new RouteAggregationsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsDataWithDefaults

`func NewRouteAggregationsDataWithDefaults() *RouteAggregationsData`

NewRouteAggregationsDataWithDefaults instantiates a new RouteAggregationsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RouteAggregationsData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteAggregationsData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteAggregationsData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteAggregationsData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *RouteAggregationsData) GetType() RouteAggregationsBaseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteAggregationsData) GetTypeOk() (*RouteAggregationsBaseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteAggregationsData) SetType(v RouteAggregationsBaseType)`

SetType sets Type field to given value.

### HasType

`func (o *RouteAggregationsData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *RouteAggregationsData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteAggregationsData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteAggregationsData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *RouteAggregationsData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *RouteAggregationsData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RouteAggregationsData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RouteAggregationsData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RouteAggregationsData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RouteAggregationsData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RouteAggregationsData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RouteAggregationsData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RouteAggregationsData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *RouteAggregationsData) GetState() RouteAggregationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RouteAggregationsData) GetStateOk() (*RouteAggregationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RouteAggregationsData) SetState(v RouteAggregationState)`

SetState sets State field to given value.

### HasState

`func (o *RouteAggregationsData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChange

`func (o *RouteAggregationsData) GetChange() RouteAggregationsChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *RouteAggregationsData) GetChangeOk() (*RouteAggregationsChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *RouteAggregationsData) SetChange(v RouteAggregationsChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *RouteAggregationsData) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *RouteAggregationsData) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *RouteAggregationsData) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *RouteAggregationsData) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *RouteAggregationsData) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetRulesCount

`func (o *RouteAggregationsData) GetRulesCount() int32`

GetRulesCount returns the RulesCount field if non-nil, zero value otherwise.

### GetRulesCountOk

`func (o *RouteAggregationsData) GetRulesCountOk() (*int32, bool)`

GetRulesCountOk returns a tuple with the RulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesCount

`func (o *RouteAggregationsData) SetRulesCount(v int32)`

SetRulesCount sets RulesCount field to given value.

### HasRulesCount

`func (o *RouteAggregationsData) HasRulesCount() bool`

HasRulesCount returns a boolean if a field has been set.

### GetProject

`func (o *RouteAggregationsData) GetProject() RouteAggregationsDataProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteAggregationsData) GetProjectOk() (*RouteAggregationsDataProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteAggregationsData) SetProject(v RouteAggregationsDataProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *RouteAggregationsData) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetChangeLog

`func (o *RouteAggregationsData) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *RouteAggregationsData) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *RouteAggregationsData) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *RouteAggregationsData) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


