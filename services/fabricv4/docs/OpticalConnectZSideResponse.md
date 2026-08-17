# OpticalConnectZSideResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Unique identifier of the patch panel. | [optional] 
**PatchPanelPortA** | Pointer to **string** | Specify the desired port number for Port A. &lt;br&gt; When ports are not provided, next available ports will be used.  | [optional] 
**PatchPanelPortB** | Pointer to **string** | Specify the desired port number for Port B. &lt;br&gt; When ports are not provided, next available ports will be used. &lt;br&gt; Required for Connector type FC, SC and ST only.  | [optional] 
**ConnectorType** | Pointer to [**OpticalConnectPatchPanelFieldsConnectorType**](OpticalConnectPatchPanelFieldsConnectorType.md) |  | [optional] 
**CageUniqueSpaceId** | Pointer to **string** | Unique identifier of the cage. | [optional] 
**CabinetUniqueSpaceId** | Pointer to **string** | Unique identifier of the cabinet. | [optional] 
**Location** | Pointer to [**OpticalConnectLocation**](OpticalConnectLocation.md) |  | [optional] 
**Loa** | Pointer to [**OpticalConnectLOA**](OpticalConnectLOA.md) |  | [optional] 

## Methods

### NewOpticalConnectZSideResponse

`func NewOpticalConnectZSideResponse() *OpticalConnectZSideResponse`

NewOpticalConnectZSideResponse instantiates a new OpticalConnectZSideResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpticalConnectZSideResponseWithDefaults

`func NewOpticalConnectZSideResponseWithDefaults() *OpticalConnectZSideResponse`

NewOpticalConnectZSideResponseWithDefaults instantiates a new OpticalConnectZSideResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *OpticalConnectZSideResponse) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *OpticalConnectZSideResponse) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *OpticalConnectZSideResponse) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *OpticalConnectZSideResponse) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetPatchPanelPortA

`func (o *OpticalConnectZSideResponse) GetPatchPanelPortA() string`

GetPatchPanelPortA returns the PatchPanelPortA field if non-nil, zero value otherwise.

### GetPatchPanelPortAOk

`func (o *OpticalConnectZSideResponse) GetPatchPanelPortAOk() (*string, bool)`

GetPatchPanelPortAOk returns a tuple with the PatchPanelPortA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortA

`func (o *OpticalConnectZSideResponse) SetPatchPanelPortA(v string)`

SetPatchPanelPortA sets PatchPanelPortA field to given value.

### HasPatchPanelPortA

`func (o *OpticalConnectZSideResponse) HasPatchPanelPortA() bool`

HasPatchPanelPortA returns a boolean if a field has been set.

### GetPatchPanelPortB

`func (o *OpticalConnectZSideResponse) GetPatchPanelPortB() string`

GetPatchPanelPortB returns the PatchPanelPortB field if non-nil, zero value otherwise.

### GetPatchPanelPortBOk

`func (o *OpticalConnectZSideResponse) GetPatchPanelPortBOk() (*string, bool)`

GetPatchPanelPortBOk returns a tuple with the PatchPanelPortB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelPortB

`func (o *OpticalConnectZSideResponse) SetPatchPanelPortB(v string)`

SetPatchPanelPortB sets PatchPanelPortB field to given value.

### HasPatchPanelPortB

`func (o *OpticalConnectZSideResponse) HasPatchPanelPortB() bool`

HasPatchPanelPortB returns a boolean if a field has been set.

### GetConnectorType

`func (o *OpticalConnectZSideResponse) GetConnectorType() OpticalConnectPatchPanelFieldsConnectorType`

GetConnectorType returns the ConnectorType field if non-nil, zero value otherwise.

### GetConnectorTypeOk

`func (o *OpticalConnectZSideResponse) GetConnectorTypeOk() (*OpticalConnectPatchPanelFieldsConnectorType, bool)`

GetConnectorTypeOk returns a tuple with the ConnectorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorType

`func (o *OpticalConnectZSideResponse) SetConnectorType(v OpticalConnectPatchPanelFieldsConnectorType)`

SetConnectorType sets ConnectorType field to given value.

### HasConnectorType

`func (o *OpticalConnectZSideResponse) HasConnectorType() bool`

HasConnectorType returns a boolean if a field has been set.

### GetCageUniqueSpaceId

`func (o *OpticalConnectZSideResponse) GetCageUniqueSpaceId() string`

GetCageUniqueSpaceId returns the CageUniqueSpaceId field if non-nil, zero value otherwise.

### GetCageUniqueSpaceIdOk

`func (o *OpticalConnectZSideResponse) GetCageUniqueSpaceIdOk() (*string, bool)`

GetCageUniqueSpaceIdOk returns a tuple with the CageUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageUniqueSpaceId

`func (o *OpticalConnectZSideResponse) SetCageUniqueSpaceId(v string)`

SetCageUniqueSpaceId sets CageUniqueSpaceId field to given value.

### HasCageUniqueSpaceId

`func (o *OpticalConnectZSideResponse) HasCageUniqueSpaceId() bool`

HasCageUniqueSpaceId returns a boolean if a field has been set.

### GetCabinetUniqueSpaceId

`func (o *OpticalConnectZSideResponse) GetCabinetUniqueSpaceId() string`

GetCabinetUniqueSpaceId returns the CabinetUniqueSpaceId field if non-nil, zero value otherwise.

### GetCabinetUniqueSpaceIdOk

`func (o *OpticalConnectZSideResponse) GetCabinetUniqueSpaceIdOk() (*string, bool)`

GetCabinetUniqueSpaceIdOk returns a tuple with the CabinetUniqueSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetUniqueSpaceId

`func (o *OpticalConnectZSideResponse) SetCabinetUniqueSpaceId(v string)`

SetCabinetUniqueSpaceId sets CabinetUniqueSpaceId field to given value.

### HasCabinetUniqueSpaceId

`func (o *OpticalConnectZSideResponse) HasCabinetUniqueSpaceId() bool`

HasCabinetUniqueSpaceId returns a boolean if a field has been set.

### GetLocation

`func (o *OpticalConnectZSideResponse) GetLocation() OpticalConnectLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OpticalConnectZSideResponse) GetLocationOk() (*OpticalConnectLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OpticalConnectZSideResponse) SetLocation(v OpticalConnectLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OpticalConnectZSideResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLoa

`func (o *OpticalConnectZSideResponse) GetLoa() OpticalConnectLOA`

GetLoa returns the Loa field if non-nil, zero value otherwise.

### GetLoaOk

`func (o *OpticalConnectZSideResponse) GetLoaOk() (*OpticalConnectLOA, bool)`

GetLoaOk returns a tuple with the Loa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoa

`func (o *OpticalConnectZSideResponse) SetLoa(v OpticalConnectLOA)`

SetLoa sets Loa field to given value.

### HasLoa

`func (o *OpticalConnectZSideResponse) HasLoa() bool`

HasLoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


