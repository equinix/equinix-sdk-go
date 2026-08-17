# OpticalConnectSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**OpticalConnectSortDirection**](OpticalConnectSortDirection.md) |  | [optional] [default to OPTICALCONNECTSORTDIRECTION_DESC]
**Property** | Pointer to [**OpticalConnectSortBy**](OpticalConnectSortBy.md) |  | [optional] [default to OPTICALCONNECTSORTBY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewOpticalConnectSortCriteria

`func NewOpticalConnectSortCriteria() *OpticalConnectSortCriteria`

NewOpticalConnectSortCriteria instantiates a new OpticalConnectSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectSortCriteriaWithDefaults

`func NewOpticalConnectSortCriteriaWithDefaults() *OpticalConnectSortCriteria`

NewOpticalConnectSortCriteriaWithDefaults instantiates a new OpticalConnectSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *OpticalConnectSortCriteria) GetDirection() OpticalConnectSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *OpticalConnectSortCriteria) GetDirectionOk() (*OpticalConnectSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *OpticalConnectSortCriteria) SetDirection(v OpticalConnectSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *OpticalConnectSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *OpticalConnectSortCriteria) GetProperty() OpticalConnectSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *OpticalConnectSortCriteria) GetPropertyOk() (*OpticalConnectSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *OpticalConnectSortCriteria) SetProperty(v OpticalConnectSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *OpticalConnectSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


