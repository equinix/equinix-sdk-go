# CustomerRoute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpBlock** | [**CustomerIpBlock**](CustomerIpBlock.md) |  | 

## Methods

### NewCustomerRoute

`func NewCustomerRoute(ipBlock CustomerIpBlock, ) *CustomerRoute`

NewCustomerRoute instantiates a new CustomerRoute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRouteWithDefaults

`func NewCustomerRouteWithDefaults() *CustomerRoute`

NewCustomerRouteWithDefaults instantiates a new CustomerRoute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpBlock

`func (o *CustomerRoute) GetIpBlock() CustomerIpBlock`

GetIpBlock returns the IpBlock field if non-nil, zero value otherwise.

### GetIpBlockOk

`func (o *CustomerRoute) GetIpBlockOk() (*CustomerIpBlock, bool)`

GetIpBlockOk returns a tuple with the IpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpBlock

`func (o *CustomerRoute) SetIpBlock(v CustomerIpBlock)`

SetIpBlock sets IpBlock field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


