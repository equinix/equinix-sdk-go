# PatchPanelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Patch panel unique ID / serial number | [optional] 
**AvailablePortCount** | Pointer to **float32** | Available ports for use | [optional] 
**PatchPanelReferenceId** | Pointer to **string** | Unique reference ID associated with patch panel | [optional] 
**IfcEnabled** | Pointer to **bool** | Patch panel will support only intra facility when &#x60;true&#x60; | [optional] [default to false]
**ProvisioningType** | Pointer to [**PatchPanelDetailsProvisioningType**](PatchPanelDetailsProvisioningType.md) |  | [optional] 

## Methods

### NewPatchPanelResponse

`func NewPatchPanelResponse() *PatchPanelResponse`

NewPatchPanelResponse instantiates a new PatchPanelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPanelResponseWithDefaults

`func NewPatchPanelResponseWithDefaults() *PatchPanelResponse`

NewPatchPanelResponseWithDefaults instantiates a new PatchPanelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *PatchPanelResponse) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *PatchPanelResponse) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *PatchPanelResponse) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *PatchPanelResponse) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetAvailablePortCount

`func (o *PatchPanelResponse) GetAvailablePortCount() float32`

GetAvailablePortCount returns the AvailablePortCount field if non-nil, zero value otherwise.

### GetAvailablePortCountOk

`func (o *PatchPanelResponse) GetAvailablePortCountOk() (*float32, bool)`

GetAvailablePortCountOk returns a tuple with the AvailablePortCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePortCount

`func (o *PatchPanelResponse) SetAvailablePortCount(v float32)`

SetAvailablePortCount sets AvailablePortCount field to given value.

### HasAvailablePortCount

`func (o *PatchPanelResponse) HasAvailablePortCount() bool`

HasAvailablePortCount returns a boolean if a field has been set.

### GetPatchPanelReferenceId

`func (o *PatchPanelResponse) GetPatchPanelReferenceId() string`

GetPatchPanelReferenceId returns the PatchPanelReferenceId field if non-nil, zero value otherwise.

### GetPatchPanelReferenceIdOk

`func (o *PatchPanelResponse) GetPatchPanelReferenceIdOk() (*string, bool)`

GetPatchPanelReferenceIdOk returns a tuple with the PatchPanelReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelReferenceId

`func (o *PatchPanelResponse) SetPatchPanelReferenceId(v string)`

SetPatchPanelReferenceId sets PatchPanelReferenceId field to given value.

### HasPatchPanelReferenceId

`func (o *PatchPanelResponse) HasPatchPanelReferenceId() bool`

HasPatchPanelReferenceId returns a boolean if a field has been set.

### GetIfcEnabled

`func (o *PatchPanelResponse) GetIfcEnabled() bool`

GetIfcEnabled returns the IfcEnabled field if non-nil, zero value otherwise.

### GetIfcEnabledOk

`func (o *PatchPanelResponse) GetIfcEnabledOk() (*bool, bool)`

GetIfcEnabledOk returns a tuple with the IfcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfcEnabled

`func (o *PatchPanelResponse) SetIfcEnabled(v bool)`

SetIfcEnabled sets IfcEnabled field to given value.

### HasIfcEnabled

`func (o *PatchPanelResponse) HasIfcEnabled() bool`

HasIfcEnabled returns a boolean if a field has been set.

### GetProvisioningType

`func (o *PatchPanelResponse) GetProvisioningType() PatchPanelDetailsProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *PatchPanelResponse) GetProvisioningTypeOk() (*PatchPanelDetailsProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *PatchPanelResponse) SetProvisioningType(v PatchPanelDetailsProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.

### HasProvisioningType

`func (o *PatchPanelResponse) HasProvisioningType() bool`

HasProvisioningType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


