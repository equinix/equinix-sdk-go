# VirtualDeviceACLDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the virtual device associated with this template. | [optional] 
**Uuid** | Pointer to **string** | The unique Id of the virtual device associated with this template. | [optional] 
**InterfaceType** | Pointer to **string** | Interface type, WAN or MGMT. | [optional] 
**AclStatus** | Pointer to **string** | Device ACL status | [optional] 

## Methods

### NewVirtualDeviceACLDetails

`func NewVirtualDeviceACLDetails() *VirtualDeviceACLDetails`

NewVirtualDeviceACLDetails instantiates a new VirtualDeviceACLDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceACLDetailsWithDefaults

`func NewVirtualDeviceACLDetailsWithDefaults() *VirtualDeviceACLDetails`

NewVirtualDeviceACLDetailsWithDefaults instantiates a new VirtualDeviceACLDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualDeviceACLDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDeviceACLDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDeviceACLDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualDeviceACLDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualDeviceACLDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualDeviceACLDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualDeviceACLDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualDeviceACLDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetInterfaceType

`func (o *VirtualDeviceACLDetails) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *VirtualDeviceACLDetails) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *VirtualDeviceACLDetails) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *VirtualDeviceACLDetails) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetAclStatus

`func (o *VirtualDeviceACLDetails) GetAclStatus() string`

GetAclStatus returns the AclStatus field if non-nil, zero value otherwise.

### GetAclStatusOk

`func (o *VirtualDeviceACLDetails) GetAclStatusOk() (*string, bool)`

GetAclStatusOk returns a tuple with the AclStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclStatus

`func (o *VirtualDeviceACLDetails) SetAclStatus(v string)`

SetAclStatus sets AclStatus field to given value.

### HasAclStatus

`func (o *VirtualDeviceACLDetails) HasAclStatus() bool`

HasAclStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


