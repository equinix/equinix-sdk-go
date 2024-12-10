# SimplifiedPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Equinix assigned response attribute for an absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Type** | Pointer to [**PortType**](PortType.md) |  | [optional] 
**Id** | Pointer to **int32** | Equinix assigned response attribute for Port Id | [optional] 
**Uuid** | Pointer to **string** | Equinix assigned response attribute for  port identifier | [optional] 
**Name** | Pointer to **string** | Equinix assigned response attribute for Port name | [optional] 
**Description** | Pointer to **string** | Equinix assigned response attribute for Port description | [optional] 
**PhysicalPortsSpeed** | Pointer to **int32** | Physical Ports Speed in Mbps | [optional] 
**ConnectionsCount** | Pointer to **int32** | Equinix assigned response attribute for Connection count | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to [**PortState**](PortState.md) |  | [optional] 
**Operation** | Pointer to [**PortOperation**](PortOperation.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 
**ServiceType** | Pointer to [**PortServiceType**](PortServiceType.md) |  | [optional] 
**Bandwidth** | Pointer to **int64** | Equinix assigned response attribute for Port bandwidth in Mbps | [optional] 
**AvailableBandwidth** | Pointer to **int64** | Equinix assigned response attribute for Port available bandwidth in Mbps | [optional] 
**UsedBandwidth** | Pointer to **int64** | Equinix assigned response attribute for Port used bandwidth in Mbps | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Device** | Pointer to [**PortDevice**](PortDevice.md) |  | [optional] 
**Interface** | Pointer to [**PortInterface**](PortInterface.md) |  | [optional] 
**Tether** | Pointer to [**PortTether**](PortTether.md) |  | [optional] 
**DemarcationPoint** | Pointer to [**PortDemarcationPoint**](PortDemarcationPoint.md) |  | [optional] 
**Redundancy** | Pointer to [**PortRedundancy**](PortRedundancy.md) |  | [optional] 
**Encapsulation** | Pointer to [**PortEncapsulation**](PortEncapsulation.md) |  | [optional] 
**LagEnabled** | Pointer to **bool** | If LAG enabled | [optional] 
**Package** | Pointer to [**Package**](Package.md) |  | [optional] 
**Settings** | Pointer to [**PortSettings**](PortSettings.md) |  | [optional] 
**PhysicalPortQuantity** | Pointer to **int32** | Number of physical ports | [optional] 
**AdditionalInfo** | Pointer to [**[]PortAdditionalInfo**](PortAdditionalInfo.md) | Port additional information | [optional] 
**PhysicalPorts** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | Physical ports that implement this port | [optional] 

## Methods

### NewSimplifiedPort

`func NewSimplifiedPort() *SimplifiedPort`

NewSimplifiedPort instantiates a new SimplifiedPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedPortWithDefaults

`func NewSimplifiedPortWithDefaults() *SimplifiedPort`

NewSimplifiedPortWithDefaults instantiates a new SimplifiedPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedPort) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedPort) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedPort) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedPort) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedPort) GetType() PortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedPort) GetTypeOk() (*PortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedPort) SetType(v PortType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedPort) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SimplifiedPort) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimplifiedPort) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimplifiedPort) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SimplifiedPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedPort) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedPort) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedPort) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedPort) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SimplifiedPort) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SimplifiedPort) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SimplifiedPort) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SimplifiedPort) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalPortsSpeed

`func (o *SimplifiedPort) GetPhysicalPortsSpeed() int32`

GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field if non-nil, zero value otherwise.

### GetPhysicalPortsSpeedOk

`func (o *SimplifiedPort) GetPhysicalPortsSpeedOk() (*int32, bool)`

GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsSpeed

`func (o *SimplifiedPort) SetPhysicalPortsSpeed(v int32)`

SetPhysicalPortsSpeed sets PhysicalPortsSpeed field to given value.

### HasPhysicalPortsSpeed

`func (o *SimplifiedPort) HasPhysicalPortsSpeed() bool`

HasPhysicalPortsSpeed returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *SimplifiedPort) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *SimplifiedPort) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *SimplifiedPort) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *SimplifiedPort) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetProject

`func (o *SimplifiedPort) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *SimplifiedPort) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *SimplifiedPort) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *SimplifiedPort) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *SimplifiedPort) GetState() PortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SimplifiedPort) GetStateOk() (*PortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SimplifiedPort) SetState(v PortState)`

SetState sets State field to given value.

### HasState

`func (o *SimplifiedPort) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperation

`func (o *SimplifiedPort) GetOperation() PortOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *SimplifiedPort) GetOperationOk() (*PortOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *SimplifiedPort) SetOperation(v PortOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *SimplifiedPort) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAccount

`func (o *SimplifiedPort) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SimplifiedPort) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SimplifiedPort) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SimplifiedPort) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetServiceType

`func (o *SimplifiedPort) GetServiceType() PortServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *SimplifiedPort) GetServiceTypeOk() (*PortServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *SimplifiedPort) SetServiceType(v PortServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *SimplifiedPort) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetBandwidth

`func (o *SimplifiedPort) GetBandwidth() int64`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *SimplifiedPort) GetBandwidthOk() (*int64, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *SimplifiedPort) SetBandwidth(v int64)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *SimplifiedPort) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAvailableBandwidth

`func (o *SimplifiedPort) GetAvailableBandwidth() int64`

GetAvailableBandwidth returns the AvailableBandwidth field if non-nil, zero value otherwise.

