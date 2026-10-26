# CrossConnectActionNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentNote** | Pointer to [**CrossConnectParentNote**](CrossConnectParentNote.md) |  | [optional] 
**Comments** | **string** | Note comments. | 

## Methods

### NewCrossConnectActionNote

`func NewCrossConnectActionNote(comments string, ) *CrossConnectActionNote`

NewCrossConnectActionNote instantiates a new CrossConnectActionNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectActionNoteWithDefaults

`func NewCrossConnectActionNoteWithDefaults() *CrossConnectActionNote`

NewCrossConnectActionNoteWithDefaults instantiates a new CrossConnectActionNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentNote

`func (o *CrossConnectActionNote) GetParentNote() CrossConnectParentNote`

GetParentNote returns the ParentNote field if non-nil, zero value otherwise.

### GetParentNoteOk

`func (o *CrossConnectActionNote) GetParentNoteOk() (*CrossConnectParentNote, bool)`

GetParentNoteOk returns a tuple with the ParentNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNote

`func (o *CrossConnectActionNote) SetParentNote(v CrossConnectParentNote)`

SetParentNote sets ParentNote field to given value.

### HasParentNote

`func (o *CrossConnectActionNote) HasParentNote() bool`

HasParentNote returns a boolean if a field has been set.

### GetComments

`func (o *CrossConnectActionNote) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CrossConnectActionNote) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CrossConnectActionNote) SetComments(v string)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


