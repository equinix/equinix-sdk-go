# Port

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | [**SimplifiedAccount**](SimplifiedAccount.md) |  | 
**Type** | [**PortType**](PortType.md) |  | 
**Id** | Pointer to **int32** | Equinix assigned response attribute for Port Id | [optional] 
**Href** | Pointer to **string** | Equinix assigned response attribute for an absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix assigned response attribute for  port identifier | [optional] 
**Name** | Pointer to **string** | Equinix assigned response attribute for Port name | [optional] 
**Description** | Pointer to **string** | Equinix assigned response attribute for Port description | [optional] 
**PhysicalPortsSpeed** | **int32** | Physical Ports Speed in Mbps | 
**ConnectionsCount** | Pointer to **int32** | Equinix assigned response attribute for Connection count | [optional] 
**PhysicalPortsType** | [**PortResponsePhysicalPortsType**](PortResponsePhysicalPortsType.md) |  | 
**PhysicalPortsCount** | Pointer to **int32** |  | [optional] 
**ConnectivitySourceType** | [**PortResponseConnectivitySourceType**](PortResponseConnectivitySourceType.md) |  | 
**BmmrType** | Pointer to [**PortResponseBmmrType**](PortResponseBmmrType.md) |  | [optional] 
**Project** | Pointer to [**PortResponseProject**](PortResponseProject.md) |  | [optional] 
**State** | Pointer to [**PortState**](PortState.md) |  | [optional] 
**Order** | Pointer to [**PortOrder**](PortOrder.md) |  | [optional] 
**CvpId** | Pointer to **string** | Equinix assigned response attribute for Unique ID for a virtual port. | [optional] 
**Operation** | Pointer to [**PortOperation**](PortOperation.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**ServiceType** | Pointer to [**PortResponseServiceType**](PortResponseServiceType.md) |  | [optional] 
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
**Settings** | [**PortSettings**](PortSettings.md) |  | 
**PhysicalPortQuantity** | Pointer to **int32** | Number of physical ports | [optional] 
**Notifications** | Pointer to [**[]PortNotification**](PortNotification.md) | Notification preferences | [optional] 
**AdditionalInfo** | Pointer to [**[]PortAdditionalInfo**](PortAdditionalInfo.md) | Port additional information | [optional] 
**PhysicalPorts** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | Physical ports that implement this port | [optional] 
**Loas** | Pointer to [**[]PortLoa**](PortLoa.md) | Port Loas | [optional] 

## Methods

### NewPort

`func NewPort(account SimplifiedAccount, type_ PortType, physicalPortsSpeed int32, physicalPortsType PortResponsePhysicalPortsType, connectivitySourceType PortResponseConnectivitySourceType, location SimplifiedLocation, encapsulation PortEncapsulation, settings PortSettings, ) *Port`

NewPort instantiates a new Port object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortWithDefaults

`func NewPortWithDefaults() *Port`

NewPortWithDefaults instantiates a new Port object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *Port) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Port) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Port) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.


### GetType

`func (o *Port) GetType() PortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Port) GetTypeOk() (*PortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Port) SetType(v PortType)`

SetType sets Type field to given value.


### GetId

`func (o *Port) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Port) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Port) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Port) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *Port) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Port) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Port) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Port) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Port) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Port) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Port) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Port) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Port) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Port) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Port) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Port) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Port) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Port) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Port) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Port) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalPortsSpeed

`func (o *Port) GetPhysicalPortsSpeed() int32`

GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field if non-nil, zero value otherwise.

### GetPhysicalPortsSpeedOk

`func (o *Port) GetPhysicalPortsSpeedOk() (*int32, bool)`

GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsSpeed

`func (o *Port) SetPhysicalPortsSpeed(v int32)`

SetPhysicalPortsSpeed sets PhysicalPortsSpeed field to given value.


### GetConnectionsCount

`func (o *Port) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *Port) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *Port) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *Port) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetPhysicalPortsType

`func (o *Port) GetPhysicalPortsType() PortResponsePhysicalPortsType`

GetPhysicalPortsType returns the PhysicalPortsType field if non-nil, zero value otherwise.

### GetPhysicalPortsTypeOk

`func (o *Port) GetPhysicalPortsTypeOk() (*PortResponsePhysicalPortsType, bool)`

GetPhysicalPortsTypeOk returns a tuple with the PhysicalPortsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsType

`func (o *Port) SetPhysicalPortsType(v PortResponsePhysicalPortsType)`

SetPhysicalPortsType sets PhysicalPortsType field to given value.


### GetPhysicalPortsCount

`func (o *Port) GetPhysicalPortsCount() int32`

GetPhysicalPortsCount returns the PhysicalPortsCount field if non-nil, zero value otherwise.

