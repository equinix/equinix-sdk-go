# RouteAggregationRuleSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**RouteAggregationRuleSortDirection**](RouteAggregationRuleSortDirection.md) |  | [optional] [default to ROUTEAGGREGATIONRULESORTDIRECTION_DESC]
**Property** | Pointer to [**RouteAggregationRuleSortBy**](RouteAggregationRuleSortBy.md) |  | [optional] [default to ROUTEAGGREGATIONRULESORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewRouteAggregationRuleSortCriteria

`func NewRouteAggregationRuleSortCriteria() *RouteAggregationRuleSortCriteria`

NewRouteAggregationRuleSortCriteria instantiates a new RouteAggregationRuleSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRuleSortCriteriaWithDefaults

`func NewRouteAggregationRuleSortCriteriaWithDefaults() *RouteAggregationRuleSortCriteria`

NewRouteAggregationRuleSortCriteriaWithDefaults instantiates a new RouteAggregationRuleSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *RouteAggregationRuleSortCriteria) GetDirection() RouteAggregationRuleSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RouteAggregationRuleSortCriteria) GetDirectionOk() (*RouteAggregationRuleSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RouteAggregationRuleSortCriteria) SetDirection(v RouteAggregationRuleSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RouteAggregationRuleSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *RouteAggregationRuleSortCriteria) GetProperty() RouteAggregationRuleSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteAggregationRuleSortCriteria) GetPropertyOk() (*RouteAggregationRuleSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteAggregationRuleSortCriteria) SetProperty(v RouteAggregationRuleSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteAggregationRuleSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


