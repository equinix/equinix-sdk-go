# Interface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Interface URI | [optional] [readonly] 
**Uuid** | Pointer to **string** | Equinix-assigned Interface identifier | [optional] 
**Id** | Pointer to **int32** | Interface id | [optional] 
**Type** | Pointer to [**InterfaceType**](InterfaceType.md) |  | [optional] 
**ProjectId** | Pointer to **string** | Interface Project ID | [optional] 

## Methods

### NewInterface

`func NewInterface() *Interface`

NewInterface instantiates a new Interface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceWithDefaults

`func NewInterfaceWithDefaults() *Interface`

NewInterfaceWithDefaults instantiates a new Interface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *Interface) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *Interface) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *Interface) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *Interface) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *Interface) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Interface) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Interface) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Interface) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetId

`func (o *Interface) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Interface) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Interface) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Interface) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Interface) GetType() InterfaceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Interface) GetTypeOk() (*InterfaceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Interface) SetType(v InterfaceType)`

SetType sets Type field to given value.

### HasType

`func (o *Interface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProjectId

`func (o *Interface) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Interface) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Interface) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *Interface) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


