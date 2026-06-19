# RouteAggregationSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**RouteAggregationSortDirection**](RouteAggregationSortDirection.md) |  | [optional] [default to ROUTEAGGREGATIONSORTDIRECTION_DESC]
**Property** | Pointer to [**RouteAggregationSortBy**](RouteAggregationSortBy.md) |  | [optional] [default to ROUTEAGGREGATIONSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewRouteAggregationSortCriteria

`func NewRouteAggregationSortCriteria() *RouteAggregationSortCriteria`

NewRouteAggregationSortCriteria instantiates a new RouteAggregationSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationSortCriteriaWithDefaults

`func NewRouteAggregationSortCriteriaWithDefaults() *RouteAggregationSortCriteria`

NewRouteAggregationSortCriteriaWithDefaults instantiates a new RouteAggregationSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *RouteAggregationSortCriteria) GetDirection() RouteAggregationSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RouteAggregationSortCriteria) GetDirectionOk() (*RouteAggregationSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RouteAggregationSortCriteria) SetDirection(v RouteAggregationSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RouteAggregationSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *RouteAggregationSortCriteria) GetProperty() RouteAggregationSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteAggregationSortCriteria) GetPropertyOk() (*RouteAggregationSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteAggregationSortCriteria) SetProperty(v RouteAggregationSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteAggregationSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


