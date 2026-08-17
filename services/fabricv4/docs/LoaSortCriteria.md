# LoaSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ExchangeServiceSearchSortCriteriaDirection**](ExchangeServiceSearchSortCriteriaDirection.md) |  | [optional] [default to EXCHANGESERVICESEARCHSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to [**LoaSortCriteriaProperty**](LoaSortCriteriaProperty.md) |  | [optional] [default to LOASORTCRITERIAPROPERTY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewLoaSortCriteria

`func NewLoaSortCriteria() *LoaSortCriteria`

NewLoaSortCriteria instantiates a new LoaSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaSortCriteriaWithDefaults

`func NewLoaSortCriteriaWithDefaults() *LoaSortCriteria`

NewLoaSortCriteriaWithDefaults instantiates a new LoaSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *LoaSortCriteria) GetDirection() ExchangeServiceSearchSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LoaSortCriteria) GetDirectionOk() (*ExchangeServiceSearchSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LoaSortCriteria) SetDirection(v ExchangeServiceSearchSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LoaSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *LoaSortCriteria) GetProperty() LoaSortCriteriaProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LoaSortCriteria) GetPropertyOk() (*LoaSortCriteriaProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LoaSortCriteria) SetProperty(v LoaSortCriteriaProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *LoaSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


