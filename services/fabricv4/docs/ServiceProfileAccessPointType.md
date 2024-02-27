# ServiceProfileAccessPointType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceProfileAccessPointTypeEnum**](ServiceProfileAccessPointTypeEnum.md) |  | 
**SupportedBandwidths** | Pointer to **[]int32** |  | [optional] 
**AllowRemoteConnections** | Pointer to **bool** | Allow remote connections to Service Profile | [optional] 
**AllowCustomBandwidth** | Pointer to **bool** |  | [optional] 
**BandwidthAlertThreshold** | Pointer to **float32** | percentage of port bandwidth at which an allocation alert is generated - missing on wiki. | [optional] 
**AllowBandwidthAutoApproval** | Pointer to **bool** | Setting to enable or disable the ability of the buyer to change connection bandwidth without approval of the seller. | [optional] [default to false]
**AllowBandwidthUpgrade** | Pointer to **bool** | Availability of a bandwidth upgrade. The default is false. | [optional] 
**LinkProtocolConfig** | Pointer to [**ServiceProfileLinkProtocolConfig**](ServiceProfileLinkProtocolConfig.md) |  | [optional] 
**EnableAutoGenerateServiceKey** | Pointer to **bool** | for verizon only. | [optional] 
**ConnectionRedundancyRequired** | Pointer to **bool** | Mandate redundant connections | [optional] [default to false]
**ApiConfig** | Pointer to [**ApiConfig**](ApiConfig.md) |  | [optional] 
**ConnectionLabel** | Pointer to **string** | custom name for \&quot;Connection\&quot; | [optional] 
**AuthenticationKey** | Pointer to [**AuthenticationKey**](AuthenticationKey.md) |  | [optional] 
**Metadata** | Pointer to [**ServiceProfileMetadata**](ServiceProfileMetadata.md) |  | [optional] 

## Methods

### NewServiceProfileAccessPointType

`func NewServiceProfileAccessPointType(type_ ServiceProfileAccessPointTypeEnum, ) *ServiceProfileAccessPointType`

NewServiceProfileAccessPointType instantiates a new ServiceProfileAccessPointType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileAccessPointTypeWithDefaults

`func NewServiceProfileAccessPointTypeWithDefaults() *ServiceProfileAccessPointType`

NewServiceProfileAccessPointTypeWithDefaults instantiates a new ServiceProfileAccessPointType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ServiceProfileAccessPointType) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileAccessPointType) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileAccessPointType) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileAccessPointType) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileAccessPointType) GetType() ServiceProfileAccessPointTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileAccessPointType) GetTypeOk() (*ServiceProfileAccessPointTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileAccessPointType) SetType(v ServiceProfileAccessPointTypeEnum)`

SetType sets Type field to given value.


### GetSupportedBandwidths

`func (o *ServiceProfileAccessPointType) GetSupportedBandwidths() []int32`

GetSupportedBandwidths returns the SupportedBandwidths field if non-nil, zero value otherwise.

### GetSupportedBandwidthsOk

`func (o *ServiceProfileAccessPointType) GetSupportedBandwidthsOk() (*[]int32, bool)`

GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBandwidths

`func (o *ServiceProfileAccessPointType) SetSupportedBandwidths(v []int32)`

SetSupportedBandwidths sets SupportedBandwidths field to given value.

### HasSupportedBandwidths

`func (o *ServiceProfileAccessPointType) HasSupportedBandwidths() bool`

HasSupportedBandwidths returns a boolean if a field has been set.

### GetAllowRemoteConnections

`func (o *ServiceProfileAccessPointType) GetAllowRemoteConnections() bool`

GetAllowRemoteConnections returns the AllowRemoteConnections field if non-nil, zero value otherwise.

### GetAllowRemoteConnectionsOk

`func (o *ServiceProfileAccessPointType) GetAllowRemoteConnectionsOk() (*bool, bool)`

GetAllowRemoteConnectionsOk returns a tuple with the AllowRemoteConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteConnections

`func (o *ServiceProfileAccessPointType) SetAllowRemoteConnections(v bool)`

SetAllowRemoteConnections sets AllowRemoteConnections field to given value.

### HasAllowRemoteConnections

`func (o *ServiceProfileAccessPointType) HasAllowRemoteConnections() bool`

HasAllowRemoteConnections returns a boolean if a field has been set.

### GetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointType) GetAllowCustomBandwidth() bool`

