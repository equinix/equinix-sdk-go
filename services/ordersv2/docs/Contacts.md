# Contacts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisteredUser** | Pointer to **string** | Username of a registered user. | [optional] 
**FirstName** | Pointer to **string** | First name of the contact. | [optional] 
**LastName** | Pointer to **string** | Last name of the contact. | [optional] 
**Type** | Pointer to [**ContactsType**](ContactsType.md) |  | [optional] 
**Availability** | Pointer to [**ContactsAvailability**](ContactsAvailability.md) |  | [optional] [default to CONTACTSAVAILABILITY_ANYTIME]
**Timezone** | Pointer to [**TIMEZONELIST**](TIMEZONELIST.md) |  | [optional] 
**Details** | Pointer to [**[]ContactsDetails**](ContactsDetails.md) | Reference of the related party, could be a party reference or a party role reference | [optional] 

## Methods

### NewContacts

`func NewContacts() *Contacts`

NewContacts instantiates a new Contacts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactsWithDefaults

`func NewContactsWithDefaults() *Contacts`

NewContactsWithDefaults instantiates a new Contacts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisteredUser

`func (o *Contacts) GetRegisteredUser() string`

GetRegisteredUser returns the RegisteredUser field if non-nil, zero value otherwise.

### GetRegisteredUserOk

`func (o *Contacts) GetRegisteredUserOk() (*string, bool)`

GetRegisteredUserOk returns a tuple with the RegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUser

`func (o *Contacts) SetRegisteredUser(v string)`

SetRegisteredUser sets RegisteredUser field to given value.

### HasRegisteredUser

`func (o *Contacts) HasRegisteredUser() bool`

HasRegisteredUser returns a boolean if a field has been set.

### GetFirstName

`func (o *Contacts) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Contacts) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Contacts) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Contacts) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Contacts) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Contacts) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Contacts) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Contacts) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetType

`func (o *Contacts) GetType() ContactsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contacts) GetTypeOk() (*ContactsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contacts) SetType(v ContactsType)`

SetType sets Type field to given value.

### HasType

`func (o *Contacts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAvailability

`func (o *Contacts) GetAvailability() ContactsAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *Contacts) GetAvailabilityOk() (*ContactsAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *Contacts) SetAvailability(v ContactsAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *Contacts) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetTimezone

`func (o *Contacts) GetTimezone() TIMEZONELIST`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Contacts) GetTimezoneOk() (*TIMEZONELIST, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Contacts) SetTimezone(v TIMEZONELIST)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Contacts) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDetails

`func (o *Contacts) GetDetails() []ContactsDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Contacts) GetDetailsOk() (*[]ContactsDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Contacts) SetDetails(v []ContactsDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Contacts) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


