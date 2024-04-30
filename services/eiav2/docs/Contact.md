# Contact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | **string** |  | 
**Type** | [**ContactType**](ContactType.md) |  | 
**RegisteredUser** | **string** | Identifies (e.g., userName) a registered user. If a registered user is specified, then firstName/lastName need not be provided  | 

## Methods

### NewContact

`func NewContact(href string, type_ ContactType, registeredUser string, ) *Contact`

NewContact instantiates a new Contact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactWithDefaults

`func NewContactWithDefaults() *Contact`

NewContactWithDefaults instantiates a new Contact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Contact) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Contact) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Contact) SetHref(v string)`

SetHref sets Href field to given value.


### GetType

`func (o *Contact) GetType() ContactType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Contact) GetTypeOk() (*ContactType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Contact) SetType(v ContactType)`

SetType sets Type field to given value.


### GetRegisteredUser

`func (o *Contact) GetRegisteredUser() string`

GetRegisteredUser returns the RegisteredUser field if non-nil, zero value otherwise.

### GetRegisteredUserOk

`func (o *Contact) GetRegisteredUserOk() (*string, bool)`

GetRegisteredUserOk returns a tuple with the RegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUser

`func (o *Contact) SetRegisteredUser(v string)`

SetRegisteredUser sets RegisteredUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


