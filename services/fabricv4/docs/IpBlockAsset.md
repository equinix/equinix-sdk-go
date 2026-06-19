# IpBlockAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**IpBlockProductType**](IpBlockProductType.md) |  | 
**Uuid** | **string** | Unique identifier for the asset | 
**Href** | **string** | Resource URL path for the linked resource | 

## Methods

### NewIpBlockAsset

`func NewIpBlockAsset(type_ IpBlockProductType, uuid string, href string, ) *IpBlockAsset`

NewIpBlockAsset instantiates a new IpBlockAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpBlockAssetWithDefaults

`func NewIpBlockAssetWithDefaults() *IpBlockAsset`

NewIpBlockAssetWithDefaults instantiates a new IpBlockAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *IpBlockAsset) GetType() IpBlockProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpBlockAsset) GetTypeOk() (*IpBlockProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpBlockAsset) SetType(v IpBlockProductType)`

SetType sets Type field to given value.


### GetUuid

`func (o *IpBlockAsset) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *IpBlockAsset) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *IpBlockAsset) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetHref

`func (o *IpBlockAsset) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IpBlockAsset) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IpBlockAsset) SetHref(v string)`

SetHref sets Href field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


