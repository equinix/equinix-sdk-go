# SortCriteriaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] [default to SORTDIRECTION_DESC]
**Property** | Pointer to [**SortBy**](SortBy.md) |  | [optional] [default to SORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewSortCriteriaResponse

`func NewSortCriteriaResponse() *SortCriteriaResponse`

NewSortCriteriaResponse instantiates a new SortCriteriaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortCriteriaResponseWithDefaults

`func NewSortCriteriaResponseWithDefaults() *SortCriteriaResponse`

NewSortCriteriaResponseWithDefaults instantiates a new SortCriteriaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SortCriteriaResponse) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SortCriteriaResponse) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SortCriteriaResponse) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SortCriteriaResponse) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *SortCriteriaResponse) GetProperty() SortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SortCriteriaResponse) GetPropertyOk() (*SortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SortCriteriaResponse) SetProperty(v SortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SortCriteriaResponse) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


