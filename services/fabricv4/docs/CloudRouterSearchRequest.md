# CloudRouterSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CloudRouterFilters**](CloudRouterFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]CloudRouterSortCriteria**](CloudRouterSortCriteria.md) |  | [optional] 

## Methods

### NewCloudRouterSearchRequest

`func NewCloudRouterSearchRequest() *CloudRouterSearchRequest`

NewCloudRouterSearchRequest instantiates a new CloudRouterSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterSearchRequestWithDefaults

`func NewCloudRouterSearchRequestWithDefaults() *CloudRouterSearchRequest`

NewCloudRouterSearchRequestWithDefaults instantiates a new CloudRouterSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CloudRouterSearchRequest) GetFilter() CloudRouterFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CloudRouterSearchRequest) GetFilterOk() (*CloudRouterFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CloudRouterSearchRequest) SetFilter(v CloudRouterFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CloudRouterSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *CloudRouterSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CloudRouterSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CloudRouterSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CloudRouterSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *CloudRouterSearchRequest) GetSort() []CloudRouterSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *CloudRouterSearchRequest) GetSortOk() (*[]CloudRouterSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *CloudRouterSearchRequest) SetSort(v []CloudRouterSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *CloudRouterSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


