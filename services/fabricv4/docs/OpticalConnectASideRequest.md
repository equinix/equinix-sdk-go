# OpticalConnectASideRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Unique identifier of the patch panel. | [optional] 
**PatchPanelPortA** | Pointer to **string** | Specify the desired port number for Port A. &lt;br&gt; When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **string** | Specify the desired port number for Port B. &lt;br&gt; When ports are not provided, next available ports will be used. &lt;br&gt; Required for Connector type FC, SC and ST only.  | [optional] 
**ConnectorType** | Pointer to [**OpticalConnectPatchPanelFieldsConnectorType**](OpticalConnectPatchPanelFieldsConnectorType.md) |  | [optional] 

## Methods

### NewOpticalConnectASideRequest

`func NewOpticalConnectASideRequest() *OpticalConnectASideRequest`

NewOpticalConnectASideRequest instantiates a new OpticalConnectASideRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectASideRequestWithDefaults

`func NewOpticalConnectASideRequestWithDefaults() *OpticalConnectASideRequest`

NewOpticalConnectASideRequestWithDefaults instantiates a new OpticalConnectASideRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *OpticalConnectASideRequest) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *OpticalConnectASideRequest) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *OpticalConnectASideRequest) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *OpticalConnectASideRequest) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *OpticalConnectASideRequest) GetPatchPanelPortA() string`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *OpticalConnectASideRequest) GetPatchPanelPortAOk() (*string, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *OpticalConnectASideRequest) SetPatchPanelPortA(v string)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *OpticalConnectASideRequest) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *OpticalConnectASideRequest) GetPatchPanelPortB() string`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *OpticalConnectASideRequest) GetPatchPanelPortBOk() (*string, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *OpticalConnectASideRequest) SetPatchPanelPortB(v string)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *OpticalConnectASideRequest) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *OpticalConnectASideRequest) GetConnectorType() OpticalConnectPatchPanelFieldsConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *OpticalConnectASideRequest) GetConnectorTypeOk() (*OpticalConnectPatchPanelFieldsConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *OpticalConnectASideRequest) SetConnectorType(v OpticalConnectPatchPanelFieldsConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *OpticalConnectASideRequest) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


