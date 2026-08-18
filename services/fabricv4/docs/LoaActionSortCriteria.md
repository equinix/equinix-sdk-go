# LoaActionSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ExchangeServiceSearchSortCriteriaDirection**](ExchangeServiceSearchSortCriteriaDirection.md) |  | [optional] [default to EXCHANGESERVICESEARCHSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to [**LoaActionSortCriteriaProperty**](LoaActionSortCriteriaProperty.md) |  | [optional] [default to LOAACTIONSORTCRITERIAPROPERTY_CHANGE_LOG_CREATED_DATE_TIME]

## Methods

### NewLoaActionSortCriteria

`func NewLoaActionSortCriteria() *LoaActionSortCriteria`

NewLoaActionSortCriteria instantiates a new LoaActionSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaActionSortCriteriaWithDefaults

`func NewLoaActionSortCriteriaWithDefaults() *LoaActionSortCriteria`

NewLoaActionSortCriteriaWithDefaults instantiates a new LoaActionSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *LoaActionSortCriteria) GetDirection() ExchangeServiceSearchSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LoaActionSortCriteria) GetDirectionOk() (*ExchangeServiceSearchSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LoaActionSortCriteria) SetDirection(v ExchangeServiceSearchSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LoaActionSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *LoaActionSortCriteria) GetProperty() LoaActionSortCriteriaProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *LoaActionSortCriteria) GetPropertyOk() (*LoaActionSortCriteriaProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *LoaActionSortCriteria) SetProperty(v LoaActionSortCriteriaProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *LoaActionSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


