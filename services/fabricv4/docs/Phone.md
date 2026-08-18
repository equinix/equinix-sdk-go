# Phone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**PhoneType**](PhoneType.md) |  | [optional] 
**Number** | Pointer to **string** | Phone number | [optional] 
**ContactPreference** | Pointer to [**ContactPreference**](ContactPreference.md) |  | [optional] 

## Methods

### NewPhone

`func NewPhone() *Phone`

NewPhone instantiates a new Phone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneWithDefaults

`func NewPhoneWithDefaults() *Phone`

NewPhoneWithDefaults instantiates a new Phone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Phone) GetType() PhoneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Phone) GetTypeOk() (*PhoneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Phone) SetType(v PhoneType)`

SetType sets Type field to given value.

### HasType

`func (o *Phone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *Phone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Phone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Phone) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Phone) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetContactPreference

`func (o *Phone) GetContactPreference() ContactPreference`

GetContactPreference returns the ContactPreference field if non-nil, zero value otherwise.

### GetContactPreferenceOk

`func (o *Phone) GetContactPreferenceOk() (*ContactPreference, bool)`

GetContactPreferenceOk returns a tuple with the ContactPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPreference

`func (o *Phone) SetContactPreference(v ContactPreference)`

SetContactPreference sets ContactPreference field to given value.

### HasContactPreference

`func (o *Phone) HasContactPreference() bool`

HasContactPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


