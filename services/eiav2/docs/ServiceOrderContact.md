# ServiceOrderContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ContactType**](ContactType.md) |  | 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Availability** | Pointer to [**ContactRequestAvailability**](ContactRequestAvailability.md) |  | [optional] 
**Details** | Pointer to [**[]ContactRequestDetails**](ContactRequestDetails.md) |  | [optional] 
**RegisteredUser** | Pointer to **string** | Identifies (e.g., userName) a registered user. If a registered user is specified, then firstName/lastName need not be provided  | [optional] 

## Methods

### NewServiceOrderContact

`func NewServiceOrderContact(type_ ContactType, ) *ServiceOrderContact`

NewServiceOrderContact instantiates a new ServiceOrderContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceOrderContactWithDefaults

`func NewServiceOrderContactWithDefaults() *ServiceOrderContact`

NewServiceOrderContactWithDefaults instantiates a new ServiceOrderContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceOrderContact) GetType() ContactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceOrderContact) GetTypeOk() (*ContactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceOrderContact) SetType(v ContactType)`

SetType sets Type field to given value.


### GetFirstName

`func (o *ServiceOrderContact) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ServiceOrderContact) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ServiceOrderContact) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ServiceOrderContact) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ServiceOrderContact) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ServiceOrderContact) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ServiceOrderContact) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ServiceOrderContact) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTimezone

`func (o *ServiceOrderContact) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ServiceOrderContact) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ServiceOrderContact) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ServiceOrderContact) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetNotes

`func (o *ServiceOrderContact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ServiceOrderContact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ServiceOrderContact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ServiceOrderContact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAvailability

`func (o *ServiceOrderContact) GetAvailability() ContactRequestAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ServiceOrderContact) GetAvailabilityOk() (*ContactRequestAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ServiceOrderContact) SetAvailability(v ContactRequestAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ServiceOrderContact) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetDetails

`func (o *ServiceOrderContact) GetDetails() []ContactRequestDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServiceOrderContact) GetDetailsOk() (*[]ContactRequestDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServiceOrderContact) SetDetails(v []ContactRequestDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServiceOrderContact) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetRegisteredUser

`func (o *ServiceOrderContact) GetRegisteredUser() string`

GetRegisteredUser returns the RegisteredUser field if non-nil, zero value otherwise.

### GetRegisteredUserOk

`func (o *ServiceOrderContact) GetRegisteredUserOk() (*string, bool)`

GetRegisteredUserOk returns a tuple with the RegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUser

`func (o *ServiceOrderContact) SetRegisteredUser(v string)`

SetRegisteredUser sets RegisteredUser field to given value.

### HasRegisteredUser

`func (o *ServiceOrderContact) HasRegisteredUser() bool`

HasRegisteredUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


