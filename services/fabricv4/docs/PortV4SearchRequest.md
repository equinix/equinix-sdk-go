# PortV4SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | [**PortExpression**](PortExpression.md) |  | 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]PortSortCriteria**](PortSortCriteria.md) |  | [optional] 

## Methods

### NewPortV4SearchRequest

`func NewPortV4SearchRequest(filter PortExpression, ) *PortV4SearchRequest`

NewPortV4SearchRequest instantiates a new PortV4SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortV4SearchRequestWithDefaults

`func NewPortV4SearchRequestWithDefaults() *PortV4SearchRequest`

NewPortV4SearchRequestWithDefaults instantiates a new PortV4SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PortV4SearchRequest) GetFilter() PortExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PortV4SearchRequest) GetFilterOk() (*PortExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PortV4SearchRequest) SetFilter(v PortExpression)`

SetFilter sets Filter field to given value.


### GetPagination

`func (o *PortV4SearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PortV4SearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PortV4SearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PortV4SearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *PortV4SearchRequest) GetSort() []PortSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PortV4SearchRequest) GetSortOk() (*[]PortSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PortV4SearchRequest) SetSort(v []PortSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *PortV4SearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


