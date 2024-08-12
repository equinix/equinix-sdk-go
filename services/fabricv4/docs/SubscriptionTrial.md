# SubscriptionTrial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**ExpiryDateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSubscriptionTrial

`func NewSubscriptionTrial() *SubscriptionTrial`

NewSubscriptionTrial instantiates a new SubscriptionTrial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionTrialWithDefaults

`func NewSubscriptionTrialWithDefaults() *SubscriptionTrial`

NewSubscriptionTrialWithDefaults instantiates a new SubscriptionTrial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SubscriptionTrial) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SubscriptionTrial) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SubscriptionTrial) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SubscriptionTrial) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExpiryDateTime

`func (o *SubscriptionTrial) GetExpiryDateTime() time.Time`

GetExpiryDateTime returns the ExpiryDateTime field if non-nil, zero value otherwise.

### GetExpiryDateTimeOk

`func (o *SubscriptionTrial) GetExpiryDateTimeOk() (*time.Time, bool)`

GetExpiryDateTimeOk returns a tuple with the ExpiryDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDateTime

`func (o *SubscriptionTrial) SetExpiryDateTime(v time.Time)`

SetExpiryDateTime sets ExpiryDateTime field to given value.

### HasExpiryDateTime

`func (o *SubscriptionTrial) HasExpiryDateTime() bool`

HasExpiryDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


