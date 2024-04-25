# ConnectionLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ConnectionLinkType**](ConnectionLinkType.md) |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewConnectionLink

`func NewConnectionLink() *ConnectionLink`

NewConnectionLink instantiates a new ConnectionLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionLinkWithDefaults

`func NewConnectionLinkWithDefaults() *ConnectionLink`

NewConnectionLinkWithDefaults instantiates a new ConnectionLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ConnectionLink) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ConnectionLink) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ConnectionLink) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ConnectionLink) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *ConnectionLink) GetType() ConnectionLinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionLink) GetTypeOk() (*ConnectionLinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionLink) SetType(v ConnectionLinkType)`

SetType sets Type field to given value.

### HasType

`func (o *ConnectionLink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *ConnectionLink) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ConnectionLink) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ConnectionLink) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ConnectionLink) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


