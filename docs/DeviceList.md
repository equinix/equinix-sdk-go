# DeviceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]Device**](Device.md) |  | [optional] 
**Meta** | Pointer to [**Meta**](Meta.md) |  | [optional] 

## Methods

### NewDeviceList

`func NewDeviceList() *DeviceList`

NewDeviceList instantiates a new DeviceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceListWithDefaults

`func NewDeviceListWithDefaults() *DeviceList`

NewDeviceListWithDefaults instantiates a new DeviceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *DeviceList) GetDevices() []Device`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *DeviceList) GetDevicesOk() (*[]Device, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *DeviceList) SetDevices(v []Device)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *DeviceList) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetMeta

`func (o *DeviceList) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *DeviceList) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *DeviceList) SetMeta(v Meta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *DeviceList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


