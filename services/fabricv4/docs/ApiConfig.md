# ApiConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAvailable** | Pointer to **bool** | Setting indicating whether the API is available (true) or not (false). | [optional] [default to false]
**IntegrationId** | Pointer to **string** |  | [optional] 
**EquinixManagedPort** | Pointer to **bool** | Setting indicating that the port is managed by Equinix (true) or not (false). | [optional] [default to false]
**EquinixManagedVlan** | Pointer to **bool** | Setting indicating that the VLAN is managed by Equinix (true) or not (false). | [optional] [default to false]
**AllowOverSubscription** | Pointer to **bool** | Setting showing that oversubscription support is available (true) or not (false). The default is false. Oversubscription is the sale of more than the available network bandwidth. This practice is common and legitimate. After all, many customers use less bandwidth than they&#39;ve purchased. And network users don&#39;t consume bandwidth all at the same time. The leftover bandwidth can be sold to other customers. When demand surges, operational and engineering resources can be shifted to accommodate the load.  | [optional] [default to false]
**OverSubscriptionLimit** | Pointer to **int32** | A cap on oversubscription. | [optional] [default to 1]
**BandwidthFromApi** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewApiConfig

`func NewApiConfig() *ApiConfig`

NewApiConfig instantiates a new ApiConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiConfigWithDefaults

`func NewApiConfigWithDefaults() *ApiConfig`

NewApiConfigWithDefaults instantiates a new ApiConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAvailable

`func (o *ApiConfig) GetApiAvailable() bool`

GetApiAvailable returns the ApiAvailable field if non-nil, zero value otherwise.

### GetApiAvailableOk

`func (o *ApiConfig) GetApiAvailableOk() (*bool, bool)`

GetApiAvailableOk returns a tuple with the ApiAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAvailable

`func (o *ApiConfig) SetApiAvailable(v bool)`

SetApiAvailable sets ApiAvailable field to given value.

### HasApiAvailable

`func (o *ApiConfig) HasApiAvailable() bool`

HasApiAvailable returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ApiConfig) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ApiConfig) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ApiConfig) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ApiConfig) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetEquinixManagedPort

`func (o *ApiConfig) GetEquinixManagedPort() bool`

GetEquinixManagedPort returns the EquinixManagedPort field if non-nil, zero value otherwise.

### GetEquinixManagedPortOk

`func (o *ApiConfig) GetEquinixManagedPortOk() (*bool, bool)`

GetEquinixManagedPortOk returns a tuple with the EquinixManagedPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixManagedPort

`func (o *ApiConfig) SetEquinixManagedPort(v bool)`

SetEquinixManagedPort sets EquinixManagedPort field to given value.

### HasEquinixManagedPort

`func (o *ApiConfig) HasEquinixManagedPort() bool`

HasEquinixManagedPort returns a boolean if a field has been set.

### GetEquinixManagedVlan

`func (o *ApiConfig) GetEquinixManagedVlan() bool`

GetEquinixManagedVlan returns the EquinixManagedVlan field if non-nil, zero value otherwise.

### GetEquinixManagedVlanOk

`func (o *ApiConfig) GetEquinixManagedVlanOk() (*bool, bool)`

GetEquinixManagedVlanOk returns a tuple with the EquinixManagedVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquinixManagedVlan

`func (o *ApiConfig) SetEquinixManagedVlan(v bool)`

SetEquinixManagedVlan sets EquinixManagedVlan field to given value.

### HasEquinixManagedVlan

`func (o *ApiConfig) HasEquinixManagedVlan() bool`

HasEquinixManagedVlan returns a boolean if a field has been set.

### GetAllowOverSubscription

`func (o *ApiConfig) GetAllowOverSubscription() bool`

GetAllowOverSubscription returns the AllowOverSubscription field if non-nil, zero value otherwise.

### GetAllowOverSubscriptionOk

`func (o *ApiConfig) GetAllowOverSubscriptionOk() (*bool, bool)`

GetAllowOverSubscriptionOk returns a tuple with the AllowOverSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverSubscription

`func (o *ApiConfig) SetAllowOverSubscription(v bool)`

SetAllowOverSubscription sets AllowOverSubscription field to given value.

### HasAllowOverSubscription

`func (o *ApiConfig) HasAllowOverSubscription() bool`

HasAllowOverSubscription returns a boolean if a field has been set.

### GetOverSubscriptionLimit

`func (o *ApiConfig) GetOverSubscriptionLimit() int32`

GetOverSubscriptionLimit returns the OverSubscriptionLimit field if non-nil, zero value otherwise.

### GetOverSubscriptionLimitOk

`func (o *ApiConfig) GetOverSubscriptionLimitOk() (*int32, bool)`

GetOverSubscriptionLimitOk returns a tuple with the OverSubscriptionLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverSubscriptionLimit

`func (o *ApiConfig) SetOverSubscriptionLimit(v int32)`

SetOverSubscriptionLimit sets OverSubscriptionLimit field to given value.

### HasOverSubscriptionLimit

`func (o *ApiConfig) HasOverSubscriptionLimit() bool`

HasOverSubscriptionLimit returns a boolean if a field has been set.

### GetBandwidthFromApi

`func (o *ApiConfig) GetBandwidthFromApi() bool`

GetBandwidthFromApi returns the BandwidthFromApi field if non-nil, zero value otherwise.

### GetBandwidthFromApiOk

`func (o *ApiConfig) GetBandwidthFromApiOk() (*bool, bool)`

GetBandwidthFromApiOk returns a tuple with the BandwidthFromApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthFromApi

`func (o *ApiConfig) SetBandwidthFromApi(v bool)`

SetBandwidthFromApi sets BandwidthFromApi field to given value.

### HasBandwidthFromApi

`func (o *ApiConfig) HasBandwidthFromApi() bool`

HasBandwidthFromApi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


