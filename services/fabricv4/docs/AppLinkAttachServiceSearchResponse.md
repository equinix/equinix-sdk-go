# AppLinkAttachServiceSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]AppLinkAppServiceAttachment**](AppLinkAppServiceAttachment.md) | Data returned from the API call. | [optional] 

## Methods

### NewAppLinkAttachServiceSearchResponse

`func NewAppLinkAttachServiceSearchResponse() *AppLinkAttachServiceSearchResponse`

NewAppLinkAttachServiceSearchResponse instantiates a new AppLinkAttachServiceSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAttachServiceSearchResponseWithDefaults

`func NewAppLinkAttachServiceSearchResponseWithDefaults() *AppLinkAttachServiceSearchResponse`

NewAppLinkAttachServiceSearchResponseWithDefaults instantiates a new AppLinkAttachServiceSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *AppLinkAttachServiceSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AppLinkAttachServiceSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AppLinkAttachServiceSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AppLinkAttachServiceSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *AppLinkAttachServiceSearchResponse) GetData() []AppLinkAppServiceAttachment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppLinkAttachServiceSearchResponse) GetDataOk() (*[]AppLinkAppServiceAttachment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppLinkAttachServiceSearchResponse) SetData(v []AppLinkAppServiceAttachment)`

SetData sets Data field to given value.

### HasData

`func (o *AppLinkAttachServiceSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


