# ResponseUsageInputTokensDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CachedTokens** | **int32** | The number of tokens that were retrieved from the cache.  [More on prompt caching](https://platform.openai.com/docs/guides/prompt-caching).  | 

## Methods

### NewResponseUsageInputTokensDetails

`func NewResponseUsageInputTokensDetails(cachedTokens int32, ) *ResponseUsageInputTokensDetails`

NewResponseUsageInputTokensDetails instantiates a new ResponseUsageInputTokensDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseUsageInputTokensDetailsWithDefaults

`func NewResponseUsageInputTokensDetailsWithDefaults() *ResponseUsageInputTokensDetails`

NewResponseUsageInputTokensDetailsWithDefaults instantiates a new ResponseUsageInputTokensDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCachedTokens

`func (o *ResponseUsageInputTokensDetails) GetCachedTokens() int32`

GetCachedTokens returns the CachedTokens field if non-nil, zero value otherwise.

### GetCachedTokensOk

`func (o *ResponseUsageInputTokensDetails) GetCachedTokensOk() (*int32, bool)`

GetCachedTokensOk returns a tuple with the CachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTokens

`func (o *ResponseUsageInputTokensDetails) SetCachedTokens(v int32)`

SetCachedTokens sets CachedTokens field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


