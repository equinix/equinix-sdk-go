# VirtualDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Virtual Device URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned Virtual Device identifier | [optional] 
**Name** | Pointer to **string** | Customer-assigned Virtual Device name | [optional] 
**Type** | Pointer to [**VirtualDeviceType**](VirtualDeviceType.md) |  | [optional] 
**Account** | Pointer to [**SimplifiedAccount**](SimplifiedAccount.md) |  | [optional] 

## Methods

### NewVirtualDevice

`func NewVirtualDevice() *VirtualDevice`

NewVirtualDevice instantiates a new VirtualDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceWithDefaults

`func NewVirtualDeviceWithDefaults() *VirtualDevice`

NewVirtualDeviceWithDefaults instantiates a new VirtualDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VirtualDevice) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VirtualDevice) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VirtualDevice) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VirtualDevice) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualDevice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualDevice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualDevice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *VirtualDevice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *VirtualDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *VirtualDevice) GetType() VirtualDeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualDevice) GetTypeOk() (*VirtualDeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualDevice) SetType(v VirtualDeviceType)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccount

`func (o *VirtualDevice) GetAccount() SimplifiedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *VirtualDevice) GetAccountOk() (*SimplifiedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *VirtualDevice) SetAccount(v SimplifiedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *VirtualDevice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


