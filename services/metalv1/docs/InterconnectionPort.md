# InterconnectionPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Organization** | Pointer to [**Href**](Href.md) |  | [optional] 
**Role** | Pointer to [**InterconnectionPortRole**](InterconnectionPortRole.md) |  | [optional] 
**Status** | Pointer to [**InterconnectionPortStatus**](InterconnectionPortStatus.md) |  | [optional] 
**SwitchId** | Pointer to **string** | A switch &#39;short ID&#39; | [optional] 
**VirtualCircuits** | Pointer to [**[]VirtualCircuit**](VirtualCircuit.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **int32** |  | [optional] 
**LinkStatus** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 

## Methods

### NewInterconnectionPort

`func NewInterconnectionPort() *InterconnectionPort`

NewInterconnectionPort instantiates a new InterconnectionPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionPortWithDefaults

`func NewInterconnectionPortWithDefaults() *InterconnectionPort`

NewInterconnectionPortWithDefaults instantiates a new InterconnectionPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InterconnectionPort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InterconnectionPort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InterconnectionPort) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InterconnectionPort) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *InterconnectionPort) GetOrganization() Href`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *InterconnectionPort) GetOrganizationOk() (*Href, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *InterconnectionPort) SetOrganization(v Href)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *InterconnectionPort) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRole

`func (o *InterconnectionPort) GetRole() InterconnectionPortRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *InterconnectionPort) GetRoleOk() (*InterconnectionPortRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *InterconnectionPort) SetRole(v InterconnectionPortRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *InterconnectionPort) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *InterconnectionPort) GetStatus() InterconnectionPortStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InterconnectionPort) GetStatusOk() (*InterconnectionPortStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InterconnectionPort) SetStatus(v InterconnectionPortStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InterconnectionPort) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSwitchId

`func (o *InterconnectionPort) GetSwitchId() string`

GetSwitchId returns the SwitchId field if non-nil, zero value otherwise.

### GetSwitchIdOk

`func (o *InterconnectionPort) GetSwitchIdOk() (*string, bool)`

GetSwitchIdOk returns a tuple with the SwitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchId

`func (o *InterconnectionPort) SetSwitchId(v string)`

SetSwitchId sets SwitchId field to given value.

### HasSwitchId

`func (o *InterconnectionPort) HasSwitchId() bool`

HasSwitchId returns a boolean if a field has been set.

### GetVirtualCircuits

`func (o *InterconnectionPort) GetVirtualCircuits() []VirtualCircuit`

GetVirtualCircuits returns the VirtualCircuits field if non-nil, zero value otherwise.

### GetVirtualCircuitsOk

`func (o *InterconnectionPort) GetVirtualCircuitsOk() (*[]VirtualCircuit, bool)`

GetVirtualCircuitsOk returns a tuple with the VirtualCircuits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCircuits

`func (o *InterconnectionPort) SetVirtualCircuits(v []VirtualCircuit)`

SetVirtualCircuits sets VirtualCircuits field to given value.

### HasVirtualCircuits

`func (o *InterconnectionPort) HasVirtualCircuits() bool`

HasVirtualCircuits returns a boolean if a field has been set.

### GetName

`func (o *InterconnectionPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InterconnectionPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InterconnectionPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InterconnectionPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSpeed

`func (o *InterconnectionPort) GetSpeed() int32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *InterconnectionPort) GetSpeedOk() (*int32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *InterconnectionPort) SetSpeed(v int32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *InterconnectionPort) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetLinkStatus

`func (o *InterconnectionPort) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *InterconnectionPort) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *InterconnectionPort) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *InterconnectionPort) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetHref

`func (o *InterconnectionPort) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *InterconnectionPort) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *InterconnectionPort) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *InterconnectionPort) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


