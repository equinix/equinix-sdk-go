# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Interpolated** | Pointer to **string** |  | [optional] 
**Relationships** | Pointer to [**[]Href**](Href.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ModifiedBy** | Pointer to **map[string]interface{}** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Event) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Event) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Event) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Event) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Event) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHref

`func (o *Event) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Event) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Event) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Event) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Event) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterpolated

`func (o *Event) GetInterpolated() string`

GetInterpolated returns the Interpolated field if non-nil, zero value otherwise.

### GetInterpolatedOk

`func (o *Event) GetInterpolatedOk() (*string, bool)`

GetInterpolatedOk returns a tuple with the Interpolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterpolated

`func (o *Event) SetInterpolated(v string)`

SetInterpolated sets Interpolated field to given value.

### HasInterpolated

`func (o *Event) HasInterpolated() bool`

HasInterpolated returns a boolean if a field has been set.

### GetRelationships

`func (o *Event) GetRelationships() []Href`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Event) GetRelationshipsOk() (*[]Href, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Event) SetRelationships(v []Href)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *Event) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetState

`func (o *Event) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Event) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Event) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Event) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Event) HasType() bool`

HasType returns a boolean if a field has been set.

### GetModifiedBy

`func (o *Event) GetModifiedBy() map[string]interface{}`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Event) GetModifiedByOk() (*map[string]interface{}, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Event) SetModifiedBy(v map[string]interface{})`

SetModifiedBy sets ModifiedBy field to given value.

### HasModifiedBy

`func (o *Event) HasModifiedBy() bool`

HasModifiedBy returns a boolean if a field has been set.

### GetIp

`func (o *Event) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Event) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Event) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Event) HasIp() bool`

HasIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


