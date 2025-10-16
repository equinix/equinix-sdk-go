# DeviceRMAInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]RmaDetailObject**](RmaDetailObject.md) | Array of previous RMA requests | [optional] 
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 

## Methods

### NewDeviceRMAInfo

`func NewDeviceRMAInfo() *DeviceRMAInfo`

NewDeviceRMAInfo instantiates a new DeviceRMAInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRMAInfoWithDefaults

`func NewDeviceRMAInfoWithDefaults() *DeviceRMAInfo`

NewDeviceRMAInfoWithDefaults instantiates a new DeviceRMAInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DeviceRMAInfo) GetData() []RmaDetailObject`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceRMAInfo) GetDataOk() (*[]RmaDetailObject, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceRMAInfo) SetData(v []RmaDetailObject)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceRMAInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *DeviceRMAInfo) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DeviceRMAInfo) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DeviceRMAInfo) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DeviceRMAInfo) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


