# LoaDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CageUniqueSpaceId** | Pointer to **string** | Unique identifier of the Cage. | [optional] 
**PatchPanelId** | Pointer to **string** | Unique identifier of the Patch Panel.  | [optional] 
**PatchPanelPortA** | Pointer to **int32** | Specify the desired port number. &lt;br&gt; When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **int32** | Specify the desired port number. &lt;br&gt; When ports are not provided, next available ports will be used. &lt;br&gt; Required for Connector type FC and ST only.  | [optional] 
**ConnectorType** | Pointer to [**LoaPatchPanelConnectorType**](LoaPatchPanelConnectorType.md) |  | [optional] 

## Methods

### NewLoaDemarcationPoint

`func NewLoaDemarcationPoint() *LoaDemarcationPoint`

NewLoaDemarcationPoint instantiates a new LoaDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaDemarcationPointWithDefaults

`func NewLoaDemarcationPointWithDefaults() *LoaDemarcationPoint`

NewLoaDemarcationPointWithDefaults instantiates a new LoaDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCageUniqueSpaceId

`func (o *LoaDemarcationPoint) GetCageUniqueSpaceId() string`

GetCageUniqueSpaceId returns the CageUniqueSpaceId field if non-nil, zero value otherwise.

### GetCageUniqueSpaceIdOk

`func (o *LoaDemarcationPoint) GetCageUniqueSpaceIdOk() (*string, bool)`

GetCageUniqueSpaceIdOk returns a tuple with the CageUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageUniqueSpaceId

`func (o *LoaDemarcationPoint) SetCageUniqueSpaceId(v string)`

SetCageUniqueSpaceId sets CageUniqueSpaceId field to given value.

### HasCageUniqueSpaceId

`func (o *LoaDemarcationPoint) HasCageUniqueSpaceId() bool`

HasCageUniqueSpaceId returns a boolean if a field has been set.

### GetPatchPanelId

`func (o *LoaDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *LoaDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *LoaDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *LoaDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *LoaDemarcationPoint) GetPatchPanelPortA() int32`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *LoaDemarcationPoint) GetPatchPanelPortAOk() (*int32, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *LoaDemarcationPoint) SetPatchPanelPortA(v int32)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *LoaDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *LoaDemarcationPoint) GetPatchPanelPortB() int32`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *LoaDemarcationPoint) GetPatchPanelPortBOk() (*int32, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *LoaDemarcationPoint) SetPatchPanelPortB(v int32)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *LoaDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *LoaDemarcationPoint) GetConnectorType() LoaPatchPanelConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *LoaDemarcationPoint) GetConnectorTypeOk() (*LoaPatchPanelConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *LoaDemarcationPoint) SetConnectorType(v LoaPatchPanelConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *LoaDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


