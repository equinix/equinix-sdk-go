# CrossConnectInstall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | **string** | Cross Connect serialNumber | 
**DeviceCabinet** | **string** | Cross Connect Cabinet | 
**DeviceConnectorType** | **string** | Cross Connect Device Connector Type | 
**DeviceDetails** | **string** | Cross Connect Device Details | 
**DevicePort** | **string** | Cross Connect Device Port | 
**LightLinkVerification** | Pointer to **bool** | Cross Connect Light Link Verification (Optional). Select this option if you would like a light reading provided and tx/rx verification after the cross connect is completed. In order to verify the correct transmit/receive alignment, please ensure your Z-Side Cross Connect Partner has their end fully extended to their equipment and their port is enabled. A separate billable activity will be created. | [optional] 
**ScopeOfWork** | Pointer to **string** | Scope of work | [optional] 

## Methods

### NewCrossConnectInstall

`func NewCrossConnectInstall(serialNumber string, deviceCabinet string, deviceConnectorType string, deviceDetails string, devicePort string, ) *CrossConnectInstall`

NewCrossConnectInstall instantiates a new CrossConnectInstall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectInstallWithDefaults

`func NewCrossConnectInstallWithDefaults() *CrossConnectInstall`

NewCrossConnectInstallWithDefaults instantiates a new CrossConnectInstall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *CrossConnectInstall) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CrossConnectInstall) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CrossConnectInstall) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceCabinet

`func (o *CrossConnectInstall) GetDeviceCabinet() string`

GetDeviceCabinet returns the DeviceCabinet field if non-nil, zero value otherwise.

### GetDeviceCabinetOk

`func (o *CrossConnectInstall) GetDeviceCabinetOk() (*string, bool)`

GetDeviceCabinetOk returns a tuple with the DeviceCabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCabinet

`func (o *CrossConnectInstall) SetDeviceCabinet(v string)`

SetDeviceCabinet sets DeviceCabinet field to given value.


### GetDeviceConnectorType

`func (o *CrossConnectInstall) GetDeviceConnectorType() string`

GetDeviceConnectorType returns the DeviceConnectorType field if non-nil, zero value otherwise.

### GetDeviceConnectorTypeOk

`func (o *CrossConnectInstall) GetDeviceConnectorTypeOk() (*string, bool)`

GetDeviceConnectorTypeOk returns a tuple with the DeviceConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnectorType

`func (o *CrossConnectInstall) SetDeviceConnectorType(v string)`

SetDeviceConnectorType sets DeviceConnectorType field to given value.


### GetDeviceDetails

`func (o *CrossConnectInstall) GetDeviceDetails() string`

GetDeviceDetails returns the DeviceDetails field if non-nil, zero value otherwise.

### GetDeviceDetailsOk

`func (o *CrossConnectInstall) GetDeviceDetailsOk() (*string, bool)`

GetDeviceDetailsOk returns a tuple with the DeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetails

`func (o *CrossConnectInstall) SetDeviceDetails(v string)`

SetDeviceDetails sets DeviceDetails field to given value.


### GetDevicePort

`func (o *CrossConnectInstall) GetDevicePort() string`

GetDevicePort returns the DevicePort field if non-nil, zero value otherwise.

### GetDevicePortOk

`func (o *CrossConnectInstall) GetDevicePortOk() (*string, bool)`

GetDevicePortOk returns a tuple with the DevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePort

`func (o *CrossConnectInstall) SetDevicePort(v string)`

SetDevicePort sets DevicePort field to given value.


### GetLightLinkVerification

`func (o *CrossConnectInstall) GetLightLinkVerification() bool`

GetLightLinkVerification returns the LightLinkVerification field if non-nil, zero value otherwise.

### GetLightLinkVerificationOk

`func (o *CrossConnectInstall) GetLightLinkVerificationOk() (*bool, bool)`

GetLightLinkVerificationOk returns a tuple with the LightLinkVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLightLinkVerification

`func (o *CrossConnectInstall) SetLightLinkVerification(v bool)`

SetLightLinkVerification sets LightLinkVerification field to given value.

### HasLightLinkVerification

`func (o *CrossConnectInstall) HasLightLinkVerification() bool`

HasLightLinkVerification returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *CrossConnectInstall) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *CrossConnectInstall) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *CrossConnectInstall) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.

### HasScopeOfWork

`func (o *CrossConnectInstall) HasScopeOfWork() bool`

HasScopeOfWork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


