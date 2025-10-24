# PaginationResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offset** | Pointer to **int32** | It is the starting point of the collection returned fromt the server | [optional] 
**Limit** | Pointer to **int32** | The page size | [optional] 
**Total** | Pointer to **int32** | The total number of results | [optional] 

## Methods

### NewPaginationResponseDto

`func NewPaginationResponseDto() *PaginationResponseDto`

NewPaginationResponseDto instantiates a new PaginationResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationResponseDtoWithDefaults

`func NewPaginationResponseDtoWithDefaults() *PaginationResponseDto`

NewPaginationResponseDtoWithDefaults instantiates a new PaginationResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *PaginationResponseDto) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationResponseDto) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationResponseDto) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaginationResponseDto) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationResponseDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationResponseDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationResponseDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationResponseDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *PaginationResponseDto) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationResponseDto) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationResponseDto) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginationResponseDto) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


