# NetworkSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**NetworkFilter**](NetworkFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]NetworkSortCriteria**](NetworkSortCriteria.md) |  | [optional] 

## Methods

### NewNetworkSearchRequest

`func NewNetworkSearchRequest() *NetworkSearchRequest`

NewNetworkSearchRequest instantiates a new NetworkSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSearchRequestWithDefaults

`func NewNetworkSearchRequestWithDefaults() *NetworkSearchRequest`

NewNetworkSearchRequestWithDefaults instantiates a new NetworkSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *NetworkSearchRequest) GetFilter() NetworkFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *NetworkSearchRequest) GetFilterOk() (*NetworkFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *NetworkSearchRequest) SetFilter(v NetworkFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *NetworkSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *NetworkSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NetworkSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NetworkSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *NetworkSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *NetworkSearchRequest) GetSort() []NetworkSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *NetworkSearchRequest) GetSortOk() (*[]NetworkSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *NetworkSearchRequest) SetSort(v []NetworkSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *NetworkSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


