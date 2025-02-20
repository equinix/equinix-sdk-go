# MessageType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to [**[]Asset**](Asset.md) | List of asset message type to subscribe | [optional] 
**CustomAlert** | Pointer to [**[]CustomAlert**](CustomAlert.md) | List of custom alert message type to subscribe | [optional] 
**Environmental** | Pointer to [**[]Environmental**](Environmental.md) | List of environmental message type to subscribe | [optional] 
**MeteredPower** | Pointer to [**[]MeteredPower**](MeteredPower.md) | List of metered power message type to subscribe | [optional] 
**Power** | Pointer to [**[]PowerMessage**](PowerMessage.md) | List of power message type to subscribe | [optional] 
**SystemAlert** | Pointer to [**[]SystemAlert**](SystemAlert.md) | List of system alert message type to subscribe | [optional] 

## Methods

### NewMessageType

`func NewMessageType() *MessageType`

NewMessageType instantiates a new MessageType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTypeWithDefaults

`func NewMessageTypeWithDefaults() *MessageType`

NewMessageTypeWithDefaults instantiates a new MessageType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *MessageType) GetAsset() []Asset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *MessageType) GetAssetOk() (*[]Asset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *MessageType) SetAsset(v []Asset)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *MessageType) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetCustomAlert

`func (o *MessageType) GetCustomAlert() []CustomAlert`

GetCustomAlert returns the CustomAlert field if non-nil, zero value otherwise.

### GetCustomAlertOk

`func (o *MessageType) GetCustomAlertOk() (*[]CustomAlert, bool)`

GetCustomAlertOk returns a tuple with the CustomAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAlert

`func (o *MessageType) SetCustomAlert(v []CustomAlert)`

SetCustomAlert sets CustomAlert field to given value.

### HasCustomAlert

`func (o *MessageType) HasCustomAlert() bool`

HasCustomAlert returns a boolean if a field has been set.

### GetEnvironmental

`func (o *MessageType) GetEnvironmental() []Environmental`

GetEnvironmental returns the Environmental field if non-nil, zero value otherwise.

### GetEnvironmentalOk

`func (o *MessageType) GetEnvironmentalOk() (*[]Environmental, bool)`

GetEnvironmentalOk returns a tuple with the Environmental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmental

`func (o *MessageType) SetEnvironmental(v []Environmental)`

SetEnvironmental sets Environmental field to given value.

### HasEnvironmental

`func (o *MessageType) HasEnvironmental() bool`

HasEnvironmental returns a boolean if a field has been set.

### GetMeteredPower

`func (o *MessageType) GetMeteredPower() []MeteredPower`

GetMeteredPower returns the MeteredPower field if non-nil, zero value otherwise.

### GetMeteredPowerOk

`func (o *MessageType) GetMeteredPowerOk() (*[]MeteredPower, bool)`

GetMeteredPowerOk returns a tuple with the MeteredPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeteredPower

`func (o *MessageType) SetMeteredPower(v []MeteredPower)`

SetMeteredPower sets MeteredPower field to given value.

### HasMeteredPower

`func (o *MessageType) HasMeteredPower() bool`

HasMeteredPower returns a boolean if a field has been set.

### GetPower

`func (o *MessageType) GetPower() []PowerMessage`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *MessageType) GetPowerOk() (*[]PowerMessage, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *MessageType) SetPower(v []PowerMessage)`

SetPower sets Power field to given value.

### HasPower

`func (o *MessageType) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetSystemAlert

`func (o *MessageType) GetSystemAlert() []SystemAlert`

GetSystemAlert returns the SystemAlert field if non-nil, zero value otherwise.

### GetSystemAlertOk

`func (o *MessageType) GetSystemAlertOk() (*[]SystemAlert, bool)`

GetSystemAlertOk returns a tuple with the SystemAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAlert

`func (o *MessageType) SetSystemAlert(v []SystemAlert)`

SetSystemAlert sets SystemAlert field to given value.

### HasSystemAlert

`func (o *MessageType) HasSystemAlert() bool`

HasSystemAlert returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


