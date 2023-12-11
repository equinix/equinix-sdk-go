# VrfIpReservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressFamily** | Pointer to **int32** |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**Href**](Href.md) |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MetalGateway** | Pointer to [**MetalGatewayLite**](MetalGatewayLite.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Project** | Pointer to [**Project**](Project.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**VrfIpReservationType**](VrfIpReservationType.md) |  | 
**Vrf** | [**Vrf**](Vrf.md) |  | 
**Public** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Customdata** | Pointer to **map[string]interface{}** |  | [optional] 
**Bill** | Pointer to **bool** |  | [optional] 
**ProjectLite** | Pointer to [**Project**](Project.md) |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**Metro** | Pointer to [**Metro**](Metro.md) |  | [optional] 

## Methods

### NewVrfIpReservation

`func NewVrfIpReservation(type_ VrfIpReservationType, vrf Vrf, ) *VrfIpReservation`

NewVrfIpReservation instantiates a new VrfIpReservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVrfIpReservationWithDefaults

`func NewVrfIpReservationWithDefaults() *VrfIpReservation`

NewVrfIpReservationWithDefaults instantiates a new VrfIpReservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressFamily

`func (o *VrfIpReservation) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *VrfIpReservation) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *VrfIpReservation) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *VrfIpReservation) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetCidr

`func (o *VrfIpReservation) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *VrfIpReservation) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *VrfIpReservation) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *VrfIpReservation) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VrfIpReservation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VrfIpReservation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VrfIpReservation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VrfIpReservation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedBy

`func (o *VrfIpReservation) GetCreatedBy() Href`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VrfIpReservation) GetCreatedByOk() (*Href, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VrfIpReservation) SetCreatedBy(v Href)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VrfIpReservation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDetails

`func (o *VrfIpReservation) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VrfIpReservation) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VrfIpReservation) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *VrfIpReservation) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetHref

`func (o *VrfIpReservation) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VrfIpReservation) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VrfIpReservation) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VrfIpReservation) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *VrfIpReservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VrfIpReservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VrfIpReservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VrfIpReservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetalGateway

`func (o *VrfIpReservation) GetMetalGateway() MetalGatewayLite`

GetMetalGateway returns the MetalGateway field if non-nil, zero value otherwise.

### GetMetalGatewayOk

`func (o *VrfIpReservation) GetMetalGatewayOk() (*MetalGatewayLite, bool)`

GetMetalGatewayOk returns a tuple with the MetalGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetalGateway

`func (o *VrfIpReservation) SetMetalGateway(v MetalGatewayLite)`

SetMetalGateway sets MetalGateway field to given value.

### HasMetalGateway

`func (o *VrfIpReservation) HasMetalGateway() bool`

HasMetalGateway returns a boolean if a field has been set.

### GetNetmask

`func (o *VrfIpReservation) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *VrfIpReservation) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *VrfIpReservation) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *VrfIpReservation) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *VrfIpReservation) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *VrfIpReservation) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *VrfIpReservation) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *VrfIpReservation) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetProject

`func (o *VrfIpReservation) GetProject() Project`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VrfIpReservation) GetProjectOk() (*Project, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VrfIpReservation) SetProject(v Project)`

SetProject sets Project field to given value.

### HasProject

`func (o *VrfIpReservation) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetState

`func (o *VrfIpReservation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VrfIpReservation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VrfIpReservation) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *VrfIpReservation) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *VrfIpReservation) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VrfIpReservation) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VrfIpReservation) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VrfIpReservation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *VrfIpReservation) GetType() VrfIpReservationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VrfIpReservation) GetTypeOk() (*VrfIpReservationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VrfIpReservation) SetType(v VrfIpReservationType)`

SetType sets Type field to given value.


### GetVrf

`func (o *VrfIpReservation) GetVrf() Vrf`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *VrfIpReservation) GetVrfOk() (*Vrf, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *VrfIpReservation) SetVrf(v Vrf)`

SetVrf sets Vrf field to given value.


### GetPublic

`func (o *VrfIpReservation) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *VrfIpReservation) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *VrfIpReservation) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *VrfIpReservation) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetManagement

`func (o *VrfIpReservation) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *VrfIpReservation) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *VrfIpReservation) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *VrfIpReservation) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetManageable

`func (o *VrfIpReservation) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *VrfIpReservation) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *VrfIpReservation) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *VrfIpReservation) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetCustomdata

`func (o *VrfIpReservation) GetCustomdata() map[string]interface{}`

GetCustomdata returns the Customdata field if non-nil, zero value otherwise.

### GetCustomdataOk

`func (o *VrfIpReservation) GetCustomdataOk() (*map[string]interface{}, bool)`

GetCustomdataOk returns a tuple with the Customdata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomdata

`func (o *VrfIpReservation) SetCustomdata(v map[string]interface{})`

SetCustomdata sets Customdata field to given value.

### HasCustomdata

`func (o *VrfIpReservation) HasCustomdata() bool`

HasCustomdata returns a boolean if a field has been set.

### GetBill

`func (o *VrfIpReservation) GetBill() bool`

GetBill returns the Bill field if non-nil, zero value otherwise.

### GetBillOk

`func (o *VrfIpReservation) GetBillOk() (*bool, bool)`

GetBillOk returns a tuple with the Bill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBill

`func (o *VrfIpReservation) SetBill(v bool)`

SetBill sets Bill field to given value.

### HasBill

`func (o *VrfIpReservation) HasBill() bool`

HasBill returns a boolean if a field has been set.

### GetProjectLite

`func (o *VrfIpReservation) GetProjectLite() Project`

GetProjectLite returns the ProjectLite field if non-nil, zero value otherwise.

### GetProjectLiteOk

`func (o *VrfIpReservation) GetProjectLiteOk() (*Project, bool)`

GetProjectLiteOk returns a tuple with the ProjectLite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectLite

`func (o *VrfIpReservation) SetProjectLite(v Project)`

SetProjectLite sets ProjectLite field to given value.

### HasProjectLite

`func (o *VrfIpReservation) HasProjectLite() bool`

HasProjectLite returns a boolean if a field has been set.

### GetAddress

`func (o *VrfIpReservation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VrfIpReservation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VrfIpReservation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VrfIpReservation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGateway

`func (o *VrfIpReservation) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *VrfIpReservation) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *VrfIpReservation) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *VrfIpReservation) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMetro

`func (o *VrfIpReservation) GetMetro() Metro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *VrfIpReservation) GetMetroOk() (*Metro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *VrfIpReservation) SetMetro(v Metro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *VrfIpReservation) HasMetro() bool`

HasMetro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


