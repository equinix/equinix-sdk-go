# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | [**EnvironmentAssetDetails**](EnvironmentAssetDetails.md) |  | 
**DataQuality** | Pointer to **string** | data quality: Good | Bad | Uncertain | [optional] 
**Ibx** | **string** | ibx | 
**Reading** | [**EnvironmentValueWithUnit**](EnvironmentValueWithUnit.md) |  | 
**ReadingTime** | Pointer to **string** | message reading time | [optional] 
**StreamId** | **string** | unique message id | 
**Tag** | [**EnvironmentTagDetails**](EnvironmentTagDetails.md) |  | 

## Methods

### NewEnvironment

`func NewEnvironment(asset EnvironmentAssetDetails, ibx string, reading EnvironmentValueWithUnit, streamId string, tag EnvironmentTagDetails, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *Environment) GetAsset() EnvironmentAssetDetails`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Environment) GetAssetOk() (*EnvironmentAssetDetails, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Environment) SetAsset(v EnvironmentAssetDetails)`

SetAsset sets Asset field to given value.


### GetDataQuality

`func (o *Environment) GetDataQuality() string`

GetDataQuality returns the DataQuality field if non-nil, zero value otherwise.

### GetDataQualityOk

`func (o *Environment) GetDataQualityOk() (*string, bool)`

GetDataQualityOk returns a tuple with the DataQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataQuality

`func (o *Environment) SetDataQuality(v string)`

SetDataQuality sets DataQuality field to given value.

### HasDataQuality

`func (o *Environment) HasDataQuality() bool`

HasDataQuality returns a boolean if a field has been set.

### GetIbx

`func (o *Environment) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *Environment) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *Environment) SetIbx(v string)`

SetIbx sets Ibx field to given value.


### GetReading

`func (o *Environment) GetReading() EnvironmentValueWithUnit`

GetReading returns the Reading field if non-nil, zero value otherwise.

### GetReadingOk

`func (o *Environment) GetReadingOk() (*EnvironmentValueWithUnit, bool)`

GetReadingOk returns a tuple with the Reading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReading

`func (o *Environment) SetReading(v EnvironmentValueWithUnit)`

SetReading sets Reading field to given value.


### GetReadingTime

`func (o *Environment) GetReadingTime() string`

GetReadingTime returns the ReadingTime field if non-nil, zero value otherwise.

### GetReadingTimeOk

`func (o *Environment) GetReadingTimeOk() (*string, bool)`

GetReadingTimeOk returns a tuple with the ReadingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadingTime

`func (o *Environment) SetReadingTime(v string)`

SetReadingTime sets ReadingTime field to given value.

### HasReadingTime

`func (o *Environment) HasReadingTime() bool`

HasReadingTime returns a boolean if a field has been set.

### GetStreamId

`func (o *Environment) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *Environment) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *Environment) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.


### GetTag

`func (o *Environment) GetTag() EnvironmentTagDetails`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Environment) GetTagOk() (*EnvironmentTagDetails, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Environment) SetTag(v EnvironmentTagDetails)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


