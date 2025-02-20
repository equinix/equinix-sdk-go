# EnvironmentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**EnvironmentDataPayLoad**](EnvironmentDataPayLoad.md) |  | [optional] 
**Status** | Pointer to [**EnvironmentDataStatus**](EnvironmentDataStatus.md) |  | [optional] 

## Methods

### NewEnvironmentData

`func NewEnvironmentData() *EnvironmentData`

NewEnvironmentData instantiates a new EnvironmentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentDataWithDefaults

`func NewEnvironmentDataWithDefaults() *EnvironmentData`

NewEnvironmentDataWithDefaults instantiates a new EnvironmentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *EnvironmentData) GetPayLoad() EnvironmentDataPayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *EnvironmentData) GetPayLoadOk() (*EnvironmentDataPayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *EnvironmentData) SetPayLoad(v EnvironmentDataPayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *EnvironmentData) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentData) GetStatus() EnvironmentDataStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentData) GetStatusOk() (*EnvironmentDataStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentData) SetStatus(v EnvironmentDataStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


