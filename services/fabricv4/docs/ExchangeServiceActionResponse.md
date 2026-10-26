# ExchangeServiceActionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ExchangeServiceActionRequestType**](ExchangeServiceActionRequestType.md) |  | 
**Href** | Pointer to **string** | Internet Exchange Service Action URI | [optional] 
**Uuid** | Pointer to **string** | Action identifier | [optional] 
**Description** | Pointer to **string** | Action description | [optional] 
**State** | Pointer to [**ExchangeServiceActionResponseAllOfState**](ExchangeServiceActionResponseAllOfState.md) |  | [optional] 
**Changelog** | Pointer to [**PlatformChangelog**](PlatformChangelog.md) |  | [optional] 

## Methods

### NewExchangeServiceActionResponse

`func NewExchangeServiceActionResponse(type_ ExchangeServiceActionRequestType, ) *ExchangeServiceActionResponse`

NewExchangeServiceActionResponse instantiates a new ExchangeServiceActionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeServiceActionResponseWithDefaults

`func NewExchangeServiceActionResponseWithDefaults() *ExchangeServiceActionResponse`

NewExchangeServiceActionResponseWithDefaults instantiates a new ExchangeServiceActionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExchangeServiceActionResponse) GetType() ExchangeServiceActionRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeServiceActionResponse) GetTypeOk() (*ExchangeServiceActionRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeServiceActionResponse) SetType(v ExchangeServiceActionRequestType)`

SetType sets Type field to given value.


### GetHref

`func (o *ExchangeServiceActionResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ExchangeServiceActionResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ExchangeServiceActionResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ExchangeServiceActionResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *ExchangeServiceActionResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ExchangeServiceActionResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ExchangeServiceActionResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ExchangeServiceActionResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetDescription

`func (o *ExchangeServiceActionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExchangeServiceActionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExchangeServiceActionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExchangeServiceActionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *ExchangeServiceActionResponse) GetState() ExchangeServiceActionResponseAllOfState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ExchangeServiceActionResponse) GetStateOk() (*ExchangeServiceActionResponseAllOfState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ExchangeServiceActionResponse) SetState(v ExchangeServiceActionResponseAllOfState)`

SetState sets State field to given value.

### HasState

`func (o *ExchangeServiceActionResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChangelog

`func (o *ExchangeServiceActionResponse) GetChangelog() PlatformChangelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *ExchangeServiceActionResponse) GetChangelogOk() (*PlatformChangelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *ExchangeServiceActionResponse) SetChangelog(v PlatformChangelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *ExchangeServiceActionResponse) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


