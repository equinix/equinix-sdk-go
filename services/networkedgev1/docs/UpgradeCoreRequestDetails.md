# UpgradeCoreRequestDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Core** | Pointer to **float32** | Core requested for the device | [optional] 
**UpgradePeerDevice** | Pointer to **bool** | Whether the peer device should be upgraded or not. | [optional] 

## Methods

### NewUpgradeCoreRequestDetails

`func NewUpgradeCoreRequestDetails() *UpgradeCoreRequestDetails`

NewUpgradeCoreRequestDetails instantiates a new UpgradeCoreRequestDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeCoreRequestDetailsWithDefaults

`func NewUpgradeCoreRequestDetailsWithDefaults() *UpgradeCoreRequestDetails`

NewUpgradeCoreRequestDetailsWithDefaults instantiates a new UpgradeCoreRequestDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCore

`func (o *UpgradeCoreRequestDetails) GetCore() float32`

GetCore returns the Core field if non-nil, zero value otherwise.

### GetCoreOk

`func (o *UpgradeCoreRequestDetails) GetCoreOk() (*float32, bool)`

GetCoreOk returns a tuple with the Core field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCore

`func (o *UpgradeCoreRequestDetails) SetCore(v float32)`

SetCore sets Core field to given value.

### HasCore

`func (o *UpgradeCoreRequestDetails) HasCore() bool`

HasCore returns a boolean if a field has been set.

### GetUpgradePeerDevice

`func (o *UpgradeCoreRequestDetails) GetUpgradePeerDevice() bool`

GetUpgradePeerDevice returns the UpgradePeerDevice field if non-nil, zero value otherwise.

### GetUpgradePeerDeviceOk

`func (o *UpgradeCoreRequestDetails) GetUpgradePeerDeviceOk() (*bool, bool)`

GetUpgradePeerDeviceOk returns a tuple with the UpgradePeerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradePeerDevice

`func (o *UpgradeCoreRequestDetails) SetUpgradePeerDevice(v bool)`

SetUpgradePeerDevice sets UpgradePeerDevice field to given value.

### HasUpgradePeerDevice

`func (o *UpgradeCoreRequestDetails) HasUpgradePeerDevice() bool`

HasUpgradePeerDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


