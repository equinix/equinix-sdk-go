# RouteAggregationsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]RouteAggregationsData**](RouteAggregationsData.md) | List of Route Aggregations | [optional] 

## Methods

### NewRouteAggregationsSearchResponse

`func NewRouteAggregationsSearchResponse() *RouteAggregationsSearchResponse`

NewRouteAggregationsSearchResponse instantiates a new RouteAggregationsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationsSearchResponseWithDefaults

`func NewRouteAggregationsSearchResponseWithDefaults() *RouteAggregationsSearchResponse`

NewRouteAggregationsSearchResponseWithDefaults instantiates a new RouteAggregationsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RouteAggregationsSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteAggregationsSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteAggregationsSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteAggregationsSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *RouteAggregationsSearchResponse) GetData() []RouteAggregationsData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteAggregationsSearchResponse) GetDataOk() (*[]RouteAggregationsData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteAggregationsSearchResponse) SetData(v []RouteAggregationsData)`

SetData sets Data field to given value.

### HasData

`func (o *RouteAggregationsSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


