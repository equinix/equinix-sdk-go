# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlwaysPxe** | Pointer to **bool** |  | [optional] 
**BillingCycle** | Pointer to **string** |  | [optional] 
**BondingMode** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**DeviceCreatedBy**](DeviceCreatedBy.md) |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] [default to {}]
**Description** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**Facility**](Facility.md) |  | [optional] 
**FirmwareSetId** | Pointer to **string** | The UUID of the firmware set to associate with the device. | [optional] 
**HardwareReservation** | Pointer to [**HardwareReservation**](HardwareReservation.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**IpAddresses** | Pointer to [**[]IPAssignment**](IPAssignment.md) |  | [optional] 
**IpxeScriptUrl** | Pointer to **string** |  | [optional] 
**Iqn** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** | Prevents accidental deletion of this resource when set to true. | [optional] 
**Metro** | Pointer to [**DeviceMetro**](DeviceMetro.md) |  | [optional] 
**NetworkFrozen** | Pointer to **bool** | Whether network mode changes such as converting to/from Layer2 or Layer3 networking, bonding or disbonding network interfaces are permitted for the device. | [optional] 
**NetworkPorts** | Pointer to [**[]Port**](Port.md) | By default, servers at Equinix Metal are configured in a “bonded” mode using LACP (Link Aggregation Control Protocol). Each 2-NIC server is configured with a single bond (namely bond0) with both interfaces eth0 and eth1 as members of the bond in a default Layer 3 mode. Some device plans may have a different number of ports and bonds available. | [optional] 
**OperatingSystem** | Pointer to [**OperatingSystem**](OperatingSystem.md) |  | [optional] 
**Actions** | Pointer to [**[]DeviceActionsInner**](DeviceActionsInner.md) | Actions supported by the device instance. | [optional] 
**Plan** | Pointer to [**Plan**](Plan.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectLite** | Pointer to [**DeviceProjectLite**](DeviceProjectLite.md) |  | [optional] 
**ProvisioningEvents** | Pointer to [**[]Event**](Event.md) |  | [optional] 
**ProvisioningPercentage** | Pointer to **float32** | Only visible while device provisioning | [optional] 
**RootPassword** | Pointer to **string** | Root password is automatically generated when server is provisioned and it is removed after 24 hours | [optional] 
**ShortId** | Pointer to **string** |  | [optional] 
**SpotInstance** | Pointer to **bool** | Whether or not the device is a spot instance. | [optional] 
**SpotPriceMax** | Pointer to **float32** | The maximum price per hour you are willing to pay to keep this spot instance.  If you are outbid, the termination will be set allowing two minutes before shutdown. | [optional] 
**SshKeys** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**State** | Pointer to [**DeviceState**](DeviceState.md) |  | [optional] 
**Storage** | Pointer to [**Storage**](Storage.md) |  | [optional] 
**SwitchUuid** | Pointer to **string** | Switch short id. This can be used to determine if two devices are connected to the same switch, for example. | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TerminationTime** | Pointer to **time.Time** | When the device will be terminated. If you don&#39;t supply timezone info, the timestamp is assumed to be in UTC.  This is commonly set in advance for ephemeral spot market instances but this field may also be set with on-demand and reservation instances to automatically delete the resource at a given time. The termination time can also be used to release a hardware reservation instance at a given time, keeping the reservation open for other uses.  On a spot market device, the termination time will be set automatically when outbid. | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Userdata** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**Sos** | Pointer to **string** | Hostname used to connect to the instance via the SOS (Serial over SSH) out-of-band console. | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlwaysPxe

`func (o *Device) GetAlwaysPxe() bool`

GetAlwaysPxe returns the AlwaysPxe field if non-nil, zero value otherwise.

### GetAlwaysPxeOk

`func (o *Device) GetAlwaysPxeOk() (*bool, bool)`

GetAlwaysPxeOk returns a tuple with the AlwaysPxe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlwaysPxe

`func (o *Device) SetAlwaysPxe(v bool)`

SetAlwaysPxe sets AlwaysPxe field to given value.

### HasAlwaysPxe

`func (o *Device) HasAlwaysPxe() bool`

HasAlwaysPxe returns a boolean if a field has been set.

### GetBillingCycle

`func (o *Device) GetBillingCycle() string`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *Device) GetBillingCycleOk() (*string, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *Device) SetBillingCycle(v string)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *Device) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetBondingMode

`func (o *Device) GetBondingMode() int32`

GetBondingMode returns the BondingMode field if non-nil, zero value otherwise.

### GetBondingModeOk

`func (o *Device) GetBondingModeOk() (*int32, bool)`

GetBondingModeOk returns a tuple with the BondingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondingMode

`func (o *Device) SetBondingMode(v int32)`

SetBondingMode sets BondingMode field to given value.

### HasBondingMode

`func (o *Device) HasBondingMode() bool`

HasBondingMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Device) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Device) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Device) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Device) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *Device) GetCreatedBy() DeviceCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Device) GetCreatedByOk() (*DeviceCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Device) SetCreatedBy(v DeviceCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Device) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCustomdata

`func (o *Device) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *Device) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *Device) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *Device) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetDescription

`func (o *Device) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Device) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Device) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Device) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacility

`func (o *Device) GetFacility() Facility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *Device) GetFacilityOk() (*Facility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *Device) SetFacility(v Facility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *Device) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetFirmwareSetId

`func (o *Device) GetFirmwareSetId() string`

GetFirmwareSetId returns the FirmwareSetId field if non-nil, zero value otherwise.

### GetFirmwareSetIdOk

`func (o *Device) GetFirmwareSetIdOk() (*string, bool)`

GetFirmwareSetIdOk returns a tuple with the FirmwareSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwareSetId

`func (o *Device) SetFirmwareSetId(v string)`

SetFirmwareSetId sets FirmwareSetId field to given value.

### HasFirmwareSetId

`func (o *Device) HasFirmwareSetId() bool`

HasFirmwareSetId returns a boolean if a field has been set.

### GetHardwareReservation

`func (o *Device) GetHardwareReservation() HardwareReservation`

GetHardwareReservation returns the HardwareReservation field if non-nil, zero value otherwise.

### GetHardwareReservationOk

`func (o *Device) GetHardwareReservationOk() (*HardwareReservation, bool)`

GetHardwareReservationOk returns a tuple with the HardwareReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareReservation

`func (o *Device) SetHardwareReservation(v HardwareReservation)`

SetHardwareReservation sets HardwareReservation field to given value.

### HasHardwareReservation

`func (o *Device) HasHardwareReservation() bool`

HasHardwareReservation returns a boolean if a field has been set.

### GetHostname

`func (o *Device) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *Device) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *Device) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *Device) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetHref

`func (o *Device) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Device) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Device) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Device) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Device) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *Device) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Device) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Device) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *Device) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetIpAddresses

`func (o *Device) GetIpAddresses() []IPAssignment`

GetIpAddresses returns the IpAddresses field if non-nil, zero value otherwise.

### GetIpAddressesOk

`func (o *Device) GetIpAddressesOk() (*[]IPAssignment, bool)`

GetIpAddressesOk returns a tuple with the IpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddresses

`func (o *Device) SetIpAddresses(v []IPAssignment)`

SetIpAddresses sets IpAddresses field to given value.

### HasIpAddresses

`func (o *Device) HasIpAddresses() bool`

HasIpAddresses returns a boolean if a field has been set.

### GetIpxeScriptUrl

`func (o *Device) GetIpxeScriptUrl() string`

GetIpxeScriptUrl returns the IpxeScriptUrl field if non-nil, zero value otherwise.

### GetIpxeScriptUrlOk

`func (o *Device) GetIpxeScriptUrlOk() (*string, bool)`

GetIpxeScriptUrlOk returns a tuple with the IpxeScriptUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpxeScriptUrl

`func (o *Device) SetIpxeScriptUrl(v string)`

SetIpxeScriptUrl sets IpxeScriptUrl field to given value.

### HasIpxeScriptUrl

`func (o *Device) HasIpxeScriptUrl() bool`

HasIpxeScriptUrl returns a boolean if a field has been set.

### GetIqn

`func (o *Device) GetIqn() string`

GetIqn returns the Iqn field if non-nil, zero value otherwise.

### GetIqnOk

`func (o *Device) GetIqnOk() (*string, bool)`

GetIqnOk returns a tuple with the Iqn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqn

`func (o *Device) SetIqn(v string)`

SetIqn sets Iqn field to given value.

### HasIqn

`func (o *Device) HasIqn() bool`

HasIqn returns a boolean if a field has been set.

### GetLocked

`func (o *Device) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Device) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Device) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Device) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMetro

