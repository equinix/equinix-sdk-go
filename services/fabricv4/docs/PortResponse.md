# PortResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PortType**](PortType.md) |  | [optional] 
**Id** | Pointer to **int32** | Equinix assigned response attribute for Port Id | [optional] 
**Href** | Pointer to **string** | Equinix assigned response attribute for an absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix assigned response attribute for  port identifier | [optional] 
**Name** | Pointer to **string** | Equinix assigned response attribute for Port name | [optional] 
**Description** | Pointer to **string** | Equinix assigned response attribute for Port description | [optional] 
**PhysicalPortsSpeed** | Pointer to **int32** | Physical Ports Speed in Mbps | [optional] 
**ConnectionsCount** | Pointer to **int32** | Equinix assigned response attribute for Connection count | [optional] 
**PhysicalPortsType** | Pointer to [**PortResponsePhysicalPortsType**](PortResponsePhysicalPortsType.md) |  | [optional] 
**PhysicalPortsCount** | Pointer to **int32** |  | [optional] 
**ConnectivitySourceType** | Pointer to [**PortResponseConnectivitySourceType**](PortResponseConnectivitySourceType.md) |  | [optional] 
**BmmrType** | Pointer to [**PortResponseBmmrType**](PortResponseBmmrType.md) |  | [optional] 
**Project** | Pointer to [**PortResponseProject**](PortResponseProject.md) |  | [optional] 
**State** | Pointer to [**PortState**](PortState.md) |  | [optional] 
**Order** | Pointer to [**PortOrder**](PortOrder.md) |  | [optional] 
**CvpId** | Pointer to **string** | Equinix assigned response attribute for Unique ID for a virtual port. | [optional] 
**Operation** | Pointer to [**PortOperation**](PortOperation.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccountPortResponse**](SimplifiedAccountPortResponse.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 
**ServiceType** | Pointer to [**PortResponseServiceType**](PortResponseServiceType.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port bandwidth in Mbps | [optional] 
**AvailableBandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port available bandwidth in Mbps | [optional] 
**UsedBandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port used bandwidth in Mbps | [optional] 
**Location** | Pointer to [**SimplifiedLocation**](SimplifiedLocation.md) |  | [optional] 
**Device** | Pointer to [**PortDevice**](PortDevice.md) |  | [optional] 
**Interface** | Pointer to [**PortInterface**](PortInterface.md) |  | [optional] 
**DemarcationPointIbx** | Pointer to **string** | A-side/Equinix ibx | [optional] 
**TetherIbx** | Pointer to **string** | z-side/Equinix ibx | [optional] 
**DemarcationPoint** | Pointer to [**PortDemarcationPoint**](PortDemarcationPoint.md) |  | [optional] 
**Redundancy** | Pointer to [**PortRedundancy**](PortRedundancy.md) |  | [optional] 
**Encapsulation** | Pointer to [**PortEncapsulation**](PortEncapsulation.md) |  | [optional] 
**LagEnabled** | Pointer to **bool** | If LAG enabled | [optional] 
**Lag** | Pointer to [**PortLag**](PortLag.md) |  | [optional] 
**Asn** | Pointer to **int32** | Port ASN | [optional] 
**Settings** | Pointer to [**PortSettings**](PortSettings.md) |  | [optional] 
**PhysicalPortQuantity** | Pointer to **int32** | Number of physical ports | [optional] 
**Notifications** | Pointer to [**[]PortNotification**](PortNotification.md) | Notification preferences | [optional] 
**AdditionalInfo** | Pointer to [**[]PortAdditionalInfo**](PortAdditionalInfo.md) | Port additional information | [optional] 
**PhysicalPorts** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | Physical ports that implement this port | [optional] 
**Loas** | Pointer to [**[]PortLoa**](PortLoa.md) | Port Loas | [optional] 

## Methods

### NewPortResponse

`func NewPortResponse() *PortResponse`

NewPortResponse instantiates a new PortResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortResponseWithDefaults

`func NewPortResponseWithDefaults() *PortResponse`

NewPortResponseWithDefaults instantiates a new PortResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PortResponse) GetType() PortType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortResponse) GetTypeOk() (*PortType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortResponse) SetType(v PortType)`

SetType sets Type field to given value.

### HasType

`func (o *PortResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *PortResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *PortResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PortResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PortResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PortResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *PortResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PortResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PortResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PortResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PortResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PortResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPhysicalPortsSpeed

`func (o *PortResponse) GetPhysicalPortsSpeed() int32`

GetPhysicalPortsSpeed returns the PhysicalPortsSpeed field if non-nil, zero value otherwise.

### GetPhysicalPortsSpeedOk

`func (o *PortResponse) GetPhysicalPortsSpeedOk() (*int32, bool)`

GetPhysicalPortsSpeedOk returns a tuple with the PhysicalPortsSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsSpeed

`func (o *PortResponse) SetPhysicalPortsSpeed(v int32)`

SetPhysicalPortsSpeed sets PhysicalPortsSpeed field to given value.

### HasPhysicalPortsSpeed

`func (o *PortResponse) HasPhysicalPortsSpeed() bool`

HasPhysicalPortsSpeed returns a boolean if a field has been set.

### GetConnectionsCount

`func (o *PortResponse) GetConnectionsCount() int32`

GetConnectionsCount returns the ConnectionsCount field if non-nil, zero value otherwise.

### GetConnectionsCountOk

`func (o *PortResponse) GetConnectionsCountOk() (*int32, bool)`

GetConnectionsCountOk returns a tuple with the ConnectionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionsCount

`func (o *PortResponse) SetConnectionsCount(v int32)`

SetConnectionsCount sets ConnectionsCount field to given value.

### HasConnectionsCount

`func (o *PortResponse) HasConnectionsCount() bool`

HasConnectionsCount returns a boolean if a field has been set.

### GetPhysicalPortsType

`func (o *PortResponse) GetPhysicalPortsType() PortResponsePhysicalPortsType`

GetPhysicalPortsType returns the PhysicalPortsType field if non-nil, zero value otherwise.

### GetPhysicalPortsTypeOk

`func (o *PortResponse) GetPhysicalPortsTypeOk() (*PortResponsePhysicalPortsType, bool)`

GetPhysicalPortsTypeOk returns a tuple with the PhysicalPortsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsType

`func (o *PortResponse) SetPhysicalPortsType(v PortResponsePhysicalPortsType)`

SetPhysicalPortsType sets PhysicalPortsType field to given value.

### HasPhysicalPortsType

`func (o *PortResponse) HasPhysicalPortsType() bool`

HasPhysicalPortsType returns a boolean if a field has been set.

### GetPhysicalPortsCount

`func (o *PortResponse) GetPhysicalPortsCount() int32`

GetPhysicalPortsCount returns the PhysicalPortsCount field if non-nil, zero value otherwise.

### GetPhysicalPortsCountOk

`func (o *PortResponse) GetPhysicalPortsCountOk() (*int32, bool)`

GetPhysicalPortsCountOk returns a tuple with the PhysicalPortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsCount

`func (o *PortResponse) SetPhysicalPortsCount(v int32)`

SetPhysicalPortsCount sets PhysicalPortsCount field to given value.

### HasPhysicalPortsCount

`func (o *PortResponse) HasPhysicalPortsCount() bool`

HasPhysicalPortsCount returns a boolean if a field has been set.

### GetConnectivitySourceType

`func (o *PortResponse) GetConnectivitySourceType() PortResponseConnectivitySourceType`

GetConnectivitySourceType returns the ConnectivitySourceType field if non-nil, zero value otherwise.

### GetConnectivitySourceTypeOk

`func (o *PortResponse) GetConnectivitySourceTypeOk() (*PortResponseConnectivitySourceType, bool)`

GetConnectivitySourceTypeOk returns a tuple with the ConnectivitySourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivitySourceType

`func (o *PortResponse) SetConnectivitySourceType(v PortResponseConnectivitySourceType)`

SetConnectivitySourceType sets ConnectivitySourceType field to given value.

### HasConnectivitySourceType

`func (o *PortResponse) HasConnectivitySourceType() bool`

HasConnectivitySourceType returns a boolean if a field has been set.

### GetBmmrType

`func (o *PortResponse) GetBmmrType() PortResponseBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *PortResponse) GetBmmrTypeOk() (*PortResponseBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *PortResponse) SetBmmrType(v PortResponseBmmrType)`

SetBmmrType sets BmmrType field to given value.

### HasBmmrType

`func (o *PortResponse) HasBmmrType() bool`

HasBmmrType returns a boolean if a field has been set.

### GetProject

`func (o *PortResponse) GetProject() PortResponseProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *PortResponse) GetProjectOk() (*PortResponseProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *PortResponse) SetProject(v PortResponseProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *PortResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *PortResponse) GetState() PortState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PortResponse) GetStateOk() (*PortState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PortResponse) SetState(v PortState)`

SetState sets State field to given value.

### HasState

`func (o *PortResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOrder

`func (o *PortResponse) GetOrder() PortOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PortResponse) GetOrderOk() (*PortOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PortResponse) SetOrder(v PortOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PortResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetCvpId

`func (o *PortResponse) GetCvpId() string`

GetCvpId returns the CvpId field if non-nil, zero value otherwise.

### GetCvpIdOk

`func (o *PortResponse) GetCvpIdOk() (*string, bool)`

GetCvpIdOk returns a tuple with the CvpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvpId

`func (o *PortResponse) SetCvpId(v string)`

SetCvpId sets CvpId field to given value.

### HasCvpId

`func (o *PortResponse) HasCvpId() bool`

HasCvpId returns a boolean if a field has been set.

### GetOperation

`func (o *PortResponse) GetOperation() PortOperation`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *PortResponse) GetOperationOk() (*PortOperation, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *PortResponse) SetOperation(v PortOperation)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *PortResponse) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetAccount

