# ServiceTokenSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**ServiceTokenSearchExpression**](ServiceTokenSearchExpression.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 

## Methods

### NewServiceTokenSearchRequest

`func NewServiceTokenSearchRequest() *ServiceTokenSearchRequest`

NewServiceTokenSearchRequest instantiates a new ServiceTokenSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenSearchRequestWithDefaults

`func NewServiceTokenSearchRequestWithDefaults() *ServiceTokenSearchRequest`

NewServiceTokenSearchRequestWithDefaults instantiates a new ServiceTokenSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *ServiceTokenSearchRequest) GetFilter() ServiceTokenSearchExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ServiceTokenSearchRequest) GetFilterOk() (*ServiceTokenSearchExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ServiceTokenSearchRequest) SetFilter(v ServiceTokenSearchExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ServiceTokenSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *ServiceTokenSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceTokenSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceTokenSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceTokenSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


