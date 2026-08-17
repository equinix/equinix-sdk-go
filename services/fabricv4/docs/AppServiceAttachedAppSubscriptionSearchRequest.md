# AppServiceAttachedAppSubscriptionSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**AppServiceAttachedAppSubscriptionFilters**](AppServiceAttachedAppSubscriptionFilters.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationRequest**](PaginationRequest.md) |  | [optional] 
**Sort** | Pointer to [**[]AppServiceAttachedAppSubscriptionSortCriteria**](AppServiceAttachedAppSubscriptionSortCriteria.md) |  | [optional] 

## Methods

### NewAppServiceAttachedAppSubscriptionSearchRequest

`func NewAppServiceAttachedAppSubscriptionSearchRequest() *AppServiceAttachedAppSubscriptionSearchRequest`

NewAppServiceAttachedAppSubscriptionSearchRequest instantiates a new AppServiceAttachedAppSubscriptionSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppServiceAttachedAppSubscriptionSearchRequestWithDefaults

`func NewAppServiceAttachedAppSubscriptionSearchRequestWithDefaults() *AppServiceAttachedAppSubscriptionSearchRequest`

NewAppServiceAttachedAppSubscriptionSearchRequestWithDefaults instantiates a new AppServiceAttachedAppSubscriptionSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetFilter() AppServiceAttachedAppSubscriptionFilters`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetFilterOk() (*AppServiceAttachedAppSubscriptionFilters, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) SetFilter(v AppServiceAttachedAppSubscriptionFilters)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetPagination() PaginationRequest`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetPaginationOk() (*PaginationRequest, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) SetPagination(v PaginationRequest)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetSort() []AppServiceAttachedAppSubscriptionSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) GetSortOk() (*[]AppServiceAttachedAppSubscriptionSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) SetSort(v []AppServiceAttachedAppSubscriptionSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AppServiceAttachedAppSubscriptionSearchRequest) HasSort() bool`

HasSort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


