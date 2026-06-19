# IpBlocksSearchRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Filter** | Pointer to [**IpBlockFilter**](IpBlockFilter.md) |  | [optional] 
**Sort** | Pointer to [**[]SearchSortItem**](SearchSortItem.md) |  | [optional] 

## Methods

### NewIpBlocksSearchRequestBody

`func NewIpBlocksSearchRequestBody() *IpBlocksSearchRequestBody`

NewIpBlocksSearchRequestBody instantiates a new IpBlocksSearchRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlocksSearchRequestBodyWithDefaults

`func NewIpBlocksSearchRequestBodyWithDefaults() *IpBlocksSearchRequestBody`

NewIpBlocksSearchRequestBodyWithDefaults instantiates a new IpBlocksSearchRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *IpBlocksSearchRequestBody) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IpBlocksSearchRequestBody) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IpBlocksSearchRequestBody) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *IpBlocksSearchRequestBody) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetFilter

`func (o *IpBlocksSearchRequestBody) GetFilter() IpBlockFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IpBlocksSearchRequestBody) GetFilterOk() (*IpBlockFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IpBlocksSearchRequestBody) SetFilter(v IpBlockFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IpBlocksSearchRequestBody) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetSort

`func (o *IpBlocksSearchRequestBody) GetSort() []SearchSortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IpBlocksSearchRequestBody) GetSortOk() (*[]SearchSortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IpBlocksSearchRequestBody) SetSort(v []SearchSortItem)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IpBlocksSearchRequestBody) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


