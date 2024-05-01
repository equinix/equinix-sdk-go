# ContactRequest

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

## Methods

### NewContactRequest

`func NewContactRequest(type_ ContactType, ) *ContactRequest`

NewContactRequest instantiates a new ContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactRequestWithDefaults

`func NewContactRequestWithDefaults() *ContactRequest`

NewContactRequestWithDefaults instantiates a new ContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ContactRequest) GetType() ContactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContactRequest) GetTypeOk() (*ContactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContactRequest) SetType(v ContactType)`

SetType sets Type field to given value.


### GetFirstName

`func (o *ContactRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ContactRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ContactRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ContactRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ContactRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ContactRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ContactRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ContactRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetTimezone

`func (o *ContactRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ContactRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ContactRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ContactRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetNotes

`func (o *ContactRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ContactRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ContactRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ContactRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAvailability

`func (o *ContactRequest) GetAvailability() ContactRequestAvailability`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *ContactRequest) GetAvailabilityOk() (*ContactRequestAvailability, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *ContactRequest) SetAvailability(v ContactRequestAvailability)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *ContactRequest) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetDetails

`func (o *ContactRequest) GetDetails() []ContactRequestDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ContactRequest) GetDetailsOk() (*[]ContactRequestDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ContactRequest) SetDetails(v []ContactRequestDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ContactRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


