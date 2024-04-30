# RoutingProtocolCustomerRouteIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportPolicy** | **map[string]interface{}** |  | 
**Prefix** | Pointer to **string** | Subnet prefix  | [optional] 
**PrefixLength** | **int32** | Determines the size of ip subnet | 
**IpBlock** | Pointer to [**IpBlockReadModel**](IpBlockReadModel.md) |  | [optional] 

## Methods

### NewRoutingProtocolCustomerRouteIpv6

`func NewRoutingProtocolCustomerRouteIpv6(importPolicy map[string]interface{}, prefixLength int32, ) *RoutingProtocolCustomerRouteIpv6`

NewRoutingProtocolCustomerRouteIpv6 instantiates a new RoutingProtocolCustomerRouteIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingProtocolCustomerRouteIpv6WithDefaults

`func NewRoutingProtocolCustomerRouteIpv6WithDefaults() *RoutingProtocolCustomerRouteIpv6`

NewRoutingProtocolCustomerRouteIpv6WithDefaults instantiates a new RoutingProtocolCustomerRouteIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportPolicy

`func (o *RoutingProtocolCustomerRouteIpv6) GetImportPolicy() map[string]interface{}`

GetImportPolicy returns the ImportPolicy field if non-nil, zero value otherwise.

### GetImportPolicyOk

`func (o *RoutingProtocolCustomerRouteIpv6) GetImportPolicyOk() (*map[string]interface{}, bool)`

GetImportPolicyOk returns a tuple with the ImportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPolicy

`func (o *RoutingProtocolCustomerRouteIpv6) SetImportPolicy(v map[string]interface{})`

SetImportPolicy sets ImportPolicy field to given value.


### GetPrefix

`func (o *RoutingProtocolCustomerRouteIpv6) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *RoutingProtocolCustomerRouteIpv6) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *RoutingProtocolCustomerRouteIpv6) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *RoutingProtocolCustomerRouteIpv6) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrefixLength

`func (o *RoutingProtocolCustomerRouteIpv6) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *RoutingProtocolCustomerRouteIpv6) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *RoutingProtocolCustomerRouteIpv6) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.


### GetIpBlock

`func (o *RoutingProtocolCustomerRouteIpv6) GetIpBlock() IpBlockReadModel`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *RoutingProtocolCustomerRouteIpv6) GetIpBlockOk() (*IpBlockReadModel, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *RoutingProtocolCustomerRouteIpv6) SetIpBlock(v IpBlockReadModel)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *RoutingProtocolCustomerRouteIpv6) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


