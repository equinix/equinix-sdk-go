# PageResponseDtoMetroResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]MetroResponse**](MetroResponse.md) |  | [optional] 
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 

## Methods

### NewPageResponseDtoMetroResponse

`func NewPageResponseDtoMetroResponse() *PageResponseDtoMetroResponse`

NewPageResponseDtoMetroResponse instantiates a new PageResponseDtoMetroResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseDtoMetroResponseWithDefaults

`func NewPageResponseDtoMetroResponseWithDefaults() *PageResponseDtoMetroResponse`

NewPageResponseDtoMetroResponseWithDefaults instantiates a new PageResponseDtoMetroResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageResponseDtoMetroResponse) GetData() []MetroResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageResponseDtoMetroResponse) GetDataOk() (*[]MetroResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageResponseDtoMetroResponse) SetData(v []MetroResponse)`

SetData sets Data field to given value.

### HasData

`func (o *PageResponseDtoMetroResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PageResponseDtoMetroResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PageResponseDtoMetroResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PageResponseDtoMetroResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PageResponseDtoMetroResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


