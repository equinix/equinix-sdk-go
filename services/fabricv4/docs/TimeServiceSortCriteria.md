# TimeServiceSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**TimeServiceSortDirection**](TimeServiceSortDirection.md) |  | [optional] [default to TIMESERVICESORTDIRECTION_DESC]
**Property** | Pointer to [**TimeServiceSortBy**](TimeServiceSortBy.md) |  | [optional] [default to TIMESERVICESORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewTimeServiceSortCriteria

`func NewTimeServiceSortCriteria() *TimeServiceSortCriteria`

NewTimeServiceSortCriteria instantiates a new TimeServiceSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeServiceSortCriteriaWithDefaults

`func NewTimeServiceSortCriteriaWithDefaults() *TimeServiceSortCriteria`

NewTimeServiceSortCriteriaWithDefaults instantiates a new TimeServiceSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *TimeServiceSortCriteria) GetDirection() TimeServiceSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *TimeServiceSortCriteria) GetDirectionOk() (*TimeServiceSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *TimeServiceSortCriteria) SetDirection(v TimeServiceSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *TimeServiceSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *TimeServiceSortCriteria) GetProperty() TimeServiceSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *TimeServiceSortCriteria) GetPropertyOk() (*TimeServiceSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *TimeServiceSortCriteria) SetProperty(v TimeServiceSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *TimeServiceSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


