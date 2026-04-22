# RouteFilterRulesSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**RouteFilterRulesFilter**](RouteFilterRulesFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]RouteFilterRuleSortCriteria**](RouteFilterRuleSortCriteria.md) |  | [optional] 

## Methods

### NewRouteFilterRulesSearchRequest

`func NewRouteFilterRulesSearchRequest() *RouteFilterRulesSearchRequest`

NewRouteFilterRulesSearchRequest instantiates a new RouteFilterRulesSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesSearchRequestWithDefaults

`func NewRouteFilterRulesSearchRequestWithDefaults() *RouteFilterRulesSearchRequest`

NewRouteFilterRulesSearchRequestWithDefaults instantiates a new RouteFilterRulesSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *RouteFilterRulesSearchRequest) GetFilter() RouteFilterRulesFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *RouteFilterRulesSearchRequest) GetFilterOk() (*RouteFilterRulesFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *RouteFilterRulesSearchRequest) SetFilter(v RouteFilterRulesFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *RouteFilterRulesSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *RouteFilterRulesSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteFilterRulesSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteFilterRulesSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteFilterRulesSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *RouteFilterRulesSearchRequest) GetSort() []RouteFilterRuleSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *RouteFilterRulesSearchRequest) GetSortOk() (*[]RouteFilterRuleSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *RouteFilterRulesSearchRequest) SetSort(v []RouteFilterRuleSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *RouteFilterRulesSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


