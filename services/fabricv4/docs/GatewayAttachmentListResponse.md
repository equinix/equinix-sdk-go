# GatewayAttachmentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | [**Pagination**](Pagination.md) |  | 
**Data** | [**[]GatewayAttachmentResponse**](GatewayAttachmentResponse.md) |  | 

## Methods

### NewGatewayAttachmentListResponse

`func NewGatewayAttachmentListResponse(pagination Pagination, data []GatewayAttachmentResponse, ) *GatewayAttachmentListResponse`

NewGatewayAttachmentListResponse instantiates a new GatewayAttachmentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayAttachmentListResponseWithDefaults

`func NewGatewayAttachmentListResponseWithDefaults() *GatewayAttachmentListResponse`

NewGatewayAttachmentListResponseWithDefaults instantiates a new GatewayAttachmentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *GatewayAttachmentListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GatewayAttachmentListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GatewayAttachmentListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetData

`func (o *GatewayAttachmentListResponse) GetData() []GatewayAttachmentResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GatewayAttachmentListResponse) GetDataOk() (*[]GatewayAttachmentResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GatewayAttachmentListResponse) SetData(v []GatewayAttachmentResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


