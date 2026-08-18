# InterconnectSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**InterconnectSortDirection**](InterconnectSortDirection.md) |  | [optional] [default to INTERCONNECTSORTDIRECTION_DESC]
**Property** | Pointer to [**InterconnectSortBy**](InterconnectSortBy.md) |  | [optional] [default to INTERCONNECTSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewInterconnectSortCriteria

`func NewInterconnectSortCriteria() *InterconnectSortCriteria`

NewInterconnectSortCriteria instantiates a new InterconnectSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectSortCriteriaWithDefaults

`func NewInterconnectSortCriteriaWithDefaults() *InterconnectSortCriteria`

NewInterconnectSortCriteriaWithDefaults instantiates a new InterconnectSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *InterconnectSortCriteria) GetDirection() InterconnectSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InterconnectSortCriteria) GetDirectionOk() (*InterconnectSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InterconnectSortCriteria) SetDirection(v InterconnectSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InterconnectSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *InterconnectSortCriteria) GetProperty() InterconnectSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *InterconnectSortCriteria) GetPropertyOk() (*InterconnectSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *InterconnectSortCriteria) SetProperty(v InterconnectSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *InterconnectSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


