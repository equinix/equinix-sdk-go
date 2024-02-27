# RouteFiltersSearchBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**RouteFiltersSearchBaseFilter**](RouteFiltersSearchBaseFilter.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]SortItem**](SortItem.md) |  | [optional] 

## Methods

### NewRouteFiltersSearchBase

`func NewRouteFiltersSearchBase() *RouteFiltersSearchBase`

NewRouteFiltersSearchBase instantiates a new RouteFiltersSearchBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFiltersSearchBaseWithDefaults

`func NewRouteFiltersSearchBaseWithDefaults() *RouteFiltersSearchBase`

NewRouteFiltersSearchBaseWithDefaults instantiates a new RouteFiltersSearchBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteFiltersSearchBase) GetFilter() RouteFiltersSearchBaseFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteFiltersSearchBase) GetFilterOk() (*RouteFiltersSearchBaseFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteFiltersSearchBase) SetFilter(v RouteFiltersSearchBaseFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteFiltersSearchBase) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteFiltersSearchBase) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteFiltersSearchBase) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteFiltersSearchBase) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteFiltersSearchBase) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteFiltersSearchBase) GetSort() []SortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteFiltersSearchBase) GetSortOk() (*[]SortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteFiltersSearchBase) SetSort(v []SortItem)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteFiltersSearchBase) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


