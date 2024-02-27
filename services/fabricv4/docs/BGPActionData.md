# BGPActionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Routing Protocol URI | [optional] 
**Uuid** | Pointer to **string** | Routing protocol identifier | [optional] 
**Type** | Pointer to [**BGPActions**](BGPActions.md) |  | [optional] 
**Description** | Pointer to **string** | BGP action description | [optional] 
**State** | Pointer to [**BGPActionStates**](BGPActionStates.md) |  | [optional] 
**Changelog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewBGPActionData

`func NewBGPActionData() *BGPActionData`

NewBGPActionData instantiates a new BGPActionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBGPActionDataWithDefaults

`func NewBGPActionDataWithDefaults() *BGPActionData`

NewBGPActionDataWithDefaults instantiates a new BGPActionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *BGPActionData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *BGPActionData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *BGPActionData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *BGPActionData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *BGPActionData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BGPActionData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BGPActionData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BGPActionData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *BGPActionData) GetType() BGPActions`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BGPActionData) GetTypeOk() (*BGPActions, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BGPActionData) SetType(v BGPActions)`

SetType sets Type field to given value.

### HasType

`func (o *BGPActionData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *BGPActionData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BGPActionData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BGPActionData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BGPActionData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetState

`func (o *BGPActionData) GetState() BGPActionStates`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BGPActionData) GetStateOk() (*BGPActionStates, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BGPActionData) SetState(v BGPActionStates)`

SetState sets State field to given value.

### HasState

`func (o *BGPActionData) HasState() bool`

HasState returns a boolean if a field has been set.

### GetChangelog

`func (o *BGPActionData) GetChangelog() Changelog`

GetChangelog returns the Changelog field if non-nil, zero value otherwise.

### GetChangelogOk

`func (o *BGPActionData) GetChangelogOk() (*Changelog, bool)`

GetChangelogOk returns a tuple with the Changelog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangelog

`func (o *BGPActionData) SetChangelog(v Changelog)`

SetChangelog sets Changelog field to given value.

### HasChangelog

`func (o *BGPActionData) HasChangelog() bool`

HasChangelog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


