# CrossConnectPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CrossConnectPhoneType**](CrossConnectPhoneType.md) |  | [optional] 
**Number** | Pointer to **string** | Phone number. | [optional] 
**ContactPreference** | Pointer to [**CrossConnectContactPreference**](CrossConnectContactPreference.md) |  | [optional] 

## Methods

### NewCrossConnectPhone

`func NewCrossConnectPhone() *CrossConnectPhone`

NewCrossConnectPhone instantiates a new CrossConnectPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectPhoneWithDefaults

`func NewCrossConnectPhoneWithDefaults() *CrossConnectPhone`

NewCrossConnectPhoneWithDefaults instantiates a new CrossConnectPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CrossConnectPhone) GetType() CrossConnectPhoneType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectPhone) GetTypeOk() (*CrossConnectPhoneType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectPhone) SetType(v CrossConnectPhoneType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectPhone) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *CrossConnectPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *CrossConnectPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *CrossConnectPhone) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *CrossConnectPhone) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetContactPreference

`func (o *CrossConnectPhone) GetContactPreference() CrossConnectContactPreference`

GetContactPreference returns the ContactPreference field if non-nil, zero value otherwise.

### GetContactPreferenceOk

`func (o *CrossConnectPhone) GetContactPreferenceOk() (*CrossConnectContactPreference, bool)`

GetContactPreferenceOk returns a tuple with the ContactPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactPreference

`func (o *CrossConnectPhone) SetContactPreference(v CrossConnectContactPreference)`

SetContactPreference sets ContactPreference field to given value.

### HasContactPreference

`func (o *CrossConnectPhone) HasContactPreference() bool`

HasContactPreference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