GetAllowCustomBandwidth returns the AllowCustomBandwidth field if non-nil, zero value otherwise.

### GetAllowCustomBandwidthOk

`func (o *ServiceProfileAccessPointType) GetAllowCustomBandwidthOk() (*bool, bool)`

GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointType) SetAllowCustomBandwidth(v bool)`

SetAllowCustomBandwidth sets AllowCustomBandwidth field to given value.

### HasAllowCustomBandwidth

`func (o *ServiceProfileAccessPointType) HasAllowCustomBandwidth() bool`

HasAllowCustomBandwidth returns a boolean if a field has been set.

### GetBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointType) GetBandwidthAlertThreshold() float32`

GetBandwidthAlertThreshold returns the BandwidthAlertThreshold field if non-nil, zero value otherwise.

### GetBandwidthAlertThresholdOk

`func (o *ServiceProfileAccessPointType) GetBandwidthAlertThresholdOk() (*float32, bool)`

GetBandwidthAlertThresholdOk returns a tuple with the BandwidthAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointType) SetBandwidthAlertThreshold(v float32)`

SetBandwidthAlertThreshold sets BandwidthAlertThreshold field to given value.

### HasBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointType) HasBandwidthAlertThreshold() bool`

HasBandwidthAlertThreshold returns a boolean if a field has been set.

### GetAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointType) GetAllowBandwidthAutoApproval() bool`

GetAllowBandwidthAutoApproval returns the AllowBandwidthAutoApproval field if non-nil, zero value otherwise.

### GetAllowBandwidthAutoApprovalOk

`func (o *ServiceProfileAccessPointType) GetAllowBandwidthAutoApprovalOk() (*bool, bool)`

GetAllowBandwidthAutoApprovalOk returns a tuple with the AllowBandwidthAutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointType) SetAllowBandwidthAutoApproval(v bool)`

SetAllowBandwidthAutoApproval sets AllowBandwidthAutoApproval field to given value.

### HasAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointType) HasAllowBandwidthAutoApproval() bool`

HasAllowBandwidthAutoApproval returns a boolean if a field has been set.

### GetAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointType) GetAllowBandwidthUpgrade() bool`

GetAllowBandwidthUpgrade returns the AllowBandwidthUpgrade field if non-nil, zero value otherwise.

### GetAllowBandwidthUpgradeOk

`func (o *ServiceProfileAccessPointType) GetAllowBandwidthUpgradeOk() (*bool, bool)`

GetAllowBandwidthUpgradeOk returns a tuple with the AllowBandwidthUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointType) SetAllowBandwidthUpgrade(v bool)`

SetAllowBandwidthUpgrade sets AllowBandwidthUpgrade field to given value.

### HasAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointType) HasAllowBandwidthUpgrade() bool`

HasAllowBandwidthUpgrade returns a boolean if a field has been set.

### GetLinkProtocolConfig

`func (o *ServiceProfileAccessPointType) GetLinkProtocolConfig() ServiceProfileLinkProtocolConfig`

GetLinkProtocolConfig returns the LinkProtocolConfig field if non-nil, zero value otherwise.

### GetLinkProtocolConfigOk

`func (o *ServiceProfileAccessPointType) GetLinkProtocolConfigOk() (*ServiceProfileLinkProtocolConfig, bool)`

GetLinkProtocolConfigOk returns a tuple with the LinkProtocolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkProtocolConfig

