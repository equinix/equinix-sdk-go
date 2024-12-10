# RouteAggregationsSearchBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**RouteAggregationsSearchBaseFilter**](RouteAggregationsSearchBaseFilter.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteAggregationSortItem**](RouteAggregationSortItem.md) |  | [optional] 

## Methods

### NewRouteAggregationsSearchBase

`func NewRouteAggregationsSearchBase() *RouteAggregationsSearchBase`

NewRouteAggregationsSearchBase instantiates a new RouteAggregationsSearchBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsSearchBaseWithDefaults

`func NewRouteAggregationsSearchBaseWithDefaults() *RouteAggregationsSearchBase`

NewRouteAggregationsSearchBaseWithDefaults instantiates a new RouteAggregationsSearchBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteAggregationsSearchBase) GetFilter() RouteAggregationsSearchBaseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteAggregationsSearchBase) GetFilterOk() (*RouteAggregationsSearchBaseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteAggregationsSearchBase) SetFilter(v RouteAggregationsSearchBaseFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteAggregationsSearchBase) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteAggregationsSearchBase) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteAggregationsSearchBase) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteAggregationsSearchBase) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteAggregationsSearchBase) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteAggregationsSearchBase) GetSort() []RouteAggregationSortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteAggregationsSearchBase) GetSortOk() (*[]RouteAggregationSortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteAggregationsSearchBase) SetSort(v []RouteAggregationSortItem)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteAggregationsSearchBase) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


