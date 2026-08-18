# ExchangeServiceSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ExchangeServiceSearchExpression**](ExchangeServiceSearchExpression.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]ExchangeServiceSearchSortCriteria**](ExchangeServiceSearchSortCriteria.md) |  | [optional] 

## Methods

### NewExchangeServiceSearchRequest

`func NewExchangeServiceSearchRequest() *ExchangeServiceSearchRequest`

NewExchangeServiceSearchRequest instantiates a new ExchangeServiceSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceSearchRequestWithDefaults

`func NewExchangeServiceSearchRequestWithDefaults() *ExchangeServiceSearchRequest`

NewExchangeServiceSearchRequestWithDefaults instantiates a new ExchangeServiceSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ExchangeServiceSearchRequest) GetFilter() ExchangeServiceSearchExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ExchangeServiceSearchRequest) GetFilterOk() (*ExchangeServiceSearchExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ExchangeServiceSearchRequest) SetFilter(v ExchangeServiceSearchExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ExchangeServiceSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *ExchangeServiceSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ExchangeServiceSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ExchangeServiceSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ExchangeServiceSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *ExchangeServiceSearchRequest) GetSort() []ExchangeServiceSearchSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *ExchangeServiceSearchRequest) GetSortOk() (*[]ExchangeServiceSearchSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *ExchangeServiceSearchRequest) SetSort(v []ExchangeServiceSearchSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *ExchangeServiceSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


