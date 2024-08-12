# SubscriptionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] 
**Marketplace** | Pointer to [**SubscriptionResponseMarketplace**](SubscriptionResponseMarketplace.md) |  | [optional] 
**OfferType** | Pointer to [**SubscriptionResponseOfferType**](SubscriptionResponseOfferType.md) |  | [optional] 
**IsAutoRenew** | Pointer to **bool** |  | [optional] 
**OfferId** | Pointer to **string** |  | [optional] 
**Trial** | Pointer to [**SubscriptionTrial**](SubscriptionTrial.md) |  | [optional] 
**SubscriptionKey** | Pointer to **string** |  | [optional] 
**Entitlements** | Pointer to [**[]SubscriptionEntitlementResponse**](SubscriptionEntitlementResponse.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewSubscriptionResponse

`func NewSubscriptionResponse() *SubscriptionResponse`

NewSubscriptionResponse instantiates a new SubscriptionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionResponseWithDefaults

`func NewSubscriptionResponseWithDefaults() *SubscriptionResponse`

NewSubscriptionResponseWithDefaults instantiates a new SubscriptionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SubscriptionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SubscriptionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SubscriptionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SubscriptionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SubscriptionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SubscriptionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SubscriptionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SubscriptionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionResponse) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionResponse) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionResponse) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMarketplace

`func (o *SubscriptionResponse) GetMarketplace() SubscriptionResponseMarketplace`

GetMarketplace returns the Marketplace field if non-nil, zero value otherwise.

### GetMarketplaceOk

`func (o *SubscriptionResponse) GetMarketplaceOk() (*SubscriptionResponseMarketplace, bool)`

GetMarketplaceOk returns a tuple with the Marketplace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketplace

`func (o *SubscriptionResponse) SetMarketplace(v SubscriptionResponseMarketplace)`

SetMarketplace sets Marketplace field to given value.

### HasMarketplace

`func (o *SubscriptionResponse) HasMarketplace() bool`

HasMarketplace returns a boolean if a field has been set.

### GetOfferType

`func (o *SubscriptionResponse) GetOfferType() SubscriptionResponseOfferType`

GetOfferType returns the OfferType field if non-nil, zero value otherwise.

### GetOfferTypeOk

`func (o *SubscriptionResponse) GetOfferTypeOk() (*SubscriptionResponseOfferType, bool)`

GetOfferTypeOk returns a tuple with the OfferType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferType

`func (o *SubscriptionResponse) SetOfferType(v SubscriptionResponseOfferType)`

SetOfferType sets OfferType field to given value.

### HasOfferType

`func (o *SubscriptionResponse) HasOfferType() bool`

HasOfferType returns a boolean if a field has been set.

### GetIsAutoRenew

`func (o *SubscriptionResponse) GetIsAutoRenew() bool`

GetIsAutoRenew returns the IsAutoRenew field if non-nil, zero value otherwise.

### GetIsAutoRenewOk

`func (o *SubscriptionResponse) GetIsAutoRenewOk() (*bool, bool)`

GetIsAutoRenewOk returns a tuple with the IsAutoRenew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRenew

`func (o *SubscriptionResponse) SetIsAutoRenew(v bool)`

SetIsAutoRenew sets IsAutoRenew field to given value.

### HasIsAutoRenew

`func (o *SubscriptionResponse) HasIsAutoRenew() bool`

HasIsAutoRenew returns a boolean if a field has been set.

### GetOfferId

`func (o *SubscriptionResponse) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *SubscriptionResponse) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *SubscriptionResponse) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *SubscriptionResponse) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetTrial

`func (o *SubscriptionResponse) GetTrial() SubscriptionTrial`

GetTrial returns the Trial field if non-nil, zero value otherwise.

### GetTrialOk

`func (o *SubscriptionResponse) GetTrialOk() (*SubscriptionTrial, bool)`

GetTrialOk returns a tuple with the Trial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrial

`func (o *SubscriptionResponse) SetTrial(v SubscriptionTrial)`

SetTrial sets Trial field to given value.

### HasTrial

`func (o *SubscriptionResponse) HasTrial() bool`

HasTrial returns a boolean if a field has been set.

### GetSubscriptionKey

`func (o *SubscriptionResponse) GetSubscriptionKey() string`

GetSubscriptionKey returns the SubscriptionKey field if non-nil, zero value otherwise.

### GetSubscriptionKeyOk

`func (o *SubscriptionResponse) GetSubscriptionKeyOk() (*string, bool)`

GetSubscriptionKeyOk returns a tuple with the SubscriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionKey

`func (o *SubscriptionResponse) SetSubscriptionKey(v string)`

SetSubscriptionKey sets SubscriptionKey field to given value.

### HasSubscriptionKey

`func (o *SubscriptionResponse) HasSubscriptionKey() bool`

HasSubscriptionKey returns a boolean if a field has been set.

### GetEntitlements

`func (o *SubscriptionResponse) GetEntitlements() []SubscriptionEntitlementResponse`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *SubscriptionResponse) GetEntitlementsOk() (*[]SubscriptionEntitlementResponse, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *SubscriptionResponse) SetEntitlements(v []SubscriptionEntitlementResponse)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *SubscriptionResponse) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetChangelog

`func (o *SubscriptionResponse) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *SubscriptionResponse) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *SubscriptionResponse) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *SubscriptionResponse) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


