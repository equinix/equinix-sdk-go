# TagPointData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**[]TagPointDataArrayCurrent**](TagPointDataArrayCurrent.md) |  | [optional] 
**Status** | Pointer to [**TagPointDataStatus**](TagPointDataStatus.md) |  | [optional] 

## Methods

### NewTagPointData

`func NewTagPointData() *TagPointData`

NewTagPointData instantiates a new TagPointData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPointDataWithDefaults

`func NewTagPointDataWithDefaults() *TagPointData`

NewTagPointDataWithDefaults instantiates a new TagPointData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *TagPointData) GetPayLoad() []TagPointDataArrayCurrent`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *TagPointData) GetPayLoadOk() (*[]TagPointDataArrayCurrent, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *TagPointData) SetPayLoad(v []TagPointDataArrayCurrent)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *TagPointData) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *TagPointData) GetStatus() TagPointDataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TagPointData) GetStatusOk() (*TagPointDataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TagPointData) SetStatus(v TagPointDataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TagPointData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


