# PatchEquipmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CabinetUniqueSpaceId** | **string** | Unique identifier of the Cabinet. | 
**Port** | **int32** | Port number. | 
**ConnectorType** | **string** | Connector Type. | 

## Methods

### NewPatchEquipmentRequest

`func NewPatchEquipmentRequest(cabinetUniqueSpaceId string, port int32, connectorType string, ) *PatchEquipmentRequest`

NewPatchEquipmentRequest instantiates a new PatchEquipmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchEquipmentRequestWithDefaults

`func NewPatchEquipmentRequestWithDefaults() *PatchEquipmentRequest`

NewPatchEquipmentRequestWithDefaults instantiates a new PatchEquipmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCabinetUniqueSpaceId

`func (o *PatchEquipmentRequest) GetCabinetUniqueSpaceId() string`

GetCabinetUniqueSpaceId returns the CabinetUniqueSpaceId field if non-nil, zero value otherwise.

### GetCabinetUniqueSpaceIdOk

`func (o *PatchEquipmentRequest) GetCabinetUniqueSpaceIdOk() (*string, bool)`

GetCabinetUniqueSpaceIdOk returns a tuple with the CabinetUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetUniqueSpaceId

`func (o *PatchEquipmentRequest) SetCabinetUniqueSpaceId(v string)`

SetCabinetUniqueSpaceId sets CabinetUniqueSpaceId field to given value.


### GetPort

`func (o *PatchEquipmentRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PatchEquipmentRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PatchEquipmentRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetConnectorType

`func (o *PatchEquipmentRequest) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *PatchEquipmentRequest) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *PatchEquipmentRequest) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


