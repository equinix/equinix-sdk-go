# CloudRouterActionsSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]CloudRouterActionResponse**](CloudRouterActionResponse.md) | Data returned from the API call. | [optional] 

## Methods

### NewCloudRouterActionsSearchResponse

`func NewCloudRouterActionsSearchResponse() *CloudRouterActionsSearchResponse`

NewCloudRouterActionsSearchResponse instantiates a new CloudRouterActionsSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRouterActionsSearchResponseWithDefaults

`func NewCloudRouterActionsSearchResponseWithDefaults() *CloudRouterActionsSearchResponse`

NewCloudRouterActionsSearchResponseWithDefaults instantiates a new CloudRouterActionsSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *CloudRouterActionsSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *CloudRouterActionsSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *CloudRouterActionsSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *CloudRouterActionsSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *CloudRouterActionsSearchResponse) GetData() []CloudRouterActionResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudRouterActionsSearchResponse) GetDataOk() (*[]CloudRouterActionResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudRouterActionsSearchResponse) SetData(v []CloudRouterActionResponse)`

SetData sets Data field to given value.

### HasData

`func (o *CloudRouterActionsSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


