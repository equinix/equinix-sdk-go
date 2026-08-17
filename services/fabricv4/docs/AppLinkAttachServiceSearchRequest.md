# AppLinkAttachServiceSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AppLinkAttachServiceFilters**](AppLinkAttachServiceFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]AppLinkAttachServiceSortCriteria**](AppLinkAttachServiceSortCriteria.md) |  | [optional] 

## Methods

### NewAppLinkAttachServiceSearchRequest

`func NewAppLinkAttachServiceSearchRequest() *AppLinkAttachServiceSearchRequest`

NewAppLinkAttachServiceSearchRequest instantiates a new AppLinkAttachServiceSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachServiceSearchRequestWithDefaults

`func NewAppLinkAttachServiceSearchRequestWithDefaults() *AppLinkAttachServiceSearchRequest`

NewAppLinkAttachServiceSearchRequestWithDefaults instantiates a new AppLinkAttachServiceSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AppLinkAttachServiceSearchRequest) GetFilter() AppLinkAttachServiceFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppLinkAttachServiceSearchRequest) GetFilterOk() (*AppLinkAttachServiceFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppLinkAttachServiceSearchRequest) SetFilter(v AppLinkAttachServiceFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppLinkAttachServiceSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *AppLinkAttachServiceSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppLinkAttachServiceSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppLinkAttachServiceSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppLinkAttachServiceSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *AppLinkAttachServiceSearchRequest) GetSort() []AppLinkAttachServiceSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AppLinkAttachServiceSearchRequest) GetSortOk() (*[]AppLinkAttachServiceSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AppLinkAttachServiceSearchRequest) SetSort(v []AppLinkAttachServiceSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AppLinkAttachServiceSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


