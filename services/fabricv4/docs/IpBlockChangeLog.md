# IpBlockChangeLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedDateTime** | **time.Time** |  | 
**UpdatedDateTime** | Pointer to **time.Time** |  | [optional] 
**DeletedDateTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIpBlockChangeLog

`func NewIpBlockChangeLog(createdDateTime time.Time, ) *IpBlockChangeLog`

NewIpBlockChangeLog instantiates a new IpBlockChangeLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockChangeLogWithDefaults

`func NewIpBlockChangeLogWithDefaults() *IpBlockChangeLog`

NewIpBlockChangeLogWithDefaults instantiates a new IpBlockChangeLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDateTime

`func (o *IpBlockChangeLog) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *IpBlockChangeLog) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *IpBlockChangeLog) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.


### GetUpdatedDateTime

`func (o *IpBlockChangeLog) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *IpBlockChangeLog) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *IpBlockChangeLog) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *IpBlockChangeLog) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetDeletedDateTime

`func (o *IpBlockChangeLog) GetDeletedDateTime() time.Time`

GetDeletedDateTime returns the DeletedDateTime field if non-nil, zero value otherwise.

### GetDeletedDateTimeOk

`func (o *IpBlockChangeLog) GetDeletedDateTimeOk() (*time.Time, bool)`

GetDeletedDateTimeOk returns a tuple with the DeletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedDateTime

`func (o *IpBlockChangeLog) SetDeletedDateTime(v time.Time)`

SetDeletedDateTime sets DeletedDateTime field to given value.

### HasDeletedDateTime

`func (o *IpBlockChangeLog) HasDeletedDateTime() bool`

HasDeletedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


