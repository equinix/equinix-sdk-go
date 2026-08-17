# ExchangeServiceResponseBgp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerPeerIp** | Pointer to **string** |  | [optional] 
**DomainName** | Pointer to **string** |  | [optional] 
**PrimaryRouteServerIp** | Pointer to **string** |  | [optional] 
**SecondaryRouteServerIp** | Pointer to **string** |  | [optional] 
**AsSet** | Pointer to **string** |  | [optional] 
**MlpeEnabled** | Pointer to **bool** |  | [optional] 
**Md5AuthKey** | Pointer to **string** |  | [optional] 
**RcMd5AuthKey** | Pointer to **string** |  | [optional] 
**Prefixes** | Pointer to **[]string** | List of IP prefix | [optional] 
**MaxPrefixLimit** | Pointer to **int64** |  | [optional] 
**PrependSelfEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewExchangeServiceResponseBgp

`func NewExchangeServiceResponseBgp() *ExchangeServiceResponseBgp`

NewExchangeServiceResponseBgp instantiates a new ExchangeServiceResponseBgp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceResponseBgpWithDefaults

`func NewExchangeServiceResponseBgpWithDefaults() *ExchangeServiceResponseBgp`

NewExchangeServiceResponseBgpWithDefaults instantiates a new ExchangeServiceResponseBgp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerPeerIp

`func (o *ExchangeServiceResponseBgp) GetCustomerPeerIp() string`

GetCustomerPeerIp returns the CustomerPeerIp field if non-nil, zero value otherwise.

### GetCustomerPeerIpOk

`func (o *ExchangeServiceResponseBgp) GetCustomerPeerIpOk() (*string, bool)`

GetCustomerPeerIpOk returns a tuple with the CustomerPeerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPeerIp

`func (o *ExchangeServiceResponseBgp) SetCustomerPeerIp(v string)`

SetCustomerPeerIp sets CustomerPeerIp field to given value.

### HasCustomerPeerIp

`func (o *ExchangeServiceResponseBgp) HasCustomerPeerIp() bool`

HasCustomerPeerIp returns a boolean if a field has been set.

### GetDomainName

`func (o *ExchangeServiceResponseBgp) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *ExchangeServiceResponseBgp) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *ExchangeServiceResponseBgp) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *ExchangeServiceResponseBgp) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetPrimaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) GetPrimaryRouteServerIp() string`

GetPrimaryRouteServerIp returns the PrimaryRouteServerIp field if non-nil, zero value otherwise.

### GetPrimaryRouteServerIpOk

`func (o *ExchangeServiceResponseBgp) GetPrimaryRouteServerIpOk() (*string, bool)`

GetPrimaryRouteServerIpOk returns a tuple with the PrimaryRouteServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) SetPrimaryRouteServerIp(v string)`

SetPrimaryRouteServerIp sets PrimaryRouteServerIp field to given value.

### HasPrimaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) HasPrimaryRouteServerIp() bool`

HasPrimaryRouteServerIp returns a boolean if a field has been set.

### GetSecondaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) GetSecondaryRouteServerIp() string`

GetSecondaryRouteServerIp returns the SecondaryRouteServerIp field if non-nil, zero value otherwise.

### GetSecondaryRouteServerIpOk

`func (o *ExchangeServiceResponseBgp) GetSecondaryRouteServerIpOk() (*string, bool)`

