# DeviceBackupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUuid** | **string** | The unique Id of a virtual device. | 
**Name** | **string** | The name of backup. | 

## Methods

### NewDeviceBackupCreateRequest

`func NewDeviceBackupCreateRequest(deviceUuid string, name string, ) *DeviceBackupCreateRequest`

NewDeviceBackupCreateRequest instantiates a new DeviceBackupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceBackupCreateRequestWithDefaults

`func NewDeviceBackupCreateRequestWithDefaults() *DeviceBackupCreateRequest`

NewDeviceBackupCreateRequestWithDefaults instantiates a new DeviceBackupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUuid

`func (o *DeviceBackupCreateRequest) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *DeviceBackupCreateRequest) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *DeviceBackupCreateRequest) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetName

`func (o *DeviceBackupCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceBackupCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceBackupCreateRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


