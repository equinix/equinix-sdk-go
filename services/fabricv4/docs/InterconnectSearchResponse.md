# InterconnectSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]InterconnectSortCriteriaResponse**](InterconnectSortCriteriaResponse.md) |  | [optional] 
**Data** | Pointer to [**[]Interconnect**](Interconnect.md) |  | [optional] 

## Methods

### NewInterconnectSearchResponse

`func NewInterconnectSearchResponse() *InterconnectSearchResponse`

NewInterconnectSearchResponse instantiates a new InterconnectSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectSearchResponseWithDefaults

`func NewInterconnectSearchResponseWithDefaults() *InterconnectSearchResponse`

NewInterconnectSearchResponseWithDefaults instantiates a new InterconnectSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *InterconnectSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InterconnectSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InterconnectSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *InterconnectSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *InterconnectSearchResponse) GetSort() []InterconnectSortCriteriaResponse`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *InterconnectSearchResponse) GetSortOk() (*[]InterconnectSortCriteriaResponse, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *InterconnectSearchResponse) SetSort(v []InterconnectSortCriteriaResponse)`

SetSort sets Sort field to given value.

### HasSort

`func (o *InterconnectSearchResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *InterconnectSearchResponse) GetData() []Interconnect`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InterconnectSearchResponse) GetDataOk() (*[]Interconnect, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InterconnectSearchResponse) SetData(v []Interconnect)`

SetData sets Data field to given value.

### HasData

`func (o *InterconnectSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


