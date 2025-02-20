# PowerDataError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayLoad** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to [**PowerDataErrorStatus**](PowerDataErrorStatus.md) |  | [optional] 

## Methods

### NewPowerDataError

`func NewPowerDataError() *PowerDataError`

NewPowerDataError instantiates a new PowerDataError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPowerDataErrorWithDefaults

`func NewPowerDataErrorWithDefaults() *PowerDataError`

NewPowerDataErrorWithDefaults instantiates a new PowerDataError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayLoad

`func (o *PowerDataError) GetPayLoad() map[string]interface{}`

GetPayLoad returns the PayLoad field if non-nil, zero value otherwise.

### GetPayLoadOk

`func (o *PowerDataError) GetPayLoadOk() (*map[string]interface{}, bool)`

GetPayLoadOk returns a tuple with the PayLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayLoad

`func (o *PowerDataError) SetPayLoad(v map[string]interface{})`

SetPayLoad sets PayLoad field to given value.

### HasPayLoad

`func (o *PowerDataError) HasPayLoad() bool`

HasPayLoad returns a boolean if a field has been set.

### GetStatus

`func (o *PowerDataError) GetStatus() PowerDataErrorStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PowerDataError) GetStatusOk() (*PowerDataErrorStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PowerDataError) SetStatus(v PowerDataErrorStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PowerDataError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


