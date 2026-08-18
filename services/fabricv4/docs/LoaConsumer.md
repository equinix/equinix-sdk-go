# LoaConsumer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Consumer URI | [optional] 
**Uuid** | Pointer to **string** | Consumer Identifier | [optional] 
**Type** | Pointer to [**LoaProductType**](LoaProductType.md) |  | [optional] 
**ChangeLog** | Pointer to [**LoaChangelog**](LoaChangelog.md) |  | [optional] 

## Methods

### NewLoaConsumer

`func NewLoaConsumer() *LoaConsumer`

NewLoaConsumer instantiates a new LoaConsumer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaConsumerWithDefaults

`func NewLoaConsumerWithDefaults() *LoaConsumer`

NewLoaConsumerWithDefaults instantiates a new LoaConsumer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LoaConsumer) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LoaConsumer) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LoaConsumer) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LoaConsumer) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LoaConsumer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LoaConsumer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LoaConsumer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LoaConsumer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *LoaConsumer) GetType() LoaProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaConsumer) GetTypeOk() (*LoaProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaConsumer) SetType(v LoaProductType)`

SetType sets Type field to given value.

### HasType

`func (o *LoaConsumer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetChangeLog

`func (o *LoaConsumer) GetChangeLog() LoaChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *LoaConsumer) GetChangeLogOk() (*LoaChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *LoaConsumer) SetChangeLog(v LoaChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *LoaConsumer) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


