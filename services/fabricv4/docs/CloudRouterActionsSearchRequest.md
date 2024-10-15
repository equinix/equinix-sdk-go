# CloudRouterActionsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CloudRouterActionsSearchFilters**](CloudRouterActionsSearchFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]CloudRouterActionsSearchSortCriteria**](CloudRouterActionsSearchSortCriteria.md) |  | [optional] 

## Methods

### NewCloudRouterActionsSearchRequest

`func NewCloudRouterActionsSearchRequest() *CloudRouterActionsSearchRequest`

NewCloudRouterActionsSearchRequest instantiates a new CloudRouterActionsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionsSearchRequestWithDefaults

`func NewCloudRouterActionsSearchRequestWithDefaults() *CloudRouterActionsSearchRequest`

NewCloudRouterActionsSearchRequestWithDefaults instantiates a new CloudRouterActionsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CloudRouterActionsSearchRequest) GetFilter() CloudRouterActionsSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CloudRouterActionsSearchRequest) GetFilterOk() (*CloudRouterActionsSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CloudRouterActionsSearchRequest) SetFilter(v CloudRouterActionsSearchFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CloudRouterActionsSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *CloudRouterActionsSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CloudRouterActionsSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CloudRouterActionsSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CloudRouterActionsSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *CloudRouterActionsSearchRequest) GetSort() []CloudRouterActionsSearchSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CloudRouterActionsSearchRequest) GetSortOk() (*[]CloudRouterActionsSearchSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CloudRouterActionsSearchRequest) SetSort(v []CloudRouterActionsSearchSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CloudRouterActionsSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


