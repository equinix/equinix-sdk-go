# GatewaySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**SearchFilter**](SearchFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]GatewaySortCriteria**](GatewaySortCriteria.md) |  | [optional] 

## Methods

### NewGatewaySearchRequest

`func NewGatewaySearchRequest() *GatewaySearchRequest`

NewGatewaySearchRequest instantiates a new GatewaySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySearchRequestWithDefaults

`func NewGatewaySearchRequestWithDefaults() *GatewaySearchRequest`

NewGatewaySearchRequestWithDefaults instantiates a new GatewaySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *GatewaySearchRequest) GetFilter() SearchFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GatewaySearchRequest) GetFilterOk() (*SearchFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GatewaySearchRequest) SetFilter(v SearchFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GatewaySearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *GatewaySearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GatewaySearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GatewaySearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GatewaySearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *GatewaySearchRequest) GetSort() []GatewaySortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GatewaySearchRequest) GetSortOk() (*[]GatewaySortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GatewaySearchRequest) SetSort(v []GatewaySortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GatewaySearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


