# SubscriptionAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type of the subscription asset ( XF_ROUTER ,IP_VC, IPWAN_VC ) | [optional] 
**Package** | Pointer to [**SubscriptionRouterPackageType**](SubscriptionRouterPackageType.md) |  | [optional] 
**Bandwidth** | Pointer to **int32** | Bandwidth of the asset in Mbps | [optional] 

## Methods

### NewSubscriptionAsset

`func NewSubscriptionAsset() *SubscriptionAsset`

NewSubscriptionAsset instantiates a new SubscriptionAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAssetWithDefaults

`func NewSubscriptionAssetWithDefaults() *SubscriptionAsset`

NewSubscriptionAssetWithDefaults instantiates a new SubscriptionAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPackage

`func (o *SubscriptionAsset) GetPackage() SubscriptionRouterPackageType`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *SubscriptionAsset) GetPackageOk() (*SubscriptionRouterPackageType, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *SubscriptionAsset) SetPackage(v SubscriptionRouterPackageType)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *SubscriptionAsset) HasPackage() bool`

HasPackage returns a boolean if a field has been set.

### GetBandwidth

`func (o *SubscriptionAsset) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *SubscriptionAsset) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *SubscriptionAsset) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *SubscriptionAsset) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