`func (o *PortResponse) GetAccount() SimplifiedAccountPortResponse`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PortResponse) GetAccountOk() (*SimplifiedAccountPortResponse, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PortResponse) SetAccount(v SimplifiedAccountPortResponse)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PortResponse) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetChangelog

`func (o *PortResponse) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *PortResponse) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *PortResponse) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *PortResponse) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.

### GetServiceType

`func (o *PortResponse) GetServiceType() PortResponseServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PortResponse) GetServiceTypeOk() (*PortResponseServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PortResponse) SetServiceType(v PortResponseServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PortResponse) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetBandwidth

`func (o *PortResponse) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *PortResponse) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *PortResponse) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *PortResponse) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetAvailableBandwidth

`func (o *PortResponse) GetAvailableBandwidth() int32`

GetAvailableBandwidth returns the AvailableBandwidth field if non-nil, zero value otherwise.

### GetAvailableBandwidthOk

`func (o *PortResponse) GetAvailableBandwidthOk() (*int32, bool)`

GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBandwidth

`func (o *PortResponse) SetAvailableBandwidth(v int32)`

SetAvailableBandwidth sets AvailableBandwidth field to given value.

### HasAvailableBandwidth

`func (o *PortResponse) HasAvailableBandwidth() bool`

HasAvailableBandwidth returns a boolean if a field has been set.

### GetUsedBandwidth

`func (o *PortResponse) GetUsedBandwidth() int32`

GetUsedBandwidth returns the UsedBandwidth field if non-nil, zero value otherwise.

### GetUsedBandwidthOk

`func (o *PortResponse) GetUsedBandwidthOk() (*int32, bool)`

GetUsedBandwidthOk returns a tuple with the UsedBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBandwidth

`func (o *PortResponse) SetUsedBandwidth(v int32)`

SetUsedBandwidth sets UsedBandwidth field to given value.

### HasUsedBandwidth

`func (o *PortResponse) HasUsedBandwidth() bool`

HasUsedBandwidth returns a boolean if a field has been set.

### GetLocation

`func (o *PortResponse) GetLocation() SimplifiedLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PortResponse) GetLocationOk() (*SimplifiedLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PortResponse) SetLocation(v SimplifiedLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PortResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDevice

`func (o *PortResponse) GetDevice() PortDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *PortResponse) GetDeviceOk() (*PortDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *PortResponse) SetDevice(v PortDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *PortResponse) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterface

`func (o *PortResponse) GetInterface() PortInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *PortResponse) GetInterfaceOk() (*PortInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *PortResponse) SetInterface(v PortInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *PortResponse) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetDemarcationPointIbx

