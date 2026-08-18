# LoaSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]LoaSortCriteria**](LoaSortCriteria.md) |  | [optional] 
**Data** | Pointer to [**[]LoaResponse**](LoaResponse.md) |  | [optional] 

## Methods

### NewLoaSearchResponse

`func NewLoaSearchResponse() *LoaSearchResponse`

NewLoaSearchResponse instantiates a new LoaSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaSearchResponseWithDefaults

`func NewLoaSearchResponseWithDefaults() *LoaSearchResponse`

NewLoaSearchResponseWithDefaults instantiates a new LoaSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *LoaSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LoaSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LoaSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LoaSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *LoaSearchResponse) GetSort() []LoaSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *LoaSearchResponse) GetSortOk() (*[]LoaSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *LoaSearchResponse) SetSort(v []LoaSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *LoaSearchResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *LoaSearchResponse) GetData() []LoaResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoaSearchResponse) GetDataOk() (*[]LoaResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoaSearchResponse) SetData(v []LoaResponse)`

SetData sets Data field to given value.

### HasData

`func (o *LoaSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


