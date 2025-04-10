# PortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Equinix assigned response attribute for an absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Type** | [**PortType**](PortType.md) |  | 
**Id** | Pointer to **int32** | Equinix assigned response attribute for Port Id | [optional] 
**Uuid** | Pointer to **string** | Equinix assigned response attribute for  port identifier | [optional] 
**Name** | Pointer to **string** | Equinix assigned response attribute for Port name | [optional] 
**Description** | Pointer to **string** | Equinix assigned response attribute for Port description | [optional] 
**PhysicalPortsSpeed** | **int32** | Physical Ports Speed in Mbps | 
**ConnectionsCount** | Pointer to **int32** | Equinix assigned response attribute for Connection count | [optional] 
**PhysicalPortsType** | [**PortPhysicalPortsType**](PortPhysicalPortsType.md) |  | 
**PhysicalPortsCount** | Pointer to **int32** |  | [optional] 
**ConnectivitySourceType** | [**PortConnectivitySourceType**](PortConnectivitySourceType.md) |  | 
**BmmrType** | Pointer to [**PortBmmrType**](PortBmmrType.md) |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to [**PortState**](PortState.md) |  | [optional] 
**Order** | Pointer to [**PortOrder**](PortOrder.md) |  | [optional] 
**Operation** | Pointer to [**PortOperation**](PortOperation.md) |  | [optional] 
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Change** | Pointer to [**PortChange**](PortChange.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**ServiceType** | Pointer to [**PortServiceType**](PortServiceType.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port bandwidth in Mbps | [optional] 
**AvailableBandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port available bandwidth in Mbps | [optional] 
**UsedBandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port used bandwidth in Mbps | [optional] 
**Location** | [**SimplifiedLocation**](SimplifiedLocation.md) |  | 
**Device** | Pointer to [**PortDevice**](PortDevice.md) |  | [optional] 
**Interface** | Pointer to [**PortInterface**](PortInterface.md) |  | [optional] 
**DemarcationPointIbx** | Pointer to **string** | A-side/Equinix ibx | [optional] 
**TetherIbx** | Pointer to **string** | z-side/Equinix ibx | [optional] 
**DemarcationPoint** | Pointer to [**PortDemarcationPoint**](PortDemarcationPoint.md) |  | [optional] 
**Redundancy** | Pointer to [**PortRedundancy**](PortRedundancy.md) |  | [optional] 
**Encapsulation** | [**PortEncapsulation**](PortEncapsulation.md) |  | 
**LagEnabled** | Pointer to **bool** | If LAG enabled | [optional] 
**Lag** | Pointer to [**PortLag**](PortLag.md) |  | [optional] 
**Asn** | Pointer to **int32** | Port ASN | [optional] 
**Package** | Pointer to [**Package**](Package.md) |  | [optional] 
**Settings** | [**PortSettings**](PortSettings.md) |  | 
**PhysicalPortQuantity** | Pointer to **int32** | Number of physical ports | [optional] 
**Notifications** | Pointer to [**[]PortNotification**](PortNotification.md) | Notification preferences | [optional] 
**AdditionalInfo** | Pointer to [**[]PortAdditionalInfo**](PortAdditionalInfo.md) | Port additional information | [optional] 
**EndCustomer** | Pointer to [**EndCustomer**](EndCustomer.md) |  | [optional] 
**PhysicalPorts** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | Physical ports that implement this port | [optional] 
**Loas** | Pointer to [**[]PortLoa**](PortLoa.md) | Port Loas | [optional] 

## Methods

### NewPortRequest

`func NewPortRequest(type_ PortType, physicalPortsSpeed int32, physicalPortsType PortPhysicalPortsType, connectivitySourceType PortConnectivitySourceType, account SimplifiedAccount, location SimplifiedLocation, encapsulation PortEncapsulation, settings PortSettings, ) *PortRequest`

NewPortRequest instantiates a new PortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortRequestWithDefaults

`func NewPortRequestWithDefaults() *PortRequest`

NewPortRequestWithDefaults instantiates a new PortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PortRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PortRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PortRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PortRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *PortRequest) GetType() PortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortRequest) GetTypeOk() (*PortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortRequest) SetType(v PortType)`

SetType sets Type field to given value.


### GetId

`func (o *PortRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *PortRequest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PortRequest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PortRequest) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PortRequest) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PortRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PortRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalPortsSpeed

