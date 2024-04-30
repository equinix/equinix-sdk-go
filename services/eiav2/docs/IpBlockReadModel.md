# IpBlockReadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Ip block URI | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Type** | [**IpBlockReadModelType**](IpBlockReadModelType.md) |  | 

## Methods

### NewIpBlockReadModel

`func NewIpBlockReadModel(type_ IpBlockReadModelType, ) *IpBlockReadModel`

NewIpBlockReadModel instantiates a new IpBlockReadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockReadModelWithDefaults

`func NewIpBlockReadModelWithDefaults() *IpBlockReadModel`

NewIpBlockReadModelWithDefaults instantiates a new IpBlockReadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *IpBlockReadModel) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IpBlockReadModel) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IpBlockReadModel) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *IpBlockReadModel) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *IpBlockReadModel) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlockReadModel) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlockReadModel) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *IpBlockReadModel) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *IpBlockReadModel) GetType() IpBlockReadModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpBlockReadModel) GetTypeOk() (*IpBlockReadModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpBlockReadModel) SetType(v IpBlockReadModelType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


