# StaticRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**RoutingProtocolType**](RoutingProtocolType.md) |  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewStaticRoutingProtocolRequest

`func NewStaticRoutingProtocolRequest(type_ RoutingProtocolType, ) *StaticRoutingProtocolRequest`

NewStaticRoutingProtocolRequest instantiates a new StaticRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingProtocolRequestWithDefaults

`func NewStaticRoutingProtocolRequestWithDefaults() *StaticRoutingProtocolRequest`

NewStaticRoutingProtocolRequestWithDefaults instantiates a new StaticRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *StaticRoutingProtocolRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StaticRoutingProtocolRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StaticRoutingProtocolRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StaticRoutingProtocolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *StaticRoutingProtocolRequest) GetType() RoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StaticRoutingProtocolRequest) GetTypeOk() (*RoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StaticRoutingProtocolRequest) SetType(v RoutingProtocolType)`

SetType sets Type field to given value.


### GetName

`func (o *StaticRoutingProtocolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticRoutingProtocolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticRoutingProtocolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StaticRoutingProtocolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StaticRoutingProtocolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StaticRoutingProtocolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StaticRoutingProtocolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StaticRoutingProtocolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *StaticRoutingProtocolRequest) GetIpv4() RoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *StaticRoutingProtocolRequest) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *StaticRoutingProtocolRequest) SetIpv4(v RoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *StaticRoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *StaticRoutingProtocolRequest) GetIpv6() RoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *StaticRoutingProtocolRequest) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *StaticRoutingProtocolRequest) SetIpv6(v RoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *StaticRoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


