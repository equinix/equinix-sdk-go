# LinkProtocolIpv4Ipv6Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkPrefix** | Pointer to **string** | Link subnet prefix | [optional] 
**LocalIfaceIp** | Pointer to **string** | Prefix datatype when linkPrefix not specified | [optional] 
**RemoteIfaceIp** | Pointer to **string** | Equinix-side link interface address | [optional] 

## Methods

### NewLinkProtocolIpv4Ipv6Config

`func NewLinkProtocolIpv4Ipv6Config() *LinkProtocolIpv4Ipv6Config`

NewLinkProtocolIpv4Ipv6Config instantiates a new LinkProtocolIpv4Ipv6Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkProtocolIpv4Ipv6ConfigWithDefaults

`func NewLinkProtocolIpv4Ipv6ConfigWithDefaults() *LinkProtocolIpv4Ipv6Config`

NewLinkProtocolIpv4Ipv6ConfigWithDefaults instantiates a new LinkProtocolIpv4Ipv6Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkPrefix

`func (o *LinkProtocolIpv4Ipv6Config) GetLinkPrefix() string`

GetLinkPrefix returns the LinkPrefix field if non-nil, zero value otherwise.

### GetLinkPrefixOk

`func (o *LinkProtocolIpv4Ipv6Config) GetLinkPrefixOk() (*string, bool)`

GetLinkPrefixOk returns a tuple with the LinkPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkPrefix

`func (o *LinkProtocolIpv4Ipv6Config) SetLinkPrefix(v string)`

SetLinkPrefix sets LinkPrefix field to given value.

### HasLinkPrefix

`func (o *LinkProtocolIpv4Ipv6Config) HasLinkPrefix() bool`

HasLinkPrefix returns a boolean if a field has been set.

### GetLocalIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) GetLocalIfaceIp() string`

GetLocalIfaceIp returns the LocalIfaceIp field if non-nil, zero value otherwise.

### GetLocalIfaceIpOk

`func (o *LinkProtocolIpv4Ipv6Config) GetLocalIfaceIpOk() (*string, bool)`

GetLocalIfaceIpOk returns a tuple with the LocalIfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) SetLocalIfaceIp(v string)`

SetLocalIfaceIp sets LocalIfaceIp field to given value.

### HasLocalIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) HasLocalIfaceIp() bool`

HasLocalIfaceIp returns a boolean if a field has been set.

### GetRemoteIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) GetRemoteIfaceIp() string`

GetRemoteIfaceIp returns the RemoteIfaceIp field if non-nil, zero value otherwise.

### GetRemoteIfaceIpOk

`func (o *LinkProtocolIpv4Ipv6Config) GetRemoteIfaceIpOk() (*string, bool)`

GetRemoteIfaceIpOk returns a tuple with the RemoteIfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) SetRemoteIfaceIp(v string)`

SetRemoteIfaceIp sets RemoteIfaceIp field to given value.

### HasRemoteIfaceIp

`func (o *LinkProtocolIpv4Ipv6Config) HasRemoteIfaceIp() bool`

HasRemoteIfaceIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


