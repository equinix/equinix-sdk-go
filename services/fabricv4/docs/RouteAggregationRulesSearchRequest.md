# RouteAggregationRulesSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**RouteAggregationRulesFilter**](RouteAggregationRulesFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteAggregationRuleSortCriteria**](RouteAggregationRuleSortCriteria.md) |  | [optional] 

## Methods

### NewRouteAggregationRulesSearchRequest

`func NewRouteAggregationRulesSearchRequest() *RouteAggregationRulesSearchRequest`

NewRouteAggregationRulesSearchRequest instantiates a new RouteAggregationRulesSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesSearchRequestWithDefaults

`func NewRouteAggregationRulesSearchRequestWithDefaults() *RouteAggregationRulesSearchRequest`

NewRouteAggregationRulesSearchRequestWithDefaults instantiates a new RouteAggregationRulesSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteAggregationRulesSearchRequest) GetFilter() RouteAggregationRulesFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteAggregationRulesSearchRequest) GetFilterOk() (*RouteAggregationRulesFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteAggregationRulesSearchRequest) SetFilter(v RouteAggregationRulesFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteAggregationRulesSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteAggregationRulesSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteAggregationRulesSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteAggregationRulesSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteAggregationRulesSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteAggregationRulesSearchRequest) GetSort() []RouteAggregationRuleSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteAggregationRulesSearchRequest) GetSortOk() (*[]RouteAggregationRuleSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteAggregationRulesSearchRequest) SetSort(v []RouteAggregationRuleSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteAggregationRulesSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


