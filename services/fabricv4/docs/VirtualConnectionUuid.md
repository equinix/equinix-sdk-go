# VirtualConnectionUuid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Connection URI | [optional] [readonly] 
**Type** | Pointer to **string** | Connection Type | [optional] 
**Uuid** | **string** | Connection UUID. | 

## Methods

### NewVirtualConnectionUuid

`func NewVirtualConnectionUuid(uuid string, ) *VirtualConnectionUuid`

NewVirtualConnectionUuid instantiates a new VirtualConnectionUuid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualConnectionUuidWithDefaults

`func NewVirtualConnectionUuidWithDefaults() *VirtualConnectionUuid`

NewVirtualConnectionUuidWithDefaults instantiates a new VirtualConnectionUuid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *VirtualConnectionUuid) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VirtualConnectionUuid) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VirtualConnectionUuid) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *VirtualConnectionUuid) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *VirtualConnectionUuid) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualConnectionUuid) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualConnectionUuid) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualConnectionUuid) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *VirtualConnectionUuid) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *VirtualConnectionUuid) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *VirtualConnectionUuid) SetUuid(v string)`

SetUuid sets Uuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


