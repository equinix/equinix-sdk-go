# Note

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the note | [optional] 
**ReferenceId** | Pointer to **string** | Unique identifier for referencing notes. | [optional] 
**CreatedDateTime** | Pointer to **time.Time** | Date and time of the note creation in UTC timezone | [optional] 
**Text** | Pointer to **string** | The content of the note | [optional] 
**Author** | Pointer to **string** | Author of the notes | [optional] 
**Type** | Pointer to [**NoteType**](NoteType.md) |  | [optional] 
**Attachments** | Pointer to [**[]Attachment**](Attachment.md) | File(s) attached to the orders. To learn about including attachments in your request, refer to Attachments API. | [optional] 

## Methods

### NewNote

`func NewNote() *Note`

NewNote instantiates a new Note object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNoteWithDefaults

`func NewNoteWithDefaults() *Note`

NewNoteWithDefaults instantiates a new Note object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Note) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Note) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Note) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Note) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReferenceId

`func (o *Note) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Note) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Note) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Note) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetCreatedDateTime

`func (o *Note) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *Note) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *Note) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *Note) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### GetText

`func (o *Note) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Note) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Note) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Note) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAuthor

`func (o *Note) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Note) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Note) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Note) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetType

`func (o *Note) GetType() NoteType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Note) GetTypeOk() (*NoteType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Note) SetType(v NoteType)`

SetType sets Type field to given value.

### HasType

`func (o *Note) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttachments

`func (o *Note) GetAttachments() []Attachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Note) GetAttachmentsOk() (*[]Attachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Note) SetAttachments(v []Attachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Note) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


