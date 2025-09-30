# PageResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 

## Methods

### NewPageResponseDto

`func NewPageResponseDto() *PageResponseDto`

NewPageResponseDto instantiates a new PageResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseDtoWithDefaults

`func NewPageResponseDtoWithDefaults() *PageResponseDto`

NewPageResponseDtoWithDefaults instantiates a new PageResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PageResponseDto) GetContent() []map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PageResponseDto) GetContentOk() (*[]map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PageResponseDto) SetContent(v []map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *PageResponseDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetData

`func (o *PageResponseDto) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageResponseDto) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageResponseDto) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *PageResponseDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PageResponseDto) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PageResponseDto) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PageResponseDto) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PageResponseDto) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


