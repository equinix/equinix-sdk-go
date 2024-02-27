# GetRouteFilterRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]RouteFilterRulesData**](RouteFilterRulesData.md) | List of Route Filter Rules | [optional] 

## Methods

### NewGetRouteFilterRulesResponse

`func NewGetRouteFilterRulesResponse() *GetRouteFilterRulesResponse`

NewGetRouteFilterRulesResponse instantiates a new GetRouteFilterRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRouteFilterRulesResponseWithDefaults

`func NewGetRouteFilterRulesResponseWithDefaults() *GetRouteFilterRulesResponse`

NewGetRouteFilterRulesResponseWithDefaults instantiates a new GetRouteFilterRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetRouteFilterRulesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetRouteFilterRulesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetRouteFilterRulesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetRouteFilterRulesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetRouteFilterRulesResponse) GetData() []RouteFilterRulesData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetRouteFilterRulesResponse) GetDataOk() (*[]RouteFilterRulesData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetRouteFilterRulesResponse) SetData(v []RouteFilterRulesData)`

SetData sets Data field to given value.

### HasData

`func (o *GetRouteFilterRulesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


