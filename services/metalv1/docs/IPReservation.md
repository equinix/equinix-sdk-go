# IPReservation

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
**Metro** | Pointer to [**IPReservationMetro**](IPReservationMetro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**ProjectLite** | Pointer to [**Href**](Href.md) |  | [optional] 
**RequestedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**IPReservationType**](IPReservationType.md) |  | 

## Methods

### NewIPReservation

`func NewIPReservation(type_ IPReservationType, ) *IPReservation`

NewIPReservation instantiates a new IPReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPReservationWithDefaults

`func NewIPReservationWithDefaults() *IPReservation`

NewIPReservationWithDefaults instantiates a new IPReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddon

`func (o *IPReservation) GetAddon() bool`

GetAddon returns the Addon field if non-nil, zero value otherwise.

### GetAddonOk

`func (o *IPReservation) GetAddonOk() (*bool, bool)`

GetAddonOk returns a tuple with the Addon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddon

`func (o *IPReservation) SetAddon(v bool)`

SetAddon sets Addon field to given value.

### HasAddon

`func (o *IPReservation) HasAddon() bool`

HasAddon returns a boolean if a field has been set.

### GetAddress

`func (o *IPReservation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPReservation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPReservation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPReservation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *IPReservation) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPReservation) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPReservation) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPReservation) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignments

`func (o *IPReservation) GetAssignments() []IPAssignment`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *IPReservation) GetAssignmentsOk() (*[]IPAssignment, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *IPReservation) SetAssignments(v []IPAssignment)`

SetAssignments sets Assignments field to given value.

### HasAssignments

`func (o *IPReservation) HasAssignments() bool`

HasAssignments returns a boolean if a field has been set.

### GetAvailable

`func (o *IPReservation) GetAvailable() string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *IPReservation) GetAvailableOk() (*string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *IPReservation) SetAvailable(v string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *IPReservation) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetBill

`func (o *IPReservation) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *IPReservation) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *IPReservation) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *IPReservation) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetCidr

`func (o *IPReservation) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IPReservation) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IPReservation) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IPReservation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPReservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPReservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPReservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomdata

`func (o *IPReservation) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *IPReservation) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *IPReservation) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *IPReservation) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetEnabled

`func (o *IPReservation) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IPReservation) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IPReservation) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IPReservation) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDetails

`func (o *IPReservation) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *IPReservation) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *IPReservation) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *IPReservation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetFacility

`func (o *IPReservation) GetFacility() IPReservationFacility`

GetFacility returns the Facility field if non-nil, zero value otherwise.

### GetFacilityOk

`func (o *IPReservation) GetFacilityOk() (*IPReservationFacility, bool)`

GetFacilityOk returns a tuple with the Facility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacility

`func (o *IPReservation) SetFacility(v IPReservationFacility)`

SetFacility sets Facility field to given value.

### HasFacility

`func (o *IPReservation) HasFacility() bool`

HasFacility returns a boolean if a field has been set.

### GetGateway

`func (o *IPReservation) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IPReservation) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IPReservation) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IPReservation) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *IPReservation) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *IPReservation) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *IPReservation) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *IPReservation) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *IPReservation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPReservation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPReservation) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IPReservation) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *IPReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *IPReservation) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *IPReservation) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *IPReservation) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *IPReservation) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *IPReservation) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *IPReservation) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *IPReservation) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *IPReservation) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetalGateway

`func (o *IPReservation) GetMetalGateway() MetalGatewayLite`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *IPReservation) GetMetalGatewayOk() (*MetalGatewayLite, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *IPReservation) SetMetalGateway(v MetalGatewayLite)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *IPReservation) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetMetro

`func (o *IPReservation) GetMetro() IPReservationMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *IPReservation) GetMetroOk() (*IPReservationMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *IPReservation) SetMetro(v IPReservationMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *IPReservation) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *IPReservation) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IPReservation) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IPReservation) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IPReservation) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *IPReservation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IPReservation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IPReservation) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IPReservation) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *IPReservation) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *IPReservation) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *IPReservation) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *IPReservation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectLite

`func (o *IPReservation) GetProjectLite() Href`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *IPReservation) GetProjectLiteOk() (*Href, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *IPReservation) SetProjectLite(v Href)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *IPReservation) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetRequestedBy

`func (o *IPReservation) GetRequestedBy() Href`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *IPReservation) GetRequestedByOk() (*Href, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *IPReservation) SetRequestedBy(v Href)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *IPReservation) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetPublic

`func (o *IPReservation) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPReservation) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPReservation) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPReservation) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *IPReservation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPReservation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPReservation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IPReservation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *IPReservation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IPReservation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IPReservation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IPReservation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *IPReservation) GetType() IPReservationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPReservation) GetTypeOk() (*IPReservationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPReservation) SetType(v IPReservationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


