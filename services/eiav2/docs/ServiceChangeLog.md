# ServiceChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **string** |  | 
**CreatedByFullName** | **string** |  | 
**CreatedByEmail** | **string** |  | 
**CreatedDateTime** | **time.Time** |  | 
**UpdatedBy** | **string** |  | 
**UpdatedByFullName** | **string** |  | 
**UpdatedByEmail** | **string** |  | 
**UpdatedDateTime** | **time.Time** |  | 
**DeletedBy** | Pointer to **string** |  | [optional] 
**DeletedByFullName** | Pointer to **string** |  | [optional] 
**DeletedByEmail** | Pointer to **string** |  | [optional] 
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewServiceChangeLog

`func NewServiceChangeLog(createdBy string, createdByFullName string, createdByEmail string, createdDateTime time.Time, updatedBy string, updatedByFullName string, updatedByEmail string, updatedDateTime time.Time, ) *ServiceChangeLog`

NewServiceChangeLog instantiates a new ServiceChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceChangeLogWithDefaults

`func NewServiceChangeLogWithDefaults() *ServiceChangeLog`

NewServiceChangeLogWithDefaults instantiates a new ServiceChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *ServiceChangeLog) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ServiceChangeLog) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ServiceChangeLog) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedByFullName

`func (o *ServiceChangeLog) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *ServiceChangeLog) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *ServiceChangeLog) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.


### GetCreatedByEmail

`func (o *ServiceChangeLog) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *ServiceChangeLog) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *ServiceChangeLog) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.


### GetCreatedDateTime

`func (o *ServiceChangeLog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ServiceChangeLog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ServiceChangeLog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.


### GetUpdatedBy

`func (o *ServiceChangeLog) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ServiceChangeLog) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ServiceChangeLog) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetUpdatedByFullName

`func (o *ServiceChangeLog) GetUpdatedByFullName() string`

GetUpdatedByFullName returns the UpdatedByFullName field if non-nil, zero value otherwise.

### GetUpdatedByFullNameOk

`func (o *ServiceChangeLog) GetUpdatedByFullNameOk() (*string, bool)`

GetUpdatedByFullNameOk returns a tuple with the UpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByFullName

`func (o *ServiceChangeLog) SetUpdatedByFullName(v string)`

SetUpdatedByFullName sets UpdatedByFullName field to given value.


### GetUpdatedByEmail

`func (o *ServiceChangeLog) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *ServiceChangeLog) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *ServiceChangeLog) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.


### GetUpdatedDateTime

`func (o *ServiceChangeLog) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *ServiceChangeLog) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *ServiceChangeLog) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetDeletedBy

`func (o *ServiceChangeLog) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *ServiceChangeLog) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *ServiceChangeLog) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *ServiceChangeLog) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByFullName

`func (o *ServiceChangeLog) GetDeletedByFullName() string`

GetDeletedByFullName returns the DeletedByFullName field if non-nil, zero value otherwise.

### GetDeletedByFullNameOk

`func (o *ServiceChangeLog) GetDeletedByFullNameOk() (*string, bool)`

GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByFullName

`func (o *ServiceChangeLog) SetDeletedByFullName(v string)`

SetDeletedByFullName sets DeletedByFullName field to given value.

### HasDeletedByFullName

`func (o *ServiceChangeLog) HasDeletedByFullName() bool`

HasDeletedByFullName returns a boolean if a field has been set.

### GetDeletedByEmail

`func (o *ServiceChangeLog) GetDeletedByEmail() string`

GetDeletedByEmail returns the DeletedByEmail field if non-nil, zero value otherwise.

### GetDeletedByEmailOk

`func (o *ServiceChangeLog) GetDeletedByEmailOk() (*string, bool)`

GetDeletedByEmailOk returns a tuple with the DeletedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByEmail

`func (o *ServiceChangeLog) SetDeletedByEmail(v string)`

SetDeletedByEmail sets DeletedByEmail field to given value.

### HasDeletedByEmail

`func (o *ServiceChangeLog) HasDeletedByEmail() bool`

HasDeletedByEmail returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *ServiceChangeLog) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *ServiceChangeLog) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *ServiceChangeLog) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *ServiceChangeLog) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


