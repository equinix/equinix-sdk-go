# StreamAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Stream Asset URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned access point identifier | [optional] 
**Type** | Pointer to [**StreamAssetType**](StreamAssetType.md) |  | [optional] 
**MetricsEnabled** | Pointer to **bool** | enable metric | [optional] 
**AttachmentStatus** | Pointer to [**StreamAssetAttachmentStatus**](StreamAssetAttachmentStatus.md) |  | [optional] 

## Methods

### NewStreamAsset

`func NewStreamAsset() *StreamAsset`

NewStreamAsset instantiates a new StreamAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAssetWithDefaults

`func NewStreamAssetWithDefaults() *StreamAsset`

NewStreamAssetWithDefaults instantiates a new StreamAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamAsset) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamAsset) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamAsset) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamAsset) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *StreamAsset) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamAsset) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamAsset) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamAsset) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamAsset) GetType() StreamAssetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamAsset) GetTypeOk() (*StreamAssetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamAsset) SetType(v StreamAssetType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamAsset) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetricsEnabled

`func (o *StreamAsset) GetMetricsEnabled() bool`

GetMetricsEnabled returns the MetricsEnabled field if non-nil, zero value otherwise.

### GetMetricsEnabledOk

`func (o *StreamAsset) GetMetricsEnabledOk() (*bool, bool)`

GetMetricsEnabledOk returns a tuple with the MetricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsEnabled

`func (o *StreamAsset) SetMetricsEnabled(v bool)`

SetMetricsEnabled sets MetricsEnabled field to given value.

### HasMetricsEnabled

`func (o *StreamAsset) HasMetricsEnabled() bool`

HasMetricsEnabled returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *StreamAsset) GetAttachmentStatus() StreamAssetAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *StreamAsset) GetAttachmentStatusOk() (*StreamAssetAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *StreamAsset) SetAttachmentStatus(v StreamAssetAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *StreamAsset) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


