# PatchPanelDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchPanelId** | Pointer to **string** | Patch panel ID / serial number | [optional] 
**Ibx** | Pointer to **string** | IBX location code | [optional] 
**CageId** | Pointer to **string** | Cage ID | [optional] 
**CabinetId** | Pointer to **string** | Cabinet ID | [optional] 
**AccountNumber** | Pointer to **string** | Customer cage account number | [optional] 
**AccountName** | Pointer to **string** | Customer cage account name | [optional] 
**DedicatedMediaType** | Pointer to **string** | Type of dedicated media | [optional] 
**PreWired** | Pointer to **bool** | Prewired status of cross connects. When &#x60;true&#x60;, cross connects are prewired | [optional] 
**Type** | Pointer to [**PatchPanelDetailsType**](PatchPanelDetailsType.md) |  | [optional] 
**IfcEnabled** | Pointer to **bool** | Identifies if intra-facility cable (IFC) connection is required for patch panel. When &#x60;true&#x60;, IFC connection is required. | [optional] 
**RackLocations** | Pointer to **string** | Rack location of patch panel | [optional] 
**InstallLocations** | Pointer to [**PatchPanelDetailsInstallLocations**](PatchPanelDetailsInstallLocations.md) |  | [optional] 
**InstallationRequired** | Pointer to **bool** | If &#x60;true&#x60;, Equinix will install cross connect from demarcation panel to customer equipment.  | [optional] [default to false]
**CircuitAvailable** | Pointer to **bool** | if &#x60;true&#x60;, circuits are available for patch panel | [optional] [default to false]
**AvailablePorts** | Pointer to **[]float32** | Individually identified ports that are available out of the total number of ports. | [optional] 
**ConnectionServices** | Pointer to [**[]ConnectionServicesDetailsInner**](ConnectionServicesDetailsInner.md) |  | [optional] 
**UsedPortsDetails** | Pointer to [**[]UserPortDetails**](UserPortDetails.md) |  | [optional] 
**ProvisioningType** | Pointer to [**PatchPanelDetailsProvisioningType**](PatchPanelDetailsProvisioningType.md) |  | [optional] 

## Methods

### NewPatchPanelDetails

`func NewPatchPanelDetails() *PatchPanelDetails`

NewPatchPanelDetails instantiates a new PatchPanelDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchPanelDetailsWithDefaults

`func NewPatchPanelDetailsWithDefaults() *PatchPanelDetails`

NewPatchPanelDetailsWithDefaults instantiates a new PatchPanelDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchPanelId

`func (o *PatchPanelDetails) GetPatchPanelId() string`

GetPatchPanelId returns the PatchPanelId field if non-nil, zero value otherwise.

### GetPatchPanelIdOk

`func (o *PatchPanelDetails) GetPatchPanelIdOk() (*string, bool)`

GetPatchPanelIdOk returns a tuple with the PatchPanelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchPanelId

`func (o *PatchPanelDetails) SetPatchPanelId(v string)`

SetPatchPanelId sets PatchPanelId field to given value.

### HasPatchPanelId

`func (o *PatchPanelDetails) HasPatchPanelId() bool`

HasPatchPanelId returns a boolean if a field has been set.

### GetIbx

`func (o *PatchPanelDetails) GetIbx() string`

GetIbx returns the Ibx field if non-nil, zero value otherwise.

### GetIbxOk

`func (o *PatchPanelDetails) GetIbxOk() (*string, bool)`

GetIbxOk returns a tuple with the Ibx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbx

`func (o *PatchPanelDetails) SetIbx(v string)`

SetIbx sets Ibx field to given value.

### HasIbx

`func (o *PatchPanelDetails) HasIbx() bool`

HasIbx returns a boolean if a field has been set.

### GetCageId

`func (o *PatchPanelDetails) GetCageId() string`

GetCageId returns the CageId field if non-nil, zero value otherwise.

### GetCageIdOk

`func (o *PatchPanelDetails) GetCageIdOk() (*string, bool)`

GetCageIdOk returns a tuple with the CageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCageId

`func (o *PatchPanelDetails) SetCageId(v string)`

SetCageId sets CageId field to given value.

### HasCageId

`func (o *PatchPanelDetails) HasCageId() bool`

HasCageId returns a boolean if a field has been set.

### GetCabinetId

`func (o *PatchPanelDetails) GetCabinetId() string`

GetCabinetId returns the CabinetId field if non-nil, zero value otherwise.

