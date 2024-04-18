# PageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 

## Methods

### NewPageResponse

`func NewPageResponse() *PageResponse`

NewPageResponse instantiates a new PageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseWithDefaults

`func NewPageResponseWithDefaults() *PageResponse`

NewPageResponseWithDefaults instantiates a new PageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageResponse) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageResponse) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageResponse) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PageResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PageResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PageResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PageResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


