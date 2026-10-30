# AsnSearchResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Sort** | Pointer to [**[]AsnSearchSortCriteria**](AsnSearchSortCriteria.md) |  | [optional] 
**Data** | [**[]Asn**](Asn.md) |  | 

## Methods

### NewAsnSearchResponseBody

`func NewAsnSearchResponseBody(pagination Pagination, data []Asn, ) *AsnSearchResponseBody`

NewAsnSearchResponseBody instantiates a new AsnSearchResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAsnSearchResponseBodyWithDefaults

`func NewAsnSearchResponseBodyWithDefaults() *AsnSearchResponseBody`

NewAsnSearchResponseBodyWithDefaults instantiates a new AsnSearchResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AsnSearchResponseBody) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AsnSearchResponseBody) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AsnSearchResponseBody) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetSort

`func (o *AsnSearchResponseBody) GetSort() []AsnSearchSortCriteria`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *AsnSearchResponseBody) GetSortOk() (*[]AsnSearchSortCriteria, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *AsnSearchResponseBody) SetSort(v []AsnSearchSortCriteria)`

SetSort sets Sort field to given value.

### HasSort

`func (o *AsnSearchResponseBody) HasSort() bool`

HasSort returns a boolean if a field has been set.

### GetData

`func (o *AsnSearchResponseBody) GetData() []Asn`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AsnSearchResponseBody) GetDataOk() (*[]Asn, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AsnSearchResponseBody) SetData(v []Asn)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


