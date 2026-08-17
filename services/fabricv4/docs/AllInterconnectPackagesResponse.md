# AllInterconnectPackagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | [**[]InterconnectPackage**](InterconnectPackage.md) | List of Interconnect Packages | 

## Methods

### NewAllInterconnectPackagesResponse

`func NewAllInterconnectPackagesResponse(data []InterconnectPackage, ) *AllInterconnectPackagesResponse`

NewAllInterconnectPackagesResponse instantiates a new AllInterconnectPackagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllInterconnectPackagesResponseWithDefaults

`func NewAllInterconnectPackagesResponseWithDefaults() *AllInterconnectPackagesResponse`

NewAllInterconnectPackagesResponseWithDefaults instantiates a new AllInterconnectPackagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AllInterconnectPackagesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AllInterconnectPackagesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AllInterconnectPackagesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AllInterconnectPackagesResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AllInterconnectPackagesResponse) GetData() []InterconnectPackage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AllInterconnectPackagesResponse) GetDataOk() (*[]InterconnectPackage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AllInterconnectPackagesResponse) SetData(v []InterconnectPackage)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


