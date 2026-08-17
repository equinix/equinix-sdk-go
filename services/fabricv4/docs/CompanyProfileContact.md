# CompanyProfileContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CompanyProfileContactType**](CompanyProfileContactType.md) |  | [optional] 
**Contacts** | Pointer to [**[]CompanyProfileContactContacts**](CompanyProfileContactContacts.md) |  | [optional] 

## Methods

### NewCompanyProfileContact

`func NewCompanyProfileContact() *CompanyProfileContact`

NewCompanyProfileContact instantiates a new CompanyProfileContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyProfileContactWithDefaults

`func NewCompanyProfileContactWithDefaults() *CompanyProfileContact`

NewCompanyProfileContactWithDefaults instantiates a new CompanyProfileContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CompanyProfileContact) GetType() CompanyProfileContactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyProfileContact) GetTypeOk() (*CompanyProfileContactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyProfileContact) SetType(v CompanyProfileContactType)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyProfileContact) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContacts

`func (o *CompanyProfileContact) GetContacts() []CompanyProfileContactContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *CompanyProfileContact) GetContactsOk() (*[]CompanyProfileContactContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *CompanyProfileContact) SetContacts(v []CompanyProfileContactContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *CompanyProfileContact) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


