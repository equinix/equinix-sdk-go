# InterconnectSortCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**InterconnectSortDirectionResponse**](InterconnectSortDirectionResponse.md) |  | [optional] [default to INTERCONNECTSORTDIRECTIONRESPONSE_DESC]
**Property** | Pointer to [**InterconnectSortByResponse**](InterconnectSortByResponse.md) |  | [optional] [default to INTERCONNECTSORTBYRESPONSE_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewInterconnectSortCriteriaResponse

`func NewInterconnectSortCriteriaResponse() *InterconnectSortCriteriaResponse`

NewInterconnectSortCriteriaResponse instantiates a new InterconnectSortCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectSortCriteriaResponseWithDefaults

`func NewInterconnectSortCriteriaResponseWithDefaults() *InterconnectSortCriteriaResponse`

NewInterconnectSortCriteriaResponseWithDefaults instantiates a new InterconnectSortCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *InterconnectSortCriteriaResponse) GetDirection() InterconnectSortDirectionResponse`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InterconnectSortCriteriaResponse) GetDirectionOk() (*InterconnectSortDirectionResponse, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InterconnectSortCriteriaResponse) SetDirection(v InterconnectSortDirectionResponse)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InterconnectSortCriteriaResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *InterconnectSortCriteriaResponse) GetProperty() InterconnectSortByResponse`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *InterconnectSortCriteriaResponse) GetPropertyOk() (*InterconnectSortByResponse, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *InterconnectSortCriteriaResponse) SetProperty(v InterconnectSortByResponse)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *InterconnectSortCriteriaResponse) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