`func (o *PortRequest) GetPhysicalPortsSpeed() int32`

GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field if non-nil, zero value otherwise.

### GetPhysicalPortsSpeedOk

`func (o *PortRequest) GetPhysicalPortsSpeedOk() (*int32, bool)`

GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsSpeed

`func (o *PortRequest) SetPhysicalPortsSpeed(v int32)`

SetPhysicalPortsSpeed sets PhysicalPortsSpeed field to given value.


### GetConnectionsCount

`func (o *PortRequest) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *PortRequest) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *PortRequest) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *PortRequest) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetPhysicalPortsType

`func (o *PortRequest) GetPhysicalPortsType() PortPhysicalPortsType`

GetPhysicalPortsType returns the PhysicalPortsType field if non-nil, zero value otherwise.

### GetPhysicalPortsTypeOk

`func (o *PortRequest) GetPhysicalPortsTypeOk() (*PortPhysicalPortsType, bool)`

GetPhysicalPortsTypeOk returns a tuple with the PhysicalPortsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsType

`func (o *PortRequest) SetPhysicalPortsType(v PortPhysicalPortsType)`

SetPhysicalPortsType sets PhysicalPortsType field to given value.


### GetPhysicalPortsCount

`func (o *PortRequest) GetPhysicalPortsCount() int32`

GetPhysicalPortsCount returns the PhysicalPortsCount field if non-nil, zero value otherwise.

### GetPhysicalPortsCountOk

`func (o *PortRequest) GetPhysicalPortsCountOk() (*int32, bool)`

GetPhysicalPortsCountOk returns a tuple with the PhysicalPortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsCount

`func (o *PortRequest) SetPhysicalPortsCount(v int32)`

SetPhysicalPortsCount sets PhysicalPortsCount field to given value.

### HasPhysicalPortsCount

`func (o *PortRequest) HasPhysicalPortsCount() bool`

HasPhysicalPortsCount returns a boolean if a field has been set.

### GetConnectivitySourceType

`func (o *PortRequest) GetConnectivitySourceType() PortConnectivitySourceType`

GetConnectivitySourceType returns the ConnectivitySourceType field if non-nil, zero value otherwise.

### GetConnectivitySourceTypeOk

`func (o *PortRequest) GetConnectivitySourceTypeOk() (*PortConnectivitySourceType, bool)`

GetConnectivitySourceTypeOk returns a tuple with the ConnectivitySourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivitySourceType

`func (o *PortRequest) SetConnectivitySourceType(v PortConnectivitySourceType)`

SetConnectivitySourceType sets ConnectivitySourceType field to given value.


### GetBmmrType

`func (o *PortRequest) GetBmmrType() PortBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *PortRequest) GetBmmrTypeOk() (*PortBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *PortRequest) SetBmmrType(v PortBmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *PortRequest) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetProject

`func (o *PortRequest) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PortRequest) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PortRequest) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *PortRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *PortRequest) GetState() PortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortRequest) GetStateOk() (*PortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortRequest) SetState(v PortState)`

SetState sets State field to given value.

### HasState

`func (o *PortRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOrder

`func (o *PortRequest) GetOrder() PortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PortRequest) GetOrderOk() (*PortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PortRequest) SetOrder(v PortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PortRequest) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOperation

`func (o *PortRequest) GetOperation() PortOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PortRequest) GetOperationOk() (*PortOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PortRequest) SetOperation(v PortOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PortRequest) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAccount

`func (o *PortRequest) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PortRequest) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PortRequest) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetChange

`func (o *PortRequest) GetChange() PortChange`

GetChange returns the Change field if non-nil, zero value otherwise.

### GetChangeOk

`func (o *PortRequest) GetChangeOk() (*PortChange, bool)`

GetChangeOk returns a tuple with the Change field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChange

`func (o *PortRequest) SetChange(v PortChange)`

SetChange sets Change field to given value.

### HasChange

`func (o *PortRequest) HasChange() bool`

