# GetAllStreamSubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]StreamSubscription**](StreamSubscription.md) | Data returned from the API call. | [optional] 

## Methods

### NewGetAllStreamSubscriptionResponse

`func NewGetAllStreamSubscriptionResponse() *GetAllStreamSubscriptionResponse`

NewGetAllStreamSubscriptionResponse instantiates a new GetAllStreamSubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllStreamSubscriptionResponseWithDefaults

`func NewGetAllStreamSubscriptionResponseWithDefaults() *GetAllStreamSubscriptionResponse`

NewGetAllStreamSubscriptionResponseWithDefaults instantiates a new GetAllStreamSubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GetAllStreamSubscriptionResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAllStreamSubscriptionResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAllStreamSubscriptionResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAllStreamSubscriptionResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *GetAllStreamSubscriptionResponse) GetData() []StreamSubscription`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllStreamSubscriptionResponse) GetDataOk() (*[]StreamSubscription, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllStreamSubscriptionResponse) SetData(v []StreamSubscription)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllStreamSubscriptionResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


