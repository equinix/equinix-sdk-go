# PortCapacitySearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**PortCapacityExpression**](PortCapacityExpression.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewPortCapacitySearchRequest

`func NewPortCapacitySearchRequest() *PortCapacitySearchRequest`

NewPortCapacitySearchRequest instantiates a new PortCapacitySearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortCapacitySearchRequestWithDefaults

`func NewPortCapacitySearchRequestWithDefaults() *PortCapacitySearchRequest`

NewPortCapacitySearchRequestWithDefaults instantiates a new PortCapacitySearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *PortCapacitySearchRequest) GetFilter() PortCapacityExpression`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PortCapacitySearchRequest) GetFilterOk() (*PortCapacityExpression, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PortCapacitySearchRequest) SetFilter(v PortCapacityExpression)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PortCapacitySearchRequest) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetPagination

`func (o *PortCapacitySearchRequest) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PortCapacitySearchRequest) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PortCapacitySearchRequest) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PortCapacitySearchRequest) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


