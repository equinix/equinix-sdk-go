# RouteTableEntrySearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]RouteTableEntry**](RouteTableEntry.md) | Data returned from the API call. | [optional] 

## Methods

### NewRouteTableEntrySearchResponse

`func NewRouteTableEntrySearchResponse() *RouteTableEntrySearchResponse`

NewRouteTableEntrySearchResponse instantiates a new RouteTableEntrySearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableEntrySearchResponseWithDefaults

`func NewRouteTableEntrySearchResponseWithDefaults() *RouteTableEntrySearchResponse`

NewRouteTableEntrySearchResponseWithDefaults instantiates a new RouteTableEntrySearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *RouteTableEntrySearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RouteTableEntrySearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RouteTableEntrySearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *RouteTableEntrySearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *RouteTableEntrySearchResponse) GetData() []RouteTableEntry`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteTableEntrySearchResponse) GetDataOk() (*[]RouteTableEntry, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteTableEntrySearchResponse) SetData(v []RouteTableEntry)`

SetData sets Data field to given value.

### HasData

`func (o *RouteTableEntrySearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


