# ActivationKeyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | Provider Encoded activation key | [optional] 
**ProviderId** | Pointer to **string** | AWS Connection identifier | [optional] 
**AccountId** | Pointer to **string** | Account identifier | [optional] 
**Bandwidth** | Pointer to **int32** | Bandwidth in Mbps | [optional] 
**Region** | Pointer to **string** | Cloud provider region identifier | [optional] 

## Methods

### NewActivationKeyDetails

`func NewActivationKeyDetails() *ActivationKeyDetails`

NewActivationKeyDetails instantiates a new ActivationKeyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivationKeyDetailsWithDefaults

`func NewActivationKeyDetailsWithDefaults() *ActivationKeyDetails`

NewActivationKeyDetailsWithDefaults instantiates a new ActivationKeyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *ActivationKeyDetails) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActivationKeyDetails) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActivationKeyDetails) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActivationKeyDetails) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetProviderId

`func (o *ActivationKeyDetails) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ActivationKeyDetails) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ActivationKeyDetails) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ActivationKeyDetails) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetAccountId

`func (o *ActivationKeyDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ActivationKeyDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ActivationKeyDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ActivationKeyDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetBandwidth

`func (o *ActivationKeyDetails) GetBandwidth() int32`

GetBandwidth returns the Bandwidth field if non-nil, zero value otherwise.

### GetBandwidthOk

`func (o *ActivationKeyDetails) GetBandwidthOk() (*int32, bool)`

GetBandwidthOk returns a tuple with the Bandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidth

`func (o *ActivationKeyDetails) SetBandwidth(v int32)`

SetBandwidth sets Bandwidth field to given value.

### HasBandwidth

`func (o *ActivationKeyDetails) HasBandwidth() bool`

HasBandwidth returns a boolean if a field has been set.

### GetRegion

`func (o *ActivationKeyDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ActivationKeyDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ActivationKeyDetails) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ActivationKeyDetails) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


