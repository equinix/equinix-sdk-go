# AttachTagResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URL to the attached tag | [optional] 
**Type** | **string** | Type of the tag or attachment | 
**Uuid** | **string** | Unique identifier for the tag | 
**AttachmentStatus** | Pointer to **string** | Status of the attachment operation | [optional] 

## Methods

### NewAttachTagResponse

`func NewAttachTagResponse(type_ string, uuid string, ) *AttachTagResponse`

NewAttachTagResponse instantiates a new AttachTagResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachTagResponseWithDefaults

`func NewAttachTagResponseWithDefaults() *AttachTagResponse`

NewAttachTagResponseWithDefaults instantiates a new AttachTagResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AttachTagResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AttachTagResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AttachTagResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AttachTagResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AttachTagResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachTagResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachTagResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUuid

`func (o *AttachTagResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AttachTagResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AttachTagResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetAttachmentStatus

`func (o *AttachTagResponse) GetAttachmentStatus() string`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AttachTagResponse) GetAttachmentStatusOk() (*string, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AttachTagResponse) SetAttachmentStatus(v string)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AttachTagResponse) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


