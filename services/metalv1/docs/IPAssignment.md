# IPAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**AddressFamily** | Pointer to **int32** |  | [optional] 
**AssignedTo** | Pointer to [**Href**](Href.md) |  | [optional] 
**Cidr** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**GlobalIp** | Pointer to **bool** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Manageable** | Pointer to **bool** |  | [optional] 
**Management** | Pointer to **bool** |  | [optional] 
**Metro** | Pointer to [**IPAssignmentMetro**](IPAssignmentMetro.md) |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**ParentBlock** | Pointer to [**ParentBlock**](ParentBlock.md) |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 
**State** | Pointer to [**IPAssignmentState**](IPAssignmentState.md) |  | [optional] 
**NextHop** | Pointer to **string** | Only set when this is a Metal Gateway Elastic IP Assignment.  The IP address within the Metal Gateway to which requests to the Elastic IP are forwarded.  | [optional] 
**Type** | Pointer to [**IPAssignmentType**](IPAssignmentType.md) |  | [optional] 

## Methods

### NewIPAssignment

`func NewIPAssignment() *IPAssignment`

NewIPAssignment instantiates a new IPAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAssignmentWithDefaults

`func NewIPAssignmentWithDefaults() *IPAssignment`

NewIPAssignmentWithDefaults instantiates a new IPAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IPAssignment) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IPAssignment) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IPAssignment) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IPAssignment) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressFamily

`func (o *IPAssignment) GetAddressFamily() int32`

GetAddressFamily returns the AddressFamily field if non-nil, zero value otherwise.

### GetAddressFamilyOk

`func (o *IPAssignment) GetAddressFamilyOk() (*int32, bool)`

GetAddressFamilyOk returns a tuple with the AddressFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressFamily

`func (o *IPAssignment) SetAddressFamily(v int32)`

SetAddressFamily sets AddressFamily field to given value.

### HasAddressFamily

`func (o *IPAssignment) HasAddressFamily() bool`

HasAddressFamily returns a boolean if a field has been set.

### GetAssignedTo

`func (o *IPAssignment) GetAssignedTo() Href`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *IPAssignment) GetAssignedToOk() (*Href, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *IPAssignment) SetAssignedTo(v Href)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *IPAssignment) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetCidr

`func (o *IPAssignment) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *IPAssignment) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *IPAssignment) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *IPAssignment) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IPAssignment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IPAssignment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IPAssignment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IPAssignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *IPAssignment) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IPAssignment) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IPAssignment) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IPAssignment) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGateway

`func (o *IPAssignment) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *IPAssignment) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *IPAssignment) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *IPAssignment) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetGlobalIp

`func (o *IPAssignment) GetGlobalIp() bool`

GetGlobalIp returns the GlobalIp field if non-nil, zero value otherwise.

### GetGlobalIpOk

`func (o *IPAssignment) GetGlobalIpOk() (*bool, bool)`

GetGlobalIpOk returns a tuple with the GlobalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIp

`func (o *IPAssignment) SetGlobalIp(v bool)`

SetGlobalIp sets GlobalIp field to given value.

### HasGlobalIp

`func (o *IPAssignment) HasGlobalIp() bool`

HasGlobalIp returns a boolean if a field has been set.

### GetHref

`func (o *IPAssignment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPAssignment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPAssignment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IPAssignment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *IPAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IPAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManageable

`func (o *IPAssignment) GetManageable() bool`

GetManageable returns the Manageable field if non-nil, zero value otherwise.

### GetManageableOk

`func (o *IPAssignment) GetManageableOk() (*bool, bool)`

GetManageableOk returns a tuple with the Manageable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageable

`func (o *IPAssignment) SetManageable(v bool)`

SetManageable sets Manageable field to given value.

### HasManageable

`func (o *IPAssignment) HasManageable() bool`

HasManageable returns a boolean if a field has been set.

### GetManagement

`func (o *IPAssignment) GetManagement() bool`

GetManagement returns the Management field if non-nil, zero value otherwise.

### GetManagementOk

`func (o *IPAssignment) GetManagementOk() (*bool, bool)`

GetManagementOk returns a tuple with the Management field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagement

`func (o *IPAssignment) SetManagement(v bool)`

SetManagement sets Management field to given value.

### HasManagement

`func (o *IPAssignment) HasManagement() bool`

HasManagement returns a boolean if a field has been set.

### GetMetro

`func (o *IPAssignment) GetMetro() IPAssignmentMetro`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *IPAssignment) GetMetroOk() (*IPAssignmentMetro, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *IPAssignment) SetMetro(v IPAssignmentMetro)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *IPAssignment) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetNetmask

`func (o *IPAssignment) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *IPAssignment) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *IPAssignment) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *IPAssignment) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *IPAssignment) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *IPAssignment) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *IPAssignment) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *IPAssignment) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetParentBlock

`func (o *IPAssignment) GetParentBlock() ParentBlock`

GetParentBlock returns the ParentBlock field if non-nil, zero value otherwise.

### GetParentBlockOk

`func (o *IPAssignment) GetParentBlockOk() (*ParentBlock, bool)`

GetParentBlockOk returns a tuple with the ParentBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBlock

`func (o *IPAssignment) SetParentBlock(v ParentBlock)`

SetParentBlock sets ParentBlock field to given value.

### HasParentBlock

`func (o *IPAssignment) HasParentBlock() bool`

HasParentBlock returns a boolean if a field has been set.

### GetPublic

`func (o *IPAssignment) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *IPAssignment) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *IPAssignment) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *IPAssignment) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetState

`func (o *IPAssignment) GetState() IPAssignmentState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IPAssignment) GetStateOk() (*IPAssignmentState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IPAssignment) SetState(v IPAssignmentState)`

SetState sets State field to given value.

### HasState

`func (o *IPAssignment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetNextHop

`func (o *IPAssignment) GetNextHop() string`

GetNextHop returns the NextHop field if non-nil, zero value otherwise.

### GetNextHopOk

`func (o *IPAssignment) GetNextHopOk() (*string, bool)`

GetNextHopOk returns a tuple with the NextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextHop

`func (o *IPAssignment) SetNextHop(v string)`

SetNextHop sets NextHop field to given value.

### HasNextHop

`func (o *IPAssignment) HasNextHop() bool`

HasNextHop returns a boolean if a field has been set.

### GetType

`func (o *IPAssignment) GetType() IPAssignmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPAssignment) GetTypeOk() (*IPAssignmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPAssignment) SetType(v IPAssignmentType)`

SetType sets Type field to given value.

### HasType

`func (o *IPAssignment) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


