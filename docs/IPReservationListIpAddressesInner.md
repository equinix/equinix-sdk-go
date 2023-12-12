# IPReservationListIpAddressesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addon** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Assignments** | Pointer to [**[]IPAssignment**](IPAssignment.md) |  | [optional] 
**Available** | Pointer to **string** |  | [optional] 
**Bill** | Pointer to **bool** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Facility** | Pointer to [**IPReservationFacility**](IPReservationFacility.md) |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**GlobalIp** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**MetalGateway** | Pointer to [**MetalGatewayLite**](MetalGatewayLite.md) |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectLite** | Pointer to [**Project**](Project.md) |  | [optional] 
**RequestedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**VrfIpReservationType**](VrfIpReservationType.md) |  | 
**CreatedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Vrf** | [**Vrf**](Vrf.md) |  | 

## Methods

### NewIPReservationListIpAddressesInner

`func NewIPReservationListIpAddressesInner(type_ VrfIpReservationType, vrf Vrf, ) *IPReservationListIpAddressesInner`

NewIPReservationListIpAddressesInner instantiates a new IPReservationListIpAddressesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReservationListIpAddressesInnerWithDefaults

`func NewIPReservationListIpAddressesInnerWithDefaults() *IPReservationListIpAddressesInner`

NewIPReservationListIpAddressesInnerWithDefaults instantiates a new IPReservationListIpAddressesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *IPReservationListIpAddressesInner) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *IPReservationListIpAddressesInner) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *IPReservationListIpAddressesInner) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *IPReservationListIpAddressesInner) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddress

`func (o *IPReservationListIpAddressesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPReservationListIpAddressesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPReservationListIpAddressesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPReservationListIpAddressesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *IPReservationListIpAddressesInner) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPReservationListIpAddressesInner) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPReservationListIpAddressesInner) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPReservationListIpAddressesInner) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignments

`func (o *IPReservationListIpAddressesInner) GetAssignments() []IPAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *IPReservationListIpAddressesInner) GetAssignmentsOk() (*[]IPAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *IPReservationListIpAddressesInner) SetAssignments(v []IPAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *IPReservationListIpAddressesInner) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *IPReservationListIpAddressesInner) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *IPReservationListIpAddressesInner) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *IPReservationListIpAddressesInner) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *IPReservationListIpAddressesInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *IPReservationListIpAddressesInner) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *IPReservationListIpAddressesInner) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *IPReservationListIpAddressesInner) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *IPReservationListIpAddressesInner) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCidr

`func (o *IPReservationListIpAddressesInner) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IPReservationListIpAddressesInner) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IPReservationListIpAddressesInner) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IPReservationListIpAddressesInner) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPReservationListIpAddressesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPReservationListIpAddressesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPReservationListIpAddressesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPReservationListIpAddressesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *IPReservationListIpAddressesInner) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *IPReservationListIpAddressesInner) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *IPReservationListIpAddressesInner) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *IPReservationListIpAddressesInner) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEnabled

`func (o *IPReservationListIpAddressesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IPReservationListIpAddressesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IPReservationListIpAddressesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IPReservationListIpAddressesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetails

`func (o *IPReservationListIpAddressesInner) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IPReservationListIpAddressesInner) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IPReservationListIpAddressesInner) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IPReservationListIpAddressesInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *IPReservationListIpAddressesInner) GetFacility() IPReservationFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *IPReservationListIpAddressesInner) GetFacilityOk() (*IPReservationFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *IPReservationListIpAddressesInner) SetFacility(v IPReservationFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *IPReservationListIpAddressesInner) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetGateway

`func (o *IPReservationListIpAddressesInner) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IPReservationListIpAddressesInner) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IPReservationListIpAddressesInner) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IPReservationListIpAddressesInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *IPReservationListIpAddressesInner) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *IPReservationListIpAddressesInner) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *IPReservationListIpAddressesInner) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *IPReservationListIpAddressesInner) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *IPReservationListIpAddressesInner) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPReservationListIpAddressesInner) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPReservationListIpAddressesInner) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IPReservationListIpAddressesInner) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *IPReservationListIpAddressesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPReservationListIpAddressesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPReservationListIpAddressesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPReservationListIpAddressesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *IPReservationListIpAddressesInner) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *IPReservationListIpAddressesInner) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *IPReservationListIpAddressesInner) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *IPReservationListIpAddressesInner) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *IPReservationListIpAddressesInner) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *IPReservationListIpAddressesInner) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *IPReservationListIpAddressesInner) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *IPReservationListIpAddressesInner) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetalGateway

`func (o *IPReservationListIpAddressesInner) GetMetalGateway() MetalGatewayLite`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *IPReservationListIpAddressesInner) GetMetalGatewayOk() (*MetalGatewayLite, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *IPReservationListIpAddressesInner) SetMetalGateway(v MetalGatewayLite)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *IPReservationListIpAddressesInner) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *IPReservationListIpAddressesInner) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *IPReservationListIpAddressesInner) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *IPReservationListIpAddressesInner) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *IPReservationListIpAddressesInner) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *IPReservationListIpAddressesInner) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IPReservationListIpAddressesInner) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IPReservationListIpAddressesInner) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IPReservationListIpAddressesInner) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *IPReservationListIpAddressesInner) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IPReservationListIpAddressesInner) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IPReservationListIpAddressesInner) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IPReservationListIpAddressesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *IPReservationListIpAddressesInner) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IPReservationListIpAddressesInner) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IPReservationListIpAddressesInner) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *IPReservationListIpAddressesInner) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *IPReservationListIpAddressesInner) GetProjectLite() Project`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *IPReservationListIpAddressesInner) GetProjectLiteOk() (*Project, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *IPReservationListIpAddressesInner) SetProjectLite(v Project)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *IPReservationListIpAddressesInner) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *IPReservationListIpAddressesInner) GetRequestedBy() Href`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *IPReservationListIpAddressesInner) GetRequestedByOk() (*Href, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *IPReservationListIpAddressesInner) SetRequestedBy(v Href)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *IPReservationListIpAddressesInner) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetPublic

`func (o *IPReservationListIpAddressesInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPReservationListIpAddressesInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPReservationListIpAddressesInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPReservationListIpAddressesInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *IPReservationListIpAddressesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPReservationListIpAddressesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPReservationListIpAddressesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IPReservationListIpAddressesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *IPReservationListIpAddressesInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPReservationListIpAddressesInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPReservationListIpAddressesInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPReservationListIpAddressesInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *IPReservationListIpAddressesInner) GetType() VrfIpReservationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPReservationListIpAddressesInner) GetTypeOk() (*VrfIpReservationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPReservationListIpAddressesInner) SetType(v VrfIpReservationType)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *IPReservationListIpAddressesInner) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IPReservationListIpAddressesInner) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IPReservationListIpAddressesInner) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *IPReservationListIpAddressesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVrf

`func (o *IPReservationListIpAddressesInner) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *IPReservationListIpAddressesInner) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *IPReservationListIpAddressesInner) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


