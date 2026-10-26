# AppProcessorSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AppProcessor**](AppProcessor.md) | Data returned from the API call. | [optional] 

## Methods

### NewAppProcessorSearchResponse

`func NewAppProcessorSearchResponse() *AppProcessorSearchResponse`

NewAppProcessorSearchResponse instantiates a new AppProcessorSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppProcessorSearchResponseWithDefaults

`func NewAppProcessorSearchResponseWithDefaults() *AppProcessorSearchResponse`

NewAppProcessorSearchResponseWithDefaults instantiates a new AppProcessorSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AppProcessorSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppProcessorSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppProcessorSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppProcessorSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AppProcessorSearchResponse) GetData() []AppProcessor`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppProcessorSearchResponse) GetDataOk() (*[]AppProcessor, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppProcessorSearchResponse) SetData(v []AppProcessor)`

SetData sets Data field to given value.

### HasData

`func (o *AppProcessorSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