`func (o *PortResponse) GetDemarcationPointIbx() string`

GetDemarcationPointIbx returns the DemarcationPointIbx field if non-nil, zero value otherwise.

### GetDemarcationPointIbxOk

`func (o *PortResponse) GetDemarcationPointIbxOk() (*string, bool)`

GetDemarcationPointIbxOk returns a tuple with the DemarcationPointIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPointIbx

`func (o *PortResponse) SetDemarcationPointIbx(v string)`

SetDemarcationPointIbx sets DemarcationPointIbx field to given value.

### HasDemarcationPointIbx

`func (o *PortResponse) HasDemarcationPointIbx() bool`

HasDemarcationPointIbx returns a boolean if a field has been set.

### GetTetherIbx

`func (o *PortResponse) GetTetherIbx() string`

GetTetherIbx returns the TetherIbx field if non-nil, zero value otherwise.

### GetTetherIbxOk

`func (o *PortResponse) GetTetherIbxOk() (*string, bool)`

GetTetherIbxOk returns a tuple with the TetherIbx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTetherIbx

`func (o *PortResponse) SetTetherIbx(v string)`

SetTetherIbx sets TetherIbx field to given value.

### HasTetherIbx

`func (o *PortResponse) HasTetherIbx() bool`

HasTetherIbx returns a boolean if a field has been set.

