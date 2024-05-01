# CustomerIpBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrefixLength** | **int32** | The IPv6 routing prefix, sometimes called a subnet mask, is the network portion of an IP address. The prefix length is an integer between 1 and 128 (inclusive) that represents the number of bits set to 1, such as /24 or /60. | 

## Methods

### NewCustomerIpBlock

`func NewCustomerIpBlock(prefixLength int32, ) *CustomerIpBlock`

NewCustomerIpBlock instantiates a new CustomerIpBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerIpBlockWithDefaults

`func NewCustomerIpBlockWithDefaults() *CustomerIpBlock`

NewCustomerIpBlockWithDefaults instantiates a new CustomerIpBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefixLength

`func (o *CustomerIpBlock) GetPrefixLength() int32`

GetPrefixLength returns the PrefixLength field if non-nil, zero value otherwise.

### GetPrefixLengthOk

`func (o *CustomerIpBlock) GetPrefixLengthOk() (*int32, bool)`

GetPrefixLengthOk returns a tuple with the PrefixLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLength

`func (o *CustomerIpBlock) SetPrefixLength(v int32)`

SetPrefixLength sets PrefixLength field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


