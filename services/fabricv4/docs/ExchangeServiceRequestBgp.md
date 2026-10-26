# ExchangeServiceRequestBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainName** | Pointer to **string** | Domain name for BGP | [optional] 
**AsSet** | Pointer to **string** | AS-SET used to filter the prefixes announced by the customer | [optional] 
**MlpeEnabled** | Pointer to **bool** | Whether Multi-Lateral Peering Exchange (MLPE) is enabled | [optional] [default to false]
**Md5AuthKey** | Pointer to **string** | MD5 key for BGP | [optional] 
**RcMd5AuthKey** | Pointer to **string** | RC MD5 key for BGP | [optional] 
**Prefixes** | Pointer to **[]string** | List of IP prefix | [optional] 
**MaxPrefixLimit** | Pointer to **int64** | Max Prefix Limit number | [optional] 
**PrependSelfEnabled** | Pointer to **bool** | Prepend Self Enabled | [optional] [default to true]

## Methods

### NewExchangeServiceRequestBgp

`func NewExchangeServiceRequestBgp() *ExchangeServiceRequestBgp`

NewExchangeServiceRequestBgp instantiates a new ExchangeServiceRequestBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceRequestBgpWithDefaults

`func NewExchangeServiceRequestBgpWithDefaults() *ExchangeServiceRequestBgp`

NewExchangeServiceRequestBgpWithDefaults instantiates a new ExchangeServiceRequestBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainName

`func (o *ExchangeServiceRequestBgp) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ExchangeServiceRequestBgp) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ExchangeServiceRequestBgp) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ExchangeServiceRequestBgp) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetAsSet

`func (o *ExchangeServiceRequestBgp) GetAsSet() string`

GetAsSet returns the AsSet field if non-nil, zero value otherwise.

### GetAsSetOk

`func (o *ExchangeServiceRequestBgp) GetAsSetOk() (*string, bool)`

GetAsSetOk returns a tuple with the AsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSet

`func (o *ExchangeServiceRequestBgp) SetAsSet(v string)`

SetAsSet sets AsSet field to given value.

### HasAsSet

`func (o *ExchangeServiceRequestBgp) HasAsSet() bool`

HasAsSet returns a boolean if a field has been set.

### GetMlpeEnabled

`func (o *ExchangeServiceRequestBgp) GetMlpeEnabled() bool`

GetMlpeEnabled returns the MlpeEnabled field if non-nil, zero value otherwise.

### GetMlpeEnabledOk

`func (o *ExchangeServiceRequestBgp) GetMlpeEnabledOk() (*bool, bool)`

GetMlpeEnabledOk returns a tuple with the MlpeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlpeEnabled

`func (o *ExchangeServiceRequestBgp) SetMlpeEnabled(v bool)`

SetMlpeEnabled sets MlpeEnabled field to given value.

### HasMlpeEnabled

`func (o *ExchangeServiceRequestBgp) HasMlpeEnabled() bool`

HasMlpeEnabled returns a boolean if a field has been set.

### GetMd5AuthKey

`func (o *ExchangeServiceRequestBgp) GetMd5AuthKey() string`

GetMd5AuthKey returns the Md5AuthKey field if non-nil, zero value otherwise.

### GetMd5AuthKeyOk

`func (o *ExchangeServiceRequestBgp) GetMd5AuthKeyOk() (*string, bool)`

GetMd5AuthKeyOk returns a tuple with the Md5AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthKey

`func (o *ExchangeServiceRequestBgp) SetMd5AuthKey(v string)`

SetMd5AuthKey sets Md5AuthKey field to given value.

### HasMd5AuthKey

`func (o *ExchangeServiceRequestBgp) HasMd5AuthKey() bool`

HasMd5AuthKey returns a boolean if a field has been set.

### GetRcMd5AuthKey

`func (o *ExchangeServiceRequestBgp) GetRcMd5AuthKey() string`

GetRcMd5AuthKey returns the RcMd5AuthKey field if non-nil, zero value otherwise.

### GetRcMd5AuthKeyOk

`func (o *ExchangeServiceRequestBgp) GetRcMd5AuthKeyOk() (*string, bool)`

GetRcMd5AuthKeyOk returns a tuple with the RcMd5AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcMd5AuthKey

`func (o *ExchangeServiceRequestBgp) SetRcMd5AuthKey(v string)`

SetRcMd5AuthKey sets RcMd5AuthKey field to given value.

### HasRcMd5AuthKey

`func (o *ExchangeServiceRequestBgp) HasRcMd5AuthKey() bool`

HasRcMd5AuthKey returns a boolean if a field has been set.

### GetPrefixes

`func (o *ExchangeServiceRequestBgp) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *ExchangeServiceRequestBgp) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *ExchangeServiceRequestBgp) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *ExchangeServiceRequestBgp) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetMaxPrefixLimit

`func (o *ExchangeServiceRequestBgp) GetMaxPrefixLimit() int64`

GetMaxPrefixLimit returns the MaxPrefixLimit field if non-nil, zero value otherwise.

### GetMaxPrefixLimitOk

`func (o *ExchangeServiceRequestBgp) GetMaxPrefixLimitOk() (*int64, bool)`

GetMaxPrefixLimitOk returns a tuple with the MaxPrefixLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefixLimit

`func (o *ExchangeServiceRequestBgp) SetMaxPrefixLimit(v int64)`

SetMaxPrefixLimit sets MaxPrefixLimit field to given value.

### HasMaxPrefixLimit

`func (o *ExchangeServiceRequestBgp) HasMaxPrefixLimit() bool`

HasMaxPrefixLimit returns a boolean if a field has been set.

### GetPrependSelfEnabled

`func (o *ExchangeServiceRequestBgp) GetPrependSelfEnabled() bool`

GetPrependSelfEnabled returns the PrependSelfEnabled field if non-nil, zero value otherwise.

### GetPrependSelfEnabledOk

`func (o *ExchangeServiceRequestBgp) GetPrependSelfEnabledOk() (*bool, bool)`

GetPrependSelfEnabledOk returns a tuple with the PrependSelfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependSelfEnabled

`func (o *ExchangeServiceRequestBgp) SetPrependSelfEnabled(v bool)`

SetPrependSelfEnabled sets PrependSelfEnabled field to given value.

### HasPrependSelfEnabled

`func (o *ExchangeServiceRequestBgp) HasPrependSelfEnabled() bool`

HasPrependSelfEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


