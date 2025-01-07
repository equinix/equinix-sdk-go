# TagPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataQuality** | **string** | data quality: Good | Bad | Uncertain | 
**Ibx** | **string** | ibx | 
**Reading** | [**TagPointValueWithUnit**](TagPointValueWithUnit.md) |  | 
**ReadingTime** | **string** | message reading time | 
**StreamId** | **string** | unique message id | 
**Tag** | [**TagDetails**](TagDetails.md) |  | 

## Methods

### NewTagPoint

`func NewTagPoint(dataQuality string, ibx string, reading TagPointValueWithUnit, readingTime string, streamId string, tag TagDetails, ) *TagPoint`

NewTagPoint instantiates a new TagPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagPointWithDefaults

`func NewTagPointWithDefaults() *TagPoint`

NewTagPointWithDefaults instantiates a new TagPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataQuality

`func (o *TagPoint) GetDataQuality() string`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *TagPoint) GetDataQualityOk() (*string, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *TagPoint) SetDataQuality(v string)`

SetDataQuality sets DataQuality field to given value.


### GetIbx

`func (o *TagPoint) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *TagPoint) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *TagPoint) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetReading

`func (o *TagPoint) GetReading() TagPointValueWithUnit`

GetReading returns the Reading field if non-nil, zero value otherwise.

### GetReadingOk

`func (o *TagPoint) GetReadingOk() (*TagPointValueWithUnit, bool)`

GetReadingOk returns a tuple with the Reading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReading

`func (o *TagPoint) SetReading(v TagPointValueWithUnit)`

SetReading sets Reading field to given value.


### GetReadingTime

`func (o *TagPoint) GetReadingTime() string`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *TagPoint) GetReadingTimeOk() (*string, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *TagPoint) SetReadingTime(v string)`

SetReadingTime sets ReadingTime field to given value.


### GetStreamId

`func (o *TagPoint) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *TagPoint) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *TagPoint) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetTag

`func (o *TagPoint) GetTag() TagDetails`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TagPoint) GetTagOk() (*TagDetails, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TagPoint) SetTag(v TagDetails)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


