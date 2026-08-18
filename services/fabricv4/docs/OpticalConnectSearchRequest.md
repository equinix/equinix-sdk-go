# OpticalConnectSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**OpticalConnectFilters**](OpticalConnectFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]OpticalConnectSortCriteria**](OpticalConnectSortCriteria.md) |  | [optional] 

## Methods

### NewOpticalConnectSearchRequest

`func NewOpticalConnectSearchRequest() *OpticalConnectSearchRequest`

NewOpticalConnectSearchRequest instantiates a new OpticalConnectSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectSearchRequestWithDefaults

`func NewOpticalConnectSearchRequestWithDefaults() *OpticalConnectSearchRequest`

NewOpticalConnectSearchRequestWithDefaults instantiates a new OpticalConnectSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *OpticalConnectSearchRequest) GetFilter() OpticalConnectFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OpticalConnectSearchRequest) GetFilterOk() (*OpticalConnectFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OpticalConnectSearchRequest) SetFilter(v OpticalConnectFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *OpticalConnectSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *OpticalConnectSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OpticalConnectSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OpticalConnectSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *OpticalConnectSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *OpticalConnectSearchRequest) GetSort() []OpticalConnectSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *OpticalConnectSearchRequest) GetSortOk() (*[]OpticalConnectSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *OpticalConnectSearchRequest) SetSort(v []OpticalConnectSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *OpticalConnectSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


