# ServiceTokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ServiceToken**](ServiceToken.md) | List of Service Tokens | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewServiceTokens

`func NewServiceTokens() *ServiceTokens`

NewServiceTokens instantiates a new ServiceTokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokensWithDefaults

`func NewServiceTokensWithDefaults() *ServiceTokens`

NewServiceTokensWithDefaults instantiates a new ServiceTokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ServiceTokens) GetData() []ServiceToken`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ServiceTokens) GetDataOk() (*[]ServiceToken, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ServiceTokens) SetData(v []ServiceToken)`

SetData sets Data field to given value.

### HasData

`func (o *ServiceTokens) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ServiceTokens) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ServiceTokens) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ServiceTokens) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ServiceTokens) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


