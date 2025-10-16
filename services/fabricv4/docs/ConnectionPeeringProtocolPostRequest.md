# ConnectionPeeringProtocolPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ConnectionPeeringProtocolPostRequestType**](ConnectionPeeringProtocolPostRequestType.md) |  | 
**Name** | **string** | Customer-provided peering protocol name | 
**Description** | **string** | Customer-provided peering protocol description | 
**CustomerAsn** | **int64** | Customer ASN | 
**MacAddress** | **string** | MAC address of the peering protocol | 
**BgpIpv4** | [**PeeringConnectionIpv4**](PeeringConnectionIpv4.md) |  | 
**BgpIpv6** | [**PeeringConnectionIpv6**](PeeringConnectionIpv6.md) |  | 

## Methods

### NewConnectionPeeringProtocolPostRequest

`func NewConnectionPeeringProtocolPostRequest(type_ ConnectionPeeringProtocolPostRequestType, name string, description string, customerAsn int64, macAddress string, bgpIpv4 PeeringConnectionIpv4, bgpIpv6 PeeringConnectionIpv6, ) *ConnectionPeeringProtocolPostRequest`

NewConnectionPeeringProtocolPostRequest instantiates a new ConnectionPeeringProtocolPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPeeringProtocolPostRequestWithDefaults

`func NewConnectionPeeringProtocolPostRequestWithDefaults() *ConnectionPeeringProtocolPostRequest`

NewConnectionPeeringProtocolPostRequestWithDefaults instantiates a new ConnectionPeeringProtocolPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionPeeringProtocolPostRequest) GetType() ConnectionPeeringProtocolPostRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionPeeringProtocolPostRequest) GetTypeOk() (*ConnectionPeeringProtocolPostRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionPeeringProtocolPostRequest) SetType(v ConnectionPeeringProtocolPostRequestType)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectionPeeringProtocolPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionPeeringProtocolPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionPeeringProtocolPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConnectionPeeringProtocolPostRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionPeeringProtocolPostRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionPeeringProtocolPostRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCustomerAsn

`func (o *ConnectionPeeringProtocolPostRequest) GetCustomerAsn() int64`

GetCustomerAsn returns the CustomerAsn field if non-nil, zero value otherwise.

### GetCustomerAsnOk

`func (o *ConnectionPeeringProtocolPostRequest) GetCustomerAsnOk() (*int64, bool)`

GetCustomerAsnOk returns a tuple with the CustomerAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAsn

`func (o *ConnectionPeeringProtocolPostRequest) SetCustomerAsn(v int64)`

SetCustomerAsn sets CustomerAsn field to given value.


### GetMacAddress

`func (o *ConnectionPeeringProtocolPostRequest) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *ConnectionPeeringProtocolPostRequest) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *ConnectionPeeringProtocolPostRequest) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.


### GetBgpIpv4

`func (o *ConnectionPeeringProtocolPostRequest) GetBgpIpv4() PeeringConnectionIpv4`

GetBgpIpv4 returns the BgpIpv4 field if non-nil, zero value otherwise.

### GetBgpIpv4Ok

`func (o *ConnectionPeeringProtocolPostRequest) GetBgpIpv4Ok() (*PeeringConnectionIpv4, bool)`

GetBgpIpv4Ok returns a tuple with the BgpIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv4

`func (o *ConnectionPeeringProtocolPostRequest) SetBgpIpv4(v PeeringConnectionIpv4)`

SetBgpIpv4 sets BgpIpv4 field to given value.


### GetBgpIpv6

`func (o *ConnectionPeeringProtocolPostRequest) GetBgpIpv6() PeeringConnectionIpv6`

GetBgpIpv6 returns the BgpIpv6 field if non-nil, zero value otherwise.

### GetBgpIpv6Ok

`func (o *ConnectionPeeringProtocolPostRequest) GetBgpIpv6Ok() (*PeeringConnectionIpv6, bool)`

GetBgpIpv6Ok returns a tuple with the BgpIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpIpv6

`func (o *ConnectionPeeringProtocolPostRequest) SetBgpIpv6(v PeeringConnectionIpv6)`

SetBgpIpv6 sets BgpIpv6 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


