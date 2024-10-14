# VirtualDeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VirtualDeviceInterfaceType**](VirtualDeviceInterfaceType.md) |  | [optional] 
**Id** | Pointer to **int32** | Network Edge assigned identifier | [optional] 
**Uuid** | Pointer to **string** | Interface identifier | [optional] 

## Methods

### NewVirtualDeviceInterface

`func NewVirtualDeviceInterface() *VirtualDeviceInterface`

NewVirtualDeviceInterface instantiates a new VirtualDeviceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceInterfaceWithDefaults

`func NewVirtualDeviceInterfaceWithDefaults() *VirtualDeviceInterface`

NewVirtualDeviceInterfaceWithDefaults instantiates a new VirtualDeviceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VirtualDeviceInterface) GetType() VirtualDeviceInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualDeviceInterface) GetTypeOk() (*VirtualDeviceInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualDeviceInterface) SetType(v VirtualDeviceInterfaceType)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualDeviceInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *VirtualDeviceInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualDeviceInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualDeviceInterface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualDeviceInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualDeviceInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualDeviceInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualDeviceInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualDeviceInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


