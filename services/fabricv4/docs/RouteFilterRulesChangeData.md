# RouteFilterRulesChangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | Current outcome of the change flow | [optional] 
**CreatedBy** | Pointer to **string** | Created by User Key | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Set when change flow starts | [optional] 
**UpdatedBy** | Pointer to **string** | Updated by User Key | [optional] 
**UpdatedDateTime** | Pointer to **time.Time** | Set when change object is updated | [optional] 
**Data** | Pointer to [**RouteFilterRulesChangeOperation**](RouteFilterRulesChangeOperation.md) |  | [optional] 
**Uuid** | **string** | Uniquely identifies a change | 
**Type** | [**RouteFilterRulesChangeType**](RouteFilterRulesChangeType.md) |  | 
**Href** | Pointer to **string** | Route Filter Change URI | [optional] 

## Methods

### NewRouteFilterRulesChangeData

`func NewRouteFilterRulesChangeData(uuid string, type_ RouteFilterRulesChangeType, ) *RouteFilterRulesChangeData`

NewRouteFilterRulesChangeData instantiates a new RouteFilterRulesChangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteFilterRulesChangeDataWithDefaults

`func NewRouteFilterRulesChangeDataWithDefaults() *RouteFilterRulesChangeData`

NewRouteFilterRulesChangeDataWithDefaults instantiates a new RouteFilterRulesChangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RouteFilterRulesChangeData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RouteFilterRulesChangeData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RouteFilterRulesChangeData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RouteFilterRulesChangeData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RouteFilterRulesChangeData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RouteFilterRulesChangeData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RouteFilterRulesChangeData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *RouteFilterRulesChangeData) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *RouteFilterRulesChangeData) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *RouteFilterRulesChangeData) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *RouteFilterRulesChangeData) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *RouteFilterRulesChangeData) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *RouteFilterRulesChangeData) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *RouteFilterRulesChangeData) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *RouteFilterRulesChangeData) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *RouteFilterRulesChangeData) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedDateTime

`func (o *RouteFilterRulesChangeData) GetUpdatedDateTime() time.Time`

GetUpdatedDateTime returns the UpdatedDateTime field if non-nil, zero value otherwise.

### GetUpdatedDateTimeOk

`func (o *RouteFilterRulesChangeData) GetUpdatedDateTimeOk() (*time.Time, bool)`

GetUpdatedDateTimeOk returns a tuple with the UpdatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDateTime

`func (o *RouteFilterRulesChangeData) SetUpdatedDateTime(v time.Time)`

SetUpdatedDateTime sets UpdatedDateTime field to given value.

### HasUpdatedDateTime

`func (o *RouteFilterRulesChangeData) HasUpdatedDateTime() bool`

HasUpdatedDateTime returns a boolean if a field has been set.

### GetData

`func (o *RouteFilterRulesChangeData) GetData() RouteFilterRulesChangeOperation`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RouteFilterRulesChangeData) GetDataOk() (*RouteFilterRulesChangeOperation, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RouteFilterRulesChangeData) SetData(v RouteFilterRulesChangeOperation)`

SetData sets Data field to given value.

### HasData

`func (o *RouteFilterRulesChangeData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetUuid

`func (o *RouteFilterRulesChangeData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RouteFilterRulesChangeData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RouteFilterRulesChangeData) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *RouteFilterRulesChangeData) GetType() RouteFilterRulesChangeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteFilterRulesChangeData) GetTypeOk() (*RouteFilterRulesChangeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteFilterRulesChangeData) SetType(v RouteFilterRulesChangeType)`

SetType sets Type field to given value.


### GetHref

`func (o *RouteFilterRulesChangeData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RouteFilterRulesChangeData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RouteFilterRulesChangeData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RouteFilterRulesChangeData) HasHref() bool`

HasHref returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


