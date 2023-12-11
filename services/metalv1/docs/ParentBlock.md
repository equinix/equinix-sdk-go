# ParentBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **int32** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Netmask** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 

## Methods

### NewParentBlock

`func NewParentBlock() *ParentBlock`

NewParentBlock instantiates a new ParentBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParentBlockWithDefaults

`func NewParentBlockWithDefaults() *ParentBlock`

NewParentBlockWithDefaults instantiates a new ParentBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *ParentBlock) GetCidr() int32`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *ParentBlock) GetCidrOk() (*int32, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *ParentBlock) SetCidr(v int32)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *ParentBlock) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetHref

`func (o *ParentBlock) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ParentBlock) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ParentBlock) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ParentBlock) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetNetmask

`func (o *ParentBlock) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *ParentBlock) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *ParentBlock) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *ParentBlock) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetNetwork

`func (o *ParentBlock) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ParentBlock) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ParentBlock) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ParentBlock) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


