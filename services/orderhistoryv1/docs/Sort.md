# Sort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to [**SortName**](SortName.md) |  | [optional] 
**Direction** | Pointer to [**SortDirection**](SortDirection.md) |  | [optional] 

## Methods

### NewSort

`func NewSort() *Sort`

NewSort instantiates a new Sort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSortWithDefaults

`func NewSortWithDefaults() *Sort`

NewSortWithDefaults instantiates a new Sort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Sort) GetName() SortName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Sort) GetNameOk() (*SortName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Sort) SetName(v SortName)`

SetName sets Name field to given value.

### HasName

`func (o *Sort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDirection

`func (o *Sort) GetDirection() SortDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Sort) GetDirectionOk() (*SortDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Sort) SetDirection(v SortDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Sort) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


