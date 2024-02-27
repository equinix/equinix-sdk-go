# SortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] [default to SORTDIRECTION_DESC]
**Property** | Pointer to [**SortBy**](SortBy.md) |  | [optional] [default to SORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewSortCriteria

`func NewSortCriteria() *SortCriteria`

NewSortCriteria instantiates a new SortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortCriteriaWithDefaults

`func NewSortCriteriaWithDefaults() *SortCriteria`

NewSortCriteriaWithDefaults instantiates a new SortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SortCriteria) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SortCriteria) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SortCriteria) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *SortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *SortCriteria) GetProperty() SortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SortCriteria) GetPropertyOk() (*SortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SortCriteria) SetProperty(v SortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *SortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


