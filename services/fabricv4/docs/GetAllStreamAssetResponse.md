# GetAllStreamAssetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]StreamAsset**](StreamAsset.md) | Data returned from the API call. | [optional] 

## Methods

### NewGetAllStreamAssetResponse

`func NewGetAllStreamAssetResponse() *GetAllStreamAssetResponse`

NewGetAllStreamAssetResponse instantiates a new GetAllStreamAssetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllStreamAssetResponseWithDefaults

`func NewGetAllStreamAssetResponseWithDefaults() *GetAllStreamAssetResponse`

NewGetAllStreamAssetResponseWithDefaults instantiates a new GetAllStreamAssetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetAllStreamAssetResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAllStreamAssetResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAllStreamAssetResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAllStreamAssetResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetAllStreamAssetResponse) GetData() []StreamAsset`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllStreamAssetResponse) GetDataOk() (*[]StreamAsset, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllStreamAssetResponse) SetData(v []StreamAsset)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllStreamAssetResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


