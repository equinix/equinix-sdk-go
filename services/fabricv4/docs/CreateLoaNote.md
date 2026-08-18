# CreateLoaNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | **string** | The note content to add to this LOA. Notes are visible to both the issuer and requestor organizations. Use notes to communicate updates, clarifications, or additional context about the LOA.  | 

## Methods

### NewCreateLoaNote

`func NewCreateLoaNote(comments string, ) *CreateLoaNote`

NewCreateLoaNote instantiates a new CreateLoaNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoaNoteWithDefaults

`func NewCreateLoaNoteWithDefaults() *CreateLoaNote`

NewCreateLoaNoteWithDefaults instantiates a new CreateLoaNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *CreateLoaNote) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CreateLoaNote) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CreateLoaNote) SetComments(v string)`

SetComments sets Comments field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


