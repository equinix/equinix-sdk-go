# CrossConnectZSideRequestDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Patch Panel Identifier. | [optional] 
**PatchPanelPortA** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used. Required for Connector type FC and ST only.  | [optional] 
**ConnectorType** | Pointer to **string** | Connector Type. | [optional] 
**CageUniqueSpaceId** | Pointer to **string** | Unique identifier of the Cage. | [optional] 
**CabinetUniqueSpaceId** | Pointer to **string** | Unique identifier of the Cabinet. | [optional] 

## Methods

### NewCrossConnectZSideRequestDemarcationPoint

`func NewCrossConnectZSideRequestDemarcationPoint() *CrossConnectZSideRequestDemarcationPoint`

NewCrossConnectZSideRequestDemarcationPoint instantiates a new CrossConnectZSideRequestDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectZSideRequestDemarcationPointWithDefaults

`func NewCrossConnectZSideRequestDemarcationPointWithDefaults() *CrossConnectZSideRequestDemarcationPoint`

NewCrossConnectZSideRequestDemarcationPointWithDefaults instantiates a new CrossConnectZSideRequestDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *CrossConnectZSideRequestDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *CrossConnectZSideRequestDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelPortA() int32`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelPortAOk() (*int32, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *CrossConnectZSideRequestDemarcationPoint) SetPatchPanelPortA(v int32)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *CrossConnectZSideRequestDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelPortB() int32`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetPatchPanelPortBOk() (*int32, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *CrossConnectZSideRequestDemarcationPoint) SetPatchPanelPortB(v int32)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *CrossConnectZSideRequestDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *CrossConnectZSideRequestDemarcationPoint) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CrossConnectZSideRequestDemarcationPoint) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CrossConnectZSideRequestDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetCageUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) GetCageUniqueSpaceId() string`

GetCageUniqueSpaceId returns the CageUniqueSpaceId field if non-nil, zero value otherwise.

### GetCageUniqueSpaceIdOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetCageUniqueSpaceIdOk() (*string, bool)`

GetCageUniqueSpaceIdOk returns a tuple with the CageUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) SetCageUniqueSpaceId(v string)`

SetCageUniqueSpaceId sets CageUniqueSpaceId field to given value.

### HasCageUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) HasCageUniqueSpaceId() bool`

HasCageUniqueSpaceId returns a boolean if a field has been set.

### GetCabinetUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) GetCabinetUniqueSpaceId() string`

GetCabinetUniqueSpaceId returns the CabinetUniqueSpaceId field if non-nil, zero value otherwise.

### GetCabinetUniqueSpaceIdOk

`func (o *CrossConnectZSideRequestDemarcationPoint) GetCabinetUniqueSpaceIdOk() (*string, bool)`

GetCabinetUniqueSpaceIdOk returns a tuple with the CabinetUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) SetCabinetUniqueSpaceId(v string)`

SetCabinetUniqueSpaceId sets CabinetUniqueSpaceId field to given value.

### HasCabinetUniqueSpaceId

`func (o *CrossConnectZSideRequestDemarcationPoint) HasCabinetUniqueSpaceId() bool`

HasCabinetUniqueSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


