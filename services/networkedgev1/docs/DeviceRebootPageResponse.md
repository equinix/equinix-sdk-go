# DeviceRebootPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 
**Data** | Pointer to [**[]DeviceRebootResponse**](DeviceRebootResponse.md) |  | [optional] 

## Methods

### NewDeviceRebootPageResponse

`func NewDeviceRebootPageResponse() *DeviceRebootPageResponse`

NewDeviceRebootPageResponse instantiates a new DeviceRebootPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRebootPageResponseWithDefaults

`func NewDeviceRebootPageResponseWithDefaults() *DeviceRebootPageResponse`

NewDeviceRebootPageResponseWithDefaults instantiates a new DeviceRebootPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *DeviceRebootPageResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DeviceRebootPageResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DeviceRebootPageResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DeviceRebootPageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *DeviceRebootPageResponse) GetData() []DeviceRebootResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceRebootPageResponse) GetDataOk() (*[]DeviceRebootResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceRebootPageResponse) SetData(v []DeviceRebootResponse)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceRebootPageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


