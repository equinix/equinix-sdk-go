# ServiceProfileAccessPointTypeCOLO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | [**ServiceProfileAccessPointTypeEnum**](ServiceProfileAccessPointTypeEnum.md) |  | 
**SupportedBandwidths** | Pointer to **[]int32** |  | [optional] 
**AllowRemoteConnections** | Pointer to **bool** | Setting to allow or prohibit remote connections to the service profile. | [optional] [default to false]
**AllowCustomBandwidth** | Pointer to **bool** | Setting to enable or disable the ability of the buyer to customize the bandwidth. | [optional] [default to false]
**BandwidthAlertThreshold** | Pointer to **float32** | percentage of port bandwidth at which an allocation alert is generated - missing on wiki. | [optional] 
**AllowBandwidthAutoApproval** | Pointer to **bool** | Setting to enable or disable the ability of the buyer to change connection bandwidth without approval of the seller. | [optional] [default to false]
**AllowBandwidthUpgrade** | Pointer to **bool** | Availability of a bandwidth upgrade. The default is false. | [optional] 
**LinkProtocolConfig** | Pointer to [**ServiceProfileLinkProtocolConfig**](ServiceProfileLinkProtocolConfig.md) |  | [optional] 
**EnableAutoGenerateServiceKey** | Pointer to **bool** | for verizon only. | [optional] 
**ConnectionRedundancyRequired** | Pointer to **bool** | Mandate redundant connections | [optional] [default to false]
**SelectiveRedundancy** | Pointer to **bool** | Optional redundant connections | [optional] [default to false]
**ApiConfig** | Pointer to [**ApiConfig**](ApiConfig.md) |  | [optional] 
**ConnectionLabel** | Pointer to **string** | custom name for \&quot;Connection\&quot; | [optional] 
**AuthenticationKey** | Pointer to [**AuthenticationKey**](AuthenticationKey.md) |  | [optional] 
**Metadata** | Pointer to [**ServiceProfileMetadata**](ServiceProfileMetadata.md) |  | [optional] 

## Methods

### NewServiceProfileAccessPointTypeCOLO

`func NewServiceProfileAccessPointTypeCOLO(type_ ServiceProfileAccessPointTypeEnum, ) *ServiceProfileAccessPointTypeCOLO`

NewServiceProfileAccessPointTypeCOLO instantiates a new ServiceProfileAccessPointTypeCOLO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProfileAccessPointTypeCOLOWithDefaults

`func NewServiceProfileAccessPointTypeCOLOWithDefaults() *ServiceProfileAccessPointTypeCOLO`

NewServiceProfileAccessPointTypeCOLOWithDefaults instantiates a new ServiceProfileAccessPointTypeCOLO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *ServiceProfileAccessPointTypeCOLO) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ServiceProfileAccessPointTypeCOLO) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ServiceProfileAccessPointTypeCOLO) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *ServiceProfileAccessPointTypeCOLO) GetType() ServiceProfileAccessPointTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetTypeOk() (*ServiceProfileAccessPointTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceProfileAccessPointTypeCOLO) SetType(v ServiceProfileAccessPointTypeEnum)`

SetType sets Type field to given value.


### GetSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeCOLO) GetSupportedBandwidths() []int32`

GetSupportedBandwidths returns the SupportedBandwidths field if non-nil, zero value otherwise.

### GetSupportedBandwidthsOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetSupportedBandwidthsOk() (*[]int32, bool)`

GetSupportedBandwidthsOk returns a tuple with the SupportedBandwidths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeCOLO) SetSupportedBandwidths(v []int32)`

SetSupportedBandwidths sets SupportedBandwidths field to given value.

### HasSupportedBandwidths

`func (o *ServiceProfileAccessPointTypeCOLO) HasSupportedBandwidths() bool`

HasSupportedBandwidths returns a boolean if a field has been set.

### GetAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowRemoteConnections() bool`

GetAllowRemoteConnections returns the AllowRemoteConnections field if non-nil, zero value otherwise.

### GetAllowRemoteConnectionsOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowRemoteConnectionsOk() (*bool, bool)`

GetAllowRemoteConnectionsOk returns a tuple with the AllowRemoteConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeCOLO) SetAllowRemoteConnections(v bool)`

SetAllowRemoteConnections sets AllowRemoteConnections field to given value.

### HasAllowRemoteConnections

`func (o *ServiceProfileAccessPointTypeCOLO) HasAllowRemoteConnections() bool`

HasAllowRemoteConnections returns a boolean if a field has been set.

### GetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowCustomBandwidth() bool`

