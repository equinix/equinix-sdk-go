# SshUserPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**PaginationResponseDto**](PaginationResponseDto.md) |  | [optional] 
**Data** | Pointer to [**[]SshUserInfoVerbose**](SshUserInfoVerbose.md) |  | [optional] 

## Methods

### NewSshUserPageResponse

`func NewSshUserPageResponse() *SshUserPageResponse`

NewSshUserPageResponse instantiates a new SshUserPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshUserPageResponseWithDefaults

`func NewSshUserPageResponseWithDefaults() *SshUserPageResponse`

NewSshUserPageResponseWithDefaults instantiates a new SshUserPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SshUserPageResponse) GetPagination() PaginationResponseDto`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SshUserPageResponse) GetPaginationOk() (*PaginationResponseDto, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SshUserPageResponse) SetPagination(v PaginationResponseDto)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SshUserPageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *SshUserPageResponse) GetData() []SshUserInfoVerbose`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SshUserPageResponse) GetDataOk() (*[]SshUserInfoVerbose, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SshUserPageResponse) SetData(v []SshUserInfoVerbose)`

SetData sets Data field to given value.

### HasData

`func (o *SshUserPageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


