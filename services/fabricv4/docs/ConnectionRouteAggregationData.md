# ConnectionRouteAggregationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Route Aggregation URI | [optional] 
**Type** | Pointer to [**ConnectionRouteAggregationDataType**](ConnectionRouteAggregationDataType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Route Aggregation identifier | [optional] 
**AttachmentStatus** | Pointer to [**ConnectionRouteAggregationDataAttachmentStatus**](ConnectionRouteAggregationDataAttachmentStatus.md) |  | [optional] 

## Methods

### NewConnectionRouteAggregationData

`func NewConnectionRouteAggregationData() *ConnectionRouteAggregationData`

NewConnectionRouteAggregationData instantiates a new ConnectionRouteAggregationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionRouteAggregationDataWithDefaults

`func NewConnectionRouteAggregationDataWithDefaults() *ConnectionRouteAggregationData`

NewConnectionRouteAggregationDataWithDefaults instantiates a new ConnectionRouteAggregationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ConnectionRouteAggregationData) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectionRouteAggregationData) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectionRouteAggregationData) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectionRouteAggregationData) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ConnectionRouteAggregationData) GetType() ConnectionRouteAggregationDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionRouteAggregationData) GetTypeOk() (*ConnectionRouteAggregationDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionRouteAggregationData) SetType(v ConnectionRouteAggregationDataType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionRouteAggregationData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ConnectionRouteAggregationData) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConnectionRouteAggregationData) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConnectionRouteAggregationData) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConnectionRouteAggregationData) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAttachmentStatus

`func (o *ConnectionRouteAggregationData) GetAttachmentStatus() ConnectionRouteAggregationDataAttachmentStatus`

GetAttachmentStatus returns the AttachmentStatus field if non-nil, zero value otherwise.

### GetAttachmentStatusOk

`func (o *ConnectionRouteAggregationData) GetAttachmentStatusOk() (*ConnectionRouteAggregationDataAttachmentStatus, bool)`

GetAttachmentStatusOk returns a tuple with the AttachmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentStatus

`func (o *ConnectionRouteAggregationData) SetAttachmentStatus(v ConnectionRouteAggregationDataAttachmentStatus)`

SetAttachmentStatus sets AttachmentStatus field to given value.

### HasAttachmentStatus

`func (o *ConnectionRouteAggregationData) HasAttachmentStatus() bool`

HasAttachmentStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