`func (o *Device) GetMetro() DeviceMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *Device) GetMetroOk() (*DeviceMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *Device) SetMetro(v DeviceMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *Device) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetworkFrozen

`func (o *Device) GetNetworkFrozen() bool`

GetNetworkFrozen returns the NetworkFrozen field if non-nil, zero value otherwise.

### GetNetworkFrozenOk

`func (o *Device) GetNetworkFrozenOk() (*bool, bool)`

GetNetworkFrozenOk returns a tuple with the NetworkFrozen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFrozen

`func (o *Device) SetNetworkFrozen(v bool)`

SetNetworkFrozen sets NetworkFrozen field to given value.

### HasNetworkFrozen

`func (o *Device) HasNetworkFrozen() bool`

HasNetworkFrozen returns a boolean if a field has been set.

### GetNetworkPorts

`func (o *Device) GetNetworkPorts() []Port`

GetNetworkPorts returns the NetworkPorts field if non-nil, zero value otherwise.

### GetNetworkPortsOk

`func (o *Device) GetNetworkPortsOk() (*[]Port, bool)`

GetNetworkPortsOk returns a tuple with the NetworkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPorts

`func (o *Device) SetNetworkPorts(v []Port)`

SetNetworkPorts sets NetworkPorts field to given value.

### HasNetworkPorts

`func (o *Device) HasNetworkPorts() bool`

HasNetworkPorts returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *Device) GetOperatingSystem() OperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *Device) GetOperatingSystemOk() (*OperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *Device) SetOperatingSystem(v OperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *Device) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetActions

`func (o *Device) GetActions() []DeviceActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Device) GetActionsOk() (*[]DeviceActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Device) SetActions(v []DeviceActionsInner)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Device) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetPlan

`func (o *Device) GetPlan() Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *Device) GetPlanOk() (*Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *Device) SetPlan(v Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *Device) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetProject

`func (o *Device) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Device) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Device) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *Device) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *Device) GetProjectLite() DeviceProjectLite`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *Device) GetProjectLiteOk() (*DeviceProjectLite, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *Device) SetProjectLite(v DeviceProjectLite)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *Device) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetProvisioningEvents

`func (o *Device) GetProvisioningEvents() []Event`

GetProvisioningEvents returns the ProvisioningEvents field if non-nil, zero value otherwise.

### GetProvisioningEventsOk

`func (o *Device) GetProvisioningEventsOk() (*[]Event, bool)`

GetProvisioningEventsOk returns a tuple with the ProvisioningEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningEvents

`func (o *Device) SetProvisioningEvents(v []Event)`

SetProvisioningEvents sets ProvisioningEvents field to given value.

### HasProvisioningEvents

`func (o *Device) HasProvisioningEvents() bool`

HasProvisioningEvents returns a boolean if a field has been set.

### GetProvisioningPercentage

`func (o *Device) GetProvisioningPercentage() float32`

GetProvisioningPercentage returns the ProvisioningPercentage field if non-nil, zero value otherwise.

### GetProvisioningPercentageOk

`func (o *Device) GetProvisioningPercentageOk() (*float32, bool)`

GetProvisioningPercentageOk returns a tuple with the ProvisioningPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningPercentage

`func (o *Device) SetProvisioningPercentage(v float32)`

SetProvisioningPercentage sets ProvisioningPercentage field to given value.

### HasProvisioningPercentage

`func (o *Device) HasProvisioningPercentage() bool`

HasProvisioningPercentage returns a boolean if a field has been set.

### GetRootPassword

`func (o *Device) GetRootPassword() string`

GetRootPassword returns the RootPassword field if non-nil, zero value otherwise.

### GetRootPasswordOk

`func (o *Device) GetRootPasswordOk() (*string, bool)`

GetRootPasswordOk returns a tuple with the RootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPassword

`func (o *Device) SetRootPassword(v string)`

SetRootPassword sets RootPassword field to given value.

### HasRootPassword

`func (o *Device) HasRootPassword() bool`

HasRootPassword returns a boolean if a field has been set.

### GetShortId

`func (o *Device) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *Device) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *Device) SetShortId(v string)`

SetShortId sets ShortId field to given value.

### HasShortId

`func (o *Device) HasShortId() bool`

HasShortId returns a boolean if a field has been set.

### GetSpotInstance

`func (o *Device) GetSpotInstance() bool`

GetSpotInstance returns the SpotInstance field if non-nil, zero value otherwise.

### GetSpotInstanceOk

`func (o *Device) GetSpotInstanceOk() (*bool, bool)`

GetSpotInstanceOk returns a tuple with the SpotInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotInstance

`func (o *Device) SetSpotInstance(v bool)`

SetSpotInstance sets SpotInstance field to given value.

### HasSpotInstance

`func (o *Device) HasSpotInstance() bool`

HasSpotInstance returns a boolean if a field has been set.

### GetSpotPriceMax

`func (o *Device) GetSpotPriceMax() float32`

GetSpotPriceMax returns the SpotPriceMax field if non-nil, zero value otherwise.

### GetSpotPriceMaxOk

`func (o *Device) GetSpotPriceMaxOk() (*float32, bool)`

GetSpotPriceMaxOk returns a tuple with the SpotPriceMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotPriceMax

`func (o *Device) SetSpotPriceMax(v float32)`

SetSpotPriceMax sets SpotPriceMax field to given value.

### HasSpotPriceMax

`func (o *Device) HasSpotPriceMax() bool`

HasSpotPriceMax returns a boolean if a field has been set.

### GetSshKeys

`func (o *Device) GetSshKeys() []Href`

GetSshKeys returns the SshKeys field if non-nil, zero value otherwise.

### GetSshKeysOk

`func (o *Device) GetSshKeysOk() (*[]Href, bool)`

GetSshKeysOk returns a tuple with the SshKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeys

`func (o *Device) SetSshKeys(v []Href)`

SetSshKeys sets SshKeys field to given value.

### HasSshKeys

`func (o *Device) HasSshKeys() bool`

HasSshKeys returns a boolean if a field has been set.

### GetState

`func (o *Device) GetState() DeviceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Device) GetStateOk() (*DeviceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Device) SetState(v DeviceState)`

SetState sets State field to given value.

### HasState

`func (o *Device) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorage

