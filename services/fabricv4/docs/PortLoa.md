# PortLoa

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | uuid | [optional] 
**Href** | Pointer to **string** | Loa uri. | [optional] [readonly] 
**Type** | Pointer to [**PortLoaType**](PortLoaType.md) |  | [optional] 

## Methods

### NewPortLoa

`func NewPortLoa() *PortLoa`

NewPortLoa instantiates a new PortLoa object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortLoaWithDefaults

`func NewPortLoaWithDefaults() *PortLoa`

NewPortLoaWithDefaults instantiates a new PortLoa object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PortLoa) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PortLoa) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PortLoa) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PortLoa) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetHref

`func (o *PortLoa) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PortLoa) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PortLoa) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PortLoa) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetType

`func (o *PortLoa) GetType() PortLoaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortLoa) GetTypeOk() (*PortLoaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortLoa) SetType(v PortLoaType)`

SetType sets Type field to given value.

### HasType

`func (o *PortLoa) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


