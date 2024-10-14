# SimplifiedVirtualDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | url to entity | [optional] 
**Uuid** | Pointer to **string** | Network Edge assigned Virtual Device Identifier | [optional] 
**Type** | Pointer to [**SimplifiedVirtualDeviceType**](SimplifiedVirtualDeviceType.md) |  | [optional] 
**Name** | Pointer to **string** | Customer-assigned Virtual Device name | [optional] 
**Cluster** | Pointer to **string** | Virtual Device Cluster Information | [optional] 

## Methods

### NewSimplifiedVirtualDevice

`func NewSimplifiedVirtualDevice() *SimplifiedVirtualDevice`

NewSimplifiedVirtualDevice instantiates a new SimplifiedVirtualDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimplifiedVirtualDeviceWithDefaults

`func NewSimplifiedVirtualDeviceWithDefaults() *SimplifiedVirtualDevice`

NewSimplifiedVirtualDeviceWithDefaults instantiates a new SimplifiedVirtualDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *SimplifiedVirtualDevice) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SimplifiedVirtualDevice) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SimplifiedVirtualDevice) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *SimplifiedVirtualDevice) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetUuid

`func (o *SimplifiedVirtualDevice) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SimplifiedVirtualDevice) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SimplifiedVirtualDevice) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SimplifiedVirtualDevice) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetType

`func (o *SimplifiedVirtualDevice) GetType() SimplifiedVirtualDeviceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimplifiedVirtualDevice) GetTypeOk() (*SimplifiedVirtualDeviceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimplifiedVirtualDevice) SetType(v SimplifiedVirtualDeviceType)`

SetType sets Type field to given value.

### HasType

`func (o *SimplifiedVirtualDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *SimplifiedVirtualDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SimplifiedVirtualDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SimplifiedVirtualDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SimplifiedVirtualDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCluster

`func (o *SimplifiedVirtualDevice) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *SimplifiedVirtualDevice) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *SimplifiedVirtualDevice) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *SimplifiedVirtualDevice) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


