# LoaNoteDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**LoaNoteDetailsType**](LoaNoteDetailsType.md) |  | [optional] 
**Href** | Pointer to **string** | URI to note resource. | [optional] 
**Uuid** | Pointer to **string** | Unique identifier of this note. | [optional] 
**Comments** | Pointer to **string** | Content of the note as submitted by the user. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time when the note was created. | [optional] 

## Methods

### NewLoaNoteDetails

`func NewLoaNoteDetails() *LoaNoteDetails`

NewLoaNoteDetails instantiates a new LoaNoteDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoaNoteDetailsWithDefaults

`func NewLoaNoteDetailsWithDefaults() *LoaNoteDetails`

NewLoaNoteDetailsWithDefaults instantiates a new LoaNoteDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LoaNoteDetails) GetType() LoaNoteDetailsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoaNoteDetails) GetTypeOk() (*LoaNoteDetailsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoaNoteDetails) SetType(v LoaNoteDetailsType)`

SetType sets Type field to given value.

### HasType

`func (o *LoaNoteDetails) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *LoaNoteDetails) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LoaNoteDetails) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LoaNoteDetails) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LoaNoteDetails) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *LoaNoteDetails) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LoaNoteDetails) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LoaNoteDetails) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *LoaNoteDetails) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetComments

`func (o *LoaNoteDetails) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *LoaNoteDetails) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *LoaNoteDetails) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *LoaNoteDetails) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *LoaNoteDetails) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *LoaNoteDetails) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *LoaNoteDetails) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *LoaNoteDetails) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


