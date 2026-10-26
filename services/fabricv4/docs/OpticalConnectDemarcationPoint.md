# OpticalConnectDemarcationPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Unique identifier of the patch panel. | [optional] 
**PatchPanelPortA** | Pointer to **string** | Specify the desired port number for Port A. &lt;br&gt; When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **string** | Specify the desired port number for Port B. &lt;br&gt; When ports are not provided, next available ports will be used. &lt;br&gt; Required for Connector type FC, SC and ST only.  | [optional] 
**ConnectorType** | Pointer to [**OpticalConnectDemarcationPointConnectorType**](OpticalConnectDemarcationPointConnectorType.md) |  | [optional] 

## Methods

### NewOpticalConnectDemarcationPoint

`func NewOpticalConnectDemarcationPoint() *OpticalConnectDemarcationPoint`

NewOpticalConnectDemarcationPoint instantiates a new OpticalConnectDemarcationPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectDemarcationPointWithDefaults

`func NewOpticalConnectDemarcationPointWithDefaults() *OpticalConnectDemarcationPoint`

NewOpticalConnectDemarcationPointWithDefaults instantiates a new OpticalConnectDemarcationPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *OpticalConnectDemarcationPoint) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *OpticalConnectDemarcationPoint) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelPortA() string`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelPortAOk() (*string, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *OpticalConnectDemarcationPoint) SetPatchPanelPortA(v string)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *OpticalConnectDemarcationPoint) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelPortB() string`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *OpticalConnectDemarcationPoint) GetPatchPanelPortBOk() (*string, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *OpticalConnectDemarcationPoint) SetPatchPanelPortB(v string)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *OpticalConnectDemarcationPoint) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *OpticalConnectDemarcationPoint) GetConnectorType() OpticalConnectDemarcationPointConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *OpticalConnectDemarcationPoint) GetConnectorTypeOk() (*OpticalConnectDemarcationPointConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *OpticalConnectDemarcationPoint) SetConnectorType(v OpticalConnectDemarcationPointConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *OpticalConnectDemarcationPoint) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


