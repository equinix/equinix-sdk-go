# DeviceRebootResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceUUID** | Pointer to **string** | Unique Id of the device. | [optional] 
**Status** | Pointer to **string** | The status of the reboot. | [optional] 
**RequestedBy** | Pointer to **string** | Requested by | [optional] 
**RequestedDate** | Pointer to **string** | Requested date | [optional] 
**CompletiondDate** | Pointer to **string** | Requested date | [optional] 

## Methods

### NewDeviceRebootResponse

`func NewDeviceRebootResponse() *DeviceRebootResponse`

NewDeviceRebootResponse instantiates a new DeviceRebootResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRebootResponseWithDefaults

`func NewDeviceRebootResponseWithDefaults() *DeviceRebootResponse`

NewDeviceRebootResponseWithDefaults instantiates a new DeviceRebootResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceUUID

`func (o *DeviceRebootResponse) GetDeviceUUID() string`

GetDeviceUUID returns the DeviceUUID field if non-nil, zero value otherwise.

### GetDeviceUUIDOk

`func (o *DeviceRebootResponse) GetDeviceUUIDOk() (*string, bool)`

GetDeviceUUIDOk returns a tuple with the DeviceUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUUID

`func (o *DeviceRebootResponse) SetDeviceUUID(v string)`

SetDeviceUUID sets DeviceUUID field to given value.

### HasDeviceUUID

`func (o *DeviceRebootResponse) HasDeviceUUID() bool`

HasDeviceUUID returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceRebootResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceRebootResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceRebootResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceRebootResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequestedBy

`func (o *DeviceRebootResponse) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *DeviceRebootResponse) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *DeviceRebootResponse) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *DeviceRebootResponse) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetRequestedDate

`func (o *DeviceRebootResponse) GetRequestedDate() string`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *DeviceRebootResponse) GetRequestedDateOk() (*string, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *DeviceRebootResponse) SetRequestedDate(v string)`

SetRequestedDate sets RequestedDate field to given value.

### HasRequestedDate

`func (o *DeviceRebootResponse) HasRequestedDate() bool`

HasRequestedDate returns a boolean if a field has been set.

### GetCompletiondDate

`func (o *DeviceRebootResponse) GetCompletiondDate() string`

GetCompletiondDate returns the CompletiondDate field if non-nil, zero value otherwise.

### GetCompletiondDateOk

`func (o *DeviceRebootResponse) GetCompletiondDateOk() (*string, bool)`

GetCompletiondDateOk returns a tuple with the CompletiondDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletiondDate

`func (o *DeviceRebootResponse) SetCompletiondDate(v string)`

SetCompletiondDate sets CompletiondDate field to given value.

### HasCompletiondDate

`func (o *DeviceRebootResponse) HasCompletiondDate() bool`

HasCompletiondDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


