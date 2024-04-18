# GETConnectionsPageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLastPage** | Pointer to **bool** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 
**IsFirstPage** | Pointer to **bool** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**PageNumber** | Pointer to **int32** |  | [optional] 
**Content** | Pointer to [**[]GETConnectionByUuidResponse**](GETConnectionByUuidResponse.md) |  | [optional] 

## Methods

### NewGETConnectionsPageResponse

`func NewGETConnectionsPageResponse() *GETConnectionsPageResponse`

NewGETConnectionsPageResponse instantiates a new GETConnectionsPageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGETConnectionsPageResponseWithDefaults

`func NewGETConnectionsPageResponseWithDefaults() *GETConnectionsPageResponse`

NewGETConnectionsPageResponseWithDefaults instantiates a new GETConnectionsPageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLastPage

`func (o *GETConnectionsPageResponse) GetIsLastPage() bool`

GetIsLastPage returns the IsLastPage field if non-nil, zero value otherwise.

### GetIsLastPageOk

`func (o *GETConnectionsPageResponse) GetIsLastPageOk() (*bool, bool)`

GetIsLastPageOk returns a tuple with the IsLastPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLastPage

`func (o *GETConnectionsPageResponse) SetIsLastPage(v bool)`

SetIsLastPage sets IsLastPage field to given value.

### HasIsLastPage

`func (o *GETConnectionsPageResponse) HasIsLastPage() bool`

HasIsLastPage returns a boolean if a field has been set.

### GetTotalCount

`func (o *GETConnectionsPageResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GETConnectionsPageResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GETConnectionsPageResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *GETConnectionsPageResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetIsFirstPage

`func (o *GETConnectionsPageResponse) GetIsFirstPage() bool`

GetIsFirstPage returns the IsFirstPage field if non-nil, zero value otherwise.

### GetIsFirstPageOk

`func (o *GETConnectionsPageResponse) GetIsFirstPageOk() (*bool, bool)`

GetIsFirstPageOk returns a tuple with the IsFirstPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFirstPage

`func (o *GETConnectionsPageResponse) SetIsFirstPage(v bool)`

SetIsFirstPage sets IsFirstPage field to given value.

### HasIsFirstPage

`func (o *GETConnectionsPageResponse) HasIsFirstPage() bool`

HasIsFirstPage returns a boolean if a field has been set.

### GetPageSize

`func (o *GETConnectionsPageResponse) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *GETConnectionsPageResponse) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *GETConnectionsPageResponse) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *GETConnectionsPageResponse) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetPageNumber

`func (o *GETConnectionsPageResponse) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *GETConnectionsPageResponse) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *GETConnectionsPageResponse) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *GETConnectionsPageResponse) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetContent

`func (o *GETConnectionsPageResponse) GetContent() []GETConnectionByUuidResponse`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GETConnectionsPageResponse) GetContentOk() (*[]GETConnectionByUuidResponse, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GETConnectionsPageResponse) SetContent(v []GETConnectionByUuidResponse)`

SetContent sets Content field to given value.

### HasContent

`func (o *GETConnectionsPageResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


