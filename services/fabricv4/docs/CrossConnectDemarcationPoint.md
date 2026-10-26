# CrossConnectDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Patch Panel Identifier. | [optional] 
**PatchPanelPortA** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **int32** | Specify the desired port number. When ports are not provided, next available ports will be used. Required for Connector type FC and ST only.  | [optional] 
**ConnectorType** | Pointer to **string** | Connector Type. | [optional] 

## Methods

### NewCrossConnectDemarcationPoint

`func NewCrossConnectDemarcationPoint() *CrossConnectDemarcationPoint`

NewCrossConnectDemarcationPoint instantiates a new CrossConnectDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectDemarcationPointWithDefaults

`func NewCrossConnectDemarcationPointWithDefaults() *CrossConnectDemarcationPoint`

NewCrossConnectDemarcationPointWithDefaults instantiates a new CrossConnectDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *CrossConnectDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *CrossConnectDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *CrossConnectDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *CrossConnectDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *CrossConnectDemarcationPoint) GetPatchPanelPortA() int32`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *CrossConnectDemarcationPoint) GetPatchPanelPortAOk() (*int32, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *CrossConnectDemarcationPoint) SetPatchPanelPortA(v int32)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *CrossConnectDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *CrossConnectDemarcationPoint) GetPatchPanelPortB() int32`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *CrossConnectDemarcationPoint) GetPatchPanelPortBOk() (*int32, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *CrossConnectDemarcationPoint) SetPatchPanelPortB(v int32)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *CrossConnectDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *CrossConnectDemarcationPoint) GetConnectorType() string`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *CrossConnectDemarcationPoint) GetConnectorTypeOk() (*string, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *CrossConnectDemarcationPoint) SetConnectorType(v string)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *CrossConnectDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


