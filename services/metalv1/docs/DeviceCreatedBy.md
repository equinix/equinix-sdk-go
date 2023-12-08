# DeviceCreatedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarThumbUrl** | Pointer to **string** | Avatar thumbnail URL of the User | [optional] 
**CreatedAt** | Pointer to **time.Time** | When the user was created | [optional] 
**Email** | Pointer to **string** | Primary email address of the User | [optional] 
**FirstName** | Pointer to **string** | First name of the User | [optional] 
**FullName** | Pointer to **string** | Full name of the User | [optional] 
**Href** | Pointer to **string** | API URL uniquely representing the User | [optional] 
**Id** | **string** | ID of the User | 
**LastName** | Pointer to **string** | Last name of the User | [optional] 
**ShortId** | **string** | Short ID of the User | 
**UpdatedAt** | Pointer to **time.Time** | When the user details were last updated | [optional] 

## Methods

### NewDeviceCreatedBy

`func NewDeviceCreatedBy(id string, shortId string, ) *DeviceCreatedBy`

NewDeviceCreatedBy instantiates a new DeviceCreatedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreatedByWithDefaults

`func NewDeviceCreatedByWithDefaults() *DeviceCreatedBy`

NewDeviceCreatedByWithDefaults instantiates a new DeviceCreatedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarThumbUrl

`func (o *DeviceCreatedBy) GetAvatarThumbUrl() string`

GetAvatarThumbUrl returns the AvatarThumbUrl field if non-nil, zero value otherwise.

### GetAvatarThumbUrlOk

`func (o *DeviceCreatedBy) GetAvatarThumbUrlOk() (*string, bool)`

GetAvatarThumbUrlOk returns a tuple with the AvatarThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarThumbUrl

`func (o *DeviceCreatedBy) SetAvatarThumbUrl(v string)`

SetAvatarThumbUrl sets AvatarThumbUrl field to given value.

### HasAvatarThumbUrl

`func (o *DeviceCreatedBy) HasAvatarThumbUrl() bool`

HasAvatarThumbUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DeviceCreatedBy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeviceCreatedBy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeviceCreatedBy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DeviceCreatedBy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEmail

`func (o *DeviceCreatedBy) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DeviceCreatedBy) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DeviceCreatedBy) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DeviceCreatedBy) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *DeviceCreatedBy) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *DeviceCreatedBy) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *DeviceCreatedBy) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *DeviceCreatedBy) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFullName

`func (o *DeviceCreatedBy) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DeviceCreatedBy) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DeviceCreatedBy) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DeviceCreatedBy) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetHref

`func (o *DeviceCreatedBy) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DeviceCreatedBy) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DeviceCreatedBy) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DeviceCreatedBy) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *DeviceCreatedBy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceCreatedBy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceCreatedBy) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *DeviceCreatedBy) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *DeviceCreatedBy) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *DeviceCreatedBy) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *DeviceCreatedBy) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetShortId

`func (o *DeviceCreatedBy) GetShortId() string`

GetShortId returns the ShortId field if non-nil, zero value otherwise.

### GetShortIdOk

`func (o *DeviceCreatedBy) GetShortIdOk() (*string, bool)`

GetShortIdOk returns a tuple with the ShortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortId

`func (o *DeviceCreatedBy) SetShortId(v string)`

SetShortId sets ShortId field to given value.


### GetUpdatedAt

`func (o *DeviceCreatedBy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeviceCreatedBy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeviceCreatedBy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeviceCreatedBy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


