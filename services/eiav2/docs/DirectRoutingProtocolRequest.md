# DirectRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**RoutingProtocolType**](RoutingProtocolType.md) |  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**Ipv4** | Pointer to [**DirectRoutingProtocolIpv4Request**](DirectRoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**DirectRoutingProtocolIpv6Request**](DirectRoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewDirectRoutingProtocolRequest

`func NewDirectRoutingProtocolRequest(type_ RoutingProtocolType, ) *DirectRoutingProtocolRequest`

NewDirectRoutingProtocolRequest instantiates a new DirectRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectRoutingProtocolRequestWithDefaults

`func NewDirectRoutingProtocolRequestWithDefaults() *DirectRoutingProtocolRequest`

NewDirectRoutingProtocolRequestWithDefaults instantiates a new DirectRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *DirectRoutingProtocolRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DirectRoutingProtocolRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DirectRoutingProtocolRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DirectRoutingProtocolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *DirectRoutingProtocolRequest) GetType() RoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DirectRoutingProtocolRequest) GetTypeOk() (*RoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DirectRoutingProtocolRequest) SetType(v RoutingProtocolType)`

SetType sets Type field to given value.


### GetName

`func (o *DirectRoutingProtocolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DirectRoutingProtocolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DirectRoutingProtocolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DirectRoutingProtocolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DirectRoutingProtocolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DirectRoutingProtocolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DirectRoutingProtocolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DirectRoutingProtocolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIpv4

`func (o *DirectRoutingProtocolRequest) GetIpv4() DirectRoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *DirectRoutingProtocolRequest) GetIpv4Ok() (*DirectRoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *DirectRoutingProtocolRequest) SetIpv4(v DirectRoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *DirectRoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *DirectRoutingProtocolRequest) GetIpv6() DirectRoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *DirectRoutingProtocolRequest) GetIpv6Ok() (*DirectRoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *DirectRoutingProtocolRequest) SetIpv6(v DirectRoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *DirectRoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


