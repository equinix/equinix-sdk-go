# PortRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PortType**](PortType.md) |  | 
**PhysicalPortsSpeed** | **int32** | Physical Ports Speed in Mbps | 
**PhysicalPortsType** | [**PortRequestPhysicalPortsType**](PortRequestPhysicalPortsType.md) |  | 
**PhysicalPortsCount** | Pointer to **int32** |  | [optional] 
**ConnectivitySourceType** | [**PortConnectivitySourceType**](PortConnectivitySourceType.md) |  | 
**BmmrType** | Pointer to [**PortRequestBmmrType**](PortRequestBmmrType.md) |  | [optional] 
**Project** | [**Project**](Project.md) |  | 
**Order** | Pointer to [**PortOrder**](PortOrder.md) |  | [optional] 
**Account** | [**SimplifiedAccountRequest**](SimplifiedAccountRequest.md) |  | 
**ServiceType** | Pointer to [**PortServiceType**](PortServiceType.md) |  | [optional] 
**ServiceCode** | Pointer to [**PortServiceCode**](PortServiceCode.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Equinix assigned response attribute for Port bandwidth in Mbps | [optional] 
**Location** | [**SimplifiedLocationRequest**](SimplifiedLocationRequest.md) |  | 
**DemarcationPointIbx** | Pointer to **string** | A-side/Equinix ibx | [optional] 
**TetherIbx** | Pointer to **string** | z-side/Equinix ibx | [optional] 
**DemarcationPoint** | Pointer to [**PortDemarcationPoint**](PortDemarcationPoint.md) |  | [optional] 
**Redundancy** | Pointer to [**PortRedundancy**](PortRedundancy.md) |  | [optional] 
**Encapsulation** | [**PortEncapsulation**](PortEncapsulation.md) |  | 
**LagEnabled** | Pointer to **bool** | Indicates whether Link Aggregation Group (LAG) is enabled on this port | [optional] 
**Package** | Pointer to [**Package**](Package.md) |  | [optional] 
**Settings** | Pointer to [**PortSettings**](PortSettings.md) |  | [optional] 
**Notifications** | Pointer to [**[]PortNotification**](PortNotification.md) | Notification preferences | [optional] 
**PhysicalPorts** | Pointer to [**[]PhysicalPort**](PhysicalPort.md) | Physical ports that implement this port | [optional] 
**Loas** | Pointer to [**[]PortLoa**](PortLoa.md) | Port Loas | [optional] 

## Methods

### NewPortRequest

`func NewPortRequest(type_ PortType, physicalPortsSpeed int32, physicalPortsType PortRequestPhysicalPortsType, connectivitySourceType PortConnectivitySourceType, project Project, account SimplifiedAccountRequest, location SimplifiedLocationRequest, encapsulation PortEncapsulation, ) *PortRequest`

NewPortRequest instantiates a new PortRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortRequestWithDefaults

`func NewPortRequestWithDefaults() *PortRequest`

NewPortRequestWithDefaults instantiates a new PortRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetPhysicalPortsType

`func (o *PortRequest) GetPhysicalPortsType() PortRequestPhysicalPortsType`

GetPhysicalPortsType returns the PhysicalPortsType field if non-nil, zero value otherwise.

### GetPhysicalPortsTypeOk

`func (o *PortRequest) GetPhysicalPortsTypeOk() (*PortRequestPhysicalPortsType, bool)`

GetPhysicalPortsTypeOk returns a tuple with the PhysicalPortsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalPortsType

`func (o *PortRequest) SetPhysicalPortsType(v PortRequestPhysicalPortsType)`

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

`func (o *PortRequest) GetBmmrType() PortRequestBmmrType`

GetBmmrType returns the BmmrType field if non-nil, zero value otherwise.

### GetBmmrTypeOk

`func (o *PortRequest) GetBmmrTypeOk() (*PortRequestBmmrType, bool)`

GetBmmrTypeOk returns a tuple with the BmmrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmmrType

`func (o *PortRequest) SetBmmrType(v PortRequestBmmrType)`

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

### GetAccount

`func (o *PortRequest) GetAccount() SimplifiedAccountRequest`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PortRequest) GetAccountOk() (*SimplifiedAccountRequest, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PortRequest) SetAccount(v SimplifiedAccountRequest)`

SetAccount sets Account field to given value.


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

### GetServiceCode

`func (o *PortRequest) GetServiceCode() PortServiceCode`

GetServiceCode returns the ServiceCode field if non-nil, zero value otherwise.

### GetServiceCodeOk

`func (o *PortRequest) GetServiceCodeOk() (*PortServiceCode, bool)`

GetServiceCodeOk returns a tuple with the ServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCode

`func (o *PortRequest) SetServiceCode(v PortServiceCode)`

SetServiceCode sets ServiceCode field to given value.

### HasServiceCode

`func (o *PortRequest) HasServiceCode() bool`

HasServiceCode returns a boolean if a field has been set.

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

### GetLocation

`func (o *PortRequest) GetLocation() SimplifiedLocationRequest`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PortRequest) GetLocationOk() (*SimplifiedLocationRequest, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PortRequest) SetLocation(v SimplifiedLocationRequest)`

SetLocation sets Location field to given value.


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

### HasSettings

`func (o *PortRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

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