### GetDemarcationPoint

`func (o *PortResponse) GetDemarcationPoint() PortDemarcationPoint`

GetDemarcationPoint returns the DemarcationPoint field if non-nil, zero value otherwise.

### GetDemarcationPointOk

`func (o *PortResponse) GetDemarcationPointOk() (*PortDemarcationPoint, bool)`

GetDemarcationPointOk returns a tuple with the DemarcationPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemarcationPoint

`func (o *PortResponse) SetDemarcationPoint(v PortDemarcationPoint)`

SetDemarcationPoint sets DemarcationPoint field to given value.

### HasDemarcationPoint

`func (o *PortResponse) HasDemarcationPoint() bool`

HasDemarcationPoint returns a boolean if a field has been set.

### GetRedundancy

`func (o *PortResponse) GetRedundancy() PortRedundancy`

GetRedundancy returns the Redundancy field if non-nil, zero value otherwise.

### GetRedundancyOk

`func (o *PortResponse) GetRedundancyOk() (*PortRedundancy, bool)`

GetRedundancyOk returns a tuple with the Redundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedundancy

`func (o *PortResponse) SetRedundancy(v PortRedundancy)`

SetRedundancy sets Redundancy field to given value.

### HasRedundancy

`func (o *PortResponse) HasRedundancy() bool`

HasRedundancy returns a boolean if a field has been set.

### GetEncapsulation

`func (o *PortResponse) GetEncapsulation() PortEncapsulation`

GetEncapsulation returns the Encapsulation field if non-nil, zero value otherwise.

### GetEncapsulationOk

`func (o *PortResponse) GetEncapsulationOk() (*PortEncapsulation, bool)`

GetEncapsulationOk returns a tuple with the Encapsulation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulation

`func (o *PortResponse) SetEncapsulation(v PortEncapsulation)`

SetEncapsulation sets Encapsulation field to given value.

