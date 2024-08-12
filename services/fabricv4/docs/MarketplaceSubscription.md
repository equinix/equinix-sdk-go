# MarketplaceSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Marketplace Subscription URI | [optional] [readonly] 
**Type** | Pointer to [**MarketplaceSubscriptionType**](MarketplaceSubscriptionType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned Marketplace Subscription identifier | [optional] 

## Methods

### NewMarketplaceSubscription

`func NewMarketplaceSubscription() *MarketplaceSubscription`

NewMarketplaceSubscription instantiates a new MarketplaceSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceSubscriptionWithDefaults

`func NewMarketplaceSubscriptionWithDefaults() *MarketplaceSubscription`

NewMarketplaceSubscriptionWithDefaults instantiates a new MarketplaceSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *MarketplaceSubscription) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MarketplaceSubscription) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MarketplaceSubscription) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *MarketplaceSubscription) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *MarketplaceSubscription) GetType() MarketplaceSubscriptionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MarketplaceSubscription) GetTypeOk() (*MarketplaceSubscriptionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MarketplaceSubscription) SetType(v MarketplaceSubscriptionType)`

SetType sets Type field to given value.

### HasType

`func (o *MarketplaceSubscription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *MarketplaceSubscription) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MarketplaceSubscription) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MarketplaceSubscription) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MarketplaceSubscription) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


