# SecondaryDeviceDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeactivationKey** | Pointer to **string** | Object that holds the secondary deactivation key for a redundant device. If you do not provide a deactivation key for a Palo Alto BYOL device, you must contact your vendor to release the license. | [optional] 

## Methods

### NewSecondaryDeviceDeleteRequest

`func NewSecondaryDeviceDeleteRequest() *SecondaryDeviceDeleteRequest`

NewSecondaryDeviceDeleteRequest instantiates a new SecondaryDeviceDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryDeviceDeleteRequestWithDefaults

`func NewSecondaryDeviceDeleteRequestWithDefaults() *SecondaryDeviceDeleteRequest`

NewSecondaryDeviceDeleteRequestWithDefaults instantiates a new SecondaryDeviceDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeactivationKey

`func (o *SecondaryDeviceDeleteRequest) GetDeactivationKey() string`

GetDeactivationKey returns the DeactivationKey field if non-nil, zero value otherwise.

### GetDeactivationKeyOk

`func (o *SecondaryDeviceDeleteRequest) GetDeactivationKeyOk() (*string, bool)`

GetDeactivationKeyOk returns a tuple with the DeactivationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivationKey

`func (o *SecondaryDeviceDeleteRequest) SetDeactivationKey(v string)`

SetDeactivationKey sets DeactivationKey field to given value.

### HasDeactivationKey

`func (o *SecondaryDeviceDeleteRequest) HasDeactivationKey() bool`

HasDeactivationKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


