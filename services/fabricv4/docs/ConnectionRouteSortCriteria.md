# ConnectionRouteSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ConnectionRouteEntrySortDirection**](ConnectionRouteEntrySortDirection.md) |  | [optional] [default to CONNECTIONROUTEENTRYSORTDIRECTION_DESC]
**Property** | Pointer to [**ConnectionRouteEntrySortBy**](ConnectionRouteEntrySortBy.md) |  | [optional] [default to CONNECTIONROUTEENTRYSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewConnectionRouteSortCriteria

`func NewConnectionRouteSortCriteria() *ConnectionRouteSortCriteria`

NewConnectionRouteSortCriteria instantiates a new ConnectionRouteSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRouteSortCriteriaWithDefaults

`func NewConnectionRouteSortCriteriaWithDefaults() *ConnectionRouteSortCriteria`

NewConnectionRouteSortCriteriaWithDefaults instantiates a new ConnectionRouteSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ConnectionRouteSortCriteria) GetDirection() ConnectionRouteEntrySortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ConnectionRouteSortCriteria) GetDirectionOk() (*ConnectionRouteEntrySortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ConnectionRouteSortCriteria) SetDirection(v ConnectionRouteEntrySortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ConnectionRouteSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *ConnectionRouteSortCriteria) GetProperty() ConnectionRouteEntrySortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ConnectionRouteSortCriteria) GetPropertyOk() (*ConnectionRouteEntrySortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ConnectionRouteSortCriteria) SetProperty(v ConnectionRouteEntrySortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ConnectionRouteSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


