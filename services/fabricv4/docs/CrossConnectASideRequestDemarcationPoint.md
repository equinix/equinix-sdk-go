# CrossConnectASideRequestDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | **string** | Patch Panel Identifier. | 
**PatchPanelPortA** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used. Required for Connector type FC and ST only.  | [optional] 
**ConnectorType** | **string** | Connector Type. | 

## Methods

### NewCrossConnectASideRequestDemarcationPoint

`func NewCrossConnectASideRequestDemarcationPoint(patchPanelId string, connectorType string, ) *CrossConnectASideRequestDemarcationPoint`

NewCrossConnectASideRequestDemarcationPoint instantiates a new CrossConnectASideRequestDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectASideRequestDemarcationPointWithDefaults

`func NewCrossConnectASideRequestDemarcationPointWithDefaults() *CrossConnectASideRequestDemarcationPoint`

NewCrossConnectASideRequestDemarcationPointWithDefaults instantiates a new CrossConnectASideRequestDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *CrossConnectASideRequestDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.


### GetPatchPanelPortA

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelPortA() int32`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelPortAOk() (*int32, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *CrossConnectASideRequestDemarcationPoint) SetPatchPanelPortA(v int32)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *CrossConnectASideRequestDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelPortB() int32`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *CrossConnectASideRequestDemarcationPoint) GetPatchPanelPortBOk() (*int32, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *CrossConnectASideRequestDemarcationPoint) SetPatchPanelPortB(v int32)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *CrossConnectASideRequestDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *CrossConnectASideRequestDemarcationPoint) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CrossConnectASideRequestDemarcationPoint) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CrossConnectASideRequestDemarcationPoint) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


