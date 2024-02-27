# NetworkSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**NetworkSortDirection**](NetworkSortDirection.md) |  | [optional] [default to NETWORKSORTDIRECTION_DESC]
**Property** | Pointer to [**NetworkSortBy**](NetworkSortBy.md) |  | [optional] [default to NETWORKSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewNetworkSortCriteria

`func NewNetworkSortCriteria() *NetworkSortCriteria`

NewNetworkSortCriteria instantiates a new NetworkSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkSortCriteriaWithDefaults

`func NewNetworkSortCriteriaWithDefaults() *NetworkSortCriteria`

NewNetworkSortCriteriaWithDefaults instantiates a new NetworkSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *NetworkSortCriteria) GetDirection() NetworkSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *NetworkSortCriteria) GetDirectionOk() (*NetworkSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *NetworkSortCriteria) SetDirection(v NetworkSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *NetworkSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *NetworkSortCriteria) GetProperty() NetworkSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *NetworkSortCriteria) GetPropertyOk() (*NetworkSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *NetworkSortCriteria) SetProperty(v NetworkSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *NetworkSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


