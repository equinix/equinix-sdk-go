# SearchSortItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | [**SearchSortItemDirection**](SearchSortItemDirection.md) |  | 
**Property** | **string** |  | 

## Methods

### NewSearchSortItem

`func NewSearchSortItem(direction SearchSortItemDirection, property string, ) *SearchSortItem`

NewSearchSortItem instantiates a new SearchSortItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchSortItemWithDefaults

`func NewSearchSortItemWithDefaults() *SearchSortItem`

NewSearchSortItemWithDefaults instantiates a new SearchSortItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *SearchSortItem) GetDirection() SearchSortItemDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *SearchSortItem) GetDirectionOk() (*SearchSortItemDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *SearchSortItem) SetDirection(v SearchSortItemDirection)`

SetDirection sets Direction field to given value.


### GetProperty

`func (o *SearchSortItem) GetProperty() string`

GetProperty returns the Property field if non-nil, zero value otherwise.

### GetPropertyOk

`func (o *SearchSortItem) GetPropertyOk() (*string, bool)`

GetPropertyOk returns a tuple with the Property field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperty

`func (o *SearchSortItem) SetProperty(v string)`

SetProperty sets Property field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


