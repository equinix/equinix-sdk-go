# ServiceTokenConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ServiceTokenConnectionType**](ServiceTokenConnectionType.md) |  | 
**Href** | Pointer to **string** | An absolute URL that is the subject of the link&#39;s context. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned connection identifier | [optional] 
**AllowRemoteConnection** | Pointer to **bool** | Authorization to connect remotely | [optional] [default to false]
**AllowCustomBandwidth** | Pointer to **bool** | Allow custom bandwidth value | [optional] [default to false]
**BandwidthLimit** | Pointer to **int32** | Connection bandwidth limit in Mbps | [optional] 
**SupportedBandwidths** | Pointer to **[]int32** | List of permitted bandwidths. | [optional] 
**ASide** | Pointer to [**ServiceTokenSide**](ServiceTokenSide.md) |  | [optional] 
**ZSide** | Pointer to [**ServiceTokenSide**](ServiceTokenSide.md) |  | [optional] 

## Methods

### NewServiceTokenConnection

`func NewServiceTokenConnection(type_ ServiceTokenConnectionType, ) *ServiceTokenConnection`

NewServiceTokenConnection instantiates a new ServiceTokenConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTokenConnectionWithDefaults

`func NewServiceTokenConnectionWithDefaults() *ServiceTokenConnection`

NewServiceTokenConnectionWithDefaults instantiates a new ServiceTokenConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceTokenConnection) GetType() ServiceTokenConnectionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceTokenConnection) GetTypeOk() (*ServiceTokenConnectionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceTokenConnection) SetType(v ServiceTokenConnectionType)`

SetType sets Type field to given value.


### GetHref

`func (o *ServiceTokenConnection) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ServiceTokenConnection) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ServiceTokenConnection) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ServiceTokenConnection) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *ServiceTokenConnection) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceTokenConnection) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceTokenConnection) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceTokenConnection) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAllowRemoteConnection

`func (o *ServiceTokenConnection) GetAllowRemoteConnection() bool`

GetAllowRemoteConnection returns the AllowRemoteConnection field if non-nil, zero value otherwise.

### GetAllowRemoteConnectionOk

`func (o *ServiceTokenConnection) GetAllowRemoteConnectionOk() (*bool, bool)`

GetAllowRemoteConnectionOk returns a tuple with the AllowRemoteConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteConnection

`func (o *ServiceTokenConnection) SetAllowRemoteConnection(v bool)`

SetAllowRemoteConnection sets AllowRemoteConnection field to given value.

### HasAllowRemoteConnection

`func (o *ServiceTokenConnection) HasAllowRemoteConnection() bool`

HasAllowRemoteConnection returns a boolean if a field has been set.

### GetAllowCustomBandwidth

`func (o *ServiceTokenConnection) GetAllowCustomBandwidth() bool`

GetAllowCustomBandwidth returns the AllowCustomBandwidth field if non-nil, zero value otherwise.

### GetAllowCustomBandwidthOk

`func (o *ServiceTokenConnection) GetAllowCustomBandwidthOk() (*bool, bool)`

GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomBandwidth

`func (o *ServiceTokenConnection) SetAllowCustomBandwidth(v bool)`

SetAllowCustomBandwidth sets AllowCustomBandwidth field to given value.

### HasAllowCustomBandwidth

`func (o *ServiceTokenConnection) HasAllowCustomBandwidth() bool`

HasAllowCustomBandwidth returns a boolean if a field has been set.

### GetBandwidthLimit

`func (o *ServiceTokenConnection) GetBandwidthLimit() int32`

GetBandwidthLimit returns the BandwidthLimit field if non-nil, zero value otherwise.

### GetBandwidthLimitOk

`func (o *ServiceTokenConnection) GetBandwidthLimitOk() (*int32, bool)`

GetBandwidthLimitOk returns a tuple with the BandwidthLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthLimit

`func (o *ServiceTokenConnection) SetBandwidthLimit(v int32)`

SetBandwidthLimit sets BandwidthLimit field to given value.

### HasBandwidthLimit

`func (o *ServiceTokenConnection) HasBandwidthLimit() bool`

HasBandwidthLimit returns a boolean if a field has been set.

### GetSupportedBandwidths

`func (o *ServiceTokenConnection) GetSupportedBandwidths() []int32`

GetSupportedBandwidths returns the SupportedBandwidths field if non-nil, zero value otherwise.

### GetSupportedBandwidthsOk

`func (o *ServiceTokenConnection) GetSupportedBandwidthsOk() (*[]int32, bool)`

GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBandwidths

`func (o *ServiceTokenConnection) SetSupportedBandwidths(v []int32)`

SetSupportedBandwidths sets SupportedBandwidths field to given value.

### HasSupportedBandwidths

`func (o *ServiceTokenConnection) HasSupportedBandwidths() bool`

HasSupportedBandwidths returns a boolean if a field has been set.

### GetASide

`func (o *ServiceTokenConnection) GetASide() ServiceTokenSide`

GetASide returns the ASide field if non-nil, zero value otherwise.

### GetASideOk

`func (o *ServiceTokenConnection) GetASideOk() (*ServiceTokenSide, bool)`

GetASideOk returns a tuple with the ASide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetASide

`func (o *ServiceTokenConnection) SetASide(v ServiceTokenSide)`

SetASide sets ASide field to given value.

### HasASide

`func (o *ServiceTokenConnection) HasASide() bool`

HasASide returns a boolean if a field has been set.

### GetZSide

`func (o *ServiceTokenConnection) GetZSide() ServiceTokenSide`

GetZSide returns the ZSide field if non-nil, zero value otherwise.

### GetZSideOk

`func (o *ServiceTokenConnection) GetZSideOk() (*ServiceTokenSide, bool)`

GetZSideOk returns a tuple with the ZSide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZSide

`func (o *ServiceTokenConnection) SetZSide(v ServiceTokenSide)`

SetZSide sets ZSide field to given value.

### HasZSide

`func (o *ServiceTokenConnection) HasZSide() bool`

HasZSide returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


