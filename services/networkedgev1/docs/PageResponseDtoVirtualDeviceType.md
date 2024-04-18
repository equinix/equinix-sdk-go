# PageResponseDtoVirtualDeviceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]VirtualDeviceType**](VirtualDeviceType.md) | Array of available virtual device types | [optional] 
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 

## Methods

### NewPageResponseDtoVirtualDeviceType

`func NewPageResponseDtoVirtualDeviceType() *PageResponseDtoVirtualDeviceType`

NewPageResponseDtoVirtualDeviceType instantiates a new PageResponseDtoVirtualDeviceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageResponseDtoVirtualDeviceTypeWithDefaults

`func NewPageResponseDtoVirtualDeviceTypeWithDefaults() *PageResponseDtoVirtualDeviceType`

NewPageResponseDtoVirtualDeviceTypeWithDefaults instantiates a new PageResponseDtoVirtualDeviceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PageResponseDtoVirtualDeviceType) GetData() []VirtualDeviceType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageResponseDtoVirtualDeviceType) GetDataOk() (*[]VirtualDeviceType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageResponseDtoVirtualDeviceType) SetData(v []VirtualDeviceType)`

SetData sets Data field to given value.

### HasData

`func (o *PageResponseDtoVirtualDeviceType) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *PageResponseDtoVirtualDeviceType) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PageResponseDtoVirtualDeviceType) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PageResponseDtoVirtualDeviceType) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PageResponseDtoVirtualDeviceType) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


