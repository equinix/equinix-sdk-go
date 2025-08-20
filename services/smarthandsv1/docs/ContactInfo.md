# ContactInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactType** | [**ContactInfoContactType**](ContactInfoContactType.md) |  | 
**UserName** | Pointer to **string** | User Name | [optional] 
**Name** | Pointer to **string** | Full Name of the Contact, Eg. &#39;First_Name&#39; + &#39;Last_Name&#39; | [optional] 
**Email** | Pointer to **string** | Valid email address of the contact | [optional] 
**WorkPhoneCountryCode** | Pointer to **string** | Country Code of Primary Phone of the contact | [optional] 
**WorkPhone** | Pointer to **string** | Primary Phone of the contact | [optional] 
**WorkPhonePrefToCall** | Pointer to [**ContactInfoWorkPhonePrefToCall**](ContactInfoWorkPhonePrefToCall.md) |  | [optional] 
**MobilePhoneCountryCode** | Pointer to **string** | Country Code of Mobile Phone of the contact | [optional] 
**MobilePhone** | Pointer to **string** | Mobile Phone of the contact | [optional] 
**MobilePhonePrefToCall** | Pointer to [**ContactInfoWorkPhonePrefToCall**](ContactInfoWorkPhonePrefToCall.md) |  | [optional] 
**WorkPhoneTimeZone** | Pointer to **string** | Work Phone TimeZone | [optional] 
**MobilePhoneTimeZone** | Pointer to **string** | Mobile Phone TimeZone | [optional] 

## Methods

### NewContactInfo

`func NewContactInfo(contactType ContactInfoContactType, ) *ContactInfo`

NewContactInfo instantiates a new ContactInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactInfoWithDefaults

`func NewContactInfoWithDefaults() *ContactInfo`

NewContactInfoWithDefaults instantiates a new ContactInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactType

`func (o *ContactInfo) GetContactType() ContactInfoContactType`

GetContactType returns the ContactType field if non-nil, zero value otherwise.

### GetContactTypeOk

`func (o *ContactInfo) GetContactTypeOk() (*ContactInfoContactType, bool)`

GetContactTypeOk returns a tuple with the ContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactType

`func (o *ContactInfo) SetContactType(v ContactInfoContactType)`

SetContactType sets ContactType field to given value.


### GetUserName

`func (o *ContactInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *ContactInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *ContactInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *ContactInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetName

`func (o *ContactInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ContactInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ContactInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ContactInfo) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ContactInfo) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWorkPhoneCountryCode

`func (o *ContactInfo) GetWorkPhoneCountryCode() string`

GetWorkPhoneCountryCode returns the WorkPhoneCountryCode field if non-nil, zero value otherwise.

### GetWorkPhoneCountryCodeOk

`func (o *ContactInfo) GetWorkPhoneCountryCodeOk() (*string, bool)`

GetWorkPhoneCountryCodeOk returns a tuple with the WorkPhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhoneCountryCode

`func (o *ContactInfo) SetWorkPhoneCountryCode(v string)`

SetWorkPhoneCountryCode sets WorkPhoneCountryCode field to given value.

### HasWorkPhoneCountryCode

`func (o *ContactInfo) HasWorkPhoneCountryCode() bool`

HasWorkPhoneCountryCode returns a boolean if a field has been set.

### GetWorkPhone

`func (o *ContactInfo) GetWorkPhone() string`

GetWorkPhone returns the WorkPhone field if non-nil, zero value otherwise.

### GetWorkPhoneOk

`func (o *ContactInfo) GetWorkPhoneOk() (*string, bool)`

GetWorkPhoneOk returns a tuple with the WorkPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhone

`func (o *ContactInfo) SetWorkPhone(v string)`

SetWorkPhone sets WorkPhone field to given value.

### HasWorkPhone

`func (o *ContactInfo) HasWorkPhone() bool`

HasWorkPhone returns a boolean if a field has been set.

### GetWorkPhonePrefToCall

`func (o *ContactInfo) GetWorkPhonePrefToCall() ContactInfoWorkPhonePrefToCall`

GetWorkPhonePrefToCall returns the WorkPhonePrefToCall field if non-nil, zero value otherwise.