### GetPhysicalPortsCountOk

`func (o *Port) GetPhysicalPortsCountOk() (*int32, bool)`

GetPhysicalPortsCountOk returns a tuple with the PhysicalPortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsCount

`func (o *Port) SetPhysicalPortsCount(v int32)`

SetPhysicalPortsCount sets PhysicalPortsCount field to given value.

### HasPhysicalPortsCount

`func (o *Port) HasPhysicalPortsCount() bool`

HasPhysicalPortsCount returns a boolean if a field has been set.

### GetConnectivitySourceType

`func (o *Port) GetConnectivitySourceType() PortResponseConnectivitySourceType`

GetConnectivitySourceType returns the ConnectivitySourceType field if non-nil, zero value otherwise.

### GetConnectivitySourceTypeOk

`func (o *Port) GetConnectivitySourceTypeOk() (*PortResponseConnectivitySourceType, bool)`

GetConnectivitySourceTypeOk returns a tuple with the ConnectivitySourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivitySourceType

`func (o *Port) SetConnectivitySourceType(v PortResponseConnectivitySourceType)`

SetConnectivitySourceType sets ConnectivitySourceType field to given value.


### GetBmmrType

`func (o *Port) GetBmmrType() PortResponseBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *Port) GetBmmrTypeOk() (*PortResponseBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *Port) SetBmmrType(v PortResponseBmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *Port) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetProject

`func (o *Port) GetProject() PortResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *Port) GetProjectOk() (*PortResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *Port) SetProject(v PortResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *Port) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *Port) GetState() PortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Port) GetStateOk() (*PortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Port) SetState(v PortState)`

SetState sets State field to given value.

### HasState

`func (o *Port) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOrder

`func (o *Port) GetOrder() PortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *Port) GetOrderOk() (*PortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *Port) SetOrder(v PortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *Port) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCvpId

`func (o *Port) GetCvpId() string`

GetCvpId returns the CvpId field if non-nil, zero value otherwise.

### GetCvpIdOk

`func (o *Port) GetCvpIdOk() (*string, bool)`

GetCvpIdOk returns a tuple with the CvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpId

`func (o *Port) SetCvpId(v string)`

SetCvpId sets CvpId field to given value.

### HasCvpId

`func (o *Port) HasCvpId() bool`

HasCvpId returns a boolean if a field has been set.

### GetOperation

`func (o *Port) GetOperation() PortOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Port) GetOperationOk() (*PortOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Port) SetOperation(v PortOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Port) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetChangelog

`func (o *Port) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *Port) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *Port) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *Port) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetServiceType

`func (o *Port) GetServiceType() PortResponseServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Port) GetServiceTypeOk() (*PortResponseServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Port) SetServiceType(v PortResponseServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *Port) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetBandwidth

`func (o *Port) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *Port) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *Port) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *Port) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAvailableBandwidth

`func (o *Port) GetAvailableBandwidth() int32`

GetAvailableBandwidth returns the AvailableBandwidth field if non-nil, zero value otherwise.

### GetAvailableBandwidthOk

`func (o *Port) GetAvailableBandwidthOk() (*int32, bool)`

GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBandwidth

`func (o *Port) SetAvailableBandwidth(v int32)`

SetAvailableBandwidth sets AvailableBandwidth field to given value.

### HasAvailableBandwidth

`func (o *Port) HasAvailableBandwidth() bool`

HasAvailableBandwidth returns a boolean if a field has been set.

### GetUsedBandwidth

`func (o *Port) GetUsedBandwidth() int32`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *Port) GetUsedBandwidthOk() (*int32, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *Port) SetUsedBandwidth(v int32)`

SetUsedBandwidth sets UsedBandwidth field to given value.

### HasUsedBandwidth

`func (o *Port) HasUsedBandwidth() bool`

HasUsedBandwidth returns a boolean if a field has been set.

### GetLocation

`func (o *Port) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Port) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Port) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.


### GetDevice

`func (o *Port) GetDevice() PortDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Port) GetDeviceOk() (*PortDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Port) SetDevice(v PortDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *Port) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterface

