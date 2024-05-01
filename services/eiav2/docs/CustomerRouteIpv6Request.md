# CustomerRouteIpv6Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | Pointer to [**IpBlockIpv6Request**](IpBlockIpv6Request.md) |  | [optional] 
**Prefix** | Pointer to **string** | Subnet prefix  | [optional] 

## Methods

### NewCustomerRouteIpv6Request

`func NewCustomerRouteIpv6Request() *CustomerRouteIpv6Request`

NewCustomerRouteIpv6Request instantiates a new CustomerRouteIpv6Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRouteIpv6RequestWithDefaults

`func NewCustomerRouteIpv6RequestWithDefaults() *CustomerRouteIpv6Request`

NewCustomerRouteIpv6RequestWithDefaults instantiates a new CustomerRouteIpv6Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpBlock

`func (o *CustomerRouteIpv6Request) GetIpBlock() IpBlockIpv6Request`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *CustomerRouteIpv6Request) GetIpBlockOk() (*IpBlockIpv6Request, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *CustomerRouteIpv6Request) SetIpBlock(v IpBlockIpv6Request)`

SetIpBlock sets IpBlock field to given value.

### HasIpBlock

`func (o *CustomerRouteIpv6Request) HasIpBlock() bool`

HasIpBlock returns a boolean if a field has been set.

### GetPrefix

`func (o *CustomerRouteIpv6Request) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *CustomerRouteIpv6Request) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *CustomerRouteIpv6Request) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *CustomerRouteIpv6Request) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


