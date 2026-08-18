# LoaActionSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]LoaActionSortCriteria**](LoaActionSortCriteria.md) |  | [optional] 
**Data** | Pointer to [**[]LoaActionResponse**](LoaActionResponse.md) |  | [optional] 

## Methods

### NewLoaActionSearchResponse

`func NewLoaActionSearchResponse() *LoaActionSearchResponse`

NewLoaActionSearchResponse instantiates a new LoaActionSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionSearchResponseWithDefaults

`func NewLoaActionSearchResponseWithDefaults() *LoaActionSearchResponse`

NewLoaActionSearchResponseWithDefaults instantiates a new LoaActionSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *LoaActionSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LoaActionSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LoaActionSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LoaActionSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *LoaActionSearchResponse) GetSort() []LoaActionSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoaActionSearchResponse) GetSortOk() (*[]LoaActionSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoaActionSearchResponse) SetSort(v []LoaActionSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoaActionSearchResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *LoaActionSearchResponse) GetData() []LoaActionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoaActionSearchResponse) GetDataOk() (*[]LoaActionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoaActionSearchResponse) SetData(v []LoaActionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *LoaActionSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


