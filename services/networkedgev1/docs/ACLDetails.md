# ACLDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceType** | Pointer to **string** | Interface type, whether MGMT or WAN. | [optional] 
**Uuid** | Pointer to **string** | The unique ID of ACL. | [optional] 

## Methods

### NewACLDetails

`func NewACLDetails() *ACLDetails`

NewACLDetails instantiates a new ACLDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewACLDetailsWithDefaults

`func NewACLDetailsWithDefaults() *ACLDetails`

NewACLDetailsWithDefaults instantiates a new ACLDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceType

`func (o *ACLDetails) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *ACLDetails) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *ACLDetails) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *ACLDetails) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetUuid

`func (o *ACLDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ACLDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ACLDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ACLDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


