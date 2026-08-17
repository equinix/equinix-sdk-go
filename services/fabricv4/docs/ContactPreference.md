# ContactPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | Timezone of the contact preference | [optional] 
**Availability** | Pointer to [**ContactPreferenceAvailability**](ContactPreferenceAvailability.md) |  | [optional] 

## Methods

### NewContactPreference

`func NewContactPreference() *ContactPreference`

NewContactPreference instantiates a new ContactPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactPreferenceWithDefaults

`func NewContactPreferenceWithDefaults() *ContactPreference`

NewContactPreferenceWithDefaults instantiates a new ContactPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *ContactPreference) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ContactPreference) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ContactPreference) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ContactPreference) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAvailability

`func (o *ContactPreference) GetAvailability() ContactPreferenceAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ContactPreference) GetAvailabilityOk() (*ContactPreferenceAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ContactPreference) SetAvailability(v ContactPreferenceAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ContactPreference) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