### GetCabinetIdOk

`func (o *PatchPanelDetails) GetCabinetIdOk() (*string, bool)`

GetCabinetIdOk returns a tuple with the CabinetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabinetId

`func (o *PatchPanelDetails) SetCabinetId(v string)`

SetCabinetId sets CabinetId field to given value.

### HasCabinetId

`func (o *PatchPanelDetails) HasCabinetId() bool`

HasCabinetId returns a boolean if a field has been set.

### GetAccountNumber

`func (o *PatchPanelDetails) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *PatchPanelDetails) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *PatchPanelDetails) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *PatchPanelDetails) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAccountName

`func (o *PatchPanelDetails) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *PatchPanelDetails) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *PatchPanelDetails) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *PatchPanelDetails) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetDedicatedMediaType

`func (o *PatchPanelDetails) GetDedicatedMediaType() string`

GetDedicatedMediaType returns the DedicatedMediaType field if non-nil, zero value otherwise.

### GetDedicatedMediaTypeOk

`func (o *PatchPanelDetails) GetDedicatedMediaTypeOk() (*string, bool)`

GetDedicatedMediaTypeOk returns a tuple with the DedicatedMediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicatedMediaType

`func (o *PatchPanelDetails) SetDedicatedMediaType(v string)`

SetDedicatedMediaType sets DedicatedMediaType field to given value.

### HasDedicatedMediaType

`func (o *PatchPanelDetails) HasDedicatedMediaType() bool`

HasDedicatedMediaType returns a boolean if a field has been set.

### GetPreWired

`func (o *PatchPanelDetails) GetPreWired() bool`

GetPreWired returns the PreWired field if non-nil, zero value otherwise.

### GetPreWiredOk

`func (o *PatchPanelDetails) GetPreWiredOk() (*bool, bool)`

GetPreWiredOk returns a tuple with the PreWired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreWired

`func (o *PatchPanelDetails) SetPreWired(v bool)`

SetPreWired sets PreWired field to given value.

### HasPreWired

`func (o *PatchPanelDetails) HasPreWired() bool`

HasPreWired returns a boolean if a field has been set.

### GetType

