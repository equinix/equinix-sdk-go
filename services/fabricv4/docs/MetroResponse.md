# MetroResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]Metro**](Metro.md) | List of Fabric Metros. | [optional] 

## Methods

### NewMetroResponse

`func NewMetroResponse() *MetroResponse`

NewMetroResponse instantiates a new MetroResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetroResponseWithDefaults

`func NewMetroResponseWithDefaults() *MetroResponse`

NewMetroResponseWithDefaults instantiates a new MetroResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *MetroResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *MetroResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *MetroResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *MetroResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *MetroResponse) GetData() []Metro`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetroResponse) GetDataOk() (*[]Metro, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetroResponse) SetData(v []Metro)`

SetData sets Data field to given value.

### HasData

`func (o *MetroResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


