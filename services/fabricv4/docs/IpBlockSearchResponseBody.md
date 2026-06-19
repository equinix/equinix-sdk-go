# IpBlockSearchResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Sort** | Pointer to [**[]SearchSortItem**](SearchSortItem.md) |  | [optional] 
**Data** | [**[]IpBlock**](IpBlock.md) |  | 

## Methods

### NewIpBlockSearchResponseBody

`func NewIpBlockSearchResponseBody(pagination Pagination, data []IpBlock, ) *IpBlockSearchResponseBody`

NewIpBlockSearchResponseBody instantiates a new IpBlockSearchResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockSearchResponseBodyWithDefaults

`func NewIpBlockSearchResponseBodyWithDefaults() *IpBlockSearchResponseBody`

NewIpBlockSearchResponseBodyWithDefaults instantiates a new IpBlockSearchResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *IpBlockSearchResponseBody) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *IpBlockSearchResponseBody) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *IpBlockSearchResponseBody) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetSort

`func (o *IpBlockSearchResponseBody) GetSort() []SearchSortItem`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *IpBlockSearchResponseBody) GetSortOk() (*[]SearchSortItem, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *IpBlockSearchResponseBody) SetSort(v []SearchSortItem)`

SetSort sets Sort field to given value.

### HasSort

`func (o *IpBlockSearchResponseBody) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *IpBlockSearchResponseBody) GetData() []IpBlock`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IpBlockSearchResponseBody) GetDataOk() (*[]IpBlock, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IpBlockSearchResponseBody) SetData(v []IpBlock)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


