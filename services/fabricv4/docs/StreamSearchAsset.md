# StreamSearchAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream Asset URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**StreamUuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to **string** | Asset type | [optional] 
**MetricsEnabled** | Pointer to **bool** | enable metric | [optional] 
**AttachmentStatus** | Pointer to [**StreamAssetAttachmentStatus**](StreamAssetAttachmentStatus.md) |  | [optional] 
**ProjectId** | Pointer to **string** | project ic | [optional] 

## Methods

### NewStreamSearchAsset

`func NewStreamSearchAsset() *StreamSearchAsset`

NewStreamSearchAsset instantiates a new StreamSearchAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamSearchAssetWithDefaults

`func NewStreamSearchAssetWithDefaults() *StreamSearchAsset`

NewStreamSearchAssetWithDefaults instantiates a new StreamSearchAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamSearchAsset) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamSearchAsset) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamSearchAsset) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamSearchAsset) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *StreamSearchAsset) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamSearchAsset) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamSearchAsset) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamSearchAsset) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStreamUuid

`func (o *StreamSearchAsset) GetStreamUuid() string`

GetStreamUuid returns the StreamUuid field if non-nil, zero value otherwise.

### GetStreamUuidOk

`func (o *StreamSearchAsset) GetStreamUuidOk() (*string, bool)`

GetStreamUuidOk returns a tuple with the StreamUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamUuid

`func (o *StreamSearchAsset) SetStreamUuid(v string)`

SetStreamUuid sets StreamUuid field to given value.

### HasStreamUuid

`func (o *StreamSearchAsset) HasStreamUuid() bool`

HasStreamUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamSearchAsset) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamSearchAsset) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamSearchAsset) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StreamSearchAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *StreamSearchAsset) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *StreamSearchAsset) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *StreamSearchAsset) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *StreamSearchAsset) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *StreamSearchAsset) GetAttachmentStatus() StreamAssetAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *StreamSearchAsset) GetAttachmentStatusOk() (*StreamAssetAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *StreamSearchAsset) SetAttachmentStatus(v StreamAssetAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *StreamSearchAsset) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetProjectId

`func (o *StreamSearchAsset) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *StreamSearchAsset) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *StreamSearchAsset) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *StreamSearchAsset) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


