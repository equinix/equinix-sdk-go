# StreamSearchSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | [**StreamSearchSortDirection**](StreamSearchSortDirection.md) |  | [default to STREAMSEARCHSORTDIRECTION_DESC]
**Property** | [**StreamSearchSortBy**](StreamSearchSortBy.md) |  | [default to STREAMSEARCHSORTBY_CREATED_DATE_TIME]

## Methods

### NewStreamSearchSortCriteria

`func NewStreamSearchSortCriteria(direction StreamSearchSortDirection, property StreamSearchSortBy, ) *StreamSearchSortCriteria`

NewStreamSearchSortCriteria instantiates a new StreamSearchSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSearchSortCriteriaWithDefaults

`func NewStreamSearchSortCriteriaWithDefaults() *StreamSearchSortCriteria`

NewStreamSearchSortCriteriaWithDefaults instantiates a new StreamSearchSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *StreamSearchSortCriteria) GetDirection() StreamSearchSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *StreamSearchSortCriteria) GetDirectionOk() (*StreamSearchSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *StreamSearchSortCriteria) SetDirection(v StreamSearchSortDirection)`

SetDirection sets Direction field to given value.


### GetProperty

`func (o *StreamSearchSortCriteria) GetProperty() StreamSearchSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamSearchSortCriteria) GetPropertyOk() (*StreamSearchSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamSearchSortCriteria) SetProperty(v StreamSearchSortBy)`

SetProperty sets Property field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


