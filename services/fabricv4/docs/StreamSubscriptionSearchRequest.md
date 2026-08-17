# StreamSubscriptionSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**StreamSubscriptionSearchFilters**](StreamSubscriptionSearchFilters.md) |  | 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]StreamSubscriptionSearchSortCriteria**](StreamSubscriptionSearchSortCriteria.md) |  | [optional] 

## Methods

### NewStreamSubscriptionSearchRequest

`func NewStreamSubscriptionSearchRequest(filter StreamSubscriptionSearchFilters, ) *StreamSubscriptionSearchRequest`

NewStreamSubscriptionSearchRequest instantiates a new StreamSubscriptionSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSearchRequestWithDefaults

`func NewStreamSubscriptionSearchRequestWithDefaults() *StreamSubscriptionSearchRequest`

NewStreamSubscriptionSearchRequestWithDefaults instantiates a new StreamSubscriptionSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *StreamSubscriptionSearchRequest) GetFilter() StreamSubscriptionSearchFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *StreamSubscriptionSearchRequest) GetFilterOk() (*StreamSubscriptionSearchFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *StreamSubscriptionSearchRequest) SetFilter(v StreamSubscriptionSearchFilters)`

SetFilter sets Filter field to given value.


### GetPagination

`func (o *StreamSubscriptionSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *StreamSubscriptionSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *StreamSubscriptionSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *StreamSubscriptionSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *StreamSubscriptionSearchRequest) GetSort() []StreamSubscriptionSearchSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *StreamSubscriptionSearchRequest) GetSortOk() (*[]StreamSubscriptionSearchSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *StreamSubscriptionSearchRequest) SetSort(v []StreamSubscriptionSearchSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *StreamSubscriptionSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


