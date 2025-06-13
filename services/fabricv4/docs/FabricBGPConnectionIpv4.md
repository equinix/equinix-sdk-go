# FabricBGPConnectionIpv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerIp** | Pointer to **string** | Customer side peering ip | [optional] 
**EquinixIp** | **string** | Equinix side peering ip | 

## Methods

### NewFabricBGPConnectionIpv4

`func NewFabricBGPConnectionIpv4(equinixIp string, ) *FabricBGPConnectionIpv4`

NewFabricBGPConnectionIpv4 instantiates a new FabricBGPConnectionIpv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricBGPConnectionIpv4WithDefaults

`func NewFabricBGPConnectionIpv4WithDefaults() *FabricBGPConnectionIpv4`

NewFabricBGPConnectionIpv4WithDefaults instantiates a new FabricBGPConnectionIpv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerIp

`func (o *FabricBGPConnectionIpv4) GetCustomerIp() string`

GetCustomerIp returns the CustomerIp field if non-nil, zero value otherwise.

### GetCustomerIpOk

`func (o *FabricBGPConnectionIpv4) GetCustomerIpOk() (*string, bool)`

GetCustomerIpOk returns a tuple with the CustomerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIp

`func (o *FabricBGPConnectionIpv4) SetCustomerIp(v string)`

SetCustomerIp sets CustomerIp field to given value.

### HasCustomerIp

`func (o *FabricBGPConnectionIpv4) HasCustomerIp() bool`

HasCustomerIp returns a boolean if a field has been set.

### GetEquinixIp

`func (o *FabricBGPConnectionIpv4) GetEquinixIp() string`

GetEquinixIp returns the EquinixIp field if non-nil, zero value otherwise.

### GetEquinixIpOk

`func (o *FabricBGPConnectionIpv4) GetEquinixIpOk() (*string, bool)`

GetEquinixIpOk returns a tuple with the EquinixIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixIp

`func (o *FabricBGPConnectionIpv4) SetEquinixIp(v string)`

SetEquinixIp sets EquinixIp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


