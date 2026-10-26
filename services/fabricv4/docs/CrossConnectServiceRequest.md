# CrossConnectServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**CrossConnectServiceType**](CrossConnectServiceType.md) |  | 
**Device** | Pointer to [**PatchEquipmentRequest**](PatchEquipmentRequest.md) | Patch equipment device details. Required for the XC_PATCH_EQUIPMENT service.  | [optional] 
**Description** | Pointer to **string** | Additional notes for the service. Required for the XC_PATCH_EQUIPMENT service.  | [optional] 

## Methods

### NewCrossConnectServiceRequest

`func NewCrossConnectServiceRequest(type_ CrossConnectServiceType, ) *CrossConnectServiceRequest`

NewCrossConnectServiceRequest instantiates a new CrossConnectServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectServiceRequestWithDefaults

`func NewCrossConnectServiceRequestWithDefaults() *CrossConnectServiceRequest`

NewCrossConnectServiceRequestWithDefaults instantiates a new CrossConnectServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossConnectServiceRequest) GetType() CrossConnectServiceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectServiceRequest) GetTypeOk() (*CrossConnectServiceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectServiceRequest) SetType(v CrossConnectServiceType)`

SetType sets Type field to given value.


### GetDevice

`func (o *CrossConnectServiceRequest) GetDevice() PatchEquipmentRequest`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CrossConnectServiceRequest) GetDeviceOk() (*PatchEquipmentRequest, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CrossConnectServiceRequest) SetDevice(v PatchEquipmentRequest)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *CrossConnectServiceRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDescription

`func (o *CrossConnectServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CrossConnectServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CrossConnectServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CrossConnectServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


