# CrossConnectAttachmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | URI to the attachment resource. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Attachment resource Identifier. | [optional] [readonly] 
**Type** | Pointer to [**CrossConnectAttachmentResponseType**](CrossConnectAttachmentResponseType.md) |  | [optional] 
**Source** | Pointer to [**CrossConnectAttachmentResponseSource**](CrossConnectAttachmentResponseSource.md) |  | [optional] 
**AttachmentId** | Pointer to **string** | Attachment Identifier. | [optional] [readonly] 
**DownloadHref** | Pointer to **string** | URI to download the attachment file content from the Equinix attachment service.  | [optional] [readonly] 
**ChangeLog** | Pointer to [**CrossConnectChangelog**](CrossConnectChangelog.md) |  | [optional] 

## Methods

### NewCrossConnectAttachmentResponse

`func NewCrossConnectAttachmentResponse() *CrossConnectAttachmentResponse`

NewCrossConnectAttachmentResponse instantiates a new CrossConnectAttachmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrossConnectAttachmentResponseWithDefaults

`func NewCrossConnectAttachmentResponseWithDefaults() *CrossConnectAttachmentResponse`

NewCrossConnectAttachmentResponseWithDefaults instantiates a new CrossConnectAttachmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *CrossConnectAttachmentResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CrossConnectAttachmentResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CrossConnectAttachmentResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CrossConnectAttachmentResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *CrossConnectAttachmentResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CrossConnectAttachmentResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CrossConnectAttachmentResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CrossConnectAttachmentResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *CrossConnectAttachmentResponse) GetType() CrossConnectAttachmentResponseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CrossConnectAttachmentResponse) GetTypeOk() (*CrossConnectAttachmentResponseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CrossConnectAttachmentResponse) SetType(v CrossConnectAttachmentResponseType)`

SetType sets Type field to given value.

### HasType

`func (o *CrossConnectAttachmentResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *CrossConnectAttachmentResponse) GetSource() CrossConnectAttachmentResponseSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CrossConnectAttachmentResponse) GetSourceOk() (*CrossConnectAttachmentResponseSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CrossConnectAttachmentResponse) SetSource(v CrossConnectAttachmentResponseSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *CrossConnectAttachmentResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetAttachmentId

`func (o *CrossConnectAttachmentResponse) GetAttachmentId() string`

GetAttachmentId returns the AttachmentId field if non-nil, zero value otherwise.

### GetAttachmentIdOk

`func (o *CrossConnectAttachmentResponse) GetAttachmentIdOk() (*string, bool)`

GetAttachmentIdOk returns a tuple with the AttachmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentId

`func (o *CrossConnectAttachmentResponse) SetAttachmentId(v string)`

SetAttachmentId sets AttachmentId field to given value.

### HasAttachmentId

`func (o *CrossConnectAttachmentResponse) HasAttachmentId() bool`

HasAttachmentId returns a boolean if a field has been set.

### GetDownloadHref

`func (o *CrossConnectAttachmentResponse) GetDownloadHref() string`

GetDownloadHref returns the DownloadHref field if non-nil, zero value otherwise.

### GetDownloadHrefOk

`func (o *CrossConnectAttachmentResponse) GetDownloadHrefOk() (*string, bool)`

GetDownloadHrefOk returns a tuple with the DownloadHref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadHref

`func (o *CrossConnectAttachmentResponse) SetDownloadHref(v string)`

SetDownloadHref sets DownloadHref field to given value.

### HasDownloadHref

`func (o *CrossConnectAttachmentResponse) HasDownloadHref() bool`

HasDownloadHref returns a boolean if a field has been set.

### GetChangeLog

`func (o *CrossConnectAttachmentResponse) GetChangeLog() CrossConnectChangelog`

GetChangeLog returns the ChangeLog field if non-nil, zero value otherwise.

### GetChangeLogOk

`func (o *CrossConnectAttachmentResponse) GetChangeLogOk() (*CrossConnectChangelog, bool)`

GetChangeLogOk returns a tuple with the ChangeLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeLog

`func (o *CrossConnectAttachmentResponse) SetChangeLog(v CrossConnectChangelog)`

SetChangeLog sets ChangeLog field to given value.

### HasChangeLog

`func (o *CrossConnectAttachmentResponse) HasChangeLog() bool`

HasChangeLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