HasChange returns a boolean if a field has been set.

### GetChangeLog

`func (o *PortRequest) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *PortRequest) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *PortRequest) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *PortRequest) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.

### GetServiceType

`func (o *PortRequest) GetServiceType() PortServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PortRequest) GetServiceTypeOk() (*PortServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PortRequest) SetServiceType(v PortServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PortRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetBandwidth

`func (o *PortRequest) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *PortRequest) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *PortRequest) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *PortRequest) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAvailableBandwidth

`func (o *PortRequest) GetAvailableBandwidth() int32`

GetAvailableBandwidth returns the AvailableBandwidth field if non-nil, zero value otherwise.

### GetAvailableBandwidthOk

`func (o *PortRequest) GetAvailableBandwidthOk() (*int32, bool)`

GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBandwidth

`func (o *PortRequest) SetAvailableBandwidth(v int32)`

SetAvailableBandwidth sets AvailableBandwidth field to given value.

### HasAvailableBandwidth

`func (o *PortRequest) HasAvailableBandwidth() bool`

HasAvailableBandwidth returns a boolean if a field has been set.

### GetUsedBandwidth

`func (o *PortRequest) GetUsedBandwidth() int32`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *PortRequest) GetUsedBandwidthOk() (*int32, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *PortRequest) SetUsedBandwidth(v int32)`

SetUsedBandwidth sets UsedBandwidth field to given value.

### HasUsedBandwidth

`func (o *PortRequest) HasUsedBandwidth() bool`

HasUsedBandwidth returns a boolean if a field has been set.

### GetLocation

`func (o *PortRequest) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PortRequest) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PortRequest) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.


### GetDevice

`func (o *PortRequest) GetDevice() PortDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PortRequest) GetDeviceOk() (*PortDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PortRequest) SetDevice(v PortDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PortRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterface

`func (o *PortRequest) GetInterface() PortInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PortRequest) GetInterfaceOk() (*PortInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PortRequest) SetInterface(v PortInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PortRequest) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetDemarcationPointIbx

`func (o *PortRequest) GetDemarcationPointIbx() string`

GetDemarcationPointIbx returns the DemarcationPointIbx field if non-nil, zero value otherwise.

### GetDemarcationPointIbxOk

`func (o *PortRequest) GetDemarcationPointIbxOk() (*string, bool)`

GetDemarcationPointIbxOk returns a tuple with the DemarcationPointIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPointIbx

`func (o *PortRequest) SetDemarcationPointIbx(v string)`

SetDemarcationPointIbx sets DemarcationPointIbx field to given value.

### HasDemarcationPointIbx

`func (o *PortRequest) HasDemarcationPointIbx() bool`

HasDemarcationPointIbx returns a boolean if a field has been set.

### GetTetherIbx

`func (o *PortRequest) GetTetherIbx() string`

GetTetherIbx returns the TetherIbx field if non-nil, zero value otherwise.

### GetTetherIbxOk

`func (o *PortRequest) GetTetherIbxOk() (*string, bool)`

GetTetherIbxOk returns a tuple with the TetherIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTetherIbx

`func (o *PortRequest) SetTetherIbx(v string)`

SetTetherIbx sets TetherIbx field to given value.

### HasTetherIbx

`func (o *PortRequest) HasTetherIbx() bool`

HasTetherIbx returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *PortRequest) GetDemarcationPoint() PortDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *PortRequest) GetDemarcationPointOk() (*PortDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *PortRequest) SetDemarcationPoint(v PortDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *PortRequest) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetRedundancy

`func (o *PortRequest) GetRedundancy() PortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *PortRequest) GetRedundancyOk() (*PortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *PortRequest) SetRedundancy(v PortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *PortRequest) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEncapsulation

