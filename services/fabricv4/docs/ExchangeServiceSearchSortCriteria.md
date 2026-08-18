# ExchangeServiceSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**ExchangeServiceSearchSortCriteriaDirection**](ExchangeServiceSearchSortCriteriaDirection.md) |  | [optional] [default to EXCHANGESERVICESEARCHSORTCRITERIADIRECTION_DESC]
**Property** | Pointer to [**ExchangeServiceSearchSortCriteriaProperty**](ExchangeServiceSearchSortCriteriaProperty.md) |  | [optional] [default to EXCHANGESERVICESEARCHSORTCRITERIAPROPERTY_CHANGE_LOG_UPDATED_DATE_TIME]

## Methods

### NewExchangeServiceSearchSortCriteria

`func NewExchangeServiceSearchSortCriteria() *ExchangeServiceSearchSortCriteria`

NewExchangeServiceSearchSortCriteria instantiates a new ExchangeServiceSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceSearchSortCriteriaWithDefaults

`func NewExchangeServiceSearchSortCriteriaWithDefaults() *ExchangeServiceSearchSortCriteria`

NewExchangeServiceSearchSortCriteriaWithDefaults instantiates a new ExchangeServiceSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *ExchangeServiceSearchSortCriteria) GetDirection() ExchangeServiceSearchSortCriteriaDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ExchangeServiceSearchSortCriteria) GetDirectionOk() (*ExchangeServiceSearchSortCriteriaDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ExchangeServiceSearchSortCriteria) SetDirection(v ExchangeServiceSearchSortCriteriaDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ExchangeServiceSearchSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *ExchangeServiceSearchSortCriteria) GetProperty() ExchangeServiceSearchSortCriteriaProperty`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *ExchangeServiceSearchSortCriteria) GetPropertyOk() (*ExchangeServiceSearchSortCriteriaProperty, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *ExchangeServiceSearchSortCriteria) SetProperty(v ExchangeServiceSearchSortCriteriaProperty)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *ExchangeServiceSearchSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


