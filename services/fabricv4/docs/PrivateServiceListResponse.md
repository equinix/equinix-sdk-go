# PrivateServiceListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]PrivateService**](PrivateService.md) | List of private services | [optional] 

## Methods

### NewPrivateServiceListResponse

`func NewPrivateServiceListResponse() *PrivateServiceListResponse`

NewPrivateServiceListResponse instantiates a new PrivateServiceListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateServiceListResponseWithDefaults

`func NewPrivateServiceListResponseWithDefaults() *PrivateServiceListResponse`

NewPrivateServiceListResponseWithDefaults instantiates a new PrivateServiceListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PrivateServiceListResponse) GetData() []PrivateService`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PrivateServiceListResponse) GetDataOk() (*[]PrivateService, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PrivateServiceListResponse) SetData(v []PrivateService)`

SetData sets Data field to given value.

### HasData

`func (o *PrivateServiceListResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


