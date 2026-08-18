# AppLinkAppServiceAttachmentChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | Uniquely identifies a change | [optional] 
**Type** | [**AppLinkChangeType**](AppLinkChangeType.md) |  | 
**Status** | Pointer to [**PortChangeStatus**](PortChangeStatus.md) |  | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedDateTime** | **time.Time** | Set when change object is updated | 
**Data** | Pointer to [**[]AppLinkChangeOperation**](AppLinkChangeOperation.md) |  | [optional] 

## Methods

### NewAppLinkAppServiceAttachmentChange

`func NewAppLinkAppServiceAttachmentChange(type_ AppLinkChangeType, updatedDateTime time.Time, ) *AppLinkAppServiceAttachmentChange`

NewAppLinkAppServiceAttachmentChange instantiates a new AppLinkAppServiceAttachmentChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAppServiceAttachmentChangeWithDefaults

`func NewAppLinkAppServiceAttachmentChangeWithDefaults() *AppLinkAppServiceAttachmentChange`

NewAppLinkAppServiceAttachmentChangeWithDefaults instantiates a new AppLinkAppServiceAttachmentChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppLinkAppServiceAttachmentChange) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAppServiceAttachmentChange) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAppServiceAttachmentChange) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AppLinkAppServiceAttachmentChange) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAppServiceAttachmentChange) GetType() AppLinkChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAppServiceAttachmentChange) GetTypeOk() (*AppLinkChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAppServiceAttachmentChange) SetType(v AppLinkChangeType)`

SetType sets Type field to given value.


### GetStatus

`func (o *AppLinkAppServiceAttachmentChange) GetStatus() PortChangeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AppLinkAppServiceAttachmentChange) GetStatusOk() (*PortChangeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AppLinkAppServiceAttachmentChange) SetStatus(v PortChangeStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AppLinkAppServiceAttachmentChange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *AppLinkAppServiceAttachmentChange) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AppLinkAppServiceAttachmentChange) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AppLinkAppServiceAttachmentChange) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AppLinkAppServiceAttachmentChange) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *AppLinkAppServiceAttachmentChange) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *AppLinkAppServiceAttachmentChange) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *AppLinkAppServiceAttachmentChange) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.


### GetData

`func (o *AppLinkAppServiceAttachmentChange) GetData() []AppLinkChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppLinkAppServiceAttachmentChange) GetDataOk() (*[]AppLinkChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppLinkAppServiceAttachmentChange) SetData(v []AppLinkChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *AppLinkAppServiceAttachmentChange) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


