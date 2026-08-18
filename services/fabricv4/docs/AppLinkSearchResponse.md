# AppLinkSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AppLink**](AppLink.md) | Data returned from the API call. | [optional] 

## Methods

### NewAppLinkSearchResponse

`func NewAppLinkSearchResponse() *AppLinkSearchResponse`

NewAppLinkSearchResponse instantiates a new AppLinkSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkSearchResponseWithDefaults

`func NewAppLinkSearchResponseWithDefaults() *AppLinkSearchResponse`

NewAppLinkSearchResponseWithDefaults instantiates a new AppLinkSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AppLinkSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppLinkSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppLinkSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppLinkSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AppLinkSearchResponse) GetData() []AppLink`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppLinkSearchResponse) GetDataOk() (*[]AppLink, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppLinkSearchResponse) SetData(v []AppLink)`

SetData sets Data field to given value.

### HasData

`func (o *AppLinkSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