`func (o *PatchPanelDetails) GetType() PatchPanelDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PatchPanelDetails) GetTypeOk() (*PatchPanelDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PatchPanelDetails) SetType(v PatchPanelDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *PatchPanelDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIfcEnabled

`func (o *PatchPanelDetails) GetIfcEnabled() bool`

GetIfcEnabled returns the IfcEnabled field if non-nil, zero value otherwise.

### GetIfcEnabledOk

`func (o *PatchPanelDetails) GetIfcEnabledOk() (*bool, bool)`

GetIfcEnabledOk returns a tuple with the IfcEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfcEnabled

`func (o *PatchPanelDetails) SetIfcEnabled(v bool)`

SetIfcEnabled sets IfcEnabled field to given value.

### HasIfcEnabled

`func (o *PatchPanelDetails) HasIfcEnabled() bool`

HasIfcEnabled returns a boolean if a field has been set.

### GetRackLocations

`func (o *PatchPanelDetails) GetRackLocations() string`

GetRackLocations returns the RackLocations field if non-nil, zero value otherwise.

### GetRackLocationsOk

`func (o *PatchPanelDetails) GetRackLocationsOk() (*string, bool)`

GetRackLocationsOk returns a tuple with the RackLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRackLocations

`func (o *PatchPanelDetails) SetRackLocations(v string)`

SetRackLocations sets RackLocations field to given value.

### HasRackLocations

`func (o *PatchPanelDetails) HasRackLocations() bool`

HasRackLocations returns a boolean if a field has been set.

### GetInstallLocations

`func (o *PatchPanelDetails) GetInstallLocations() PatchPanelDetailsInstallLocations`

GetInstallLocations returns the InstallLocations field if non-nil, zero value otherwise.

### GetInstallLocationsOk

`func (o *PatchPanelDetails) GetInstallLocationsOk() (*PatchPanelDetailsInstallLocations, bool)`

GetInstallLocationsOk returns a tuple with the InstallLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallLocations

`func (o *PatchPanelDetails) SetInstallLocations(v PatchPanelDetailsInstallLocations)`

SetInstallLocations sets InstallLocations field to given value.

### HasInstallLocations

`func (o *PatchPanelDetails) HasInstallLocations() bool`

HasInstallLocations returns a boolean if a field has been set.

### GetInstallationRequired

`func (o *PatchPanelDetails) GetInstallationRequired() bool`

GetInstallationRequired returns the InstallationRequired field if non-nil, zero value otherwise.

### GetInstallationRequiredOk

`func (o *PatchPanelDetails) GetInstallationRequiredOk() (*bool, bool)`

GetInstallationRequiredOk returns a tuple with the InstallationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationRequired

`func (o *PatchPanelDetails) SetInstallationRequired(v bool)`

SetInstallationRequired sets InstallationRequired field to given value.

### HasInstallationRequired

`func (o *PatchPanelDetails) HasInstallationRequired() bool`

HasInstallationRequired returns a boolean if a field has been set.

### GetCircuitAvailable

`func (o *PatchPanelDetails) GetCircuitAvailable() bool`

GetCircuitAvailable returns the CircuitAvailable field if non-nil, zero value otherwise.

### GetCircuitAvailableOk

`func (o *PatchPanelDetails) GetCircuitAvailableOk() (*bool, bool)`

GetCircuitAvailableOk returns a tuple with the CircuitAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitAvailable

`func (o *PatchPanelDetails) SetCircuitAvailable(v bool)`

SetCircuitAvailable sets CircuitAvailable field to given value.

### HasCircuitAvailable

`func (o *PatchPanelDetails) HasCircuitAvailable() bool`

HasCircuitAvailable returns a boolean if a field has been set.

### GetAvailablePorts

`func (o *PatchPanelDetails) GetAvailablePorts() []float32`

GetAvailablePorts returns the AvailablePorts field if non-nil, zero value otherwise.

### GetAvailablePortsOk

`func (o *PatchPanelDetails) GetAvailablePortsOk() (*[]float32, bool)`

GetAvailablePortsOk returns a tuple with the AvailablePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePorts

`func (o *PatchPanelDetails) SetAvailablePorts(v []float32)`

SetAvailablePorts sets AvailablePorts field to given value.

### HasAvailablePorts

`func (o *PatchPanelDetails) HasAvailablePorts() bool`

HasAvailablePorts returns a boolean if a field has been set.

### GetConnectionServices

`func (o *PatchPanelDetails) GetConnectionServices() []ConnectionServicesDetailsInner`

GetConnectionServices returns the ConnectionServices field if non-nil, zero value otherwise.

### GetConnectionServicesOk

`func (o *PatchPanelDetails) GetConnectionServicesOk() (*[]ConnectionServicesDetailsInner, bool)`

GetConnectionServicesOk returns a tuple with the ConnectionServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServices

`func (o *PatchPanelDetails) SetConnectionServices(v []ConnectionServicesDetailsInner)`

SetConnectionServices sets ConnectionServices field to given value.

### HasConnectionServices

`func (o *PatchPanelDetails) HasConnectionServices() bool`

HasConnectionServices returns a boolean if a field has been set.

### GetUsedPortsDetails

`func (o *PatchPanelDetails) GetUsedPortsDetails() []UserPortDetails`

GetUsedPortsDetails returns the UsedPortsDetails field if non-nil, zero value otherwise.

### GetUsedPortsDetailsOk

`func (o *PatchPanelDetails) GetUsedPortsDetailsOk() (*[]UserPortDetails, bool)`

GetUsedPortsDetailsOk returns a tuple with the UsedPortsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedPortsDetails

`func (o *PatchPanelDetails) SetUsedPortsDetails(v []UserPortDetails)`

SetUsedPortsDetails sets UsedPortsDetails field to given value.

### HasUsedPortsDetails

`func (o *PatchPanelDetails) HasUsedPortsDetails() bool`

HasUsedPortsDetails returns a boolean if a field has been set.

### GetProvisioningType

`func (o *PatchPanelDetails) GetProvisioningType() PatchPanelDetailsProvisioningType`

GetProvisioningType returns the ProvisioningType field if non-nil, zero value otherwise.

### GetProvisioningTypeOk

`func (o *PatchPanelDetails) GetProvisioningTypeOk() (*PatchPanelDetailsProvisioningType, bool)`

GetProvisioningTypeOk returns a tuple with the ProvisioningType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningType

`func (o *PatchPanelDetails) SetProvisioningType(v PatchPanelDetailsProvisioningType)`

SetProvisioningType sets ProvisioningType field to given value.

### HasProvisioningType

`func (o *PatchPanelDetails) HasProvisioningType() bool`

HasProvisioningType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


