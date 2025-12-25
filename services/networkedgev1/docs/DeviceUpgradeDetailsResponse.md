# DeviceUpgradeDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Unique Id of the upgrade. | [optional] 
**VirtualDeviceUuid** | Pointer to **string** | Unique Id of the device. | [optional] 
**Status** | Pointer to **string** | The status of the upgrade. REQUEST_ACCEPTED, IN_PROGRESS, SUCCESS, FAILED, CANCELLED | [optional] 
**RequestedDate** | Pointer to **string** | Requested date | [optional] 
**CompletiondDate** | Pointer to **string** | Requested date | [optional] 
**RequestedBy** | Pointer to **string** | Requested by. | [optional] 

## Methods

### NewDeviceUpgradeDetailsResponse

`func NewDeviceUpgradeDetailsResponse() *DeviceUpgradeDetailsResponse`

NewDeviceUpgradeDetailsResponse instantiates a new DeviceUpgradeDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpgradeDetailsResponseWithDefaults

`func NewDeviceUpgradeDetailsResponseWithDefaults() *DeviceUpgradeDetailsResponse`

NewDeviceUpgradeDetailsResponseWithDefaults instantiates a new DeviceUpgradeDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *DeviceUpgradeDetailsResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *DeviceUpgradeDetailsResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *DeviceUpgradeDetailsResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *DeviceUpgradeDetailsResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVirtualDeviceUuid

`func (o *DeviceUpgradeDetailsResponse) GetVirtualDeviceUuid() string`

GetVirtualDeviceUuid returns the VirtualDeviceUuid field if non-nil, zero value otherwise.

### GetVirtualDeviceUuidOk

`func (o *DeviceUpgradeDetailsResponse) GetVirtualDeviceUuidOk() (*string, bool)`

GetVirtualDeviceUuidOk returns a tuple with the VirtualDeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDeviceUuid

`func (o *DeviceUpgradeDetailsResponse) SetVirtualDeviceUuid(v string)`

SetVirtualDeviceUuid sets VirtualDeviceUuid field to given value.

### HasVirtualDeviceUuid

`func (o *DeviceUpgradeDetailsResponse) HasVirtualDeviceUuid() bool`

HasVirtualDeviceUuid returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceUpgradeDetailsResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceUpgradeDetailsResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceUpgradeDetailsResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceUpgradeDetailsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRequestedDate

`func (o *DeviceUpgradeDetailsResponse) GetRequestedDate() string`

GetRequestedDate returns the RequestedDate field if non-nil, zero value otherwise.

### GetRequestedDateOk

`func (o *DeviceUpgradeDetailsResponse) GetRequestedDateOk() (*string, bool)`

GetRequestedDateOk returns a tuple with the RequestedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDate

`func (o *DeviceUpgradeDetailsResponse) SetRequestedDate(v string)`

SetRequestedDate sets RequestedDate field to given value.

### HasRequestedDate

`func (o *DeviceUpgradeDetailsResponse) HasRequestedDate() bool`

HasRequestedDate returns a boolean if a field has been set.

### GetCompletiondDate

`func (o *DeviceUpgradeDetailsResponse) GetCompletiondDate() string`

GetCompletiondDate returns the CompletiondDate field if non-nil, zero value otherwise.

### GetCompletiondDateOk

`func (o *DeviceUpgradeDetailsResponse) GetCompletiondDateOk() (*string, bool)`

GetCompletiondDateOk returns a tuple with the CompletiondDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletiondDate

`func (o *DeviceUpgradeDetailsResponse) SetCompletiondDate(v string)`

SetCompletiondDate sets CompletiondDate field to given value.

### HasCompletiondDate

`func (o *DeviceUpgradeDetailsResponse) HasCompletiondDate() bool`

HasCompletiondDate returns a boolean if a field has been set.

### GetRequestedBy

`func (o *DeviceUpgradeDetailsResponse) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *DeviceUpgradeDetailsResponse) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *DeviceUpgradeDetailsResponse) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *DeviceUpgradeDetailsResponse) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