`func (o *ServiceProfileAccessPointType) SetLinkProtocolConfig(v ServiceProfileLinkProtocolConfig)`

SetLinkProtocolConfig sets LinkProtocolConfig field to given value.

### HasLinkProtocolConfig

`func (o *ServiceProfileAccessPointType) HasLinkProtocolConfig() bool`

HasLinkProtocolConfig returns a boolean if a field has been set.

### GetEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointType) GetEnableAutoGenerateServiceKey() bool`

GetEnableAutoGenerateServiceKey returns the EnableAutoGenerateServiceKey field if non-nil, zero value otherwise.

### GetEnableAutoGenerateServiceKeyOk

`func (o *ServiceProfileAccessPointType) GetEnableAutoGenerateServiceKeyOk() (*bool, bool)`

GetEnableAutoGenerateServiceKeyOk returns a tuple with the EnableAutoGenerateServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointType) SetEnableAutoGenerateServiceKey(v bool)`

SetEnableAutoGenerateServiceKey sets EnableAutoGenerateServiceKey field to given value.

### HasEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointType) HasEnableAutoGenerateServiceKey() bool`

HasEnableAutoGenerateServiceKey returns a boolean if a field has been set.

### GetConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointType) GetConnectionRedundancyRequired() bool`

GetConnectionRedundancyRequired returns the ConnectionRedundancyRequired field if non-nil, zero value otherwise.

### GetConnectionRedundancyRequiredOk

`func (o *ServiceProfileAccessPointType) GetConnectionRedundancyRequiredOk() (*bool, bool)`

GetConnectionRedundancyRequiredOk returns a tuple with the ConnectionRedundancyRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointType) SetConnectionRedundancyRequired(v bool)`

SetConnectionRedundancyRequired sets ConnectionRedundancyRequired field to given value.

### HasConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointType) HasConnectionRedundancyRequired() bool`

HasConnectionRedundancyRequired returns a boolean if a field has been set.

### GetApiConfig

`func (o *ServiceProfileAccessPointType) GetApiConfig() ApiConfig`

GetApiConfig returns the ApiConfig field if non-nil, zero value otherwise.

### GetApiConfigOk

`func (o *ServiceProfileAccessPointType) GetApiConfigOk() (*ApiConfig, bool)`

GetApiConfigOk returns a tuple with the ApiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConfig

`func (o *ServiceProfileAccessPointType) SetApiConfig(v ApiConfig)`

SetApiConfig sets ApiConfig field to given value.

### HasApiConfig

`func (o *ServiceProfileAccessPointType) HasApiConfig() bool`

HasApiConfig returns a boolean if a field has been set.

### GetConnectionLabel

`func (o *ServiceProfileAccessPointType) GetConnectionLabel() string`

GetConnectionLabel returns the ConnectionLabel field if non-nil, zero value otherwise.

### GetConnectionLabelOk

`func (o *ServiceProfileAccessPointType) GetConnectionLabelOk() (*string, bool)`

GetConnectionLabelOk returns a tuple with the ConnectionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLabel

`func (o *ServiceProfileAccessPointType) SetConnectionLabel(v string)`

SetConnectionLabel sets ConnectionLabel field to given value.

### HasConnectionLabel

`func (o *ServiceProfileAccessPointType) HasConnectionLabel() bool`

HasConnectionLabel returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *ServiceProfileAccessPointType) GetAuthenticationKey() AuthenticationKey`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *ServiceProfileAccessPointType) GetAuthenticationKeyOk() (*AuthenticationKey, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *ServiceProfileAccessPointType) SetAuthenticationKey(v AuthenticationKey)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *ServiceProfileAccessPointType) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceProfileAccessPointType) GetMetadata() ServiceProfileMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceProfileAccessPointType) GetMetadataOk() (*ServiceProfileMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceProfileAccessPointType) SetMetadata(v ServiceProfileMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceProfileAccessPointType) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


