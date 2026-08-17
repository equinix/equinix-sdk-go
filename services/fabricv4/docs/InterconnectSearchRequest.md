# InterconnectSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**InterconnectFilter**](InterconnectFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]InterconnectSortCriteria**](InterconnectSortCriteria.md) |  | [optional] 

## Methods

### NewInterconnectSearchRequest

`func NewInterconnectSearchRequest() *InterconnectSearchRequest`

NewInterconnectSearchRequest instantiates a new InterconnectSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectSearchRequestWithDefaults

`func NewInterconnectSearchRequestWithDefaults() *InterconnectSearchRequest`

NewInterconnectSearchRequestWithDefaults instantiates a new InterconnectSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *InterconnectSearchRequest) GetFilter() InterconnectFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InterconnectSearchRequest) GetFilterOk() (*InterconnectFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InterconnectSearchRequest) SetFilter(v InterconnectFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InterconnectSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *InterconnectSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InterconnectSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InterconnectSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InterconnectSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *InterconnectSearchRequest) GetSort() []InterconnectSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *InterconnectSearchRequest) GetSortOk() (*[]InterconnectSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *InterconnectSearchRequest) SetSort(v []InterconnectSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *InterconnectSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


