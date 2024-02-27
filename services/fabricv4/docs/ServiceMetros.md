# ServiceMetros

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceMetro**](ServiceMetro.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewServiceMetros

`func NewServiceMetros() *ServiceMetros`

NewServiceMetros instantiates a new ServiceMetros object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceMetrosWithDefaults

`func NewServiceMetrosWithDefaults() *ServiceMetros`

NewServiceMetrosWithDefaults instantiates a new ServiceMetros object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceMetros) GetData() []ServiceMetro`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceMetros) GetDataOk() (*[]ServiceMetro, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceMetros) SetData(v []ServiceMetro)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceMetros) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ServiceMetros) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceMetros) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceMetros) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceMetros) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