`func (o *PortRequest) GetEncapsulation() PortEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PortRequest) GetEncapsulationOk() (*PortEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PortRequest) SetEncapsulation(v PortEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.


### GetLagEnabled

`func (o *PortRequest) GetLagEnabled() bool`

GetLagEnabled returns the LagEnabled field if non-nil, zero value otherwise.

### GetLagEnabledOk

`func (o *PortRequest) GetLagEnabledOk() (*bool, bool)`

GetLagEnabledOk returns a tuple with the LagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagEnabled

`func (o *PortRequest) SetLagEnabled(v bool)`

SetLagEnabled sets LagEnabled field to given value.

### HasLagEnabled

`func (o *PortRequest) HasLagEnabled() bool`

HasLagEnabled returns a boolean if a field has been set.

### GetLag

`func (o *PortRequest) GetLag() PortLag`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *PortRequest) GetLagOk() (*PortLag, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *PortRequest) SetLag(v PortLag)`

SetLag sets Lag field to given value.

### HasLag

`func (o *PortRequest) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetAsn

`func (o *PortRequest) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PortRequest) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PortRequest) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PortRequest) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetPackage

`func (o *PortRequest) GetPackage() Package`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *PortRequest) GetPackageOk() (*Package, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *PortRequest) SetPackage(v Package)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *PortRequest) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetSettings

`func (o *PortRequest) GetSettings() PortSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PortRequest) GetSettingsOk() (*PortSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PortRequest) SetSettings(v PortSettings)`

SetSettings sets Settings field to given value.


### GetPhysicalPortQuantity

`func (o *PortRequest) GetPhysicalPortQuantity() int32`

GetPhysicalPortQuantity returns the PhysicalPortQuantity field if non-nil, zero value otherwise.

### GetPhysicalPortQuantityOk

`func (o *PortRequest) GetPhysicalPortQuantityOk() (*int32, bool)`

GetPhysicalPortQuantityOk returns a tuple with the PhysicalPortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortQuantity

`func (o *PortRequest) SetPhysicalPortQuantity(v int32)`

SetPhysicalPortQuantity sets PhysicalPortQuantity field to given value.

### HasPhysicalPortQuantity

`func (o *PortRequest) HasPhysicalPortQuantity() bool`

HasPhysicalPortQuantity returns a boolean if a field has been set.

### GetNotifications

`func (o *PortRequest) GetNotifications() []PortNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PortRequest) GetNotificationsOk() (*[]PortNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PortRequest) SetNotifications(v []PortNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PortRequest) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PortRequest) GetAdditionalInfo() []PortAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PortRequest) GetAdditionalInfoOk() (*[]PortAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PortRequest) SetAdditionalInfo(v []PortAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PortRequest) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetEndCustomer

`func (o *PortRequest) GetEndCustomer() EndCustomer`

GetEndCustomer returns the EndCustomer field if non-nil, zero value otherwise.

### GetEndCustomerOk

`func (o *PortRequest) GetEndCustomerOk() (*EndCustomer, bool)`

GetEndCustomerOk returns a tuple with the EndCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCustomer

`func (o *PortRequest) SetEndCustomer(v EndCustomer)`

SetEndCustomer sets EndCustomer field to given value.

### HasEndCustomer

`func (o *PortRequest) HasEndCustomer() bool`

HasEndCustomer returns a boolean if a field has been set.

### GetPhysicalPorts

`func (o *PortRequest) GetPhysicalPorts() []PhysicalPort`

GetPhysicalPorts returns the PhysicalPorts field if non-nil, zero value otherwise.

### GetPhysicalPortsOk

`func (o *PortRequest) GetPhysicalPortsOk() (*[]PhysicalPort, bool)`

GetPhysicalPortsOk returns a tuple with the PhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPorts

`func (o *PortRequest) SetPhysicalPorts(v []PhysicalPort)`

SetPhysicalPorts sets PhysicalPorts field to given value.

### HasPhysicalPorts

`func (o *PortRequest) HasPhysicalPorts() bool`

HasPhysicalPorts returns a boolean if a field has been set.

### GetLoas

`func (o *PortRequest) GetLoas() []PortLoa`

GetLoas returns the Loas field if non-nil, zero value otherwise.

### GetLoasOk

`func (o *PortRequest) GetLoasOk() (*[]PortLoa, bool)`

GetLoasOk returns a tuple with the Loas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoas

`func (o *PortRequest) SetLoas(v []PortLoa)`

SetLoas sets Loas field to given value.

### HasLoas

`func (o *PortRequest) HasLoas() bool`

HasLoas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


