# LinkProtocolGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]LinkProtocolResponse**](LinkProtocolResponse.md) |  | [optional] 

## Methods

### NewLinkProtocolGetResponse

`func NewLinkProtocolGetResponse() *LinkProtocolGetResponse`

NewLinkProtocolGetResponse instantiates a new LinkProtocolGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolGetResponseWithDefaults

`func NewLinkProtocolGetResponseWithDefaults() *LinkProtocolGetResponse`

NewLinkProtocolGetResponseWithDefaults instantiates a new LinkProtocolGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *LinkProtocolGetResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *LinkProtocolGetResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *LinkProtocolGetResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *LinkProtocolGetResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *LinkProtocolGetResponse) GetData() []LinkProtocolResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LinkProtocolGetResponse) GetDataOk() (*[]LinkProtocolResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LinkProtocolGetResponse) SetData(v []LinkProtocolResponse)`

SetData sets Data field to given value.

### HasData

`func (o *LinkProtocolGetResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


