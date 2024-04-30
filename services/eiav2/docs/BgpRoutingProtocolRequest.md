# BgpRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerAsnRange** | Pointer to [**BgpRoutingProtocolRequestAllOfCustomerAsnRange**](BgpRoutingProtocolRequestAllOfCustomerAsnRange.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer Autonomous System Number  | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**BgpRoutingProtocolRequestAllOfExportPolicy**](BgpRoutingProtocolRequestAllOfExportPolicy.md) |  | 
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewBgpRoutingProtocolRequest

`func NewBgpRoutingProtocolRequest(exportPolicy BgpRoutingProtocolRequestAllOfExportPolicy, ) *BgpRoutingProtocolRequest`

NewBgpRoutingProtocolRequest instantiates a new BgpRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpRoutingProtocolRequestWithDefaults

`func NewBgpRoutingProtocolRequestWithDefaults() *BgpRoutingProtocolRequest`

NewBgpRoutingProtocolRequestWithDefaults instantiates a new BgpRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerAsnRange

`func (o *BgpRoutingProtocolRequest) GetCustomerAsnRange() BgpRoutingProtocolRequestAllOfCustomerAsnRange`

GetCustomerAsnRange returns the CustomerAsnRange field if non-nil, zero value otherwise.

### GetCustomerAsnRangeOk

`func (o *BgpRoutingProtocolRequest) GetCustomerAsnRangeOk() (*BgpRoutingProtocolRequestAllOfCustomerAsnRange, bool)`

GetCustomerAsnRangeOk returns a tuple with the CustomerAsnRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsnRange

`func (o *BgpRoutingProtocolRequest) SetCustomerAsnRange(v BgpRoutingProtocolRequestAllOfCustomerAsnRange)`

SetCustomerAsnRange sets CustomerAsnRange field to given value.

### HasCustomerAsnRange

`func (o *BgpRoutingProtocolRequest) HasCustomerAsnRange() bool`

HasCustomerAsnRange returns a boolean if a field has been set.

### GetCustomerAsn

`func (o *BgpRoutingProtocolRequest) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *BgpRoutingProtocolRequest) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *BgpRoutingProtocolRequest) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.

### HasCustomerAsn

`func (o *BgpRoutingProtocolRequest) HasCustomerAsn() bool`

HasCustomerAsn returns a boolean if a field has been set.

### GetBgpAuthKey

`func (o *BgpRoutingProtocolRequest) GetBgpAuthKey() string`

GetBgpAuthKey returns the BgpAuthKey field if non-nil, zero value otherwise.

### GetBgpAuthKeyOk

`func (o *BgpRoutingProtocolRequest) GetBgpAuthKeyOk() (*string, bool)`

GetBgpAuthKeyOk returns a tuple with the BgpAuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAuthKey

`func (o *BgpRoutingProtocolRequest) SetBgpAuthKey(v string)`

SetBgpAuthKey sets BgpAuthKey field to given value.

### HasBgpAuthKey

`func (o *BgpRoutingProtocolRequest) HasBgpAuthKey() bool`

HasBgpAuthKey returns a boolean if a field has been set.

### GetExportPolicy

`func (o *BgpRoutingProtocolRequest) GetExportPolicy() BgpRoutingProtocolRequestAllOfExportPolicy`

GetExportPolicy returns the ExportPolicy field if non-nil, zero value otherwise.

### GetExportPolicyOk

`func (o *BgpRoutingProtocolRequest) GetExportPolicyOk() (*BgpRoutingProtocolRequestAllOfExportPolicy, bool)`

GetExportPolicyOk returns a tuple with the ExportPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportPolicy

`func (o *BgpRoutingProtocolRequest) SetExportPolicy(v BgpRoutingProtocolRequestAllOfExportPolicy)`

SetExportPolicy sets ExportPolicy field to given value.


### GetIpv4

`func (o *BgpRoutingProtocolRequest) GetIpv4() RoutingProtocolIpv4Request`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *BgpRoutingProtocolRequest) GetIpv4Ok() (*RoutingProtocolIpv4Request, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *BgpRoutingProtocolRequest) SetIpv4(v RoutingProtocolIpv4Request)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *BgpRoutingProtocolRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *BgpRoutingProtocolRequest) GetIpv6() RoutingProtocolIpv6Request`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *BgpRoutingProtocolRequest) GetIpv6Ok() (*RoutingProtocolIpv6Request, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *BgpRoutingProtocolRequest) SetIpv6(v RoutingProtocolIpv6Request)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *BgpRoutingProtocolRequest) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