### HasEncapsulation

`func (o *PortResponse) HasEncapsulation() bool`

HasEncapsulation returns a boolean if a field has been set.

### GetLagEnabled

`func (o *PortResponse) GetLagEnabled() bool`

GetLagEnabled returns the LagEnabled field if non-nil, zero value otherwise.

### GetLagEnabledOk

`func (o *PortResponse) GetLagEnabledOk() (*bool, bool)`

GetLagEnabledOk returns a tuple with the LagEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagEnabled

`func (o *PortResponse) SetLagEnabled(v bool)`

SetLagEnabled sets LagEnabled field to given value.

### HasLagEnabled

`func (o *PortResponse) HasLagEnabled() bool`

HasLagEnabled returns a boolean if a field has been set.

### GetLag

`func (o *PortResponse) GetLag() PortLag`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *PortResponse) GetLagOk() (*PortLag, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *PortResponse) SetLag(v PortLag)`

SetLag sets Lag field to given value.

### HasLag

`func (o *PortResponse) HasLag() bool`

HasLag returns a boolean if a field has been set.

### GetAsn

`func (o *PortResponse) GetAsn() int32`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *PortResponse) GetAsnOk() (*int32, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *PortResponse) SetAsn(v int32)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *PortResponse) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetSettings

`func (o *PortResponse) GetSettings() PortSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *PortResponse) GetSettingsOk() (*PortSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *PortResponse) SetSettings(v PortSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *PortResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetPhysicalPortQuantity

`func (o *PortResponse) GetPhysicalPortQuantity() int32`

GetPhysicalPortQuantity returns the PhysicalPortQuantity field if non-nil, zero value otherwise.

### GetPhysicalPortQuantityOk

`func (o *PortResponse) GetPhysicalPortQuantityOk() (*int32, bool)`

GetPhysicalPortQuantityOk returns a tuple with the PhysicalPortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortQuantity

`func (o *PortResponse) SetPhysicalPortQuantity(v int32)`

SetPhysicalPortQuantity sets PhysicalPortQuantity field to given value.

### HasPhysicalPortQuantity

`func (o *PortResponse) HasPhysicalPortQuantity() bool`

HasPhysicalPortQuantity returns a boolean if a field has been set.

### GetNotifications

`func (o *PortResponse) GetNotifications() []PortNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *PortResponse) GetNotificationsOk() (*[]PortNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *PortResponse) SetNotifications(v []PortNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *PortResponse) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *PortResponse) GetAdditionalInfo() []PortAdditionalInfo`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *PortResponse) GetAdditionalInfoOk() (*[]PortAdditionalInfo, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *PortResponse) SetAdditionalInfo(v []PortAdditionalInfo)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *PortResponse) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetPhysicalPorts

`func (o *PortResponse) GetPhysicalPorts() []PhysicalPort`

GetPhysicalPorts returns the PhysicalPorts field if non-nil, zero value otherwise.

### GetPhysicalPortsOk

`func (o *PortResponse) GetPhysicalPortsOk() (*[]PhysicalPort, bool)`

GetPhysicalPortsOk returns a tuple with the PhysicalPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPorts

`func (o *PortResponse) SetPhysicalPorts(v []PhysicalPort)`

SetPhysicalPorts sets PhysicalPorts field to given value.

### HasPhysicalPorts

`func (o *PortResponse) HasPhysicalPorts() bool`

HasPhysicalPorts returns a boolean if a field has been set.

### GetLoas

`func (o *PortResponse) GetLoas() []PortLoa`

GetLoas returns the Loas field if non-nil, zero value otherwise.

### GetLoasOk

`func (o *PortResponse) GetLoasOk() (*[]PortLoa, bool)`

GetLoasOk returns a tuple with the Loas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoas

`func (o *PortResponse) SetLoas(v []PortLoa)`

SetLoas sets Loas field to given value.

### HasLoas

`func (o *PortResponse) HasLoas() bool`

HasLoas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


