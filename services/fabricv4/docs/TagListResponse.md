# TagListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TagResponse**](TagResponse.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewTagListResponse

`func NewTagListResponse() *TagListResponse`

NewTagListResponse instantiates a new TagListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagListResponseWithDefaults

`func NewTagListResponseWithDefaults() *TagListResponse`

NewTagListResponseWithDefaults instantiates a new TagListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TagListResponse) GetData() []TagResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TagListResponse) GetDataOk() (*[]TagResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TagListResponse) SetData(v []TagResponse)`

SetData sets Data field to given value.

### HasData

`func (o *TagListResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *TagListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *TagListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *TagListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *TagListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


