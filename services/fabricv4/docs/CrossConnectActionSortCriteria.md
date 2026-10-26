# CrossConnectActionSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**CrossConnectSortCriteriaDirection**](CrossConnectSortCriteriaDirection.md) |  | [optional] [default to CROSSCONNECTSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to [**CrossConnectActionSortCriteriaProperty**](CrossConnectActionSortCriteriaProperty.md) |  | [optional] [default to CROSSCONNECTACTIONSORTCRITERIAPROPERTY_CHANGE_LOG_CREATED_DATE_TIME]

## Methods

### NewCrossConnectActionSortCriteria

`func NewCrossConnectActionSortCriteria() *CrossConnectActionSortCriteria`

NewCrossConnectActionSortCriteria instantiates a new CrossConnectActionSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionSortCriteriaWithDefaults

`func NewCrossConnectActionSortCriteriaWithDefaults() *CrossConnectActionSortCriteria`

NewCrossConnectActionSortCriteriaWithDefaults instantiates a new CrossConnectActionSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *CrossConnectActionSortCriteria) GetDirection() CrossConnectSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CrossConnectActionSortCriteria) GetDirectionOk() (*CrossConnectSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CrossConnectActionSortCriteria) SetDirection(v CrossConnectSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CrossConnectActionSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *CrossConnectActionSortCriteria) GetProperty() CrossConnectActionSortCriteriaProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *CrossConnectActionSortCriteria) GetPropertyOk() (*CrossConnectActionSortCriteriaProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *CrossConnectActionSortCriteria) SetProperty(v CrossConnectActionSortCriteriaProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *CrossConnectActionSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


