# DeviceACLPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 
**Data** | Pointer to [**[]DeviceACLTemplatesResponse**](DeviceACLTemplatesResponse.md) |  | [optional] 

## Methods

### NewDeviceACLPageResponse

`func NewDeviceACLPageResponse() *DeviceACLPageResponse`

NewDeviceACLPageResponse instantiates a new DeviceACLPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceACLPageResponseWithDefaults

`func NewDeviceACLPageResponseWithDefaults() *DeviceACLPageResponse`

NewDeviceACLPageResponseWithDefaults instantiates a new DeviceACLPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *DeviceACLPageResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *DeviceACLPageResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *DeviceACLPageResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *DeviceACLPageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *DeviceACLPageResponse) GetData() []DeviceACLTemplatesResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DeviceACLPageResponse) GetDataOk() (*[]DeviceACLTemplatesResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DeviceACLPageResponse) SetData(v []DeviceACLTemplatesResponse)`

SetData sets Data field to given value.

### HasData

`func (o *DeviceACLPageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


