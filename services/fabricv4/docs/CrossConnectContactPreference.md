# CrossConnectContactPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timezone** | Pointer to **string** | Timezone. | [optional] 
**Availability** | Pointer to [**CrossConnectContactAvailability**](CrossConnectContactAvailability.md) |  | [optional] 

## Methods

### NewCrossConnectContactPreference

`func NewCrossConnectContactPreference() *CrossConnectContactPreference`

NewCrossConnectContactPreference instantiates a new CrossConnectContactPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectContactPreferenceWithDefaults

`func NewCrossConnectContactPreferenceWithDefaults() *CrossConnectContactPreference`

NewCrossConnectContactPreferenceWithDefaults instantiates a new CrossConnectContactPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimezone

`func (o *CrossConnectContactPreference) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CrossConnectContactPreference) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CrossConnectContactPreference) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CrossConnectContactPreference) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetAvailability

`func (o *CrossConnectContactPreference) GetAvailability() CrossConnectContactAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *CrossConnectContactPreference) GetAvailabilityOk() (*CrossConnectContactAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *CrossConnectContactPreference) SetAvailability(v CrossConnectContactAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *CrossConnectContactPreference) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


