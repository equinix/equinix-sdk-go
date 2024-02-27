# PackageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]CloudRouterPackage**](CloudRouterPackage.md) | Data returned from the API call. | [optional] 

## Methods

### NewPackageResponse

`func NewPackageResponse() *PackageResponse`

NewPackageResponse instantiates a new PackageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageResponseWithDefaults

`func NewPackageResponseWithDefaults() *PackageResponse`

NewPackageResponseWithDefaults instantiates a new PackageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *PackageResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *PackageResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *PackageResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *PackageResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *PackageResponse) GetData() []CloudRouterPackage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PackageResponse) GetDataOk() (*[]CloudRouterPackage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PackageResponse) SetData(v []CloudRouterPackage)`

SetData sets Data field to given value.

### HasData

`func (o *PackageResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


