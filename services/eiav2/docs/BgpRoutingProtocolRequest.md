# BgpRoutingProtocolRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | Pointer to **[]string** |  | [optional] 
**Type** | [**RoutingProtocolType**](RoutingProtocolType.md) |  | 
**Name** | Pointer to **string** | Name of the routing protocol instance.  | [optional] 
**Description** | Pointer to **string** | Description of the routing protocol instance  | [optional] 
**CustomerAsnRange** | Pointer to [**BgpRoutingProtocolRequestAllOfCustomerAsnRange**](BgpRoutingProtocolRequestAllOfCustomerAsnRange.md) |  | [optional] 
**CustomerAsn** | Pointer to **int64** | Customer Autonomous System Number  | [optional] 
**BgpAuthKey** | Pointer to **string** | BGP authentication key  | [optional] 
**ExportPolicy** | [**BgpRoutingProtocolRequestAllOfExportPolicy**](BgpRoutingProtocolRequestAllOfExportPolicy.md) |  | 
**Ipv4** | Pointer to [**RoutingProtocolIpv4Request**](RoutingProtocolIpv4Request.md) |  | [optional] 
**Ipv6** | Pointer to [**RoutingProtocolIpv6Request**](RoutingProtocolIpv6Request.md) |  | [optional] 

## Methods

### NewBgpRoutingProtocolRequest

`func NewBgpRoutingProtocolRequest(type_ RoutingProtocolType, exportPolicy BgpRoutingProtocolRequestAllOfExportPolicy, ) *BgpRoutingProtocolRequest`

NewBgpRoutingProtocolRequest instantiates a new BgpRoutingProtocolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBgpRoutingProtocolRequestWithDefaults

`func NewBgpRoutingProtocolRequestWithDefaults() *BgpRoutingProtocolRequest`

NewBgpRoutingProtocolRequestWithDefaults instantiates a new BgpRoutingProtocolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *BgpRoutingProtocolRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BgpRoutingProtocolRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BgpRoutingProtocolRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BgpRoutingProtocolRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *BgpRoutingProtocolRequest) GetType() RoutingProtocolType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BgpRoutingProtocolRequest) GetTypeOk() (*RoutingProtocolType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BgpRoutingProtocolRequest) SetType(v RoutingProtocolType)`

SetType sets Type field to given value.


### GetName

`func (o *BgpRoutingProtocolRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BgpRoutingProtocolRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BgpRoutingProtocolRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BgpRoutingProtocolRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *BgpRoutingProtocolRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BgpRoutingProtocolRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BgpRoutingProtocolRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BgpRoutingProtocolRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

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


