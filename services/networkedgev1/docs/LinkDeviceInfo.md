# LinkDeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** | The ASN number of the device. The request will fail if you provide a new ASN for a device that already has an ASN. | [optional] 
**DeviceUuid** | **string** | device | 
**InterfaceId** | Pointer to **int32** | Any available interface of the device. | [optional] 

## Methods

### NewLinkDeviceInfo

`func NewLinkDeviceInfo(deviceUuid string, ) *LinkDeviceInfo`

NewLinkDeviceInfo instantiates a new LinkDeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkDeviceInfoWithDefaults

`func NewLinkDeviceInfoWithDefaults() *LinkDeviceInfo`

NewLinkDeviceInfoWithDefaults instantiates a new LinkDeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *LinkDeviceInfo) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *LinkDeviceInfo) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *LinkDeviceInfo) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *LinkDeviceInfo) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetDeviceUuid

`func (o *LinkDeviceInfo) GetDeviceUuid() string`

GetDeviceUuid returns the DeviceUuid field if non-nil, zero value otherwise.

### GetDeviceUuidOk

`func (o *LinkDeviceInfo) GetDeviceUuidOk() (*string, bool)`

GetDeviceUuidOk returns a tuple with the DeviceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceUuid

`func (o *LinkDeviceInfo) SetDeviceUuid(v string)`

SetDeviceUuid sets DeviceUuid field to given value.


### GetInterfaceId

`func (o *LinkDeviceInfo) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *LinkDeviceInfo) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *LinkDeviceInfo) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *LinkDeviceInfo) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


