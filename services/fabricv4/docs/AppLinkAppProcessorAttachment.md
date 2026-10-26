# AppLinkAppProcessorAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Resource URI | [optional] [readonly] 
**Type** | Pointer to [**AppProcessorType**](AppProcessorType.md) |  | [optional] [default to APPPROCESSORTYPE_APP_PROCESSOR]
**Uuid** | **string** | App Processor UUID | 
**Config** | Pointer to **string** | Base64-encoded configuration for the processor attachment | [optional] 
**Endpoint** | Pointer to **string** | Endpoint DNS name assigned to the attached App Processor | [optional] 
**AttachmentStatus** | Pointer to [**AppLinkAttachState**](AppLinkAttachState.md) |  | [optional] 
**ChangeLog** | Pointer to [**Changelog**](Changelog.md) |  | [optional] 

## Methods

### NewAppLinkAppProcessorAttachment

`func NewAppLinkAppProcessorAttachment(uuid string, ) *AppLinkAppProcessorAttachment`

NewAppLinkAppProcessorAttachment instantiates a new AppLinkAppProcessorAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppLinkAppProcessorAttachmentWithDefaults

`func NewAppLinkAppProcessorAttachmentWithDefaults() *AppLinkAppProcessorAttachment`

NewAppLinkAppProcessorAttachmentWithDefaults instantiates a new AppLinkAppProcessorAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *AppLinkAppProcessorAttachment) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AppLinkAppProcessorAttachment) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AppLinkAppProcessorAttachment) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *AppLinkAppProcessorAttachment) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *AppLinkAppProcessorAttachment) GetType() AppProcessorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppLinkAppProcessorAttachment) GetTypeOk() (*AppProcessorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppLinkAppProcessorAttachment) SetType(v AppProcessorType)`

SetType sets Type field to given value.

### HasType

`func (o *AppLinkAppProcessorAttachment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *AppLinkAppProcessorAttachment) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppLinkAppProcessorAttachment) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppLinkAppProcessorAttachment) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetConfig

`func (o *AppLinkAppProcessorAttachment) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AppLinkAppProcessorAttachment) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AppLinkAppProcessorAttachment) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AppLinkAppProcessorAttachment) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEndpoint

`func (o *AppLinkAppProcessorAttachment) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AppLinkAppProcessorAttachment) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AppLinkAppProcessorAttachment) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AppLinkAppProcessorAttachment) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *AppLinkAppProcessorAttachment) GetAttachmentStatus() AppLinkAttachState`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *AppLinkAppProcessorAttachment) GetAttachmentStatusOk() (*AppLinkAttachState, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *AppLinkAppProcessorAttachment) SetAttachmentStatus(v AppLinkAttachState)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *AppLinkAppProcessorAttachment) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetChangeLog

`func (o *AppLinkAppProcessorAttachment) GetChangeLog() Changelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *AppLinkAppProcessorAttachment) GetChangeLogOk() (*Changelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *AppLinkAppProcessorAttachment) SetChangeLog(v Changelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *AppLinkAppProcessorAttachment) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


