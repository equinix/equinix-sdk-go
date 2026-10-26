# CrossConnectNoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Note URI. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Note Identifier. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectNoteResponseType**](CrossConnectNoteResponseType.md) |  | [optional] 
**Source** | Pointer to [**CrossConnectNoteResponseSource**](CrossConnectNoteResponseSource.md) |  | [optional] 
**Comments** | Pointer to **string** | Note comments. | [optional] [readonly] 
**ParentNote** | Pointer to [**CrossConnectParentNote**](CrossConnectParentNote.md) | The note this note is responding to. Returned when the note was added in response to an Equinix note.  | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] 

## Methods

### NewCrossConnectNoteResponse

`func NewCrossConnectNoteResponse() *CrossConnectNoteResponse`

NewCrossConnectNoteResponse instantiates a new CrossConnectNoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectNoteResponseWithDefaults

`func NewCrossConnectNoteResponseWithDefaults() *CrossConnectNoteResponse`

NewCrossConnectNoteResponseWithDefaults instantiates a new CrossConnectNoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectNoteResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectNoteResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectNoteResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectNoteResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectNoteResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectNoteResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectNoteResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectNoteResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectNoteResponse) GetType() CrossConnectNoteResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectNoteResponse) GetTypeOk() (*CrossConnectNoteResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectNoteResponse) SetType(v CrossConnectNoteResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectNoteResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *CrossConnectNoteResponse) GetSource() CrossConnectNoteResponseSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CrossConnectNoteResponse) GetSourceOk() (*CrossConnectNoteResponseSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CrossConnectNoteResponse) SetSource(v CrossConnectNoteResponseSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CrossConnectNoteResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetComments

`func (o *CrossConnectNoteResponse) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CrossConnectNoteResponse) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CrossConnectNoteResponse) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CrossConnectNoteResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetParentNote

`func (o *CrossConnectNoteResponse) GetParentNote() CrossConnectParentNote`

GetParentNote returns the ParentNote field if non-nil, zero value otherwise.

### GetParentNoteOk

`func (o *CrossConnectNoteResponse) GetParentNoteOk() (*CrossConnectParentNote, bool)`

GetParentNoteOk returns a tuple with the ParentNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNote

`func (o *CrossConnectNoteResponse) SetParentNote(v CrossConnectParentNote)`

SetParentNote sets ParentNote field to given value.

### HasParentNote

`func (o *CrossConnectNoteResponse) HasParentNote() bool`

HasParentNote returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectNoteResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectNoteResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectNoteResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectNoteResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


