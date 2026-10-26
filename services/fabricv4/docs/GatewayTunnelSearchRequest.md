# GatewayTunnelSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**SearchFilter**](SearchFilter.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]GatewayTunnelSortCriteria**](GatewayTunnelSortCriteria.md) |  | [optional] 

## Methods

### NewGatewayTunnelSearchRequest

`func NewGatewayTunnelSearchRequest() *GatewayTunnelSearchRequest`

NewGatewayTunnelSearchRequest instantiates a new GatewayTunnelSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTunnelSearchRequestWithDefaults

`func NewGatewayTunnelSearchRequestWithDefaults() *GatewayTunnelSearchRequest`

NewGatewayTunnelSearchRequestWithDefaults instantiates a new GatewayTunnelSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *GatewayTunnelSearchRequest) GetFilter() SearchFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GatewayTunnelSearchRequest) GetFilterOk() (*SearchFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GatewayTunnelSearchRequest) SetFilter(v SearchFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GatewayTunnelSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *GatewayTunnelSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GatewayTunnelSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GatewayTunnelSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GatewayTunnelSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *GatewayTunnelSearchRequest) GetSort() []GatewayTunnelSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *GatewayTunnelSearchRequest) GetSortOk() (*[]GatewayTunnelSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *GatewayTunnelSearchRequest) SetSort(v []GatewayTunnelSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *GatewayTunnelSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


