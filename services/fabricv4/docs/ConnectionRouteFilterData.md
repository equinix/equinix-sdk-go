# ConnectionRouteFilterData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Filter URI | [optional] 
**Type** | Pointer to [**ConnectionRouteFilterDataType**](ConnectionRouteFilterDataType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Filter identifier | [optional] 
**AttachmentStatus** | Pointer to [**ConnectionRouteFilterDataAttachmentStatus**](ConnectionRouteFilterDataAttachmentStatus.md) |  | [optional] 
**Direction** | Pointer to [**ConnectionRouteFilterDataDirection**](ConnectionRouteFilterDataDirection.md) |  | [optional] 

## Methods

### NewConnectionRouteFilterData

`func NewConnectionRouteFilterData() *ConnectionRouteFilterData`

NewConnectionRouteFilterData instantiates a new ConnectionRouteFilterData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRouteFilterDataWithDefaults

`func NewConnectionRouteFilterDataWithDefaults() *ConnectionRouteFilterData`

NewConnectionRouteFilterDataWithDefaults instantiates a new ConnectionRouteFilterData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ConnectionRouteFilterData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectionRouteFilterData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectionRouteFilterData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectionRouteFilterData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ConnectionRouteFilterData) GetType() ConnectionRouteFilterDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionRouteFilterData) GetTypeOk() (*ConnectionRouteFilterDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionRouteFilterData) SetType(v ConnectionRouteFilterDataType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionRouteFilterData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ConnectionRouteFilterData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConnectionRouteFilterData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConnectionRouteFilterData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConnectionRouteFilterData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *ConnectionRouteFilterData) GetAttachmentStatus() ConnectionRouteFilterDataAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *ConnectionRouteFilterData) GetAttachmentStatusOk() (*ConnectionRouteFilterDataAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *ConnectionRouteFilterData) SetAttachmentStatus(v ConnectionRouteFilterDataAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *ConnectionRouteFilterData) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.

### GetDirection

`func (o *ConnectionRouteFilterData) GetDirection() ConnectionRouteFilterDataDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ConnectionRouteFilterData) GetDirectionOk() (*ConnectionRouteFilterDataDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ConnectionRouteFilterData) SetDirection(v ConnectionRouteFilterDataDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ConnectionRouteFilterData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


