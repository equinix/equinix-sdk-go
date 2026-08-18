# StreamSubscriptionSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | [**StreamSubscriptionSearchSortDirection**](StreamSubscriptionSearchSortDirection.md) |  | [default to STREAMSUBSCRIPTIONSEARCHSORTDIRECTION_DESC]
**Property** | [**StreamSubscriptionSearchSortBy**](StreamSubscriptionSearchSortBy.md) |  | [default to STREAMSUBSCRIPTIONSEARCHSORTBY_CREATED_DATE_TIME]

## Methods

### NewStreamSubscriptionSearchSortCriteria

`func NewStreamSubscriptionSearchSortCriteria(direction StreamSubscriptionSearchSortDirection, property StreamSubscriptionSearchSortBy, ) *StreamSubscriptionSearchSortCriteria`

NewStreamSubscriptionSearchSortCriteria instantiates a new StreamSubscriptionSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSubscriptionSearchSortCriteriaWithDefaults

`func NewStreamSubscriptionSearchSortCriteriaWithDefaults() *StreamSubscriptionSearchSortCriteria`

NewStreamSubscriptionSearchSortCriteriaWithDefaults instantiates a new StreamSubscriptionSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *StreamSubscriptionSearchSortCriteria) GetDirection() StreamSubscriptionSearchSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *StreamSubscriptionSearchSortCriteria) GetDirectionOk() (*StreamSubscriptionSearchSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *StreamSubscriptionSearchSortCriteria) SetDirection(v StreamSubscriptionSearchSortDirection)`

SetDirection sets Direction field to given value.


### GetProperty

`func (o *StreamSubscriptionSearchSortCriteria) GetProperty() StreamSubscriptionSearchSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamSubscriptionSearchSortCriteria) GetPropertyOk() (*StreamSubscriptionSearchSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamSubscriptionSearchSortCriteria) SetProperty(v StreamSubscriptionSearchSortBy)`

SetProperty sets Property field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


