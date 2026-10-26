# StreamAssetInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Network Edge Device interface URI | [optional] [readonly] 
**Id** | Pointer to **int32** | Network Edge Device interface identifier | [optional] 
**Uuid** | Pointer to **string** | Network Edge Device interface UUID | [optional] 
**Type** | Pointer to [**StreamAssetInterfaceType**](StreamAssetInterfaceType.md) |  | [optional] 
**State** | Pointer to **string** | Network Edge Device interface usage status | [optional] 
**AttachmentStatus** | Pointer to [**StreamAssetInterfaceAttachmentStatus**](StreamAssetInterfaceAttachmentStatus.md) |  | [optional] 

## Methods

### NewStreamAssetInterface

`func NewStreamAssetInterface() *StreamAssetInterface`

NewStreamAssetInterface instantiates a new StreamAssetInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamAssetInterfaceWithDefaults

`func NewStreamAssetInterfaceWithDefaults() *StreamAssetInterface`

NewStreamAssetInterfaceWithDefaults instantiates a new StreamAssetInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *StreamAssetInterface) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StreamAssetInterface) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StreamAssetInterface) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *StreamAssetInterface) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *StreamAssetInterface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamAssetInterface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamAssetInterface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StreamAssetInterface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *StreamAssetInterface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *StreamAssetInterface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *StreamAssetInterface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *StreamAssetInterface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *StreamAssetInterface) GetType() StreamAssetInterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StreamAssetInterface) GetTypeOk() (*StreamAssetInterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StreamAssetInterface) SetType(v StreamAssetInterfaceType)`

SetType sets Type field to given value.

### HasType

`func (o *StreamAssetInterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetState

`func (o *StreamAssetInterface) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *StreamAssetInterface) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *StreamAssetInterface) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *StreamAssetInterface) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *StreamAssetInterface) GetAttachmentStatus() StreamAssetInterfaceAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *StreamAssetInterface) GetAttachmentStatusOk() (*StreamAssetInterfaceAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *StreamAssetInterface) SetAttachmentStatus(v StreamAssetInterfaceAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *StreamAssetInterface) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