GetAllowCustomBandwidth returns the AllowCustomBandwidth field if non-nil, zero value otherwise.

### GetAllowCustomBandwidthOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowCustomBandwidthOk() (*bool, bool)`

GetAllowCustomBandwidthOk returns a tuple with the AllowCustomBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeCOLO) SetAllowCustomBandwidth(v bool)`

SetAllowCustomBandwidth sets AllowCustomBandwidth field to given value.

### HasAllowCustomBandwidth

`func (o *ServiceProfileAccessPointTypeCOLO) HasAllowCustomBandwidth() bool`

HasAllowCustomBandwidth returns a boolean if a field has been set.

### GetBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointTypeCOLO) GetBandwidthAlertThreshold() float32`

GetBandwidthAlertThreshold returns the BandwidthAlertThreshold field if non-nil, zero value otherwise.

### GetBandwidthAlertThresholdOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetBandwidthAlertThresholdOk() (*float32, bool)`

GetBandwidthAlertThresholdOk returns a tuple with the BandwidthAlertThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointTypeCOLO) SetBandwidthAlertThreshold(v float32)`

SetBandwidthAlertThreshold sets BandwidthAlertThreshold field to given value.

### HasBandwidthAlertThreshold

`func (o *ServiceProfileAccessPointTypeCOLO) HasBandwidthAlertThreshold() bool`

HasBandwidthAlertThreshold returns a boolean if a field has been set.

### GetAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthAutoApproval() bool`

GetAllowBandwidthAutoApproval returns the AllowBandwidthAutoApproval field if non-nil, zero value otherwise.

### GetAllowBandwidthAutoApprovalOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthAutoApprovalOk() (*bool, bool)`

GetAllowBandwidthAutoApprovalOk returns a tuple with the AllowBandwidthAutoApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointTypeCOLO) SetAllowBandwidthAutoApproval(v bool)`

SetAllowBandwidthAutoApproval sets AllowBandwidthAutoApproval field to given value.

### HasAllowBandwidthAutoApproval

`func (o *ServiceProfileAccessPointTypeCOLO) HasAllowBandwidthAutoApproval() bool`

HasAllowBandwidthAutoApproval returns a boolean if a field has been set.

### GetAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthUpgrade() bool`

GetAllowBandwidthUpgrade returns the AllowBandwidthUpgrade field if non-nil, zero value otherwise.

### GetAllowBandwidthUpgradeOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetAllowBandwidthUpgradeOk() (*bool, bool)`

GetAllowBandwidthUpgradeOk returns a tuple with the AllowBandwidthUpgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointTypeCOLO) SetAllowBandwidthUpgrade(v bool)`

SetAllowBandwidthUpgrade sets AllowBandwidthUpgrade field to given value.

### HasAllowBandwidthUpgrade

`func (o *ServiceProfileAccessPointTypeCOLO) HasAllowBandwidthUpgrade() bool`

HasAllowBandwidthUpgrade returns a boolean if a field has been set.

### GetLinkProtocolConfig

`func (o *ServiceProfileAccessPointTypeCOLO) GetLinkProtocolConfig() ServiceProfileLinkProtocolConfig`

GetLinkProtocolConfig returns the LinkProtocolConfig field if non-nil, zero value otherwise.

### GetLinkProtocolConfigOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetLinkProtocolConfigOk() (*ServiceProfileLinkProtocolConfig, bool)`

GetLinkProtocolConfigOk returns a tuple with the LinkProtocolConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkProtocolConfig

`func (o *ServiceProfileAccessPointTypeCOLO) SetLinkProtocolConfig(v ServiceProfileLinkProtocolConfig)`

SetLinkProtocolConfig sets LinkProtocolConfig field to given value.

### HasLinkProtocolConfig

`func (o *ServiceProfileAccessPointTypeCOLO) HasLinkProtocolConfig() bool`

HasLinkProtocolConfig returns a boolean if a field has been set.

### GetEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointTypeCOLO) GetEnableAutoGenerateServiceKey() bool`

GetEnableAutoGenerateServiceKey returns the EnableAutoGenerateServiceKey field if non-nil, zero value otherwise.

### GetEnableAutoGenerateServiceKeyOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetEnableAutoGenerateServiceKeyOk() (*bool, bool)`

