# InternetAccessServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Sort** | [**[]SearchSortItem**](SearchSortItem.md) |  | 
**Data** | [**[]InternetAccessService**](InternetAccessService.md) |  | 

## Methods

### NewInternetAccessServices

`func NewInternetAccessServices(pagination Pagination, sort []SearchSortItem, data []InternetAccessService, ) *InternetAccessServices`

NewInternetAccessServices instantiates a new InternetAccessServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetAccessServicesWithDefaults

`func NewInternetAccessServicesWithDefaults() *InternetAccessServices`

NewInternetAccessServicesWithDefaults instantiates a new InternetAccessServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *InternetAccessServices) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *InternetAccessServices) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *InternetAccessServices) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetSort

`func (o *InternetAccessServices) GetSort() []SearchSortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *InternetAccessServices) GetSortOk() (*[]SearchSortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *InternetAccessServices) SetSort(v []SearchSortItem)`

SetSort sets Sort field to given value.


### GetData

`func (o *InternetAccessServices) GetData() []InternetAccessService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternetAccessServices) GetDataOk() (*[]InternetAccessService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternetAccessServices) SetData(v []InternetAccessService)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


