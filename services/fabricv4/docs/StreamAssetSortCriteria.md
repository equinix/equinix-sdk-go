# StreamAssetSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to [**StreamAssetSortDirection**](StreamAssetSortDirection.md) |  | [optional] [default to STREAMASSETSORTDIRECTION_DESC]
**Property** | Pointer to [**StreamAssetSortBy**](StreamAssetSortBy.md) |  | [optional] [default to STREAMASSETSORTBY_UUID]

## Methods

### NewStreamAssetSortCriteria

`func NewStreamAssetSortCriteria() *StreamAssetSortCriteria`

NewStreamAssetSortCriteria instantiates a new StreamAssetSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAssetSortCriteriaWithDefaults

`func NewStreamAssetSortCriteriaWithDefaults() *StreamAssetSortCriteria`

NewStreamAssetSortCriteriaWithDefaults instantiates a new StreamAssetSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *StreamAssetSortCriteria) GetDirection() StreamAssetSortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *StreamAssetSortCriteria) GetDirectionOk() (*StreamAssetSortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *StreamAssetSortCriteria) SetDirection(v StreamAssetSortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *StreamAssetSortCriteria) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetProperty

`func (o *StreamAssetSortCriteria) GetProperty() StreamAssetSortBy`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *StreamAssetSortCriteria) GetPropertyOk() (*StreamAssetSortBy, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *StreamAssetSortCriteria) SetProperty(v StreamAssetSortBy)`

SetProperty sets Property field to given value.

### HasProperty

`func (o *StreamAssetSortCriteria) HasProperty() bool`

HasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


