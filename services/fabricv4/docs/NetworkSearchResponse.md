# NetworkSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Sort** | Pointer to [**[]NetworkSortCriteriaResponse**](NetworkSortCriteriaResponse.md) |  | [optional] 
**Data** | Pointer to [**[]Network**](Network.md) |  | [optional] 

## Methods

### NewNetworkSearchResponse

`func NewNetworkSearchResponse() *NetworkSearchResponse`

NewNetworkSearchResponse instantiates a new NetworkSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSearchResponseWithDefaults

`func NewNetworkSearchResponseWithDefaults() *NetworkSearchResponse`

NewNetworkSearchResponseWithDefaults instantiates a new NetworkSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *NetworkSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *NetworkSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *NetworkSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *NetworkSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSort

`func (o *NetworkSearchResponse) GetSort() []NetworkSortCriteriaResponse`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *NetworkSearchResponse) GetSortOk() (*[]NetworkSortCriteriaResponse, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *NetworkSearchResponse) SetSort(v []NetworkSortCriteriaResponse)`

SetSort sets Sort field to given value.

### HasSort

`func (o *NetworkSearchResponse) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *NetworkSearchResponse) GetData() []Network`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NetworkSearchResponse) GetDataOk() (*[]Network, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NetworkSearchResponse) SetData(v []Network)`

SetData sets Data field to given value.

### HasData

`func (o *NetworkSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


