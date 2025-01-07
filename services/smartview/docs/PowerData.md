# PowerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to [**PowerDataPayLoad**](PowerDataPayLoad.md) |  | [optional] 
**Status** | Pointer to [**AssetDetailResponseStatus**](AssetDetailResponseStatus.md) |  | [optional] 

## Methods

### NewPowerData

`func NewPowerData() *PowerData`

NewPowerData instantiates a new PowerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerDataWithDefaults

`func NewPowerDataWithDefaults() *PowerData`

NewPowerDataWithDefaults instantiates a new PowerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *PowerData) GetPayLoad() PowerDataPayLoad`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *PowerData) GetPayLoadOk() (*PowerDataPayLoad, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *PowerData) SetPayLoad(v PowerDataPayLoad)`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *PowerData) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *PowerData) GetStatus() AssetDetailResponseStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerData) GetStatusOk() (*AssetDetailResponseStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerData) SetStatus(v AssetDetailResponseStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


