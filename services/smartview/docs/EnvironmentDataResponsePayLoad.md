# EnvironmentDataResponsePayLoad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]EnvironmentDataForArray**](EnvironmentDataForArray.md) |  | [optional] 
**TotalCount** | Pointer to **int32** | total number of data values | [optional] 

## Methods

### NewEnvironmentDataResponsePayLoad

`func NewEnvironmentDataResponsePayLoad() *EnvironmentDataResponsePayLoad`

NewEnvironmentDataResponsePayLoad instantiates a new EnvironmentDataResponsePayLoad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataResponsePayLoadWithDefaults

`func NewEnvironmentDataResponsePayLoadWithDefaults() *EnvironmentDataResponsePayLoad`

NewEnvironmentDataResponsePayLoadWithDefaults instantiates a new EnvironmentDataResponsePayLoad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EnvironmentDataResponsePayLoad) GetData() []EnvironmentDataForArray`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EnvironmentDataResponsePayLoad) GetDataOk() (*[]EnvironmentDataForArray, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EnvironmentDataResponsePayLoad) SetData(v []EnvironmentDataForArray)`

SetData sets Data field to given value.

### HasData

`func (o *EnvironmentDataResponsePayLoad) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTotalCount

`func (o *EnvironmentDataResponsePayLoad) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EnvironmentDataResponsePayLoad) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EnvironmentDataResponsePayLoad) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *EnvironmentDataResponsePayLoad) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


