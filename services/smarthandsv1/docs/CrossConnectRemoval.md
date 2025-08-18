# CrossConnectRemoval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SerialNumber** | **string** | Cross Connect serialNumber | 
**DeviceCabinet** | **string** | Cross Connect Cabinet | 
**DeviceConnectorType** | **string** | Cross Connect Device Connector Type | 
**DeviceDetails** | **string** | Cross Connect Device Details | 
**DevicePort** | **string** | Cross Connect Device Port | 
**RemovePatchCableWithLiveTraffic** | Pointer to **bool** | Proceed with the de-install if live traffic is detected (Optional). By selecting this option, I agree to the terms of removal with live traffic. Please be advised that Equinix will complete the requested removal based on your instruction and will not be responsible for any service outages resulting from this removal. | [optional] 
**ScopeOfWork** | Pointer to **string** | Scope of work | [optional] 

## Methods

### NewCrossConnectRemoval

`func NewCrossConnectRemoval(serialNumber string, deviceCabinet string, deviceConnectorType string, deviceDetails string, devicePort string, ) *CrossConnectRemoval`

NewCrossConnectRemoval instantiates a new CrossConnectRemoval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectRemovalWithDefaults

`func NewCrossConnectRemovalWithDefaults() *CrossConnectRemoval`

NewCrossConnectRemovalWithDefaults instantiates a new CrossConnectRemoval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSerialNumber

`func (o *CrossConnectRemoval) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CrossConnectRemoval) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CrossConnectRemoval) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetDeviceCabinet

`func (o *CrossConnectRemoval) GetDeviceCabinet() string`

GetDeviceCabinet returns the DeviceCabinet field if non-nil, zero value otherwise.

### GetDeviceCabinetOk

`func (o *CrossConnectRemoval) GetDeviceCabinetOk() (*string, bool)`

GetDeviceCabinetOk returns a tuple with the DeviceCabinet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCabinet

`func (o *CrossConnectRemoval) SetDeviceCabinet(v string)`

SetDeviceCabinet sets DeviceCabinet field to given value.


### GetDeviceConnectorType

`func (o *CrossConnectRemoval) GetDeviceConnectorType() string`

GetDeviceConnectorType returns the DeviceConnectorType field if non-nil, zero value otherwise.

### GetDeviceConnectorTypeOk

`func (o *CrossConnectRemoval) GetDeviceConnectorTypeOk() (*string, bool)`

GetDeviceConnectorTypeOk returns a tuple with the DeviceConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceConnectorType

`func (o *CrossConnectRemoval) SetDeviceConnectorType(v string)`

SetDeviceConnectorType sets DeviceConnectorType field to given value.


### GetDeviceDetails

`func (o *CrossConnectRemoval) GetDeviceDetails() string`

GetDeviceDetails returns the DeviceDetails field if non-nil, zero value otherwise.

### GetDeviceDetailsOk

`func (o *CrossConnectRemoval) GetDeviceDetailsOk() (*string, bool)`

GetDeviceDetailsOk returns a tuple with the DeviceDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDetails

`func (o *CrossConnectRemoval) SetDeviceDetails(v string)`

SetDeviceDetails sets DeviceDetails field to given value.


### GetDevicePort

`func (o *CrossConnectRemoval) GetDevicePort() string`

GetDevicePort returns the DevicePort field if non-nil, zero value otherwise.

### GetDevicePortOk

`func (o *CrossConnectRemoval) GetDevicePortOk() (*string, bool)`

GetDevicePortOk returns a tuple with the DevicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePort

`func (o *CrossConnectRemoval) SetDevicePort(v string)`

SetDevicePort sets DevicePort field to given value.


### GetRemovePatchCableWithLiveTraffic

`func (o *CrossConnectRemoval) GetRemovePatchCableWithLiveTraffic() bool`

GetRemovePatchCableWithLiveTraffic returns the RemovePatchCableWithLiveTraffic field if non-nil, zero value otherwise.

### GetRemovePatchCableWithLiveTrafficOk

`func (o *CrossConnectRemoval) GetRemovePatchCableWithLiveTrafficOk() (*bool, bool)`

GetRemovePatchCableWithLiveTrafficOk returns a tuple with the RemovePatchCableWithLiveTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovePatchCableWithLiveTraffic

`func (o *CrossConnectRemoval) SetRemovePatchCableWithLiveTraffic(v bool)`

SetRemovePatchCableWithLiveTraffic sets RemovePatchCableWithLiveTraffic field to given value.

### HasRemovePatchCableWithLiveTraffic

`func (o *CrossConnectRemoval) HasRemovePatchCableWithLiveTraffic() bool`

HasRemovePatchCableWithLiveTraffic returns a boolean if a field has been set.

### GetScopeOfWork

`func (o *CrossConnectRemoval) GetScopeOfWork() string`

GetScopeOfWork returns the ScopeOfWork field if non-nil, zero value otherwise.

### GetScopeOfWorkOk

`func (o *CrossConnectRemoval) GetScopeOfWorkOk() (*string, bool)`

GetScopeOfWorkOk returns a tuple with the ScopeOfWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOfWork

`func (o *CrossConnectRemoval) SetScopeOfWork(v string)`

SetScopeOfWork sets ScopeOfWork field to given value.

### HasScopeOfWork

`func (o *CrossConnectRemoval) HasScopeOfWork() bool`

HasScopeOfWork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


