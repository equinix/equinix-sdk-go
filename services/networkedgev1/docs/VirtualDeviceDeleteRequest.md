# VirtualDeviceDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeactivationKey** | Pointer to **string** | If you do not provide a deactivation key for a Palo Alto BYOL device, you must contact your vendor to release the license. | [optional] 
**Secondary** | Pointer to [**SecondaryDeviceDeleteRequest**](SecondaryDeviceDeleteRequest.md) |  | [optional] 

## Methods

### NewVirtualDeviceDeleteRequest

`func NewVirtualDeviceDeleteRequest() *VirtualDeviceDeleteRequest`

NewVirtualDeviceDeleteRequest instantiates a new VirtualDeviceDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualDeviceDeleteRequestWithDefaults

`func NewVirtualDeviceDeleteRequestWithDefaults() *VirtualDeviceDeleteRequest`

NewVirtualDeviceDeleteRequestWithDefaults instantiates a new VirtualDeviceDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeactivationKey

`func (o *VirtualDeviceDeleteRequest) GetDeactivationKey() string`

GetDeactivationKey returns the DeactivationKey field if non-nil, zero value otherwise.

### GetDeactivationKeyOk

`func (o *VirtualDeviceDeleteRequest) GetDeactivationKeyOk() (*string, bool)`

GetDeactivationKeyOk returns a tuple with the DeactivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationKey

`func (o *VirtualDeviceDeleteRequest) SetDeactivationKey(v string)`

SetDeactivationKey sets DeactivationKey field to given value.

### HasDeactivationKey

`func (o *VirtualDeviceDeleteRequest) HasDeactivationKey() bool`

HasDeactivationKey returns a boolean if a field has been set.

### GetSecondary

`func (o *VirtualDeviceDeleteRequest) GetSecondary() SecondaryDeviceDeleteRequest`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *VirtualDeviceDeleteRequest) GetSecondaryOk() (*SecondaryDeviceDeleteRequest, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *VirtualDeviceDeleteRequest) SetSecondary(v SecondaryDeviceDeleteRequest)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *VirtualDeviceDeleteRequest) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