GetSecondaryRouteServerIpOk returns a tuple with the SecondaryRouteServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) SetSecondaryRouteServerIp(v string)`

SetSecondaryRouteServerIp sets SecondaryRouteServerIp field to given value.

### HasSecondaryRouteServerIp

`func (o *ExchangeServiceResponseBgp) HasSecondaryRouteServerIp() bool`

HasSecondaryRouteServerIp returns a boolean if a field has been set.

### GetAsSet

`func (o *ExchangeServiceResponseBgp) GetAsSet() string`

GetAsSet returns the AsSet field if non-nil, zero value otherwise.

### GetAsSetOk

`func (o *ExchangeServiceResponseBgp) GetAsSetOk() (*string, bool)`

GetAsSetOk returns a tuple with the AsSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsSet

`func (o *ExchangeServiceResponseBgp) SetAsSet(v string)`

SetAsSet sets AsSet field to given value.

### HasAsSet

`func (o *ExchangeServiceResponseBgp) HasAsSet() bool`

HasAsSet returns a boolean if a field has been set.

### GetMlpeEnabled

`func (o *ExchangeServiceResponseBgp) GetMlpeEnabled() bool`

GetMlpeEnabled returns the MlpeEnabled field if non-nil, zero value otherwise.

### GetMlpeEnabledOk

`func (o *ExchangeServiceResponseBgp) GetMlpeEnabledOk() (*bool, bool)`

GetMlpeEnabledOk returns a tuple with the MlpeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlpeEnabled

`func (o *ExchangeServiceResponseBgp) SetMlpeEnabled(v bool)`

SetMlpeEnabled sets MlpeEnabled field to given value.

### HasMlpeEnabled

`func (o *ExchangeServiceResponseBgp) HasMlpeEnabled() bool`

HasMlpeEnabled returns a boolean if a field has been set.

### GetMd5AuthKey

`func (o *ExchangeServiceResponseBgp) GetMd5AuthKey() string`

GetMd5AuthKey returns the Md5AuthKey field if non-nil, zero value otherwise.

### GetMd5AuthKeyOk

`func (o *ExchangeServiceResponseBgp) GetMd5AuthKeyOk() (*string, bool)`

GetMd5AuthKeyOk returns a tuple with the Md5AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5AuthKey

`func (o *ExchangeServiceResponseBgp) SetMd5AuthKey(v string)`

SetMd5AuthKey sets Md5AuthKey field to given value.

### HasMd5AuthKey

`func (o *ExchangeServiceResponseBgp) HasMd5AuthKey() bool`

HasMd5AuthKey returns a boolean if a field has been set.

### GetRcMd5AuthKey

`func (o *ExchangeServiceResponseBgp) GetRcMd5AuthKey() string`

GetRcMd5AuthKey returns the RcMd5AuthKey field if non-nil, zero value otherwise.

### GetRcMd5AuthKeyOk

`func (o *ExchangeServiceResponseBgp) GetRcMd5AuthKeyOk() (*string, bool)`

GetRcMd5AuthKeyOk returns a tuple with the RcMd5AuthKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcMd5AuthKey

`func (o *ExchangeServiceResponseBgp) SetRcMd5AuthKey(v string)`

SetRcMd5AuthKey sets RcMd5AuthKey field to given value.

### HasRcMd5AuthKey

`func (o *ExchangeServiceResponseBgp) HasRcMd5AuthKey() bool`

HasRcMd5AuthKey returns a boolean if a field has been set.

### GetPrefixes

`func (o *ExchangeServiceResponseBgp) GetPrefixes() []string`

GetPrefixes returns the Prefixes field if non-nil, zero value otherwise.

### GetPrefixesOk

`func (o *ExchangeServiceResponseBgp) GetPrefixesOk() (*[]string, bool)`

GetPrefixesOk returns a tuple with the Prefixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixes

`func (o *ExchangeServiceResponseBgp) SetPrefixes(v []string)`

SetPrefixes sets Prefixes field to given value.

### HasPrefixes

`func (o *ExchangeServiceResponseBgp) HasPrefixes() bool`

HasPrefixes returns a boolean if a field has been set.

### GetMaxPrefixLimit

`func (o *ExchangeServiceResponseBgp) GetMaxPrefixLimit() int64`

GetMaxPrefixLimit returns the MaxPrefixLimit field if non-nil, zero value otherwise.

### GetMaxPrefixLimitOk

`func (o *ExchangeServiceResponseBgp) GetMaxPrefixLimitOk() (*int64, bool)`

GetMaxPrefixLimitOk returns a tuple with the MaxPrefixLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrefixLimit

`func (o *ExchangeServiceResponseBgp) SetMaxPrefixLimit(v int64)`

SetMaxPrefixLimit sets MaxPrefixLimit field to given value.

### HasMaxPrefixLimit

`func (o *ExchangeServiceResponseBgp) HasMaxPrefixLimit() bool`

HasMaxPrefixLimit returns a boolean if a field has been set.

### GetPrependSelfEnabled

`func (o *ExchangeServiceResponseBgp) GetPrependSelfEnabled() bool`

GetPrependSelfEnabled returns the PrependSelfEnabled field if non-nil, zero value otherwise.

### GetPrependSelfEnabledOk

`func (o *ExchangeServiceResponseBgp) GetPrependSelfEnabledOk() (*bool, bool)`

GetPrependSelfEnabledOk returns a tuple with the PrependSelfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrependSelfEnabled

`func (o *ExchangeServiceResponseBgp) SetPrependSelfEnabled(v bool)`

SetPrependSelfEnabled sets PrependSelfEnabled field to given value.

### HasPrependSelfEnabled

`func (o *ExchangeServiceResponseBgp) HasPrependSelfEnabled() bool`

HasPrependSelfEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


