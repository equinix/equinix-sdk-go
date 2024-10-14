# Ipv4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **string** | Primary Timing Server IP Address | 
**Secondary** | **string** | Secondary Timing Server IP Address | 
**NetworkMask** | **string** | Network Mask | 
**DefaultGateway** | Pointer to **string** | Gateway Interface IP address | [optional] 

## Methods

### NewIpv4

`func NewIpv4(primary string, secondary string, networkMask string, ) *Ipv4`

NewIpv4 instantiates a new Ipv4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4WithDefaults

`func NewIpv4WithDefaults() *Ipv4`

NewIpv4WithDefaults instantiates a new Ipv4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *Ipv4) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Ipv4) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Ipv4) SetPrimary(v string)`

SetPrimary sets Primary field to given value.


### GetSecondary

`func (o *Ipv4) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *Ipv4) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *Ipv4) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.


### GetNetworkMask

`func (o *Ipv4) GetNetworkMask() string`

GetNetworkMask returns the NetworkMask field if non-nil, zero value otherwise.

### GetNetworkMaskOk

`func (o *Ipv4) GetNetworkMaskOk() (*string, bool)`

GetNetworkMaskOk returns a tuple with the NetworkMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMask

`func (o *Ipv4) SetNetworkMask(v string)`

SetNetworkMask sets NetworkMask field to given value.


### GetDefaultGateway

`func (o *Ipv4) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *Ipv4) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *Ipv4) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *Ipv4) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


