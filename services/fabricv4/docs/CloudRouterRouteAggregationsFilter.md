# CloudRouterRouteAggregationsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to [**[]CloudRouterRouteAggregationExpression**](CloudRouterRouteAggregationExpression.md) |  | [optional] 
**Or** | Pointer to [**[]CloudRouterRouteAggregationExpression**](CloudRouterRouteAggregationExpression.md) |  | [optional] 

## Methods

### NewCloudRouterRouteAggregationsFilter

`func NewCloudRouterRouteAggregationsFilter() *CloudRouterRouteAggregationsFilter`

NewCloudRouterRouteAggregationsFilter instantiates a new CloudRouterRouteAggregationsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterRouteAggregationsFilterWithDefaults

`func NewCloudRouterRouteAggregationsFilterWithDefaults() *CloudRouterRouteAggregationsFilter`

NewCloudRouterRouteAggregationsFilterWithDefaults instantiates a new CloudRouterRouteAggregationsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *CloudRouterRouteAggregationsFilter) GetAnd() []CloudRouterRouteAggregationExpression`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *CloudRouterRouteAggregationsFilter) GetAndOk() (*[]CloudRouterRouteAggregationExpression, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *CloudRouterRouteAggregationsFilter) SetAnd(v []CloudRouterRouteAggregationExpression)`

SetAnd sets And field to given value.

### HasAnd

`func (o *CloudRouterRouteAggregationsFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetOr

`func (o *CloudRouterRouteAggregationsFilter) GetOr() []CloudRouterRouteAggregationExpression`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *CloudRouterRouteAggregationsFilter) GetOrOk() (*[]CloudRouterRouteAggregationExpression, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *CloudRouterRouteAggregationsFilter) SetOr(v []CloudRouterRouteAggregationExpression)`

SetOr sets Or field to given value.

### HasOr

`func (o *CloudRouterRouteAggregationsFilter) HasOr() bool`

HasOr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


