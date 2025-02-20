# TagPointMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TagPoint**](TagPoint.md) |  | 
**Type** | **string** | message type | [default to "tag-point"]

## Methods

### NewTagPointMessageData

`func NewTagPointMessageData(data TagPoint, type_ string, ) *TagPointMessageData`

NewTagPointMessageData instantiates a new TagPointMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPointMessageDataWithDefaults

`func NewTagPointMessageDataWithDefaults() *TagPointMessageData`

NewTagPointMessageDataWithDefaults instantiates a new TagPointMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TagPointMessageData) GetData() TagPoint`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TagPointMessageData) GetDataOk() (*TagPoint, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TagPointMessageData) SetData(v TagPoint)`

SetData sets Data field to given value.


### GetType

`func (o *TagPointMessageData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TagPointMessageData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TagPointMessageData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


