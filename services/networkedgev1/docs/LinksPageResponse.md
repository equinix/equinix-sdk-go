# LinksPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 
**Data** | Pointer to [**[]DeviceLinkGroupDto**](DeviceLinkGroupDto.md) |  | [optional] 

## Methods

### NewLinksPageResponse

`func NewLinksPageResponse() *LinksPageResponse`

NewLinksPageResponse instantiates a new LinksPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinksPageResponseWithDefaults

`func NewLinksPageResponseWithDefaults() *LinksPageResponse`

NewLinksPageResponseWithDefaults instantiates a new LinksPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *LinksPageResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LinksPageResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LinksPageResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LinksPageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *LinksPageResponse) GetData() []DeviceLinkGroupDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinksPageResponse) GetDataOk() (*[]DeviceLinkGroupDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinksPageResponse) SetData(v []DeviceLinkGroupDto)`

SetData sets Data field to given value.

### HasData

`func (o *LinksPageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