`func (o *Port) GetInterface() PortInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Port) GetInterfaceOk() (*PortInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Port) SetInterface(v PortInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *Port) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetDemarcationPointIbx

`func (o *Port) GetDemarcationPointIbx() string`

GetDemarcationPointIbx returns the DemarcationPointIbx field if non-nil, zero value otherwise.

### GetDemarcationPointIbxOk

`func (o *Port) GetDemarcationPointIbxOk() (*string, bool)`

GetDemarcationPointIbxOk returns a tuple with the DemarcationPointIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPointIbx

`func (o *Port) SetDemarcationPointIbx(v string)`

SetDemarcationPointIbx sets DemarcationPointIbx field to given value.

### HasDemarcationPointIbx

`func (o *Port) HasDemarcationPointIbx() bool`

HasDemarcationPointIbx returns a boolean if a field has been set.

### GetTetherIbx

`func (o *Port) GetTetherIbx() string`

GetTetherIbx returns the TetherIbx field if non-nil, zero value otherwise.

### GetTetherIbxOk

`func (o *Port) GetTetherIbxOk() (*string, bool)`

GetTetherIbxOk returns a tuple with the TetherIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTetherIbx

`func (o *Port) SetTetherIbx(v string)`

SetTetherIbx sets TetherIbx field to given value.

### HasTetherIbx

`func (o *Port) HasTetherIbx() bool`

HasTetherIbx returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *Port) GetDemarcationPoint() PortDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *Port) GetDemarcationPointOk() (*PortDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *Port) SetDemarcationPoint(v PortDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *Port) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetRedundancy

`func (o *Port) GetRedundancy() PortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *Port) GetRedundancyOk() (*PortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *Port) SetRedundancy(v PortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *Port) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEncapsulation

`func (o *Port) GetEncapsulation() PortEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *Port) GetEncapsulationOk() (*PortEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *Port) SetEncapsulation(v PortEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.


### GetLagEnabled

`func (o *Port) GetLagEnabled() bool`

GetLagEnabled returns the LagEnabled field if non-nil, zero value otherwise.

### GetLagEnabledOk

`func (o *Port) GetLagEnabledOk() (*bool, bool)`

GetLagEnabledOk returns a tuple with the LagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagEnabled

`func (o *Port) SetLagEnabled(v bool)`

SetLagEnabled sets LagEnabled field to given value.

### HasLagEnabled

`func (o *Port) HasLagEnabled() bool`

HasLagEnabled returns a boolean if a field has been set.

### GetLag

`func (o *Port) GetLag() PortLag`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *Port) GetLagOk() (*PortLag, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *Port) SetLag(v PortLag)`

SetLag sets Lag field to given value.

### HasLag

`func (o *Port) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetAsn

`func (o *Port) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *Port) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *Port) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *Port) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetSettings

`func (o *Port) GetSettings() PortSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Port) GetSettingsOk() (*PortSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Port) SetSettings(v PortSettings)`

SetSettings sets Settings field to given value.


### GetPhysicalPortQuantity

`func (o *Port) GetPhysicalPortQuantity() int32`

GetPhysicalPortQuantity returns the PhysicalPortQuantity field if non-nil, zero value otherwise.

### GetPhysicalPortQuantityOk

`func (o *Port) GetPhysicalPortQuantityOk() (*int32, bool)`

GetPhysicalPortQuantityOk returns a tuple with the PhysicalPortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortQuantity

`func (o *Port) SetPhysicalPortQuantity(v int32)`

SetPhysicalPortQuantity sets PhysicalPortQuantity field to given value.

### HasPhysicalPortQuantity

`func (o *Port) HasPhysicalPortQuantity() bool`

HasPhysicalPortQuantity returns a boolean if a field has been set.

### GetNotifications

`func (o *Port) GetNotifications() []PortNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Port) GetNotificationsOk() (*[]PortNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Port) SetNotifications(v []PortNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Port) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *Port) GetAdditionalInfo() []PortAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *Port) GetAdditionalInfoOk() (*[]PortAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *Port) SetAdditionalInfo(v []PortAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *Port) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPhysicalPorts

`func (o *Port) GetPhysicalPorts() []PhysicalPort`

GetPhysicalPorts returns the PhysicalPorts field if non-nil, zero value otherwise.

### GetPhysicalPortsOk

`func (o *Port) GetPhysicalPortsOk() (*[]PhysicalPort, bool)`

GetPhysicalPortsOk returns a tuple with the PhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPorts

`func (o *Port) SetPhysicalPorts(v []PhysicalPort)`

SetPhysicalPorts sets PhysicalPorts field to given value.

### HasPhysicalPorts

`func (o *Port) HasPhysicalPorts() bool`

HasPhysicalPorts returns a boolean if a field has been set.

### GetLoas

`func (o *Port) GetLoas() []PortLoa`

GetLoas returns the Loas field if non-nil, zero value otherwise.

### GetLoasOk

`func (o *Port) GetLoasOk() (*[]PortLoa, bool)`

GetLoasOk returns a tuple with the Loas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoas

`func (o *Port) SetLoas(v []PortLoa)`

SetLoas sets Loas field to given value.

### HasLoas

`func (o *Port) HasLoas() bool`

HasLoas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


