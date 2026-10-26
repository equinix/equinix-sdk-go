# SearchSchedulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]CrossConnectScheduleSortCriteria**](CrossConnectScheduleSortCriteria.md) |  | [optional] 
**Data** | Pointer to [**[]CrossConnectScheduleAvailability**](CrossConnectScheduleAvailability.md) |  | [optional] [readonly] 

## Methods

### NewSearchSchedulesResponse

`func NewSearchSchedulesResponse() *SearchSchedulesResponse`

NewSearchSchedulesResponse instantiates a new SearchSchedulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSchedulesResponseWithDefaults

`func NewSearchSchedulesResponseWithDefaults() *SearchSchedulesResponse`

NewSearchSchedulesResponseWithDefaults instantiates a new SearchSchedulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *SearchSchedulesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *SearchSchedulesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *SearchSchedulesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *SearchSchedulesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *SearchSchedulesResponse) GetSort() []CrossConnectScheduleSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *SearchSchedulesResponse) GetSortOk() (*[]CrossConnectScheduleSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *SearchSchedulesResponse) SetSort(v []CrossConnectScheduleSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *SearchSchedulesResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *SearchSchedulesResponse) GetData() []CrossConnectScheduleAvailability`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SearchSchedulesResponse) GetDataOk() (*[]CrossConnectScheduleAvailability, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SearchSchedulesResponse) SetData(v []CrossConnectScheduleAvailability)`

SetData sets Data field to given value.

### HasData

`func (o *SearchSchedulesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


