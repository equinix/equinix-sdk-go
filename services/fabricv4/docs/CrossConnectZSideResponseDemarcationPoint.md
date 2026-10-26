# CrossConnectZSideResponseDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Patch Panel Identifier. | [optional] 
**PatchPanelPortA** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used. Required for Connector type FC and ST only.  | [optional] 
**ConnectorType** | Pointer to **string** | Connector Type. | [optional] 
**CageUniqueSpaceId** | Pointer to **string** | Unique identifier of the Cage. | [optional] [readonly] 
**CabinetUniqueSpaceId** | Pointer to **string** | Unique identifier of the Cabinet. | [optional] [readonly] 

## Methods

### NewCrossConnectZSideResponseDemarcationPoint

`func NewCrossConnectZSideResponseDemarcationPoint() *CrossConnectZSideResponseDemarcationPoint`

NewCrossConnectZSideResponseDemarcationPoint instantiates a new CrossConnectZSideResponseDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectZSideResponseDemarcationPointWithDefaults

`func NewCrossConnectZSideResponseDemarcationPointWithDefaults() *CrossConnectZSideResponseDemarcationPoint`

NewCrossConnectZSideResponseDemarcationPointWithDefaults instantiates a new CrossConnectZSideResponseDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *CrossConnectZSideResponseDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *CrossConnectZSideResponseDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelPortA() int32`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelPortAOk() (*int32, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *CrossConnectZSideResponseDemarcationPoint) SetPatchPanelPortA(v int32)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *CrossConnectZSideResponseDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelPortB() int32`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetPatchPanelPortBOk() (*int32, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *CrossConnectZSideResponseDemarcationPoint) SetPatchPanelPortB(v int32)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *CrossConnectZSideResponseDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *CrossConnectZSideResponseDemarcationPoint) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CrossConnectZSideResponseDemarcationPoint) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CrossConnectZSideResponseDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetCageUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) GetCageUniqueSpaceId() string`

GetCageUniqueSpaceId returns the CageUniqueSpaceId field if non-nil, zero value otherwise.

### GetCageUniqueSpaceIdOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetCageUniqueSpaceIdOk() (*string, bool)`

GetCageUniqueSpaceIdOk returns a tuple with the CageUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) SetCageUniqueSpaceId(v string)`

SetCageUniqueSpaceId sets CageUniqueSpaceId field to given value.

### HasCageUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) HasCageUniqueSpaceId() bool`

HasCageUniqueSpaceId returns a boolean if a field has been set.

### GetCabinetUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) GetCabinetUniqueSpaceId() string`

GetCabinetUniqueSpaceId returns the CabinetUniqueSpaceId field if non-nil, zero value otherwise.

### GetCabinetUniqueSpaceIdOk

`func (o *CrossConnectZSideResponseDemarcationPoint) GetCabinetUniqueSpaceIdOk() (*string, bool)`

GetCabinetUniqueSpaceIdOk returns a tuple with the CabinetUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) SetCabinetUniqueSpaceId(v string)`

SetCabinetUniqueSpaceId sets CabinetUniqueSpaceId field to given value.

### HasCabinetUniqueSpaceId

`func (o *CrossConnectZSideResponseDemarcationPoint) HasCabinetUniqueSpaceId() bool`

HasCabinetUniqueSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


