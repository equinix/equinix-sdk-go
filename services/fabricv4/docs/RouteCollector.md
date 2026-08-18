# RouteCollector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asn** | Pointer to **int64** | Route Collector ASN | [optional] 
**IpV4** | Pointer to **string** | Route Collector ipV4 address | [optional] 
**IpV6** | Pointer to **string** | Route Collector ipV6 address | [optional] 

## Methods

### NewRouteCollector

`func NewRouteCollector() *RouteCollector`

NewRouteCollector instantiates a new RouteCollector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteCollectorWithDefaults

`func NewRouteCollectorWithDefaults() *RouteCollector`

NewRouteCollectorWithDefaults instantiates a new RouteCollector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsn

`func (o *RouteCollector) GetAsn() int64`

GetAsn returns the Asn field if non-nil, zero value otherwise.

### GetAsnOk

`func (o *RouteCollector) GetAsnOk() (*int64, bool)`

GetAsnOk returns a tuple with the Asn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsn

`func (o *RouteCollector) SetAsn(v int64)`

SetAsn sets Asn field to given value.

### HasAsn

`func (o *RouteCollector) HasAsn() bool`

HasAsn returns a boolean if a field has been set.

### GetIpV4

`func (o *RouteCollector) GetIpV4() string`

GetIpV4 returns the IpV4 field if non-nil, zero value otherwise.

### GetIpV4Ok

`func (o *RouteCollector) GetIpV4Ok() (*string, bool)`

GetIpV4Ok returns a tuple with the IpV4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV4

`func (o *RouteCollector) SetIpV4(v string)`

SetIpV4 sets IpV4 field to given value.

### HasIpV4

`func (o *RouteCollector) HasIpV4() bool`

HasIpV4 returns a boolean if a field has been set.

### GetIpV6

`func (o *RouteCollector) GetIpV6() string`

GetIpV6 returns the IpV6 field if non-nil, zero value otherwise.

### GetIpV6Ok

`func (o *RouteCollector) GetIpV6Ok() (*string, bool)`

GetIpV6Ok returns a tuple with the IpV6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6

`func (o *RouteCollector) SetIpV6(v string)`

SetIpV6 sets IpV6 field to given value.

### HasIpV6

`func (o *RouteCollector) HasIpV6() bool`

HasIpV6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