### GetWorkPhonePrefToCallOk

`func (o *ContactInfo) GetWorkPhonePrefToCallOk() (*ContactInfoWorkPhonePrefToCall, bool)`

GetWorkPhonePrefToCallOk returns a tuple with the WorkPhonePrefToCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhonePrefToCall

`func (o *ContactInfo) SetWorkPhonePrefToCall(v ContactInfoWorkPhonePrefToCall)`

SetWorkPhonePrefToCall sets WorkPhonePrefToCall field to given value.

### HasWorkPhonePrefToCall

`func (o *ContactInfo) HasWorkPhonePrefToCall() bool`

HasWorkPhonePrefToCall returns a boolean if a field has been set.

### GetMobilePhoneCountryCode

`func (o *ContactInfo) GetMobilePhoneCountryCode() string`

GetMobilePhoneCountryCode returns the MobilePhoneCountryCode field if non-nil, zero value otherwise.

### GetMobilePhoneCountryCodeOk

`func (o *ContactInfo) GetMobilePhoneCountryCodeOk() (*string, bool)`

GetMobilePhoneCountryCodeOk returns a tuple with the MobilePhoneCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneCountryCode

`func (o *ContactInfo) SetMobilePhoneCountryCode(v string)`

SetMobilePhoneCountryCode sets MobilePhoneCountryCode field to given value.

### HasMobilePhoneCountryCode

`func (o *ContactInfo) HasMobilePhoneCountryCode() bool`

HasMobilePhoneCountryCode returns a boolean if a field has been set.

### GetMobilePhone

`func (o *ContactInfo) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *ContactInfo) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *ContactInfo) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *ContactInfo) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetMobilePhonePrefToCall

`func (o *ContactInfo) GetMobilePhonePrefToCall() ContactInfoWorkPhonePrefToCall`

GetMobilePhonePrefToCall returns the MobilePhonePrefToCall field if non-nil, zero value otherwise.

### GetMobilePhonePrefToCallOk

`func (o *ContactInfo) GetMobilePhonePrefToCallOk() (*ContactInfoWorkPhonePrefToCall, bool)`

GetMobilePhonePrefToCallOk returns a tuple with the MobilePhonePrefToCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhonePrefToCall

`func (o *ContactInfo) SetMobilePhonePrefToCall(v ContactInfoWorkPhonePrefToCall)`

SetMobilePhonePrefToCall sets MobilePhonePrefToCall field to given value.

### HasMobilePhonePrefToCall

`func (o *ContactInfo) HasMobilePhonePrefToCall() bool`

HasMobilePhonePrefToCall returns a boolean if a field has been set.

### GetWorkPhoneTimeZone

`func (o *ContactInfo) GetWorkPhoneTimeZone() string`

GetWorkPhoneTimeZone returns the WorkPhoneTimeZone field if non-nil, zero value otherwise.

### GetWorkPhoneTimeZoneOk

`func (o *ContactInfo) GetWorkPhoneTimeZoneOk() (*string, bool)`

GetWorkPhoneTimeZoneOk returns a tuple with the WorkPhoneTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkPhoneTimeZone

`func (o *ContactInfo) SetWorkPhoneTimeZone(v string)`

SetWorkPhoneTimeZone sets WorkPhoneTimeZone field to given value.

### HasWorkPhoneTimeZone

`func (o *ContactInfo) HasWorkPhoneTimeZone() bool`

HasWorkPhoneTimeZone returns a boolean if a field has been set.

### GetMobilePhoneTimeZone

`func (o *ContactInfo) GetMobilePhoneTimeZone() string`

GetMobilePhoneTimeZone returns the MobilePhoneTimeZone field if non-nil, zero value otherwise.

### GetMobilePhoneTimeZoneOk

`func (o *ContactInfo) GetMobilePhoneTimeZoneOk() (*string, bool)`

GetMobilePhoneTimeZoneOk returns a tuple with the MobilePhoneTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneTimeZone

`func (o *ContactInfo) SetMobilePhoneTimeZone(v string)`

SetMobilePhoneTimeZone sets MobilePhoneTimeZone field to given value.

### HasMobilePhoneTimeZone

`func (o *ContactInfo) HasMobilePhoneTimeZone() bool`

HasMobilePhoneTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


