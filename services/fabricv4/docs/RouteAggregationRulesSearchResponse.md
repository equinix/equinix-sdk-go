# RouteAggregationRulesSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]RouteAggregationRulesData**](RouteAggregationRulesData.md) | List of route aggregation rules | [optional] 

## Methods

### NewRouteAggregationRulesSearchResponse

`func NewRouteAggregationRulesSearchResponse() *RouteAggregationRulesSearchResponse`

NewRouteAggregationRulesSearchResponse instantiates a new RouteAggregationRulesSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteAggregationRulesSearchResponseWithDefaults

`func NewRouteAggregationRulesSearchResponseWithDefaults() *RouteAggregationRulesSearchResponse`

NewRouteAggregationRulesSearchResponseWithDefaults instantiates a new RouteAggregationRulesSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RouteAggregationRulesSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteAggregationRulesSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteAggregationRulesSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteAggregationRulesSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *RouteAggregationRulesSearchResponse) GetData() []RouteAggregationRulesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteAggregationRulesSearchResponse) GetDataOk() (*[]RouteAggregationRulesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteAggregationRulesSearchResponse) SetData(v []RouteAggregationRulesData)`

SetData sets Data field to given value.

### HasData

`func (o *RouteAggregationRulesSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


