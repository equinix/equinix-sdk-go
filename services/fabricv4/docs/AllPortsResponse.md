# AllPortsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]Port**](Port.md) | GET All User Port Across Fabric Metros | [optional] 

## Methods

### NewAllPortsResponse

`func NewAllPortsResponse() *AllPortsResponse`

NewAllPortsResponse instantiates a new AllPortsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllPortsResponseWithDefaults

`func NewAllPortsResponseWithDefaults() *AllPortsResponse`

NewAllPortsResponseWithDefaults instantiates a new AllPortsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AllPortsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AllPortsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AllPortsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AllPortsResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AllPortsResponse) GetData() []Port`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AllPortsResponse) GetDataOk() (*[]Port, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AllPortsResponse) SetData(v []Port)`

SetData sets Data field to given value.

### HasData

`func (o *AllPortsResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


