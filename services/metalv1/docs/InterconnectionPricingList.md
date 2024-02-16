# InterconnectionPricingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProviderPricing** | Pointer to [**[]InterconnectionPricingListProviderPricingInner**](InterconnectionPricingListProviderPricingInner.md) | Pricing information per connection provider. | [optional] 

## Methods

### NewInterconnectionPricingList

`func NewInterconnectionPricingList() *InterconnectionPricingList`

NewInterconnectionPricingList instantiates a new InterconnectionPricingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterconnectionPricingListWithDefaults

`func NewInterconnectionPricingListWithDefaults() *InterconnectionPricingList`

NewInterconnectionPricingListWithDefaults instantiates a new InterconnectionPricingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviderPricing

`func (o *InterconnectionPricingList) GetProviderPricing() []InterconnectionPricingListProviderPricingInner`

GetProviderPricing returns the ProviderPricing field if non-nil, zero value otherwise.

### GetProviderPricingOk

`func (o *InterconnectionPricingList) GetProviderPricingOk() (*[]InterconnectionPricingListProviderPricingInner, bool)`

GetProviderPricingOk returns a tuple with the ProviderPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderPricing

`func (o *InterconnectionPricingList) SetProviderPricing(v []InterconnectionPricingListProviderPricingInner)`

SetProviderPricing sets ProviderPricing field to given value.

### HasProviderPricing

`func (o *InterconnectionPricingList) HasProviderPricing() bool`

HasProviderPricing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


