# ChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **string** | Account number of the account used for object creation | 
**CreatedByFullName** | **string** | Full name of the account used for object creation | 
**CreatedByEmail** | **string** | E-mail of the account used for object creation | 
**CreatedDateTime** | **time.Time** | Date and time of object creation | 
**UpdatedBy** | **string** | Account number of the account that updated the object last time | 
**UpdatedByFullName** | **string** | Full name of the account that updated the object last time | 
**UpdatedByEmail** | **string** | E-mail of the account that updated the object last time | 
**UpdatedDateTime** | **time.Time** | Date and time of the account that updated the object last time | 
**DeletedBy** | Pointer to **string** | Account number of the account that updated the object last time | [optional] 
**DeletedByFullName** | Pointer to **string** | Full name of the account that updated the object last time | [optional] 
**DeletedByEmail** | Pointer to **string** | E-mail of the account that updated the object last time | [optional] 
**DeletedDateTime** | Pointer to **time.Time** | Date and time of the account that updated the object last time | [optional] 

## Methods

### NewChangeLog

`func NewChangeLog(createdBy string, createdByFullName string, createdByEmail string, createdDateTime time.Time, updatedBy string, updatedByFullName string, updatedByEmail string, updatedDateTime time.Time, ) *ChangeLog`

NewChangeLog instantiates a new ChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeLogWithDefaults

`func NewChangeLogWithDefaults() *ChangeLog`

NewChangeLogWithDefaults instantiates a new ChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ChangeLog) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ChangeLog) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ChangeLog) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedByFullName

`func (o *ChangeLog) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *ChangeLog) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *ChangeLog) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.


### GetCreatedByEmail

`func (o *ChangeLog) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *ChangeLog) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *ChangeLog) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.


### GetCreatedDateTime

`func (o *ChangeLog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ChangeLog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ChangeLog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.


### GetUpdatedBy

`func (o *ChangeLog) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ChangeLog) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ChangeLog) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetUpdatedByFullName

`func (o *ChangeLog) GetUpdatedByFullName() string`

GetUpdatedByFullName returns the UpdatedByFullName field if non-nil, zero value otherwise.

### GetUpdatedByFullNameOk

`func (o *ChangeLog) GetUpdatedByFullNameOk() (*string, bool)`

GetUpdatedByFullNameOk returns a tuple with the UpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByFullName

`func (o *ChangeLog) SetUpdatedByFullName(v string)`

SetUpdatedByFullName sets UpdatedByFullName field to given value.


### GetUpdatedByEmail

`func (o *ChangeLog) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *ChangeLog) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *ChangeLog) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.


### GetUpdatedDateTime

`func (o *ChangeLog) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *ChangeLog) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *ChangeLog) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetDeletedBy

`func (o *ChangeLog) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ChangeLog) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ChangeLog) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ChangeLog) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByFullName

`func (o *ChangeLog) GetDeletedByFullName() string`

GetDeletedByFullName returns the DeletedByFullName field if non-nil, zero value otherwise.

### GetDeletedByFullNameOk

`func (o *ChangeLog) GetDeletedByFullNameOk() (*string, bool)`

GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByFullName

`func (o *ChangeLog) SetDeletedByFullName(v string)`

SetDeletedByFullName sets DeletedByFullName field to given value.

### HasDeletedByFullName

`func (o *ChangeLog) HasDeletedByFullName() bool`

HasDeletedByFullName returns a boolean if a field has been set.

### GetDeletedByEmail

`func (o *ChangeLog) GetDeletedByEmail() string`

GetDeletedByEmail returns the DeletedByEmail field if non-nil, zero value otherwise.

### GetDeletedByEmailOk

`func (o *ChangeLog) GetDeletedByEmailOk() (*string, bool)`

GetDeletedByEmailOk returns a tuple with the DeletedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByEmail

`func (o *ChangeLog) SetDeletedByEmail(v string)`

SetDeletedByEmail sets DeletedByEmail field to given value.

### HasDeletedByEmail

`func (o *ChangeLog) HasDeletedByEmail() bool`

HasDeletedByEmail returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *ChangeLog) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *ChangeLog) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *ChangeLog) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *ChangeLog) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


