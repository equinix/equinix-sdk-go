# OpticalConnectServiceSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Data** | Pointer to [**[]OpticalConnectResponse**](OpticalConnectResponse.md) | Data returned from the API call. | [optional] 

## Methods

### NewOpticalConnectServiceSearchResponse

`func NewOpticalConnectServiceSearchResponse() *OpticalConnectServiceSearchResponse`

NewOpticalConnectServiceSearchResponse instantiates a new OpticalConnectServiceSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectServiceSearchResponseWithDefaults

`func NewOpticalConnectServiceSearchResponseWithDefaults() *OpticalConnectServiceSearchResponse`

NewOpticalConnectServiceSearchResponseWithDefaults instantiates a new OpticalConnectServiceSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPagination

`func (o *OpticalConnectServiceSearchResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *OpticalConnectServiceSearchResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *OpticalConnectServiceSearchResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *OpticalConnectServiceSearchResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetData

`func (o *OpticalConnectServiceSearchResponse) GetData() []OpticalConnectResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OpticalConnectServiceSearchResponse) GetDataOk() (*[]OpticalConnectResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OpticalConnectServiceSearchResponse) SetData(v []OpticalConnectResponse)`

SetData sets Data field to given value.

### HasData

`func (o *OpticalConnectServiceSearchResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


