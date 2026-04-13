# RouteFilterRulesSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]RouteFilterRulesData**](RouteFilterRulesData.md) | List of route filter rules | [optional] 

## Methods

### NewRouteFilterRulesSearchResponse

`func NewRouteFilterRulesSearchResponse() *RouteFilterRulesSearchResponse`

NewRouteFilterRulesSearchResponse instantiates a new RouteFilterRulesSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesSearchResponseWithDefaults

`func NewRouteFilterRulesSearchResponseWithDefaults() *RouteFilterRulesSearchResponse`

NewRouteFilterRulesSearchResponseWithDefaults instantiates a new RouteFilterRulesSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RouteFilterRulesSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteFilterRulesSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteFilterRulesSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteFilterRulesSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *RouteFilterRulesSearchResponse) GetData() []RouteFilterRulesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteFilterRulesSearchResponse) GetDataOk() (*[]RouteFilterRulesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteFilterRulesSearchResponse) SetData(v []RouteFilterRulesData)`

SetData sets Data field to given value.

### HasData

`func (o *RouteFilterRulesSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


