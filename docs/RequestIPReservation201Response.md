# RequestIPReservation201Response

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

### NewRequestIPReservation201Response

`func NewRequestIPReservation201Response(type_ VrfIpReservationType, vrf Vrf, ) *RequestIPReservation201Response`

NewRequestIPReservation201Response instantiates a new RequestIPReservation201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestIPReservation201ResponseWithDefaults

`func NewRequestIPReservation201ResponseWithDefaults() *RequestIPReservation201Response`

NewRequestIPReservation201ResponseWithDefaults instantiates a new RequestIPReservation201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *RequestIPReservation201Response) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *RequestIPReservation201Response) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *RequestIPReservation201Response) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *RequestIPReservation201Response) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddress

`func (o *RequestIPReservation201Response) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RequestIPReservation201Response) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RequestIPReservation201Response) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RequestIPReservation201Response) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *RequestIPReservation201Response) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *RequestIPReservation201Response) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *RequestIPReservation201Response) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *RequestIPReservation201Response) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignments

`func (o *RequestIPReservation201Response) GetAssignments() []IPAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *RequestIPReservation201Response) GetAssignmentsOk() (*[]IPAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *RequestIPReservation201Response) SetAssignments(v []IPAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *RequestIPReservation201Response) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *RequestIPReservation201Response) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *RequestIPReservation201Response) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *RequestIPReservation201Response) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *RequestIPReservation201Response) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *RequestIPReservation201Response) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *RequestIPReservation201Response) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *RequestIPReservation201Response) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *RequestIPReservation201Response) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCidr

`func (o *RequestIPReservation201Response) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *RequestIPReservation201Response) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *RequestIPReservation201Response) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *RequestIPReservation201Response) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RequestIPReservation201Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RequestIPReservation201Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RequestIPReservation201Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RequestIPReservation201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *RequestIPReservation201Response) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *RequestIPReservation201Response) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *RequestIPReservation201Response) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *RequestIPReservation201Response) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEnabled

`func (o *RequestIPReservation201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestIPReservation201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestIPReservation201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RequestIPReservation201Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetails

`func (o *RequestIPReservation201Response) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RequestIPReservation201Response) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RequestIPReservation201Response) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *RequestIPReservation201Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *RequestIPReservation201Response) GetFacility() IPReservationFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *RequestIPReservation201Response) GetFacilityOk() (*IPReservationFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *RequestIPReservation201Response) SetFacility(v IPReservationFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *RequestIPReservation201Response) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetGateway

`func (o *RequestIPReservation201Response) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *RequestIPReservation201Response) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *RequestIPReservation201Response) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *RequestIPReservation201Response) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *RequestIPReservation201Response) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *RequestIPReservation201Response) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *RequestIPReservation201Response) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *RequestIPReservation201Response) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *RequestIPReservation201Response) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RequestIPReservation201Response) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RequestIPReservation201Response) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RequestIPReservation201Response) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RequestIPReservation201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestIPReservation201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestIPReservation201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RequestIPReservation201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *RequestIPReservation201Response) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *RequestIPReservation201Response) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *RequestIPReservation201Response) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *RequestIPReservation201Response) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *RequestIPReservation201Response) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *RequestIPReservation201Response) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *RequestIPReservation201Response) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *RequestIPReservation201Response) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetalGateway

`func (o *RequestIPReservation201Response) GetMetalGateway() MetalGatewayLite`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *RequestIPReservation201Response) GetMetalGatewayOk() (*MetalGatewayLite, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *RequestIPReservation201Response) SetMetalGateway(v MetalGatewayLite)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *RequestIPReservation201Response) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *RequestIPReservation201Response) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *RequestIPReservation201Response) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *RequestIPReservation201Response) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *RequestIPReservation201Response) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *RequestIPReservation201Response) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *RequestIPReservation201Response) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *RequestIPReservation201Response) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *RequestIPReservation201Response) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *RequestIPReservation201Response) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *RequestIPReservation201Response) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *RequestIPReservation201Response) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *RequestIPReservation201Response) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *RequestIPReservation201Response) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RequestIPReservation201Response) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RequestIPReservation201Response) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *RequestIPReservation201Response) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *RequestIPReservation201Response) GetProjectLite() Project`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *RequestIPReservation201Response) GetProjectLiteOk() (*Project, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *RequestIPReservation201Response) SetProjectLite(v Project)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *RequestIPReservation201Response) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *RequestIPReservation201Response) GetRequestedBy() Href`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RequestIPReservation201Response) GetRequestedByOk() (*Href, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RequestIPReservation201Response) SetRequestedBy(v Href)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *RequestIPReservation201Response) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetPublic

`func (o *RequestIPReservation201Response) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *RequestIPReservation201Response) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *RequestIPReservation201Response) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *RequestIPReservation201Response) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *RequestIPReservation201Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RequestIPReservation201Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RequestIPReservation201Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RequestIPReservation201Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *RequestIPReservation201Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *RequestIPReservation201Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *RequestIPReservation201Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *RequestIPReservation201Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *RequestIPReservation201Response) GetType() VrfIpReservationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestIPReservation201Response) GetTypeOk() (*VrfIpReservationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestIPReservation201Response) SetType(v VrfIpReservationType)`

SetType sets Type field to given value.


### GetCreatedBy

`func (o *RequestIPReservation201Response) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestIPReservation201Response) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestIPReservation201Response) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RequestIPReservation201Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetVrf

`func (o *RequestIPReservation201Response) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *RequestIPReservation201Response) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *RequestIPReservation201Response) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


