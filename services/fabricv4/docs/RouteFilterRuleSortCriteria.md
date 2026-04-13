# RouteFilterRuleSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**RouteFilterRuleSortDirection**](RouteFilterRuleSortDirection.md) |  | [optional] [default to ROUTEFILTERRULESORTDIRECTION_DESC]
**Property** | Pointer to [**RouteFilterRuleSortBy**](RouteFilterRuleSortBy.md) |  | [optional] [default to ROUTEFILTERRULESORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewRouteFilterRuleSortCriteria

`func NewRouteFilterRuleSortCriteria() *RouteFilterRuleSortCriteria`

NewRouteFilterRuleSortCriteria instantiates a new RouteFilterRuleSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRuleSortCriteriaWithDefaults

`func NewRouteFilterRuleSortCriteriaWithDefaults() *RouteFilterRuleSortCriteria`

NewRouteFilterRuleSortCriteriaWithDefaults instantiates a new RouteFilterRuleSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *RouteFilterRuleSortCriteria) GetDirection() RouteFilterRuleSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RouteFilterRuleSortCriteria) GetDirectionOk() (*RouteFilterRuleSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RouteFilterRuleSortCriteria) SetDirection(v RouteFilterRuleSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RouteFilterRuleSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *RouteFilterRuleSortCriteria) GetProperty() RouteFilterRuleSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteFilterRuleSortCriteria) GetPropertyOk() (*RouteFilterRuleSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteFilterRuleSortCriteria) SetProperty(v RouteFilterRuleSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteFilterRuleSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


