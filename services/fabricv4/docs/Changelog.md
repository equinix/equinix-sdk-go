# Changelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | Pointer to **string** | Created by User Key | [optional] 
**CreatedByFullName** | Pointer to **string** | Created by User Full Name | [optional] 
**CreatedByEmail** | Pointer to **string** | Created by User Email Address | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Created by Date and Time | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by User Key | [optional] 
**UpdatedByFullName** | Pointer to **string** | Updated by User Full Name | [optional] 
**UpdatedByEmail** | Pointer to **string** | Updated by User Email Address | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Updated by Date and Time | [optional] 
**DeletedBy** | Pointer to **string** | Deleted by User Key | [optional] 
**DeletedByFullName** | Pointer to **string** | Deleted by User Full Name | [optional] 
**DeletedByEmail** | Pointer to **string** | Deleted by User Email Address | [optional] 
**DeletedDateTime** | Pointer to **time.Time** | Deleted by Date and Time | [optional] 

## Methods

### NewChangelog

`func NewChangelog() *Changelog`

NewChangelog instantiates a new Changelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangelogWithDefaults

`func NewChangelogWithDefaults() *Changelog`

NewChangelogWithDefaults instantiates a new Changelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *Changelog) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Changelog) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Changelog) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Changelog) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByFullName

`func (o *Changelog) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *Changelog) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *Changelog) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.

### HasCreatedByFullName

`func (o *Changelog) HasCreatedByFullName() bool`

HasCreatedByFullName returns a boolean if a field has been set.

### GetCreatedByEmail

`func (o *Changelog) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *Changelog) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *Changelog) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *Changelog) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Changelog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Changelog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Changelog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Changelog) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Changelog) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Changelog) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Changelog) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Changelog) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedByFullName

`func (o *Changelog) GetUpdatedByFullName() string`

GetUpdatedByFullName returns the UpdatedByFullName field if non-nil, zero value otherwise.

### GetUpdatedByFullNameOk

`func (o *Changelog) GetUpdatedByFullNameOk() (*string, bool)`

GetUpdatedByFullNameOk returns a tuple with the UpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByFullName

`func (o *Changelog) SetUpdatedByFullName(v string)`

SetUpdatedByFullName sets UpdatedByFullName field to given value.

### HasUpdatedByFullName

`func (o *Changelog) HasUpdatedByFullName() bool`

HasUpdatedByFullName returns a boolean if a field has been set.

### GetUpdatedByEmail

`func (o *Changelog) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *Changelog) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *Changelog) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.

### HasUpdatedByEmail

`func (o *Changelog) HasUpdatedByEmail() bool`

HasUpdatedByEmail returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *Changelog) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *Changelog) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *Changelog) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *Changelog) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Changelog) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Changelog) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Changelog) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Changelog) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByFullName

`func (o *Changelog) GetDeletedByFullName() string`

GetDeletedByFullName returns the DeletedByFullName field if non-nil, zero value otherwise.

### GetDeletedByFullNameOk

`func (o *Changelog) GetDeletedByFullNameOk() (*string, bool)`

GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByFullName

`func (o *Changelog) SetDeletedByFullName(v string)`

SetDeletedByFullName sets DeletedByFullName field to given value.

### HasDeletedByFullName

`func (o *Changelog) HasDeletedByFullName() bool`

HasDeletedByFullName returns a boolean if a field has been set.

### GetDeletedByEmail

`func (o *Changelog) GetDeletedByEmail() string`

GetDeletedByEmail returns the DeletedByEmail field if non-nil, zero value otherwise.

### GetDeletedByEmailOk

`func (o *Changelog) GetDeletedByEmailOk() (*string, bool)`

GetDeletedByEmailOk returns a tuple with the DeletedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByEmail

`func (o *Changelog) SetDeletedByEmail(v string)`

SetDeletedByEmail sets DeletedByEmail field to given value.

### HasDeletedByEmail

`func (o *Changelog) HasDeletedByEmail() bool`

HasDeletedByEmail returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *Changelog) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *Changelog) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *Changelog) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *Changelog) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


