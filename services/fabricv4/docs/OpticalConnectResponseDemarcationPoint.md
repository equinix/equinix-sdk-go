# OpticalConnectResponseDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Unique identifier of the patch panel. | [optional] 
**PatchPanelPortA** | Pointer to **string** | Specify the desired port number for Port A. &lt;br&gt; When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **string** | Specify the desired port number for Port B. &lt;br&gt; When ports are not provided, next available ports will be used. &lt;br&gt; Required for Connector type FC, SC and ST only.  | [optional] 
**ConnectorType** | Pointer to [**OpticalConnectDemarcationPointConnectorType**](OpticalConnectDemarcationPointConnectorType.md) |  | [optional] 
**CageUniqueSpaceId** | Pointer to **string** | Unique identifier of the cage. | [optional] 
**CabinetUniqueSpaceId** | Pointer to **string** | Unique identifier of the cabinet. | [optional] 

## Methods

### NewOpticalConnectResponseDemarcationPoint

`func NewOpticalConnectResponseDemarcationPoint() *OpticalConnectResponseDemarcationPoint`

NewOpticalConnectResponseDemarcationPoint instantiates a new OpticalConnectResponseDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectResponseDemarcationPointWithDefaults

`func NewOpticalConnectResponseDemarcationPointWithDefaults() *OpticalConnectResponseDemarcationPoint`

NewOpticalConnectResponseDemarcationPointWithDefaults instantiates a new OpticalConnectResponseDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *OpticalConnectResponseDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *OpticalConnectResponseDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelPortA() string`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelPortAOk() (*string, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *OpticalConnectResponseDemarcationPoint) SetPatchPanelPortA(v string)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *OpticalConnectResponseDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelPortB() string`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *OpticalConnectResponseDemarcationPoint) GetPatchPanelPortBOk() (*string, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *OpticalConnectResponseDemarcationPoint) SetPatchPanelPortB(v string)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *OpticalConnectResponseDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *OpticalConnectResponseDemarcationPoint) GetConnectorType() OpticalConnectDemarcationPointConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *OpticalConnectResponseDemarcationPoint) GetConnectorTypeOk() (*OpticalConnectDemarcationPointConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *OpticalConnectResponseDemarcationPoint) SetConnectorType(v OpticalConnectDemarcationPointConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *OpticalConnectResponseDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetCageUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) GetCageUniqueSpaceId() string`

GetCageUniqueSpaceId returns the CageUniqueSpaceId field if non-nil, zero value otherwise.

### GetCageUniqueSpaceIdOk

`func (o *OpticalConnectResponseDemarcationPoint) GetCageUniqueSpaceIdOk() (*string, bool)`

GetCageUniqueSpaceIdOk returns a tuple with the CageUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) SetCageUniqueSpaceId(v string)`

SetCageUniqueSpaceId sets CageUniqueSpaceId field to given value.

### HasCageUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) HasCageUniqueSpaceId() bool`

HasCageUniqueSpaceId returns a boolean if a field has been set.

### GetCabinetUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) GetCabinetUniqueSpaceId() string`

GetCabinetUniqueSpaceId returns the CabinetUniqueSpaceId field if non-nil, zero value otherwise.

### GetCabinetUniqueSpaceIdOk

`func (o *OpticalConnectResponseDemarcationPoint) GetCabinetUniqueSpaceIdOk() (*string, bool)`

GetCabinetUniqueSpaceIdOk returns a tuple with the CabinetUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) SetCabinetUniqueSpaceId(v string)`

SetCabinetUniqueSpaceId sets CabinetUniqueSpaceId field to given value.

### HasCabinetUniqueSpaceId

`func (o *OpticalConnectResponseDemarcationPoint) HasCabinetUniqueSpaceId() bool`

HasCabinetUniqueSpaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