`func (o *Device) GetStorage() Storage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Device) GetStorageOk() (*Storage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Device) SetStorage(v Storage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Device) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSwitchUuid

`func (o *Device) GetSwitchUuid() string`

GetSwitchUuid returns the SwitchUuid field if non-nil, zero value otherwise.

### GetSwitchUuidOk

`func (o *Device) GetSwitchUuidOk() (*string, bool)`

GetSwitchUuidOk returns a tuple with the SwitchUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchUuid

`func (o *Device) SetSwitchUuid(v string)`

SetSwitchUuid sets SwitchUuid field to given value.

### HasSwitchUuid

`func (o *Device) HasSwitchUuid() bool`

HasSwitchUuid returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTerminationTime

`func (o *Device) GetTerminationTime() time.Time`

GetTerminationTime returns the TerminationTime field if non-nil, zero value otherwise.

### GetTerminationTimeOk

`func (o *Device) GetTerminationTimeOk() (*time.Time, bool)`

GetTerminationTimeOk returns a tuple with the TerminationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminationTime

`func (o *Device) SetTerminationTime(v time.Time)`

SetTerminationTime sets TerminationTime field to given value.

### HasTerminationTime

`func (o *Device) HasTerminationTime() bool`

HasTerminationTime returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Device) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Device) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Device) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Device) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUser

`func (o *Device) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Device) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Device) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Device) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserdata

`func (o *Device) GetUserdata() string`

GetUserdata returns the Userdata field if non-nil, zero value otherwise.

### GetUserdataOk

`func (o *Device) GetUserdataOk() (*string, bool)`

GetUserdataOk returns a tuple with the Userdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdata

`func (o *Device) SetUserdata(v string)`

SetUserdata sets Userdata field to given value.

### HasUserdata

`func (o *Device) HasUserdata() bool`

HasUserdata returns a boolean if a field has been set.

### GetVolumes

`func (o *Device) GetVolumes() []Href`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *Device) GetVolumesOk() (*[]Href, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *Device) SetVolumes(v []Href)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *Device) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetSos

`func (o *Device) GetSos() string`

GetSos returns the Sos field if non-nil, zero value otherwise.

### GetSosOk

`func (o *Device) GetSosOk() (*string, bool)`

GetSosOk returns a tuple with the Sos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSos

`func (o *Device) SetSos(v string)`

SetSos sets Sos field to given value.

### HasSos

`func (o *Device) HasSos() bool`

HasSos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


