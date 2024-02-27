# ConnectionAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**Actions**](Actions.md) |  | 
**Href** | **string** | Connection action URI | [readonly] 
**Uuid** | **string** | Equinix-assigned connection identifier | 
**Description** | Pointer to **string** | Connection rejection reason detail | [optional] 
**Data** | [**ConnectionAcceptanceData**](ConnectionAcceptanceData.md) |  | 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewConnectionAction

`func NewConnectionAction(type_ Actions, href string, uuid string, data ConnectionAcceptanceData, ) *ConnectionAction`

NewConnectionAction instantiates a new ConnectionAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionActionWithDefaults

`func NewConnectionActionWithDefaults() *ConnectionAction`

NewConnectionActionWithDefaults instantiates a new ConnectionAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionAction) GetType() Actions`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionAction) GetTypeOk() (*Actions, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionAction) SetType(v Actions)`

SetType sets Type field to given value.


### GetHref

`func (o *ConnectionAction) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectionAction) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectionAction) SetHref(v string)`

SetHref sets Href field to given value.


### GetUuid

`func (o *ConnectionAction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConnectionAction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConnectionAction) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetDescription

`func (o *ConnectionAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetData

`func (o *ConnectionAction) GetData() ConnectionAcceptanceData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ConnectionAction) GetDataOk() (*ConnectionAcceptanceData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ConnectionAction) SetData(v ConnectionAcceptanceData)`

SetData sets Data field to given value.


### GetChangeLog

`func (o *ConnectionAction) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *ConnectionAction) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *ConnectionAction) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *ConnectionAction) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


