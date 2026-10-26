# CrossConnectScheduleSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**CrossConnectSortCriteriaDirection**](CrossConnectSortCriteriaDirection.md) |  | [optional] [default to CROSSCONNECTSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to [**CrossConnectScheduleSortCriteriaProperty**](CrossConnectScheduleSortCriteriaProperty.md) |  | [optional] [default to CROSSCONNECTSCHEDULESORTCRITERIAPROPERTY_START_DATE_TIME]

## Methods

### NewCrossConnectScheduleSortCriteria

`func NewCrossConnectScheduleSortCriteria() *CrossConnectScheduleSortCriteria`

NewCrossConnectScheduleSortCriteria instantiates a new CrossConnectScheduleSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectScheduleSortCriteriaWithDefaults

`func NewCrossConnectScheduleSortCriteriaWithDefaults() *CrossConnectScheduleSortCriteria`

NewCrossConnectScheduleSortCriteriaWithDefaults instantiates a new CrossConnectScheduleSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CrossConnectScheduleSortCriteria) GetDirection() CrossConnectSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CrossConnectScheduleSortCriteria) GetDirectionOk() (*CrossConnectSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CrossConnectScheduleSortCriteria) SetDirection(v CrossConnectSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CrossConnectScheduleSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *CrossConnectScheduleSortCriteria) GetProperty() CrossConnectScheduleSortCriteriaProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CrossConnectScheduleSortCriteria) GetPropertyOk() (*CrossConnectScheduleSortCriteriaProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CrossConnectScheduleSortCriteria) SetProperty(v CrossConnectScheduleSortCriteriaProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CrossConnectScheduleSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


