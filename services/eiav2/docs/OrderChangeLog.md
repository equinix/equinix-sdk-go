# OrderChangeLog

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

### NewOrderChangeLog

`func NewOrderChangeLog(createdBy string, createdByFullName string, createdByEmail string, createdDateTime time.Time, updatedBy string, updatedByFullName string, updatedByEmail string, updatedDateTime time.Time, ) *OrderChangeLog`

NewOrderChangeLog instantiates a new OrderChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderChangeLogWithDefaults

`func NewOrderChangeLogWithDefaults() *OrderChangeLog`

NewOrderChangeLogWithDefaults instantiates a new OrderChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedBy

`func (o *OrderChangeLog) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *OrderChangeLog) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *OrderChangeLog) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedByFullName

`func (o *OrderChangeLog) GetCreatedByFullName() string`

GetCreatedByFullName returns the CreatedByFullName field if non-nil, zero value otherwise.

### GetCreatedByFullNameOk

`func (o *OrderChangeLog) GetCreatedByFullNameOk() (*string, bool)`

GetCreatedByFullNameOk returns a tuple with the CreatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByFullName

`func (o *OrderChangeLog) SetCreatedByFullName(v string)`

SetCreatedByFullName sets CreatedByFullName field to given value.


### GetCreatedByEmail

`func (o *OrderChangeLog) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *OrderChangeLog) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *OrderChangeLog) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.


### GetCreatedDateTime

`func (o *OrderChangeLog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *OrderChangeLog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *OrderChangeLog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.


### GetUpdatedBy

`func (o *OrderChangeLog) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *OrderChangeLog) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *OrderChangeLog) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.


### GetUpdatedByFullName

`func (o *OrderChangeLog) GetUpdatedByFullName() string`

GetUpdatedByFullName returns the UpdatedByFullName field if non-nil, zero value otherwise.

### GetUpdatedByFullNameOk

`func (o *OrderChangeLog) GetUpdatedByFullNameOk() (*string, bool)`

GetUpdatedByFullNameOk returns a tuple with the UpdatedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByFullName

`func (o *OrderChangeLog) SetUpdatedByFullName(v string)`

SetUpdatedByFullName sets UpdatedByFullName field to given value.


### GetUpdatedByEmail

`func (o *OrderChangeLog) GetUpdatedByEmail() string`

GetUpdatedByEmail returns the UpdatedByEmail field if non-nil, zero value otherwise.

### GetUpdatedByEmailOk

`func (o *OrderChangeLog) GetUpdatedByEmailOk() (*string, bool)`

GetUpdatedByEmailOk returns a tuple with the UpdatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedByEmail

`func (o *OrderChangeLog) SetUpdatedByEmail(v string)`

SetUpdatedByEmail sets UpdatedByEmail field to given value.


### GetUpdatedDateTime

`func (o *OrderChangeLog) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *OrderChangeLog) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *OrderChangeLog) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetDeletedBy

`func (o *OrderChangeLog) GetDeletedBy() string`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *OrderChangeLog) GetDeletedByOk() (*string, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *OrderChangeLog) SetDeletedBy(v string)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *OrderChangeLog) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedByFullName

`func (o *OrderChangeLog) GetDeletedByFullName() string`

GetDeletedByFullName returns the DeletedByFullName field if non-nil, zero value otherwise.

### GetDeletedByFullNameOk

`func (o *OrderChangeLog) GetDeletedByFullNameOk() (*string, bool)`

GetDeletedByFullNameOk returns a tuple with the DeletedByFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByFullName

`func (o *OrderChangeLog) SetDeletedByFullName(v string)`

SetDeletedByFullName sets DeletedByFullName field to given value.

### HasDeletedByFullName

`func (o *OrderChangeLog) HasDeletedByFullName() bool`

HasDeletedByFullName returns a boolean if a field has been set.

### GetDeletedByEmail

`func (o *OrderChangeLog) GetDeletedByEmail() string`

GetDeletedByEmail returns the DeletedByEmail field if non-nil, zero value otherwise.

### GetDeletedByEmailOk

`func (o *OrderChangeLog) GetDeletedByEmailOk() (*string, bool)`

GetDeletedByEmailOk returns a tuple with the DeletedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedByEmail

`func (o *OrderChangeLog) SetDeletedByEmail(v string)`

SetDeletedByEmail sets DeletedByEmail field to given value.

### HasDeletedByEmail

`func (o *OrderChangeLog) HasDeletedByEmail() bool`

HasDeletedByEmail returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *OrderChangeLog) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *OrderChangeLog) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *OrderChangeLog) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *OrderChangeLog) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


