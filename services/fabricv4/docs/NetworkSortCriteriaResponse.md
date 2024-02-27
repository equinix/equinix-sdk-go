# NetworkSortCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**NetworkSortDirectionResponse**](NetworkSortDirectionResponse.md) |  | [optional] [default to NETWORKSORTDIRECTIONRESPONSE_DESC]
**Property** | Pointer to [**NetworkSortByResponse**](NetworkSortByResponse.md) |  | [optional] [default to NETWORKSORTBYRESPONSE_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewNetworkSortCriteriaResponse

`func NewNetworkSortCriteriaResponse() *NetworkSortCriteriaResponse`

NewNetworkSortCriteriaResponse instantiates a new NetworkSortCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSortCriteriaResponseWithDefaults

`func NewNetworkSortCriteriaResponseWithDefaults() *NetworkSortCriteriaResponse`

NewNetworkSortCriteriaResponseWithDefaults instantiates a new NetworkSortCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *NetworkSortCriteriaResponse) GetDirection() NetworkSortDirectionResponse`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkSortCriteriaResponse) GetDirectionOk() (*NetworkSortDirectionResponse, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkSortCriteriaResponse) SetDirection(v NetworkSortDirectionResponse)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkSortCriteriaResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *NetworkSortCriteriaResponse) GetProperty() NetworkSortByResponse`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *NetworkSortCriteriaResponse) GetPropertyOk() (*NetworkSortByResponse, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *NetworkSortCriteriaResponse) SetProperty(v NetworkSortByResponse)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *NetworkSortCriteriaResponse) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


