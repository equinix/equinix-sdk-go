# RouteTableEntrySortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**RouteTableEntrySortDirection**](RouteTableEntrySortDirection.md) |  | [optional] [default to ROUTETABLEENTRYSORTDIRECTION_DESC]
**Property** | Pointer to [**RouteTableEntrySortBy**](RouteTableEntrySortBy.md) |  | [optional] [default to ROUTETABLEENTRYSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewRouteTableEntrySortCriteria

`func NewRouteTableEntrySortCriteria() *RouteTableEntrySortCriteria`

NewRouteTableEntrySortCriteria instantiates a new RouteTableEntrySortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTableEntrySortCriteriaWithDefaults

`func NewRouteTableEntrySortCriteriaWithDefaults() *RouteTableEntrySortCriteria`

NewRouteTableEntrySortCriteriaWithDefaults instantiates a new RouteTableEntrySortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *RouteTableEntrySortCriteria) GetDirection() RouteTableEntrySortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RouteTableEntrySortCriteria) GetDirectionOk() (*RouteTableEntrySortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RouteTableEntrySortCriteria) SetDirection(v RouteTableEntrySortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RouteTableEntrySortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *RouteTableEntrySortCriteria) GetProperty() RouteTableEntrySortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *RouteTableEntrySortCriteria) GetPropertyOk() (*RouteTableEntrySortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *RouteTableEntrySortCriteria) SetProperty(v RouteTableEntrySortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *RouteTableEntrySortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


