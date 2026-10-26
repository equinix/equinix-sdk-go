# ExchangeServiceReplaceOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | [**ExchangeServiceOpEnum**](ExchangeServiceOpEnum.md) |  | 
**Path** | **string** | ### A JSON Pointer path.   * &#x60;/name&#x60; - Name of the Internet Exchange Service   * &#x60;/description&#x60; - Description of the Internet Exchange Service   * &#x60;routingProtocol/macAddress&#x60; - MAC Address   * &#x60;routingProtocol/bgpIpv4/domainName&#x60; - Domain Name for BGP IPv4   * &#x60;routingProtocol/bgpIpv4/asSet&#x60; - AS Set for BGP IPv4   * &#x60;routingProtocol/bgpIpv4/md5AuthKey&#x60; - MD5 Auth Key for BGP IPv4   * &#x60;routingProtocol/bgpIpv4/rcMd5AuthKey&#x60; - RC MD5 Auth Key for BGP IPv4   * &#x60;routingProtocol/bgpIpv4/ipPrefixes&#x60; - IP Prefixes for BGP IPv4   * &#x60;routingProtocol/bgpIpv6/mlpeEnabled&#x60; - MLPE Enabled for BGP IPv4   * &#x60;routingProtocol/bgpIpv6/domainName&#x60;- Domain Name for BGP IPv6   * &#x60;routingProtocol/bgpIpv6/asSet&#x60;- AS Set for BGP IPv6   * &#x60;routingProtocol/bgpIpv6/md5AuthKey&#x60; - MD5 Auth Key for BGP IPv6   * &#x60;routingProtocol/bgpIpv6/rcMd5AuthKey&#x60; - RC MD5 Auth Key for BGP IPv6   * &#x60;routingProtocol/bgpIpv6/ipPrefixes&#x60; - IP Prefixes for BGP IPv6   * &#x60;routingProtocol/bgpIpv6/mlpeEnabled&#x60; - MLPE Enabled for BGP IPv6  | 
**Value** | **interface{}** | value to replace with | 

## Methods

### NewExchangeServiceReplaceOperation

`func NewExchangeServiceReplaceOperation(op ExchangeServiceOpEnum, path string, value interface{}, ) *ExchangeServiceReplaceOperation`

NewExchangeServiceReplaceOperation instantiates a new ExchangeServiceReplaceOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceReplaceOperationWithDefaults

`func NewExchangeServiceReplaceOperationWithDefaults() *ExchangeServiceReplaceOperation`

NewExchangeServiceReplaceOperationWithDefaults instantiates a new ExchangeServiceReplaceOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ExchangeServiceReplaceOperation) GetOp() ExchangeServiceOpEnum`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ExchangeServiceReplaceOperation) GetOpOk() (*ExchangeServiceOpEnum, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ExchangeServiceReplaceOperation) SetOp(v ExchangeServiceOpEnum)`

SetOp sets Op field to given value.


### GetPath

`func (o *ExchangeServiceReplaceOperation) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ExchangeServiceReplaceOperation) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ExchangeServiceReplaceOperation) SetPath(v string)`

SetPath sets Path field to given value.


### GetValue

`func (o *ExchangeServiceReplaceOperation) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ExchangeServiceReplaceOperation) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ExchangeServiceReplaceOperation) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *ExchangeServiceReplaceOperation) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ExchangeServiceReplaceOperation) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


