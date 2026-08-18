# EnvironmentActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Environment action URI | [optional] [readonly] 
**Type** | Pointer to [**EnvironmentActionTypeEnum**](EnvironmentActionTypeEnum.md) |  | [optional] 
**Uuid** | Pointer to **string** | Equinix-assigned action identifier | [optional] 
**State** | Pointer to [**EnvironmentActionStateEnum**](EnvironmentActionStateEnum.md) |  | [optional] 
**KeyDetails** | Pointer to [**ActivationKeyDetails**](ActivationKeyDetails.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewEnvironmentActionResponse

`func NewEnvironmentActionResponse() *EnvironmentActionResponse`

NewEnvironmentActionResponse instantiates a new EnvironmentActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentActionResponseWithDefaults

`func NewEnvironmentActionResponseWithDefaults() *EnvironmentActionResponse`

NewEnvironmentActionResponseWithDefaults instantiates a new EnvironmentActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *EnvironmentActionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *EnvironmentActionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *EnvironmentActionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *EnvironmentActionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentActionResponse) GetType() EnvironmentActionTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentActionResponse) GetTypeOk() (*EnvironmentActionTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentActionResponse) SetType(v EnvironmentActionTypeEnum)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentActionResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *EnvironmentActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EnvironmentActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EnvironmentActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EnvironmentActionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentActionResponse) GetState() EnvironmentActionStateEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentActionResponse) GetStateOk() (*EnvironmentActionStateEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentActionResponse) SetState(v EnvironmentActionStateEnum)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentActionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetKeyDetails

`func (o *EnvironmentActionResponse) GetKeyDetails() ActivationKeyDetails`

GetKeyDetails returns the KeyDetails field if non-nil, zero value otherwise.

### GetKeyDetailsOk

`func (o *EnvironmentActionResponse) GetKeyDetailsOk() (*ActivationKeyDetails, bool)`

GetKeyDetailsOk returns a tuple with the KeyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDetails

`func (o *EnvironmentActionResponse) SetKeyDetails(v ActivationKeyDetails)`

SetKeyDetails sets KeyDetails field to given value.

### HasKeyDetails

`func (o *EnvironmentActionResponse) HasKeyDetails() bool`

HasKeyDetails returns a boolean if a field has been set.

### GetChangeLog

`func (o *EnvironmentActionResponse) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *EnvironmentActionResponse) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *EnvironmentActionResponse) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *EnvironmentActionResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