GetEnableAutoGenerateServiceKeyOk returns a tuple with the EnableAutoGenerateServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointTypeCOLO) SetEnableAutoGenerateServiceKey(v bool)`

SetEnableAutoGenerateServiceKey sets EnableAutoGenerateServiceKey field to given value.

### HasEnableAutoGenerateServiceKey

`func (o *ServiceProfileAccessPointTypeCOLO) HasEnableAutoGenerateServiceKey() bool`

HasEnableAutoGenerateServiceKey returns a boolean if a field has been set.

### GetConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionRedundancyRequired() bool`

GetConnectionRedundancyRequired returns the ConnectionRedundancyRequired field if non-nil, zero value otherwise.

### GetConnectionRedundancyRequiredOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionRedundancyRequiredOk() (*bool, bool)`

GetConnectionRedundancyRequiredOk returns a tuple with the ConnectionRedundancyRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointTypeCOLO) SetConnectionRedundancyRequired(v bool)`

SetConnectionRedundancyRequired sets ConnectionRedundancyRequired field to given value.

### HasConnectionRedundancyRequired

`func (o *ServiceProfileAccessPointTypeCOLO) HasConnectionRedundancyRequired() bool`

HasConnectionRedundancyRequired returns a boolean if a field has been set.

### GetSelectiveRedundancy

`func (o *ServiceProfileAccessPointTypeCOLO) GetSelectiveRedundancy() bool`

GetSelectiveRedundancy returns the SelectiveRedundancy field if non-nil, zero value otherwise.

### GetSelectiveRedundancyOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetSelectiveRedundancyOk() (*bool, bool)`

GetSelectiveRedundancyOk returns a tuple with the SelectiveRedundancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectiveRedundancy

`func (o *ServiceProfileAccessPointTypeCOLO) SetSelectiveRedundancy(v bool)`

SetSelectiveRedundancy sets SelectiveRedundancy field to given value.

### HasSelectiveRedundancy

`func (o *ServiceProfileAccessPointTypeCOLO) HasSelectiveRedundancy() bool`

HasSelectiveRedundancy returns a boolean if a field has been set.

### GetApiConfig

`func (o *ServiceProfileAccessPointTypeCOLO) GetApiConfig() ApiConfig`

GetApiConfig returns the ApiConfig field if non-nil, zero value otherwise.

### GetApiConfigOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetApiConfigOk() (*ApiConfig, bool)`

GetApiConfigOk returns a tuple with the ApiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiConfig

`func (o *ServiceProfileAccessPointTypeCOLO) SetApiConfig(v ApiConfig)`

SetApiConfig sets ApiConfig field to given value.

### HasApiConfig

`func (o *ServiceProfileAccessPointTypeCOLO) HasApiConfig() bool`

HasApiConfig returns a boolean if a field has been set.

### GetConnectionLabel

`func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionLabel() string`

GetConnectionLabel returns the ConnectionLabel field if non-nil, zero value otherwise.

### GetConnectionLabelOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetConnectionLabelOk() (*string, bool)`

GetConnectionLabelOk returns a tuple with the ConnectionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLabel

`func (o *ServiceProfileAccessPointTypeCOLO) SetConnectionLabel(v string)`

SetConnectionLabel sets ConnectionLabel field to given value.

### HasConnectionLabel

`func (o *ServiceProfileAccessPointTypeCOLO) HasConnectionLabel() bool`

HasConnectionLabel returns a boolean if a field has been set.

### GetAuthenticationKey

`func (o *ServiceProfileAccessPointTypeCOLO) GetAuthenticationKey() AuthenticationKey`

GetAuthenticationKey returns the AuthenticationKey field if non-nil, zero value otherwise.

### GetAuthenticationKeyOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetAuthenticationKeyOk() (*AuthenticationKey, bool)`

GetAuthenticationKeyOk returns a tuple with the AuthenticationKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationKey

`func (o *ServiceProfileAccessPointTypeCOLO) SetAuthenticationKey(v AuthenticationKey)`

SetAuthenticationKey sets AuthenticationKey field to given value.

### HasAuthenticationKey

`func (o *ServiceProfileAccessPointTypeCOLO) HasAuthenticationKey() bool`

HasAuthenticationKey returns a boolean if a field has been set.

### GetMetadata

`func (o *ServiceProfileAccessPointTypeCOLO) GetMetadata() ServiceProfileMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ServiceProfileAccessPointTypeCOLO) GetMetadataOk() (*ServiceProfileMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ServiceProfileAccessPointTypeCOLO) SetMetadata(v ServiceProfileMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ServiceProfileAccessPointTypeCOLO) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


