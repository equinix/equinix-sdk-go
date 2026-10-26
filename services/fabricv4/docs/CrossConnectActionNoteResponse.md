# CrossConnectActionNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Note URI. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Note Identifier. | [optional] [readonly] 
**ParentNote** | Pointer to [**CrossConnectParentNote**](CrossConnectParentNote.md) | The note this note is responding to. Returned when the note was added in response to an Equinix note.  | [optional] [readonly] 
**Comments** | Pointer to **string** | Note comments. | [optional] [readonly] 

## Methods

### NewCrossConnectActionNoteResponse

`func NewCrossConnectActionNoteResponse() *CrossConnectActionNoteResponse`

NewCrossConnectActionNoteResponse instantiates a new CrossConnectActionNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionNoteResponseWithDefaults

`func NewCrossConnectActionNoteResponseWithDefaults() *CrossConnectActionNoteResponse`

NewCrossConnectActionNoteResponseWithDefaults instantiates a new CrossConnectActionNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectActionNoteResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectActionNoteResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectActionNoteResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectActionNoteResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectActionNoteResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectActionNoteResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectActionNoteResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectActionNoteResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetParentNote

`func (o *CrossConnectActionNoteResponse) GetParentNote() CrossConnectParentNote`

GetParentNote returns the ParentNote field if non-nil, zero value otherwise.

### GetParentNoteOk

`func (o *CrossConnectActionNoteResponse) GetParentNoteOk() (*CrossConnectParentNote, bool)`

GetParentNoteOk returns a tuple with the ParentNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNote

`func (o *CrossConnectActionNoteResponse) SetParentNote(v CrossConnectParentNote)`

SetParentNote sets ParentNote field to given value.

### HasParentNote

`func (o *CrossConnectActionNoteResponse) HasParentNote() bool`

HasParentNote returns a boolean if a field has been set.

### GetComments

`func (o *CrossConnectActionNoteResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CrossConnectActionNoteResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CrossConnectActionNoteResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CrossConnectActionNoteResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


