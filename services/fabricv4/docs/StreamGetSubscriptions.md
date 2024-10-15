# StreamGetSubscriptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream Get Stream Subscriptions URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to [**StreamGetSubscriptionsType**](StreamGetSubscriptionsType.md) |  | [optional] 

## Methods

### NewStreamGetSubscriptions

`func NewStreamGetSubscriptions() *StreamGetSubscriptions`

NewStreamGetSubscriptions instantiates a new StreamGetSubscriptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGetSubscriptionsWithDefaults

`func NewStreamGetSubscriptionsWithDefaults() *StreamGetSubscriptions`

NewStreamGetSubscriptionsWithDefaults instantiates a new StreamGetSubscriptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamGetSubscriptions) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamGetSubscriptions) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamGetSubscriptions) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamGetSubscriptions) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *StreamGetSubscriptions) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamGetSubscriptions) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamGetSubscriptions) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamGetSubscriptions) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamGetSubscriptions) GetType() StreamGetSubscriptionsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamGetSubscriptions) GetTypeOk() (*StreamGetSubscriptionsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamGetSubscriptions) SetType(v StreamGetSubscriptionsType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamGetSubscriptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


