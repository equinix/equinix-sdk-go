# AttachPrivateServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URL to the attached private service | [optional] 
**Type** | **string** | Type of the private service or attachment | 
**Uuid** | **string** | Unique identifier for the private service | 
**AttachmentStatus** | Pointer to **string** | Status of the attachment operation | [optional] 

## Methods

### NewAttachPrivateServiceResponse

`func NewAttachPrivateServiceResponse(type_ string, uuid string, ) *AttachPrivateServiceResponse`

NewAttachPrivateServiceResponse instantiates a new AttachPrivateServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachPrivateServiceResponseWithDefaults

`func NewAttachPrivateServiceResponseWithDefaults() *AttachPrivateServiceResponse`

NewAttachPrivateServiceResponseWithDefaults instantiates a new AttachPrivateServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AttachPrivateServiceResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AttachPrivateServiceResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AttachPrivateServiceResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AttachPrivateServiceResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AttachPrivateServiceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AttachPrivateServiceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AttachPrivateServiceResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUuid

`func (o *AttachPrivateServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AttachPrivateServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AttachPrivateServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetAttachmentStatus

`func (o *AttachPrivateServiceResponse) GetAttachmentStatus() string`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AttachPrivateServiceResponse) GetAttachmentStatusOk() (*string, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AttachPrivateServiceResponse) SetAttachmentStatus(v string)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AttachPrivateServiceResponse) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