### GetAvailableBandwidthOk

`func (o *SimplifiedPort) GetAvailableBandwidthOk() (*int64, bool)`

GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBandwidth

`func (o *SimplifiedPort) SetAvailableBandwidth(v int64)`

SetAvailableBandwidth sets AvailableBandwidth field to given value.

### HasAvailableBandwidth

`func (o *SimplifiedPort) HasAvailableBandwidth() bool`

HasAvailableBandwidth returns a boolean if a field has been set.

### GetUsedBandwidth

`func (o *SimplifiedPort) GetUsedBandwidth() int64`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *SimplifiedPort) GetUsedBandwidthOk() (*int64, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *SimplifiedPort) SetUsedBandwidth(v int64)`

SetUsedBandwidth sets UsedBandwidth field to given value.

### HasUsedBandwidth

`func (o *SimplifiedPort) HasUsedBandwidth() bool`

HasUsedBandwidth returns a boolean if a field has been set.

### GetLocation

`func (o *SimplifiedPort) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SimplifiedPort) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SimplifiedPort) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SimplifiedPort) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDevice

`func (o *SimplifiedPort) GetDevice() PortDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *SimplifiedPort) GetDeviceOk() (*PortDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *SimplifiedPort) SetDevice(v PortDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *SimplifiedPort) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterface

`func (o *SimplifiedPort) GetInterface() PortInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *SimplifiedPort) GetInterfaceOk() (*PortInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *SimplifiedPort) SetInterface(v PortInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *SimplifiedPort) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetTether

`func (o *SimplifiedPort) GetTether() PortTether`

GetTether returns the Tether field if non-nil, zero value otherwise.

### GetTetherOk

`func (o *SimplifiedPort) GetTetherOk() (*PortTether, bool)`

GetTetherOk returns a tuple with the Tether field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTether

`func (o *SimplifiedPort) SetTether(v PortTether)`

SetTether sets Tether field to given value.

### HasTether

`func (o *SimplifiedPort) HasTether() bool`

HasTether returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *SimplifiedPort) GetDemarcationPoint() PortDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *SimplifiedPort) GetDemarcationPointOk() (*PortDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *SimplifiedPort) SetDemarcationPoint(v PortDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *SimplifiedPort) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetRedundancy

`func (o *SimplifiedPort) GetRedundancy() PortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *SimplifiedPort) GetRedundancyOk() (*PortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *SimplifiedPort) SetRedundancy(v PortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *SimplifiedPort) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEncapsulation

`func (o *SimplifiedPort) GetEncapsulation() PortEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *SimplifiedPort) GetEncapsulationOk() (*PortEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *SimplifiedPort) SetEncapsulation(v PortEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *SimplifiedPort) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetLagEnabled

`func (o *SimplifiedPort) GetLagEnabled() bool`

GetLagEnabled returns the LagEnabled field if non-nil, zero value otherwise.

### GetLagEnabledOk

`func (o *SimplifiedPort) GetLagEnabledOk() (*bool, bool)`

GetLagEnabledOk returns a tuple with the LagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagEnabled

`func (o *SimplifiedPort) SetLagEnabled(v bool)`

SetLagEnabled sets LagEnabled field to given value.

### HasLagEnabled

`func (o *SimplifiedPort) HasLagEnabled() bool`

HasLagEnabled returns a boolean if a field has been set.

### GetPackage

`func (o *SimplifiedPort) GetPackage() Package`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SimplifiedPort) GetPackageOk() (*Package, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SimplifiedPort) SetPackage(v Package)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SimplifiedPort) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSettings

`func (o *SimplifiedPort) GetSettings() PortSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SimplifiedPort) GetSettingsOk() (*PortSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SimplifiedPort) SetSettings(v PortSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SimplifiedPort) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPhysicalPortQuantity

`func (o *SimplifiedPort) GetPhysicalPortQuantity() int32`

GetPhysicalPortQuantity returns the PhysicalPortQuantity field if non-nil, zero value otherwise.

### GetPhysicalPortQuantityOk

`func (o *SimplifiedPort) GetPhysicalPortQuantityOk() (*int32, bool)`

GetPhysicalPortQuantityOk returns a tuple with the PhysicalPortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortQuantity

`func (o *SimplifiedPort) SetPhysicalPortQuantity(v int32)`

SetPhysicalPortQuantity sets PhysicalPortQuantity field to given value.

### HasPhysicalPortQuantity

`func (o *SimplifiedPort) HasPhysicalPortQuantity() bool`

HasPhysicalPortQuantity returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *SimplifiedPort) GetAdditionalInfo() []PortAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *SimplifiedPort) GetAdditionalInfoOk() (*[]PortAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *SimplifiedPort) SetAdditionalInfo(v []PortAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *SimplifiedPort) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPhysicalPorts

`func (o *SimplifiedPort) GetPhysicalPorts() []PhysicalPort`

GetPhysicalPorts returns the PhysicalPorts field if non-nil, zero value otherwise.

### GetPhysicalPortsOk

`func (o *SimplifiedPort) GetPhysicalPortsOk() (*[]PhysicalPort, bool)`

GetPhysicalPortsOk returns a tuple with the PhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPorts

`func (o *SimplifiedPort) SetPhysicalPorts(v []PhysicalPort)`

SetPhysicalPorts sets PhysicalPorts field to given value.

### HasPhysicalPorts

`func (o *SimplifiedPort) HasPhysicalPorts() bool`

HasPhysicalPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


