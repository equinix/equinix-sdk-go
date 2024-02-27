# ServiceProfileAccessPointTypeVD

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceProfileAccessPointTypeEnum**](ServiceProfileAccessPointTypeEnum.md) |  | 
**SupportedBandwidths** | Pointer to **[]int32** |  | [optional] 
**AllowRemoteConnections** | Pointer to **bool** | Allow remote connections to Service Profile | [optional] 
**AllowCustomBandwidth** | Pointer to **bool** |  | [optional] 

## Methods

### NewServiceProfileAccessPointTypeVD

`func NewServiceProfileAccessPointTypeVD(type_ ServiceProfileAccessPointTypeEnum, ) *ServiceProfileAccessPointTypeVD`

NewServiceProfileAccessPointTypeVD instantiates a new ServiceProfileAccessPointTypeVD object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileAccessPointTypeVDWithDefaults

`func NewServiceProfileAccessPointTypeVDWithDefaults() *ServiceProfileAccessPointTypeVD`

NewServiceProfileAccessPointTypeVDWithDefaults instantiates a new ServiceProfileAccessPointTypeVD object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ServiceProfileAccessPointTypeVD) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileAccessPointTypeVD) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileAccessPointTypeVD) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileAccessPointTypeVD) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileAccessPointTypeVD) GetType() ServiceProfileAccessPointTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileAccessPointTypeVD) GetTypeOk() (*ServiceProfileAccessPointTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileAccessPointTypeVD) SetType(v ServiceProfileAccessPointTypeEnum)`

SetType sets Type field to given value.


### GetSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeVD) GetSupportedBandwidths() []int32`

GetSupportedBandwidths returns the SupportedBandwidths field if non-nil, zero value otherwise.

### GetSupportedBandwidthsOk

`func (o *ServiceProfileAccessPointTypeVD) GetSupportedBandwidthsOk() (*[]int32, bool)`

GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeVD) SetSupportedBandwidths(v []int32)`

SetSupportedBandwidths sets SupportedBandwidths field to given value.

### HasSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeVD) HasSupportedBandwidths() bool`

HasSupportedBandwidths returns a boolean if a field has been set.

### GetAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeVD) GetAllowRemoteConnections() bool`

GetAllowRemoteConnections returns the AllowRemoteConnections field if non-nil, zero value otherwise.

### GetAllowRemoteConnectionsOk

`func (o *ServiceProfileAccessPointTypeVD) GetAllowRemoteConnectionsOk() (*bool, bool)`

GetAllowRemoteConnectionsOk returns a tuple with the AllowRemoteConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeVD) SetAllowRemoteConnections(v bool)`

SetAllowRemoteConnections sets AllowRemoteConnections field to given value.

### HasAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeVD) HasAllowRemoteConnections() bool`

HasAllowRemoteConnections returns a boolean if a field has been set.

### GetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeVD) GetAllowCustomBandwidth() bool`

GetAllowCustomBandwidth returns the AllowCustomBandwidth field if non-nil, zero value otherwise.

### GetAllowCustomBandwidthOk

`func (o *ServiceProfileAccessPointTypeVD) GetAllowCustomBandwidthOk() (*bool, bool)`

GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeVD) SetAllowCustomBandwidth(v bool)`

SetAllowCustomBandwidth sets AllowCustomBandwidth field to given value.

### HasAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeVD) HasAllowCustomBandwidth() bool`

HasAllowCustomBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


