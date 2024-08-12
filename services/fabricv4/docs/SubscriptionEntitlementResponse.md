# SubscriptionEntitlementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Subscription Entitlement Id | [optional] 
**QuantityEntitled** | Pointer to **int32** |  | [optional] 
**QuantityConsumed** | Pointer to **int32** |  | [optional] 
**QuantityAvailable** | Pointer to **int32** |  | [optional] 
**Asset** | Pointer to [**SubscriptionAsset**](SubscriptionAsset.md) |  | [optional] 

## Methods

### NewSubscriptionEntitlementResponse

`func NewSubscriptionEntitlementResponse() *SubscriptionEntitlementResponse`

NewSubscriptionEntitlementResponse instantiates a new SubscriptionEntitlementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEntitlementResponseWithDefaults

`func NewSubscriptionEntitlementResponseWithDefaults() *SubscriptionEntitlementResponse`

NewSubscriptionEntitlementResponseWithDefaults instantiates a new SubscriptionEntitlementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SubscriptionEntitlementResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SubscriptionEntitlementResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SubscriptionEntitlementResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SubscriptionEntitlementResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetQuantityEntitled

`func (o *SubscriptionEntitlementResponse) GetQuantityEntitled() int32`

GetQuantityEntitled returns the QuantityEntitled field if non-nil, zero value otherwise.

### GetQuantityEntitledOk

`func (o *SubscriptionEntitlementResponse) GetQuantityEntitledOk() (*int32, bool)`

GetQuantityEntitledOk returns a tuple with the QuantityEntitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityEntitled

`func (o *SubscriptionEntitlementResponse) SetQuantityEntitled(v int32)`

SetQuantityEntitled sets QuantityEntitled field to given value.

### HasQuantityEntitled

`func (o *SubscriptionEntitlementResponse) HasQuantityEntitled() bool`

HasQuantityEntitled returns a boolean if a field has been set.

### GetQuantityConsumed

`func (o *SubscriptionEntitlementResponse) GetQuantityConsumed() int32`

GetQuantityConsumed returns the QuantityConsumed field if non-nil, zero value otherwise.

### GetQuantityConsumedOk

`func (o *SubscriptionEntitlementResponse) GetQuantityConsumedOk() (*int32, bool)`

GetQuantityConsumedOk returns a tuple with the QuantityConsumed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityConsumed

`func (o *SubscriptionEntitlementResponse) SetQuantityConsumed(v int32)`

SetQuantityConsumed sets QuantityConsumed field to given value.

### HasQuantityConsumed

`func (o *SubscriptionEntitlementResponse) HasQuantityConsumed() bool`

HasQuantityConsumed returns a boolean if a field has been set.

### GetQuantityAvailable

`func (o *SubscriptionEntitlementResponse) GetQuantityAvailable() int32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *SubscriptionEntitlementResponse) GetQuantityAvailableOk() (*int32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *SubscriptionEntitlementResponse) SetQuantityAvailable(v int32)`

SetQuantityAvailable sets QuantityAvailable field to given value.

### HasQuantityAvailable

`func (o *SubscriptionEntitlementResponse) HasQuantityAvailable() bool`

HasQuantityAvailable returns a boolean if a field has been set.

### GetAsset

`func (o *SubscriptionEntitlementResponse) GetAsset() SubscriptionAsset`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *SubscriptionEntitlementResponse) GetAssetOk() (*SubscriptionAsset, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *SubscriptionEntitlementResponse) SetAsset(v SubscriptionAsset)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *SubscriptionEntitlementResponse) HasAsset() bool`

HasAsset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


